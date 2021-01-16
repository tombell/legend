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
	listen          = flag.String("listen", "0.0.0.0:8888", "")
	pollingInterval = flag.Duration("pollingInterval", time.Second*5, "")
	version         = flag.Bool("version", false, "")
)

func main() {
	flag.Parse()

	if *version {
		fmt.Fprintf(os.Stdout, "legend %s (%s)\n", Version, Commit)
		os.Exit(0)
	}

	logger := log.New(os.Stderr, "", log.LstdFlags)

	options := &legend.Options{
		Logger:   logger,
		Listen:   *listen,
		Interval: *pollingInterval,
	}

	if err := legend.Run(options); err != nil {
		logger.Printf("error: %s", err)
	}
}
