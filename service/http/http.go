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

func (h *httpService) GetJobs() ([]model.ExtJob, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://api.dataatwork.org/v1/jobs", nil)
	fmt.Println(req)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "*/*")

	fmt.Println(req)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	fmt.Println(resp.Body)
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(bodyBytes))
	var response model.APIResult
	var newJobs []model.ExtJob
	json.Unmarshal(bodyBytes, &newJobs)
	response.Results = newJobs
	fmt.Println(response)
	fmt.Println(newJobs)
	return newJobs, nil
}
