package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/segmentio/kafka-go"
	// https://github.com/Shopify/sarama
	// https://pkg.go.dev/github.com/twmb/kafka-go/pkg/kgo
)

func sendKafka(nsg Flatevent) {
	fmt.Println("Kafka sending")
	topic := "insights-logs-networksecuritygroupflowevent"
	partition := 0

	// TODO: use config
	kfk, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	Error(err)

	nsgjson, err := json.Marshal(nsg)
	Error(err)
	fmt.Println(string(nsgjson))

	_, err = kfk.WriteMessages(kafka.Message{
		Value: nsgjson,
	},
	)
	Error(err)

	err = kfk.Close()
	Error(err)
}
