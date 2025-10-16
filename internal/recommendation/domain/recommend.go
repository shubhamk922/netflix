package domain

import (
	"strings"
	"sync"
)

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

type response struct {
	difference int
	movie      Movie
}

func (r TitleSimilarityEngine) Recommed2(movies []Movie, title string) ([]Movie, error) {

	titleFreqId := GenerateFreqId(title)
	maxDifference := r.acceptance
	var result []Movie
	ch := make(chan response, len(movies))
	var wg sync.WaitGroup

	for _, movie := range movies {
		wg.Add(1)
		go func(m Movie) {
			defer wg.Done()
			difference := diff(titleFreqId, movie.charFreqId)
			ch <- response{difference: difference, movie: movie}
		}(movie)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for res := range ch {
		if res.difference == maxDifference {
			result = append(result, res.movie)
		} else if res.difference < maxDifference {
			result = []Movie{}
			result = append(result, res.movie)
			maxDifference = res.difference
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
