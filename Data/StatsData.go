package data

import models "Mutant-Api/Models"

type StatsData interface {
	AddStats(statsDto models.MutantStatsDTO) error
	GetStats() (models.MutantStatsDTO, error)
}

func NewMutantStatsRepository() StatsData {
	return newStatsMongoData()
}
