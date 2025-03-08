package format

import "time"

// struct for (temporary) storing event in flat format as json compatbile struct, to easily transform the event to csv or ecs
type Flatevent struct {
	Time              time.Time `json:"time"`
	FlowLogVersion    string    `json:"flowlogversion"`
	FlowLogGUID       string    `json:"flowlogguid"`
	SystemID          string    `json:"systemId"`
	MACAdress         string    `json:"macAddress"`
	Category          string    `json:"category"`
	ResourceID        string    `json:"resourceId"`
	FlowLogResourceID string    `json:"flowlogresourceid"`
	TargetResourceID  string    `json:"targetresourceid"`
	OperationName     string    `json:"operationName"`
	Version           int       `json:"Version"`
	ACLID             string    `json:"aclid"`
	Rule              string    `json:"rule"`
	Mac               string    `json:"mac"`
	Unixtime          string    `json:"unixtime"`
	SrcIP             string    `json:"srcip"`
	DstIP             string    `json:"dstip"`
	SrcPort           string    `json:"srcport"`
	DstPort           string    `json:"dstport"`
	Proto             string    `json:"proto"` // I would have liked Proto to be an int, but nsgflowlogs uses the letters T and U for TCP and UDP instead
	Direction         string    `json:"direction"`
	Action            string    `json:"action"`
	State             string    `json:"state"`
	Encryption        bool      `json:"encryption"`
	SrcPackets        int       `json:"srcpackets"`
	SrcBytes          int       `json:"srcbytes"`
	DstPackets        int       `json:"dstpackets"`
	DstBytes          int       `json:"dstbytes"`
}
