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
	queue := make(chan format.Flatevent, config.Qsize)
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

	// Start output workers - Note: Currently only one output can consume from the queue
	// For multiple outputs, need to implement a fan-out pattern with separate channels
	// TODO: Should prefix the queue with the source format (json, json_lines, lines?)
	// Note: flatevents work from both nsgflowlogs and vnetflowlogs (both are implemented)
	fmt.Printf("Config outputs: %v\n", config.Output)
	workersStarted := 0
	for _, out := range config.Output {
		fmt.Printf("Processing output: %s\n", out)
		switch out {
		case "elasticsearch":
			fmt.Println("Launching Elasticsearch worker goroutine...")
			go output.ElasticsearchWorker(queue)
			workersStarted++
		case "kafka":
			// go output.KafkaWorker(queue)
		case "eventhub":
			// go output.EventHubWorker(queue)
		case "mqtt":
			// output.MQTTWorker(queue)
		case "ampq":
			// output.AMPQWorker(queue)
		case "zeromq":
			// output.ZEROWorker(queue)
		case "keyval":
			// output.KeyvalWorker(queue)
		case "redis":
			// output.RedisWorker(queue)
		case "fluent":
			// output.FluentWorker(queue)
		case "fluxdb":
			// output.FluxWorker(queue)
		case "file":
			// output.AppendFileWorker(queue)
		case "stdout":
			// output.StdoutWorker(queue)
		case "summary":
			// output.StatisticsWorker(queue)
		}
	}

	if workersStarted == 0 {
		fmt.Println("WARNING: No workers started! Check your config.Output settings")
	} else {
		fmt.Printf("Started %d worker(s)\n", workersStarted)
	}

	// Keep main running
	select {}
}
