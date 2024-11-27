package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
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
var config Config

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
	config = configHandler()
	//configPrint(blob)

	// Shutdown handler, if stop signal comes, process last messages in the queue, but stop inflow
	//fetchmore := true
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Read flatevents from the blobstorage and add them to the queue
	queue := make(chan Flatevent, 10000)
	defer close(queue)

	go func() {
		sig := <-sigs
		fmt.Println(sig)
		//fetchmore = false
		fmt.Println("stopping after processing the last events, now in the queue: ", len(queue))
		for len(queue) > 0 {
			time.Sleep(10 * time.Second)
			fmt.Println("still in the queue: ", len(queue))
		}
		os.Exit(99)
	}()

	// Read flatevents from the blobstorage and add them to the queue
	go blobworker(queue)

	// Read from the queue and decide what to do with the output
	send(queue)
}

func send(queue <-chan Flatevent) {
	// Read from the queue and decide what to do with the output
	for {
		nsg := <-queue
		//fmt.Println(format("csv", nsg))
		// TODO: filter?

		for index, output := range config.Output {
			fmt.Println(index, output)
			switch output {
			case "elasticsearch":
				sendElasticsearch(nsg)
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
			case "keyval":
				sendKeyval(nsg)
			case "redis":
				sendRedis(nsg)
			case "file":
				appendFile(nsg)
			case "stdout":
				stdout(nsg)
			case "summary":
				statistics(nsg)
			}
		}
	}
}
