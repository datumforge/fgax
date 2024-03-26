// Code generated by mockery. DO NOT EDIT.

package client

import (
	context "context"

	client "github.com/openfga/go-sdk/client"

	mock "github.com/stretchr/testify/mock"
)

// MockSdkClientListRelationsRequestInterface is an autogenerated mock type for the SdkClientListRelationsRequestInterface type
type MockSdkClientListRelationsRequestInterface struct {
	mock.Mock
}

type MockSdkClientListRelationsRequestInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockSdkClientListRelationsRequestInterface) EXPECT() *MockSdkClientListRelationsRequestInterface_Expecter {
	return &MockSdkClientListRelationsRequestInterface_Expecter{mock: &_m.Mock}
}

// Body provides a mock function with given fields: body
func (_m *MockSdkClientListRelationsRequestInterface) Body(body client.ClientListRelationsRequest) client.SdkClientListRelationsRequestInterface {
	ret := _m.Called(body)

	if len(ret) == 0 {
		panic("no return value specified for Body")
	}

	var r0 client.SdkClientListRelationsRequestInterface
	if rf, ok := ret.Get(0).(func(client.ClientListRelationsRequest) client.SdkClientListRelationsRequestInterface); ok {
		r0 = rf(body)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(client.SdkClientListRelationsRequestInterface)
		}
	}

	return r0
}

// MockSdkClientListRelationsRequestInterface_Body_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Body'
type MockSdkClientListRelationsRequestInterface_Body_Call struct {
	*mock.Call
}

// Body is a helper method to define mock.On call
//   - body client.ClientListRelationsRequest
func (_e *MockSdkClientListRelationsRequestInterface_Expecter) Body(body interface{}) *MockSdkClientListRelationsRequestInterface_Body_Call {
	return &MockSdkClientListRelationsRequestInterface_Body_Call{Call: _e.mock.On("Body", body)}
}

func (_c *MockSdkClientListRelationsRequestInterface_Body_Call) Run(run func(body client.ClientListRelationsRequest)) *MockSdkClientListRelationsRequestInterface_Body_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(client.ClientListRelationsRequest))
	})
	return _c
}

func (_c *MockSdkClientListRelationsRequestInterface_Body_Call) Return(_a0 client.SdkClientListRelationsRequestInterface) *MockSdkClientListRelationsRequestInterface_Body_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSdkClientListRelationsRequestInterface_Body_Call) RunAndReturn(run func(client.ClientListRelationsRequest) client.SdkClientListRelationsRequestInterface) *MockSdkClientListRelationsRequestInterface_Body_Call {
	_c.Call.Return(run)
	return _c
}

// Execute provides a mock function with given fields:
func (_m *MockSdkClientListRelationsRequestInterface) Execute() (*client.ClientListRelationsResponse, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Execute")
	}

	var r0 *client.ClientListRelationsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func() (*client.ClientListRelationsResponse, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *client.ClientListRelationsResponse); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.ClientListRelationsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSdkClientListRelationsRequestInterface_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type MockSdkClientListRelationsRequestInterface_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
func (_e *MockSdkClientListRelationsRequestInterface_Expecter) Execute() *MockSdkClientListRelationsRequestInterface_Execute_Call {
	return &MockSdkClientListRelationsRequestInterface_Execute_Call{Call: _e.mock.On("Execute")}
}

