package main

import (
	"log"
	"net/http"

	"movies/controller"
	"movies/infrastructure/router"
	csvservice "movies/service/csv"
	"movies/usecase"
)

func main() {

	csvService := csvservice.New()
	usecase := usecase.New(csvService)
	controller := controller.New(usecase)

	router := router.New(controller)
	r := router.InitRouter()
	log.Fatal(http.ListenAndServe(":3000", r))
}
