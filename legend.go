package legend

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/tombell/go-rekordbox"

	"github.com/tombell/legend/pkg/api"
	"github.com/tombell/legend/pkg/monitor"
	"github.com/tombell/legend/pkg/playlist"
)

const rekordboxPath = "/Applications/rekordbox 6/rekordbox.app"

// Options contains the configurable properties.
type Options struct {
	Logger   *log.Logger
	Listen   string
	Interval time.Duration
}

// Run starts the API server, and rekordbox polling monitor.
func Run(options *Options) error {
	db, err := rekordbox.OpenDatabase(rekordboxPath)
	if err != nil {
		return fmt.Errorf("rekordbox open database failed: %w", err)
	}

	playlist := playlist.New()

	errCh := make(chan error, 1)

	mon := monitor.New(options.Logger, db, options.Interval, playlist)
	srv := api.New(options.Logger, playlist, options.Listen)

	go mon.Run(errCh)
	go srv.Start(errCh)

	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-errCh:
		if err == http.ErrServerClosed {
			return nil
		}
		return err
	case <-ch:
		break
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		return fmt.Errorf("server shutdown failed: %w", err)
	}

	return nil
}
