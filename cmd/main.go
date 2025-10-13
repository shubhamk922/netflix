package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	httpadapter "example.com/netflix/internal/infrastructure/http"
	"example.com/netflix/internal/infrastructure/logger"
	"example.com/netflix/internal/recommendation/adapter/in/controller"
	"example.com/netflix/internal/recommendation/adapter/out/persistence"
	"example.com/netflix/internal/recommendation/application/service"
	"example.com/netflix/internal/recommendation/domain"
)

func main() {

	domain_service := domain.NewTitleSimilarityEngine()
	repo := persistence.NewRecommednationAdapter()
	logger := logger.NewLogger()
	usercase := service.NewRecommendationService(domain_service, repo, logger)
	c := controller.NewRecommendationController(usercase)
	// Initialize mux and map routes
	mux := http.NewServeMux()
	mux.HandleFunc("/movies/title/:id", c.GetAllSimilarMovies)

	port := "8080"
	server := httpadapter.NewServer(port, mux)

	// server start
	server.Prestart()
	err := server.Start()
	if err != nil {
		log.Fatalf("error is starting server %+v", err)
	}
	log.Printf("Server is lsitening on port %s ", port)

	// so as to avoid shutdown the main function
	done := make(chan struct{})
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
	// Now Lets write a graceful shutdown function
	go func() {
		<-interrupt
		server.Shutdown()
		done <- struct{}{}
	}()

	<-done
}
