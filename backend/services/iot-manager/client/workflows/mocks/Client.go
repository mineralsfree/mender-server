// Copyright 2022 Northern.tech AS
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

	mock "github.com/stretchr/testify/mock"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// CheckHealth provides a mock function with given fields: ctx
func (_m *Client) CheckHealth(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for CheckHealth")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ProvisionExternalDevice provides a mock function with given fields: ctx, devID, config
func (_m *Client) ProvisionExternalDevice(ctx context.Context, devID string, config map[string]string) error {
	ret := _m.Called(ctx, devID, config)

	if len(ret) == 0 {
		panic("no return value specified for ProvisionExternalDevice")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, map[string]string) error); ok {
		r0 = rf(ctx, devID, config)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewClient creates a new instance of Client. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *Client {
	mock := &Client{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
