package test

import (
	"testing"

	"github.com/duvansuo/Mutant-API/services"

	data "github.com/duvansuo/Mutant-API/Data"

	"github.com/stretchr/testify/assert"
)

func TestAddHuman(t *testing.T) {
	repo := data.NewStatsTestData()
	service := services.NewStatsService(repo)
	service.AddStats(false)

	statsDTO := service.GetStats()
	assert.Equal(t, 1, statsDTO.CountHuman)
	assert.Equal(t, float32(0), statsDTO.Ratio)
	assert.Equal(t, 0, statsDTO.CountMutant)
}

func TestAddMutant(t *testing.T) {
	repo := data.NewStatsTestData()
	service := services.NewStatsService(repo)
	service.AddStats(true)

	statsDTO := service.GetStats()
	assert.Equal(t, 0, statsDTO.CountHuman)
	assert.Equal(t, float32(1), statsDTO.Ratio)
	assert.Equal(t, 1, statsDTO.CountMutant)
}

func TestAddStats2Human1Mutant(t *testing.T) {
	repo := data.NewStatsTestData()
	service := services.NewStatsService(repo)
	service.AddStats(false)
	service.AddStats(false)
	service.AddStats(true)

	statsDTO := service.GetStats()
	assert.Equal(t, 2, statsDTO.CountHuman)
	assert.Equal(t, float32(0.5), statsDTO.Ratio)
	assert.Equal(t, 1, statsDTO.CountMutant)
}
