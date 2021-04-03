package httpservice

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"errors"
	"jobs/domain/model"
)

type httpService struct{
	apiUrl string
}

// NewHTTPService interface
type NewHTTPService interface {
	GetJobs() ([]model.ExtJob, error)
}

// New function
func New(apiUrl string) *httpService {
	return &httpService{apiUrl: apiUrl}
}

// GetJobs function
func (h *httpService) GetJobs() ([]model.ExtJob, error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, h.apiUrl, nil)
	if err != nil {
		return nil, errors.New("There was an error instantiating the request to API")
	}
	req.Header.Add("Accept", "*/*")
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New("There was an performing the request to the API")
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("There was an error reading from the response of the aPI")
	}
	var response model.APIResult
	var newJobs []model.ExtJob
	json.Unmarshal(bodyBytes, &newJobs)
	response.Results = newJobs
	return newJobs, nil
}
