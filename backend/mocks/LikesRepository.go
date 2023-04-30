// Code generated by mockery v2.23.1. DO NOT EDIT.

package mocks

import (
	models "github.com/VolunteerOne/volunteer-one-app/backend/models"
	mock "github.com/stretchr/testify/mock"
)

// LikesRepository is an autogenerated mock type for the LikesRepository type
type LikesRepository struct {
	mock.Mock
}

// AllLikes provides a mock function with given fields:
func (_m *LikesRepository) AllLikes() ([]models.Likes, error) {
	ret := _m.Called()

	var r0 []models.Likes
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]models.Likes, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []models.Likes); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Likes)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateLike provides a mock function with given fields: Like
func (_m *LikesRepository) CreateLike(Like models.Likes) (models.Likes, error) {
	ret := _m.Called(Like)

	var r0 models.Likes
	var r1 error
	if rf, ok := ret.Get(0).(func(models.Likes) (models.Likes, error)); ok {
		return rf(Like)
	}
	if rf, ok := ret.Get(0).(func(models.Likes) models.Likes); ok {
		r0 = rf(Like)
	} else {
		r0 = ret.Get(0).(models.Likes)
	}

	if rf, ok := ret.Get(1).(func(models.Likes) error); ok {
		r1 = rf(Like)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteLike provides a mock function with given fields: Like
func (_m *LikesRepository) DeleteLike(Like models.Likes) error {
	ret := _m.Called(Like)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Likes) error); ok {
		r0 = rf(Like)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindLike provides a mock function with given fields: id
func (_m *LikesRepository) FindLike(id string) (models.Likes, error) {
	ret := _m.Called(id)

	var r0 models.Likes
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (models.Likes, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) models.Likes); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(models.Likes)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLikes provides a mock function with given fields: id
func (_m *LikesRepository) GetLikes(id string) (int64, error) {
	ret := _m.Called(id)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (int64, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) int64); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewLikesRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewLikesRepository creates a new instance of LikesRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewLikesRepository(t mockConstructorTestingTNewLikesRepository) *LikesRepository {
	mock := &LikesRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}