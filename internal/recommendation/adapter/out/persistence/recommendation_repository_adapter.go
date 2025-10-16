package persistence

import (
	"context"

	"example.com/netflix/internal/recommendation/domain"
)

type RecommendationAdapter struct {
	movies []domain.Movie
}

func NewRecommednationAdapter() *RecommendationAdapter {
	var movies []domain.Movie

	movies = append(movies, *domain.NewMovie("The manz"))
	movies = append(movies, *domain.NewMovie("The man"))
	movies = append(movies, *domain.NewMovie("Tha nem"))
	movies = append(movies, *domain.NewMovie("Running fox"))
	movies = append(movies, *domain.NewMovie("Th man"))
	movies = append(movies, *domain.NewMovie("Deep code"))
	movies = append(movies, *domain.NewMovie("Sky fall"))
	return &RecommendationAdapter{
		movies: movies,
	}
}

func (r *RecommendationAdapter) GetAllMovies(ctx context.Context) ([]domain.Movie, error) {
	return r.movies, nil
}

func (r *RecommendationAdapter) GetAllMoviesByFrqId(ctx context.Context, freqId string) ([]domain.Movie, error) {
	var result []domain.Movie
	for _, movie := range r.movies {
		if movie.FreqId() == freqId {
			result = append(result, movie)
		}
	}
	return result, nil
}
