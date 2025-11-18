package output

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"janmg.com/blob-to-queue/common"
	"janmg.com/blob-to-queue/format"
)

var (
	amqpConn    *amqp.Connection
	amqpChannel *amqp.Channel
	amqpOnce    sync.Once
	amqpInitErr error
)

// initAMQP initializes the AMQP connection and channel once
func initAMQP() {
	amqpOnce.Do(func() {
		config := common.ConfigHandler()

		// Default AMQP URL - should be configurable
		amqpURL := "amqp://guest:guest@localhost:5672/"
		if config.Format != "" {
			// You can add amqp URL to config if needed
			// amqpURL = config.AMQPURL
		}

		// Connect to RabbitMQ
		conn, err := amqp.Dial(amqpURL)
		if err != nil {
			amqpInitErr = fmt.Errorf("failed to connect to AMQP: %w", err)
			log.Printf("AMQP connection error: %v", amqpInitErr)
			return
		}
		amqpConn = conn

		// Create a channel
		ch, err := conn.Channel()
		if err != nil {
			amqpInitErr = fmt.Errorf("failed to open AMQP channel: %w", err)
			log.Printf("AMQP channel error: %v", amqpInitErr)
			return
		}
		amqpChannel = ch

		// Declare exchange (topic exchange for flexible routing)
		exchangeName := "flowlogs"
		err = ch.ExchangeDeclare(
			exchangeName, // name
			"topic",      // type
			true,         // durable
			false,        // auto-deleted
			false,        // internal
			false,        // no-wait
			nil,          // arguments
		)
		if err != nil {
			amqpInitErr = fmt.Errorf("failed to declare exchange: %w", err)
			log.Printf("AMQP exchange error: %v", amqpInitErr)
			return
		}

		log.Printf("AMQP connected to %s, exchange: %s", amqpURL, exchangeName)
	})
}

func SendAMPQ(nsg format.Flatevent) {
	initAMQP()

	if amqpInitErr != nil {
		fmt.Printf("AMQP not initialized: %v\n", amqpInitErr)
		return
	}

	// Convert the event to JSON
	nsgjson, err := json.Marshal(nsg)
	if err != nil {
		fmt.Printf("Failed to marshal event to JSON: %v\n", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Routing key based on resource type (customize as needed)
	routingKey := "flowlog.nsg"
	if nsg.Category == "NetworkSecurityGroupFlowEvent" {
		routingKey = "flowlog.nsg"
	} else if nsg.Category == "VirtualNetworkFlowEvent" {
		routingKey = "flowlog.vnet"
	}

	// Publish message
	err = amqpChannel.PublishWithContext(
		ctx,
		"flowlogs", // exchange
		routingKey, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType:  "application/json",
			Body:         nsgjson,
			DeliveryMode: amqp.Persistent, // Make messages persistent
			Timestamp:    time.Now(),
		},
	)
	if err != nil {
		fmt.Printf("Failed to publish AMQP message: %v\n", err)
		return
	}

	fmt.Println("AMQP message sent")
}

// CloseAMQP should be called on shutdown
func CloseAMQP() error {
	if amqpChannel != nil {
		if err := amqpChannel.Close(); err != nil {
			return err
		}
	}
	if amqpConn != nil {
		return amqpConn.Close()
	}
	return nil
}
