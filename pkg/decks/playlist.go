package decks

import (
	"sync"

	"github.com/tombell/go-rekordbox"
)

// Playlist represents a set of decks inside rekordbox that have played tracks.
type Playlist struct {
	sync.Mutex

	listeners map[chan bool]bool

	Deck *Deck
}

// New returns an initialised Playlist model.
func New() *Playlist {
	return &Playlist{
		Deck:      newDeck(),
		listeners: make(map[chan bool]bool, 0),
	}
}

// AddNotificationChannel adds a channel to be notified when the decks are
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

// Notify will notify the correct deck which is playing or has played the track,
// and have it update the current track and append to that decks track history.
// This will also create any decks which do not exist internally.
func (d *Playlist) Notify(track *rekordbox.Track) {
	d.Lock()
	defer d.Unlock()

	d.Deck.notify(track)

	for ch := range d.listeners {
		ch <- true
	}
}
