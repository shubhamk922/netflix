package in

import "example.com/netflix/internal/domain"

type RecommendationPort interface {
	Recommend() ([]domain.Movie, error)
}
