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

	queue := make(chan flatevent, 1000)
	go worker(queue)
	send(queue)

}

func send(queue <-chan flatevent) {
	// loop through lookup, prep desired format
	//csv:=
	//ecs:=
	nsg := <-queue
	sendAzure(nsg)
	/*
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
				summary(nsg)
		}
	*/
}

func sendAMPQ(nsg string) {
	fmt.Println("AMPQ sending")
}

func sendMQTT(nsg string) {
	fmt.Println("MQTT sending")
}

func appendFile(nsg string) {
	fmt.Println(nsg)
}

func stdout(nsg string) {
	fmt.Println(nsg)
}

func sumout(nsg string) {
	fmt.Println(nsg)
}
