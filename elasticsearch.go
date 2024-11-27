package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	elasticsearch "github.com/elastic/go-elasticsearch/v8"
)

func sendElasticsearch(nsg Flatevent) {
	fmt.Println("Elasticsearch sending not yet implemented")
	client, _ := elasticsearch.NewClient(elasticsearch.Config{
		CloudID: "<CloudID>",
		APIKey:  "<ApiKey>",
	})
	data, _ := json.Marshal(format("json", nsg))
	client.Index("my_index", bytes.NewReader(data))
}
