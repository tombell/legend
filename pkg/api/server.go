package api

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"

	"github.com/tombell/legend/pkg/playlist"
	"github.com/tombell/legend/pkg/web"
)

// Server serves the status of the playlist to clients via websockets.
type Server struct {
	logger   *log.Logger
	playlist *playlist.Playlist
	server   *http.Server
}

// New returns an initialised Server with the given playlist.
func New(logger *log.Logger, playlist *playlist.Playlist, listen string) *Server {
	server := &http.Server{
		Addr:         listen,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return &Server{
		logger:   logger,
		playlist: playlist,
		server:   server,
	}
}

// Start registers the handler, and start listening for incoming connections.
func (s *Server) Start(ch chan error) {
	s.logger.Println("registering http handler...")

	mux := http.NewServeMux()
	mux.Handle("/public/", http.FileServer(http.FS(web.FS)))
	mux.Handle("/", s.handler())

	s.server.Handler = mux

	s.logger.Println("starting api server...")

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

	s.playlist.AddNotificationChannel(ch)

	defer func() {
		s.playlist.RemoveNotificationChannel(ch)
		conn.Close()
	}()

	if err := conn.WriteJSON(s.playlist.BuildResponse()); err != nil {
		return
	}

	for {
		select {
		case <-ch:
			if err := conn.WriteJSON(s.playlist.BuildResponse()); err != nil {
				break
			}
		}
	}
}
