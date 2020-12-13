package api

import "github.com/tombell/legend/pkg/playlist"

type trackResponse struct {
	Artist   string `json:"artist"`
	Name     string `json:"name"`
	AlbumArt string `json:"album_art"`
}

type deckResponse struct {
	Current *trackResponse   `json:"current"`
	History []*trackResponse `json:"history"`
}

type response struct {
	Deck *deckResponse `json:"deck"`
}

func buildResponse(deck *playlist.Deck) *response {
	resp := &response{}

	resp.Deck = &deckResponse{
		History: make([]*trackResponse, 0, len(deck.History)),
	}

	if deck.Current != nil {
		resp.Deck.Current = &trackResponse{
			Artist:   deck.Current.Artist,
			Name:     deck.Current.Name,
			AlbumArt: deck.Current.ImagePath,
		}
	}

	for _, track := range deck.History {
		resp.Deck.History = append(resp.Deck.History, &trackResponse{
			Artist:   track.Artist,
			Name:     track.Name,
			AlbumArt: track.ImagePath,
		})
	}

	return resp
}
