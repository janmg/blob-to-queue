package output

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/go-zeromq/zmq4"
	"janmg.com/blob-to-queue/common"
	"janmg.com/blob-to-queue/format"
)

var (
	zmqSocket  zmq4.Socket
	zmqOnce    sync.Once
	zmqInitErr error
)

// initZeroMQ initializes the ZeroMQ socket once
func initZeroMQ() {
	zmqOnce.Do(func() {
		config := common.ConfigHandler()

		// Create a PUB socket (can be changed to PUSH, REQ, etc. based on needs)
		zmqSocket = zmq4.NewPub(context.Background())

		// Default endpoint - should be configurable
		endpoint := "tcp://localhost:5555"
		if config.Format != "" {
			// You can add zeromq endpoint to config if needed
			// endpoint = config.ZeroMQEndpoint
		}

		err := zmqSocket.Listen(endpoint)
		if err != nil {
			zmqInitErr = err
			log.Printf("Failed to initialize ZeroMQ socket: %v", err)
			return
		}

		log.Printf("ZeroMQ publisher listening on %s", endpoint)
	})
}

func SendZERO(nsg format.Flatevent) {
	initZeroMQ()

	if zmqInitErr != nil {
		fmt.Printf("ZeroMQ not initialized: %v\n", zmqInitErr)
		return
	}

	// Convert the event to JSON
	nsgjson, err := json.Marshal(nsg)
	if err != nil {
		fmt.Printf("Failed to marshal event to JSON: %v\n", err)
		return
	}

	// Create a ZeroMQ message
	msg := zmq4.NewMsgFrom(nsgjson)

	// Send the message
	err = zmqSocket.Send(msg)
	if err != nil {
		fmt.Printf("Failed to send ZeroMQ message: %v\n", err)
		return
	}

	fmt.Println("ZeroMQ message sent")
}

// CloseZeroMQ should be called on shutdown
func CloseZeroMQ() error {
	if zmqSocket != nil {
		return zmqSocket.Close()
	}
	return nil
}
