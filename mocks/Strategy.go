// Code generated by mockery v2.32.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	provider "github.com/app-net-interface/awi-infra-guard/provider"
)

// Strategy is an autogenerated mock type for the Strategy type
type Strategy struct {
	mock.Mock
}

type Strategy_Expecter struct {
	mock *mock.Mock
}

func (_m *Strategy) EXPECT() *Strategy_Expecter {
	return &Strategy_Expecter{mock: &_m.Mock}
}

// GetAllProviders provides a mock function with given fields:
func (_m *Strategy) GetAllProviders() []provider.CloudProvider {
	ret := _m.Called()

	var r0 []provider.CloudProvider
	if rf, ok := ret.Get(0).(func() []provider.CloudProvider); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]provider.CloudProvider)
		}
	}

	return r0
}

// Strategy_GetAllProviders_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllProviders'
type Strategy_GetAllProviders_Call struct {
	*mock.Call
}

// GetAllProviders is a helper method to define mock.On call
func (_e *Strategy_Expecter) GetAllProviders() *Strategy_GetAllProviders_Call {
	return &Strategy_GetAllProviders_Call{Call: _e.mock.On("GetAllProviders")}
}

func (_c *Strategy_GetAllProviders_Call) Run(run func()) *Strategy_GetAllProviders_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Strategy_GetAllProviders_Call) Return(_a0 []provider.CloudProvider) *Strategy_GetAllProviders_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Strategy_GetAllProviders_Call) RunAndReturn(run func() []provider.CloudProvider) *Strategy_GetAllProviders_Call {
	_c.Call.Return(run)
	return _c
}

// GetKubernetesProvider provides a mock function with given fields:
func (_m *Strategy) GetKubernetesProvider() (provider.Kubernetes, error) {
	ret := _m.Called()

	var r0 provider.Kubernetes
	var r1 error
	if rf, ok := ret.Get(0).(func() (provider.Kubernetes, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() provider.Kubernetes); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(provider.Kubernetes)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Strategy_GetKubernetesProvider_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetKubernetesProvider'
type Strategy_GetKubernetesProvider_Call struct {
	*mock.Call
}

// GetKubernetesProvider is a helper method to define mock.On call
func (_e *Strategy_Expecter) GetKubernetesProvider() *Strategy_GetKubernetesProvider_Call {
	return &Strategy_GetKubernetesProvider_Call{Call: _e.mock.On("GetKubernetesProvider")}
}

func (_c *Strategy_GetKubernetesProvider_Call) Run(run func()) *Strategy_GetKubernetesProvider_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Strategy_GetKubernetesProvider_Call) Return(_a0 provider.Kubernetes, _a1 error) *Strategy_GetKubernetesProvider_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Strategy_GetKubernetesProvider_Call) RunAndReturn(run func() (provider.Kubernetes, error)) *Strategy_GetKubernetesProvider_Call {
	_c.Call.Return(run)
	return _c
}

// GetProvider provides a mock function with given fields: ctx, cloud
func (_m *Strategy) GetProvider(ctx context.Context, cloud string) (provider.CloudProvider, error) {
	ret := _m.Called(ctx, cloud)

	var r0 provider.CloudProvider
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (provider.CloudProvider, error)); ok {
		return rf(ctx, cloud)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) provider.CloudProvider); ok {
		r0 = rf(ctx, cloud)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(provider.CloudProvider)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, cloud)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Strategy_GetProvider_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetProvider'
type Strategy_GetProvider_Call struct {
	*mock.Call
}

// GetProvider is a helper method to define mock.On call
//   - ctx context.Context
//   - cloud string
func (_e *Strategy_Expecter) GetProvider(ctx interface{}, cloud interface{}) *Strategy_GetProvider_Call {
	return &Strategy_GetProvider_Call{Call: _e.mock.On("GetProvider", ctx, cloud)}
}

func (_c *Strategy_GetProvider_Call) Run(run func(ctx context.Context, cloud string)) *Strategy_GetProvider_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *Strategy_GetProvider_Call) Return(_a0 provider.CloudProvider, _a1 error) *Strategy_GetProvider_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Strategy_GetProvider_Call) RunAndReturn(run func(context.Context, string) (provider.CloudProvider, error)) *Strategy_GetProvider_Call {
	_c.Call.Return(run)
	return _c
}

// RefreshState provides a mock function with given fields: ctx
func (_m *Strategy) RefreshState(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Strategy_RefreshState_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RefreshState'
type Strategy_RefreshState_Call struct {
	*mock.Call
}

// RefreshState is a helper method to define mock.On call
//   - ctx context.Context
func (_e *Strategy_Expecter) RefreshState(ctx interface{}) *Strategy_RefreshState_Call {
	return &Strategy_RefreshState_Call{Call: _e.mock.On("RefreshState", ctx)}
}

func (_c *Strategy_RefreshState_Call) Run(run func(ctx context.Context)) *Strategy_RefreshState_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Strategy_RefreshState_Call) Return(_a0 error) *Strategy_RefreshState_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Strategy_RefreshState_Call) RunAndReturn(run func(context.Context) error) *Strategy_RefreshState_Call {
	_c.Call.Return(run)
	return _c
}

// NewStrategy creates a new instance of Strategy. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStrategy(t interface {
	mock.TestingT
	Cleanup(func())
}) *Strategy {
	mock := &Strategy{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
