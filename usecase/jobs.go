package usecase

import (
	"jobs/domain/model"
	csvservice "jobs/service/csv"
	httpservice "jobs/service/http"
)

const pathFile = "./csv/jobs.csv"

// JobUsecase struct
type JobUsecase struct {
	csvService  csvservice.NewCsvService
	httpService httpservice.NewHTTPService
}

// NewJobUsecase interface
type NewJobUsecase interface {
	GetJobs() ([]model.Job, error)
	GetJobsFromAPI() (*[]model.ExtJob, error)
}

// New function
func New(s csvservice.NewCsvService, h httpservice.NewHTTPService) *JobUsecase {
	return &JobUsecase{s, h}
}

// GetJobs function
func (us *JobUsecase) GetJobs() ([]model.Job, error) {
	f, err := us.csvService.Open(pathFile)

	if err != nil {
		return nil, err
	}
	return us.csvService.GetJobs(f)
}

// GetJobsFromAPI function
func (us *JobUsecase) GetJobsFromAPI() (*[]model.ExtJob, error) {
	newJobs, err := us.httpService.GetJobs()

	if err != nil {
		return nil, err
	}

	errorCsv := us.csvService.StoreJobs(&newJobs)

	if errorCsv != nil {
		return nil, errorCsv
	}

	return &newJobs, nil
}
