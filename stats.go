package main

import "fmt"

type summary struct {
	unixtime int
	socket   string
	inbytes  int
	outbytes int
	inpack   int
	outpack  int
}

func statistics(time int, proto string, state string, direction string, src string, dst string) {
	stats := make(map[string]summary) // K socket, V summary
	socket := proto + "/" + src + "-" + dst
	if direction == "O" {
		socket = proto + "/" + dst + "-" + src
	}

	if proto == "T" {
		if state == "B" {
			stats[socket] = summary{}
		} else {
			// find socket from map
			_, found := stats[socket]
			if found {
				stats[socket] = summary{}
			} else {
				// not B and not found ... mmh
				stats[socket] = summary{}
				fmt.Println(fmt.Sprintf("not B and not previously found? %s", socket))
			}

		}
		// lookup from map src+"-"+dst and dst+"-"+src, check time
	}

	if proto == "U" {
		// lookup from map src+"-"+dst and dst+"-"+src, check time
		_, found := stats[socket]
		if !found {
			//stats[socket] = summary{time, socket, inbytes, outbytes, inpack, outpack}
		}
		// if new add src+"-"+dst to map
	}
	if direction == "I" {
		// add src byte and packs to incoming
		// add dst byte and packs to outgoing
	}
	// save packets to map
	// keep timer to find expired connections?
	/*
		if now.Unix()-stats.time > int64(4*60) {
			delete(stats[socket])
		}
	*/
}
