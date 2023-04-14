// Code generated by MockGen. DO NOT EDIT.
// Source: instance.go

// Package instance is a generated GoMock package.
package instance

import (
	context "context"
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	github "github.com/google/go-github/github"
	oauth2 "golang.org/x/oauth2"
)

// MockInstance is a mock of Instance interface.
type MockInstance struct {
	ctrl     *gomock.Controller
	recorder *MockInstanceMockRecorder
}

// MockInstanceMockRecorder is the mock recorder for MockInstance.
type MockInstanceMockRecorder struct {
	mock *MockInstance
}

// NewMockInstance creates a new mock instance.
func NewMockInstance(ctrl *gomock.Controller) *MockInstance {
	mock := &MockInstance{ctrl: ctrl}
	mock.recorder = &MockInstanceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInstance) EXPECT() *MockInstanceMockRecorder {
	return m.recorder
}

// GetName mocks base method.
func (m *MockInstance) GetName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName.
func (mr *MockInstanceMockRecorder) GetName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockInstance)(nil).GetName))
}

// GetOrganization mocks base method.
func (m *MockInstance) GetOrganization() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrganization")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetOrganization indicates an expected call of GetOrganization.
func (mr *MockInstanceMockRecorder) GetOrganization() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrganization", reflect.TypeOf((*MockInstance)(nil).GetOrganization))
}

// OAuthCallback mocks base method.
func (m *MockInstance) OAuthCallback(ctx context.Context, state, code string) (*oauth2.Token, *github.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OAuthCallback", ctx, state, code)
	ret0, _ := ret[0].(*oauth2.Token)
	ret1, _ := ret[1].(*github.User)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// OAuthCallback indicates an expected call of OAuthCallback.
func (mr *MockInstanceMockRecorder) OAuthCallback(ctx, state, code interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OAuthCallback", reflect.TypeOf((*MockInstance)(nil).OAuthCallback), ctx, state, code)
}

// OAuthIsAuthenticated mocks base method.
func (m *MockInstance) OAuthIsAuthenticated(ctx context.Context, token *oauth2.Token) (*github.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OAuthIsAuthenticated", ctx, token)
	ret0, _ := ret[0].(*github.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OAuthIsAuthenticated indicates an expected call of OAuthIsAuthenticated.
func (mr *MockInstanceMockRecorder) OAuthIsAuthenticated(ctx, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OAuthIsAuthenticated", reflect.TypeOf((*MockInstance)(nil).OAuthIsAuthenticated), ctx, token)
}

// OAuthLoginURL mocks base method.
func (m *MockInstance) OAuthLoginURL() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OAuthLoginURL")
	ret0, _ := ret[0].(string)
	return ret0
}

// OAuthLoginURL indicates an expected call of OAuthLoginURL.
func (mr *MockInstanceMockRecorder) OAuthLoginURL() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OAuthLoginURL", reflect.TypeOf((*MockInstance)(nil).OAuthLoginURL))
}

// TokenFromCookie mocks base method.
func (m *MockInstance) TokenFromCookie(r *http.Request) (*oauth2.Token, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TokenFromCookie", r)
	ret0, _ := ret[0].(*oauth2.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TokenFromCookie indicates an expected call of TokenFromCookie.
func (mr *MockInstanceMockRecorder) TokenFromCookie(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TokenFromCookie", reflect.TypeOf((*MockInstance)(nil).TokenFromCookie), r)
}

// TokenToCookie mocks base method.
func (m *MockInstance) TokenToCookie(token *oauth2.Token) (*http.Cookie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TokenToCookie", token)
	ret0, _ := ret[0].(*http.Cookie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TokenToCookie indicates an expected call of TokenToCookie.
func (mr *MockInstanceMockRecorder) TokenToCookie(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TokenToCookie", reflect.TypeOf((*MockInstance)(nil).TokenToCookie), token)
}
