package controllers

import (
	models "Mutant-Api/Models"
	"encoding/json"
	"net/http"

	"Mutant-Api/services"
)

var mutantDTO models.MutantDto

type mutantMuxController struct {
	mutantService services.MutantService
	statService   services.StatsService
}

func NewMutantMuxController(mutantService services.MutantService, statService services.StatsService) MutantController {
	return &mutantMuxController{
		mutantService,
		statService,
	}
}

func (controller *mutantMuxController) IsMutant(w http.ResponseWriter, r *http.Request) {

	err := json.NewDecoder(r.Body).Decode(&mutantDTO)
	if err != nil {
		Response(w, http.StatusBadRequest, err)
		return
	}

	mutant := controller.mutantService.NewMutant(mutantDTO.Dna)
	isMutant, err := mutant.IsMutant()
	_ = controller.statService.AddStats(isMutant)
	if err != nil {
		Response(w, http.StatusInternalServerError, err)
		return
	}

	if err != nil {
		Response(w, http.StatusInternalServerError, err)
		return
	}
	if isMutant {
		Response(w, http.StatusOK, struct{}{})
	} else {
		Response(w, http.StatusForbidden, struct{}{})
	}
}
