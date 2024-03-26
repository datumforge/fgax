// Code generated by mockery. DO NOT EDIT.

package client

import (
	context "context"

	client "github.com/openfga/go-sdk/client"

	mock "github.com/stretchr/testify/mock"
)

// MockSdkClientDeleteStoreRequestInterface is an autogenerated mock type for the SdkClientDeleteStoreRequestInterface type
type MockSdkClientDeleteStoreRequestInterface struct {
	mock.Mock
}

type MockSdkClientDeleteStoreRequestInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockSdkClientDeleteStoreRequestInterface) EXPECT() *MockSdkClientDeleteStoreRequestInterface_Expecter {
	return &MockSdkClientDeleteStoreRequestInterface_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields:
func (_m *MockSdkClientDeleteStoreRequestInterface) Execute() (*client.ClientDeleteStoreResponse, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Execute")
	}

	var r0 *client.ClientDeleteStoreResponse
	var r1 error
	if rf, ok := ret.Get(0).(func() (*client.ClientDeleteStoreResponse, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *client.ClientDeleteStoreResponse); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.ClientDeleteStoreResponse)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSdkClientDeleteStoreRequestInterface_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type MockSdkClientDeleteStoreRequestInterface_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
func (_e *MockSdkClientDeleteStoreRequestInterface_Expecter) Execute() *MockSdkClientDeleteStoreRequestInterface_Execute_Call {
	return &MockSdkClientDeleteStoreRequestInterface_Execute_Call{Call: _e.mock.On("Execute")}
}

func (_c *MockSdkClientDeleteStoreRequestInterface_Execute_Call) Run(run func()) *MockSdkClientDeleteStoreRequestInterface_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSdkClientDeleteStoreRequestInterface_Execute_Call) Return(_a0 *client.ClientDeleteStoreResponse, _a1 error) *MockSdkClientDeleteStoreRequestInterface_Execute_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSdkClientDeleteStoreRequestInterface_Execute_Call) RunAndReturn(run func() (*client.ClientDeleteStoreResponse, error)) *MockSdkClientDeleteStoreRequestInterface_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// GetContext provides a mock function with given fields:
func (_m *MockSdkClientDeleteStoreRequestInterface) GetContext() context.Context {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetContext")
	}

	var r0 context.Context
	if rf, ok := ret.Get(0).(func() context.Context); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	return r0
}

// MockSdkClientDeleteStoreRequestInterface_GetContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetContext'
type MockSdkClientDeleteStoreRequestInterface_GetContext_Call struct {
	*mock.Call
}

// GetContext is a helper method to define mock.On call
func (_e *MockSdkClientDeleteStoreRequestInterface_Expecter) GetContext() *MockSdkClientDeleteStoreRequestInterface_GetContext_Call {
	return &MockSdkClientDeleteStoreRequestInterface_GetContext_Call{Call: _e.mock.On("GetContext")}
}

func (_c *MockSdkClientDeleteStoreRequestInterface_GetContext_Call) Run(run func()) *MockSdkClientDeleteStoreRequestInterface_GetContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSdkClientDeleteStoreRequestInterface_GetContext_Call) Return(_a0 context.Context) *MockSdkClientDeleteStoreRequestInterface_GetContext_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSdkClientDeleteStoreRequestInterface_GetContext_Call) RunAndReturn(run func() context.Context) *MockSdkClientDeleteStoreRequestInterface_GetContext_Call {
	_c.Call.Return(run)
	return _c
}

// GetOptions provides a mock function with given fields:
func (_m *MockSdkClientDeleteStoreRequestInterface) GetOptions() *client.ClientDeleteStoreOptions {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetOptions")
	}

	var r0 *client.ClientDeleteStoreOptions
	if rf, ok := ret.Get(0).(func() *client.ClientDeleteStoreOptions); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.ClientDeleteStoreOptions)
		}
	}

	return r0
}

// MockSdkClientDeleteStoreRequestInterface_GetOptions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetOptions'
type MockSdkClientDeleteStoreRequestInterface_GetOptions_Call struct {
	*mock.Call
}

// GetOptions is a helper method to define mock.On call
func (_e *MockSdkClientDeleteStoreRequestInterface_Expecter) GetOptions() *MockSdkClientDeleteStoreRequestInterface_GetOptions_Call {
	return &MockSdkClientDeleteStoreRequestInterface_GetOptions_Call{Call: _e.mock.On("GetOptions")}
}

func (_c *MockSdkClientDeleteStoreRequestInterface_GetOptions_Call) Run(run func()) *MockSdkClientDeleteStoreRequestInterface_GetOptions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSdkClientDeleteStoreRequestInterface_GetOptions_Call) Return(_a0 *client.ClientDeleteStoreOptions) *MockSdkClientDeleteStoreRequestInterface_GetOptions_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSdkClientDeleteStoreRequestInterface_GetOptions_Call) RunAndReturn(run func() *client.ClientDeleteStoreOptions) *MockSdkClientDeleteStoreRequestInterface_GetOptions_Call {
	_c.Call.Return(run)
	return _c
}

// Options provides a mock function with given fields: options
func (_m *MockSdkClientDeleteStoreRequestInterface) Options(options client.ClientDeleteStoreOptions) client.SdkClientDeleteStoreRequestInterface {
	ret := _m.Called(options)

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 client.SdkClientDeleteStoreRequestInterface
	if rf, ok := ret.Get(0).(func(client.ClientDeleteStoreOptions) client.SdkClientDeleteStoreRequestInterface); ok {
		r0 = rf(options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(client.SdkClientDeleteStoreRequestInterface)
		}
	}

	return r0
}

// MockSdkClientDeleteStoreRequestInterface_Options_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Options'
type MockSdkClientDeleteStoreRequestInterface_Options_Call struct {
	*mock.Call
}

// Options is a helper method to define mock.On call
//   - options client.ClientDeleteStoreOptions
func (_e *MockSdkClientDeleteStoreRequestInterface_Expecter) Options(options interface{}) *MockSdkClientDeleteStoreRequestInterface_Options_Call {
	return &MockSdkClientDeleteStoreRequestInterface_Options_Call{Call: _e.mock.On("Options", options)}
}

func (_c *MockSdkClientDeleteStoreRequestInterface_Options_Call) Run(run func(options client.ClientDeleteStoreOptions)) *MockSdkClientDeleteStoreRequestInterface_Options_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(client.ClientDeleteStoreOptions))
	})
	return _c
}

func (_c *MockSdkClientDeleteStoreRequestInterface_Options_Call) Return(_a0 client.SdkClientDeleteStoreRequestInterface) *MockSdkClientDeleteStoreRequestInterface_Options_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSdkClientDeleteStoreRequestInterface_Options_Call) RunAndReturn(run func(client.ClientDeleteStoreOptions) client.SdkClientDeleteStoreRequestInterface) *MockSdkClientDeleteStoreRequestInterface_Options_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockSdkClientDeleteStoreRequestInterface creates a new instance of MockSdkClientDeleteStoreRequestInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockSdkClientDeleteStoreRequestInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockSdkClientDeleteStoreRequestInterface {
	mock := &MockSdkClientDeleteStoreRequestInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
