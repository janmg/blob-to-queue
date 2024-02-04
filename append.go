package main

import "os"

// /SUBSCRIPTIONS/F5DD6E2D-1F42-4F54-B3BD-DBF595138C59/RESOURCEGROUPS/VM/PROVIDERS/MICROSOFT.NETWORK/NETWORKSECURITYGROUPS/OCTOBER-NSG/y=2023/m=10/d=31/h=18/m=00/macAddress=002248A31CA3/PT1H.json
func appendFile(nsg flatevent) {
	file, err := os.OpenFile("filename.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	Error(err)
	if _, err = file.WriteString(format("csv", nsg)); err != nil {
		Error(err)
		defer file.Close()
	}
}
