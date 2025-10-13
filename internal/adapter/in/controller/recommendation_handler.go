package controller

import (
	"net/http"

	"example.com/netflix/internal/application/port/in"
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
	return
}
