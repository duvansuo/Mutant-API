package controllers

import (
	// "net/http"

	"net/http"

	"Mutant-Api/services"
)

type MutantController interface {
	IsMutant(w http.ResponseWriter, r *http.Request)
}

func ValidateMutantController(mutantService services.MutantService) MutantController {
	return validateMutantController(mutantService)
}
