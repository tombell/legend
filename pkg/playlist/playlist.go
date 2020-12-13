package playlist

import (
	"sync"

	"github.com/tombell/go-rekordbox"
)

// Playlist represents the current track and track history of rekordbox.
type Playlist struct {
	sync.Mutex

	listeners map[chan bool]bool

	Current *rekordbox.Track
	History []*rekordbox.Track
}

// New returns an initialised Playlist model.
func New() *Playlist {
	return &Playlist{
		listeners: make(map[chan bool]bool, 0),
		Current:   nil,
		History:   make([]*rekordbox.Track, 0),
	}
}

// AddNotificationChannel adds a channel to be notified when the playlist is
// updated.
func (d *Playlist) AddNotificationChannel(ch chan bool) {
	d.Lock()
	defer d.Unlock()

	d.listeners[ch] = true
}

// RemoveNotificationChannel removes a channel from the list of notification
// channels.
func (d *Playlist) RemoveNotificationChannel(ch chan bool) {
	d.Lock()
	defer d.Unlock()

	delete(d.listeners, ch)
}

// Notify will notify will update the current track and history if the track has
// changed from the current.
func (d *Playlist) Notify(track *rekordbox.Track) {
	d.Lock()
	defer d.Unlock()

	if d.Current != nil && d.Current.ID == track.ID {
		return
	}

	d.Current = track
	d.History = append(d.History, track)

	for ch := range d.listeners {
		ch <- true
	}
}
