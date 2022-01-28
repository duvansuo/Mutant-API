package data

import models "github.com/duvansuo/Mutant-API/Models"

type StatsTestData struct {
	Err            error
	MutantStatsDTO models.MutantStatsDTO
}

func NewStatsTestData() *StatsTestData {
	return &StatsTestData{
		Err: nil,
	}
}

func (data *StatsTestData) AddStats(statsDto models.MutantStatsDTO) error {
	data.MutantStatsDTO = statsDto
	return data.Err
}
func (data *StatsTestData) GetStats() (models.MutantStatsDTO, error) {
	return data.MutantStatsDTO, data.Err
}
