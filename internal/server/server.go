package server

import (
	"context"
	"net/http"
)

type Server struct {
	httpserver *http.Server
}

func (s *Server) Run(addr string, handler http.Handler) error {
	s.httpserver = &http.Server{
		Addr:    addr,
		Handler: handler,
	}

	return s.httpserver.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpserver.Shutdown(ctx)
}
