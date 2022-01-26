package controllers

import (
	"Mutant-Api/services"
	"net/http"
)

type statMuxController struct {
	statsService services.StatsService
}

func NewStatMuxController(statsService services.StatsService) StatController {
	return &statMuxController{
		statsService,
	}
}
func (controller *statMuxController) MutantStats(w http.ResponseWriter, r *http.Request) {
	statsDTO := controller.statsService.GetStats()
	Response(w, http.StatusOK, statsDTO)
}
