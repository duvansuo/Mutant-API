package main

import (
	"Mutant-Api/controllers"
	"Mutant-Api/routing"
	"Mutant-Api/services"
	"log"
	"net/http"

	data "github.com/duvansuo/Mutant-API/Data"
)

var (
	mutantService    = services.MutantService{}
	mutantStatData   = data.NewMutantStatsRepository()
	statService      = services.NewStatsService(mutantStatData)
	mutantController = controllers.NewMutantMuxController(mutantService, statService)
	statController   = controllers.NewStatMuxController(statService)
)

func main() {
	router := routing.NewRouter(mutantController, statController)
	server := http.ListenAndServe(":8080", router)

	log.Fatal(server)
}
