package api

import (
	"net/http"

	"github.com/gorilla/websocket"
)

// Server ...
type Server struct {
	server *http.Server
}

// New ...
func New() *Server {
	return nil
}

func (s *Server) register(conn *websocket.Conn) {
	ch := make(chan bool, 1)

	for {
		select {
		case <-ch:
			// TODO: add data to channel
		}
	}

	conn.Close()
}
