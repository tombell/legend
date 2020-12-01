package decks

import "github.com/tombell/go-rekordbox"

// Deck is a single deck in rekordbox, which has a currently playing track, and
// history of played tracks.
type Deck struct {
	ID      int
	Current *rekordbox.Track
	History []*rekordbox.Track
}

func newDeck(id int) *Deck {
	return &Deck{
		ID:      id,
		Current: nil,
		History: make([]*rekordbox.Track, 0),
	}
}

func (d *Deck) notify(track *rekordbox.Track) {
	if d.Current != nil && d.Current.ID == track.ID {
		return
	}

	d.Current = track
	d.History = append(d.History, track)

}
