/**
 * @fileoverview gRPC-Web generated client stub for infra
 * @enhanceable
 * @public
 */

// Code generated by protoc-gen-grpc-web. DO NOT EDIT.
// versions:
// 	protoc-gen-grpc-web v1.4.2
// 	protoc              v4.22.2
// source: access_control.proto


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');


var types_pb = require('./types_pb.js')
const proto = {};
proto.infra = require('./access_control_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.infra.AccessControlServiceClient =
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
proto.infra.AccessControlServicePromiseClient =
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
 *   !proto.infra.AddInboundAllowRuleInVPCRequest,
 *   !proto.infra.AddInboundAllowRuleInVPCResponse>}
 */
const methodDescriptor_AccessControlService_AddInboundAllowRuleInVPC = new grpc.web.MethodDescriptor(
  '/infra.AccessControlService/AddInboundAllowRuleInVPC',
  grpc.web.MethodType.UNARY,
  proto.infra.AddInboundAllowRuleInVPCRequest,
  proto.infra.AddInboundAllowRuleInVPCResponse,
  /**
   * @param {!proto.infra.AddInboundAllowRuleInVPCRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.infra.AddInboundAllowRuleInVPCResponse.deserializeBinary
);


/**
 * @param {!proto.infra.AddInboundAllowRuleInVPCRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.infra.AddInboundAllowRuleInVPCResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.infra.AddInboundAllowRuleInVPCResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.infra.AccessControlServiceClient.prototype.addInboundAllowRuleInVPC =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/infra.AccessControlService/AddInboundAllowRuleInVPC',
      request,
      metadata || {},
      methodDescriptor_AccessControlService_AddInboundAllowRuleInVPC,
      callback);
};


/**
 * @param {!proto.infra.AddInboundAllowRuleInVPCRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.infra.AddInboundAllowRuleInVPCResponse>}
 *     Promise that resolves to the response
 */
proto.infra.AccessControlServicePromiseClient.prototype.addInboundAllowRuleInVPC =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/infra.AccessControlService/AddInboundAllowRuleInVPC',
      request,
      metadata || {},
      methodDescriptor_AccessControlService_AddInboundAllowRuleInVPC);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.infra.AddInboundAllowRuleByLabelsMatchRequest,
 *   !proto.infra.AddInboundAllowRuleByLabelsMatchResponse>}
 */
const methodDescriptor_AccessControlService_AddInboundAllowRuleByLabelsMatch = new grpc.web.MethodDescriptor(
  '/infra.AccessControlService/AddInboundAllowRuleByLabelsMatch',
  grpc.web.MethodType.UNARY,
  proto.infra.AddInboundAllowRuleByLabelsMatchRequest,
  proto.infra.AddInboundAllowRuleByLabelsMatchResponse,
  /**
   * @param {!proto.infra.AddInboundAllowRuleByLabelsMatchRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.infra.AddInboundAllowRuleByLabelsMatchResponse.deserializeBinary
);


/**
 * @param {!proto.infra.AddInboundAllowRuleByLabelsMatchRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.infra.AddInboundAllowRuleByLabelsMatchResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.infra.AddInboundAllowRuleByLabelsMatchResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.infra.AccessControlServiceClient.prototype.addInboundAllowRuleByLabelsMatch =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/infra.AccessControlService/AddInboundAllowRuleByLabelsMatch',
      request,
      metadata || {},
      methodDescriptor_AccessControlService_AddInboundAllowRuleByLabelsMatch,
      callback);
};


/**
 * @param {!proto.infra.AddInboundAllowRuleByLabelsMatchRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.infra.AddInboundAllowRuleByLabelsMatchResponse>}
 *     Promise that resolves to the response
 */
proto.infra.AccessControlServicePromiseClient.prototype.addInboundAllowRuleByLabelsMatch =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/infra.AccessControlService/AddInboundAllowRuleByLabelsMatch',
      request,
      metadata || {},
      methodDescriptor_AccessControlService_AddInboundAllowRuleByLabelsMatch);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.infra.AddInboundAllowRuleBySubnetMatchRequest,
 *   !proto.infra.AddInboundAllowRuleBySubnetMatchResponse>}
 */
