package usecase

import (
	"reflect"
	"testing"
	"jobs/domain/model"
	csvservice "jobs/service/csv"
	httpservice "jobs/service/http"
	"jobs/service/mock"
	
	"github.com/golang/mock/gomock"
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

func TestJobUsecase_GetJobsConcurrently(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCsvService := mock.NewMockNewCsvService(ctrl)

	mockCsvService.EXPECT().GetJobs().Return([]model.Job{
		{ID: 1, Title: "Software Engineer 1", NormalizedTitle: "software engineer 1"},
		{ID: 2, Title: "Software Engineer 2", NormalizedTitle: "software engineer 2"},
		{ID: 3, Title: "Software Engineer 3", NormalizedTitle: "software engineer 3"},
	}, nil)

	tests := []struct {
		name           string
		csvService     csvservice.NewCsvService
		typeNumber     string
		items          int
		itemsPerWorker int
		want           []model.Job
		wantErr        error
	}{
		{
			name:           "Concurrent Jobs test: odd, items: 1, itemsPerWorker: 1",
			csvService:     mockCsvService,
			typeNumber:     "odd",
			items:          1,
			itemsPerWorker: 1,
			want: []model.Job{
				{
					ID:   1,
					Title: "Software Engineer 1",
					NormalizedTitle:  "software engineer 1",
				},
			},
			wantErr: nil,
		},
		{
			name:           "Concurrent Jobs test: even, items: 3, itemsPerWorker: 2",
			csvService:     mockCsvService,
			typeNumber:     "even",
			items:          3,
			itemsPerWorker: 2,
			want: []model.Job{
				{ID: 2, Title: "Software Engineer 2", NormalizedTitle: "software engineer 2"},
				{ID: 4, Title: "1st Pressman", NormalizedTitle: "1st pressman"},
				{ID: 6, Title: "21 Dealer", NormalizedTitle: "21 dealer"},
			},
			wantErr: nil,
		},
		{
			name:           "Concurrent Jobs test: even, items: 10, itemsPerWorker: 3",
			csvService:     mockCsvService,
			typeNumber:     "odd",
			items:          10,
			itemsPerWorker: 3,
			want: []model.Job{
				{ID: 1, Title: "Software Engineer 1", NormalizedTitle: "software engineer 1"},
				{ID: 3, Title: "Software Engineer 3", NormalizedTitle: "software engineer 1"},
				{ID: 7, Title: "2nd Grade Teacher", NormalizedTitle: "2nd grade teacher"},
				{ID: 9, Title: "2 Year Olds Preschool Teacher", NormalizedTitle: "2 year olds preschool teacher"},
				{ID: 5, Title: "1st Pressman On Web Press", NormalizedTitle: "1st pressman on web press"},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			us := &JobUsecase{
				csvService: tt.csvService,
			}
			got, gotErr := us.GetJobsConcurrently(tt.typeNumber, tt.items, tt.itemsPerWorker)
			assert.ElementsMatch(t, got, tt.want)
			assert.Equal(t, gotErr, tt.wantErr)
		})
	}
}