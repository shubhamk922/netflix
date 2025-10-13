package domain

type TitleSimilarityEngine struct {
}

func NewTitleSimilarityEngine() TitleSimilarityEngine {
	return TitleSimilarityEngine{}
}

func (r TitleSimilarityEngine) Recommed() ([]Movie, error) {
	return []Movie{}, nil
}
