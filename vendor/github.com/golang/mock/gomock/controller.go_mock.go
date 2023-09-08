// Code generated by MockGen. DO NOT EDIT.
// Source: vendor/github.com/golang/mock/gomock/controller.go

// Package mock_gomock is a generated GoMock package.
package mock_gomock

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockTestReporter is a mock of TestReporter interface.
type MockTestReporter struct {
	ctrl     *gomock.Controller
	recorder *MockTestReporterMockRecorder
}

// MockTestReporterMockRecorder is the mock recorder for MockTestReporter.
type MockTestReporterMockRecorder struct {
	mock *MockTestReporter
}

// NewMockTestReporter creates a new mock instance.
func NewMockTestReporter(ctrl *gomock.Controller) *MockTestReporter {
	mock := &MockTestReporter{ctrl: ctrl}
	mock.recorder = &MockTestReporterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTestReporter) EXPECT() *MockTestReporterMockRecorder {
	return m.recorder
}

// Errorf mocks base method.
func (m *MockTestReporter) Errorf(format string, args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Errorf", varargs...)
}

// Errorf indicates an expected call of Errorf.
func (mr *MockTestReporterMockRecorder) Errorf(format interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Errorf", reflect.TypeOf((*MockTestReporter)(nil).Errorf), varargs...)
}

// Fatalf mocks base method.
func (m *MockTestReporter) Fatalf(format string, args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Fatalf", varargs...)
}

// Fatalf indicates an expected call of Fatalf.
func (mr *MockTestReporterMockRecorder) Fatalf(format interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fatalf", reflect.TypeOf((*MockTestReporter)(nil).Fatalf), varargs...)
}

// MockTestHelper is a mock of TestHelper interface.
type MockTestHelper struct {
	ctrl     *gomock.Controller
	recorder *MockTestHelperMockRecorder
}

// MockTestHelperMockRecorder is the mock recorder for MockTestHelper.
type MockTestHelperMockRecorder struct {
	mock *MockTestHelper
}

// NewMockTestHelper creates a new mock instance.
func NewMockTestHelper(ctrl *gomock.Controller) *MockTestHelper {
	mock := &MockTestHelper{ctrl: ctrl}
	mock.recorder = &MockTestHelperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTestHelper) EXPECT() *MockTestHelperMockRecorder {
	return m.recorder
}

// Errorf mocks base method.
func (m *MockTestHelper) Errorf(format string, args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Errorf", varargs...)
}

// Errorf indicates an expected call of Errorf.
func (mr *MockTestHelperMockRecorder) Errorf(format interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Errorf", reflect.TypeOf((*MockTestHelper)(nil).Errorf), varargs...)
}

// Fatalf mocks base method.
func (m *MockTestHelper) Fatalf(format string, args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Fatalf", varargs...)
}

// Fatalf indicates an expected call of Fatalf.
func (mr *MockTestHelperMockRecorder) Fatalf(format interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fatalf", reflect.TypeOf((*MockTestHelper)(nil).Fatalf), varargs...)
}

// Helper mocks base method.
func (m *MockTestHelper) Helper() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Helper")
}

// Helper indicates an expected call of Helper.
func (mr *MockTestHelperMockRecorder) Helper() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Helper", reflect.TypeOf((*MockTestHelper)(nil).Helper))
}
