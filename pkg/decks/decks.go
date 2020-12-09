package decks

import (
	"sync"

	"github.com/tombell/go-rekordbox"
)

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

// All returns all the known decks.
func (d *Decks) All() map[int]*Deck {
	d.Lock()
	defer d.Unlock()

	all := make(map[int]*Deck, 0)

	for _, deck := range d.decks {
		all[deck.ID] = deck
	}

	return all
}

// Notify will notify the correct deck which is playing or has played the track,
// and have it update the current track and append to that decks track history.
// This will also create any decks which do not exist internally.
func (d *Decks) Notify(track *rekordbox.Track) {
	d.Lock()
	defer d.Unlock()

	deckID := 1

	// XXX: right now we cannot get which deck the track played on, so since I
	// only play using two decks, and I always play the first track on deck 1,
	// let's assume, if the track number is even, it's played on deck 2.
	if track.Number%2 == 0 {
		deckID = 2
	}

	deck, ok := d.decks[deckID]
	if !ok {
		d.decks[deckID] = newDeck(deckID)
		deck = d.decks[deckID]
	}

	deck.notify(track)

	for ch := range d.listeners {
		ch <- true
	}
}
