package router

import (
	"movies/controller"

	"github.com/gorilla/mux"
)

// Router struct
type Router struct {
	controller controller.NewMovieController
}

// IRouter interface
type IRouter interface {
	InitRouter() *mux.Router
}

// New function
func New(c controller.NewMovieController) *Router {
	return &Router{c}
}

// InitRouter function
func (router *Router) InitRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/movies", router.controller.GetMovies).Methods("GET")

	return r
}
