package routing

import (
	"Mutant-API/controllers"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/mutant", controllers.IsMutant).Methods("POST")
	r.HandleFunc("/stats", controllers.GetStats).Methods("GET")

	return r
}
