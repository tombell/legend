package monitor

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/tombell/go-rekordbox"

	"github.com/tombell/legend/pkg/decks"
)

// Monitor ...
type Monitor struct {
	db       *sql.DB
	interval time.Duration
	decks    *decks.Decks
}

// New ...
func New(db *sql.DB, interval time.Duration) *Monitor {
	return &Monitor{
		db:       db,
		interval: interval,
	}
}

// Run ...
func (m *Monitor) Run(ch chan error) {
	track, err := m.fetchTrack()
	if err != nil {
		ch <- err
		return
	}

	m.decks.Notify(track)

	tick := time.Tick(m.interval)

	for range tick {
		track, err := m.fetchTrack()
		if err != nil {
			ch <- err
			continue
		}

		m.decks.Notify(track)
	}
}

func (m *Monitor) fetchTrack() (*rekordbox.Track, error) {
	track, err := rekordbox.GetRecentTrack(m.db)
	if err != nil {
		return nil, fmt.Errorf("rekordbox get recent track failed: %w", err)
	}

	return track, nil
}
