/**
 * @fileoverview gRPC-Web generated client stub for infra
 * @enhanceable
 * @public
 */

// Code generated by protoc-gen-grpc-web. DO NOT EDIT.
// versions:
// 	protoc-gen-grpc-web v1.4.2
// 	protoc              v3.20.3
// source: cloud.proto


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as cloud_pb from './cloud_pb';


export class CloudProviderServiceClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: any; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname.replace(/\/+$/, '');
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodDescriptorListAccounts = new grpcWeb.MethodDescriptor(
    '/infra.CloudProviderService/ListAccounts',
    grpcWeb.MethodType.UNARY,
    cloud_pb.ListAccountsRequest,
    cloud_pb.ListAccountsResponse,
    (request: cloud_pb.ListAccountsRequest) => {
      return request.serializeBinary();
    },
    cloud_pb.ListAccountsResponse.deserializeBinary
  );

  listAccounts(
    request: cloud_pb.ListAccountsRequest,
    metadata: grpcWeb.Metadata | null): Promise<cloud_pb.ListAccountsResponse>;

  listAccounts(
    request: cloud_pb.ListAccountsRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: cloud_pb.ListAccountsResponse) => void): grpcWeb.ClientReadableStream<cloud_pb.ListAccountsResponse>;

  listAccounts(
    request: cloud_pb.ListAccountsRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: cloud_pb.ListAccountsResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/infra.CloudProviderService/ListAccounts',
        request,
        metadata || {},
        this.methodDescriptorListAccounts,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/infra.CloudProviderService/ListAccounts',
    request,
    metadata || {},
    this.methodDescriptorListAccounts);
  }

  methodDescriptorListVPC = new grpcWeb.MethodDescriptor(
    '/infra.CloudProviderService/ListVPC',
    grpcWeb.MethodType.UNARY,
    cloud_pb.ListVPCRequest,
    cloud_pb.ListVPCResponse,
    (request: cloud_pb.ListVPCRequest) => {
      return request.serializeBinary();
    },
    cloud_pb.ListVPCResponse.deserializeBinary
  );

  listVPC(
    request: cloud_pb.ListVPCRequest,
    metadata: grpcWeb.Metadata | null): Promise<cloud_pb.ListVPCResponse>;

  listVPC(
    request: cloud_pb.ListVPCRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: cloud_pb.ListVPCResponse) => void): grpcWeb.ClientReadableStream<cloud_pb.ListVPCResponse>;

  listVPC(
    request: cloud_pb.ListVPCRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: cloud_pb.ListVPCResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/infra.CloudProviderService/ListVPC',
        request,
        metadata || {},
        this.methodDescriptorListVPC,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/infra.CloudProviderService/ListVPC',
    request,
    metadata || {},
    this.methodDescriptorListVPC);
  }

  methodDescriptorListInstances = new grpcWeb.MethodDescriptor(
    '/infra.CloudProviderService/ListInstances',
    grpcWeb.MethodType.UNARY,
    cloud_pb.ListInstancesRequest,
    cloud_pb.ListInstancesResponse,
    (request: cloud_pb.ListInstancesRequest) => {
      return request.serializeBinary();
    },
    cloud_pb.ListInstancesResponse.deserializeBinary
  );

  listInstances(
    request: cloud_pb.ListInstancesRequest,
    metadata: grpcWeb.Metadata | null): Promise<cloud_pb.ListInstancesResponse>;

  listInstances(
    request: cloud_pb.ListInstancesRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: cloud_pb.ListInstancesResponse) => void): grpcWeb.ClientReadableStream<cloud_pb.ListInstancesResponse>;

  listInstances(
    request: cloud_pb.ListInstancesRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: cloud_pb.ListInstancesResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/infra.CloudProviderService/ListInstances',
        request,
        metadata || {},
        this.methodDescriptorListInstances,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/infra.CloudProviderService/ListInstances',
    request,
    metadata || {},
    this.methodDescriptorListInstances);
  }

  methodDescriptorGetSubnet = new grpcWeb.MethodDescriptor(
    '/infra.CloudProviderService/GetSubnet',
    grpcWeb.MethodType.UNARY,
    cloud_pb.GetSubnetRequest,
    cloud_pb.GetSubnetResponse,
    (request: cloud_pb.GetSubnetRequest) => {
      return request.serializeBinary();
    },
    cloud_pb.GetSubnetResponse.deserializeBinary
  );

  getSubnet(
    request: cloud_pb.GetSubnetRequest,
    metadata: grpcWeb.Metadata | null): Promise<cloud_pb.GetSubnetResponse>;

  getSubnet(
    request: cloud_pb.GetSubnetRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: cloud_pb.GetSubnetResponse) => void): grpcWeb.ClientReadableStream<cloud_pb.GetSubnetResponse>;

  getSubnet(
    request: cloud_pb.GetSubnetRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: cloud_pb.GetSubnetResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/infra.CloudProviderService/GetSubnet',
        request,
        metadata || {},
        this.methodDescriptorGetSubnet,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/infra.CloudProviderService/GetSubnet',
    request,
    metadata || {},
    this.methodDescriptorGetSubnet);
  }

  methodDescriptorListSubnets = new grpcWeb.MethodDescriptor(
    '/infra.CloudProviderService/ListSubnets',
    grpcWeb.MethodType.UNARY,
    cloud_pb.ListSubnetsRequest,
    cloud_pb.ListSubnetsResponse,
    (request: cloud_pb.ListSubnetsRequest) => {
      return request.serializeBinary();
    },
    cloud_pb.ListSubnetsResponse.deserializeBinary
  );

  listSubnets(
    request: cloud_pb.ListSubnetsRequest,
    metadata: grpcWeb.Metadata | null): Promise<cloud_pb.ListSubnetsResponse>;

  listSubnets(
    request: cloud_pb.ListSubnetsRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: cloud_pb.ListSubnetsResponse) => void): grpcWeb.ClientReadableStream<cloud_pb.ListSubnetsResponse>;

  listSubnets(
    request: cloud_pb.ListSubnetsRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: cloud_pb.ListSubnetsResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/infra.CloudProviderService/ListSubnets',
        request,
        metadata || {},
        this.methodDescriptorListSubnets,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/infra.CloudProviderService/ListSubnets',
    request,
    metadata || {},
    this.methodDescriptorListSubnets);
  }

  methodDescriptorListACLs = new grpcWeb.MethodDescriptor(
    '/infra.CloudProviderService/ListACLs',
    grpcWeb.MethodType.UNARY,
    cloud_pb.ListACLsRequest,
    cloud_pb.ListACLsResponse,
    (request: cloud_pb.ListACLsRequest) => {
      return request.serializeBinary();
    },
    cloud_pb.ListACLsResponse.deserializeBinary
  );

  listACLs(
    request: cloud_pb.ListACLsRequest,
    metadata: grpcWeb.Metadata | null): Promise<cloud_pb.ListACLsResponse>;

  listACLs(
    request: cloud_pb.ListACLsRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: cloud_pb.ListACLsResponse) => void): grpcWeb.ClientReadableStream<cloud_pb.ListACLsResponse>;

  listACLs(
    request: cloud_pb.ListACLsRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: cloud_pb.ListACLsResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/infra.CloudProviderService/ListACLs',
        request,
        metadata || {},
        this.methodDescriptorListACLs,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/infra.CloudProviderService/ListACLs',
    request,
    metadata || {},
    this.methodDescriptorListACLs);
  }

  methodDescriptorListSecurityGroups = new grpcWeb.MethodDescriptor(
    '/infra.CloudProviderService/ListSecurityGroups',
    grpcWeb.MethodType.UNARY,
    cloud_pb.ListSecurityGroupsRequest,
    cloud_pb.ListSecurityGroupsResponse,
    (request: cloud_pb.ListSecurityGroupsRequest) => {
      return request.serializeBinary();
    },
    cloud_pb.ListSecurityGroupsResponse.deserializeBinary
  );

  listSecurityGroups(
    request: cloud_pb.ListSecurityGroupsRequest,
    metadata: grpcWeb.Metadata | null): Promise<cloud_pb.ListSecurityGroupsResponse>;

  listSecurityGroups(
    request: cloud_pb.ListSecurityGroupsRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: cloud_pb.ListSecurityGroupsResponse) => void): grpcWeb.ClientReadableStream<cloud_pb.ListSecurityGroupsResponse>;

  listSecurityGroups(
    request: cloud_pb.ListSecurityGroupsRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: cloud_pb.ListSecurityGroupsResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/infra.CloudProviderService/ListSecurityGroups',
        request,
        metadata || {},
        this.methodDescriptorListSecurityGroups,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/infra.CloudProviderService/ListSecurityGroups',
    request,
    metadata || {},
    this.methodDescriptorListSecurityGroups);
  }

  methodDescriptorListRouteTables = new grpcWeb.MethodDescriptor(
    '/infra.CloudProviderService/ListRouteTables',
    grpcWeb.MethodType.UNARY,
    cloud_pb.ListRouteTablesRequest,
    cloud_pb.ListRouteTablesResponse,
    (request: cloud_pb.ListRouteTablesRequest) => {
      return request.serializeBinary();
    },
    cloud_pb.ListRouteTablesResponse.deserializeBinary
  );

  listRouteTables(
    request: cloud_pb.ListRouteTablesRequest,
    metadata: grpcWeb.Metadata | null): Promise<cloud_pb.ListRouteTablesResponse>;

  listRouteTables(
    request: cloud_pb.ListRouteTablesRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: cloud_pb.ListRouteTablesResponse) => void): grpcWeb.ClientReadableStream<cloud_pb.ListRouteTablesResponse>;

  listRouteTables(
    request: cloud_pb.ListRouteTablesRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: cloud_pb.ListRouteTablesResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/infra.CloudProviderService/ListRouteTables',
        request,
        metadata || {},
        this.methodDescriptorListRouteTables,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/infra.CloudProviderService/ListRouteTables',
    request,
    metadata || {},
    this.methodDescriptorListRouteTables);
  }

  methodDescriptorListNATGateways = new grpcWeb.MethodDescriptor(
    '/infra.CloudProviderService/ListNATGateways',
    grpcWeb.MethodType.UNARY,
    cloud_pb.ListNATGatewaysRequest,
    cloud_pb.ListNATGatewaysResponse,
    (request: cloud_pb.ListNATGatewaysRequest) => {
      return request.serializeBinary();
    },
    cloud_pb.ListNATGatewaysResponse.deserializeBinary
  );

  listNATGateways(
    request: cloud_pb.ListNATGatewaysRequest,
    metadata: grpcWeb.Metadata | null): Promise<cloud_pb.ListNATGatewaysResponse>;

  listNATGateways(
    request: cloud_pb.ListNATGatewaysRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: cloud_pb.ListNATGatewaysResponse) => void): grpcWeb.ClientReadableStream<cloud_pb.ListNATGatewaysResponse>;

  listNATGateways(
    request: cloud_pb.ListNATGatewaysRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: cloud_pb.ListNATGatewaysResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/infra.CloudProviderService/ListNATGateways',
        request,
        metadata || {},
        this.methodDescriptorListNATGateways,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/infra.CloudProviderService/ListNATGateways',
    request,
    metadata || {},
    this.methodDescriptorListNATGateways);
  }

  methodDescriptorListRouters = new grpcWeb.MethodDescriptor(
    '/infra.CloudProviderService/ListRouters',
    grpcWeb.MethodType.UNARY,
    cloud_pb.ListRoutersRequest,
    cloud_pb.ListRoutersResponse,
    (request: cloud_pb.ListRoutersRequest) => {
      return request.serializeBinary();
    },
    cloud_pb.ListRoutersResponse.deserializeBinary
  );

  listRouters(
    request: cloud_pb.ListRoutersRequest,
    metadata: grpcWeb.Metadata | null): Promise<cloud_pb.ListRoutersResponse>;

  listRouters(
    request: cloud_pb.ListRoutersRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: cloud_pb.ListRoutersResponse) => void): grpcWeb.ClientReadableStream<cloud_pb.ListRoutersResponse>;

  listRouters(
    request: cloud_pb.ListRoutersRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: cloud_pb.ListRoutersResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/infra.CloudProviderService/ListRouters',
        request,
        metadata || {},
        this.methodDescriptorListRouters,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/infra.CloudProviderService/ListRouters',
    request,
    metadata || {},
    this.methodDescriptorListRouters);
  }

  methodDescriptorListInternetGateways = new grpcWeb.MethodDescriptor(
    '/infra.CloudProviderService/ListInternetGateways',
    grpcWeb.MethodType.UNARY,
    cloud_pb.ListInternetGatewaysRequest,
    cloud_pb.ListInternetGatewaysResponse,
    (request: cloud_pb.ListInternetGatewaysRequest) => {
      return request.serializeBinary();
    },
    cloud_pb.ListInternetGatewaysResponse.deserializeBinary
  );

  listInternetGateways(
    request: cloud_pb.ListInternetGatewaysRequest,
    metadata: grpcWeb.Metadata | null): Promise<cloud_pb.ListInternetGatewaysResponse>;

  listInternetGateways(
    request: cloud_pb.ListInternetGatewaysRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: cloud_pb.ListInternetGatewaysResponse) => void): grpcWeb.ClientReadableStream<cloud_pb.ListInternetGatewaysResponse>;

  listInternetGateways(
    request: cloud_pb.ListInternetGatewaysRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: cloud_pb.ListInternetGatewaysResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/infra.CloudProviderService/ListInternetGateways',
        request,
        metadata || {},
        this.methodDescriptorListInternetGateways,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/infra.CloudProviderService/ListInternetGateways',
    request,
    metadata || {},
    this.methodDescriptorListInternetGateways);
  }

  methodDescriptorListVPCEndpoints = new grpcWeb.MethodDescriptor(
    '/infra.CloudProviderService/ListVPCEndpoints',
    grpcWeb.MethodType.UNARY,
    cloud_pb.ListVPCEndpointsRequest,
    cloud_pb.ListVPCEndpointsResponse,
    (request: cloud_pb.ListVPCEndpointsRequest) => {
      return request.serializeBinary();
    },
    cloud_pb.ListVPCEndpointsResponse.deserializeBinary
  );

  listVPCEndpoints(
    request: cloud_pb.ListVPCEndpointsRequest,
    metadata: grpcWeb.Metadata | null): Promise<cloud_pb.ListVPCEndpointsResponse>;

  listVPCEndpoints(
    request: cloud_pb.ListVPCEndpointsRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: cloud_pb.ListVPCEndpointsResponse) => void): grpcWeb.ClientReadableStream<cloud_pb.ListVPCEndpointsResponse>;

  listVPCEndpoints(
    request: cloud_pb.ListVPCEndpointsRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: cloud_pb.ListVPCEndpointsResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/infra.CloudProviderService/ListVPCEndpoints',
        request,
        metadata || {},
        this.methodDescriptorListVPCEndpoints,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/infra.CloudProviderService/ListVPCEndpoints',
    request,
    metadata || {},
    this.methodDescriptorListVPCEndpoints);
  }

  methodDescriptorGetVPCIDForCIDR = new grpcWeb.MethodDescriptor(
    '/infra.CloudProviderService/GetVPCIDForCIDR',
    grpcWeb.MethodType.UNARY,
    cloud_pb.GetVPCIDForCIDRRequest,
    cloud_pb.GetVPCIDForCIDRResponse,
    (request: cloud_pb.GetVPCIDForCIDRRequest) => {
      return request.serializeBinary();
    },
    cloud_pb.GetVPCIDForCIDRResponse.deserializeBinary
  );

  getVPCIDForCIDR(
    request: cloud_pb.GetVPCIDForCIDRRequest,
    metadata: grpcWeb.Metadata | null): Promise<cloud_pb.GetVPCIDForCIDRResponse>;

  getVPCIDForCIDR(
    request: cloud_pb.GetVPCIDForCIDRRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: cloud_pb.GetVPCIDForCIDRResponse) => void): grpcWeb.ClientReadableStream<cloud_pb.GetVPCIDForCIDRResponse>;

  getVPCIDForCIDR(
    request: cloud_pb.GetVPCIDForCIDRRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: cloud_pb.GetVPCIDForCIDRResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/infra.CloudProviderService/GetVPCIDForCIDR',
        request,
        metadata || {},
        this.methodDescriptorGetVPCIDForCIDR,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/infra.CloudProviderService/GetVPCIDForCIDR',
    request,
    metadata || {},
    this.methodDescriptorGetVPCIDForCIDR);
  }

  methodDescriptorGetCIDRsForLabels = new grpcWeb.MethodDescriptor(
    '/infra.CloudProviderService/GetCIDRsForLabels',
    grpcWeb.MethodType.UNARY,
    cloud_pb.GetCIDRsForLabelsRequest,
    cloud_pb.GetCIDRsForLabelsResponse,
    (request: cloud_pb.GetCIDRsForLabelsRequest) => {
      return request.serializeBinary();
    },
    cloud_pb.GetCIDRsForLabelsResponse.deserializeBinary
  );

  getCIDRsForLabels(
    request: cloud_pb.GetCIDRsForLabelsRequest,
    metadata: grpcWeb.Metadata | null): Promise<cloud_pb.GetCIDRsForLabelsResponse>;

  getCIDRsForLabels(
    request: cloud_pb.GetCIDRsForLabelsRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: cloud_pb.GetCIDRsForLabelsResponse) => void): grpcWeb.ClientReadableStream<cloud_pb.GetCIDRsForLabelsResponse>;

  getCIDRsForLabels(
    request: cloud_pb.GetCIDRsForLabelsRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: cloud_pb.GetCIDRsForLabelsResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/infra.CloudProviderService/GetCIDRsForLabels',
        request,
        metadata || {},
        this.methodDescriptorGetCIDRsForLabels,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/infra.CloudProviderService/GetCIDRsForLabels',
    request,
    metadata || {},
    this.methodDescriptorGetCIDRsForLabels);
  }

  methodDescriptorGetIPsForLabels = new grpcWeb.MethodDescriptor(
    '/infra.CloudProviderService/GetIPsForLabels',
    grpcWeb.MethodType.UNARY,
    cloud_pb.GetIPsForLabelsRequest,
    cloud_pb.GetIPsForLabelsResponse,
    (request: cloud_pb.GetIPsForLabelsRequest) => {
      return request.serializeBinary();
    },
    cloud_pb.GetIPsForLabelsResponse.deserializeBinary
  );

  getIPsForLabels(
    request: cloud_pb.GetIPsForLabelsRequest,
    metadata: grpcWeb.Metadata | null): Promise<cloud_pb.GetIPsForLabelsResponse>;

  getIPsForLabels(
    request: cloud_pb.GetIPsForLabelsRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: cloud_pb.GetIPsForLabelsResponse) => void): grpcWeb.ClientReadableStream<cloud_pb.GetIPsForLabelsResponse>;

  getIPsForLabels(
    request: cloud_pb.GetIPsForLabelsRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: cloud_pb.GetIPsForLabelsResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/infra.CloudProviderService/GetIPsForLabels',
        request,
        metadata || {},
        this.methodDescriptorGetIPsForLabels,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/infra.CloudProviderService/GetIPsForLabels',
    request,
    metadata || {},
    this.methodDescriptorGetIPsForLabels);
  }

  methodDescriptorGetInstancesForLabels = new grpcWeb.MethodDescriptor(
    '/infra.CloudProviderService/GetInstancesForLabels',
    grpcWeb.MethodType.UNARY,
    cloud_pb.GetInstancesForLabelsRequest,
    cloud_pb.GetInstancesForLabelsResponse,
    (request: cloud_pb.GetInstancesForLabelsRequest) => {
      return request.serializeBinary();
    },
    cloud_pb.GetInstancesForLabelsResponse.deserializeBinary
  );

  getInstancesForLabels(
    request: cloud_pb.GetInstancesForLabelsRequest,
    metadata: grpcWeb.Metadata | null): Promise<cloud_pb.GetInstancesForLabelsResponse>;

  getInstancesForLabels(
    request: cloud_pb.GetInstancesForLabelsRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: cloud_pb.GetInstancesForLabelsResponse) => void): grpcWeb.ClientReadableStream<cloud_pb.GetInstancesForLabelsResponse>;

  getInstancesForLabels(
    request: cloud_pb.GetInstancesForLabelsRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: cloud_pb.GetInstancesForLabelsResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/infra.CloudProviderService/GetInstancesForLabels',
        request,
        metadata || {},
        this.methodDescriptorGetInstancesForLabels,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/infra.CloudProviderService/GetInstancesForLabels',
    request,
    metadata || {},
    this.methodDescriptorGetInstancesForLabels);
  }

  methodDescriptorGetVPCIDWithTag = new grpcWeb.MethodDescriptor(
    '/infra.CloudProviderService/GetVPCIDWithTag',
    grpcWeb.MethodType.UNARY,
    cloud_pb.GetVPCIDWithTagRequest,
    cloud_pb.GetVPCIDWithTagResponse,
    (request: cloud_pb.GetVPCIDWithTagRequest) => {
      return request.serializeBinary();
    },
    cloud_pb.GetVPCIDWithTagResponse.deserializeBinary
  );

  getVPCIDWithTag(
    request: cloud_pb.GetVPCIDWithTagRequest,
    metadata: grpcWeb.Metadata | null): Promise<cloud_pb.GetVPCIDWithTagResponse>;

  getVPCIDWithTag(
    request: cloud_pb.GetVPCIDWithTagRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: cloud_pb.GetVPCIDWithTagResponse) => void): grpcWeb.ClientReadableStream<cloud_pb.GetVPCIDWithTagResponse>;

  getVPCIDWithTag(
    request: cloud_pb.GetVPCIDWithTagRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: cloud_pb.GetVPCIDWithTagResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/infra.CloudProviderService/GetVPCIDWithTag',
        request,
        metadata || {},
        this.methodDescriptorGetVPCIDWithTag,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/infra.CloudProviderService/GetVPCIDWithTag',
    request,
    metadata || {},
    this.methodDescriptorGetVPCIDWithTag);
  }

  methodDescriptorListCloudClusters = new grpcWeb.MethodDescriptor(
    '/infra.CloudProviderService/ListCloudClusters',
    grpcWeb.MethodType.UNARY,
    cloud_pb.ListCloudClustersRequest,
    cloud_pb.ListCloudClustersResponse,
    (request: cloud_pb.ListCloudClustersRequest) => {
      return request.serializeBinary();
    },
    cloud_pb.ListCloudClustersResponse.deserializeBinary
  );

  listCloudClusters(
    request: cloud_pb.ListCloudClustersRequest,
    metadata: grpcWeb.Metadata | null): Promise<cloud_pb.ListCloudClustersResponse>;

  listCloudClusters(
    request: cloud_pb.ListCloudClustersRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: cloud_pb.ListCloudClustersResponse) => void): grpcWeb.ClientReadableStream<cloud_pb.ListCloudClustersResponse>;

  listCloudClusters(
    request: cloud_pb.ListCloudClustersRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: cloud_pb.ListCloudClustersResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/infra.CloudProviderService/ListCloudClusters',
        request,
        metadata || {},
        this.methodDescriptorListCloudClusters,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/infra.CloudProviderService/ListCloudClusters',
    request,
    metadata || {},
    this.methodDescriptorListCloudClusters);
  }

  methodDescriptorSummary = new grpcWeb.MethodDescriptor(
    '/infra.CloudProviderService/Summary',
    grpcWeb.MethodType.UNARY,
    cloud_pb.SummaryRequest,
    cloud_pb.SummaryResponse,
    (request: cloud_pb.SummaryRequest) => {
      return request.serializeBinary();
    },
    cloud_pb.SummaryResponse.deserializeBinary
  );

  summary(
    request: cloud_pb.SummaryRequest,
    metadata: grpcWeb.Metadata | null): Promise<cloud_pb.SummaryResponse>;

  summary(
    request: cloud_pb.SummaryRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: cloud_pb.SummaryResponse) => void): grpcWeb.ClientReadableStream<cloud_pb.SummaryResponse>;

  summary(
    request: cloud_pb.SummaryRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: cloud_pb.SummaryResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/infra.CloudProviderService/Summary',
        request,
        metadata || {},
        this.methodDescriptorSummary,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/infra.CloudProviderService/Summary',
    request,
    metadata || {},
    this.methodDescriptorSummary);
  }

}

