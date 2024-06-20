package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
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
filter finegrained with if statements
format events
send to stream
printing out to stdout or logfile
*/

func Error(err error) {
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(2)
	}
}

func Warning(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func main() {
	fmt.Printf("blob-to-queue v1.0-dev\n")
	//var blob IBlob = configHandler()
	//blob.Print()

	// Shutdown handler, if stop signal comes, process last messages in the queue, but stop inflow
	fetchmore := true
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Read flatevents from the blobstorage and add them to the queue
	queue := make(chan flatevent, 10000)
	defer close(queue)

	go func() {
		sig := <-sigs
		fmt.Println(sig)
		fetchmore = false
		fmt.Println("stopping after processing the last events, now in the queue: ", len(queue))
		for len(queue) > 0 {
			time.Sleep(10 * time.Second)
			fmt.Println("still in the queue: ", len(queue))
		}
		os.Exit(99)
	}()

	for {
		// Read flatevents from the blobstorage and add them to the queue
		if fetchmore {
			queue := make(chan flatevent, 10000)
			go blobworker(queue)
		}
		// Read from the queue and decide what to do with the output
		send(queue)
	}
}

func send(queue <-chan flatevent) {
	// TODO: Create break handler to finish the queue
	for {
		nsg := <-queue
		// TODO: filter?
		// TODO hardcoded output, replace with configHandler info
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
		case "zeromq":
			sendZERO(nsg)
		case "file":
			appendFile(nsg)
		case "stdout":
			stdout(nsg)
		case "summary":
			statistics(nsg)
		}
	}
}
