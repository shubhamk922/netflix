package in

import (
	"context"

	"example.com/netflix/internal/recommendation/domain"
)

type RecommendationPort interface {
	Recommend(ctx context.Context) ([]domain.Movie, error)
}
