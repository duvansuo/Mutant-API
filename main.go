package main

import (
	"Mutant-Api/controllers"
	"Mutant-Api/routing"
	"Mutant-Api/services"
	"log"
	"net/http"
	"os"
)

var (
	mutantService    = services.MutantService{}
	mutantController = controllers.ValidateMutantController(mutantService)
)

func main() {
	router := routing.NewRouter()
	server := http.ListenAndServe(":"+os.Args[1], router)

	log.Fatal(server)
}
