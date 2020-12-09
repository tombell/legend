package monitor

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/tombell/go-rekordbox"

	"github.com/tombell/legend/pkg/decks"
)

// Monitor ...
type Monitor struct {
	logger   *log.Logger
	db       *sql.DB
	interval time.Duration
	decks    *decks.Decks
}

// New ...
func New(logger *log.Logger, db *sql.DB, interval time.Duration, decks *decks.Decks) *Monitor {
	return &Monitor{
		logger:   logger,
		db:       db,
		interval: interval,
		decks:    decks,
	}
}

// Run ...
func (m *Monitor) Run(ch chan error) {
	if err := m.handle(); err != nil {
		ch <- err
		return
	}

	tick := time.Tick(m.interval)

	for range tick {
		if err := m.handle(); err != nil {
			ch <- err
			continue
		}
	}
}

func (m *Monitor) handle() error {
	track, err := rekordbox.GetRecentTrack(m.db)
	if err != nil {
		return fmt.Errorf("rekordbox get recent track failed: %w", err)
	}

	m.logger.Printf("notifying decks of current track: %s - %s", track.Artist, track.Name)
	m.decks.Notify(track)

	return nil
}
