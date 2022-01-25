package controllers

import "net/http"

type StatController interface {
	MutantStats(w http.ResponseWriter, r *http.Request)
}
