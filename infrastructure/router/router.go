package router

import (
	"jobs/controller"

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
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/jobs", router.controller.GetJobs).Methods("GET")
	r.HandleFunc("/api/jobs", router.controller.GetJobsFromAPI).Methods("GET")
	return r
}
