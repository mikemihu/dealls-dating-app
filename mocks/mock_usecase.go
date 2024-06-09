// Code generated by MockGen. DO NOT EDIT.
// Source: ../internal/usecase.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	entity "dealls-dating-app/internal/entity"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockUserUC is a mock of UserUC interface.
type MockUserUC struct {
	ctrl     *gomock.Controller
	recorder *MockUserUCMockRecorder
}

// MockUserUCMockRecorder is the mock recorder for MockUserUC.
type MockUserUCMockRecorder struct {
	mock *MockUserUC
}

// NewMockUserUC creates a new mock instance.
func NewMockUserUC(ctrl *gomock.Controller) *MockUserUC {
	mock := &MockUserUC{ctrl: ctrl}
	mock.recorder = &MockUserUCMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserUC) EXPECT() *MockUserUCMockRecorder {
	return m.recorder
}

// BuyPackage mocks base method.
func (m *MockUserUC) BuyPackage(ctx context.Context, req entity.BuyPackageRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuyPackage", ctx, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// BuyPackage indicates an expected call of BuyPackage.
func (mr *MockUserUCMockRecorder) BuyPackage(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuyPackage", reflect.TypeOf((*MockUserUC)(nil).BuyPackage), ctx, req)
}

// Get mocks base method.
func (m *MockUserUC) Get(ctx context.Context, id uuid.UUID) (entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, id)
	ret0, _ := ret[0].(entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockUserUCMockRecorder) Get(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockUserUC)(nil).Get), ctx, id)
}

// GetVerifiedStatus mocks base method.
func (m *MockUserUC) GetVerifiedStatus(ctx context.Context, id uuid.UUID) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVerifiedStatus", ctx, id)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVerifiedStatus indicates an expected call of GetVerifiedStatus.
func (mr *MockUserUCMockRecorder) GetVerifiedStatus(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVerifiedStatus", reflect.TypeOf((*MockUserUC)(nil).GetVerifiedStatus), ctx, id)
}

// Login mocks base method.
func (m *MockUserUC) Login(ctx context.Context, email, password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", ctx, email, password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockUserUCMockRecorder) Login(ctx, email, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockUserUC)(nil).Login), ctx, email, password)
}

// Register mocks base method.
func (m *MockUserUC) Register(ctx context.Context, email, password string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", ctx, email, password)
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register.
func (mr *MockUserUCMockRecorder) Register(ctx, email, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockUserUC)(nil).Register), ctx, email, password)
}

// MockProfileUC is a mock of ProfileUC interface.
type MockProfileUC struct {
	ctrl     *gomock.Controller
	recorder *MockProfileUCMockRecorder
}

// MockProfileUCMockRecorder is the mock recorder for MockProfileUC.
type MockProfileUCMockRecorder struct {
	mock *MockProfileUC
}

// NewMockProfileUC creates a new mock instance.
func NewMockProfileUC(ctrl *gomock.Controller) *MockProfileUC {
	mock := &MockProfileUC{ctrl: ctrl}
	mock.recorder = &MockProfileUCMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProfileUC) EXPECT() *MockProfileUCMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockProfileUC) Get(ctx context.Context, filter entity.ProfileFilter) (entity.Profile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, filter)
	ret0, _ := ret[0].(entity.Profile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockProfileUCMockRecorder) Get(ctx, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockProfileUC)(nil).Get), ctx, filter)
}

// GetList mocks base method.
func (m *MockProfileUC) GetList(ctx context.Context, filter entity.ProfileFilter) ([]entity.Profile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetList", ctx, filter)
	ret0, _ := ret[0].([]entity.Profile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetList indicates an expected call of GetList.
func (mr *MockProfileUCMockRecorder) GetList(ctx, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetList", reflect.TypeOf((*MockProfileUC)(nil).GetList), ctx, filter)
}

// Update mocks base method.
func (m *MockProfileUC) Update(ctx context.Context, req entity.ProfileUpdateRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockProfileUCMockRecorder) Update(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockProfileUC)(nil).Update), ctx, req)
}

// MockLikeUC is a mock of LikeUC interface.
type MockLikeUC struct {
	ctrl     *gomock.Controller
	recorder *MockLikeUCMockRecorder
}

// MockLikeUCMockRecorder is the mock recorder for MockLikeUC.
type MockLikeUCMockRecorder struct {
	mock *MockLikeUC
}

// NewMockLikeUC creates a new mock instance.
func NewMockLikeUC(ctrl *gomock.Controller) *MockLikeUC {
	mock := &MockLikeUC{ctrl: ctrl}
	mock.recorder = &MockLikeUCMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLikeUC) EXPECT() *MockLikeUCMockRecorder {
	return m.recorder
}

// MutualList mocks base method.
func (m *MockLikeUC) MutualList(ctx context.Context, req entity.MutualListRequest) (entity.MutualListResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MutualList", ctx, req)
	ret0, _ := ret[0].(entity.MutualListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MutualList indicates an expected call of MutualList.
func (mr *MockLikeUCMockRecorder) MutualList(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MutualList", reflect.TypeOf((*MockLikeUC)(nil).MutualList), ctx, req)
}

// MockSwipeUC is a mock of SwipeUC interface.
type MockSwipeUC struct {
	ctrl     *gomock.Controller
	recorder *MockSwipeUCMockRecorder
}

// MockSwipeUCMockRecorder is the mock recorder for MockSwipeUC.
type MockSwipeUCMockRecorder struct {
	mock *MockSwipeUC
}

// NewMockSwipeUC creates a new mock instance.
func NewMockSwipeUC(ctrl *gomock.Controller) *MockSwipeUC {
	mock := &MockSwipeUC{ctrl: ctrl}
	mock.recorder = &MockSwipeUCMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSwipeUC) EXPECT() *MockSwipeUCMockRecorder {
	return m.recorder
}

// Like mocks base method.
func (m *MockSwipeUC) Like(ctx context.Context, req entity.SwipeRequest) (entity.SwipeResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Like", ctx, req)
	ret0, _ := ret[0].(entity.SwipeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Like indicates an expected call of Like.
func (mr *MockSwipeUCMockRecorder) Like(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Like", reflect.TypeOf((*MockSwipeUC)(nil).Like), ctx, req)
}

// NextProfile mocks base method.
func (m *MockSwipeUC) NextProfile(ctx context.Context, req entity.NextProfileRequest) (entity.NextProfileResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NextProfile", ctx, req)
	ret0, _ := ret[0].(entity.NextProfileResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NextProfile indicates an expected call of NextProfile.
func (mr *MockSwipeUCMockRecorder) NextProfile(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NextProfile", reflect.TypeOf((*MockSwipeUC)(nil).NextProfile), ctx, req)
}