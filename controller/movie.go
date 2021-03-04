package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"movies/usecase"
)

// MovieController struct
type MovieController struct {
	useCase usecase.NewMovieUsecase
}

// NewMovieController inferface
type NewMovieController interface {
	GetMovies(w http.ResponseWriter, r *http.Request)
}

// New function
func New(pc usecase.NewMovieUsecase) *MovieController {
	return &MovieController{pc}
}

// GetMovies function
func (mc *MovieController) GetMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := mc.useCase.GetMovies()
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(404)
		fmt.Fprintf(w, "There was an unexpected error.")
		return
	}

	json.NewEncoder(w).Encode(movies)
}
