package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"jobs/usecase"
)

// JobController struct
type JobController struct {
	useCase usecase.NewJobUsecase
}

// NewJobController inferface
type NewJobController interface {
	GetJobs(w http.ResponseWriter, r *http.Request)
	GetJobsFromAPI(w http.ResponseWriter, r *http.Request)
}

// New function
func New(juc usecase.NewJobUsecase) *JobController {
	return &JobController{juc}
}

// GetJobs function
func (jc *JobController) GetJobs(w http.ResponseWriter, r *http.Request) {
	jobs, err := jc.useCase.GetJobs()
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(404)
		fmt.Fprintf(w, "There was an unexpected error.")
		return
	}

	json.NewEncoder(w).Encode(jobs)
}

// GetJobsFromAPI function
func (jc *JobController) GetJobsFromAPI(
	w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	response, err := jc.useCase.GetJobsFromAPI()
	if err != nil {
		w.WriteHeader(404)
		fmt.Fprintf(w, "There was some errors, please try again.")
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}
