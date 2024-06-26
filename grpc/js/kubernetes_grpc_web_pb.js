/**
 * @fileoverview gRPC-Web generated client stub for infra
 * @enhanceable
 * @public
 */

// Code generated by protoc-gen-grpc-web. DO NOT EDIT.
// versions:
// 	protoc-gen-grpc-web v1.4.2
// 	protoc              v3.20.3
// source: kubernetes.proto


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');


var types_pb = require('./types_pb.js')
const proto = {};
proto.infra = require('./kubernetes_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.infra.KubernetesServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname.replace(/\/+$/, '');

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.infra.KubernetesServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname.replace(/\/+$/, '');

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.infra.ListClustersRequest,
 *   !proto.infra.ListClustersResponse>}
 */
const methodDescriptor_KubernetesService_ListClusters = new grpc.web.MethodDescriptor(
  '/infra.KubernetesService/ListClusters',
  grpc.web.MethodType.UNARY,
  proto.infra.ListClustersRequest,
  proto.infra.ListClustersResponse,
  /**
   * @param {!proto.infra.ListClustersRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.infra.ListClustersResponse.deserializeBinary
);


/**
 * @param {!proto.infra.ListClustersRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.infra.ListClustersResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.infra.ListClustersResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.infra.KubernetesServiceClient.prototype.listClusters =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/infra.KubernetesService/ListClusters',
      request,
      metadata || {},
      methodDescriptor_KubernetesService_ListClusters,
      callback);
};


/**
 * @param {!proto.infra.ListClustersRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.infra.ListClustersResponse>}
 *     Promise that resolves to the response
 */
proto.infra.KubernetesServicePromiseClient.prototype.listClusters =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/infra.KubernetesService/ListClusters',
      request,
      metadata || {},
      methodDescriptor_KubernetesService_ListClusters);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.infra.ListNamespacesRequest,
 *   !proto.infra.ListNamespacesResponse>}
 */
const methodDescriptor_KubernetesService_ListNamespaces = new grpc.web.MethodDescriptor(
  '/infra.KubernetesService/ListNamespaces',
  grpc.web.MethodType.UNARY,
  proto.infra.ListNamespacesRequest,
  proto.infra.ListNamespacesResponse,
  /**
   * @param {!proto.infra.ListNamespacesRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.infra.ListNamespacesResponse.deserializeBinary
);


/**
 * @param {!proto.infra.ListNamespacesRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.infra.ListNamespacesResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.infra.ListNamespacesResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.infra.KubernetesServiceClient.prototype.listNamespaces =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/infra.KubernetesService/ListNamespaces',
      request,
      metadata || {},
      methodDescriptor_KubernetesService_ListNamespaces,
      callback);
};


/**
 * @param {!proto.infra.ListNamespacesRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.infra.ListNamespacesResponse>}
 *     Promise that resolves to the response
 */
proto.infra.KubernetesServicePromiseClient.prototype.listNamespaces =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/infra.KubernetesService/ListNamespaces',
      request,
      metadata || {},
      methodDescriptor_KubernetesService_ListNamespaces);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.infra.ListNodesRequest,
 *   !proto.infra.ListNodesResponse>}
 */
const methodDescriptor_KubernetesService_ListNodes = new grpc.web.MethodDescriptor(
  '/infra.KubernetesService/ListNodes',
  grpc.web.MethodType.UNARY,
  proto.infra.ListNodesRequest,
  proto.infra.ListNodesResponse,
  /**
   * @param {!proto.infra.ListNodesRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.infra.ListNodesResponse.deserializeBinary
);


/**
 * @param {!proto.infra.ListNodesRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.infra.ListNodesResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.infra.ListNodesResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.infra.KubernetesServiceClient.prototype.listNodes =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/infra.KubernetesService/ListNodes',
      request,
      metadata || {},
      methodDescriptor_KubernetesService_ListNodes,
      callback);
};


/**
 * @param {!proto.infra.ListNodesRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.infra.ListNodesResponse>}
 *     Promise that resolves to the response
 */
proto.infra.KubernetesServicePromiseClient.prototype.listNodes =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/infra.KubernetesService/ListNodes',
      request,
      metadata || {},
      methodDescriptor_KubernetesService_ListNodes);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.infra.ListPodsRequest,
 *   !proto.infra.ListPodsResponse>}
 */
