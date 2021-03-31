package httpservice

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"jobs/domain/model"
)

type httpService struct{}

// NewHTTPService interface
type NewHTTPService interface {
	GetJobs() ([]model.ExtJob, error)
}

// New function
func New() *httpService {
	return &httpService{}
}

// GetJobs function
func (h *httpService) GetJobs() ([]model.ExtJob, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://api.dataatwork.org/v1/jobs", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "*/*")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var response model.APIResult
	var newJobs []model.ExtJob
	json.Unmarshal(bodyBytes, &newJobs)
	response.Results = newJobs
	return newJobs, nil
}
