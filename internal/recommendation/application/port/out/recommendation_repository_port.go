package out

import (
	"context"

	"example.com/netflix/internal/recommendation/domain"
)

type RecommendationRepositoty interface {
	GetAllMoviesByFrqId(ctx context.Context) ([]domain.Movie, error)

	GetAllMovies(ctx context.Context) ([]domain.Movie, error)
}
