package main

import (
	"fmt"
	"net/http"

	"example.com/netflix/internal/adapter/in/controller"
	"example.com/netflix/internal/adapter/out/persistence"
	"example.com/netflix/internal/application/service"
	"example.com/netflix/internal/domain"
	"example.com/netflix/internal/infrastructure/logger"
)

func main() {

	domain_service := domain.NewTitleSimilarityEngine()
	repo := persistence.NewRecommednationAdapter()
	logger := logger.NewLogger()
	usercase := service.NewRecommendationService(domain_service, repo, logger)
	c := controller.NewRecommendationController(usercase)

	mux := http.NewServeMux()

	mux.HandleFunc("/movies/title/:id", c.GetAllSimilarMovies)

	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Server is Listening on port 8080")
}
