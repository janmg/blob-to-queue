package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"janmg.com/blob-to-queue/common"
	"janmg.com/blob-to-queue/format"
	"janmg.com/blob-to-queue/input"
	"janmg.com/blob-to-queue/output"
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
var config common.Config

func main() {
	fmt.Printf("blob-to-queue v1.0-dev\n")
	config = common.ConfigHandler()
	//configPrint(blob)

	// Shutdown handler, if stop signal comes, process last messages in the queue, but stop inflow
	//fetchmore := true
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Read flatevents from the blobstorage and add them to the queue
	queue := make(chan format.Flatevent, 10000)
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
	// TODO: Should this be done for multiple storage accounts? or multiple directories? Or some logic between them? I think each directory should have each own worker.
	// e.g. /SUBSCRIPTIONS/F5DD6E2D-1F42-4F54-B3BD-DBF595138C59/RESOURCEGROUPS/VM/PROVIDERS/MICROSOFT.NETWORK/OCTOBER-NSG
	go input.Blobworker(queue)

	// Read from the queue and decide what to do with the output
	send(queue)
}

func send(queue <-chan format.Flatevent) {
	// Read from the queue and decide what to do with the output
	for {
		nsg := <-queue
		//fmt.Println(format("csv", nsg))
		// TODO: filter?

		for _, out := range config.Output {
			switch out {
			case "elasticsearch":
				output.SendElasticsearch(nsg)
			case "eventhub":
				output.SendAzure(nsg)
			case "kafka":
				output.SendKafka(nsg)
			case "mqtt":
				output.SendMQTT(nsg)
			case "ampq":
				output.SendAMPQ(nsg)
			case "zeromq":
				output.SendZERO(nsg)
			case "keyval":
				output.SendKeyval(nsg)
			case "redis":
				output.SendRedis(nsg)
			case "file":
				output.AppendFile(nsg)
			case "stdout":
				output.Stdout(nsg)
			case "summary":
				output.Statistics(nsg)
			}
		}
	}
}
