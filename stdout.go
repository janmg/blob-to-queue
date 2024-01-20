package main

import (
	"fmt"
)

// TODO hardcoded format, replace with configHandler info
func stdout(nsg flatevent) {
	fmt.Println(format("csv", nsg))
}
