package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func sendKafka(nsg flatevent) {
	fmt.Println("Kafka sending")
	topic := "insights-logs-networksecuritygroupflowevent"
	partition := 0

	kfk, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	handleError(err)

	nsgjson, err := json.Marshal(nsg)
	handleError(err)
	fmt.Println(string(nsgjson))

	_, err = kfk.WriteMessages(kafka.Message{
		Value: nsgjson,
	},
	)
	handleError(err)

	err = kfk.Close()
	handleError(err)
}
