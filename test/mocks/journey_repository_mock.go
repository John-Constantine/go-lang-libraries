// Code generated by MockGen. DO NOT EDIT.
// Source: ./db/repository/journey_repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	model "flow/model"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockJourneyRepository is a mock of JourneyRepository interface
type MockJourneyRepository struct {
	ctrl     *gomock.Controller
	recorder *MockJourneyRepositoryMockRecorder
}

// MockJourneyRepositoryMockRecorder is the mock recorder for MockJourneyRepository
type MockJourneyRepositoryMockRecorder struct {
	mock *MockJourneyRepository
}

// NewMockJourneyRepository creates a new mock instance
func NewMockJourneyRepository(ctrl *gomock.Controller) *MockJourneyRepository {
	mock := &MockJourneyRepository{ctrl: ctrl}
	mock.recorder = &MockJourneyRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockJourneyRepository) EXPECT() *MockJourneyRepositoryMockRecorder {
	return m.recorder
}

// FindByExternalId mocks base method
func (m *MockJourneyRepository) FindByExternalId(flowExternalId string) model.Journey {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByExternalId", flowExternalId)
	ret0, _ := ret[0].(model.Journey)
	return ret0
}

// FindByExternalId indicates an expected call of FindByExternalId
func (mr *MockJourneyRepositoryMockRecorder) FindByExternalId(flowExternalId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByExternalId", reflect.TypeOf((*MockJourneyRepository)(nil).FindByExternalId), flowExternalId)
}

// FindActiveJourneysByJourneyContext mocks base method
func (m *MockJourneyRepository) FindActiveJourneysByJourneyContext(merchantId, tenantId, channelId string) []model.Journey {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindActiveJourneysByJourneyContext", merchantId, tenantId, channelId)
	ret0, _ := ret[0].([]model.Journey)
	return ret0
}

// FindActiveJourneysByJourneyContext indicates an expected call of FindActiveJourneysByJourneyContext
func (mr *MockJourneyRepositoryMockRecorder) FindActiveJourneysByJourneyContext(merchantId, tenantId, channelId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindActiveJourneysByJourneyContext", reflect.TypeOf((*MockJourneyRepository)(nil).FindActiveJourneysByJourneyContext), merchantId, tenantId, channelId)
}
