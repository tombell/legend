package api

import (
	"context"
	"net/http"
	"time"

	"github.com/tombell/legend/pkg/decks"
)

// Server serves the status of the decks to clients via websockets.
type Server struct {
	server *http.Server
	decks  *decks.Decks
}

// New returns an initialised Server with the given decks.
func New(decks *decks.Decks) *Server {
	server := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return &Server{
		server: server,
		decks:  decks,
	}
}

// Start registers the handler, and start listening for incoming connections.
func (s *Server) Start() error {
	s.server.Handler = s.handler()
	return s.server.ListenAndServe()
}

// Shutdown shuts the running server down.
func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
