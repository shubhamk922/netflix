package http

import (
	"errors"
	"log"
	"net"
	"net/http"

	"example.com/netflix/internal/infrastructure/http/middleware"
	"example.com/netflix/internal/infrastructure/logger"
)

type GracefulServer struct {
	Server   *http.Server
	Listener net.Listener
}

func NewServer(port string, handler http.Handler) *GracefulServer {

	server := &http.Server{
		Addr:    ":" + port,
		Handler: handler,
	}
	return &GracefulServer{Server: server}
}

func (s *GracefulServer) Prestart() error {
	logger := logger.InitLogger()
	if logger == nil {
		errMsg := "logger is not initialized"
		log.Println(errMsg)
		return errors.New(errMsg)
	}
	s.Server.Handler = middleware.SimpleLogger(s.Server.Handler)
	return nil
}

func (s *GracefulServer) Start() error {
	listener, err := net.Listen("tcp", s.Server.Addr)

	if err != nil {
		return err
	}
	s.Listener = listener
	go s.Server.Serve(s.Listener)
	// log
	return nil
}

func (s *GracefulServer) Shutdown() error {
	logger.Close()
	if s.Listener != nil {
		err := s.Listener.Close()
		// we can check for all resource conection if that is close or not to gracefully shutdown
		if err != nil {
			return err
		}
	}

	return nil
}
