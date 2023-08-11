// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocksv2

import (
	context "context"

	aws "github.com/aws/aws-sdk-go-v2/aws"

	mock "github.com/stretchr/testify/mock"
)

// CredentialsProvider is an autogenerated mock type for the CredentialsProvider type
type CredentialsProvider struct {
	mock.Mock
}

// Retrieve provides a mock function with given fields: ctx
func (_m *CredentialsProvider) Retrieve(ctx context.Context) (aws.Credentials, error) {
	ret := _m.Called(ctx)

	var r0 aws.Credentials
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (aws.Credentials, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) aws.Credentials); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(aws.Credentials)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewCredentialsProvider creates a new instance of CredentialsProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCredentialsProvider(t interface {
	mock.TestingT
	Cleanup(func())
}) *CredentialsProvider {
	mock := &CredentialsProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