const methodDescriptor_AccessControlService_AddInboundAllowRuleBySubnetMatch = new grpc.web.MethodDescriptor(
  '/infra.AccessControlService/AddInboundAllowRuleBySubnetMatch',
  grpc.web.MethodType.UNARY,
  proto.infra.AddInboundAllowRuleBySubnetMatchRequest,
  proto.infra.AddInboundAllowRuleBySubnetMatchResponse,
  /**
   * @param {!proto.infra.AddInboundAllowRuleBySubnetMatchRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.infra.AddInboundAllowRuleBySubnetMatchResponse.deserializeBinary
);


/**
 * @param {!proto.infra.AddInboundAllowRuleBySubnetMatchRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.infra.AddInboundAllowRuleBySubnetMatchResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.infra.AddInboundAllowRuleBySubnetMatchResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.infra.AccessControlServiceClient.prototype.addInboundAllowRuleBySubnetMatch =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/infra.AccessControlService/AddInboundAllowRuleBySubnetMatch',
      request,
      metadata || {},
      methodDescriptor_AccessControlService_AddInboundAllowRuleBySubnetMatch,
      callback);
};


/**
 * @param {!proto.infra.AddInboundAllowRuleBySubnetMatchRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.infra.AddInboundAllowRuleBySubnetMatchResponse>}
 *     Promise that resolves to the response
 */
proto.infra.AccessControlServicePromiseClient.prototype.addInboundAllowRuleBySubnetMatch =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/infra.AccessControlService/AddInboundAllowRuleBySubnetMatch',
      request,
      metadata || {},
      methodDescriptor_AccessControlService_AddInboundAllowRuleBySubnetMatch);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.infra.AddInboundAllowRuleByInstanceIPMatchRequest,
 *   !proto.infra.AddInboundAllowRuleByInstanceIPMatchResponse>}
 */
const methodDescriptor_AccessControlService_AddInboundAllowRuleByInstanceIPMatch = new grpc.web.MethodDescriptor(
  '/infra.AccessControlService/AddInboundAllowRuleByInstanceIPMatch',
  grpc.web.MethodType.UNARY,
  proto.infra.AddInboundAllowRuleByInstanceIPMatchRequest,
  proto.infra.AddInboundAllowRuleByInstanceIPMatchResponse,
  /**
   * @param {!proto.infra.AddInboundAllowRuleByInstanceIPMatchRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.infra.AddInboundAllowRuleByInstanceIPMatchResponse.deserializeBinary
);


/**
 * @param {!proto.infra.AddInboundAllowRuleByInstanceIPMatchRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.infra.AddInboundAllowRuleByInstanceIPMatchResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.infra.AddInboundAllowRuleByInstanceIPMatchResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.infra.AccessControlServiceClient.prototype.addInboundAllowRuleByInstanceIPMatch =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/infra.AccessControlService/AddInboundAllowRuleByInstanceIPMatch',
      request,
      metadata || {},
      methodDescriptor_AccessControlService_AddInboundAllowRuleByInstanceIPMatch,
      callback);
};


/**
 * @param {!proto.infra.AddInboundAllowRuleByInstanceIPMatchRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.infra.AddInboundAllowRuleByInstanceIPMatchResponse>}
 *     Promise that resolves to the response
 */
proto.infra.AccessControlServicePromiseClient.prototype.addInboundAllowRuleByInstanceIPMatch =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/infra.AccessControlService/AddInboundAllowRuleByInstanceIPMatch',
      request,
      metadata || {},
      methodDescriptor_AccessControlService_AddInboundAllowRuleByInstanceIPMatch);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.infra.AddInboundAllowRuleForLoadBalancerByDNSRequest,
 *   !proto.infra.AddInboundAllowRuleForLoadBalancerByDNSResponse>}
 */
