package domain

import "strings"

type TitleSimilarityEngine struct {
	acceptance int
}

func NewTitleSimilarityEngine(acceptance int) TitleSimilarityEngine {
	return TitleSimilarityEngine{
		acceptance: acceptance,
	}
}

func (r TitleSimilarityEngine) Recommed(movies []Movie, title string) ([]Movie, error) {

	titleFreqId := GenerateFreqId(title)
	maxDifference := r.acceptance
	var result []Movie

	for _, movie := range movies {
		difference := diff(titleFreqId, movie.charFreqId)
		if difference == maxDifference {
			result = append(result, movie)
		} else if difference < maxDifference {
			result = []Movie{}
			result = append(result, movie)
			maxDifference = difference
		}
	}

	return result, nil
}

func diff(titleFreqId, movieId string) int {

	freq1 := strings.Split(titleFreqId, "#")
	freq2 := strings.Split(movieId, "#")
	count := 0
	for i := 0; i < 26; i++ {
		count = count + abs(int(freq1[i][0]-'a'), int(freq2[i][0]-'a'))
	}
	return count
}

func abs(i, j int) int {
	if i > j {
		return i - j
	}
	return j - i
}
