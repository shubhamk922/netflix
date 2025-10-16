package domain

import (
	"strconv"
	"strings"

	"github.com/google/uuid"
)

type Movie struct {
	name       string
	id         uuid.UUID
	charFreqId string
}

func NewMovie(name string) *Movie {

	return &Movie{
		name:       name,
		id:         uuid.New(),
		charFreqId: GenerateFreqId(name),
	}

}

func (movie *Movie) FreqId() string {
	return movie.charFreqId
}

func GenerateFreqId(title string) string {

	title = strings.ToLower(title)
	charFreq := make([]int, 26)
	for i := 0; i < 26; i++ {
		charFreq[i] = 0
	}

	for i := 0; i < len(title); i++ {
		if title[i] >= 'a' && title[i] <= 'z' {
			charFreq[int(title[i]-'a')]++
		}
	}

	var titleId string
	for i := 0; i < len(charFreq); i++ {
		n := strconv.Itoa(charFreq[i])
		titleId = titleId + n + "#"
	}
	return titleId
}
