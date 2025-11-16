package server

import (
	"fmt"
	"leetcode-rich-presence/internal/server/handlers"
	"log/slog"
	"net/http"
	"time"
)

type Server struct {
	Port      int
	Queue     chan<- handlers.Message
	startTime time.Time
	mux       *http.ServeMux
}

func NewServer(port int, queue chan<- handlers.Message) *Server {
	return &Server{
		Port:      port,
		Queue:     queue,
		startTime: time.Now(),
		mux:       http.NewServeMux(),
	}
}

func (s Server) addHandler(path string, handler handlers.Handler) {

	s.mux.HandleFunc(path, handler.Controller)
}

func (s Server) Start() error {

	s.addHandler("/", handlers.WebsocketHandler{Queue: s.Queue})

	s.addHandler("/health", handlers.HealthHandler{Start: s.startTime})

	server := &http.Server{
		Addr:    fmt.Sprintf("localhost:%d", s.Port),
		Handler: s.mux,
	}

	slog.Info("Starting server", "addr", server.Addr)

	return server.ListenAndServe()
}
