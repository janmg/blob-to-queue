package format

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gocarina/gocsv"

	"janmg.com/blob-to-queue/common"
)

func Format(format string, nsg Flatevent) string {
	var formatted string
	switch format {
	case "csv":
		formatted = format_csv(nsg)
	case "json":
		formatted = string(format_json(nsg))
	case "ecs":
		formatted = format_ecs(nsg)
	}
	return formatted
}

func format_json(nsg Flatevent) []byte {
	event_json, _ := json.Marshal(nsg)
	return event_json
}

func format_csv(nsg Flatevent) string {
	// copy from gocarina the csv reader and strip out the need for an array of structs, here I only have one event in a struct at a time and don't want the header
	// maybe nice to doublequote the values?
	// https://github.com/gocarina/gocsv/blob/master/csv.go
	nsgs := []*Flatevent{}
	nsgs = append(nsgs, &nsg)
	csvContent, err := gocsv.MarshalStringWithoutHeaders(&nsgs)
	common.Error(err)
	return csvContent
}

// struct for outputting in Logstash ECS format
func format_ecs(flat Flatevent) string {
	// https://www.elastic.co/guide/en/ecs-logging/overview/master/intro.html
	// https://www.elastic.co/guide/en/ecs/current/ecs-field-reference.html
	// https://github.com/elastic/ecs/blob/main/generated/ecs/ecs_nested.yml
	// Add to a new struct and then marshall to JSON
	var event ecs
	event.Ecs.Fields.EcsVersion.Short = "1.0.0"
	//event.Base.Fields.Timestamp = flat.Time
	event.Cloud.Fields.CloudProvider.Name = "azure"
	// Enriched from configuration?
	//event.Cloud.Fields.CloudAccountID = flat.Subscription
	//event.Cloud.Fields.CloudProjectID = flat.Environment
	//event.File.Fields.FileName.Name = flat.FileName
	event.Event.Fields.EventCategory.Name = "network"
	if flat.Action == "D" {
		event.Event.Fields.EventAction.Name = "denied"
	}
	if flat.Action == "A" {
		event.Event.Fields.EventAction.Name = "allowed"
	}
	event.Event.Fields.EventAction.Name = "default"
	event.Rule.Fields.RuleRuleset.Type = "nsg"
	event.Rule.Fields.RuleName.Name = flat.Rule
	event.Destination.Fields.DestinationMac.Name = flat.Mac
	event.Source.Fields.SourceAddress.Name = flat.SrcIP
	event.Source.Fields.SourcePort.Name = flat.SrcPort
	event.Source.Fields.SourceBytes.Name = fmt.Sprint(flat.SrcBytes)
	event.Source.Fields.SourcePackets.Name = fmt.Sprint(flat.SrcPackets)
	event.Destination.Fields.DestinationAddress.Name = flat.DstIP
	event.Destination.Fields.DestinationPort.Name = flat.DstPort
	event.Destination.Fields.DestinationBytes.Name = fmt.Sprint(flat.DstBytes)
	event.Destination.Fields.DestinationPackets.Name = fmt.Sprint(flat.DstPackets)
	var socket string
	if flat.Direction == "I" {
		event.Network.Fields.NetworkTransport.Name = "incoming"
		socket = fmt.Sprintf("%d/%s:%s-%s:%s", flat.Proto, flat.SrcIP, flat.SrcPort, flat.DstIP, flat.DstPort)
	}
	switch flat.Direction {
	case "O":
		event.Network.Fields.NetworkTransport.Name = "outgoing"
		socket = fmt.Sprintf("%d/%s:%s-%s:%s", flat.Proto, flat.DstIP, flat.DstPort, flat.SrcIP, flat.SrcPort)
	case "I":
		event.Network.Fields.NetworkTransport.Name = "incoming"
		socket = fmt.Sprintf("%d/%s:%s-%s:%s", flat.Proto, flat.SrcIP, flat.SrcPort, flat.DstIP, flat.DstPort)
	}

	switch proto_to_string(flat.Proto) {
	case "I":
		event.Network.Fields.NetworkTransport.Name = "icmp"
		socket = fmt.Sprintf("%s/%s:%s-%s:%s", strconv.Itoa(flat.Proto), flat.SrcIP, flat.SrcPort, flat.DstIP, flat.DstPort)
	case "U":
		event.Network.Fields.NetworkTransport.Name = "udp"
		socket = fmt.Sprintf("%s/%s:%s-%s:%s", strconv.Itoa(flat.Proto), flat.SrcIP, flat.SrcPort, flat.DstIP, flat.DstPort)
	case "T":
		event.Network.Fields.NetworkTransport.Name = "tcp"
		socket = fmt.Sprintf("%s/%s:%s-%s:%s", strconv.Itoa(flat.Proto), flat.SrcIP, flat.SrcPort, flat.DstIP, flat.DstPort)
	default:
		event.Network.Fields.NetworkTransport.Name = "unknown"
		socket = fmt.Sprintf("%s/unknown", strconv.Itoa(flat.Proto))
	}
	// Note: traceid uses the socket format (proto/srcIP:srcPort-dstIP:dstPort) for tracing bidirectional flows
	event.Tracing.Fields.TraceID.Name = socket
	event.Network.Fields.NetworkBytes.Name = fmt.Sprint(flat.SrcBytes + flat.DstBytes)
	event.Network.Fields.NetworkPackets.Name = fmt.Sprint(flat.SrcPackets + flat.DstPackets)

	event_json, _ := json.Marshal(event)
	fmt.Print(event_json)

	return string(event_json)
}

func proto_to_string(proto int) string {
	// https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml
	switch proto {
	case 1:
	case 58:
		return "I"
	case 6:
		return "T"
	case 17:
		return "U"
	default:
		x := strconv.Itoa(proto)
		return x
	}
	return "X"
}

func proto_to_int(proto string) int {
	// https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml
	switch proto {
	case "I":
		return 58
	case "T":
		return 6
	case "U":
		return 17
	default:
		return 0
	}
}
