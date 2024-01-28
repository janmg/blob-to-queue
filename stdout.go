package main

import (
	"fmt"
)

// TODO hardcoded format, replace with configHandler info
func stdout(nsg flatevent) {
	fmt.Print(format("csv", nsg))
}
