// Code generated by mockery. DO NOT EDIT.

package client

import (
	context "context"

	client "github.com/openfga/go-sdk/client"

	mock "github.com/stretchr/testify/mock"

	openfga "github.com/openfga/go-sdk"
)

// MockSdkClientListUsersRequestInterface is an autogenerated mock type for the SdkClientListUsersRequestInterface type
type MockSdkClientListUsersRequestInterface struct {
	mock.Mock
}

type MockSdkClientListUsersRequestInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockSdkClientListUsersRequestInterface) EXPECT() *MockSdkClientListUsersRequestInterface_Expecter {
	return &MockSdkClientListUsersRequestInterface_Expecter{mock: &_m.Mock}
}

// Body provides a mock function with given fields: body
func (_m *MockSdkClientListUsersRequestInterface) Body(body client.ClientListUsersRequest) client.SdkClientListUsersRequestInterface {
	ret := _m.Called(body)

	if len(ret) == 0 {
		panic("no return value specified for Body")
	}

	var r0 client.SdkClientListUsersRequestInterface
	if rf, ok := ret.Get(0).(func(client.ClientListUsersRequest) client.SdkClientListUsersRequestInterface); ok {
		r0 = rf(body)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(client.SdkClientListUsersRequestInterface)
		}
	}

	return r0
}

// MockSdkClientListUsersRequestInterface_Body_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Body'
type MockSdkClientListUsersRequestInterface_Body_Call struct {
	*mock.Call
}

// Body is a helper method to define mock.On call
//   - body client.ClientListUsersRequest
func (_e *MockSdkClientListUsersRequestInterface_Expecter) Body(body interface{}) *MockSdkClientListUsersRequestInterface_Body_Call {
	return &MockSdkClientListUsersRequestInterface_Body_Call{Call: _e.mock.On("Body", body)}
}

func (_c *MockSdkClientListUsersRequestInterface_Body_Call) Run(run func(body client.ClientListUsersRequest)) *MockSdkClientListUsersRequestInterface_Body_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(client.ClientListUsersRequest))
	})
	return _c
}

func (_c *MockSdkClientListUsersRequestInterface_Body_Call) Return(_a0 client.SdkClientListUsersRequestInterface) *MockSdkClientListUsersRequestInterface_Body_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSdkClientListUsersRequestInterface_Body_Call) RunAndReturn(run func(client.ClientListUsersRequest) client.SdkClientListUsersRequestInterface) *MockSdkClientListUsersRequestInterface_Body_Call {
	_c.Call.Return(run)
	return _c
}

// Execute provides a mock function with given fields:
func (_m *MockSdkClientListUsersRequestInterface) Execute() (*openfga.ListUsersResponse, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Execute")
	}

	var r0 *openfga.ListUsersResponse
	var r1 error
	if rf, ok := ret.Get(0).(func() (*openfga.ListUsersResponse, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *openfga.ListUsersResponse); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*openfga.ListUsersResponse)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSdkClientListUsersRequestInterface_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type MockSdkClientListUsersRequestInterface_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
func (_e *MockSdkClientListUsersRequestInterface_Expecter) Execute() *MockSdkClientListUsersRequestInterface_Execute_Call {
	return &MockSdkClientListUsersRequestInterface_Execute_Call{Call: _e.mock.On("Execute")}
}

func (_c *MockSdkClientListUsersRequestInterface_Execute_Call) Run(run func()) *MockSdkClientListUsersRequestInterface_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSdkClientListUsersRequestInterface_Execute_Call) Return(_a0 *openfga.ListUsersResponse, _a1 error) *MockSdkClientListUsersRequestInterface_Execute_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSdkClientListUsersRequestInterface_Execute_Call) RunAndReturn(run func() (*openfga.ListUsersResponse, error)) *MockSdkClientListUsersRequestInterface_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// GetAuthorizationModelIdOverride provides a mock function with given fields:
func (_m *MockSdkClientListUsersRequestInterface) GetAuthorizationModelIdOverride() *string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAuthorizationModelIdOverride")
	}

	var r0 *string
	if rf, ok := ret.Get(0).(func() *string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*string)
		}
	}

	return r0
}

// MockSdkClientListUsersRequestInterface_GetAuthorizationModelIdOverride_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAuthorizationModelIdOverride'
type MockSdkClientListUsersRequestInterface_GetAuthorizationModelIdOverride_Call struct {
	*mock.Call
}

