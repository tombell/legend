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
	pollingInterval = flag.Duration("pollingInterval", time.Second*5, "")
	version         = flag.Bool("version", false, "")
)

func main() {
	flag.Parse()

	if *version {
		fmt.Fprintf(os.Stdout, "legend %s (%s)\n", Version, Commit)
		os.Exit(0)
	}

	logger := log.New(os.Stderr, "[legend] ", log.LstdFlags)

	options := &legend.Options{
		Logger:   logger,
		Listen:   *listen,
		Interval: *pollingInterval,
	}

	if err := legend.Run(options); err != nil {
		logger.Printf("error: %s", err)
	}
}
