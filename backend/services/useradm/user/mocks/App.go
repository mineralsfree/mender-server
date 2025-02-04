// Copyright 2023 Northern.tech AS
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

// Code generated by mockery v2.45.1. DO NOT EDIT.

package mocks

import (
	context "context"

	jwt "github.com/mendersoftware/mender-server/services/useradm/jwt"
	mock "github.com/stretchr/testify/mock"

	model "github.com/mendersoftware/mender-server/services/useradm/model"

	useradm "github.com/mendersoftware/mender-server/services/useradm/user"
)

// App is an autogenerated mock type for the App type
type App struct {
	mock.Mock
}

// CreateTenant provides a mock function with given fields: ctx, tenant
func (_m *App) CreateTenant(ctx context.Context, tenant model.NewTenant) error {
	ret := _m.Called(ctx, tenant)

	if len(ret) == 0 {
		panic("no return value specified for CreateTenant")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.NewTenant) error); ok {
		r0 = rf(ctx, tenant)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateUser provides a mock function with given fields: ctx, u
func (_m *App) CreateUser(ctx context.Context, u *model.User) error {
	ret := _m.Called(ctx, u)

	if len(ret) == 0 {
		panic("no return value specified for CreateUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.User) error); ok {
		r0 = rf(ctx, u)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateUserInternal provides a mock function with given fields: ctx, u
func (_m *App) CreateUserInternal(ctx context.Context, u *model.UserInternal) error {
	ret := _m.Called(ctx, u)

	if len(ret) == 0 {
		panic("no return value specified for CreateUserInternal")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.UserInternal) error); ok {
		r0 = rf(ctx, u)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteToken provides a mock function with given fields: ctx, id
func (_m *App) DeleteToken(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteToken")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteTokens provides a mock function with given fields: ctx, tenantId, userId
func (_m *App) DeleteTokens(ctx context.Context, tenantId string, userId string) error {
	ret := _m.Called(ctx, tenantId, userId)

	if len(ret) == 0 {
		panic("no return value specified for DeleteTokens")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, tenantId, userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteUser provides a mock function with given fields: ctx, id
func (_m *App) DeleteUser(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetPersonalAccessTokens provides a mock function with given fields: ctx, userID
func (_m *App) GetPersonalAccessTokens(ctx context.Context, userID string) ([]model.PersonalAccessToken, error) {
	ret := _m.Called(ctx, userID)

	if len(ret) == 0 {
		panic("no return value specified for GetPersonalAccessTokens")
	}

	var r0 []model.PersonalAccessToken
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]model.PersonalAccessToken, error)); ok {
		return rf(ctx, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []model.PersonalAccessToken); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.PersonalAccessToken)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPlanBinding provides a mock function with given fields: ctx
func (_m *App) GetPlanBinding(ctx context.Context) (*model.PlanBindingDetails, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetPlanBinding")
	}

	var r0 *model.PlanBindingDetails
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*model.PlanBindingDetails, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *model.PlanBindingDetails); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.PlanBindingDetails)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPlans provides a mock function with given fields: ctx, skip, limit
func (_m *App) GetPlans(ctx context.Context, skip int, limit int) []model.Plan {
	ret := _m.Called(ctx, skip, limit)

	if len(ret) == 0 {
		panic("no return value specified for GetPlans")
	}

	var r0 []model.Plan
	if rf, ok := ret.Get(0).(func(context.Context, int, int) []model.Plan); ok {
		r0 = rf(ctx, skip, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Plan)
		}
	}

	return r0
}

// GetUser provides a mock function with given fields: ctx, id
func (_m *App) GetUser(ctx context.Context, id string) (*model.User, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetUser")
	}

	var r0 *model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*model.User, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.User); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUsers provides a mock function with given fields: ctx, fltr
func (_m *App) GetUsers(ctx context.Context, fltr model.UserFilter) ([]model.User, error) {
	ret := _m.Called(ctx, fltr)

	if len(ret) == 0 {
		panic("no return value specified for GetUsers")
	}

	var r0 []model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.UserFilter) ([]model.User, error)); ok {
		return rf(ctx, fltr)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.UserFilter) []model.User); ok {
		r0 = rf(ctx, fltr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.UserFilter) error); ok {
		r1 = rf(ctx, fltr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HealthCheck provides a mock function with given fields: ctx
func (_m *App) HealthCheck(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for HealthCheck")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IssuePersonalAccessToken provides a mock function with given fields: ctx, tr
func (_m *App) IssuePersonalAccessToken(ctx context.Context, tr *model.TokenRequest) (string, error) {
	ret := _m.Called(ctx, tr)

	if len(ret) == 0 {
		panic("no return value specified for IssuePersonalAccessToken")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.TokenRequest) (string, error)); ok {
		return rf(ctx, tr)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.TokenRequest) string); ok {
		r0 = rf(ctx, tr)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.TokenRequest) error); ok {
		r1 = rf(ctx, tr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: ctx, email, pass, options
func (_m *App) Login(ctx context.Context, email model.Email, pass string, options *useradm.LoginOptions) (*jwt.Token, error) {
	ret := _m.Called(ctx, email, pass, options)

	if len(ret) == 0 {
		panic("no return value specified for Login")
	}

	var r0 *jwt.Token
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.Email, string, *useradm.LoginOptions) (*jwt.Token, error)); ok {
		return rf(ctx, email, pass, options)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.Email, string, *useradm.LoginOptions) *jwt.Token); ok {
		r0 = rf(ctx, email, pass, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*jwt.Token)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.Email, string, *useradm.LoginOptions) error); ok {
		r1 = rf(ctx, email, pass, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Logout provides a mock function with given fields: ctx, token
func (_m *App) Logout(ctx context.Context, token *jwt.Token) error {
	ret := _m.Called(ctx, token)

	if len(ret) == 0 {
		panic("no return value specified for Logout")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *jwt.Token) error); ok {
		r0 = rf(ctx, token)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetPassword provides a mock function with given fields: ctx, u
func (_m *App) SetPassword(ctx context.Context, u model.UserUpdate) error {
	ret := _m.Called(ctx, u)

	if len(ret) == 0 {
		panic("no return value specified for SetPassword")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.UserUpdate) error); ok {
		r0 = rf(ctx, u)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SignToken provides a mock function with given fields: ctx, t
func (_m *App) SignToken(ctx context.Context, t *jwt.Token) (string, error) {
	ret := _m.Called(ctx, t)

	if len(ret) == 0 {
		panic("no return value specified for SignToken")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *jwt.Token) (string, error)); ok {
		return rf(ctx, t)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *jwt.Token) string); ok {
		r0 = rf(ctx, t)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *jwt.Token) error); ok {
		r1 = rf(ctx, t)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUser provides a mock function with given fields: ctx, id, u
func (_m *App) UpdateUser(ctx context.Context, id string, u *model.UserUpdate) error {
	ret := _m.Called(ctx, id, u)

	if len(ret) == 0 {
		panic("no return value specified for UpdateUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *model.UserUpdate) error); ok {
		r0 = rf(ctx, id, u)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Verify provides a mock function with given fields: ctx, token
func (_m *App) Verify(ctx context.Context, token *jwt.Token) error {
	ret := _m.Called(ctx, token)

	if len(ret) == 0 {
		panic("no return value specified for Verify")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *jwt.Token) error); ok {
		r0 = rf(ctx, token)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewApp creates a new instance of App. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewApp(t interface {
	mock.TestingT
	Cleanup(func())
}) *App {
	mock := &App{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
