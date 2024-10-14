// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/NishimuraTakuya-nt/go-rest-clean-plane-wire/internal/infrastructure/auth (interfaces: TokenService)

// Package mockauth is a generated GoMock package.
package mockauth

import (
	reflect "reflect"

	auth "github.com/NishimuraTakuya-nt/go-rest-clean-plane-wire/internal/infrastructure/auth"
	gomock "github.com/golang/mock/gomock"
)

// MockTokenService is a mock of TokenService interface.
type MockTokenService struct {
	ctrl     *gomock.Controller
	recorder *MockTokenServiceMockRecorder
}

// MockTokenServiceMockRecorder is the mock recorder for MockTokenService.
type MockTokenServiceMockRecorder struct {
	mock *MockTokenService
}

// NewMockTokenService creates a new mock instance.
func NewMockTokenService(ctrl *gomock.Controller) *MockTokenService {
	mock := &MockTokenService{ctrl: ctrl}
	mock.recorder = &MockTokenServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTokenService) EXPECT() *MockTokenServiceMockRecorder {
	return m.recorder
}

// GenerateToken mocks base method.
func (m *MockTokenService) GenerateToken(arg0 string, arg1 []string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateToken", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateToken indicates an expected call of GenerateToken.
func (mr *MockTokenServiceMockRecorder) GenerateToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateToken", reflect.TypeOf((*MockTokenService)(nil).GenerateToken), arg0, arg1)
}

// ValidateToken mocks base method.
func (m *MockTokenService) ValidateToken(arg0 string) (*auth.Claims, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateToken", arg0)
	ret0, _ := ret[0].(*auth.Claims)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateToken indicates an expected call of ValidateToken.
func (mr *MockTokenServiceMockRecorder) ValidateToken(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateToken", reflect.TypeOf((*MockTokenService)(nil).ValidateToken), arg0)
}
