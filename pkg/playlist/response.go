package playlist

// Track is the JSON representation of a played track in a playlist.
type Track struct {
	Number   int    `json:"number"`
	Artist   string `json:"artist"`
	Name     string `json:"name"`
	AlbumArt string `json:"album_art"`
}

// Response is the JSON representation of the playlist.
type Response struct {
	Current *Track   `json:"current"`
	History []*Track `json:"history"`
}

// BuildResponse builds the response struct up ready to send to the websocket
// clients.
func (p *Playlist) BuildResponse() *Response {
	resp := &Response{
		History: make([]*Track, 0, len(p.History)),
	}

	if p.Current != nil {
		resp.Current = &Track{
			Number:   p.Current.Number,
			Artist:   p.Current.Artist,
			Name:     p.Current.Name,
			AlbumArt: p.Current.ImagePath,
		}
	}

	for _, track := range p.History {
		resp.History = append(resp.History, &Track{
			Number:   track.Number,
			Artist:   track.Artist,
			Name:     track.Name,
			AlbumArt: track.ImagePath,
		})
	}

	return resp
}