func (_c *MockSdkClientListRelationsRequestInterface_Execute_Call) Run(run func()) *MockSdkClientListRelationsRequestInterface_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSdkClientListRelationsRequestInterface_Execute_Call) Return(_a0 *client.ClientListRelationsResponse, _a1 error) *MockSdkClientListRelationsRequestInterface_Execute_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSdkClientListRelationsRequestInterface_Execute_Call) RunAndReturn(run func() (*client.ClientListRelationsResponse, error)) *MockSdkClientListRelationsRequestInterface_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// GetAuthorizationModelIdOverride provides a mock function with given fields:
func (_m *MockSdkClientListRelationsRequestInterface) GetAuthorizationModelIdOverride() *string {
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

// MockSdkClientListRelationsRequestInterface_GetAuthorizationModelIdOverride_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAuthorizationModelIdOverride'
type MockSdkClientListRelationsRequestInterface_GetAuthorizationModelIdOverride_Call struct {
	*mock.Call
}

// GetAuthorizationModelIdOverride is a helper method to define mock.On call
func (_e *MockSdkClientListRelationsRequestInterface_Expecter) GetAuthorizationModelIdOverride() *MockSdkClientListRelationsRequestInterface_GetAuthorizationModelIdOverride_Call {
	return &MockSdkClientListRelationsRequestInterface_GetAuthorizationModelIdOverride_Call{Call: _e.mock.On("GetAuthorizationModelIdOverride")}
}

func (_c *MockSdkClientListRelationsRequestInterface_GetAuthorizationModelIdOverride_Call) Run(run func()) *MockSdkClientListRelationsRequestInterface_GetAuthorizationModelIdOverride_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSdkClientListRelationsRequestInterface_GetAuthorizationModelIdOverride_Call) Return(_a0 *string) *MockSdkClientListRelationsRequestInterface_GetAuthorizationModelIdOverride_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSdkClientListRelationsRequestInterface_GetAuthorizationModelIdOverride_Call) RunAndReturn(run func() *string) *MockSdkClientListRelationsRequestInterface_GetAuthorizationModelIdOverride_Call {
	_c.Call.Return(run)
	return _c
}

// GetBody provides a mock function with given fields:
func (_m *MockSdkClientListRelationsRequestInterface) GetBody() *client.ClientListRelationsRequest {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetBody")
	}

	var r0 *client.ClientListRelationsRequest
	if rf, ok := ret.Get(0).(func() *client.ClientListRelationsRequest); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.ClientListRelationsRequest)
		}
	}

	return r0
}

// MockSdkClientListRelationsRequestInterface_GetBody_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBody'
type MockSdkClientListRelationsRequestInterface_GetBody_Call struct {
	*mock.Call
}

// GetBody is a helper method to define mock.On call
func (_e *MockSdkClientListRelationsRequestInterface_Expecter) GetBody() *MockSdkClientListRelationsRequestInterface_GetBody_Call {
	return &MockSdkClientListRelationsRequestInterface_GetBody_Call{Call: _e.mock.On("GetBody")}
}

func (_c *MockSdkClientListRelationsRequestInterface_GetBody_Call) Run(run func()) *MockSdkClientListRelationsRequestInterface_GetBody_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSdkClientListRelationsRequestInterface_GetBody_Call) Return(_a0 *client.ClientListRelationsRequest) *MockSdkClientListRelationsRequestInterface_GetBody_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSdkClientListRelationsRequestInterface_GetBody_Call) RunAndReturn(run func() *client.ClientListRelationsRequest) *MockSdkClientListRelationsRequestInterface_GetBody_Call {
	_c.Call.Return(run)
	return _c
}

// GetContext provides a mock function with given fields:
func (_m *MockSdkClientListRelationsRequestInterface) GetContext() context.Context {
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

// MockSdkClientListRelationsRequestInterface_GetContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetContext'
type MockSdkClientListRelationsRequestInterface_GetContext_Call struct {
	*mock.Call
}

// GetContext is a helper method to define mock.On call
func (_e *MockSdkClientListRelationsRequestInterface_Expecter) GetContext() *MockSdkClientListRelationsRequestInterface_GetContext_Call {
	return &MockSdkClientListRelationsRequestInterface_GetContext_Call{Call: _e.mock.On("GetContext")}
}

func (_c *MockSdkClientListRelationsRequestInterface_GetContext_Call) Run(run func()) *MockSdkClientListRelationsRequestInterface_GetContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSdkClientListRelationsRequestInterface_GetContext_Call) Return(_a0 context.Context) *MockSdkClientListRelationsRequestInterface_GetContext_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSdkClientListRelationsRequestInterface_GetContext_Call) RunAndReturn(run func() context.Context) *MockSdkClientListRelationsRequestInterface_GetContext_Call {
	_c.Call.Return(run)
	return _c
}

