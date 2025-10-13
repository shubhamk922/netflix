package persistence

import "example.com/netflix/internal/domain"

type RecommendationAdapter struct {
}

func NewRecommednationAdapter() *RecommendationAdapter {
	return &RecommendationAdapter{}
}

func (r *RecommendationAdapter) GetAllMovies() ([]domain.Movie, error) {
	return []domain.Movie{}, nil
}

func (r *RecommendationAdapter) GetAllMoviesByFrqId() ([]domain.Movie, error) {
	return []domain.Movie{}, nil
}
