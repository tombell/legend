package api

import (
	"net/http"

	"github.com/gorilla/websocket"
)

func (s *Server) handler() http.HandlerFunc {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}

	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			// TODO: error logging
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		go s.register(conn)
	}
}

func (s *Server) register(conn *websocket.Conn) {
	ch := make(chan bool, 1)

	s.decks.AddNotificationChannel(ch)

	// TODO: add data to channel.

	for {
		select {
		case <-ch:
			// TODO: add data to channel.
		}
	}

	s.decks.RemoveNotificationChannel(ch)
	conn.Close()
}
