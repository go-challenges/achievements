// Code generated by MockGen. DO NOT EDIT.
// Source: ./src/repository/challenge.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	models "achievements/src/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockChallengesRepository is a mock of ChallengesRepository interface.
type MockChallengesRepository struct {
	ctrl     *gomock.Controller
	recorder *MockChallengesRepositoryMockRecorder
}

// MockChallengesRepositoryMockRecorder is the mock recorder for MockChallengesRepository.
type MockChallengesRepositoryMockRecorder struct {
	mock *MockChallengesRepository
}

// NewMockChallengesRepository creates a new mock instance.
func NewMockChallengesRepository(ctrl *gomock.Controller) *MockChallengesRepository {
	mock := &MockChallengesRepository{ctrl: ctrl}
	mock.recorder = &MockChallengesRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChallengesRepository) EXPECT() *MockChallengesRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockChallengesRepository) Create(arg0 *models.Challenge) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockChallengesRepositoryMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockChallengesRepository)(nil).Create), arg0)
}
