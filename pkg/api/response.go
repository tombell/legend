package api

import "github.com/tombell/legend/pkg/playlist"

type trackResponse struct {
	Artist   string `json:"artist"`
	Name     string `json:"name"`
	AlbumArt string `json:"album_art"`
}

type playlistResponse struct {
	Current *trackResponse   `json:"current"`
	History []*trackResponse `json:"history"`
}

type response struct {
	Playlist *playlistResponse `json:"playlist"`
}

func buildResponse(list *playlist.Playlist) *response {
	resp := &response{
		Playlist: &playlistResponse{
			History: make([]*trackResponse, 0, len(list.History)),
		},
	}

	if list.Current != nil {
		resp.Playlist.Current = &trackResponse{
			Artist:   list.Current.Artist,
			Name:     list.Current.Name,
			AlbumArt: list.Current.ImagePath,
		}
	}

	for _, track := range list.History {
		resp.Playlist.History = append(resp.Playlist.History, &trackResponse{
			Artist:   track.Artist,
			Name:     track.Name,
			AlbumArt: track.ImagePath,
		})
	}

	return resp
}
