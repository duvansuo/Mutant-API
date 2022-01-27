package services

import (
	"errors"
	"sync"
)

type mutant struct {
	Adn         []string
	counter     int
	lockCounter sync.Mutex
}

const (
	MAX_SUBSEQUENCES = 2
	SIZE_SUBSEQUENCE = 4
)

type MutantService struct {
}

func (MutantService) NewMutant(dna []string) *mutant {
	return &mutant{
		Adn: dna,
	}
}

func (mutant *mutant) validDimension() bool {
	return len(mutant.Adn) == len(mutant.Adn[0])
}

func (mutant *mutant) IsMutant() (bool, error) {

	if !mutant.validDimension() {
		return false, errors.New("dimensions do not match")
	}
	return true, nil
}
