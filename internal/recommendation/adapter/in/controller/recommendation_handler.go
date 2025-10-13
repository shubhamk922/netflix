package controller

import (
	"context"
	"net/http"

	"example.com/netflix/internal/recommendation/application/port/in"
)

type RecommendationController struct {
	usecase in.RecommendationPort
}

func NewRecommendationController(usecase in.RecommendationPort) *RecommendationController {

	return &RecommendationController{
		usecase: usecase,
	}
}

func (c *RecommendationController) GetAllSimilarMovies(w http.ResponseWriter, r *http.Request) {
	_ = context.Background()
	return
}
