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

	cert, err := os.ReadFile("elastic_ca.crt")
	if err != nil {
		log.Printf("Warning: Could not read elastic_ca.crt: %v", err)
		cert = nil
	}

	// TODO: Make configurable via config
	cfg := elasticsearch.Config{
		Addresses: []string{"https://10.0.0.230:9200"},
		APIKey:    "NExt==",
		CACert:    cert,
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: 30 * time.Second,
			DialContext:           (&net.Dialer{Timeout: 5 * time.Second}).DialContext,
			TLSClientConfig: &tls.Config{
				MinVersion: tls.VersionTLS12,
			},
		},
	}

	esClient, err = elasticsearch.NewClient(cfg)
	common.Warning(err)
	// error creating Elasticsearch client

	// Test connection
	res, err := esClient.Info()
	common.Warning(err)
	// error getting Elasticsearch info

	defer res.Body.Close()
	common.Warning(err)
	// error response from Elasticsearch

	// Create bulk indexer
	bi, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Index:         "nsgflowlog",
		Client:        esClient,
		NumWorkers:    4,                // Number of concurrent workers
		FlushBytes:    5e+6,             // Flush threshold in bytes (5MB)
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
	// error creating bulk indexer

	esBulkIndexer = bi
	esInitialized = true

	log.Println("Elasticsearch BulkIndexer initialized")
	return nil
}

// ElasticsearchWorker reads from the queue and bulk indexes to Elasticsearch
func ElasticsearchWorker(queue <-chan format.Flatevent) {
	// Print immediately to prove goroutine started
	log.Println("Elasticsearch worker initializing...")

	if err := initElasticsearch(); err != nil {
		log.Fatalf("Failed to initialize Elasticsearch: %v", err)
	}

	eventCount := 0
	for event := range queue {
		eventCount++
		if eventCount%100 == 1 {
			log.Printf("Elasticsearch worker received event #%d from queue", eventCount)
		}

		// Convert event to JSON
		data, err := json.Marshal(event)
		if err != nil {
			log.Printf("Error marshaling event: %v", err)
			continue
		}

		// Add to bulk indexer
		err = esBulkIndexer.Add(
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
		)
		if err != nil {
			log.Printf("Error adding document to bulk indexer: %v", err)
		}
	}

	// Close bulk indexer when channel is closed
	log.Println("Queue closed, closing bulk indexer...")
	if err := esBulkIndexer.Close(context.Background()); err != nil {
		log.Printf("Error closing bulk indexer: %v", err)
	}
	log.Printf("Elasticsearch worker finished. Total received: %d, Total indexed: %d", eventCount, atomic.LoadUint64(&indexedCount))
}
