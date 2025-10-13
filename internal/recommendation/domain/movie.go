package domain

import "github.com/google/uuid"

type Movie struct {
	name     string
	id       uuid.UUID
	charFreq string
}

func NewMovie(name string) *Movie {

	return &Movie{
		name:     name,
		id:       uuid.New(),
		charFreq: generateFreqId(),
	}

}

func generateFreqId() string {

	return ""
}
