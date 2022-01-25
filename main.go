package main

import (
	"Mutant-Api/controllers"
	"Mutant-Api/routing"
	"Mutant-Api/services"
	"log"
	"net/http"
)

var (
	mutantService    = services.MutantService{}
	mutantController = controllers.NewMutantMuxController(mutantService)
)

func main() {
	router := routing.NewRouter(mutantController)
	server := http.ListenAndServe(":8080", router)

	log.Fatal(server)
}
