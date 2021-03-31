package router

import (
	"jobs/controller"
	"net/http"
	"github.com/gorilla/mux"
)

// Router struct
type Router struct {
	controller controller.NewJobController
}

// IRouter interface
type IRouter interface {
	InitRouter() *mux.Router
}

// New function
func New(c controller.NewJobController) *Router {
	return &Router{c}
}

// InitRouter function
func (router *Router) InitRouter() *mux.Router {
	r := mux.NewRouter()
	r.Methods(http.MethodGet).
		Path("/concurrency/jobs/{type}").
		Queries("items_per_worker", "{[0-9]+}").
		Queries("items", "{[0-9]+}").
		HandlerFunc(router.controller.GetJobsConcurrently)
	r.Methods(http.MethodGet).
		Path("/concurrency/jobs/{type}").
		Queries("items", "{[0-9]+}").
		HandlerFunc(router.controller.GetJobsConcurrently)
	r.HandleFunc("/jobs", router.controller.GetJobs).Methods("GET")
	r.HandleFunc("/api/jobs", router.controller.GetJobsFromAPI).Methods("GET")
	return r
}
