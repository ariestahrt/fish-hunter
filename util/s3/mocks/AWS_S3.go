// Code generated by mockery v2.14.1. DO NOT EDIT.

package mocks

import (
	os "os"

	mock "github.com/stretchr/testify/mock"
)

// AWS_S3 is an autogenerated mock type for the AWS_S3 type
type AWS_S3 struct {
	mock.Mock
}

// DownloadFile provides a mock function with given fields: file, key
func (_m *AWS_S3) DownloadFile(file *os.File, key string) error {
	ret := _m.Called(file, key)

	var r0 error
	if rf, ok := ret.Get(0).(func(*os.File, string) error); ok {
		r0 = rf(file, key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UploadFile provides a mock function with given fields: file, key
func (_m *AWS_S3) UploadFile(file *os.File, key string) error {
	ret := _m.Called(file, key)

	var r0 error
	if rf, ok := ret.Get(0).(func(*os.File, string) error); ok {
		r0 = rf(file, key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewAWS_S3 interface {
	mock.TestingT
	Cleanup(func())
}

// NewAWS_S3 creates a new instance of AWS_S3. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAWS_S3(t mockConstructorTestingTNewAWS_S3) *AWS_S3 {
	mock := &AWS_S3{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
