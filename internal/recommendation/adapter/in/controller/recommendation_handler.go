package controller

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"example.com/netflix/internal/infrastructure/logger"
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
	ctx := context.Background()
	title := r.URL.Query().Get("search")

	if title == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid Params"))
		return
	}

	movies, err := c.usecase.Recommend(ctx, title)

	if err != nil {
		w.WriteHeader(http.StatusOK)
		logger.NewLogger().Errorf("Error in Getting Movies %+v", err)
		return
	}

	if len(movies) == 0 {
		w.WriteHeader(http.StatusOK)
		logger.NewLogger().Infof("No movies found based on search ")
		io.WriteString(w, "No Movies found based on search")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(movies)
}
