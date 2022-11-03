// Copyright 2016 Stanislav Liberman
// Copyright 2022 Steven Stern
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// If FragmentHandler changes, recreate the mock code with the below command.
// mockery --name=FragmentHandler --inpackage --structname=MockFragmentHandler --print

// Code generated by mockery v2.14.0. DO NOT EDIT.

package term

import (
	atomic "github.com/lirm/aeron-go/aeron/atomic"
	logbuffer "github.com/lirm/aeron-go/aeron/logbuffer"

	mock "github.com/stretchr/testify/mock"
)

// MockFragmentHandler is an autogenerated mock type for the FragmentHandler type
type MockFragmentHandler struct {
	mock.Mock
}

// Execute provides a mock function with given fields: buffer, offset, length, header
func (_m *MockFragmentHandler) Execute(buffer *atomic.Buffer, offset int32, length int32, header *logbuffer.Header) error {
	ret := _m.Called(buffer, offset, length, header)

	var r0 error
	if rf, ok := ret.Get(0).(func(*atomic.Buffer, int32, int32, *logbuffer.Header) error); ok {
		r0 = rf(buffer, offset, length, header)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewMockFragmentHandler interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockFragmentHandler creates a new instance of MockFragmentHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockFragmentHandler(t mockConstructorTestingTNewMockFragmentHandler) *MockFragmentHandler {
	mock := &MockFragmentHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}