// GetAuthorizationModelIdOverride is a helper method to define mock.On call
func (_e *MockSdkClientListUsersRequestInterface_Expecter) GetAuthorizationModelIdOverride() *MockSdkClientListUsersRequestInterface_GetAuthorizationModelIdOverride_Call {
	return &MockSdkClientListUsersRequestInterface_GetAuthorizationModelIdOverride_Call{Call: _e.mock.On("GetAuthorizationModelIdOverride")}
}

func (_c *MockSdkClientListUsersRequestInterface_GetAuthorizationModelIdOverride_Call) Run(run func()) *MockSdkClientListUsersRequestInterface_GetAuthorizationModelIdOverride_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSdkClientListUsersRequestInterface_GetAuthorizationModelIdOverride_Call) Return(_a0 *string) *MockSdkClientListUsersRequestInterface_GetAuthorizationModelIdOverride_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSdkClientListUsersRequestInterface_GetAuthorizationModelIdOverride_Call) RunAndReturn(run func() *string) *MockSdkClientListUsersRequestInterface_GetAuthorizationModelIdOverride_Call {
	_c.Call.Return(run)
	return _c
}

// GetBody provides a mock function with given fields:
func (_m *MockSdkClientListUsersRequestInterface) GetBody() *client.ClientListUsersRequest {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetBody")
	}

	var r0 *client.ClientListUsersRequest
	if rf, ok := ret.Get(0).(func() *client.ClientListUsersRequest); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.ClientListUsersRequest)
		}
	}

	return r0
}

// MockSdkClientListUsersRequestInterface_GetBody_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBody'
type MockSdkClientListUsersRequestInterface_GetBody_Call struct {
	*mock.Call
}

// GetBody is a helper method to define mock.On call
func (_e *MockSdkClientListUsersRequestInterface_Expecter) GetBody() *MockSdkClientListUsersRequestInterface_GetBody_Call {
	return &MockSdkClientListUsersRequestInterface_GetBody_Call{Call: _e.mock.On("GetBody")}
}

func (_c *MockSdkClientListUsersRequestInterface_GetBody_Call) Run(run func()) *MockSdkClientListUsersRequestInterface_GetBody_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSdkClientListUsersRequestInterface_GetBody_Call) Return(_a0 *client.ClientListUsersRequest) *MockSdkClientListUsersRequestInterface_GetBody_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSdkClientListUsersRequestInterface_GetBody_Call) RunAndReturn(run func() *client.ClientListUsersRequest) *MockSdkClientListUsersRequestInterface_GetBody_Call {
	_c.Call.Return(run)
	return _c
}

// GetContext provides a mock function with given fields:
func (_m *MockSdkClientListUsersRequestInterface) GetContext() context.Context {
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

// MockSdkClientListUsersRequestInterface_GetContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetContext'
type MockSdkClientListUsersRequestInterface_GetContext_Call struct {
	*mock.Call
}

// GetContext is a helper method to define mock.On call
func (_e *MockSdkClientListUsersRequestInterface_Expecter) GetContext() *MockSdkClientListUsersRequestInterface_GetContext_Call {
	return &MockSdkClientListUsersRequestInterface_GetContext_Call{Call: _e.mock.On("GetContext")}
}

func (_c *MockSdkClientListUsersRequestInterface_GetContext_Call) Run(run func()) *MockSdkClientListUsersRequestInterface_GetContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSdkClientListUsersRequestInterface_GetContext_Call) Return(_a0 context.Context) *MockSdkClientListUsersRequestInterface_GetContext_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSdkClientListUsersRequestInterface_GetContext_Call) RunAndReturn(run func() context.Context) *MockSdkClientListUsersRequestInterface_GetContext_Call {
	_c.Call.Return(run)
	return _c
}

// GetOptions provides a mock function with given fields:
func (_m *MockSdkClientListUsersRequestInterface) GetOptions() *client.ClientListUsersOptions {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetOptions")
	}

	var r0 *client.ClientListUsersOptions
	if rf, ok := ret.Get(0).(func() *client.ClientListUsersOptions); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.ClientListUsersOptions)
		}
	}

	return r0
}

// MockSdkClientListUsersRequestInterface_GetOptions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetOptions'
type MockSdkClientListUsersRequestInterface_GetOptions_Call struct {
	*mock.Call
}

