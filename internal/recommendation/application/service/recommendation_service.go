package service

import (
	"context"
	"fmt"

	"example.com/netflix/internal/recommendation/application/port/out"
	"example.com/netflix/internal/recommendation/domain"
)

const LEN_RECOMMENDED_MOVIES = 3

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

	titleFreqId := domain.GenerateFreqId(title)

	movies, err := r.repo.GetAllMoviesByFrqId(ctx, titleFreqId)

	if err != nil {
		return []domain.Movie{}, fmt.Errorf("Error in finding similar movies from db %s", err.Error())
	}

	if len(movies) < LEN_RECOMMENDED_MOVIES {

		similar_movies, err := r.engine.Recommed(movies, title)
		if err != nil {
			return movies, nil
		}

		if len(similar_movies) > len(movies) {
			return similar_movies, nil
		}

	}

	return movies, nil
}
