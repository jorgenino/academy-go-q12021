package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"jobs/usecase"
	"net/http"
	"strconv"
)

// JobController struct
type JobController struct {
	useCase usecase.NewJobUsecase
}

// NewJobController inferface
type NewJobController interface {
	GetJobs(w http.ResponseWriter, r *http.Request)
	GetJobsFromAPI(w http.ResponseWriter, r *http.Request)
	GetJobsConcurrently(w http.ResponseWriter, r *http.Request)
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

// GetJobsConcurrently function
func (jbc *JobController) GetJobsConcurrently(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	typeNumber := vars["type"]
	if typeNumber == "even" || typeNumber == "odd" {
		itemsS := r.FormValue("items")
		itemsPerWorkerS := r.FormValue("items_per_worker")
		items, _ := strconv.Atoi(r.FormValue("items"))
		itemsPerWorker, _ := strconv.Atoi(r.FormValue("items_per_worker"))
		jobs, _ := jbc.useCase.GetJobsConcurrently(typeNumber, items, itemsPerWorker)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(typeNumber + " " + itemsS + " " + itemsPerWorkerS)
		json.NewEncoder(w).Encode(&jobs)
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{ "message": "You only can use "even" or "odd"" }`)
	}
}
