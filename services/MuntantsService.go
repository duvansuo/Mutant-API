package services

import (
	"errors"
	"sync"
)

type mutant struct {
	Dna     []string
	counter int
	len     int
	lock    sync.Mutex
}

const (
	MAX_SEQUENCES = 2
	SEQUENCE      = 3
)

type MutantService struct {
}

func (MutantService) NewMutant(dna []string) *mutant {
	return &mutant{
		Dna: dna,
		len: len(dna[0]),
	}
}

func (mutant *mutant) ValidDimension() bool {
	Base := len(mutant.Dna)
	for i := 0; i < Base; i++ {
		if len(mutant.Dna[i]) != Base {
			return false
		}
	}
	return true
}

func (mutant *mutant) IsMutant() (bool, error) {
	if !mutant.ValidDimension() {
		return false, errors.New("dimensions do not match")
	}
	wg := sync.WaitGroup{}
	wg.Add(4)
	go func() {
		defer wg.Done()
		mutant.goOverHorizontal()
	}()
	go func() {
		defer wg.Done()
		mutant.goOverVertical()
	}()
	go func() {
		defer wg.Done()
		mutant.goOverDiagonal()
	}()
	go func() {
		defer wg.Done()
		mutant.goOverBackslash()
	}()
	wg.Wait()
	return mutant.counter >= MAX_SEQUENCES, nil
}

// horizontal
func (mutant *mutant) goOverHorizontal() {
	for i := 0; i < mutant.len; i++ {
		mutant.validateSequence([]byte(mutant.Dna[i]))
	}
}

//vertical
func (mutant *mutant) goOverVertical() {
	seq := []byte{}
	for i, _ := range mutant.Dna {
		for j, _ := range mutant.Dna {
			seq = append(seq, mutant.Dna[j][i])
		}
		mutant.validateSequence(seq)
		seq = []byte{}
	}

}

// Diagonal
func (mutant *mutant) goOverDiagonal() {
	var (
		seqH []byte
		seqV []byte
	)
	for i := 0; i < mutant.len-SEQUENCE; i++ {
		for j, _ := range mutant.Dna {
			count := i + j
			if count >= mutant.len {
				break
			}
			seqV = append(seqV, mutant.Dna[count][j])
			if i >= 1 {
				seqH = append(seqH, mutant.Dna[j][count])
			}
		}
		mutant.validateSequence(seqV)
		if seqH != nil {
			mutant.validateSequence(seqH)
		}
		seqH = []byte{}
		seqV = []byte{}
	}

}

// Diagonal inversa
func (mutant *mutant) goOverBackslash() {
	var (
		seqH []byte
		seqV []byte
	)
	var count, countB int

	len := mutant.len - 1
	for i := len; i >= 3; i-- {
		for j, _ := range mutant.Dna {
			count = i - j
			if countB+j > len {
				break
			}
			if count >= 0 {
				seqH = append(seqH, mutant.Dna[j][count])
			}
			if i < len {
				seqV = append(seqV, mutant.Dna[countB+j][len-j])
			}
		}
		count = 0
		countB++
		mutant.validateSequence(seqH)
		if seqV != nil {
			mutant.validateSequence(seqV)
		}
		seqV = []byte{}
		seqH = []byte{}
	}
}

// validar secuencia
func (mutant *mutant) validateSequence(Sequence []byte) {
	var count = 0
	for i := 1; i < len(Sequence); i++ {
		if mutant.counter >= MAX_SEQUENCES {
			return
		}
		if Sequence[i-1] == Sequence[i] {
			count++
		} else {
			count = 0
		}
		if count == SEQUENCE {
			mutant.lock.Lock()
			mutant.counter++
			mutant.lock.Unlock()
			count = 0
		}
	}
}
