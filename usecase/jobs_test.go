package usecase

import (
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"jobs/domain/model"
	csvservice "jobs/service/csv"
	httpservice "jobs/service/http"
	"jobs/service/mock"
)

func TestJobsUsecase_GetJobs(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCsvService := mock.NewMockNewCsvService(ctrl)
	mockCsvService.EXPECT().GetJobs().Return([]model.Job{
		{ID: 1, Title: "Software Engineer 1", NormalizedTitle: "software engineer 1"},
		{ID: 2, Title: "Software Engineer 2", NormalizedTitle: "software engineer 2"},
		{ID: 3, Title: "Software Engineer 3", NormalizedTitle: "software engineer 3"},
	}, nil)

	tests := []struct {
		name       string
		csvService csvservice.NewCsvService
		want       []model.Job
		wantErr    *error
	}{
		{
			name:       "GetJobs test",
			csvService: mockCsvService,
			want: []model.Job{
				{ID: 1, Title: "Software Engineer 1", NormalizedTitle: "software engineer 1"},
				{ID: 2, Title: "Software Engineer 2", NormalizedTitle: "software engineer 2"},
				{ID: 3, Title: "Software Engineer 3", NormalizedTitle: "software engineer 3"},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			us := &JobUsecase{
				csvService: tt.csvService,
			}
			got, _ := us.GetJobs()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JobUsecase.GetJobs() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJobUsecase_GetJobsFromAPI(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCsvService := mock.NewMockNewCsvService(ctrl)
	mockHttpService := mock.NewMockNewHTTPService(ctrl)

	mockNewJobs := []model.ExtJob{
		{Title: "3D Animator", NormalizedJobTitle: "3d animator"},
		{Title: "3D Artist", NormalizedJobTitle: "3d artist"},
		{Title: "3D Designer", NormalizedJobTitle: "3d designer"},
	}

	// we mock our methods
	mockHttpService.EXPECT().GetJobs().Return(mockNewJobs, nil)
	mockCsvService.EXPECT().StoreJobs(&mockNewJobs).Return(nil)

	tests := []struct {
		name        string
		csvService  csvservice.NewCsvService
		httpService httpservice.NewHTTPService
		want        *[]model.ExtJob
		wantErr     *error
	}{
		{
			name:        "GetJobsFromAPI test",
			csvService:  mockCsvService,
			httpService: mockHttpService,
			want:        &mockNewJobs,
			wantErr:     nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			us := &JobUsecase{
				csvService:  tt.csvService,
				httpService: tt.httpService,
			}
			got, _ := us.GetJobsFromAPI()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JobUsecase.GetJobsFromAPI() got = %v, want %v", got, tt.want)
			}
		})
	}
}
