package decks

import "sync"

// Decks represents a set of decks inside rekordbox that have played tracks.
type Decks struct {
	sync.Mutex

	decks     map[int]*Deck
	listeners map[chan bool]bool
}

// New returns an initialised Decks model.
func New() *Decks {
	return &Decks{
		decks:     make(map[int]*Deck, 0),
		listeners: make(map[chan bool]bool, 0),
	}
}

// AddNotificationChannel adds a channel to be notified when the decks are
// updated.
func (d *Decks) AddNotificationChannel(ch chan bool) {
	d.Lock()
	defer d.Unlock()

	d.listeners[ch] = true
}

// RemoveNotificationChannel removes a channel from the list of notification
// channels.
func (d *Decks) RemoveNotificationChannel(ch chan bool) {
	d.Lock()
	defer d.Unlock()

	delete(d.listeners, ch)
}
