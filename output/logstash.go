package output

import (
	"bufio"
	"encoding/json"
	"log"
	"net"
	"time"

	"janmg.com/blob-to-queue/common"
	"janmg.com/blob-to-queue/format"
)

// LogstashWorker sends JSON events (one per line) to a Logstash TCP input.
// It reads the target address from Config.Stdout.Connectionstring or defaults to localhost:5000.
func LogstashWorker(queue <-chan format.Flatevent) {
	log.Println("Logstash worker initializing...")
	config := common.ConfigHandler()

	addr := config.Stdout.Connectionstring
	if addr == "" {
		addr = "localhost:5000"
	}

	var conn net.Conn
	var writer *bufio.Writer

	dial := func() error {
		if conn != nil {
			conn.Close()
		}
		c, err := net.DialTimeout("tcp", addr, 5*time.Second)
		if err != nil {
			return err
		}
		conn = c
		writer = bufio.NewWriter(conn)
		log.Printf("Connected to Logstash at %s", addr)
		return nil
	}

	// initial connect with backoff
	for i := 0; ; i++ {
		if err := dial(); err != nil {
			wait := time.Second * time.Duration(1<<uint(min(i, 6)))
			log.Printf("Logstash connect failed: %v; retrying in %s", err, wait)
			time.Sleep(wait)
			continue
		}
		break
	}

	sent := 0
	for evt := range queue {
		// Marshal event to JSON
		data, err := json.Marshal(evt)
		if err != nil {
			log.Printf("Logstash: error marshaling event: %v", err)
			continue
		}

		// Ensure connection is available
		if conn == nil {
			// try to reconnect
			if err := dial(); err != nil {
				log.Printf("Logstash: reconnect failed: %v", err)
				time.Sleep(2 * time.Second)
				continue
			}
		}

		// Write JSON line
		_, err = writer.Write(append(data, '\n'))
		if err != nil {
			log.Printf("Logstash: write error: %v; attempting reconnect", err)
			conn.Close()
			conn = nil
			writer = nil
			// try reconnect next loop
			continue
		}

		// Flush (could be optimized)
		if err := writer.Flush(); err != nil {
			log.Printf("Logstash: flush error: %v; closing connection", err)
			conn.Close()
			conn = nil
			writer = nil
			continue
		}

		sent++
		if sent%100 == 0 {
			log.Printf("Logstash: sent %d events", sent)
		}
	}

	// Clean up when queue is closed
	if conn != nil {
		conn.Close()
	}
	log.Printf("Logstash worker finished, total sent: %d", sent)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
