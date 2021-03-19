package csvservice

import (
	"encoding/csv"
	"errors"
	"io"
	"jobs/domain/model"
	"os"
	"strconv"
)

const pathFile = "./csv/jobs.csv"

// CsvService struct
type CsvService struct{}

// NewCsvService interface
type NewCsvService interface {
	GetJobs(f *os.File) ([]model.Job, error)
	Open(path string) (*os.File, error)
	StoreJobs(*[]model.ExtJob) error
}

// New function
func New() *CsvService {
	return &CsvService{}
}

// GetJobs function
func (s *CsvService) GetJobs(f *os.File) ([]model.Job, error) {

	jobs, err := Read(f)

	if err != nil {
		return nil, err
	}

	return jobs, nil
}

// Read function
func Read(f *os.File) ([]model.Job, error) {

	reader := csv.NewReader(f)
	reader.Comma = ','
	reader.Comment = '#'
	reader.FieldsPerRecord = -1

	var jobs []model.Job = nil
	for {
		line, err := reader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}
		tempJob := model.Job{
			Title:           line[1],
			NormalizedTitle: line[2],
		}

		if line[0] != "" {
			id, err := strconv.Atoi(line[0])
			if err != nil {
				return nil, err
			}
			tempJob.ID = id
		}

		jobs = append(jobs, tempJob)
	}
	defer f.Close()

	return jobs, nil
}

// Open function
func (s *CsvService) Open(path string) (*os.File, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, errors.New("There was an error opening the file")
	}
	return f, nil
}

// Open function
func Open(path string) (*os.File, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, errors.New("There was an error opening the file")
	}
	return f, nil
}

// ReadAllLines function
func ReadAllLines(f *os.File) ([][]string, error) {
	reader := csv.NewReader(f)
	reader.Comma = ','
	reader.Comment = '#'
	reader.FieldsPerRecord = -1
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	defer f.Close()

	return lines, nil
}

// AddLine function
func AddLine(f *os.File, lines [][]string, newJobs *[]model.ExtJob) error {

	linesNumber := len(lines) + 1

	w := csv.NewWriter(f)
	for _, job := range *newJobs {
		w.Write([]string{strconv.Itoa(linesNumber), job.Title, job.NormalizedJobTitle})
		linesNumber = linesNumber + 1
	}
	defer w.Flush()

	return nil
}

// OpenAndWrite function
func OpenAndWrite(path string) (*os.File, error) {
	f, err := os.OpenFile(path, os.O_RDONLY|os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return nil, errors.New("There was an error opening the file")
	}
	return f, nil
}

// StoreJobs function
func (s *CsvService) StoreJobs(newJobs *[]model.ExtJob) error {
	f, _ := Open(pathFile) //Read only
	lines, _ := ReadAllLines(f)
	fileOpenAndWrite, _ := OpenAndWrite(pathFile) // Write

	err := AddLine(fileOpenAndWrite, lines, newJobs)
	if err != nil {
		return err
	}

	return nil
}
