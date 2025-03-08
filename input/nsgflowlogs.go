package input

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"janmg.com/blob-to-queue/format"
)

// structs for extracting nsgflowlogs
type flows struct {
	Mac        string   `json:"mac"`
	FlowTuples []string `json:"flowTuples"`
}
type properties struct {
	Version int `json:"Version"`
	Flows   []struct {
		Rule  string  `json:"rule"`
		Flows []flows `json:"flows"`
	} `json:"flows"`
}
type NSGFlowLogs struct {
	Records []struct {
		Time          time.Time  `json:"time"`
		SystemID      string     `json:"systemId"`
		MacAddress    string     `json:"macAddress"`
		Category      string     `json:"category"`
		ResourceID    string     `json:"resourceId"`
		OperationName string     `json:"operationName"`
		Properties    properties `json:"properties"`
	} `json:"records"`
}

func nsgflowlog(queue chan<- format.Flatevent, flowlogs []byte, blobname string) {
	count := 0
	/* ChatGPT says the worker should be a concurrency pool
	var wg sync.WaitGroup
	numWorkers := 4
	queue := make(chan Flatevent, 100)

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for event := range queue {
				processEvent(event) // Replace with actual function
			}
		}()
	}
	*/

	var nsgflowlogs NSGFlowLogs
	json.Unmarshal(flowlogs, &nsgflowlogs)

	/*
	   ChatGPT thinks that Unmarshal is slow and should use a Decoder
	   decoder := json.NewDecoder(bytes.NewReader(flowlogs))
	   if err := decoder.Decode(&nsgflowlogs); err != nil {
	       fmt.Println("JSON Decode error:", err)
	   }
	*/

	for _, elements := range nsgflowlogs.Records {
		var event format.Flatevent
		event.Time = elements.Time
		event.SystemID = elements.SystemID
		event.MACAdress = elements.MacAddress
		event.Category = elements.Category
		event.ResourceID = elements.ResourceID
		event.OperationName = elements.OperationName
		event.Version = elements.Properties.Version
		for _, flows := range elements.Properties.Flows {
			event.Rule = flows.Rule
			for _, flow := range flows.Flows {
				event.Mac = flow.Mac
				for _, tuples := range flow.FlowTuples {
					event = addtuples(event, tuples)
					queue <- event
					// TODO: do some wait event if channel is full?
					count++
				}
			}
		}
	}
	fmt.Println("nsgflowlog count: ", count, " in file", blobname)
}
func addtuples(event format.Flatevent, nsgflow string) format.Flatevent {
	tups := strings.Split(nsgflow, ",")
	event.Unixtime = tups[0]
	event.SrcIP = tups[1]
	event.DstIP = tups[2]
	event.SrcPort = tups[3]
	event.DstPort = tups[4]
	event.Proto = tups[5]
	event.Direction = tups[6]
	event.Action = tups[7]
	if event.Version == 2 {
		event.State = tups[8]
		event.SrcPackets = zeroIfEmpty(tups[9])
		event.SrcBytes = zeroIfEmpty(tups[10])
		event.DstPackets = zeroIfEmpty(tups[11])
		event.DstBytes = zeroIfEmpty(tups[12])
	}
	// TODO nice moment to keep some socket statistics? now doing some in stats and in ecs?
	// socket(src_ip-src_port+dst_port-dst_port, begintime, src_packets, src_bytes, dst_packets, dst_bytes)
	return event
}

func zeroIfEmpty(s string) int {
	if len(s) == 0 {
		return 0
	}
	n, err := strconv.Atoi(s)
	if err == nil {
		return n
	}
	return 0
}
