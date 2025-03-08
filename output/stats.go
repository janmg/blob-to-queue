package output

import (
	"fmt"
	"strconv"

	"janmg.com/blob-to-queue/common"
	"janmg.com/blob-to-queue/format"
)

type summary struct {
	unixtime int
	inbytes  int
	outbytes int
	inpack   int
	outpack  int
}

func Statistics(nsg format.Flatevent) {
	unix, err := strconv.Atoi(nsg.Unixtime)
	common.Error(err)
	// stats must be kept globally and exported somehow?
	var stats = make(map[string]summary) // K socket, V summary
	// Write stats every minute to a file?
	src := nsg.SrcIP + "_" + nsg.SrcPort
	dst := nsg.DstIP + "_" + nsg.DstPort

	socket := nsg.Proto + "/" + src + "-" + dst
	outbytes := nsg.SrcBytes
	inbytes := nsg.DstBytes
	outpack := nsg.SrcPackets
	inpack := nsg.DstPackets

	if nsg.Direction == "O" {
		socket = nsg.Proto + "/" + dst + "-" + src
		inbytes = nsg.SrcBytes
		outbytes = nsg.DstBytes
		inpack = nsg.SrcPackets
		outpack = nsg.DstPackets
	}

	// 1698775873,10.0.0.4,52.239.137.164,57756,443,T,O,A,B,,,,
	// 1698775879,10.0.0.4,52.239.137.164,57756,443,T,O,A,E,8,1474,18,18980
	// be more clever when it comes to TCP with B,x,E packets and UDP, and the deal with socket timeout?
	if nsg.Proto == "T" {
		if nsg.State == "B" {
			stats[socket] = summary{}
		} else {
			// find socket from map
			_, found := stats[socket]
			if found {
				stats[socket] = summary{}
			} else {
				// not B and not found ... mmh
				stats[socket] = summary{}
				stats[socket] = summary{unix, stats[socket].inbytes + inbytes, stats[socket].outbytes + outbytes, stats[socket].inpack + inpack, stats[socket].outpack + outpack}
				fmt.Printf(fmt.Sprintf("out of sync? no previous packet captured and not previously found? %s\n", socket))
			}
		}
		// lookup from map src+"-"+dst and dst+"-"+src, check time
	}

	if nsg.Proto == "U" {
		// lookup from map src+"-"+dst and dst+"-"+src, check time
		_, found := stats[socket]
		if !found {
			stats[socket] = summary{unix, 0, 0, 0, 0}
		} else {
			// if unix is expired?
			stats[socket] = summary{unix, stats[socket].inbytes + inbytes, stats[socket].outbytes + outbytes, stats[socket].inpack + inpack, stats[socket].outpack + outpack}
		}
		// if new add src+"-"+dst to map
	}

	// save packets to map
	// keep timer to find expired connections?
	/*
		if now.Unix()-stats.time > int64(4*60) {
			delete(stats[socket])
		}
	*/
}
