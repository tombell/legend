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
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		go s.register(conn)
	}
}

func (s *Server) register(conn *websocket.Conn) {
	ch := make(chan bool, 1)

	s.decks.AddNotificationChannel(ch)

	defer func() {
		s.decks.RemoveNotificationChannel(ch)
		conn.Close()
	}()

	resp := buildResponse(s.decks.All())
	if err := conn.WriteJSON(resp); err != nil {
		return
	}

	for {
		select {
		case <-ch:
			resp := buildResponse(s.decks.All())
			if err := conn.WriteJSON(resp); err != nil {
				break
			}
		}
	}
}
