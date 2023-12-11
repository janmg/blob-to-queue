package main

import (
	"fmt"
	"log"
	// https://github.com/Shopify/sarama
	// https://pkg.go.dev/github.com/twmb/kafka-go/pkg/kgo
	// https://github.com/streadway/amqp
	// https://github.com/rabbitmq/amqp091-go
)

/*
config viper config reload on fsnotify
blob listing
path filtering
keeping registry?
reading full blobs
reading partial blobs blocks, tracking start and end
detecting json, nsgflowlogs
tracking of time and read sequential file list
printing out to stdout or logfile
format events
send to stream
*/

func handleError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

var lookup []output

func main() {
	fmt.Printf("NSGFLOWLOG\n")
	//var blob IBlob = configHandler()
	//blob.Print()

	queue := make(chan flatevent, 10000)
	go blobworker(queue)
	send(queue)
}

func send(queue <-chan flatevent) {
	// loop through lookup, prep desired format
	//csv:=
	//ecs:=
	nsg := <-queue
	output := "eventhub"
	switch output {
	case "eventhub":
		sendAzure(nsg)
	case "kafka":
		sendKafka(nsg)
	case "mqtt":
		sendMQTT(nsg)
	case "ampq":
		sendAMPQ(nsg)
	case "file":
		appendFile(nsg)
	case "stdout":
		stdout(nsg)
	case "summary":
		statistics(nsg)
	}
}
