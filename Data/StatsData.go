package data

import (
	models "github.com/duvansuo/Mutant-API/Models"
)

type StatsData interface {
	AddStats(statsDto models.MutantStatsDTO) error
	GetStats() (models.MutantStatsDTO, error)
}

func NewMutantStatsRepository() StatsData {
	return newStatsMongoData()
}
