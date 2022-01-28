package main

import (
	"log"
	"net/http"

	"github.com/duvansuo/Mutant-API/services"

	"github.com/duvansuo/Mutant-API/routing"

	"github.com/duvansuo/Mutant-API/controllers"

	data "github.com/duvansuo/Mutant-API/Data"
)

var (
	mutantService    = services.MutantService{}
	mutantStatData   = data.NewMutantStatsData()
	statService      = services.NewStatsService(mutantStatData)
	mutantController = controllers.NewMutantMuxController(mutantService, statService)
	statController   = controllers.NewStatMuxController(statService)
)

func main() {
	router := routing.NewRouter(mutantController, statController)
	server := http.ListenAndServe(":8080", router)

	log.Fatal(server)
}
