package csvservice

import (
	"encoding/csv"
	"errors"
	"io"
	"movies/domain/model"
	"os"
	"strconv"
)

const pathFile = "./csv/movies.csv"

// CsvService struct
type CsvService struct{}

// NewCsvService interface
type NewCsvService interface {
	GetMovies(f *os.File) ([]model.Movie, error)
	Open(path string) (*os.File, error)
}

// New function
func New() *CsvService {
	return &CsvService{}
}

// GetMovies function
func (s *CsvService) GetMovies(f *os.File) ([]model.Movie, error) {

	movies, err := Read(f)

	if err != nil {
		return nil, err
	}

	return movies, nil
}

// Read function
func Read(f *os.File) ([]model.Movie, error) {

	reader := csv.NewReader(f)
	reader.Comma = ','
	reader.Comment = '#'
	reader.FieldsPerRecord = -1

	var movies []model.Movie = nil
	for {
		line, err := reader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}
		tempMovie := model.Movie{
			Title:    line[1],
			Director: line[2],
		}

		if line[0] != "" {
			id, err := strconv.Atoi(line[0])
			if err != nil {
				return nil, err
			}
			tempMovie.ID = id
		}

		movies = append(movies, tempMovie)
	}
	defer f.Close()

	return movies, nil
}

// Open function
func (s *CsvService) Open(path string) (*os.File, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, errors.New("There was an error opening the file")
	}
	return f, nil
}
