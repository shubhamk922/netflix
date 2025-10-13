package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"example.com/netflix/internal/infrastructure/logger"
	"example.com/netflix/internal/recommendation/adapter/in/controller"
	"example.com/netflix/internal/recommendation/adapter/out/persistence"
	"example.com/netflix/internal/recommendation/application/service"
	"example.com/netflix/internal/recommendation/domain"
)

type GracefulServer struct {
	Server   *http.Server
	Listener net.Listener
}

func NewServer(port string) *GracefulServer {
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}
	return &GracefulServer{Server: server}
}

func (s *GracefulServer) start() error {
	listener, err := net.Listen("tcp", s.Server.Addr)

	if err != nil {
		return err
	}
	s.Listener = listener
	go s.Server.Serve(s.Listener)
	// log
	return nil
}

func (s *GracefulServer) shutdown() error {
	if s.Listener != nil {
		err := s.Listener.Close()
		// we can check for all resource conection if that is close or not to gracefully shutdown
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {

	domain_service := domain.NewTitleSimilarityEngine()
	repo := persistence.NewRecommednationAdapter()
	logger := logger.NewLogger()
	usercase := service.NewRecommendationService(domain_service, repo, logger)
	c := controller.NewRecommendationController(usercase)

	port := "8080"
	server := NewServer(port)
	err := server.start()

	if err != nil {
		log.Fatalf("error is starting server %+v", err)
	}

	log.Printf("Server is lsitening on port %s ", port)

	handler := server.Server.Handler.(*http.ServeMux)
	handler.HandleFunc("/movies/title/:id", c.GetAllSimilarMovies)

	// so as to avoid shutdown the main function
	done := make(chan struct{})
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
	// Now Lets write a graceful shutdown function
	go func() {
		<-interrupt
		server.shutdown()
		done <- struct{}{}
	}()

	<-done
}
