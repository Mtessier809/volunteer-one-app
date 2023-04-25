// Code generated by mockery v2.26.0. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"

	mock "github.com/stretchr/testify/mock"

	models "github.com/VolunteerOne/volunteer-one-app/backend/models"

	uuid "github.com/google/uuid"
)

// LoginService is an autogenerated mock type for the LoginService type
type LoginService struct {
	mock.Mock
}

// ChangePassword provides a mock function with given fields: _a0, _a1
func (_m *LoginService) ChangePassword(_a0 []byte, _a1 models.Users) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte, models.Users) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CompareHashedAndUserPass provides a mock function with given fields: _a0, _a1
func (_m *LoginService) CompareHashedAndUserPass(_a0 []byte, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteRefreshToken provides a mock function with given fields: _a0
func (_m *LoginService) DeleteRefreshToken(_a0 models.Delegations) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Delegations) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindRefreshToken provides a mock function with given fields: _a0, _a1
func (_m *LoginService) FindRefreshToken(_a0 float64, _a1 models.Delegations) (models.Delegations, error) {
	ret := _m.Called(_a0, _a1)

	var r0 models.Delegations
	var r1 error
	if rf, ok := ret.Get(0).(func(float64, models.Delegations) (models.Delegations, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(float64, models.Delegations) models.Delegations); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(models.Delegations)
	}

	if rf, ok := ret.Get(1).(func(float64, models.Delegations) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindUserFromEmail provides a mock function with given fields: _a0, _a1
func (_m *LoginService) FindUserFromEmail(_a0 string, _a1 models.Users) (models.Users, error) {
	ret := _m.Called(_a0, _a1)

	var r0 models.Users
	var r1 error
	if rf, ok := ret.Get(0).(func(string, models.Users) (models.Users, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(string, models.Users) models.Users); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(models.Users)
	}

	if rf, ok := ret.Get(1).(func(string, models.Users) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GenerateJWT provides a mock function with given fields: _a0, _a1, _a2, _a3, _a4
func (_m *LoginService) GenerateJWT(_a0 uint, _a1 *jwt.NumericDate, _a2 *jwt.NumericDate, _a3 string, _a4 *gin.Context) (string, string, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3, _a4)

	var r0 string
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(uint, *jwt.NumericDate, *jwt.NumericDate, string, *gin.Context) (string, string, error)); ok {
		return rf(_a0, _a1, _a2, _a3, _a4)
	}
	if rf, ok := ret.Get(0).(func(uint, *jwt.NumericDate, *jwt.NumericDate, string, *gin.Context) string); ok {
		r0 = rf(_a0, _a1, _a2, _a3, _a4)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(uint, *jwt.NumericDate, *jwt.NumericDate, string, *gin.Context) string); ok {
		r1 = rf(_a0, _a1, _a2, _a3, _a4)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(uint, *jwt.NumericDate, *jwt.NumericDate, string, *gin.Context) error); ok {
		r2 = rf(_a0, _a1, _a2, _a3, _a4)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// HashPassword provides a mock function with given fields: _a0
func (_m *LoginService) HashPassword(_a0 []byte) ([]byte, error) {
	ret := _m.Called(_a0)

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func([]byte) ([]byte, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func([]byte) []byte); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveRefreshToken provides a mock function with given fields: _a0, _a1, _a2
func (_m *LoginService) SaveRefreshToken(_a0 uint, _a1 string, _a2 models.Delegations) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, string, models.Delegations) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveResetCodeToUser provides a mock function with given fields: _a0, _a1
func (_m *LoginService) SaveResetCodeToUser(_a0 uuid.UUID, _a1 models.Users) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(uuid.UUID, models.Users) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewLoginService interface {
	mock.TestingT
	Cleanup(func())
}

// NewLoginService creates a new instance of LoginService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewLoginService(t mockConstructorTestingTNewLoginService) *LoginService {
	mock := &LoginService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
