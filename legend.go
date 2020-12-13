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

// Run ...
func Run(logger *log.Logger, listen, rekordboxPath string, interval time.Duration) error {
	db, err := rekordbox.OpenDatabase(rekordboxPath)
	if err != nil {
		return fmt.Errorf("rekordbox open database failed: %w", err)
	}

	playlist := playlist.New()

	errCh := make(chan error, 1)

	m := monitor.New(logger, db, time.Second*30, playlist)
	s := api.New(logger, playlist, listen)

	go m.Run(errCh)
	go s.Start(errCh)

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

	if err := s.Shutdown(ctx); err != nil {
		return fmt.Errorf("server shutdown failed: %w", err)
	}

	return nil
}
