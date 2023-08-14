// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	context "context"

	port "github.com/JackDaniells/port-service/proto"
	mock "github.com/stretchr/testify/mock"
)

// PortServiceServer is an autogenerated mock type for the PortServiceServer type
type PortServiceServer struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0
func (_m *PortServiceServer) Create(_a0 port.PortService_CreateServer) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(port.PortService_CreateServer) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindByID provides a mock function with given fields: _a0, _a1
func (_m *PortServiceServer) FindByID(_a0 context.Context, _a1 *port.FindByIDRequest) (*port.Port, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *port.Port
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *port.FindByIDRequest) (*port.Port, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *port.FindByIDRequest) *port.Port); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*port.Port)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *port.FindByIDRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mustEmbedUnimplementedPortServiceServer provides a mock function with given fields:
func (_m *PortServiceServer) mustEmbedUnimplementedPortServiceServer() {
	_m.Called()
}

// NewPortServiceServer creates a new instance of PortServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPortServiceServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *PortServiceServer {
	mock := &PortServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
