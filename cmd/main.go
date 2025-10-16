package main

import (
	"log"
	"net/http"

	httpadapter "example.com/netflix/internal/infrastructure/http"
	"example.com/netflix/internal/infrastructure/logger"
	"example.com/netflix/internal/recommendation/adapter/in/controller"
	"example.com/netflix/internal/recommendation/adapter/out/persistence"
	"example.com/netflix/internal/recommendation/application/service"
	"example.com/netflix/internal/recommendation/domain"
)

func main() {

	domain_service := domain.NewTitleSimilarityEngine(3)
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
	done, err := server.Start()
	if err != nil {
		log.Fatalf("error is starting server %+v", err)
		server.Shutdown()
	}
	log.Printf("Server is lsitening on port %s ", port)
	<-done // wait for graceful shutdown
}