// GetOptions is a helper method to define mock.On call
func (_e *MockSdkClientListUsersRequestInterface_Expecter) GetOptions() *MockSdkClientListUsersRequestInterface_GetOptions_Call {
	return &MockSdkClientListUsersRequestInterface_GetOptions_Call{Call: _e.mock.On("GetOptions")}
}

func (_c *MockSdkClientListUsersRequestInterface_GetOptions_Call) Run(run func()) *MockSdkClientListUsersRequestInterface_GetOptions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSdkClientListUsersRequestInterface_GetOptions_Call) Return(_a0 *client.ClientListUsersOptions) *MockSdkClientListUsersRequestInterface_GetOptions_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSdkClientListUsersRequestInterface_GetOptions_Call) RunAndReturn(run func() *client.ClientListUsersOptions) *MockSdkClientListUsersRequestInterface_GetOptions_Call {
	_c.Call.Return(run)
	return _c
}

// GetStoreIdOverride provides a mock function with given fields:
func (_m *MockSdkClientListUsersRequestInterface) GetStoreIdOverride() *string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetStoreIdOverride")
	}

	var r0 *string
	if rf, ok := ret.Get(0).(func() *string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*string)
		}
	}

	return r0
}

// MockSdkClientListUsersRequestInterface_GetStoreIdOverride_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetStoreIdOverride'
type MockSdkClientListUsersRequestInterface_GetStoreIdOverride_Call struct {
	*mock.Call
}

// GetStoreIdOverride is a helper method to define mock.On call
func (_e *MockSdkClientListUsersRequestInterface_Expecter) GetStoreIdOverride() *MockSdkClientListUsersRequestInterface_GetStoreIdOverride_Call {
	return &MockSdkClientListUsersRequestInterface_GetStoreIdOverride_Call{Call: _e.mock.On("GetStoreIdOverride")}
}

func (_c *MockSdkClientListUsersRequestInterface_GetStoreIdOverride_Call) Run(run func()) *MockSdkClientListUsersRequestInterface_GetStoreIdOverride_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSdkClientListUsersRequestInterface_GetStoreIdOverride_Call) Return(_a0 *string) *MockSdkClientListUsersRequestInterface_GetStoreIdOverride_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSdkClientListUsersRequestInterface_GetStoreIdOverride_Call) RunAndReturn(run func() *string) *MockSdkClientListUsersRequestInterface_GetStoreIdOverride_Call {
	_c.Call.Return(run)
	return _c
}

// Options provides a mock function with given fields: options
func (_m *MockSdkClientListUsersRequestInterface) Options(options client.ClientListUsersOptions) client.SdkClientListUsersRequestInterface {
	ret := _m.Called(options)

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 client.SdkClientListUsersRequestInterface
	if rf, ok := ret.Get(0).(func(client.ClientListUsersOptions) client.SdkClientListUsersRequestInterface); ok {
		r0 = rf(options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(client.SdkClientListUsersRequestInterface)
		}
	}

	return r0
}

// MockSdkClientListUsersRequestInterface_Options_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Options'
type MockSdkClientListUsersRequestInterface_Options_Call struct {
	*mock.Call
}

// Options is a helper method to define mock.On call
//   - options client.ClientListUsersOptions
func (_e *MockSdkClientListUsersRequestInterface_Expecter) Options(options interface{}) *MockSdkClientListUsersRequestInterface_Options_Call {
	return &MockSdkClientListUsersRequestInterface_Options_Call{Call: _e.mock.On("Options", options)}
}

func (_c *MockSdkClientListUsersRequestInterface_Options_Call) Run(run func(options client.ClientListUsersOptions)) *MockSdkClientListUsersRequestInterface_Options_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(client.ClientListUsersOptions))
	})
	return _c
}

func (_c *MockSdkClientListUsersRequestInterface_Options_Call) Return(_a0 client.SdkClientListUsersRequestInterface) *MockSdkClientListUsersRequestInterface_Options_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSdkClientListUsersRequestInterface_Options_Call) RunAndReturn(run func(client.ClientListUsersOptions) client.SdkClientListUsersRequestInterface) *MockSdkClientListUsersRequestInterface_Options_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockSdkClientListUsersRequestInterface creates a new instance of MockSdkClientListUsersRequestInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockSdkClientListUsersRequestInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockSdkClientListUsersRequestInterface {
	mock := &MockSdkClientListUsersRequestInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
