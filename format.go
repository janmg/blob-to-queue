package main

import (
	"encoding/json"
	"fmt"

	"github.com/gocarina/gocsv"
)

func format(format string, nsg flatevent) string {
	var formatted string
	switch format {
	case "csv":
		formatted = format_csv(nsg)
	case "json":
		formatted = format_json(nsg)
	case "ecs":
		formatted = format_ecs(nsg)
	}
	return formatted
}

func format_json(nsg flatevent) string {
	event_json, _ := json.Marshal(nsg)
	fmt.Print(event_json)
	return string(event_json)
}

func format_csv(nsg flatevent) string {
	// copy from gocarina the csv reader and strip out the need for an array of structs, here I only have one event in a struct at a time and don't want the header
	// https://github.com/gocarina/gocsv/blob/master/csv.go
	nsgs := []*flatevent{}
	nsgs = append(nsgs, &nsg)
	csvContent, err := gocsv.MarshalString(&nsgs)
	handleError(err)
	return csvContent
}

// struct for outputting in Logstash ECS format
func format_ecs(flat flatevent) string {
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
		socket = flat.Proto + "/" + flat.SrcIP + ":" + flat.SrcPort + "-" + flat.DstIP + ":" + flat.DstPort
	}
	if flat.Action == "O" {
		event.Network.Fields.NetworkTransport.Name = "outgoing"
		socket = flat.Proto + "/" + flat.DstIP + ":" + flat.DstPort + "-" + flat.SrcIP + ":" + flat.SrcPort
	}
	// TODO add default
	if flat.Action == "U" {
		event.Network.Fields.NetworkTransport.Name = "udp"
	}
	if flat.Action == "T" {
		event.Network.Fields.NetworkTransport.Name = "tcp"
	}
	// TODO traceid is the socket T/10.0.0.5:1024-10.0.0.4:443 for the incoming direction source to destination, must flip src and dst to trace
	event.Tracing.Fields.TraceID.Name = socket
	event.Network.Fields.NetworkBytes.Name = fmt.Sprint(flat.SrcBytes + flat.DstBytes)
	event.Network.Fields.NetworkPackets.Name = fmt.Sprint(flat.SrcPackets + flat.DstPackets)

	event_json, _ := json.Marshal(event)
	fmt.Print(event_json)

	return string(event_json)
}
