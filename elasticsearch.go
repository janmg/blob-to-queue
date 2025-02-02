package main

import (
	"bytes"
	"crypto/tls"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	elasticsearch "github.com/elastic/go-elasticsearch/v8"
)

func sendElasticsearch(nsg Flatevent) {
	// convert this sender to worker that can reuse elasticsearch connections, would this require a channel ?
	// https://gobyexample.com/worker-pools
	// https://github.com/elastic/go-elasticsearch/issues/123

	cert, _ := os.ReadFile("elastic_ca.crt")
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
	Error(err)

	if err != nil {
		log.Printf("Error creating the client: %s", err)
	} else {
		//log.Println(es.Info())
		data := format_json(nsg)
		_, err := es.Index("nsgflowlog", bytes.NewReader(data))
		Error(err)
		//_, err = Client.Bulk(...).
		//fmt.Println(res)
		// 2024/12/02 20:20:31 net/http: timeout awaiting response headers
	}
}