const methodDescriptor_AccessControlService_AddInboundAllowRuleForLoadBalancerByDNS = new grpc.web.MethodDescriptor(
  '/infra.AccessControlService/AddInboundAllowRuleForLoadBalancerByDNS',
  grpc.web.MethodType.UNARY,
  proto.infra.AddInboundAllowRuleForLoadBalancerByDNSRequest,
  proto.infra.AddInboundAllowRuleForLoadBalancerByDNSResponse,
  /**
   * @param {!proto.infra.AddInboundAllowRuleForLoadBalancerByDNSRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.infra.AddInboundAllowRuleForLoadBalancerByDNSResponse.deserializeBinary
);


/**
 * @param {!proto.infra.AddInboundAllowRuleForLoadBalancerByDNSRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.infra.AddInboundAllowRuleForLoadBalancerByDNSResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.infra.AddInboundAllowRuleForLoadBalancerByDNSResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.infra.AccessControlServiceClient.prototype.addInboundAllowRuleForLoadBalancerByDNS =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/infra.AccessControlService/AddInboundAllowRuleForLoadBalancerByDNS',
      request,
      metadata || {},
      methodDescriptor_AccessControlService_AddInboundAllowRuleForLoadBalancerByDNS,
      callback);
};


/**
 * @param {!proto.infra.AddInboundAllowRuleForLoadBalancerByDNSRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.infra.AddInboundAllowRuleForLoadBalancerByDNSResponse>}
 *     Promise that resolves to the response
 */
proto.infra.AccessControlServicePromiseClient.prototype.addInboundAllowRuleForLoadBalancerByDNS =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/infra.AccessControlService/AddInboundAllowRuleForLoadBalancerByDNS',
      request,
      metadata || {},
      methodDescriptor_AccessControlService_AddInboundAllowRuleForLoadBalancerByDNS);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.infra.RemoveInboundAllowRuleFromVPCByNameRequest,
 *   !proto.infra.RemoveInboundAllowRuleFromVPCByNameResponse>}
 */
const methodDescriptor_AccessControlService_RemoveInboundAllowRuleFromVPCByName = new grpc.web.MethodDescriptor(
  '/infra.AccessControlService/RemoveInboundAllowRuleFromVPCByName',
  grpc.web.MethodType.UNARY,
  proto.infra.RemoveInboundAllowRuleFromVPCByNameRequest,
  proto.infra.RemoveInboundAllowRuleFromVPCByNameResponse,
  /**
   * @param {!proto.infra.RemoveInboundAllowRuleFromVPCByNameRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.infra.RemoveInboundAllowRuleFromVPCByNameResponse.deserializeBinary
);


/**
 * @param {!proto.infra.RemoveInboundAllowRuleFromVPCByNameRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.infra.RemoveInboundAllowRuleFromVPCByNameResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.infra.RemoveInboundAllowRuleFromVPCByNameResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.infra.AccessControlServiceClient.prototype.removeInboundAllowRuleFromVPCByName =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/infra.AccessControlService/RemoveInboundAllowRuleFromVPCByName',
      request,
      metadata || {},
      methodDescriptor_AccessControlService_RemoveInboundAllowRuleFromVPCByName,
      callback);
};


/**
 * @param {!proto.infra.RemoveInboundAllowRuleFromVPCByNameRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.infra.RemoveInboundAllowRuleFromVPCByNameResponse>}
 *     Promise that resolves to the response
 */
proto.infra.AccessControlServicePromiseClient.prototype.removeInboundAllowRuleFromVPCByName =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/infra.AccessControlService/RemoveInboundAllowRuleFromVPCByName',
      request,
      metadata || {},
      methodDescriptor_AccessControlService_RemoveInboundAllowRuleFromVPCByName);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.infra.RemoveInboundAllowRulesFromVPCByIdRequest,
 *   !proto.infra.RemoveInboundAllowRulesFromVPCByIdResponse>}
 */
