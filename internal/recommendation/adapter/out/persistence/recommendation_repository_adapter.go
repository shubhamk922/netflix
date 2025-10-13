package persistence

import (
	"context"

	"example.com/netflix/internal/recommendation/domain"
)

type RecommendationAdapter struct {
}

func NewRecommednationAdapter() *RecommendationAdapter {
	return &RecommendationAdapter{}
}

func (r *RecommendationAdapter) GetAllMovies(ctx context.Context) ([]domain.Movie, error) {
	return []domain.Movie{}, nil
}

func (r *RecommendationAdapter) GetAllMoviesByFrqId(ctx context.Context) ([]domain.Movie, error) {
	return []domain.Movie{}, nil
}
