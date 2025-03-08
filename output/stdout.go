package output

import (
	"fmt"

	"janmg.com/blob-to-queue/format"
)

// TODO hardcoded format, replace with configHandler info
func Stdout(nsg format.Flatevent) {
	fmt.Print(format.Format("csv", nsg))
}
