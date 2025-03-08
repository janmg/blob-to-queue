package common

import (
	"log"
	"os"
)

/*
Should logging be implemnented as zerolog?
https://betterstack.com/community/guides/logging/zerolog/
*/
func Error(err error) {
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(2)
	}
}

func Warning(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
