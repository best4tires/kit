package srv

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"
)

type Server struct {
	listener net.Listener
	server   *http.Server
}

func New(bind string) (*Server, error) {
	l, err := net.Listen("tcp", bind)
	if err != nil {
		return nil, fmt.Errorf("listen-tcp on %q: %w", bind, err)
	}
	s := &Server{
		listener: l,
		server:   &http.Server{},
	}
	return s, nil
}

func (s *Server) Addr() net.Addr {
	return s.listener.Addr()
}

func (s *Server) Run(handler http.Handler) error {
	s.server.Handler = handler
	return s.server.Serve(s.listener)
}

func (s *Server) RunTLS(handler http.Handler, certFile, keyFile string) error {
	s.server.Handler = handler
	return s.server.ServeTLS(s.listener, certFile, keyFile)
}

func (s *Server) Close() {
	s.server.Close()
}

func (s *Server) Shutdown(timeout time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	s.server.Shutdown(ctx)
}
