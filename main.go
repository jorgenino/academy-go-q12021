package main

import (
	"log"
	"net/http"

	"jobs/controller"
	"jobs/infrastructure/router"
	csvservice "jobs/service/csv"
	httpservice "jobs/service/http"
	"jobs/usecase"
)

func main() {

	csvService := csvservice.New()
	httpService := httpservice.New()
	usecase := usecase.New(csvService, httpService)
	controller := controller.New(usecase)

	router := router.New(controller)
	r := router.InitRouter()
	log.Fatal(http.ListenAndServe(":3000", r))
}
