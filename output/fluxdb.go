package output

import (
	"fmt"
	"log"
	"time"

	"github.com/influxdata/influxdb/client/v2"
	"janmg.com/blob-to-queue/format"
)

const (
	address  = "http://localhost:8086"
	dbname   = "square_holes"
	username = "bubba"
	password = "bumblebeetuna"
)

func SendFlux(nsg format.Flatevent) {
	fmt.Println("FluxDB sending not yet implemented")

	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     address,
		Username: username,
		Password: password,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	// Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  dbname,
		Precision: "s",
	})
	if err != nil {
		log.Fatal(err)
	}

	// Create a point and add to batch
	tags := map[string]string{"cpu": "cpu-total"}
	fields := map[string]interface{}{
		"idle":   10.1,
		"system": 53.3,
		"user":   46.6,
	}

	pt, err := client.NewPoint("cpu_usage", tags, fields, time.Now())
	if err != nil {
		log.Fatal(err)
	}
	bp.AddPoint(pt)

	// Write the batch
	if err := c.Write(bp); err != nil {
		log.Fatal(err)
	}

	// Close client resources
	if err := c.Close(); err != nil {
		log.Fatal(err)
	}
}
