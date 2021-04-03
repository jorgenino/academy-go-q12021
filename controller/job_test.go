package controller

import (
	"jobs/domain/model"
	"jobs/usecase"
	usecasemock "jobs/usecase/mock"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

// TestJobController_GetJobs function
func TestJobController_GetJobs(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	jobs := []model.Job{
		{ID: 1, Title: "Software Engineer 1", NormalizedTitle: "software engineer 1"},
		{ID: 2, Title: "Software Engineer 2", NormalizedTitle: "software engineer 2"},
		{ID: 3, Title: "Software Engineer 3", NormalizedTitle: "software engineer 3"},
	}

	mockJobUsecase := usecasemock.NewMockNewJobUsecase(ctrl)
	mockJobUsecase.EXPECT().GetJobs().Return(jobs, nil)

	request := httptest.NewRequest("GET", "/jobs", nil)
	recorder := httptest.NewRecorder()

	tests := []struct {
		name    string
		useCase usecase.NewJobUsecase
		r       *http.Request
		rr      *httptest.ResponseRecorder
		want    []model.Job
	}{
		{
			name:    "Get jobs test",
			useCase: mockJobUsecase,
			r:       request,
			rr:      recorder,
			want:    jobs,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jc := &JobController{
				useCase: tt.useCase,
			}
			handler := http.HandlerFunc(jc.GetJobs)
			handler.ServeHTTP(tt.rr, tt.r)

			if status := tt.rr.Code; status != http.StatusOK {
				t.Fatalf("handler returned wrong status code: got %v want %v",
					status, http.StatusOK)
			}

			reflect.DeepEqual(tt.rr.Body, tt.want)
		})
	}
}

// TestJobController_GetJobsFromAPI function
func TestJobController_GetJobsFromAPI(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	newJobs := &[]model.ExtJob{
		{Title: "3D Animator", NormalizedJobTitle: "3d animator"},
		{Title: "3D Artist", NormalizedJobTitle: "3d artist"},
		{Title: "3D Designer", NormalizedJobTitle: "3d designer"},
	}
	mockUsecaseJob := usecasemock.NewMockNewJobUsecase(ctrl)
	mockUsecaseJob.EXPECT().GetJobsFromAPI().Return(newJobs, nil)
	request := httptest.NewRequest("GET", "/api/jobs", nil)
	recorder := httptest.NewRecorder()

	tests := []struct {
		name    string
		useCase usecase.NewJobUsecase
		r       *http.Request
		rr      *httptest.ResponseRecorder
		want    *[]model.ExtJob
	}{
		{
			name:    "Get Jobs from API",
			useCase: mockUsecaseJob,
			r:       request,
			rr:      recorder,
			want:    newJobs,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jc := &JobController{
				useCase: tt.useCase,
			}
			jc.GetJobsFromAPI(tt.rr, tt.r)
			handler := http.HandlerFunc(jc.GetJobs)
			handler(tt.rr, tt.r)

			if status := tt.rr.Code; status != http.StatusOK {
				t.Fatalf("handler returned wrong status code: got %v want %v",
					status, http.StatusOK)
			}
			reflect.DeepEqual(tt.rr.Body, tt.want)
		})
	}
}
