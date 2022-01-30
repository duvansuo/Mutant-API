package routing

import (
	"github.com/duvansuo/Mutant-API/controllers"

	"github.com/gorilla/mux"
)

func NewRouter(mutantController controllers.MutantController, statController controllers.StatController) *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/mutant", mutantController.IsMutant).Methods("POST")
	r.HandleFunc("/stats", statController.MutantStats).Methods("GET")

	return r
}
