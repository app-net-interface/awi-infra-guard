// Copyright (c) 2023 Cisco Systems, Inc. and its affiliates
// All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http:www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: access_control.proto

package infrapb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	AccessControlService_AddInboundAllowRuleInVPC_FullMethodName                = "/infra.AccessControlService/AddInboundAllowRuleInVPC"
	AccessControlService_AddInboundAllowRuleByLabelsMatch_FullMethodName        = "/infra.AccessControlService/AddInboundAllowRuleByLabelsMatch"
	AccessControlService_AddInboundAllowRuleBySubnetMatch_FullMethodName        = "/infra.AccessControlService/AddInboundAllowRuleBySubnetMatch"
	AccessControlService_AddInboundAllowRuleByInstanceIPMatch_FullMethodName    = "/infra.AccessControlService/AddInboundAllowRuleByInstanceIPMatch"
	AccessControlService_AddInboundAllowRuleForLoadBalancerByDNS_FullMethodName = "/infra.AccessControlService/AddInboundAllowRuleForLoadBalancerByDNS"
	AccessControlService_RemoveInboundAllowRuleFromVPCByName_FullMethodName     = "/infra.AccessControlService/RemoveInboundAllowRuleFromVPCByName"
	AccessControlService_RemoveInboundAllowRulesFromVPCById_FullMethodName      = "/infra.AccessControlService/RemoveInboundAllowRulesFromVPCById"
	AccessControlService_RemoveInboundAllowRuleRulesByTags_FullMethodName       = "/infra.AccessControlService/RemoveInboundAllowRuleRulesByTags"
	AccessControlService_RefreshInboundAllowRule_FullMethodName                 = "/infra.AccessControlService/RefreshInboundAllowRule"
)

// AccessControlServiceClient is the client API for AccessControlService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccessControlServiceClient interface {
	// AddInboundAllowRuleInVPC allows given cidrs in all VPC instances. Security rules are created with name ruleName
	// and tags if they are allowed in given cloud.
	AddInboundAllowRuleInVPC(ctx context.Context, in *AddInboundAllowRuleInVPCRequest, opts ...grpc.CallOption) (*AddInboundAllowRuleInVPCResponse, error)
	// AddInboundAllowRuleByLabelsMatch allows cidrsToAllow with protocolsAndPorts to all instances which match to labels
	AddInboundAllowRuleByLabelsMatch(ctx context.Context, in *AddInboundAllowRuleByLabelsMatchRequest, opts ...grpc.CallOption) (*AddInboundAllowRuleByLabelsMatchResponse, error)
	// AddInboundAllowRuleBySubnetMatch allows cidrsToAllow with protocolsAndPorts to all instances which are within provided cloud subnets
	AddInboundAllowRuleBySubnetMatch(ctx context.Context, in *AddInboundAllowRuleBySubnetMatchRequest, opts ...grpc.CallOption) (*AddInboundAllowRuleBySubnetMatchResponse, error)
	// AddInboundAllowRuleByInstanceIPMatch allows cidrsToAllow with protocolsAndPorts to all instances which have provided instancesIPs
	AddInboundAllowRuleByInstanceIPMatch(ctx context.Context, in *AddInboundAllowRuleByInstanceIPMatchRequest, opts ...grpc.CallOption) (*AddInboundAllowRuleByInstanceIPMatchResponse, error)
	// AddInboundAllowRuleForLoadBalancerByDNS allows cidrsToAllow with protocolsAndPorts to load balancer with given DNS
	AddInboundAllowRuleForLoadBalancerByDNS(ctx context.Context, in *AddInboundAllowRuleForLoadBalancerByDNSRequest, opts ...grpc.CallOption) (*AddInboundAllowRuleForLoadBalancerByDNSResponse, error)
	RemoveInboundAllowRuleFromVPCByName(ctx context.Context, in *RemoveInboundAllowRuleFromVPCByNameRequest, opts ...grpc.CallOption) (*RemoveInboundAllowRuleFromVPCByNameResponse, error)
	RemoveInboundAllowRulesFromVPCById(ctx context.Context, in *RemoveInboundAllowRulesFromVPCByIdRequest, opts ...grpc.CallOption) (*RemoveInboundAllowRulesFromVPCByIdResponse, error)
	RemoveInboundAllowRuleRulesByTags(ctx context.Context, in *RemoveInboundAllowRuleRulesByTagsRequest, opts ...grpc.CallOption) (*RemoveInboundAllowRuleRulesByTagsResponse, error)
	// RefreshInboundAllowRule adds and removes CIDRs in rule rules and applies rule in instances matching
	// to destinationLabels or destinationPrefixes
	RefreshInboundAllowRule(ctx context.Context, in *RefreshInboundAllowRuleRequest, opts ...grpc.CallOption) (*RefreshInboundAllowRuleResponse, error)
}

type accessControlServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAccessControlServiceClient(cc grpc.ClientConnInterface) AccessControlServiceClient {
	return &accessControlServiceClient{cc}
}

func (c *accessControlServiceClient) AddInboundAllowRuleInVPC(ctx context.Context, in *AddInboundAllowRuleInVPCRequest, opts ...grpc.CallOption) (*AddInboundAllowRuleInVPCResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddInboundAllowRuleInVPCResponse)
	err := c.cc.Invoke(ctx, AccessControlService_AddInboundAllowRuleInVPC_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessControlServiceClient) AddInboundAllowRuleByLabelsMatch(ctx context.Context, in *AddInboundAllowRuleByLabelsMatchRequest, opts ...grpc.CallOption) (*AddInboundAllowRuleByLabelsMatchResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddInboundAllowRuleByLabelsMatchResponse)
	err := c.cc.Invoke(ctx, AccessControlService_AddInboundAllowRuleByLabelsMatch_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessControlServiceClient) AddInboundAllowRuleBySubnetMatch(ctx context.Context, in *AddInboundAllowRuleBySubnetMatchRequest, opts ...grpc.CallOption) (*AddInboundAllowRuleBySubnetMatchResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddInboundAllowRuleBySubnetMatchResponse)
	err := c.cc.Invoke(ctx, AccessControlService_AddInboundAllowRuleBySubnetMatch_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessControlServiceClient) AddInboundAllowRuleByInstanceIPMatch(ctx context.Context, in *AddInboundAllowRuleByInstanceIPMatchRequest, opts ...grpc.CallOption) (*AddInboundAllowRuleByInstanceIPMatchResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddInboundAllowRuleByInstanceIPMatchResponse)
	err := c.cc.Invoke(ctx, AccessControlService_AddInboundAllowRuleByInstanceIPMatch_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessControlServiceClient) AddInboundAllowRuleForLoadBalancerByDNS(ctx context.Context, in *AddInboundAllowRuleForLoadBalancerByDNSRequest, opts ...grpc.CallOption) (*AddInboundAllowRuleForLoadBalancerByDNSResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddInboundAllowRuleForLoadBalancerByDNSResponse)
	err := c.cc.Invoke(ctx, AccessControlService_AddInboundAllowRuleForLoadBalancerByDNS_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessControlServiceClient) RemoveInboundAllowRuleFromVPCByName(ctx context.Context, in *RemoveInboundAllowRuleFromVPCByNameRequest, opts ...grpc.CallOption) (*RemoveInboundAllowRuleFromVPCByNameResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RemoveInboundAllowRuleFromVPCByNameResponse)
	err := c.cc.Invoke(ctx, AccessControlService_RemoveInboundAllowRuleFromVPCByName_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessControlServiceClient) RemoveInboundAllowRulesFromVPCById(ctx context.Context, in *RemoveInboundAllowRulesFromVPCByIdRequest, opts ...grpc.CallOption) (*RemoveInboundAllowRulesFromVPCByIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RemoveInboundAllowRulesFromVPCByIdResponse)
	err := c.cc.Invoke(ctx, AccessControlService_RemoveInboundAllowRulesFromVPCById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessControlServiceClient) RemoveInboundAllowRuleRulesByTags(ctx context.Context, in *RemoveInboundAllowRuleRulesByTagsRequest, opts ...grpc.CallOption) (*RemoveInboundAllowRuleRulesByTagsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RemoveInboundAllowRuleRulesByTagsResponse)
	err := c.cc.Invoke(ctx, AccessControlService_RemoveInboundAllowRuleRulesByTags_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessControlServiceClient) RefreshInboundAllowRule(ctx context.Context, in *RefreshInboundAllowRuleRequest, opts ...grpc.CallOption) (*RefreshInboundAllowRuleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RefreshInboundAllowRuleResponse)
	err := c.cc.Invoke(ctx, AccessControlService_RefreshInboundAllowRule_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccessControlServiceServer is the server API for AccessControlService service.
// All implementations must embed UnimplementedAccessControlServiceServer
// for forward compatibility
type AccessControlServiceServer interface {
	// AddInboundAllowRuleInVPC allows given cidrs in all VPC instances. Security rules are created with name ruleName
	// and tags if they are allowed in given cloud.
	AddInboundAllowRuleInVPC(context.Context, *AddInboundAllowRuleInVPCRequest) (*AddInboundAllowRuleInVPCResponse, error)
	// AddInboundAllowRuleByLabelsMatch allows cidrsToAllow with protocolsAndPorts to all instances which match to labels
	AddInboundAllowRuleByLabelsMatch(context.Context, *AddInboundAllowRuleByLabelsMatchRequest) (*AddInboundAllowRuleByLabelsMatchResponse, error)
	// AddInboundAllowRuleBySubnetMatch allows cidrsToAllow with protocolsAndPorts to all instances which are within provided cloud subnets
	AddInboundAllowRuleBySubnetMatch(context.Context, *AddInboundAllowRuleBySubnetMatchRequest) (*AddInboundAllowRuleBySubnetMatchResponse, error)
	// AddInboundAllowRuleByInstanceIPMatch allows cidrsToAllow with protocolsAndPorts to all instances which have provided instancesIPs
	AddInboundAllowRuleByInstanceIPMatch(context.Context, *AddInboundAllowRuleByInstanceIPMatchRequest) (*AddInboundAllowRuleByInstanceIPMatchResponse, error)
	// AddInboundAllowRuleForLoadBalancerByDNS allows cidrsToAllow with protocolsAndPorts to load balancer with given DNS
	AddInboundAllowRuleForLoadBalancerByDNS(context.Context, *AddInboundAllowRuleForLoadBalancerByDNSRequest) (*AddInboundAllowRuleForLoadBalancerByDNSResponse, error)
	RemoveInboundAllowRuleFromVPCByName(context.Context, *RemoveInboundAllowRuleFromVPCByNameRequest) (*RemoveInboundAllowRuleFromVPCByNameResponse, error)
	RemoveInboundAllowRulesFromVPCById(context.Context, *RemoveInboundAllowRulesFromVPCByIdRequest) (*RemoveInboundAllowRulesFromVPCByIdResponse, error)
	RemoveInboundAllowRuleRulesByTags(context.Context, *RemoveInboundAllowRuleRulesByTagsRequest) (*RemoveInboundAllowRuleRulesByTagsResponse, error)
	// RefreshInboundAllowRule adds and removes CIDRs in rule rules and applies rule in instances matching
	// to destinationLabels or destinationPrefixes
	RefreshInboundAllowRule(context.Context, *RefreshInboundAllowRuleRequest) (*RefreshInboundAllowRuleResponse, error)
	mustEmbedUnimplementedAccessControlServiceServer()
}

// UnimplementedAccessControlServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAccessControlServiceServer struct {
}

