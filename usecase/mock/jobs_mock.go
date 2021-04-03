// Code generated by MockGen. DO NOT EDIT.
// Source: jobs.go

// Package mock_usecase is a generated GoMock package.
package mock

import (
	model "jobs/domain/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockNewJobUsecase is a mock of NewJobUsecase interface.
type MockNewJobUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockNewJobUsecaseMockRecorder
}

// MockNewJobUsecaseMockRecorder is the mock recorder for MockNewJobUsecase.
type MockNewJobUsecaseMockRecorder struct {
	mock *MockNewJobUsecase
}

// NewMockNewJobUsecase creates a new mock instance.
func NewMockNewJobUsecase(ctrl *gomock.Controller) *MockNewJobUsecase {
	mock := &MockNewJobUsecase{ctrl: ctrl}
	mock.recorder = &MockNewJobUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNewJobUsecase) EXPECT() *MockNewJobUsecaseMockRecorder {
	return m.recorder
}

// GetJobs mocks base method.
func (m *MockNewJobUsecase) GetJobs() ([]model.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetJobs")
	ret0, _ := ret[0].([]model.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetJobs indicates an expected call of GetJobs.
func (mr *MockNewJobUsecaseMockRecorder) GetJobs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJobs", reflect.TypeOf((*MockNewJobUsecase)(nil).GetJobs))
}

// GetJobsFromAPI mocks base method.
func (m *MockNewJobUsecase) GetJobsFromAPI() (*[]model.ExtJob, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetJobsFromAPI")
	ret0, _ := ret[0].(*[]model.ExtJob)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetJobsFromAPI indicates an expected call of GetJobsFromAPI.
func (mr *MockNewJobUsecaseMockRecorder) GetJobsFromAPI() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJobsFromAPI", reflect.TypeOf((*MockNewJobUsecase)(nil).GetJobsFromAPI))
}
