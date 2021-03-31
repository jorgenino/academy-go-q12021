package usecase

import (
	"math"
	"sync"
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
	GetJobs() ([]model.Job, error) // <host>/jobs
	GetJobsFromAPI() (*[]model.ExtJob, error) // <host>/api/jobs
	GetJobsConcurrently(typeNumber string, items int, itemsPerWorker int) ([]model.Job, error) // <host>/concurrency/jobs/{type}
}

// New function
func New(s csvservice.NewCsvService, h httpservice.NewHTTPService) *JobUsecase {
	return &JobUsecase{s, h}
}

// GetJobs function
func (us *JobUsecase) GetJobs() ([]model.Job, error) {
	return us.csvService.GetJobs()
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

// calculatePoolSize function
func calculatePoolSize(items int, itemsPerWorker int, totalJobs int) int {
	var poolSize int
	if items%itemsPerWorker != 0 {
		poolSize = int(math.Ceil(float64(items) / float64(itemsPerWorker)))
	} else {
		poolSize = int(items / itemsPerWorker)
	}

	// If we overpass the number of workers above the half of number
	// of items it's gonna get into an infinit looop
	if poolSize > (totalJobs / 2) {
		poolSize = totalJobs / 2
	}
	return poolSize
}

// calculateMaxJobs function
func calculateMaxJobs(totalJobs int) int {
	var maxJobs int

	if totalJobs%2 == 0 {
		maxJobs = totalJobs / 2
	} else {
		maxJobs = totalJobs/2 + 1
	}
	return maxJobs
}

// GetJobsConcurrently function
func (us *JobUsecase) GetJobsConcurrently(typeNumber string, items int, itemsPerWorker int) ([]model.Job, error) {
	jobs, err := us.csvService.GetJobs()
	if err != nil {
		return nil, err
	}
	totalJobs := len(jobs)
	poolSize := calculatePoolSize(items, itemsPerWorker, totalJobs)
	maxJobs := calculateMaxJobs(totalJobs)
	values := make(chan int)
	workerJobs := make(chan int, poolSize)
	shutdown := make(chan struct{})
	startIndex := 0
	var limit int
	limit = int(math.Ceil(float64(totalJobs) / float64(poolSize)))
	lastLimit := (totalJobs % limit)
	var wg sync.WaitGroup
	wg.Add(poolSize)
	for i := 0; i < poolSize; i++ {
		go func(workerJobs <-chan int) {
			for {
				var id int
				var limitRecalculated int
				start := <-workerJobs

				// We do need to iterate with the same limit every time.
				// on the last cycle we use the leftovers of the division (modulus)
				if limit+start >= totalJobs && lastLimit != 0 { // lastLimit can be 0, take care of that
					limitRecalculated = start + lastLimit
				} else {
					limitRecalculated = start + limit
				}

				for j := start; j < limitRecalculated; j++ {
					id = jobs[j].ID

					select {
					case values <- id:
					case <-shutdown:
						wg.Done()
						return
					}
				}
			}
		}(workerJobs)
	}
	for i := 0; i < poolSize; i++ {
		workerJobs <- startIndex
		startIndex += limit
	}
	close(workerJobs)
	var filteredJobs []model.Job = nil
	bucket := make(map[int]int, totalJobs+1)
	for elem := range values {
		if typeNumber == "even" {
			if elem%2 == 0 && bucket[elem] == 0 {
				filteredJobs = append(filteredJobs, jobs[elem-1])
				bucket[elem] = elem
			}
		} else if typeNumber == "odd" {
			if elem % 2 != 0 && bucket[elem] == 0 {
				filteredJobs = append(filteredJobs, jobs[elem-1])
				bucket[elem] = elem 
			}
		}
		if len(filteredJobs) >= items || len(filteredJobs) >= maxJobs {
			break
		}
	}
	close(shutdown)
	wg.Wait()
	return filteredJobs, nil
}
