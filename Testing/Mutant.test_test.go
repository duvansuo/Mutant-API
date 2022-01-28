package test

import (
	"testing"

	"github.com/duvansuo/Mutant-API/services"

	"github.com/stretchr/testify/assert"
)

func TestIsMutant(t *testing.T) {
	dna := []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}
	human := services.MutantService{}.NewMutant(dna)
	isMutant, _ := human.IsMutant()
	assert.Equal(t, true, isMutant)
}

func TestIsMutantNotFound(t *testing.T) {
	dna := []string{"TTGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CGCCTA", "TCACTG"}
	human := services.MutantService{}.NewMutant(dna)
	isMutant, _ := human.IsMutant()
	assert.Equal(t, false, isMutant)
}

func TestIsNotValidAdn(t *testing.T) {
	dna := []string{
		"GCATAA",
	}
	human := services.MutantService{}.NewMutant(dna)

	assert.Equal(t, false, human.ValidDimension())
}

func TestIsNotValidate(t *testing.T) {
	dna := []string{
		"GACCAA",
		"GCAGTT",
		"AAACCT",
		"GTGTAC",
		"ACATTG",
		"GAGCTTT",
	}
	human := services.MutantService{}.NewMutant(dna)

	assert.Equal(t, false, human.ValidDimension())
}
func TestIsValid(t *testing.T) {
	dna := []string{
		"GACCAA",
		"GCAGTT",
		"AAACCT",
		"GTGTAC",
		"ACATTG",
		"GAGCTT",
	}
	human := services.MutantService{}.NewMutant(dna)

	assert.Equal(t, true, human.ValidDimension())
}
