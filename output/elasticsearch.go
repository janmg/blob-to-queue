package output

import (
	"bytes"
	"crypto/tls"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	elasticsearch "github.com/elastic/go-elasticsearch/v8"

	"janmg.com/blob-to-queue/common"
	"janmg.com/blob-to-queue/format"
)

func SendElasticsearch(nsg format.Flatevent) {
	// convert this sender to worker that can reuse elasticsearch connections, would this require a channel ?
	// https://gobyexample.com/worker-pools
	// https://github.com/elastic/go-elasticsearch/issues/123

	cert, _ := os.ReadFile("elastic_ca.crt")
	// TODO: Make configurable
	cfg := elasticsearch.Config{
		//CloudID: "my-cluster:dXMtZWFzdC0xLZC5pbyRjZWM2ZjI2MWE3NGJm...",
		Addresses: []string{"https://10.0.0.247:9200"},
		APIKey:    "OEQxU2lKTUJqdmgxWnpqMGxCUU46bHg0WldieDVRYXFLcF8tSlViQkRTQQ==",
		CACert:    cert,
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Millisecond,
			DialContext:           (&net.Dialer{Timeout: 5 * time.Second}).DialContext,
			TLSClientConfig: &tls.Config{
				MinVersion: tls.VersionTLS12,
			},
		},
	}
	es, err := elasticsearch.NewClient(cfg)
	common.Error(err)

	if err != nil {
		log.Printf("Error creating the client: %s", err)
	} else {
		//log.Println(es.Info())
		data := format.Format("json", nsg)
		_, err := es.Index("nsgflowlog", bytes.NewReader([]byte(data)))
		common.Error(err)
		//_, err = Client.Bulk(...).
		//fmt.Println(res)
		// 2024/12/02 20:20:31 net/http: timeout awaiting response headers

		/* ChatGPT suggestion, but this covers multiple requests at once for bulk api, this means code improvemnts in nsgflowlogs and queuing a batch
		var bulkData strings.Builder
		for _, event := range events {
			bulkData.WriteString(fmt.Sprintf(`{"index":{}}\n%s\n`, eventJSON))
		}

		req := esapi.BulkRequest{
			Index: "nsgflowlogs",
			Body:  strings.NewReader(bulkData.String()),
		}
		*/
	}
}
