package api

import "github.com/tombell/legend/pkg/decks"

type trackResponse struct {
	Artist   string `json:"artist"`
	Name     string `json:"name"`
	AlbumArt string `json:"album_art"`
}

type deckResponse struct {
	ID      int              `json:"id"`
	Current *trackResponse   `json:"current"`
	History []*trackResponse `json:"history"`
}

type response struct {
	Decks []*deckResponse `json:"decks"`
}

func buildResponse(decks map[int]*decks.Deck) *response {
	resp := &response{Decks: make([]*deckResponse, 0, len(decks))}

	for _, deck := range decks {
		d := &deckResponse{
			ID:      deck.ID,
			History: make([]*trackResponse, 0, len(deck.History)),
		}

		if deck.Current != nil {
			d.Current = &trackResponse{
				Artist:   deck.Current.Artist,
				Name:     deck.Current.Name,
				AlbumArt: deck.Current.ImagePath,
			}
		}

		for _, track := range deck.History {
			d.History = append(d.History, &trackResponse{
				Artist:   track.Artist,
				Name:     track.Name,
				AlbumArt: track.ImagePath,
			})
		}

		resp.Decks = append(resp.Decks, d)
	}

	return resp
}
