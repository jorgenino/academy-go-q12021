package usecase

import (
	"movies/domain/model"
	csvservice "movies/service/csv"
)

const pathFile = "./csv/movies.csv"

// MovieUsecase struct
type MovieUsecase struct {
	csvService csvservice.NewCsvService
}

// NewMovieUsecase interface
type NewMovieUsecase interface {
	GetMovies() ([]model.Movie, error)
}

// New function
func New(s csvservice.NewCsvService) *MovieUsecase {
	return &MovieUsecase{s}
}

// GetMovies function
func (us *MovieUsecase) GetMovies() ([]model.Movie, error) {
	f, err := us.csvService.Open(pathFile)

	if err != nil {
		return nil, err
	}
	return us.csvService.GetMovies(f)
}
