package monitor

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/tombell/go-rekordbox"

	"github.com/tombell/legend/pkg/playlist"
)

// Monitor is a struct that polls rekordbox for the currently playing track.
type Monitor struct {
	logger   *log.Logger
	db       *sql.DB
	interval time.Duration
	playlist *playlist.Playlist
}

// New returns an initialised monitor.
func New(logger *log.Logger, db *sql.DB, interval time.Duration, playlist *playlist.Playlist) *Monitor {
	return &Monitor{
		logger:   logger,
		db:       db,
		interval: interval,
		playlist: playlist,
	}
}

// Run starts the polling of the currently playing track in rekordbox.
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

	m.logger.Printf("notifying playlist of current track: %s - %s", track.Artist, track.Name)
	m.playlist.Notify(track)

	return nil
}
