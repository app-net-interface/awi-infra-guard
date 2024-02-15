// Code generated by mockery v2.32.0. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	infrapb "github.com/app-net-interface/awi-infra-guard/grpc/go/infrapb"

	mock "github.com/stretchr/testify/mock"
)

// KubernetesServiceClient is an autogenerated mock type for the KubernetesServiceClient type
type KubernetesServiceClient struct {
	mock.Mock
}

type KubernetesServiceClient_Expecter struct {
	mock *mock.Mock
}

func (_m *KubernetesServiceClient) EXPECT() *KubernetesServiceClient_Expecter {
	return &KubernetesServiceClient_Expecter{mock: &_m.Mock}
}

// ListClusters provides a mock function with given fields: ctx, in, opts
func (_m *KubernetesServiceClient) ListClusters(ctx context.Context, in *infrapb.ListClustersRequest, opts ...grpc.CallOption) (*infrapb.ListClustersResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *infrapb.ListClustersResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *infrapb.ListClustersRequest, ...grpc.CallOption) (*infrapb.ListClustersResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *infrapb.ListClustersRequest, ...grpc.CallOption) *infrapb.ListClustersResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*infrapb.ListClustersResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *infrapb.ListClustersRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// KubernetesServiceClient_ListClusters_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListClusters'
type KubernetesServiceClient_ListClusters_Call struct {
	*mock.Call
}

// ListClusters is a helper method to define mock.On call
//   - ctx context.Context
//   - in *infrapb.ListClustersRequest
//   - opts ...grpc.CallOption
func (_e *KubernetesServiceClient_Expecter) ListClusters(ctx interface{}, in interface{}, opts ...interface{}) *KubernetesServiceClient_ListClusters_Call {
	return &KubernetesServiceClient_ListClusters_Call{Call: _e.mock.On("ListClusters",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *KubernetesServiceClient_ListClusters_Call) Run(run func(ctx context.Context, in *infrapb.ListClustersRequest, opts ...grpc.CallOption)) *KubernetesServiceClient_ListClusters_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*infrapb.ListClustersRequest), variadicArgs...)
	})
	return _c
}

func (_c *KubernetesServiceClient_ListClusters_Call) Return(_a0 *infrapb.ListClustersResponse, _a1 error) *KubernetesServiceClient_ListClusters_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *KubernetesServiceClient_ListClusters_Call) RunAndReturn(run func(context.Context, *infrapb.ListClustersRequest, ...grpc.CallOption) (*infrapb.ListClustersResponse, error)) *KubernetesServiceClient_ListClusters_Call {
	_c.Call.Return(run)
	return _c
}

// ListNamespaces provides a mock function with given fields: ctx, in, opts
func (_m *KubernetesServiceClient) ListNamespaces(ctx context.Context, in *infrapb.ListNamespacesRequest, opts ...grpc.CallOption) (*infrapb.ListNamespacesResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *infrapb.ListNamespacesResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *infrapb.ListNamespacesRequest, ...grpc.CallOption) (*infrapb.ListNamespacesResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *infrapb.ListNamespacesRequest, ...grpc.CallOption) *infrapb.ListNamespacesResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*infrapb.ListNamespacesResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *infrapb.ListNamespacesRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// KubernetesServiceClient_ListNamespaces_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListNamespaces'
type KubernetesServiceClient_ListNamespaces_Call struct {
	*mock.Call
}

// ListNamespaces is a helper method to define mock.On call
//   - ctx context.Context
//   - in *infrapb.ListNamespacesRequest
//   - opts ...grpc.CallOption
func (_e *KubernetesServiceClient_Expecter) ListNamespaces(ctx interface{}, in interface{}, opts ...interface{}) *KubernetesServiceClient_ListNamespaces_Call {
	return &KubernetesServiceClient_ListNamespaces_Call{Call: _e.mock.On("ListNamespaces",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *KubernetesServiceClient_ListNamespaces_Call) Run(run func(ctx context.Context, in *infrapb.ListNamespacesRequest, opts ...grpc.CallOption)) *KubernetesServiceClient_ListNamespaces_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*infrapb.ListNamespacesRequest), variadicArgs...)
	})
	return _c
}

func (_c *KubernetesServiceClient_ListNamespaces_Call) Return(_a0 *infrapb.ListNamespacesResponse, _a1 error) *KubernetesServiceClient_ListNamespaces_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *KubernetesServiceClient_ListNamespaces_Call) RunAndReturn(run func(context.Context, *infrapb.ListNamespacesRequest, ...grpc.CallOption) (*infrapb.ListNamespacesResponse, error)) *KubernetesServiceClient_ListNamespaces_Call {
	_c.Call.Return(run)
	return _c
}

// ListNodes provides a mock function with given fields: ctx, in, opts
func (_m *KubernetesServiceClient) ListNodes(ctx context.Context, in *infrapb.ListNodesRequest, opts ...grpc.CallOption) (*infrapb.ListNodesResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *infrapb.ListNodesResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *infrapb.ListNodesRequest, ...grpc.CallOption) (*infrapb.ListNodesResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *infrapb.ListNodesRequest, ...grpc.CallOption) *infrapb.ListNodesResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*infrapb.ListNodesResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *infrapb.ListNodesRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// KubernetesServiceClient_ListNodes_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListNodes'
type KubernetesServiceClient_ListNodes_Call struct {
	*mock.Call
}

