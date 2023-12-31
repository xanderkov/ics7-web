// Code generated by MockGen. DO NOT EDIT.
// Source: hospital/src/internal/modules/domain/dayPlan/service (interfaces: IDayPlanRepo)

// Package service is a generated GoMock package.
package service

import (
	context "context"
	dto "hospital/src/internal/modules/domain/dayPlan/dto"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIDayPlanRepo is a mock of IDayPlanRepo interface.
type MockIDayPlanRepo struct {
	ctrl     *gomock.Controller
	recorder *MockIDayPlanRepoMockRecorder
}

// MockIDayPlanRepoMockRecorder is the mock recorder for MockIDayPlanRepo.
type MockIDayPlanRepoMockRecorder struct {
	mock *MockIDayPlanRepo
}

// NewMockIDayPlanRepo creates a new mock instance.
func NewMockIDayPlanRepo(ctrl *gomock.Controller) *MockIDayPlanRepo {
	mock := &MockIDayPlanRepo{ctrl: ctrl}
	mock.recorder = &MockIDayPlanRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIDayPlanRepo) EXPECT() *MockIDayPlanRepoMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIDayPlanRepo) Create(arg0 context.Context, arg1 *dto.CalculateDay) (*dto.DayPlan, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*dto.DayPlan)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockIDayPlanRepoMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIDayPlanRepo)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockIDayPlanRepo) Delete(arg0 context.Context, arg1 int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockIDayPlanRepoMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIDayPlanRepo)(nil).Delete), arg0, arg1)
}

// GetById mocks base method.
func (m *MockIDayPlanRepo) GetById(arg0 context.Context, arg1 int32) (*dto.DayPlan, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", arg0, arg1)
	ret0, _ := ret[0].(*dto.DayPlan)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockIDayPlanRepoMockRecorder) GetById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockIDayPlanRepo)(nil).GetById), arg0, arg1)
}

// List mocks base method.
func (m *MockIDayPlanRepo) List(arg0 context.Context) (dto.DayPlans, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(dto.DayPlans)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockIDayPlanRepoMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockIDayPlanRepo)(nil).List), arg0)
}

// Update mocks base method.
func (m *MockIDayPlanRepo) Update(arg0 context.Context, arg1 int32, arg2 *dto.ChangeDay) (*dto.DayPlan, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(*dto.DayPlan)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockIDayPlanRepoMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIDayPlanRepo)(nil).Update), arg0, arg1, arg2)
}
