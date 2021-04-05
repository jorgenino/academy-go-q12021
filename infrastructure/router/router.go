package router

import (
	"github.com/gorilla/mux"
	"jobs/controller"
	"net/http"
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
// Initiates the Router object
func New(c controller.NewJobController) *Router {
	return &Router{c}
}

// InitRouter function
// Initiates the Router object endpoints 
func (router *Router) InitRouter() *mux.Router {
	r := mux.NewRouter()
	r.Methods(http.MethodGet).
		Path("/concurrency/jobs").
		Queries("type", "{[a-z]+}").
		Queries("items_per_worker", "{[0-9]+}").
		Queries("items", "{[0-9]+}").
		HandlerFunc(router.controller.GetJobsConcurrently)
	r.Methods(http.MethodGet).
		Path("/concurrency/jobs").
		Queries("type", "{[a-z]+}").
		Queries("items", "{[0-9]+}").
		HandlerFunc(router.controller.GetJobsConcurrently)
	r.HandleFunc("/jobs", router.controller.GetJobs).Methods(http.MethodGet)
	r.HandleFunc("/api/jobs", router.controller.GetJobsFromAPI).Methods(http.MethodGet)
	return r
}