// ListNodes is a helper method to define mock.On call
//   - ctx context.Context
//   - in *infrapb.ListNodesRequest
//   - opts ...grpc.CallOption
func (_e *KubernetesServiceClient_Expecter) ListNodes(ctx interface{}, in interface{}, opts ...interface{}) *KubernetesServiceClient_ListNodes_Call {
	return &KubernetesServiceClient_ListNodes_Call{Call: _e.mock.On("ListNodes",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *KubernetesServiceClient_ListNodes_Call) Run(run func(ctx context.Context, in *infrapb.ListNodesRequest, opts ...grpc.CallOption)) *KubernetesServiceClient_ListNodes_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*infrapb.ListNodesRequest), variadicArgs...)
	})
	return _c
}

func (_c *KubernetesServiceClient_ListNodes_Call) Return(_a0 *infrapb.ListNodesResponse, _a1 error) *KubernetesServiceClient_ListNodes_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *KubernetesServiceClient_ListNodes_Call) RunAndReturn(run func(context.Context, *infrapb.ListNodesRequest, ...grpc.CallOption) (*infrapb.ListNodesResponse, error)) *KubernetesServiceClient_ListNodes_Call {
	_c.Call.Return(run)
	return _c
}

// ListPods provides a mock function with given fields: ctx, in, opts
func (_m *KubernetesServiceClient) ListPods(ctx context.Context, in *infrapb.ListPodsRequest, opts ...grpc.CallOption) (*infrapb.ListPodsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *infrapb.ListPodsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *infrapb.ListPodsRequest, ...grpc.CallOption) (*infrapb.ListPodsResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *infrapb.ListPodsRequest, ...grpc.CallOption) *infrapb.ListPodsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*infrapb.ListPodsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *infrapb.ListPodsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// KubernetesServiceClient_ListPods_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListPods'
type KubernetesServiceClient_ListPods_Call struct {
	*mock.Call
}

// ListPods is a helper method to define mock.On call
//   - ctx context.Context
//   - in *infrapb.ListPodsRequest
//   - opts ...grpc.CallOption
func (_e *KubernetesServiceClient_Expecter) ListPods(ctx interface{}, in interface{}, opts ...interface{}) *KubernetesServiceClient_ListPods_Call {
	return &KubernetesServiceClient_ListPods_Call{Call: _e.mock.On("ListPods",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *KubernetesServiceClient_ListPods_Call) Run(run func(ctx context.Context, in *infrapb.ListPodsRequest, opts ...grpc.CallOption)) *KubernetesServiceClient_ListPods_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*infrapb.ListPodsRequest), variadicArgs...)
	})
	return _c
}

func (_c *KubernetesServiceClient_ListPods_Call) Return(_a0 *infrapb.ListPodsResponse, _a1 error) *KubernetesServiceClient_ListPods_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *KubernetesServiceClient_ListPods_Call) RunAndReturn(run func(context.Context, *infrapb.ListPodsRequest, ...grpc.CallOption) (*infrapb.ListPodsResponse, error)) *KubernetesServiceClient_ListPods_Call {
	_c.Call.Return(run)
	return _c
}

// ListPodsCIDRs provides a mock function with given fields: ctx, in, opts
func (_m *KubernetesServiceClient) ListPodsCIDRs(ctx context.Context, in *infrapb.ListPodsCIDRsRequest, opts ...grpc.CallOption) (*infrapb.ListPodsCIDRsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *infrapb.ListPodsCIDRsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *infrapb.ListPodsCIDRsRequest, ...grpc.CallOption) (*infrapb.ListPodsCIDRsResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *infrapb.ListPodsCIDRsRequest, ...grpc.CallOption) *infrapb.ListPodsCIDRsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*infrapb.ListPodsCIDRsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *infrapb.ListPodsCIDRsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// KubernetesServiceClient_ListPodsCIDRs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListPodsCIDRs'
type KubernetesServiceClient_ListPodsCIDRs_Call struct {
	*mock.Call
}

// ListPodsCIDRs is a helper method to define mock.On call
//   - ctx context.Context
//   - in *infrapb.ListPodsCIDRsRequest
//   - opts ...grpc.CallOption
func (_e *KubernetesServiceClient_Expecter) ListPodsCIDRs(ctx interface{}, in interface{}, opts ...interface{}) *KubernetesServiceClient_ListPodsCIDRs_Call {
	return &KubernetesServiceClient_ListPodsCIDRs_Call{Call: _e.mock.On("ListPodsCIDRs",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *KubernetesServiceClient_ListPodsCIDRs_Call) Run(run func(ctx context.Context, in *infrapb.ListPodsCIDRsRequest, opts ...grpc.CallOption)) *KubernetesServiceClient_ListPodsCIDRs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*infrapb.ListPodsCIDRsRequest), variadicArgs...)
	})
	return _c
}

