// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	port "github.com/JackDaniells/port-service/proto"
)

// PortServiceClient is an autogenerated mock type for the PortServiceClient type
type PortServiceClient struct {
	mock.Mock
}

// FindByID provides a mock function with given fields: ctx, id
func (_m *PortServiceClient) FindByID(ctx context.Context, id string) (*port.Port, error) {
	ret := _m.Called(ctx, id)

	var r0 *port.Port
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*port.Port, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *port.Port); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*port.Port)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StreamCreate provides a mock function with given fields: ctx, ports
func (_m *PortServiceClient) StreamCreate(ctx context.Context, ports []*port.Port) (*port.CreateResponse, error) {
	ret := _m.Called(ctx, ports)

	var r0 *port.CreateResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []*port.Port) (*port.CreateResponse, error)); ok {
		return rf(ctx, ports)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []*port.Port) *port.CreateResponse); ok {
		r0 = rf(ctx, ports)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*port.CreateResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []*port.Port) error); ok {
		r1 = rf(ctx, ports)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewPortServiceClient creates a new instance of PortServiceClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPortServiceClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *PortServiceClient {
	mock := &PortServiceClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
