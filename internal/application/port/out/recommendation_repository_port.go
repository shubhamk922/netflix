package out

import "example.com/netflix/internal/domain"

type RecommendationRepositoty interface {
	GetAllMoviesByFrqId() ([]domain.Movie, error)

	GetAllMovies() ([]domain.Movie, error)
}
