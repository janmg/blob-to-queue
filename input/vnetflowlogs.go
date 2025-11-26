package input

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	. "janmg.com/blob-to-queue/common"
	"janmg.com/blob-to-queue/format"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// structs for extracting vnetflowlogs
type flowrecords struct {
	Flows []struct {
		ACLID      string `json:"aclID"`
		FlowGroups []struct {
			Rule       string   `json:"rule"`
			FlowTuples []string `json:"flowTuples"`
		} `json:"flowGroups"`
	} `json:"flows"`
}

type VNETFlowLogs struct {
	Records []struct {
		Time              time.Time   `json:"time"`
		FlowLogVersion    int         `json:"flowLogVersion"`
		FlowLogGUID       string      `json:"flowLogGUID"`
		MacAddress        string      `json:"macAddress"`
		Category          string      `json:"category"`
		FlowLogResourceID string      `json:"flowLogResourceID"`
		TargetResourceID  string      `json:"targetResourceID"`
		OperationName     string      `json:"operationName"`
		FlowRecords       flowrecords `json:"flowrecords"`
	} `json:"records"`
}

func vnetflowlog(queue chan<- format.Flatevent, flowlogs []byte, blobname string) {
	count := 0

	fmt.Printf("Processing vnetflowlog from blob: %s\n", blobname)
	fmt.Printf("Raw JSON length: %d bytes\n", len(flowlogs))

	var vnetflowlogs VNETFlowLogs
	err := json.Unmarshal(flowlogs, &vnetflowlogs)
	Warning(err)
	fmt.Printf("Successfully unmarshaled vnetflowlog, Records count: %d\n", len(vnetflowlogs.Records))

	if len(vnetflowlogs.Records) == 0 {
		fmt.Println("WARNING: vnetflowlogs.Records is empty!")
		fmt.Printf("Raw JSON (first 1000 chars): %s\n", string(flowlogs[:min(1000, len(flowlogs))]))
		return
	}

	for _, elements := range vnetflowlogs.Records {
		var event format.Flatevent
		event.Time = elements.Time
		event.MacAddress = elements.MacAddress
		event.Category = elements.Category
		event.OperationName = elements.OperationName
		for _, flows := range elements.FlowRecords.Flows {
			event.ACLID = flows.ACLID
			for _, flow := range flows.FlowGroups {
				event.Rule = flow.Rule
				for _, tuples := range flow.FlowTuples {
					event = vnettuples(event, tuples)
					//fmt.Println(tuples)
					queue <- event
					// Check if queue is over 80% capacity
					queueLen := len(queue)
					queueCap := cap(queue)
					if queueLen > int(float64(queueCap)*0.8) {
						fmt.Printf("WARNING: Queue is at %d/%d (%.1f%%), pausing to prevent overflow\n", queueLen, queueCap, float64(queueLen)/float64(queueCap)*100)
						time.Sleep(5 * time.Second)
						// TODO: use a signal to the input reader to slow down reading from blob storage
					}
					count++
				}
			}
		}
	}
	fmt.Println("vnetflowlog count: ", count)
}

func vnettuples(event format.Flatevent, vnetflow string) format.Flatevent {
	tups := strings.Split(vnetflow, ",")
	event.Unixtime = tups[0]
	event.SrcIP = tups[1]
	event.DstIP = tups[2]
	event.SrcPort = tups[3]
	event.DstPort = tups[4]
	switch tups[5] {
	// Now an IANA protocol number
	case "T":
		event.Proto = 6
	case "U":
		event.Proto = 8
	case "I":
		event.Proto = 11
	}
	event.Direction = tups[6]
	event.State = tups[7]
	event.Encryption = false
	if len(tups[7]) == 2 {
		if tups[7][1] == 'X' {
			event.Encryption = true
		}
	}
	event.SrcPackets = zeroIfEmpty(tups[8])
	event.SrcBytes = zeroIfEmpty(tups[9])
	event.DstPackets = zeroIfEmpty(tups[10])
	event.DstBytes = zeroIfEmpty(tups[11])
	// TODO nice moment to keep some socket statistics? now doing some in stats and in ecs?
	// For continuation (C) and end (E) flow states, byte and packet counts are aggregate counts from the time of the previous flow's tuple record.
	// socket(src_ip-src_port+dst_port-dst_port, begintime, src_packets, src_bytes, dst_packets, dst_bytes)
	return event
}
