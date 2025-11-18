package output

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	"janmg.com/blob-to-queue/common"
	"janmg.com/blob-to-queue/format"
)

var (
	kvClient  *clientv3.Client
	kvOnce    sync.Once
	kvInitErr error
)

// initKeyval initializes the etcd client once
func initKeyval() {
	kvOnce.Do(func() {
		config := common.ConfigHandler()

		// Default etcd endpoints - should be configurable
		endpoints := []string{"localhost:2379"}
		if config.Format != "" {
			// You can add keyval endpoints to config if needed
			// endpoints = config.KeyvalEndpoints
		}

		// Create etcd client
		cli, err := clientv3.New(clientv3.Config{
			Endpoints:   endpoints,
			DialTimeout: 5 * time.Second,
		})
		if err != nil {
			kvInitErr = fmt.Errorf("failed to connect to etcd: %w", err)
			log.Printf("Keyval connection error: %v", kvInitErr)
			return
		}
		kvClient = cli

		log.Printf("Keyval (etcd) connected to %v", endpoints)
	})
}

func SendKeyval(nsg format.Flatevent) {
	initKeyval()

	if kvInitErr != nil {
		fmt.Printf("Keyval not initialized: %v\n", kvInitErr)
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

	// Create a key based on timestamp and source IP for uniqueness
	// Format: /flowlogs/{category}/{timestamp}_{srcip}_{dstip}
	key := fmt.Sprintf("/flowlogs/%s/%s_%s_%s",
		nsg.Category,
		nsg.Unixtime,
		nsg.SrcIP,
		nsg.DstIP,
	)

	// Store the event in etcd
	_, err = kvClient.Put(ctx, key, string(nsgjson))
	if err != nil {
		fmt.Printf("Failed to put key-value: %v\n", err)
		return
	}

	fmt.Printf("Keyval stored: %s\n", key)
}

// CloseKeyval should be called on shutdown
func CloseKeyval() error {
	if kvClient != nil {
		return kvClient.Close()
	}
	return nil
}
