package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/tombell/legend"
)

var (
	listen          = flag.String("listen", ":8888", "")
	rekordboxPath   = flag.String("rekordboxPath", "/Applications/rekordbox 6/rekordbox.app", "")
	pollingInterval = flag.Duration("pollingInterval", time.Second*30, "")
	version         = flag.Bool("version", false, "")
)

func main() {
	flag.Parse()

	if *version {
		fmt.Fprintf(os.Stdout, "legend %s (%s)\n", Version, Commit)
		os.Exit(0)
	}

	logger := log.New(os.Stderr, "[legend] ", log.LstdFlags)

	if err := legend.Run(logger, *listen, *rekordboxPath, *pollingInterval); err != nil {
		logger.Printf("error: %s", err)
	}
}
