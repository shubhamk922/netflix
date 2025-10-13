package service

import (
	"example.com/netflix/internal/application/port/out"
	"example.com/netflix/internal/domain"
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

func (r *RecommendationService) Recommend() ([]domain.Movie, error) {
	return []domain.Movie{}, nil
}
