package api

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"

	"github.com/tombell/legend/pkg/decks"
)

// Server serves the status of the decks to clients via websockets.
type Server struct {
	logger *log.Logger
	server *http.Server
	decks  *decks.Decks
}

// New returns an initialised Server with the given decks.
func New(logger *log.Logger, decks *decks.Decks, listen string) *Server {
	server := &http.Server{
		Addr:         listen,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return &Server{
		logger: logger,
		server: server,
		decks:  decks,
	}
}

// Start registers the handler, and start listening for incoming connections.
func (s *Server) Start(ch chan error) {
	s.logger.Println("registering http handler...")
	s.server.Handler = s.handler()
	s.logger.Println("starting api server, listening on http://localhost:8888...")
	if err := s.server.ListenAndServe(); err != nil {
		ch <- err
	}
}

// Shutdown shuts the running server down.
func (s *Server) Shutdown(ctx context.Context) error {
	s.logger.Println("shutting down api server...")
	return s.server.Shutdown(ctx)
}

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
