package services

import (
	data "github.com/duvansuo/Mutant-API/Data"
	models "github.com/duvansuo/Mutant-API/Models"
)

type StatsService struct {
	repo data.StatsData
}

func NewStatsService(repo data.StatsData) StatsService {
	return StatsService{
		repo,
	}
}

func (service *StatsService) AddStats(isMutant bool) error {
	statsDTO := service.GetStats()
	statsDTO = calculateStats(statsDTO, isMutant)
	err := service.repo.AddStats(statsDTO)
	if err != nil {
		return err
	}
	return nil
}
func (service *StatsService) GetStats() models.MutantStatsDTO {
	statsDTO, err := service.repo.GetStats()
	if err != nil {
		panic(err)

	}
	return statsDTO
}

func calculateStats(statsDTO models.MutantStatsDTO, isMutant bool) models.MutantStatsDTO {

	if isMutant {
		statsDTO.CountMutant += 1
	} else {
		statsDTO.CountHuman += 1
	}
	if statsDTO.CountHuman == 0 {
		statsDTO.Ratio = 1
	} else {
		statsDTO.Ratio = float32(statsDTO.CountMutant) / float32(statsDTO.CountHuman)
	}
	return statsDTO
}
