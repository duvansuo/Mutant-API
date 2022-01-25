package controllers

import (
	// "net/http"

	"net/http"
)

type MutantController interface {
	IsMutant(w http.ResponseWriter, r *http.Request)
}