func (UnimplementedAccessControlServiceServer) AddInboundAllowRuleInVPC(context.Context, *AddInboundAllowRuleInVPCRequest) (*AddInboundAllowRuleInVPCResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddInboundAllowRuleInVPC not implemented")
}
func (UnimplementedAccessControlServiceServer) AddInboundAllowRuleByLabelsMatch(context.Context, *AddInboundAllowRuleByLabelsMatchRequest) (*AddInboundAllowRuleByLabelsMatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddInboundAllowRuleByLabelsMatch not implemented")
}
func (UnimplementedAccessControlServiceServer) AddInboundAllowRuleBySubnetMatch(context.Context, *AddInboundAllowRuleBySubnetMatchRequest) (*AddInboundAllowRuleBySubnetMatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddInboundAllowRuleBySubnetMatch not implemented")
}
func (UnimplementedAccessControlServiceServer) AddInboundAllowRuleByInstanceIPMatch(context.Context, *AddInboundAllowRuleByInstanceIPMatchRequest) (*AddInboundAllowRuleByInstanceIPMatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddInboundAllowRuleByInstanceIPMatch not implemented")
}
func (UnimplementedAccessControlServiceServer) AddInboundAllowRuleForLoadBalancerByDNS(context.Context, *AddInboundAllowRuleForLoadBalancerByDNSRequest) (*AddInboundAllowRuleForLoadBalancerByDNSResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddInboundAllowRuleForLoadBalancerByDNS not implemented")
}
func (UnimplementedAccessControlServiceServer) RemoveInboundAllowRuleFromVPCByName(context.Context, *RemoveInboundAllowRuleFromVPCByNameRequest) (*RemoveInboundAllowRuleFromVPCByNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveInboundAllowRuleFromVPCByName not implemented")
}
func (UnimplementedAccessControlServiceServer) RemoveInboundAllowRulesFromVPCById(context.Context, *RemoveInboundAllowRulesFromVPCByIdRequest) (*RemoveInboundAllowRulesFromVPCByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveInboundAllowRulesFromVPCById not implemented")
}
func (UnimplementedAccessControlServiceServer) RemoveInboundAllowRuleRulesByTags(context.Context, *RemoveInboundAllowRuleRulesByTagsRequest) (*RemoveInboundAllowRuleRulesByTagsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveInboundAllowRuleRulesByTags not implemented")
}
func (UnimplementedAccessControlServiceServer) RefreshInboundAllowRule(context.Context, *RefreshInboundAllowRuleRequest) (*RefreshInboundAllowRuleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshInboundAllowRule not implemented")
}
func (UnimplementedAccessControlServiceServer) mustEmbedUnimplementedAccessControlServiceServer() {}

// UnsafeAccessControlServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccessControlServiceServer will
// result in compilation errors.
type UnsafeAccessControlServiceServer interface {
	mustEmbedUnimplementedAccessControlServiceServer()
}

func RegisterAccessControlServiceServer(s grpc.ServiceRegistrar, srv AccessControlServiceServer) {
	s.RegisterService(&AccessControlService_ServiceDesc, srv)
}

func _AccessControlService_AddInboundAllowRuleInVPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddInboundAllowRuleInVPCRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessControlServiceServer).AddInboundAllowRuleInVPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccessControlService_AddInboundAllowRuleInVPC_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessControlServiceServer).AddInboundAllowRuleInVPC(ctx, req.(*AddInboundAllowRuleInVPCRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessControlService_AddInboundAllowRuleByLabelsMatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddInboundAllowRuleByLabelsMatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessControlServiceServer).AddInboundAllowRuleByLabelsMatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccessControlService_AddInboundAllowRuleByLabelsMatch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessControlServiceServer).AddInboundAllowRuleByLabelsMatch(ctx, req.(*AddInboundAllowRuleByLabelsMatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessControlService_AddInboundAllowRuleBySubnetMatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddInboundAllowRuleBySubnetMatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessControlServiceServer).AddInboundAllowRuleBySubnetMatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccessControlService_AddInboundAllowRuleBySubnetMatch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessControlServiceServer).AddInboundAllowRuleBySubnetMatch(ctx, req.(*AddInboundAllowRuleBySubnetMatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessControlService_AddInboundAllowRuleByInstanceIPMatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddInboundAllowRuleByInstanceIPMatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessControlServiceServer).AddInboundAllowRuleByInstanceIPMatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccessControlService_AddInboundAllowRuleByInstanceIPMatch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessControlServiceServer).AddInboundAllowRuleByInstanceIPMatch(ctx, req.(*AddInboundAllowRuleByInstanceIPMatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessControlService_AddInboundAllowRuleForLoadBalancerByDNS_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddInboundAllowRuleForLoadBalancerByDNSRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessControlServiceServer).AddInboundAllowRuleForLoadBalancerByDNS(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccessControlService_AddInboundAllowRuleForLoadBalancerByDNS_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessControlServiceServer).AddInboundAllowRuleForLoadBalancerByDNS(ctx, req.(*AddInboundAllowRuleForLoadBalancerByDNSRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessControlService_RemoveInboundAllowRuleFromVPCByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveInboundAllowRuleFromVPCByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessControlServiceServer).RemoveInboundAllowRuleFromVPCByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccessControlService_RemoveInboundAllowRuleFromVPCByName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessControlServiceServer).RemoveInboundAllowRuleFromVPCByName(ctx, req.(*RemoveInboundAllowRuleFromVPCByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessControlService_RemoveInboundAllowRulesFromVPCById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveInboundAllowRulesFromVPCByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessControlServiceServer).RemoveInboundAllowRulesFromVPCById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccessControlService_RemoveInboundAllowRulesFromVPCById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessControlServiceServer).RemoveInboundAllowRulesFromVPCById(ctx, req.(*RemoveInboundAllowRulesFromVPCByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessControlService_RemoveInboundAllowRuleRulesByTags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveInboundAllowRuleRulesByTagsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessControlServiceServer).RemoveInboundAllowRuleRulesByTags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccessControlService_RemoveInboundAllowRuleRulesByTags_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessControlServiceServer).RemoveInboundAllowRuleRulesByTags(ctx, req.(*RemoveInboundAllowRuleRulesByTagsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessControlService_RefreshInboundAllowRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshInboundAllowRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessControlServiceServer).RefreshInboundAllowRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccessControlService_RefreshInboundAllowRule_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessControlServiceServer).RefreshInboundAllowRule(ctx, req.(*RefreshInboundAllowRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AccessControlService_ServiceDesc is the grpc.ServiceDesc for AccessControlService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccessControlService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "infra.AccessControlService",
	HandlerType: (*AccessControlServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddInboundAllowRuleInVPC",
			Handler:    _AccessControlService_AddInboundAllowRuleInVPC_Handler,
		},
		{
			MethodName: "AddInboundAllowRuleByLabelsMatch",
			Handler:    _AccessControlService_AddInboundAllowRuleByLabelsMatch_Handler,
		},
		{
			MethodName: "AddInboundAllowRuleBySubnetMatch",
			Handler:    _AccessControlService_AddInboundAllowRuleBySubnetMatch_Handler,
		},
		{
			MethodName: "AddInboundAllowRuleByInstanceIPMatch",
			Handler:    _AccessControlService_AddInboundAllowRuleByInstanceIPMatch_Handler,
		},
		{
			MethodName: "AddInboundAllowRuleForLoadBalancerByDNS",
			Handler:    _AccessControlService_AddInboundAllowRuleForLoadBalancerByDNS_Handler,
		},
		{
			MethodName: "RemoveInboundAllowRuleFromVPCByName",
			Handler:    _AccessControlService_RemoveInboundAllowRuleFromVPCByName_Handler,
		},
		{
			MethodName: "RemoveInboundAllowRulesFromVPCById",
			Handler:    _AccessControlService_RemoveInboundAllowRulesFromVPCById_Handler,
		},
		{
			MethodName: "RemoveInboundAllowRuleRulesByTags",
			Handler:    _AccessControlService_RemoveInboundAllowRuleRulesByTags_Handler,
		},
		{
			MethodName: "RefreshInboundAllowRule",
			Handler:    _AccessControlService_RefreshInboundAllowRule_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "access_control.proto",
}
