package decks

import "github.com/tombell/go-rekordbox"

// Deck is a single deck in rekordbox, which has a currently playing track, and
// history of played tracks.
type Deck struct {
	ID      int
	Current *rekordbox.Track
	History []*rekordbox.Track
}
