// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package infrapb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CloudProviderServiceClient is the client API for CloudProviderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CloudProviderServiceClient interface {
	ListAccounts(ctx context.Context, in *ListAccountsRequest, opts ...grpc.CallOption) (*ListAccountsResponse, error)
	ListVPC(ctx context.Context, in *ListVPCRequest, opts ...grpc.CallOption) (*ListVPCResponse, error)
	ListInstances(ctx context.Context, in *ListInstancesRequest, opts ...grpc.CallOption) (*ListInstancesResponse, error)
	GetSubnet(ctx context.Context, in *GetSubnetRequest, opts ...grpc.CallOption) (*GetSubnetResponse, error)
	ListSubnets(ctx context.Context, in *ListSubnetsRequest, opts ...grpc.CallOption) (*ListSubnetsResponse, error)
	ListACLs(ctx context.Context, in *ListACLsRequest, opts ...grpc.CallOption) (*ListACLsResponse, error)
	ListSecurityGroups(ctx context.Context, in *ListSecurityGroupsRequest, opts ...grpc.CallOption) (*ListSecurityGroupsResponse, error)
	ListRouteTables(ctx context.Context, in *ListRouteTablesRequest, opts ...grpc.CallOption) (*ListRouteTablesResponse, error)
	ListNATGateways(ctx context.Context, in *ListNATGatewaysRequest, opts ...grpc.CallOption) (*ListNATGatewaysResponse, error)
	ListRouters(ctx context.Context, in *ListRoutersRequest, opts ...grpc.CallOption) (*ListRoutersResponse, error)
	ListInternetGateways(ctx context.Context, in *ListInternetGatewaysRequest, opts ...grpc.CallOption) (*ListInternetGatewaysResponse, error)
	ListVPCEndpoints(ctx context.Context, in *ListVPCEndpointsRequest, opts ...grpc.CallOption) (*ListVPCEndpointsResponse, error)
	GetVPCIDForCIDR(ctx context.Context, in *GetVPCIDForCIDRRequest, opts ...grpc.CallOption) (*GetVPCIDForCIDRResponse, error)
	GetCIDRsForLabels(ctx context.Context, in *GetCIDRsForLabelsRequest, opts ...grpc.CallOption) (*GetCIDRsForLabelsResponse, error)
	GetIPsForLabels(ctx context.Context, in *GetIPsForLabelsRequest, opts ...grpc.CallOption) (*GetIPsForLabelsResponse, error)
	GetInstancesForLabels(ctx context.Context, in *GetInstancesForLabelsRequest, opts ...grpc.CallOption) (*GetInstancesForLabelsResponse, error)
	GetVPCIDWithTag(ctx context.Context, in *GetVPCIDWithTagRequest, opts ...grpc.CallOption) (*GetVPCIDWithTagResponse, error)
	ListCloudClusters(ctx context.Context, in *ListCloudClustersRequest, opts ...grpc.CallOption) (*ListCloudClustersResponse, error)
	Summary(ctx context.Context, in *SummaryRequest, opts ...grpc.CallOption) (*SummaryResponse, error)
}

type cloudProviderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCloudProviderServiceClient(cc grpc.ClientConnInterface) CloudProviderServiceClient {
	return &cloudProviderServiceClient{cc}
}

