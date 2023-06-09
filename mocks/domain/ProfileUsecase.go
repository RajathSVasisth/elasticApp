// Code generated by mockery v2.30.1. DO NOT EDIT.

package domain

import (
	context "context"

	elastic_go_app_cleancodedomain "github.com/RajathSVasisth/elastic-go-app-cleancode/domain"
	mock "github.com/stretchr/testify/mock"
)

// ProfileUsecase is an autogenerated mock type for the ProfileUsecase type
type ProfileUsecase struct {
	mock.Mock
}

// GetProfileByID provides a mock function with given fields: c, userID
func (_m *ProfileUsecase) GetProfileByID(c context.Context, userID string) (*elastic_go_app_cleancodedomain.Profile, error) {
	ret := _m.Called(c, userID)

	var r0 *elastic_go_app_cleancodedomain.Profile
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*elastic_go_app_cleancodedomain.Profile, error)); ok {
		return rf(c, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *elastic_go_app_cleancodedomain.Profile); ok {
		r0 = rf(c, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*elastic_go_app_cleancodedomain.Profile)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(c, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewProfileUsecase creates a new instance of ProfileUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewProfileUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *ProfileUsecase {
	mock := &ProfileUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
