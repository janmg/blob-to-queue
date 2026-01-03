package common

import (
	"log"
	"os"
)

/*
Should logging be implemented as zerolog or slog?
https://betterstack.com/community/guides/logging/zerolog/
slog "github.com/sagikazarmark/slog-shim"
*/
func Error(err error) {
	// Errors should log a message and stop the process
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(2)
	}
}

func Warning(warn error) {
	// Warnings log a message and continue
	if warn != nil {
		log.Fatal(warn.Error())
	}
}
