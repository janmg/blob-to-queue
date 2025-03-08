package output

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/segmentio/kafka-go"
	"janmg.com/blob-to-queue/common"
	"janmg.com/blob-to-queue/format"
	// https://github.com/Shopify/sarama
	// https://pkg.go.dev/github.com/twmb/kafka-go/pkg/kgo
)

func SendKafka(nsg format.Flatevent) {
	fmt.Println("Kafka sending")
	topic := "insights-logs-networksecuritygroupflowevent"
	partition := 0

	// TODO: use config
	kfk, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	common.Error(err)

	nsgjson, err := json.Marshal(nsg)
	common.Error(err)
	fmt.Println(string(nsgjson))

	_, err = kfk.WriteMessages(kafka.Message{
		Value: nsgjson,
	},
	)
	common.Error(err)

	err = kfk.Close()
	common.Error(err)
}