// GetOptions provides a mock function with given fields:
func (_m *MockSdkClientListRelationsRequestInterface) GetOptions() *client.ClientListRelationsOptions {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetOptions")
	}

	var r0 *client.ClientListRelationsOptions
	if rf, ok := ret.Get(0).(func() *client.ClientListRelationsOptions); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.ClientListRelationsOptions)
		}
	}

	return r0
}

// MockSdkClientListRelationsRequestInterface_GetOptions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetOptions'
type MockSdkClientListRelationsRequestInterface_GetOptions_Call struct {
	*mock.Call
}

// GetOptions is a helper method to define mock.On call
func (_e *MockSdkClientListRelationsRequestInterface_Expecter) GetOptions() *MockSdkClientListRelationsRequestInterface_GetOptions_Call {
	return &MockSdkClientListRelationsRequestInterface_GetOptions_Call{Call: _e.mock.On("GetOptions")}
}

func (_c *MockSdkClientListRelationsRequestInterface_GetOptions_Call) Run(run func()) *MockSdkClientListRelationsRequestInterface_GetOptions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSdkClientListRelationsRequestInterface_GetOptions_Call) Return(_a0 *client.ClientListRelationsOptions) *MockSdkClientListRelationsRequestInterface_GetOptions_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSdkClientListRelationsRequestInterface_GetOptions_Call) RunAndReturn(run func() *client.ClientListRelationsOptions) *MockSdkClientListRelationsRequestInterface_GetOptions_Call {
	_c.Call.Return(run)
	return _c
}

// Options provides a mock function with given fields: options
func (_m *MockSdkClientListRelationsRequestInterface) Options(options client.ClientListRelationsOptions) client.SdkClientListRelationsRequestInterface {
	ret := _m.Called(options)

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 client.SdkClientListRelationsRequestInterface
	if rf, ok := ret.Get(0).(func(client.ClientListRelationsOptions) client.SdkClientListRelationsRequestInterface); ok {
		r0 = rf(options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(client.SdkClientListRelationsRequestInterface)
		}
	}

	return r0
}

// MockSdkClientListRelationsRequestInterface_Options_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Options'
type MockSdkClientListRelationsRequestInterface_Options_Call struct {
	*mock.Call
}

// Options is a helper method to define mock.On call
//   - options client.ClientListRelationsOptions
func (_e *MockSdkClientListRelationsRequestInterface_Expecter) Options(options interface{}) *MockSdkClientListRelationsRequestInterface_Options_Call {
	return &MockSdkClientListRelationsRequestInterface_Options_Call{Call: _e.mock.On("Options", options)}
}

func (_c *MockSdkClientListRelationsRequestInterface_Options_Call) Run(run func(options client.ClientListRelationsOptions)) *MockSdkClientListRelationsRequestInterface_Options_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(client.ClientListRelationsOptions))
	})
	return _c
}

func (_c *MockSdkClientListRelationsRequestInterface_Options_Call) Return(_a0 client.SdkClientListRelationsRequestInterface) *MockSdkClientListRelationsRequestInterface_Options_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSdkClientListRelationsRequestInterface_Options_Call) RunAndReturn(run func(client.ClientListRelationsOptions) client.SdkClientListRelationsRequestInterface) *MockSdkClientListRelationsRequestInterface_Options_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockSdkClientListRelationsRequestInterface creates a new instance of MockSdkClientListRelationsRequestInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockSdkClientListRelationsRequestInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockSdkClientListRelationsRequestInterface {
	mock := &MockSdkClientListRelationsRequestInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
