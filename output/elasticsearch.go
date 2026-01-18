package output

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"log"
	"net"
	"net/http"
	"os"
	"sync/atomic"
	"time"

	elasticsearch "github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"janmg.com/blob-to-queue/common"
	"janmg.com/blob-to-queue/format"
)

var (
	esClient      *elasticsearch.Client
	esBulkIndexer esutil.BulkIndexer
	esInitialized bool
	indexedCount  uint64
)

// initElasticsearch initializes the Elasticsearch client and bulk indexer
func initElasticsearch() error {
	if esInitialized {
		return nil
	}

	// Prefer CA file configured in blob-to-queue.yaml -> elasticsearch.details[0]
	config := common.ConfigHandler()
	var cert []byte
	indexname := "nsgflowlog"

	if len(config.Elasticsearch.Details) > 0 && config.Elasticsearch.Details[0] != "" {
		caFile := config.Elasticsearch.Details[0]
		c, err := os.ReadFile(caFile)
		if err != nil {
			log.Printf("Warning: Could not read CA file %s: %v; falling back to elastic_ca.crt", caFile, err)
			cert, _ = os.ReadFile("elastic_ca.crt")
		} else {
			cert = c
		}
	} else {
		cert, _ = os.ReadFile("elastic_ca.crt")
	}

	if len(config.Elasticsearch.Details) > 1 && config.Elasticsearch.Details[1] != "" {
		indexname = config.Elasticsearch.Details[1]
	}

	// TODO: Make configurable via config.elasticsearch.connection / config.elasticsearch.token
	cfg := elasticsearch.Config{
		Addresses: []string{config.Elasticsearch.Connection},
		APIKey:    config.Elasticsearch.Token,
		CACert:    cert, // failed to verify certificate: x509: certificate signed by unknown authority (possibly because of "crypto/rsa: verification error" while trying to verify candidate authority certificate "Elasticsearch security auto-configuration HTTP CA")
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: 30 * time.Second,
			DialContext:           (&net.Dialer{Timeout: 5 * time.Second}).DialContext,
			TLSClientConfig: &tls.Config{
				MinVersion: tls.VersionTLS12,
			},
		},
	}

	var err error
	esClient, err = elasticsearch.NewClient(cfg)
	common.Warning(err)

	res, err := esClient.Info()
	common.Warning(err)
	defer res.Body.Close()

	bi, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Index:         indexname,
		Client:        esClient,
		NumWorkers:    4,                // Number of concurrent workers
		FlushBytes:    1e+7,             // Flush threshold in bytes (10MB)
		FlushInterval: 10 * time.Second, // Flush interval
		OnError: func(ctx context.Context, err error) {
			log.Printf("Bulk indexer error: %v", err)
		},
		OnFlushStart: func(ctx context.Context) context.Context {
			log.Println("Bulk indexer flush started")
			return ctx
		},
		OnFlushEnd: func(ctx context.Context) {
			log.Println("Bulk indexer flush completed")
		},
	})
	common.Warning(err)

	esBulkIndexer = bi
	esInitialized = true

	log.Println("Elasticsearch BulkIndexer initialized")
	return nil
}

// ElasticsearchWorker reads from the queue and bulk indexes to Elasticsearch
func ElasticsearchWorker(queue <-chan format.Flatevent) {
	// Print immediately to prove goroutine started
	log.Println("Elasticsearch worker initializing...")
	// Ensure Elasticsearch client & bulk indexer are initialized
	if err := initElasticsearch(); err != nil {
		log.Fatalf("Failed to initialize Elasticsearch: %v", err)
	}

	eventCount := 0
	// Buffer up to 1000 events and flush as a batch. Also flush after a short timeout
	const maxBatch = 1000
	flushTimeout := 500 * time.Millisecond
	batch := make([]format.Flatevent, 0, maxBatch)

outer:
	for {
		// Block waiting for the first event of a batch
		ev, ok := <-queue
		if !ok {
			// channel closed, flush any remaining and exit
			break outer
		}
		batch = append(batch, ev)

		// Try to collect up to maxBatch events; stop collecting on timeout or channel close
	collectLoop:
		for len(batch) < maxBatch {
			select {
			case ev, ok := <-queue:
				if !ok {
					// channel closed; proceed to flush remaining
					break collectLoop
				}
				batch = append(batch, ev)
			case <-time.After(flushTimeout):
				// no more events within timeout, flush what we have
				break collectLoop
			}
		}

		// Send batch to bulk indexer
		// Good moment to check if bulkindexer is still healthy
		stats := esBulkIndexer.Stats()
		log.Println(stats)
		for _, event := range batch {
			eventCount++
			if eventCount%1000 == 0 {
				log.Printf("Elasticsearch worker received event #%d from queue", eventCount)
			}

			data, err := json.Marshal(event)
			if err != nil {
				log.Printf("Error marshaling event: %v", err)
				continue
			}

			if err := esBulkIndexer.Add(
				context.Background(),
				esutil.BulkIndexerItem{
					Action: "index",
					Body:   bytes.NewReader(data),
					OnSuccess: func(ctx context.Context, item esutil.BulkIndexerItem, res esutil.BulkIndexerResponseItem) {
						atomic.AddUint64(&indexedCount, 1)
						count := atomic.LoadUint64(&indexedCount)
						if count%100 == 0 {
							log.Printf("Successfully indexed %d documents to Elasticsearch", count)
						}
					},
					OnFailure: func(ctx context.Context, item esutil.BulkIndexerItem, res esutil.BulkIndexerResponseItem, err error) {
						if err != nil {
							log.Printf("Error indexing document: %v", err)
						} else {
							log.Printf("Error indexing document: %s: %s", res.Error.Type, res.Error.Reason)
						}
					},
				},
			); err != nil {
				log.Printf("Error adding document to bulk indexer: %v", err)
			}
		}

		// reset batch
		batch = batch[:0]
	}

	// If channel closed but batch still has items, flush them
	if len(batch) > 0 {
		for _, event := range batch {
			eventCount++
			if eventCount%1000 == 0 {
				log.Printf("Elasticsearch worker received event #%d from queue", eventCount)
			}
			data, err := json.Marshal(event)
			if err != nil {
				log.Printf("Error marshaling event: %v", err)
				continue
			}
			if err := esBulkIndexer.Add(context.Background(), esutil.BulkIndexerItem{Action: "index", Body: bytes.NewReader(data)}); err != nil {
				log.Printf("Error adding document to bulk indexer: %v", err)
			}
		}
	}

	// Close bulk indexer when channel is closed
	log.Println("Queue closed, closing bulk indexer...")
	if err := esBulkIndexer.Close(context.Background()); err != nil {
		log.Printf("Error closing bulk indexer: %v", err)
	}
	log.Printf("Elasticsearch worker finished. Total received: %d, Total indexed: %d", eventCount, atomic.LoadUint64(&indexedCount))
}