const methodDescriptor_KubernetesService_ListPods = new grpc.web.MethodDescriptor(
  '/infra.KubernetesService/ListPods',
  grpc.web.MethodType.UNARY,
  proto.infra.ListPodsRequest,
  proto.infra.ListPodsResponse,
  /**
   * @param {!proto.infra.ListPodsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.infra.ListPodsResponse.deserializeBinary
);


/**
 * @param {!proto.infra.ListPodsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.infra.ListPodsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.infra.ListPodsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.infra.KubernetesServiceClient.prototype.listPods =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/infra.KubernetesService/ListPods',
      request,
      metadata || {},
      methodDescriptor_KubernetesService_ListPods,
      callback);
};


/**
 * @param {!proto.infra.ListPodsRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.infra.ListPodsResponse>}
 *     Promise that resolves to the response
 */
proto.infra.KubernetesServicePromiseClient.prototype.listPods =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/infra.KubernetesService/ListPods',
      request,
      metadata || {},
      methodDescriptor_KubernetesService_ListPods);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.infra.ListServicesRequest,
 *   !proto.infra.ListServicesResponse>}
 */
const methodDescriptor_KubernetesService_ListServices = new grpc.web.MethodDescriptor(
  '/infra.KubernetesService/ListServices',
  grpc.web.MethodType.UNARY,
  proto.infra.ListServicesRequest,
  proto.infra.ListServicesResponse,
  /**
   * @param {!proto.infra.ListServicesRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.infra.ListServicesResponse.deserializeBinary
);


/**
 * @param {!proto.infra.ListServicesRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.infra.ListServicesResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.infra.ListServicesResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.infra.KubernetesServiceClient.prototype.listServices =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/infra.KubernetesService/ListServices',
      request,
      metadata || {},
      methodDescriptor_KubernetesService_ListServices,
      callback);
};


/**
 * @param {!proto.infra.ListServicesRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.infra.ListServicesResponse>}
 *     Promise that resolves to the response
 */
proto.infra.KubernetesServicePromiseClient.prototype.listServices =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/infra.KubernetesService/ListServices',
      request,
      metadata || {},
      methodDescriptor_KubernetesService_ListServices);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.infra.ListPodsCIDRsRequest,
 *   !proto.infra.ListPodsCIDRsResponse>}
 */
const methodDescriptor_KubernetesService_ListPodsCIDRs = new grpc.web.MethodDescriptor(
  '/infra.KubernetesService/ListPodsCIDRs',
  grpc.web.MethodType.UNARY,
  proto.infra.ListPodsCIDRsRequest,
  proto.infra.ListPodsCIDRsResponse,
  /**
   * @param {!proto.infra.ListPodsCIDRsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.infra.ListPodsCIDRsResponse.deserializeBinary
);


/**
 * @param {!proto.infra.ListPodsCIDRsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.infra.ListPodsCIDRsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.infra.ListPodsCIDRsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.infra.KubernetesServiceClient.prototype.listPodsCIDRs =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/infra.KubernetesService/ListPodsCIDRs',
      request,
      metadata || {},
      methodDescriptor_KubernetesService_ListPodsCIDRs,
      callback);
};


/**
 * @param {!proto.infra.ListPodsCIDRsRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.infra.ListPodsCIDRsResponse>}
 *     Promise that resolves to the response
 */
proto.infra.KubernetesServicePromiseClient.prototype.listPodsCIDRs =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/infra.KubernetesService/ListPodsCIDRs',
      request,
      metadata || {},
      methodDescriptor_KubernetesService_ListPodsCIDRs);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.infra.ListServicesCIDRsRequest,
 *   !proto.infra.ListServicesCIDRsResponse>}
 */
const methodDescriptor_KubernetesService_ListServicesCIDRs = new grpc.web.MethodDescriptor(
  '/infra.KubernetesService/ListServicesCIDRs',
  grpc.web.MethodType.UNARY,
  proto.infra.ListServicesCIDRsRequest,
  proto.infra.ListServicesCIDRsResponse,
  /**
   * @param {!proto.infra.ListServicesCIDRsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.infra.ListServicesCIDRsResponse.deserializeBinary
);


/**
 * @param {!proto.infra.ListServicesCIDRsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.infra.ListServicesCIDRsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.infra.ListServicesCIDRsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.infra.KubernetesServiceClient.prototype.listServicesCIDRs =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/infra.KubernetesService/ListServicesCIDRs',
      request,
      metadata || {},
      methodDescriptor_KubernetesService_ListServicesCIDRs,
      callback);
};


/**
 * @param {!proto.infra.ListServicesCIDRsRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.infra.ListServicesCIDRsResponse>}
 *     Promise that resolves to the response
 */
proto.infra.KubernetesServicePromiseClient.prototype.listServicesCIDRs =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/infra.KubernetesService/ListServicesCIDRs',
      request,
      metadata || {},
      methodDescriptor_KubernetesService_ListServicesCIDRs);
};


module.exports = proto.infra;