func (_c *KubernetesServiceClient_ListPodsCIDRs_Call) Return(_a0 *infrapb.ListPodsCIDRsResponse, _a1 error) *KubernetesServiceClient_ListPodsCIDRs_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *KubernetesServiceClient_ListPodsCIDRs_Call) RunAndReturn(run func(context.Context, *infrapb.ListPodsCIDRsRequest, ...grpc.CallOption) (*infrapb.ListPodsCIDRsResponse, error)) *KubernetesServiceClient_ListPodsCIDRs_Call {
	_c.Call.Return(run)
	return _c
}

// ListServices provides a mock function with given fields: ctx, in, opts
func (_m *KubernetesServiceClient) ListServices(ctx context.Context, in *infrapb.ListServicesRequest, opts ...grpc.CallOption) (*infrapb.ListServicesResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *infrapb.ListServicesResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *infrapb.ListServicesRequest, ...grpc.CallOption) (*infrapb.ListServicesResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *infrapb.ListServicesRequest, ...grpc.CallOption) *infrapb.ListServicesResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*infrapb.ListServicesResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *infrapb.ListServicesRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// KubernetesServiceClient_ListServices_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListServices'
type KubernetesServiceClient_ListServices_Call struct {
	*mock.Call
}

// ListServices is a helper method to define mock.On call
//   - ctx context.Context
//   - in *infrapb.ListServicesRequest
//   - opts ...grpc.CallOption
func (_e *KubernetesServiceClient_Expecter) ListServices(ctx interface{}, in interface{}, opts ...interface{}) *KubernetesServiceClient_ListServices_Call {
	return &KubernetesServiceClient_ListServices_Call{Call: _e.mock.On("ListServices",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *KubernetesServiceClient_ListServices_Call) Run(run func(ctx context.Context, in *infrapb.ListServicesRequest, opts ...grpc.CallOption)) *KubernetesServiceClient_ListServices_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*infrapb.ListServicesRequest), variadicArgs...)
	})
	return _c
}

func (_c *KubernetesServiceClient_ListServices_Call) Return(_a0 *infrapb.ListServicesResponse, _a1 error) *KubernetesServiceClient_ListServices_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *KubernetesServiceClient_ListServices_Call) RunAndReturn(run func(context.Context, *infrapb.ListServicesRequest, ...grpc.CallOption) (*infrapb.ListServicesResponse, error)) *KubernetesServiceClient_ListServices_Call {
	_c.Call.Return(run)
	return _c
}

// ListServicesCIDRs provides a mock function with given fields: ctx, in, opts
func (_m *KubernetesServiceClient) ListServicesCIDRs(ctx context.Context, in *infrapb.ListServicesCIDRsRequest, opts ...grpc.CallOption) (*infrapb.ListServicesCIDRsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *infrapb.ListServicesCIDRsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *infrapb.ListServicesCIDRsRequest, ...grpc.CallOption) (*infrapb.ListServicesCIDRsResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *infrapb.ListServicesCIDRsRequest, ...grpc.CallOption) *infrapb.ListServicesCIDRsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*infrapb.ListServicesCIDRsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *infrapb.ListServicesCIDRsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// KubernetesServiceClient_ListServicesCIDRs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListServicesCIDRs'
type KubernetesServiceClient_ListServicesCIDRs_Call struct {
	*mock.Call
}

// ListServicesCIDRs is a helper method to define mock.On call
//   - ctx context.Context
//   - in *infrapb.ListServicesCIDRsRequest
//   - opts ...grpc.CallOption
func (_e *KubernetesServiceClient_Expecter) ListServicesCIDRs(ctx interface{}, in interface{}, opts ...interface{}) *KubernetesServiceClient_ListServicesCIDRs_Call {
	return &KubernetesServiceClient_ListServicesCIDRs_Call{Call: _e.mock.On("ListServicesCIDRs",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *KubernetesServiceClient_ListServicesCIDRs_Call) Run(run func(ctx context.Context, in *infrapb.ListServicesCIDRsRequest, opts ...grpc.CallOption)) *KubernetesServiceClient_ListServicesCIDRs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*infrapb.ListServicesCIDRsRequest), variadicArgs...)
	})
	return _c
}

func (_c *KubernetesServiceClient_ListServicesCIDRs_Call) Return(_a0 *infrapb.ListServicesCIDRsResponse, _a1 error) *KubernetesServiceClient_ListServicesCIDRs_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *KubernetesServiceClient_ListServicesCIDRs_Call) RunAndReturn(run func(context.Context, *infrapb.ListServicesCIDRsRequest, ...grpc.CallOption) (*infrapb.ListServicesCIDRsResponse, error)) *KubernetesServiceClient_ListServicesCIDRs_Call {
	_c.Call.Return(run)
	return _c
}

// NewKubernetesServiceClient creates a new instance of KubernetesServiceClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewKubernetesServiceClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *KubernetesServiceClient {
	mock := &KubernetesServiceClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