const methodDescriptor_AccessControlService_RemoveInboundAllowRulesFromVPCById = new grpc.web.MethodDescriptor(
  '/infra.AccessControlService/RemoveInboundAllowRulesFromVPCById',
  grpc.web.MethodType.UNARY,
  proto.infra.RemoveInboundAllowRulesFromVPCByIdRequest,
  proto.infra.RemoveInboundAllowRulesFromVPCByIdResponse,
  /**
   * @param {!proto.infra.RemoveInboundAllowRulesFromVPCByIdRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.infra.RemoveInboundAllowRulesFromVPCByIdResponse.deserializeBinary
);


/**
 * @param {!proto.infra.RemoveInboundAllowRulesFromVPCByIdRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.infra.RemoveInboundAllowRulesFromVPCByIdResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.infra.RemoveInboundAllowRulesFromVPCByIdResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.infra.AccessControlServiceClient.prototype.removeInboundAllowRulesFromVPCById =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/infra.AccessControlService/RemoveInboundAllowRulesFromVPCById',
      request,
      metadata || {},
      methodDescriptor_AccessControlService_RemoveInboundAllowRulesFromVPCById,
      callback);
};


/**
 * @param {!proto.infra.RemoveInboundAllowRulesFromVPCByIdRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.infra.RemoveInboundAllowRulesFromVPCByIdResponse>}
 *     Promise that resolves to the response
 */
proto.infra.AccessControlServicePromiseClient.prototype.removeInboundAllowRulesFromVPCById =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/infra.AccessControlService/RemoveInboundAllowRulesFromVPCById',
      request,
      metadata || {},
      methodDescriptor_AccessControlService_RemoveInboundAllowRulesFromVPCById);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.infra.RemoveInboundAllowRuleRulesByTagsRequest,
 *   !proto.infra.RemoveInboundAllowRuleRulesByTagsResponse>}
 */
const methodDescriptor_AccessControlService_RemoveInboundAllowRuleRulesByTags = new grpc.web.MethodDescriptor(
  '/infra.AccessControlService/RemoveInboundAllowRuleRulesByTags',
  grpc.web.MethodType.UNARY,
  proto.infra.RemoveInboundAllowRuleRulesByTagsRequest,
  proto.infra.RemoveInboundAllowRuleRulesByTagsResponse,
  /**
   * @param {!proto.infra.RemoveInboundAllowRuleRulesByTagsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.infra.RemoveInboundAllowRuleRulesByTagsResponse.deserializeBinary
);


/**
 * @param {!proto.infra.RemoveInboundAllowRuleRulesByTagsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.infra.RemoveInboundAllowRuleRulesByTagsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.infra.RemoveInboundAllowRuleRulesByTagsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.infra.AccessControlServiceClient.prototype.removeInboundAllowRuleRulesByTags =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/infra.AccessControlService/RemoveInboundAllowRuleRulesByTags',
      request,
      metadata || {},
      methodDescriptor_AccessControlService_RemoveInboundAllowRuleRulesByTags,
      callback);
};


/**
 * @param {!proto.infra.RemoveInboundAllowRuleRulesByTagsRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.infra.RemoveInboundAllowRuleRulesByTagsResponse>}
 *     Promise that resolves to the response
 */
proto.infra.AccessControlServicePromiseClient.prototype.removeInboundAllowRuleRulesByTags =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/infra.AccessControlService/RemoveInboundAllowRuleRulesByTags',
      request,
      metadata || {},
      methodDescriptor_AccessControlService_RemoveInboundAllowRuleRulesByTags);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.infra.RefreshInboundAllowRuleRequest,
 *   !proto.infra.RefreshInboundAllowRuleResponse>}
 */
const methodDescriptor_AccessControlService_RefreshInboundAllowRule = new grpc.web.MethodDescriptor(
  '/infra.AccessControlService/RefreshInboundAllowRule',
  grpc.web.MethodType.UNARY,
  proto.infra.RefreshInboundAllowRuleRequest,
  proto.infra.RefreshInboundAllowRuleResponse,
  /**
   * @param {!proto.infra.RefreshInboundAllowRuleRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.infra.RefreshInboundAllowRuleResponse.deserializeBinary
);


/**
 * @param {!proto.infra.RefreshInboundAllowRuleRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.infra.RefreshInboundAllowRuleResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.infra.RefreshInboundAllowRuleResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.infra.AccessControlServiceClient.prototype.refreshInboundAllowRule =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/infra.AccessControlService/RefreshInboundAllowRule',
      request,
      metadata || {},
      methodDescriptor_AccessControlService_RefreshInboundAllowRule,
      callback);
};


/**
 * @param {!proto.infra.RefreshInboundAllowRuleRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.infra.RefreshInboundAllowRuleResponse>}
 *     Promise that resolves to the response
 */
proto.infra.AccessControlServicePromiseClient.prototype.refreshInboundAllowRule =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/infra.AccessControlService/RefreshInboundAllowRule',
      request,
      metadata || {},
      methodDescriptor_AccessControlService_RefreshInboundAllowRule);
};


module.exports = proto.infra;

