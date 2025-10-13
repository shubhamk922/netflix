package http

import (
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

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

func (s *GracefulServer) Start() (chan bool, error) {
	listener, err := net.Listen("tcp", s.Server.Addr)

	if err != nil {
		return nil, err
	}
	s.Listener = listener
	go s.Server.Serve(s.Listener)
	// log
	// so as to avoid shutdown the main function
	done := make(chan bool, 1)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
	// Now Lets write a graceful shutdown function
	go func() {
		<-interrupt
		s.Shutdown()
		done <- true
	}()
	return done, nil
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
