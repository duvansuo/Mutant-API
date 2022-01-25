package controllers

import (
	"Mutant-Api/services"
	"net/http"
)

type statMuxController struct {
	StatsService services.MutantService
}

func (controller *mutantMuxController) MutantStats(w http.ResponseWriter, r *http.Request) {
	statsDTO := controller.StatsService.GetStats()
	Response(w, statsDTO)
}
