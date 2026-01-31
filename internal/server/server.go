package server

import (
	"log/slog"
	"net/http"
)

type Server struct {
	myServer *http.Server
}

func New(router http.Handler, port string) *Server {
	return &Server{myServer: &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}}
}

func (s *Server) Run() error {
	slog.Info("Starting server", "ADDR", s.myServer.Addr)
	return s.myServer.ListenAndServe()
}
