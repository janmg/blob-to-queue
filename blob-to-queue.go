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
filter finegrained with if statements,=
format events
send to stream
*/

func handleError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func main() {
	fmt.Printf("blob-to-queue v1.0-dev\n")
	//var blob IBlob = configHandler()
	//blob.Print()

	queue := make(chan flatevent, 10000)
	go blobworker(queue)
	send(queue)

	defer close(queue)
}

func send(queue <-chan flatevent) {
	nsg := <-queue
	// filter?
	//TODO hardcoded output, replace with configHandler info
	output := "stdout"
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
