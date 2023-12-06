package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
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

// struct for (temporary) storing event in flat format as json compatbile struct, to easily transform the event to csv or ecs
type flatevent struct {
	Time          time.Time `json:"time"`
	SystemID      string    `json:"systemId"`
	MACAdress     string    `json:"macAddress"`
	Category      string    `json:"category"`
	ResourceID    string    `json:"resourceId"`
	OperationName string    `json:"operationName"`
	Version       int       `json:"Version"`
	Rule          string    `json:"rule"`
	Mac           string    `json:"mac"`
	Unixtime      string    `json:"unixtime"`
	SrcIP         string    `json:"srcip"`
	DstIP         string    `json:"dstip"`
	SrcPort       string    `json:"srcport"`
	DstPort       string    `json:"dstport"`
	Proto         string    `json:"proto"`
	Direction     string    `json:"direction"`
	Action        string    `json:"action"`
	State         string    `json:"state"`
	SrcPackets    int       `json:"srcpackets"`
	SrcBytes      int       `json:"srcbytes"`
	DstPackets    int       `json:"dstpackets"`
	DstBytes      int       `json:"dstbytes"`
}

// struct for outputting in Logstash ECS format
type ecsevent struct {
	/*      Ecs struct {
	                Version string `json:"version"`
	        } string `json:"ecs"`
	        ...

	        // https://www.elastic.co/guide/en/ecs/current/ecs-field-reference.html
	        ecs.set("ecs.version", "1.0.0")
	        ecs.set("@timestamp", old.timestamp)
	        ecs.set("cloud.provider", "azure")
	        ecs.set("cloud.account.id", old.get("[subscription]")
	        ecs.set("cloud.project.id", old.get("[environment]")
	        ecs.set("file.name", old.get("[filename]")
	        ecs.set("event.category", "network")
	        if old.get("[decision]") == "D"
	            ecs.set("event.type", "denied")
	        else
	            ecs.set("event.type", "allowed")
	        end
	        ecs.set("event.action", "")
	        ecs.set("rule.ruleset", old.get("[nsg]")
	        ecs.set("rule.name", old.get("[rule]")
	        ecs.set("trace.id", old.get("[protocol]")+"/"+old.get("[src_ip]")+":"+old.get("[src_port]")+"-"+old.get("[dst_ip]")+":"+old.get("[dst_port]")
	        # requires logic to match sockets and flip src/dst for outgoing.
	        ecs.set("host.mac", old.get("[mac]")
	        ecs.set("source.ip", old.get("[src_ip]")
	        ecs.set("source.port", old.get("[src_port]")
	        ecs.set("source.bytes", old.get("[srcbytes]")
	        ecs.set("source.packets", old.get("[src_pack]")
	        ecs.set("destination.ip", old.get("[dst_ip]")
	        ecs.set("destination.port", old.get("[dst_port]")
	        ecs.set("destination.bytes", old.get("[dst_bytes]")
	        ecs.set("destination.packets", old.get("[dst_packets]")
	        if old.get("[protocol]") = "U"
	            ecs.set("network.transport", "udp")
	        else
	            ecs.set("network.transport", "tcp")
	        end
	        if old.get("[decision]") == "I"
	            ecs.set("network.direction", "incoming")
	        else
	            ecs.set("network.direction", "outgoing")
	        end
	        ecs.set("network.bytes", old.get("[src_bytes]")+old.get("[dst_bytes]")
	        ecs.set("network.packets", old.get("[src_packets]")+old.get("[dst_packets]")
	        return ecs
	*/
}

func nsgflowlog(queue chan<- flatevent, flowlogs []byte, blobname string) {
	count := 0

	var nsgflowlogs NSGFlowLogs
	json.Unmarshal(flowlogs, &nsgflowlogs)
	for _, elements := range nsgflowlogs.Records {
		var event flatevent
		event.Time = elements.Time
		event.Version = elements.Properties.Version
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
					fmt.Println(tuples)
					queue <- event
					// do some wait event if channel is full?
					count++
				}
			}
		}
	}
	fmt.Println(count)
}
func addtuples(event flatevent, nsgflow string) flatevent {
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
	// nice moment to keep some socket statistics
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
