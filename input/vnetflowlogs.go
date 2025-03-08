package input

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"janmg.com/blob-to-queue/format"
)

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

	var vnetflowlogs VNETFlowLogs
	json.Unmarshal(flowlogs, &vnetflowlogs)
	for _, elements := range vnetflowlogs.Records {
		var event format.Flatevent
		event.Time = elements.Time
		event.MACAdress = elements.MacAddress
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
					// do some wait event if channel is full?
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
	event.Proto = tups[5] // Now an IANA protocol number
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