func (c *cloudProviderServiceClient) ListAccounts(ctx context.Context, in *ListAccountsRequest, opts ...grpc.CallOption) (*ListAccountsResponse, error) {
	out := new(ListAccountsResponse)
	err := c.cc.Invoke(ctx, "/infra.CloudProviderService/ListAccounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudProviderServiceClient) ListVPC(ctx context.Context, in *ListVPCRequest, opts ...grpc.CallOption) (*ListVPCResponse, error) {
	out := new(ListVPCResponse)
	err := c.cc.Invoke(ctx, "/infra.CloudProviderService/ListVPC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudProviderServiceClient) ListInstances(ctx context.Context, in *ListInstancesRequest, opts ...grpc.CallOption) (*ListInstancesResponse, error) {
	out := new(ListInstancesResponse)
	err := c.cc.Invoke(ctx, "/infra.CloudProviderService/ListInstances", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudProviderServiceClient) GetSubnet(ctx context.Context, in *GetSubnetRequest, opts ...grpc.CallOption) (*GetSubnetResponse, error) {
	out := new(GetSubnetResponse)
	err := c.cc.Invoke(ctx, "/infra.CloudProviderService/GetSubnet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudProviderServiceClient) ListSubnets(ctx context.Context, in *ListSubnetsRequest, opts ...grpc.CallOption) (*ListSubnetsResponse, error) {
	out := new(ListSubnetsResponse)
	err := c.cc.Invoke(ctx, "/infra.CloudProviderService/ListSubnets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudProviderServiceClient) ListACLs(ctx context.Context, in *ListACLsRequest, opts ...grpc.CallOption) (*ListACLsResponse, error) {
	out := new(ListACLsResponse)
	err := c.cc.Invoke(ctx, "/infra.CloudProviderService/ListACLs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudProviderServiceClient) ListSecurityGroups(ctx context.Context, in *ListSecurityGroupsRequest, opts ...grpc.CallOption) (*ListSecurityGroupsResponse, error) {
	out := new(ListSecurityGroupsResponse)
	err := c.cc.Invoke(ctx, "/infra.CloudProviderService/ListSecurityGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudProviderServiceClient) ListRouteTables(ctx context.Context, in *ListRouteTablesRequest, opts ...grpc.CallOption) (*ListRouteTablesResponse, error) {
	out := new(ListRouteTablesResponse)
	err := c.cc.Invoke(ctx, "/infra.CloudProviderService/ListRouteTables", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudProviderServiceClient) ListNATGateways(ctx context.Context, in *ListNATGatewaysRequest, opts ...grpc.CallOption) (*ListNATGatewaysResponse, error) {
	out := new(ListNATGatewaysResponse)
	err := c.cc.Invoke(ctx, "/infra.CloudProviderService/ListNATGateways", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudProviderServiceClient) ListRouters(ctx context.Context, in *ListRoutersRequest, opts ...grpc.CallOption) (*ListRoutersResponse, error) {
	out := new(ListRoutersResponse)
	err := c.cc.Invoke(ctx, "/infra.CloudProviderService/ListRouters", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudProviderServiceClient) ListInternetGateways(ctx context.Context, in *ListInternetGatewaysRequest, opts ...grpc.CallOption) (*ListInternetGatewaysResponse, error) {
	out := new(ListInternetGatewaysResponse)
	err := c.cc.Invoke(ctx, "/infra.CloudProviderService/ListInternetGateways", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudProviderServiceClient) ListVPCEndpoints(ctx context.Context, in *ListVPCEndpointsRequest, opts ...grpc.CallOption) (*ListVPCEndpointsResponse, error) {
	out := new(ListVPCEndpointsResponse)
	err := c.cc.Invoke(ctx, "/infra.CloudProviderService/ListVPCEndpoints", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudProviderServiceClient) GetVPCIDForCIDR(ctx context.Context, in *GetVPCIDForCIDRRequest, opts ...grpc.CallOption) (*GetVPCIDForCIDRResponse, error) {
	out := new(GetVPCIDForCIDRResponse)
	err := c.cc.Invoke(ctx, "/infra.CloudProviderService/GetVPCIDForCIDR", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudProviderServiceClient) GetCIDRsForLabels(ctx context.Context, in *GetCIDRsForLabelsRequest, opts ...grpc.CallOption) (*GetCIDRsForLabelsResponse, error) {
	out := new(GetCIDRsForLabelsResponse)
	err := c.cc.Invoke(ctx, "/infra.CloudProviderService/GetCIDRsForLabels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudProviderServiceClient) GetIPsForLabels(ctx context.Context, in *GetIPsForLabelsRequest, opts ...grpc.CallOption) (*GetIPsForLabelsResponse, error) {
	out := new(GetIPsForLabelsResponse)
	err := c.cc.Invoke(ctx, "/infra.CloudProviderService/GetIPsForLabels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudProviderServiceClient) GetInstancesForLabels(ctx context.Context, in *GetInstancesForLabelsRequest, opts ...grpc.CallOption) (*GetInstancesForLabelsResponse, error) {
	out := new(GetInstancesForLabelsResponse)
	err := c.cc.Invoke(ctx, "/infra.CloudProviderService/GetInstancesForLabels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudProviderServiceClient) GetVPCIDWithTag(ctx context.Context, in *GetVPCIDWithTagRequest, opts ...grpc.CallOption) (*GetVPCIDWithTagResponse, error) {
	out := new(GetVPCIDWithTagResponse)
	err := c.cc.Invoke(ctx, "/infra.CloudProviderService/GetVPCIDWithTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudProviderServiceClient) ListCloudClusters(ctx context.Context, in *ListCloudClustersRequest, opts ...grpc.CallOption) (*ListCloudClustersResponse, error) {
	out := new(ListCloudClustersResponse)
	err := c.cc.Invoke(ctx, "/infra.CloudProviderService/ListCloudClusters", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudProviderServiceClient) Summary(ctx context.Context, in *SummaryRequest, opts ...grpc.CallOption) (*SummaryResponse, error) {
	out := new(SummaryResponse)
	err := c.cc.Invoke(ctx, "/infra.CloudProviderService/Summary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CloudProviderServiceServer is the server API for CloudProviderService service.
// All implementations must embed UnimplementedCloudProviderServiceServer
// for forward compatibility
type CloudProviderServiceServer interface {
	ListAccounts(context.Context, *ListAccountsRequest) (*ListAccountsResponse, error)
	ListVPC(context.Context, *ListVPCRequest) (*ListVPCResponse, error)
	ListInstances(context.Context, *ListInstancesRequest) (*ListInstancesResponse, error)
	GetSubnet(context.Context, *GetSubnetRequest) (*GetSubnetResponse, error)
	ListSubnets(context.Context, *ListSubnetsRequest) (*ListSubnetsResponse, error)
	ListACLs(context.Context, *ListACLsRequest) (*ListACLsResponse, error)
	ListSecurityGroups(context.Context, *ListSecurityGroupsRequest) (*ListSecurityGroupsResponse, error)
	ListRouteTables(context.Context, *ListRouteTablesRequest) (*ListRouteTablesResponse, error)
	ListNATGateways(context.Context, *ListNATGatewaysRequest) (*ListNATGatewaysResponse, error)
	ListRouters(context.Context, *ListRoutersRequest) (*ListRoutersResponse, error)
	ListInternetGateways(context.Context, *ListInternetGatewaysRequest) (*ListInternetGatewaysResponse, error)
	ListVPCEndpoints(context.Context, *ListVPCEndpointsRequest) (*ListVPCEndpointsResponse, error)
	GetVPCIDForCIDR(context.Context, *GetVPCIDForCIDRRequest) (*GetVPCIDForCIDRResponse, error)
	GetCIDRsForLabels(context.Context, *GetCIDRsForLabelsRequest) (*GetCIDRsForLabelsResponse, error)
	GetIPsForLabels(context.Context, *GetIPsForLabelsRequest) (*GetIPsForLabelsResponse, error)
	GetInstancesForLabels(context.Context, *GetInstancesForLabelsRequest) (*GetInstancesForLabelsResponse, error)
	GetVPCIDWithTag(context.Context, *GetVPCIDWithTagRequest) (*GetVPCIDWithTagResponse, error)
	ListCloudClusters(context.Context, *ListCloudClustersRequest) (*ListCloudClustersResponse, error)
	Summary(context.Context, *SummaryRequest) (*SummaryResponse, error)
	mustEmbedUnimplementedCloudProviderServiceServer()
}

// UnimplementedCloudProviderServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCloudProviderServiceServer struct {
}

func (UnimplementedCloudProviderServiceServer) ListAccounts(context.Context, *ListAccountsRequest) (*ListAccountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAccounts not implemented")
}
func (UnimplementedCloudProviderServiceServer) ListVPC(context.Context, *ListVPCRequest) (*ListVPCResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListVPC not implemented")
}
func (UnimplementedCloudProviderServiceServer) ListInstances(context.Context, *ListInstancesRequest) (*ListInstancesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListInstances not implemented")
}
func (UnimplementedCloudProviderServiceServer) GetSubnet(context.Context, *GetSubnetRequest) (*GetSubnetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubnet not implemented")
}
func (UnimplementedCloudProviderServiceServer) ListSubnets(context.Context, *ListSubnetsRequest) (*ListSubnetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSubnets not implemented")
}
func (UnimplementedCloudProviderServiceServer) ListACLs(context.Context, *ListACLsRequest) (*ListACLsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListACLs not implemented")
}
func (UnimplementedCloudProviderServiceServer) ListSecurityGroups(context.Context, *ListSecurityGroupsRequest) (*ListSecurityGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSecurityGroups not implemented")
}
func (UnimplementedCloudProviderServiceServer) ListRouteTables(context.Context, *ListRouteTablesRequest) (*ListRouteTablesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRouteTables not implemented")
}
func (UnimplementedCloudProviderServiceServer) ListNATGateways(context.Context, *ListNATGatewaysRequest) (*ListNATGatewaysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNATGateways not implemented")
}
func (UnimplementedCloudProviderServiceServer) ListRouters(context.Context, *ListRoutersRequest) (*ListRoutersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRouters not implemented")
}
func (UnimplementedCloudProviderServiceServer) ListInternetGateways(context.Context, *ListInternetGatewaysRequest) (*ListInternetGatewaysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListInternetGateways not implemented")
}
func (UnimplementedCloudProviderServiceServer) ListVPCEndpoints(context.Context, *ListVPCEndpointsRequest) (*ListVPCEndpointsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListVPCEndpoints not implemented")
}
func (UnimplementedCloudProviderServiceServer) GetVPCIDForCIDR(context.Context, *GetVPCIDForCIDRRequest) (*GetVPCIDForCIDRResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVPCIDForCIDR not implemented")
}
func (UnimplementedCloudProviderServiceServer) GetCIDRsForLabels(context.Context, *GetCIDRsForLabelsRequest) (*GetCIDRsForLabelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCIDRsForLabels not implemented")
}
func (UnimplementedCloudProviderServiceServer) GetIPsForLabels(context.Context, *GetIPsForLabelsRequest) (*GetIPsForLabelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIPsForLabels not implemented")
}
func (UnimplementedCloudProviderServiceServer) GetInstancesForLabels(context.Context, *GetInstancesForLabelsRequest) (*GetInstancesForLabelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInstancesForLabels not implemented")
}
func (UnimplementedCloudProviderServiceServer) GetVPCIDWithTag(context.Context, *GetVPCIDWithTagRequest) (*GetVPCIDWithTagResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVPCIDWithTag not implemented")
}
func (UnimplementedCloudProviderServiceServer) ListCloudClusters(context.Context, *ListCloudClustersRequest) (*ListCloudClustersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCloudClusters not implemented")
}
func (UnimplementedCloudProviderServiceServer) Summary(context.Context, *SummaryRequest) (*SummaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Summary not implemented")
}
func (UnimplementedCloudProviderServiceServer) mustEmbedUnimplementedCloudProviderServiceServer() {}

// UnsafeCloudProviderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CloudProviderServiceServer will
// result in compilation errors.
type UnsafeCloudProviderServiceServer interface {
	mustEmbedUnimplementedCloudProviderServiceServer()
}

func RegisterCloudProviderServiceServer(s grpc.ServiceRegistrar, srv CloudProviderServiceServer) {
	s.RegisterService(&CloudProviderService_ServiceDesc, srv)
}

func _CloudProviderService_ListAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAccountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudProviderServiceServer).ListAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infra.CloudProviderService/ListAccounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudProviderServiceServer).ListAccounts(ctx, req.(*ListAccountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudProviderService_ListVPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListVPCRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudProviderServiceServer).ListVPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infra.CloudProviderService/ListVPC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudProviderServiceServer).ListVPC(ctx, req.(*ListVPCRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudProviderService_ListInstances_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListInstancesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudProviderServiceServer).ListInstances(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infra.CloudProviderService/ListInstances",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudProviderServiceServer).ListInstances(ctx, req.(*ListInstancesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudProviderService_GetSubnet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSubnetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudProviderServiceServer).GetSubnet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infra.CloudProviderService/GetSubnet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudProviderServiceServer).GetSubnet(ctx, req.(*GetSubnetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudProviderService_ListSubnets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSubnetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudProviderServiceServer).ListSubnets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infra.CloudProviderService/ListSubnets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudProviderServiceServer).ListSubnets(ctx, req.(*ListSubnetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudProviderService_ListACLs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListACLsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudProviderServiceServer).ListACLs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infra.CloudProviderService/ListACLs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudProviderServiceServer).ListACLs(ctx, req.(*ListACLsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudProviderService_ListSecurityGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSecurityGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudProviderServiceServer).ListSecurityGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infra.CloudProviderService/ListSecurityGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudProviderServiceServer).ListSecurityGroups(ctx, req.(*ListSecurityGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudProviderService_ListRouteTables_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRouteTablesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudProviderServiceServer).ListRouteTables(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infra.CloudProviderService/ListRouteTables",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudProviderServiceServer).ListRouteTables(ctx, req.(*ListRouteTablesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudProviderService_ListNATGateways_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNATGatewaysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudProviderServiceServer).ListNATGateways(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infra.CloudProviderService/ListNATGateways",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudProviderServiceServer).ListNATGateways(ctx, req.(*ListNATGatewaysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudProviderService_ListRouters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRoutersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudProviderServiceServer).ListRouters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infra.CloudProviderService/ListRouters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudProviderServiceServer).ListRouters(ctx, req.(*ListRoutersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudProviderService_ListInternetGateways_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListInternetGatewaysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudProviderServiceServer).ListInternetGateways(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infra.CloudProviderService/ListInternetGateways",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudProviderServiceServer).ListInternetGateways(ctx, req.(*ListInternetGatewaysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudProviderService_ListVPCEndpoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListVPCEndpointsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudProviderServiceServer).ListVPCEndpoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infra.CloudProviderService/ListVPCEndpoints",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudProviderServiceServer).ListVPCEndpoints(ctx, req.(*ListVPCEndpointsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudProviderService_GetVPCIDForCIDR_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVPCIDForCIDRRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudProviderServiceServer).GetVPCIDForCIDR(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infra.CloudProviderService/GetVPCIDForCIDR",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudProviderServiceServer).GetVPCIDForCIDR(ctx, req.(*GetVPCIDForCIDRRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudProviderService_GetCIDRsForLabels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCIDRsForLabelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudProviderServiceServer).GetCIDRsForLabels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infra.CloudProviderService/GetCIDRsForLabels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudProviderServiceServer).GetCIDRsForLabels(ctx, req.(*GetCIDRsForLabelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudProviderService_GetIPsForLabels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIPsForLabelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudProviderServiceServer).GetIPsForLabels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infra.CloudProviderService/GetIPsForLabels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudProviderServiceServer).GetIPsForLabels(ctx, req.(*GetIPsForLabelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudProviderService_GetInstancesForLabels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInstancesForLabelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudProviderServiceServer).GetInstancesForLabels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infra.CloudProviderService/GetInstancesForLabels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudProviderServiceServer).GetInstancesForLabels(ctx, req.(*GetInstancesForLabelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudProviderService_GetVPCIDWithTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVPCIDWithTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudProviderServiceServer).GetVPCIDWithTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infra.CloudProviderService/GetVPCIDWithTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudProviderServiceServer).GetVPCIDWithTag(ctx, req.(*GetVPCIDWithTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudProviderService_ListCloudClusters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCloudClustersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudProviderServiceServer).ListCloudClusters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infra.CloudProviderService/ListCloudClusters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudProviderServiceServer).ListCloudClusters(ctx, req.(*ListCloudClustersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudProviderService_Summary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SummaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudProviderServiceServer).Summary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infra.CloudProviderService/Summary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudProviderServiceServer).Summary(ctx, req.(*SummaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CloudProviderService_ServiceDesc is the grpc.ServiceDesc for CloudProviderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CloudProviderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "infra.CloudProviderService",
	HandlerType: (*CloudProviderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListAccounts",
			Handler:    _CloudProviderService_ListAccounts_Handler,
		},
		{
			MethodName: "ListVPC",
			Handler:    _CloudProviderService_ListVPC_Handler,
		},
		{
			MethodName: "ListInstances",
			Handler:    _CloudProviderService_ListInstances_Handler,
		},
		{
			MethodName: "GetSubnet",
			Handler:    _CloudProviderService_GetSubnet_Handler,
		},
		{
			MethodName: "ListSubnets",
			Handler:    _CloudProviderService_ListSubnets_Handler,
		},
		{
			MethodName: "ListACLs",
			Handler:    _CloudProviderService_ListACLs_Handler,
		},
		{
			MethodName: "ListSecurityGroups",
			Handler:    _CloudProviderService_ListSecurityGroups_Handler,
		},
		{
			MethodName: "ListRouteTables",
			Handler:    _CloudProviderService_ListRouteTables_Handler,
		},
		{
			MethodName: "ListNATGateways",
			Handler:    _CloudProviderService_ListNATGateways_Handler,
		},
		{
			MethodName: "ListRouters",
			Handler:    _CloudProviderService_ListRouters_Handler,
		},
		{
			MethodName: "ListInternetGateways",
			Handler:    _CloudProviderService_ListInternetGateways_Handler,
		},
		{
			MethodName: "ListVPCEndpoints",
			Handler:    _CloudProviderService_ListVPCEndpoints_Handler,
		},
		{
			MethodName: "GetVPCIDForCIDR",
			Handler:    _CloudProviderService_GetVPCIDForCIDR_Handler,
		},
		{
			MethodName: "GetCIDRsForLabels",
			Handler:    _CloudProviderService_GetCIDRsForLabels_Handler,
		},
		{
			MethodName: "GetIPsForLabels",
			Handler:    _CloudProviderService_GetIPsForLabels_Handler,
		},
		{
			MethodName: "GetInstancesForLabels",
			Handler:    _CloudProviderService_GetInstancesForLabels_Handler,
		},
		{
			MethodName: "GetVPCIDWithTag",
			Handler:    _CloudProviderService_GetVPCIDWithTag_Handler,
		},
		{
			MethodName: "ListCloudClusters",
			Handler:    _CloudProviderService_ListCloudClusters_Handler,
		},
		{
			MethodName: "Summary",
			Handler:    _CloudProviderService_Summary_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud.proto",
}
