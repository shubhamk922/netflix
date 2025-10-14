package service

import (
	"context"

	"example.com/netflix/internal/recommendation/application/port/out"
	"example.com/netflix/internal/recommendation/domain"
)

type RecommendationService struct {
	engine domain.TitleSimilarityEngine
	repo   out.RecommendationRepositoty
	log    out.Logger
}

func NewRecommendationService(engine domain.TitleSimilarityEngine, repo out.RecommendationRepositoty, log out.Logger) *RecommendationService {
	return &RecommendationService{
		engine: engine,
		repo:   repo,
		log:    log,
	}
}

func (r *RecommendationService) Recommend(ctx context.Context, title string) ([]domain.Movie, error) {
	return []domain.Movie{}, nil
}
