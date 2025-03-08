package output

import (
	"os"

	"janmg.com/blob-to-queue/common"
	"janmg.com/blob-to-queue/format"
)

// /SUBSCRIPTIONS/F5DD6E2D-1F42-4F54-B3BD-DBF595138C59/RESOURCEGROUPS/VM/PROVIDERS/MICROSOFT.NETWORK/NETWORKSECURITYGROUPS/OCTOBER-NSG/y=2023/m=10/d=31/h=18/m=00/macAddress=002248A31CA3/PT1H.json
func AppendFile(nsg format.Flatevent) {
	file, err := os.OpenFile("filename.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	common.Error(err)
	if _, err = file.WriteString(format.Format("csv", nsg)); err != nil {
		common.Error(err)
		defer file.Close()
	}
}
