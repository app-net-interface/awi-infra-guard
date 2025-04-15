// source: cloud.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {missingRequire} reports error on implicit type usages.
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!
/* eslint-disable */
// @ts-nocheck

var jspb = require('google-protobuf');
var goog = jspb;
var global =
    (typeof globalThis !== 'undefined' && globalThis) ||
    (typeof window !== 'undefined' && window) ||
    (typeof global !== 'undefined' && global) ||
    (typeof self !== 'undefined' && self) ||
    (function () { return this; }).call(null) ||
    Function('return this')();

var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js');
goog.object.extend(proto, google_protobuf_timestamp_pb);
var google_protobuf_field_mask_pb = require('google-protobuf/google/protobuf/field_mask_pb.js');
goog.object.extend(proto, google_protobuf_field_mask_pb);
var types_pb = require('./types_pb.js');
goog.object.extend(proto, types_pb);
var types_k8s_pb = require('./types_k8s_pb.js');
goog.object.extend(proto, types_k8s_pb);
goog.exportSymbol('proto.infra.Counters', null, global);
goog.exportSymbol('proto.infra.GetCIDRsForLabelsRequest', null, global);
goog.exportSymbol('proto.infra.GetCIDRsForLabelsResponse', null, global);
goog.exportSymbol('proto.infra.GetIPsForLabelsRequest', null, global);
goog.exportSymbol('proto.infra.GetIPsForLabelsResponse', null, global);
goog.exportSymbol('proto.infra.GetInstancesForLabelsRequest', null, global);
goog.exportSymbol('proto.infra.GetInstancesForLabelsResponse', null, global);
goog.exportSymbol('proto.infra.GetSubnetRequest', null, global);
goog.exportSymbol('proto.infra.GetSubnetResponse', null, global);
goog.exportSymbol('proto.infra.GetVPCIDForCIDRRequest', null, global);
goog.exportSymbol('proto.infra.GetVPCIDForCIDRResponse', null, global);
goog.exportSymbol('proto.infra.GetVPCIDWithTagRequest', null, global);
goog.exportSymbol('proto.infra.GetVPCIDWithTagResponse', null, global);
goog.exportSymbol('proto.infra.GetVPCIndexRequest', null, global);
goog.exportSymbol('proto.infra.GetVPCIndexResponse', null, global);
goog.exportSymbol('proto.infra.GetVpcConnectivityGraphRequest', null, global);
goog.exportSymbol('proto.infra.GetVpcConnectivityGraphResponse', null, global);
goog.exportSymbol('proto.infra.ListACLsRequest', null, global);
goog.exportSymbol('proto.infra.ListACLsResponse', null, global);
goog.exportSymbol('proto.infra.ListAccountsRequest', null, global);
goog.exportSymbol('proto.infra.ListAccountsResponse', null, global);
goog.exportSymbol('proto.infra.ListCloudClustersRequest', null, global);
goog.exportSymbol('proto.infra.ListCloudClustersResponse', null, global);
goog.exportSymbol('proto.infra.ListInstancesRequest', null, global);
goog.exportSymbol('proto.infra.ListInstancesResponse', null, global);
goog.exportSymbol('proto.infra.ListInternetGatewaysRequest', null, global);
goog.exportSymbol('proto.infra.ListInternetGatewaysResponse', null, global);
goog.exportSymbol('proto.infra.ListKeyPairsRequest', null, global);
goog.exportSymbol('proto.infra.ListKeyPairsResponse', null, global);
goog.exportSymbol('proto.infra.ListLBsRequest', null, global);
goog.exportSymbol('proto.infra.ListLBsResponse', null, global);
goog.exportSymbol('proto.infra.ListNATGatewaysRequest', null, global);
goog.exportSymbol('proto.infra.ListNATGatewaysResponse', null, global);
goog.exportSymbol('proto.infra.ListNetworkInterfacesRequest', null, global);
goog.exportSymbol('proto.infra.ListNetworkInterfacesResponse', null, global);
goog.exportSymbol('proto.infra.ListPublicIPsRequest', null, global);
goog.exportSymbol('proto.infra.ListPublicIPsResponse', null, global);
goog.exportSymbol('proto.infra.ListRegionsRequest', null, global);
goog.exportSymbol('proto.infra.ListRegionsResponse', null, global);
goog.exportSymbol('proto.infra.ListRouteTablesRequest', null, global);
goog.exportSymbol('proto.infra.ListRouteTablesResponse', null, global);
goog.exportSymbol('proto.infra.ListRoutersRequest', null, global);
goog.exportSymbol('proto.infra.ListRoutersResponse', null, global);
goog.exportSymbol('proto.infra.ListSecurityGroupsRequest', null, global);
goog.exportSymbol('proto.infra.ListSecurityGroupsResponse', null, global);
goog.exportSymbol('proto.infra.ListSubnetsRequest', null, global);
goog.exportSymbol('proto.infra.ListSubnetsResponse', null, global);
goog.exportSymbol('proto.infra.ListVPCEndpointsRequest', null, global);
goog.exportSymbol('proto.infra.ListVPCEndpointsResponse', null, global);
goog.exportSymbol('proto.infra.ListVPCRequest', null, global);
goog.exportSymbol('proto.infra.ListVPCResponse', null, global);
goog.exportSymbol('proto.infra.ListVPNConcentratorsRequest', null, global);
goog.exportSymbol('proto.infra.ListVPNConcentratorsResponse', null, global);
goog.exportSymbol('proto.infra.ListVpcGraphEdgesRequest', null, global);
goog.exportSymbol('proto.infra.ListVpcGraphEdgesResponse', null, global);
goog.exportSymbol('proto.infra.ListVpcGraphNodesRequest', null, global);
goog.exportSymbol('proto.infra.ListVpcGraphNodesResponse', null, global);
goog.exportSymbol('proto.infra.SearchResourcesRequest', null, global);
goog.exportSymbol('proto.infra.SearchResourcesResponse', null, global);
goog.exportSymbol('proto.infra.StatusSummary', null, global);
goog.exportSymbol('proto.infra.SummaryRequest', null, global);
goog.exportSymbol('proto.infra.SummaryResponse', null, global);
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.infra.GetVPCIndexRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.infra.GetVPCIndexRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.infra.GetVPCIndexRequest.displayName = 'proto.infra.GetVPCIndexRequest';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.infra.GetVPCIndexResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.infra.GetVPCIndexResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.infra.GetVPCIndexResponse.displayName = 'proto.infra.GetVPCIndexResponse';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.infra.ListAccountsRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.infra.ListAccountsRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.infra.ListAccountsRequest.displayName = 'proto.infra.ListAccountsRequest';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.infra.ListAccountsResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.infra.ListAccountsResponse.repeatedFields_, null);
};
goog.inherits(proto.infra.ListAccountsResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.infra.ListAccountsResponse.displayName = 'proto.infra.ListAccountsResponse';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.infra.ListRegionsRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.infra.ListRegionsRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.infra.ListRegionsRequest.displayName = 'proto.infra.ListRegionsRequest';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.infra.ListRegionsResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.infra.ListRegionsResponse.repeatedFields_, null);
};
goog.inherits(proto.infra.ListRegionsResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.infra.ListRegionsResponse.displayName = 'proto.infra.ListRegionsResponse';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.infra.ListVPCRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.infra.ListVPCRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.infra.ListVPCRequest.displayName = 'proto.infra.ListVPCRequest';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.infra.ListVPCResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.infra.ListVPCResponse.repeatedFields_, null);
};
goog.inherits(proto.infra.ListVPCResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.infra.ListVPCResponse.displayName = 'proto.infra.ListVPCResponse';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.infra.ListInstancesRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.infra.ListInstancesRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.infra.ListInstancesRequest.displayName = 'proto.infra.ListInstancesRequest';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.infra.ListInstancesResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.infra.ListInstancesResponse.repeatedFields_, null);
};
goog.inherits(proto.infra.ListInstancesResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.infra.ListInstancesResponse.displayName = 'proto.infra.ListInstancesResponse';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.infra.ListACLsRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.infra.ListACLsRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.infra.ListACLsRequest.displayName = 'proto.infra.ListACLsRequest';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.infra.ListACLsResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.infra.ListACLsResponse.repeatedFields_, null);
};
goog.inherits(proto.infra.ListACLsResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.infra.ListACLsResponse.displayName = 'proto.infra.ListACLsResponse';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.infra.ListSecurityGroupsRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.infra.ListSecurityGroupsRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.infra.ListSecurityGroupsRequest.displayName = 'proto.infra.ListSecurityGroupsRequest';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.infra.ListSecurityGroupsResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.infra.ListSecurityGroupsResponse.repeatedFields_, null);
};
goog.inherits(proto.infra.ListSecurityGroupsResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.infra.ListSecurityGroupsResponse.displayName = 'proto.infra.ListSecurityGroupsResponse';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.infra.ListRouteTablesRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.infra.ListRouteTablesRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.infra.ListRouteTablesRequest.displayName = 'proto.infra.ListRouteTablesRequest';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.infra.ListRouteTablesResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.infra.ListRouteTablesResponse.repeatedFields_, null);
};
goog.inherits(proto.infra.ListRouteTablesResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.infra.ListRouteTablesResponse.displayName = 'proto.infra.ListRouteTablesResponse';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.infra.ListNATGatewaysRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.infra.ListNATGatewaysRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.infra.ListNATGatewaysRequest.displayName = 'proto.infra.ListNATGatewaysRequest';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.infra.ListNATGatewaysResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.infra.ListNATGatewaysResponse.repeatedFields_, null);
};
goog.inherits(proto.infra.ListNATGatewaysResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.infra.ListNATGatewaysResponse.displayName = 'proto.infra.ListNATGatewaysResponse';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.infra.ListRoutersRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.infra.ListRoutersRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.infra.ListRoutersRequest.displayName = 'proto.infra.ListRoutersRequest';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.infra.ListRoutersResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.infra.ListRoutersResponse.repeatedFields_, null);
};
goog.inherits(proto.infra.ListRoutersResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.infra.ListRoutersResponse.displayName = 'proto.infra.ListRoutersResponse';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.infra.ListInternetGatewaysRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.infra.ListInternetGatewaysRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.infra.ListInternetGatewaysRequest.displayName = 'proto.infra.ListInternetGatewaysRequest';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.infra.ListInternetGatewaysResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.infra.ListInternetGatewaysResponse.repeatedFields_, null);
};
goog.inherits(proto.infra.ListInternetGatewaysResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.infra.ListInternetGatewaysResponse.displayName = 'proto.infra.ListInternetGatewaysResponse';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.infra.ListVPCEndpointsRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.infra.ListVPCEndpointsRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.infra.ListVPCEndpointsRequest.displayName = 'proto.infra.ListVPCEndpointsRequest';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.infra.ListVPCEndpointsResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.infra.ListVPCEndpointsResponse.repeatedFields_, null);
};
goog.inherits(proto.infra.ListVPCEndpointsResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.infra.ListVPCEndpointsResponse.displayName = 'proto.infra.ListVPCEndpointsResponse';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.infra.ListPublicIPsRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.infra.ListPublicIPsRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.infra.ListPublicIPsRequest.displayName = 'proto.infra.ListPublicIPsRequest';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.infra.ListPublicIPsResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.infra.ListPublicIPsResponse.repeatedFields_, null);
};
goog.inherits(proto.infra.ListPublicIPsResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.infra.ListPublicIPsResponse.displayName = 'proto.infra.ListPublicIPsResponse';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.infra.ListLBsRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.infra.ListLBsRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.infra.ListLBsRequest.displayName = 'proto.infra.ListLBsRequest';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.infra.ListLBsResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.infra.ListLBsResponse.repeatedFields_, null);
};
goog.inherits(proto.infra.ListLBsResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.infra.ListLBsResponse.displayName = 'proto.infra.ListLBsResponse';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.infra.GetSubnetRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.infra.GetSubnetRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.infra.GetSubnetRequest.displayName = 'proto.infra.GetSubnetRequest';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.infra.GetSubnetResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.infra.GetSubnetResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.infra.GetSubnetResponse.displayName = 'proto.infra.GetSubnetResponse';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.infra.ListSubnetsRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.infra.ListSubnetsRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.infra.ListSubnetsRequest.displayName = 'proto.infra.ListSubnetsRequest';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.infra.ListSubnetsResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.infra.ListSubnetsResponse.repeatedFields_, null);
};
goog.inherits(proto.infra.ListSubnetsResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.infra.ListSubnetsResponse.displayName = 'proto.infra.ListSubnetsResponse';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.infra.ListNetworkInterfacesRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.infra.ListNetworkInterfacesRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.infra.ListNetworkInterfacesRequest.displayName = 'proto.infra.ListNetworkInterfacesRequest';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.infra.ListNetworkInterfacesResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.infra.ListNetworkInterfacesResponse.repeatedFields_, null);
};
goog.inherits(proto.infra.ListNetworkInterfacesResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.infra.ListNetworkInterfacesResponse.displayName = 'proto.infra.ListNetworkInterfacesResponse';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.infra.ListKeyPairsRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.infra.ListKeyPairsRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.infra.ListKeyPairsRequest.displayName = 'proto.infra.ListKeyPairsRequest';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.infra.ListKeyPairsResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.infra.ListKeyPairsResponse.repeatedFields_, null);
};
goog.inherits(proto.infra.ListKeyPairsResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.infra.ListKeyPairsResponse.displayName = 'proto.infra.ListKeyPairsResponse';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.infra.ListVPNConcentratorsRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.infra.ListVPNConcentratorsRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.infra.ListVPNConcentratorsRequest.displayName = 'proto.infra.ListVPNConcentratorsRequest';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.infra.ListVPNConcentratorsResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.infra.ListVPNConcentratorsResponse.repeatedFields_, null);
};
goog.inherits(proto.infra.ListVPNConcentratorsResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.infra.ListVPNConcentratorsResponse.displayName = 'proto.infra.ListVPNConcentratorsResponse';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.infra.GetVPCIDForCIDRRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.infra.GetVPCIDForCIDRRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.infra.GetVPCIDForCIDRRequest.displayName = 'proto.infra.GetVPCIDForCIDRRequest';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.infra.GetVPCIDForCIDRResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.infra.GetVPCIDForCIDRResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.infra.GetVPCIDForCIDRResponse.displayName = 'proto.infra.GetVPCIDForCIDRResponse';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.infra.GetCIDRsForLabelsRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.infra.GetCIDRsForLabelsRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.infra.GetCIDRsForLabelsRequest.displayName = 'proto.infra.GetCIDRsForLabelsRequest';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.infra.GetCIDRsForLabelsResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.infra.GetCIDRsForLabelsResponse.repeatedFields_, null);
};
goog.inherits(proto.infra.GetCIDRsForLabelsResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.infra.GetCIDRsForLabelsResponse.displayName = 'proto.infra.GetCIDRsForLabelsResponse';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.infra.GetIPsForLabelsRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.infra.GetIPsForLabelsRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.infra.GetIPsForLabelsRequest.displayName = 'proto.infra.GetIPsForLabelsRequest';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.infra.GetIPsForLabelsResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.infra.GetIPsForLabelsResponse.repeatedFields_, null);
};
goog.inherits(proto.infra.GetIPsForLabelsResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.infra.GetIPsForLabelsResponse.displayName = 'proto.infra.GetIPsForLabelsResponse';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.infra.GetInstancesForLabelsRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.infra.GetInstancesForLabelsRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.infra.GetInstancesForLabelsRequest.displayName = 'proto.infra.GetInstancesForLabelsRequest';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.infra.GetInstancesForLabelsResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.infra.GetInstancesForLabelsResponse.repeatedFields_, null);
};
goog.inherits(proto.infra.GetInstancesForLabelsResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.infra.GetInstancesForLabelsResponse.displayName = 'proto.infra.GetInstancesForLabelsResponse';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.infra.GetVPCIDWithTagRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.infra.GetVPCIDWithTagRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.infra.GetVPCIDWithTagRequest.displayName = 'proto.infra.GetVPCIDWithTagRequest';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.infra.GetVPCIDWithTagResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.infra.GetVPCIDWithTagResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.infra.GetVPCIDWithTagResponse.displayName = 'proto.infra.GetVPCIDWithTagResponse';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.infra.ListCloudClustersRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.infra.ListCloudClustersRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.infra.ListCloudClustersRequest.displayName = 'proto.infra.ListCloudClustersRequest';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.infra.ListCloudClustersResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.infra.ListCloudClustersResponse.repeatedFields_, null);
};
goog.inherits(proto.infra.ListCloudClustersResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.infra.ListCloudClustersResponse.displayName = 'proto.infra.ListCloudClustersResponse';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.infra.SummaryRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.infra.SummaryRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.infra.SummaryRequest.displayName = 'proto.infra.SummaryRequest';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.infra.Counters = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.infra.Counters, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.infra.Counters.displayName = 'proto.infra.Counters';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.infra.StatusSummary = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.infra.StatusSummary, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.infra.StatusSummary.displayName = 'proto.infra.StatusSummary';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.infra.SummaryResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.infra.SummaryResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.infra.SummaryResponse.displayName = 'proto.infra.SummaryResponse';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.infra.SearchResourcesRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.infra.SearchResourcesRequest.repeatedFields_, null);
};
goog.inherits(proto.infra.SearchResourcesRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.infra.SearchResourcesRequest.displayName = 'proto.infra.SearchResourcesRequest';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.infra.SearchResourcesResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.infra.SearchResourcesResponse.repeatedFields_, null);
};
goog.inherits(proto.infra.SearchResourcesResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.infra.SearchResourcesResponse.displayName = 'proto.infra.SearchResourcesResponse';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.infra.GetVpcConnectivityGraphRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.infra.GetVpcConnectivityGraphRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.infra.GetVpcConnectivityGraphRequest.displayName = 'proto.infra.GetVpcConnectivityGraphRequest';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.infra.GetVpcConnectivityGraphResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.infra.GetVpcConnectivityGraphResponse.repeatedFields_, null);
};
goog.inherits(proto.infra.GetVpcConnectivityGraphResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.infra.GetVpcConnectivityGraphResponse.displayName = 'proto.infra.GetVpcConnectivityGraphResponse';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.infra.ListVpcGraphNodesRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.infra.ListVpcGraphNodesRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.infra.ListVpcGraphNodesRequest.displayName = 'proto.infra.ListVpcGraphNodesRequest';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.infra.ListVpcGraphNodesResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.infra.ListVpcGraphNodesResponse.repeatedFields_, null);
};
goog.inherits(proto.infra.ListVpcGraphNodesResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.infra.ListVpcGraphNodesResponse.displayName = 'proto.infra.ListVpcGraphNodesResponse';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.infra.ListVpcGraphEdgesRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.infra.ListVpcGraphEdgesRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.infra.ListVpcGraphEdgesRequest.displayName = 'proto.infra.ListVpcGraphEdgesRequest';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.infra.ListVpcGraphEdgesResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.infra.ListVpcGraphEdgesResponse.repeatedFields_, null);
};
goog.inherits(proto.infra.ListVpcGraphEdgesResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.infra.ListVpcGraphEdgesResponse.displayName = 'proto.infra.ListVpcGraphEdgesResponse';
}



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.infra.GetVPCIndexRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.infra.GetVPCIndexRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.infra.GetVPCIndexRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.GetVPCIndexRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    provider: jspb.Message.getFieldWithDefault(msg, 1, ""),
    accountId: jspb.Message.getFieldWithDefault(msg, 2, ""),
    region: jspb.Message.getFieldWithDefault(msg, 3, ""),
    vpcId: jspb.Message.getFieldWithDefault(msg, 4, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.infra.GetVPCIndexRequest}
 */
proto.infra.GetVPCIndexRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.infra.GetVPCIndexRequest;
  return proto.infra.GetVPCIndexRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.infra.GetVPCIndexRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.infra.GetVPCIndexRequest}
 */
proto.infra.GetVPCIndexRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setProvider(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setAccountId(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setRegion(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setVpcId(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.infra.GetVPCIndexRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.infra.GetVPCIndexRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.infra.GetVPCIndexRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.GetVPCIndexRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getProvider();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getAccountId();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getRegion();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getVpcId();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
};


/**
 * optional string provider = 1;
 * @return {string}
 */
proto.infra.GetVPCIndexRequest.prototype.getProvider = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.GetVPCIndexRequest} returns this
 */
proto.infra.GetVPCIndexRequest.prototype.setProvider = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string account_id = 2;
 * @return {string}
 */
proto.infra.GetVPCIndexRequest.prototype.getAccountId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.GetVPCIndexRequest} returns this
 */
proto.infra.GetVPCIndexRequest.prototype.setAccountId = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string region = 3;
 * @return {string}
 */
proto.infra.GetVPCIndexRequest.prototype.getRegion = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.GetVPCIndexRequest} returns this
 */
proto.infra.GetVPCIndexRequest.prototype.setRegion = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string vpc_id = 4;
 * @return {string}
 */
proto.infra.GetVPCIndexRequest.prototype.getVpcId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.GetVPCIndexRequest} returns this
 */
proto.infra.GetVPCIndexRequest.prototype.setVpcId = function(value) {
  return jspb.Message.setProto3StringField(this, 4, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.infra.GetVPCIndexResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.infra.GetVPCIndexResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.infra.GetVPCIndexResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.GetVPCIndexResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    vpcIndex: (f = msg.getVpcIndex()) && types_pb.VPCIndex.toObject(includeInstance, f),
    lastSyncTime: jspb.Message.getFieldWithDefault(msg, 17, ""),
    err: (f = msg.getErr()) && types_pb.Error.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.infra.GetVPCIndexResponse}
 */
proto.infra.GetVPCIndexResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.infra.GetVPCIndexResponse;
  return proto.infra.GetVPCIndexResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.infra.GetVPCIndexResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.infra.GetVPCIndexResponse}
 */
proto.infra.GetVPCIndexResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new types_pb.VPCIndex;
      reader.readMessage(value,types_pb.VPCIndex.deserializeBinaryFromReader);
      msg.setVpcIndex(value);
      break;
    case 17:
      var value = /** @type {string} */ (reader.readString());
      msg.setLastSyncTime(value);
      break;
    case 18:
      var value = new types_pb.Error;
      reader.readMessage(value,types_pb.Error.deserializeBinaryFromReader);
      msg.setErr(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.infra.GetVPCIndexResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.infra.GetVPCIndexResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.infra.GetVPCIndexResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.GetVPCIndexResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getVpcIndex();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      types_pb.VPCIndex.serializeBinaryToWriter
    );
  }
  f = message.getLastSyncTime();
  if (f.length > 0) {
    writer.writeString(
      17,
      f
    );
  }
  f = message.getErr();
  if (f != null) {
    writer.writeMessage(
      18,
      f,
      types_pb.Error.serializeBinaryToWriter
    );
  }
};


/**
 * optional VPCIndex vpc_index = 1;
 * @return {?proto.infra.VPCIndex}
 */
proto.infra.GetVPCIndexResponse.prototype.getVpcIndex = function() {
  return /** @type{?proto.infra.VPCIndex} */ (
    jspb.Message.getWrapperField(this, types_pb.VPCIndex, 1));
};


/**
 * @param {?proto.infra.VPCIndex|undefined} value
 * @return {!proto.infra.GetVPCIndexResponse} returns this
*/
proto.infra.GetVPCIndexResponse.prototype.setVpcIndex = function(value) {
  return jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.infra.GetVPCIndexResponse} returns this
 */
proto.infra.GetVPCIndexResponse.prototype.clearVpcIndex = function() {
  return this.setVpcIndex(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.infra.GetVPCIndexResponse.prototype.hasVpcIndex = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional string last_sync_time = 17;
 * @return {string}
 */
proto.infra.GetVPCIndexResponse.prototype.getLastSyncTime = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 17, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.GetVPCIndexResponse} returns this
 */
proto.infra.GetVPCIndexResponse.prototype.setLastSyncTime = function(value) {
  return jspb.Message.setProto3StringField(this, 17, value);
};


/**
 * optional Error err = 18;
 * @return {?proto.infra.Error}
 */
proto.infra.GetVPCIndexResponse.prototype.getErr = function() {
  return /** @type{?proto.infra.Error} */ (
    jspb.Message.getWrapperField(this, types_pb.Error, 18));
};


/**
 * @param {?proto.infra.Error|undefined} value
 * @return {!proto.infra.GetVPCIndexResponse} returns this
*/
proto.infra.GetVPCIndexResponse.prototype.setErr = function(value) {
  return jspb.Message.setWrapperField(this, 18, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.infra.GetVPCIndexResponse} returns this
 */
proto.infra.GetVPCIndexResponse.prototype.clearErr = function() {
  return this.setErr(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.infra.GetVPCIndexResponse.prototype.hasErr = function() {
  return jspb.Message.getField(this, 18) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.infra.ListAccountsRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.infra.ListAccountsRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.infra.ListAccountsRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListAccountsRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    provider: jspb.Message.getFieldWithDefault(msg, 1, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.infra.ListAccountsRequest}
 */
proto.infra.ListAccountsRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.infra.ListAccountsRequest;
  return proto.infra.ListAccountsRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.infra.ListAccountsRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.infra.ListAccountsRequest}
 */
proto.infra.ListAccountsRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setProvider(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.infra.ListAccountsRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.infra.ListAccountsRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.infra.ListAccountsRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListAccountsRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getProvider();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
};


/**
 * optional string provider = 1;
 * @return {string}
 */
proto.infra.ListAccountsRequest.prototype.getProvider = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListAccountsRequest} returns this
 */
proto.infra.ListAccountsRequest.prototype.setProvider = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.infra.ListAccountsResponse.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.infra.ListAccountsResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.infra.ListAccountsResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.infra.ListAccountsResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListAccountsResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    accountsList: jspb.Message.toObjectList(msg.getAccountsList(),
    types_pb.Account.toObject, includeInstance),
    err: (f = msg.getErr()) && types_pb.Error.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.infra.ListAccountsResponse}
 */
proto.infra.ListAccountsResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.infra.ListAccountsResponse;
  return proto.infra.ListAccountsResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.infra.ListAccountsResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.infra.ListAccountsResponse}
 */
proto.infra.ListAccountsResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new types_pb.Account;
      reader.readMessage(value,types_pb.Account.deserializeBinaryFromReader);
      msg.addAccounts(value);
      break;
    case 2:
      var value = new types_pb.Error;
      reader.readMessage(value,types_pb.Error.deserializeBinaryFromReader);
      msg.setErr(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.infra.ListAccountsResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.infra.ListAccountsResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.infra.ListAccountsResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListAccountsResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getAccountsList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      types_pb.Account.serializeBinaryToWriter
    );
  }
  f = message.getErr();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      types_pb.Error.serializeBinaryToWriter
    );
  }
};


/**
 * repeated Account accounts = 1;
 * @return {!Array<!proto.infra.Account>}
 */
proto.infra.ListAccountsResponse.prototype.getAccountsList = function() {
  return /** @type{!Array<!proto.infra.Account>} */ (
    jspb.Message.getRepeatedWrapperField(this, types_pb.Account, 1));
};


/**
 * @param {!Array<!proto.infra.Account>} value
 * @return {!proto.infra.ListAccountsResponse} returns this
*/
proto.infra.ListAccountsResponse.prototype.setAccountsList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.infra.Account=} opt_value
 * @param {number=} opt_index
 * @return {!proto.infra.Account}
 */
proto.infra.ListAccountsResponse.prototype.addAccounts = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.infra.Account, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.infra.ListAccountsResponse} returns this
 */
proto.infra.ListAccountsResponse.prototype.clearAccountsList = function() {
  return this.setAccountsList([]);
};


/**
 * optional Error err = 2;
 * @return {?proto.infra.Error}
 */
proto.infra.ListAccountsResponse.prototype.getErr = function() {
  return /** @type{?proto.infra.Error} */ (
    jspb.Message.getWrapperField(this, types_pb.Error, 2));
};


/**
 * @param {?proto.infra.Error|undefined} value
 * @return {!proto.infra.ListAccountsResponse} returns this
*/
proto.infra.ListAccountsResponse.prototype.setErr = function(value) {
  return jspb.Message.setWrapperField(this, 2, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.infra.ListAccountsResponse} returns this
 */
proto.infra.ListAccountsResponse.prototype.clearErr = function() {
  return this.setErr(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.infra.ListAccountsResponse.prototype.hasErr = function() {
  return jspb.Message.getField(this, 2) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.infra.ListRegionsRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.infra.ListRegionsRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.infra.ListRegionsRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListRegionsRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    provider: jspb.Message.getFieldWithDefault(msg, 1, ""),
    accountId: jspb.Message.getFieldWithDefault(msg, 2, ""),
    creds: (f = msg.getCreds()) && types_pb.Credentials.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.infra.ListRegionsRequest}
 */
proto.infra.ListRegionsRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.infra.ListRegionsRequest;
  return proto.infra.ListRegionsRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.infra.ListRegionsRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.infra.ListRegionsRequest}
 */
proto.infra.ListRegionsRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setProvider(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setAccountId(value);
      break;
    case 3:
      var value = new types_pb.Credentials;
      reader.readMessage(value,types_pb.Credentials.deserializeBinaryFromReader);
      msg.setCreds(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.infra.ListRegionsRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.infra.ListRegionsRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.infra.ListRegionsRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListRegionsRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getProvider();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getAccountId();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getCreds();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      types_pb.Credentials.serializeBinaryToWriter
    );
  }
};


/**
 * optional string provider = 1;
 * @return {string}
 */
proto.infra.ListRegionsRequest.prototype.getProvider = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListRegionsRequest} returns this
 */
proto.infra.ListRegionsRequest.prototype.setProvider = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string account_id = 2;
 * @return {string}
 */
proto.infra.ListRegionsRequest.prototype.getAccountId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListRegionsRequest} returns this
 */
proto.infra.ListRegionsRequest.prototype.setAccountId = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional Credentials creds = 3;
 * @return {?proto.infra.Credentials}
 */
proto.infra.ListRegionsRequest.prototype.getCreds = function() {
  return /** @type{?proto.infra.Credentials} */ (
    jspb.Message.getWrapperField(this, types_pb.Credentials, 3));
};


/**
 * @param {?proto.infra.Credentials|undefined} value
 * @return {!proto.infra.ListRegionsRequest} returns this
*/
proto.infra.ListRegionsRequest.prototype.setCreds = function(value) {
  return jspb.Message.setWrapperField(this, 3, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.infra.ListRegionsRequest} returns this
 */
proto.infra.ListRegionsRequest.prototype.clearCreds = function() {
  return this.setCreds(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.infra.ListRegionsRequest.prototype.hasCreds = function() {
  return jspb.Message.getField(this, 3) != null;
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.infra.ListRegionsResponse.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.infra.ListRegionsResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.infra.ListRegionsResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.infra.ListRegionsResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListRegionsResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    regionsList: jspb.Message.toObjectList(msg.getRegionsList(),
    types_pb.Region.toObject, includeInstance),
    err: (f = msg.getErr()) && types_pb.Error.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.infra.ListRegionsResponse}
 */
proto.infra.ListRegionsResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.infra.ListRegionsResponse;
  return proto.infra.ListRegionsResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.infra.ListRegionsResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.infra.ListRegionsResponse}
 */
proto.infra.ListRegionsResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new types_pb.Region;
      reader.readMessage(value,types_pb.Region.deserializeBinaryFromReader);
      msg.addRegions(value);
      break;
    case 2:
      var value = new types_pb.Error;
      reader.readMessage(value,types_pb.Error.deserializeBinaryFromReader);
      msg.setErr(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.infra.ListRegionsResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.infra.ListRegionsResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.infra.ListRegionsResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListRegionsResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getRegionsList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      types_pb.Region.serializeBinaryToWriter
    );
  }
  f = message.getErr();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      types_pb.Error.serializeBinaryToWriter
    );
  }
};


/**
 * repeated Region regions = 1;
 * @return {!Array<!proto.infra.Region>}
 */
proto.infra.ListRegionsResponse.prototype.getRegionsList = function() {
  return /** @type{!Array<!proto.infra.Region>} */ (
    jspb.Message.getRepeatedWrapperField(this, types_pb.Region, 1));
};


/**
 * @param {!Array<!proto.infra.Region>} value
 * @return {!proto.infra.ListRegionsResponse} returns this
*/
proto.infra.ListRegionsResponse.prototype.setRegionsList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.infra.Region=} opt_value
 * @param {number=} opt_index
 * @return {!proto.infra.Region}
 */
proto.infra.ListRegionsResponse.prototype.addRegions = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.infra.Region, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.infra.ListRegionsResponse} returns this
 */
proto.infra.ListRegionsResponse.prototype.clearRegionsList = function() {
  return this.setRegionsList([]);
};


/**
 * optional Error err = 2;
 * @return {?proto.infra.Error}
 */
proto.infra.ListRegionsResponse.prototype.getErr = function() {
  return /** @type{?proto.infra.Error} */ (
    jspb.Message.getWrapperField(this, types_pb.Error, 2));
};


/**
 * @param {?proto.infra.Error|undefined} value
 * @return {!proto.infra.ListRegionsResponse} returns this
*/
proto.infra.ListRegionsResponse.prototype.setErr = function(value) {
  return jspb.Message.setWrapperField(this, 2, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.infra.ListRegionsResponse} returns this
 */
proto.infra.ListRegionsResponse.prototype.clearErr = function() {
  return this.setErr(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.infra.ListRegionsResponse.prototype.hasErr = function() {
  return jspb.Message.getField(this, 2) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.infra.ListVPCRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.infra.ListVPCRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.infra.ListVPCRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListVPCRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    provider: jspb.Message.getFieldWithDefault(msg, 1, ""),
    accountId: jspb.Message.getFieldWithDefault(msg, 2, ""),
    region: jspb.Message.getFieldWithDefault(msg, 3, ""),
    labelsMap: (f = msg.getLabelsMap()) ? f.toObject(includeInstance, undefined) : [],
    creds: (f = msg.getCreds()) && types_pb.Credentials.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.infra.ListVPCRequest}
 */
proto.infra.ListVPCRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.infra.ListVPCRequest;
  return proto.infra.ListVPCRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.infra.ListVPCRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.infra.ListVPCRequest}
 */
proto.infra.ListVPCRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setProvider(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setAccountId(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setRegion(value);
      break;
    case 4:
      var value = msg.getLabelsMap();
      reader.readMessage(value, function(message, reader) {
        jspb.Map.deserializeBinary(message, reader, jspb.BinaryReader.prototype.readString, jspb.BinaryReader.prototype.readString, null, "", "");
         });
      break;
    case 5:
      var value = new types_pb.Credentials;
      reader.readMessage(value,types_pb.Credentials.deserializeBinaryFromReader);
      msg.setCreds(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.infra.ListVPCRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.infra.ListVPCRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.infra.ListVPCRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListVPCRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getProvider();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getAccountId();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getRegion();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getLabelsMap(true);
  if (f && f.getLength() > 0) {
    f.serializeBinary(4, writer, jspb.BinaryWriter.prototype.writeString, jspb.BinaryWriter.prototype.writeString);
  }
  f = message.getCreds();
  if (f != null) {
    writer.writeMessage(
      5,
      f,
      types_pb.Credentials.serializeBinaryToWriter
    );
  }
};


/**
 * optional string provider = 1;
 * @return {string}
 */
proto.infra.ListVPCRequest.prototype.getProvider = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListVPCRequest} returns this
 */
proto.infra.ListVPCRequest.prototype.setProvider = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string account_id = 2;
 * @return {string}
 */
proto.infra.ListVPCRequest.prototype.getAccountId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListVPCRequest} returns this
 */
proto.infra.ListVPCRequest.prototype.setAccountId = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string region = 3;
 * @return {string}
 */
proto.infra.ListVPCRequest.prototype.getRegion = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListVPCRequest} returns this
 */
proto.infra.ListVPCRequest.prototype.setRegion = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * map<string, string> labels = 4;
 * @param {boolean=} opt_noLazyCreate Do not create the map if
 * empty, instead returning `undefined`
 * @return {!jspb.Map<string,string>}
 */
proto.infra.ListVPCRequest.prototype.getLabelsMap = function(opt_noLazyCreate) {
  return /** @type {!jspb.Map<string,string>} */ (
      jspb.Message.getMapField(this, 4, opt_noLazyCreate,
      null));
};


/**
 * Clears values from the map. The map will be non-null.
 * @return {!proto.infra.ListVPCRequest} returns this
 */
proto.infra.ListVPCRequest.prototype.clearLabelsMap = function() {
  this.getLabelsMap().clear();
  return this;
};


/**
 * optional Credentials creds = 5;
 * @return {?proto.infra.Credentials}
 */
proto.infra.ListVPCRequest.prototype.getCreds = function() {
  return /** @type{?proto.infra.Credentials} */ (
    jspb.Message.getWrapperField(this, types_pb.Credentials, 5));
};


/**
 * @param {?proto.infra.Credentials|undefined} value
 * @return {!proto.infra.ListVPCRequest} returns this
*/
proto.infra.ListVPCRequest.prototype.setCreds = function(value) {
  return jspb.Message.setWrapperField(this, 5, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.infra.ListVPCRequest} returns this
 */
proto.infra.ListVPCRequest.prototype.clearCreds = function() {
  return this.setCreds(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.infra.ListVPCRequest.prototype.hasCreds = function() {
  return jspb.Message.getField(this, 5) != null;
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.infra.ListVPCResponse.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.infra.ListVPCResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.infra.ListVPCResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.infra.ListVPCResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListVPCResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    vpcsList: jspb.Message.toObjectList(msg.getVpcsList(),
    types_pb.VPC.toObject, includeInstance),
    lastSyncTime: jspb.Message.getFieldWithDefault(msg, 2, ""),
    err: (f = msg.getErr()) && types_pb.Error.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.infra.ListVPCResponse}
 */
proto.infra.ListVPCResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.infra.ListVPCResponse;
  return proto.infra.ListVPCResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.infra.ListVPCResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.infra.ListVPCResponse}
 */
proto.infra.ListVPCResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new types_pb.VPC;
      reader.readMessage(value,types_pb.VPC.deserializeBinaryFromReader);
      msg.addVpcs(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setLastSyncTime(value);
      break;
    case 3:
      var value = new types_pb.Error;
      reader.readMessage(value,types_pb.Error.deserializeBinaryFromReader);
      msg.setErr(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.infra.ListVPCResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.infra.ListVPCResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.infra.ListVPCResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListVPCResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getVpcsList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      types_pb.VPC.serializeBinaryToWriter
    );
  }
  f = message.getLastSyncTime();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getErr();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      types_pb.Error.serializeBinaryToWriter
    );
  }
};


/**
 * repeated VPC vpcs = 1;
 * @return {!Array<!proto.infra.VPC>}
 */
proto.infra.ListVPCResponse.prototype.getVpcsList = function() {
  return /** @type{!Array<!proto.infra.VPC>} */ (
    jspb.Message.getRepeatedWrapperField(this, types_pb.VPC, 1));
};


/**
 * @param {!Array<!proto.infra.VPC>} value
 * @return {!proto.infra.ListVPCResponse} returns this
*/
proto.infra.ListVPCResponse.prototype.setVpcsList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.infra.VPC=} opt_value
 * @param {number=} opt_index
 * @return {!proto.infra.VPC}
 */
proto.infra.ListVPCResponse.prototype.addVpcs = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.infra.VPC, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.infra.ListVPCResponse} returns this
 */
proto.infra.ListVPCResponse.prototype.clearVpcsList = function() {
  return this.setVpcsList([]);
};


/**
 * optional string last_sync_time = 2;
 * @return {string}
 */
proto.infra.ListVPCResponse.prototype.getLastSyncTime = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListVPCResponse} returns this
 */
proto.infra.ListVPCResponse.prototype.setLastSyncTime = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional Error err = 3;
 * @return {?proto.infra.Error}
 */
proto.infra.ListVPCResponse.prototype.getErr = function() {
  return /** @type{?proto.infra.Error} */ (
    jspb.Message.getWrapperField(this, types_pb.Error, 3));
};


/**
 * @param {?proto.infra.Error|undefined} value
 * @return {!proto.infra.ListVPCResponse} returns this
*/
proto.infra.ListVPCResponse.prototype.setErr = function(value) {
  return jspb.Message.setWrapperField(this, 3, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.infra.ListVPCResponse} returns this
 */
proto.infra.ListVPCResponse.prototype.clearErr = function() {
  return this.setErr(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.infra.ListVPCResponse.prototype.hasErr = function() {
  return jspb.Message.getField(this, 3) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.infra.ListInstancesRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.infra.ListInstancesRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.infra.ListInstancesRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListInstancesRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    provider: jspb.Message.getFieldWithDefault(msg, 1, ""),
    accountId: jspb.Message.getFieldWithDefault(msg, 2, ""),
    region: jspb.Message.getFieldWithDefault(msg, 3, ""),
    vpcId: jspb.Message.getFieldWithDefault(msg, 4, ""),
    zone: jspb.Message.getFieldWithDefault(msg, 5, ""),
    labelsMap: (f = msg.getLabelsMap()) ? f.toObject(includeInstance, undefined) : [],
    creds: (f = msg.getCreds()) && types_pb.Credentials.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.infra.ListInstancesRequest}
 */
proto.infra.ListInstancesRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.infra.ListInstancesRequest;
  return proto.infra.ListInstancesRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.infra.ListInstancesRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.infra.ListInstancesRequest}
 */
proto.infra.ListInstancesRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setProvider(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setAccountId(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setRegion(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setVpcId(value);
      break;
    case 5:
      var value = /** @type {string} */ (reader.readString());
      msg.setZone(value);
      break;
    case 6:
      var value = msg.getLabelsMap();
      reader.readMessage(value, function(message, reader) {
        jspb.Map.deserializeBinary(message, reader, jspb.BinaryReader.prototype.readString, jspb.BinaryReader.prototype.readString, null, "", "");
         });
      break;
    case 7:
      var value = new types_pb.Credentials;
      reader.readMessage(value,types_pb.Credentials.deserializeBinaryFromReader);
      msg.setCreds(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.infra.ListInstancesRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.infra.ListInstancesRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.infra.ListInstancesRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListInstancesRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getProvider();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getAccountId();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getRegion();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getVpcId();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getZone();
  if (f.length > 0) {
    writer.writeString(
      5,
      f
    );
  }
  f = message.getLabelsMap(true);
  if (f && f.getLength() > 0) {
    f.serializeBinary(6, writer, jspb.BinaryWriter.prototype.writeString, jspb.BinaryWriter.prototype.writeString);
  }
  f = message.getCreds();
  if (f != null) {
    writer.writeMessage(
      7,
      f,
      types_pb.Credentials.serializeBinaryToWriter
    );
  }
};


/**
 * optional string provider = 1;
 * @return {string}
 */
proto.infra.ListInstancesRequest.prototype.getProvider = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListInstancesRequest} returns this
 */
proto.infra.ListInstancesRequest.prototype.setProvider = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string account_id = 2;
 * @return {string}
 */
proto.infra.ListInstancesRequest.prototype.getAccountId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListInstancesRequest} returns this
 */
proto.infra.ListInstancesRequest.prototype.setAccountId = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string region = 3;
 * @return {string}
 */
proto.infra.ListInstancesRequest.prototype.getRegion = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListInstancesRequest} returns this
 */
proto.infra.ListInstancesRequest.prototype.setRegion = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string vpc_id = 4;
 * @return {string}
 */
proto.infra.ListInstancesRequest.prototype.getVpcId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListInstancesRequest} returns this
 */
proto.infra.ListInstancesRequest.prototype.setVpcId = function(value) {
  return jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * optional string zone = 5;
 * @return {string}
 */
proto.infra.ListInstancesRequest.prototype.getZone = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 5, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListInstancesRequest} returns this
 */
proto.infra.ListInstancesRequest.prototype.setZone = function(value) {
  return jspb.Message.setProto3StringField(this, 5, value);
};


/**
 * map<string, string> labels = 6;
 * @param {boolean=} opt_noLazyCreate Do not create the map if
 * empty, instead returning `undefined`
 * @return {!jspb.Map<string,string>}
 */
proto.infra.ListInstancesRequest.prototype.getLabelsMap = function(opt_noLazyCreate) {
  return /** @type {!jspb.Map<string,string>} */ (
      jspb.Message.getMapField(this, 6, opt_noLazyCreate,
      null));
};


/**
 * Clears values from the map. The map will be non-null.
 * @return {!proto.infra.ListInstancesRequest} returns this
 */
proto.infra.ListInstancesRequest.prototype.clearLabelsMap = function() {
  this.getLabelsMap().clear();
  return this;
};


/**
 * optional Credentials creds = 7;
 * @return {?proto.infra.Credentials}
 */
proto.infra.ListInstancesRequest.prototype.getCreds = function() {
  return /** @type{?proto.infra.Credentials} */ (
    jspb.Message.getWrapperField(this, types_pb.Credentials, 7));
};


/**
 * @param {?proto.infra.Credentials|undefined} value
 * @return {!proto.infra.ListInstancesRequest} returns this
*/
proto.infra.ListInstancesRequest.prototype.setCreds = function(value) {
  return jspb.Message.setWrapperField(this, 7, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.infra.ListInstancesRequest} returns this
 */
proto.infra.ListInstancesRequest.prototype.clearCreds = function() {
  return this.setCreds(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.infra.ListInstancesRequest.prototype.hasCreds = function() {
  return jspb.Message.getField(this, 7) != null;
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.infra.ListInstancesResponse.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.infra.ListInstancesResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.infra.ListInstancesResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.infra.ListInstancesResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListInstancesResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    instancesList: jspb.Message.toObjectList(msg.getInstancesList(),
    types_pb.Instance.toObject, includeInstance),
    lastSyncTime: jspb.Message.getFieldWithDefault(msg, 2, ""),
    err: (f = msg.getErr()) && types_pb.Error.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.infra.ListInstancesResponse}
 */
proto.infra.ListInstancesResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.infra.ListInstancesResponse;
  return proto.infra.ListInstancesResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.infra.ListInstancesResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.infra.ListInstancesResponse}
 */
proto.infra.ListInstancesResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new types_pb.Instance;
      reader.readMessage(value,types_pb.Instance.deserializeBinaryFromReader);
      msg.addInstances(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setLastSyncTime(value);
      break;
    case 3:
      var value = new types_pb.Error;
      reader.readMessage(value,types_pb.Error.deserializeBinaryFromReader);
      msg.setErr(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.infra.ListInstancesResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.infra.ListInstancesResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.infra.ListInstancesResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListInstancesResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getInstancesList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      types_pb.Instance.serializeBinaryToWriter
    );
  }
  f = message.getLastSyncTime();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getErr();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      types_pb.Error.serializeBinaryToWriter
    );
  }
};


/**
 * repeated Instance instances = 1;
 * @return {!Array<!proto.infra.Instance>}
 */
proto.infra.ListInstancesResponse.prototype.getInstancesList = function() {
  return /** @type{!Array<!proto.infra.Instance>} */ (
    jspb.Message.getRepeatedWrapperField(this, types_pb.Instance, 1));
};


/**
 * @param {!Array<!proto.infra.Instance>} value
 * @return {!proto.infra.ListInstancesResponse} returns this
*/
proto.infra.ListInstancesResponse.prototype.setInstancesList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.infra.Instance=} opt_value
 * @param {number=} opt_index
 * @return {!proto.infra.Instance}
 */
proto.infra.ListInstancesResponse.prototype.addInstances = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.infra.Instance, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.infra.ListInstancesResponse} returns this
 */
proto.infra.ListInstancesResponse.prototype.clearInstancesList = function() {
  return this.setInstancesList([]);
};


/**
 * optional string last_sync_time = 2;
 * @return {string}
 */
proto.infra.ListInstancesResponse.prototype.getLastSyncTime = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListInstancesResponse} returns this
 */
proto.infra.ListInstancesResponse.prototype.setLastSyncTime = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional Error err = 3;
 * @return {?proto.infra.Error}
 */
proto.infra.ListInstancesResponse.prototype.getErr = function() {
  return /** @type{?proto.infra.Error} */ (
    jspb.Message.getWrapperField(this, types_pb.Error, 3));
};


/**
 * @param {?proto.infra.Error|undefined} value
 * @return {!proto.infra.ListInstancesResponse} returns this
*/
proto.infra.ListInstancesResponse.prototype.setErr = function(value) {
  return jspb.Message.setWrapperField(this, 3, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.infra.ListInstancesResponse} returns this
 */
proto.infra.ListInstancesResponse.prototype.clearErr = function() {
  return this.setErr(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.infra.ListInstancesResponse.prototype.hasErr = function() {
  return jspb.Message.getField(this, 3) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.infra.ListACLsRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.infra.ListACLsRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.infra.ListACLsRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListACLsRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    provider: jspb.Message.getFieldWithDefault(msg, 1, ""),
    accountId: jspb.Message.getFieldWithDefault(msg, 2, ""),
    region: jspb.Message.getFieldWithDefault(msg, 3, ""),
    vpcId: jspb.Message.getFieldWithDefault(msg, 4, ""),
    labelsMap: (f = msg.getLabelsMap()) ? f.toObject(includeInstance, undefined) : [],
    creds: (f = msg.getCreds()) && types_pb.Credentials.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.infra.ListACLsRequest}
 */
proto.infra.ListACLsRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.infra.ListACLsRequest;
  return proto.infra.ListACLsRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.infra.ListACLsRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.infra.ListACLsRequest}
 */
proto.infra.ListACLsRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setProvider(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setAccountId(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setRegion(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setVpcId(value);
      break;
    case 5:
      var value = msg.getLabelsMap();
      reader.readMessage(value, function(message, reader) {
        jspb.Map.deserializeBinary(message, reader, jspb.BinaryReader.prototype.readString, jspb.BinaryReader.prototype.readString, null, "", "");
         });
      break;
    case 6:
      var value = new types_pb.Credentials;
      reader.readMessage(value,types_pb.Credentials.deserializeBinaryFromReader);
      msg.setCreds(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.infra.ListACLsRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.infra.ListACLsRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.infra.ListACLsRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListACLsRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getProvider();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getAccountId();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getRegion();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getVpcId();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getLabelsMap(true);
  if (f && f.getLength() > 0) {
    f.serializeBinary(5, writer, jspb.BinaryWriter.prototype.writeString, jspb.BinaryWriter.prototype.writeString);
  }
  f = message.getCreds();
  if (f != null) {
    writer.writeMessage(
      6,
      f,
      types_pb.Credentials.serializeBinaryToWriter
    );
  }
};


/**
 * optional string provider = 1;
 * @return {string}
 */
proto.infra.ListACLsRequest.prototype.getProvider = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListACLsRequest} returns this
 */
proto.infra.ListACLsRequest.prototype.setProvider = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string account_id = 2;
 * @return {string}
 */
proto.infra.ListACLsRequest.prototype.getAccountId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListACLsRequest} returns this
 */
proto.infra.ListACLsRequest.prototype.setAccountId = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string region = 3;
 * @return {string}
 */
proto.infra.ListACLsRequest.prototype.getRegion = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListACLsRequest} returns this
 */
proto.infra.ListACLsRequest.prototype.setRegion = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string vpc_id = 4;
 * @return {string}
 */
proto.infra.ListACLsRequest.prototype.getVpcId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListACLsRequest} returns this
 */
proto.infra.ListACLsRequest.prototype.setVpcId = function(value) {
  return jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * map<string, string> labels = 5;
 * @param {boolean=} opt_noLazyCreate Do not create the map if
 * empty, instead returning `undefined`
 * @return {!jspb.Map<string,string>}
 */
proto.infra.ListACLsRequest.prototype.getLabelsMap = function(opt_noLazyCreate) {
  return /** @type {!jspb.Map<string,string>} */ (
      jspb.Message.getMapField(this, 5, opt_noLazyCreate,
      null));
};


/**
 * Clears values from the map. The map will be non-null.
 * @return {!proto.infra.ListACLsRequest} returns this
 */
proto.infra.ListACLsRequest.prototype.clearLabelsMap = function() {
  this.getLabelsMap().clear();
  return this;
};


/**
 * optional Credentials creds = 6;
 * @return {?proto.infra.Credentials}
 */
proto.infra.ListACLsRequest.prototype.getCreds = function() {
  return /** @type{?proto.infra.Credentials} */ (
    jspb.Message.getWrapperField(this, types_pb.Credentials, 6));
};


/**
 * @param {?proto.infra.Credentials|undefined} value
 * @return {!proto.infra.ListACLsRequest} returns this
*/
proto.infra.ListACLsRequest.prototype.setCreds = function(value) {
  return jspb.Message.setWrapperField(this, 6, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.infra.ListACLsRequest} returns this
 */
proto.infra.ListACLsRequest.prototype.clearCreds = function() {
  return this.setCreds(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.infra.ListACLsRequest.prototype.hasCreds = function() {
  return jspb.Message.getField(this, 6) != null;
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.infra.ListACLsResponse.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.infra.ListACLsResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.infra.ListACLsResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.infra.ListACLsResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListACLsResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    aclsList: jspb.Message.toObjectList(msg.getAclsList(),
    types_pb.ACL.toObject, includeInstance),
    lastSyncTime: jspb.Message.getFieldWithDefault(msg, 2, ""),
    err: (f = msg.getErr()) && types_pb.Error.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.infra.ListACLsResponse}
 */
proto.infra.ListACLsResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.infra.ListACLsResponse;
  return proto.infra.ListACLsResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.infra.ListACLsResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.infra.ListACLsResponse}
 */
proto.infra.ListACLsResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new types_pb.ACL;
      reader.readMessage(value,types_pb.ACL.deserializeBinaryFromReader);
      msg.addAcls(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setLastSyncTime(value);
      break;
    case 3:
      var value = new types_pb.Error;
      reader.readMessage(value,types_pb.Error.deserializeBinaryFromReader);
      msg.setErr(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.infra.ListACLsResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.infra.ListACLsResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.infra.ListACLsResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListACLsResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getAclsList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      types_pb.ACL.serializeBinaryToWriter
    );
  }
  f = message.getLastSyncTime();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getErr();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      types_pb.Error.serializeBinaryToWriter
    );
  }
};


/**
 * repeated ACL acls = 1;
 * @return {!Array<!proto.infra.ACL>}
 */
proto.infra.ListACLsResponse.prototype.getAclsList = function() {
  return /** @type{!Array<!proto.infra.ACL>} */ (
    jspb.Message.getRepeatedWrapperField(this, types_pb.ACL, 1));
};


/**
 * @param {!Array<!proto.infra.ACL>} value
 * @return {!proto.infra.ListACLsResponse} returns this
*/
proto.infra.ListACLsResponse.prototype.setAclsList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.infra.ACL=} opt_value
 * @param {number=} opt_index
 * @return {!proto.infra.ACL}
 */
proto.infra.ListACLsResponse.prototype.addAcls = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.infra.ACL, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.infra.ListACLsResponse} returns this
 */
proto.infra.ListACLsResponse.prototype.clearAclsList = function() {
  return this.setAclsList([]);
};


/**
 * optional string last_sync_time = 2;
 * @return {string}
 */
proto.infra.ListACLsResponse.prototype.getLastSyncTime = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListACLsResponse} returns this
 */
proto.infra.ListACLsResponse.prototype.setLastSyncTime = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional Error err = 3;
 * @return {?proto.infra.Error}
 */
proto.infra.ListACLsResponse.prototype.getErr = function() {
  return /** @type{?proto.infra.Error} */ (
    jspb.Message.getWrapperField(this, types_pb.Error, 3));
};


/**
 * @param {?proto.infra.Error|undefined} value
 * @return {!proto.infra.ListACLsResponse} returns this
*/
proto.infra.ListACLsResponse.prototype.setErr = function(value) {
  return jspb.Message.setWrapperField(this, 3, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.infra.ListACLsResponse} returns this
 */
proto.infra.ListACLsResponse.prototype.clearErr = function() {
  return this.setErr(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.infra.ListACLsResponse.prototype.hasErr = function() {
  return jspb.Message.getField(this, 3) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.infra.ListSecurityGroupsRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.infra.ListSecurityGroupsRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.infra.ListSecurityGroupsRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListSecurityGroupsRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    provider: jspb.Message.getFieldWithDefault(msg, 1, ""),
    accountId: jspb.Message.getFieldWithDefault(msg, 2, ""),
    region: jspb.Message.getFieldWithDefault(msg, 3, ""),
    vpcId: jspb.Message.getFieldWithDefault(msg, 4, ""),
    labelsMap: (f = msg.getLabelsMap()) ? f.toObject(includeInstance, undefined) : [],
    creds: (f = msg.getCreds()) && types_pb.Credentials.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.infra.ListSecurityGroupsRequest}
 */
proto.infra.ListSecurityGroupsRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.infra.ListSecurityGroupsRequest;
  return proto.infra.ListSecurityGroupsRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.infra.ListSecurityGroupsRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.infra.ListSecurityGroupsRequest}
 */
proto.infra.ListSecurityGroupsRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setProvider(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setAccountId(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setRegion(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setVpcId(value);
      break;
    case 5:
      var value = msg.getLabelsMap();
      reader.readMessage(value, function(message, reader) {
        jspb.Map.deserializeBinary(message, reader, jspb.BinaryReader.prototype.readString, jspb.BinaryReader.prototype.readString, null, "", "");
         });
      break;
    case 6:
      var value = new types_pb.Credentials;
      reader.readMessage(value,types_pb.Credentials.deserializeBinaryFromReader);
      msg.setCreds(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.infra.ListSecurityGroupsRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.infra.ListSecurityGroupsRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.infra.ListSecurityGroupsRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListSecurityGroupsRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getProvider();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getAccountId();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getRegion();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getVpcId();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getLabelsMap(true);
  if (f && f.getLength() > 0) {
    f.serializeBinary(5, writer, jspb.BinaryWriter.prototype.writeString, jspb.BinaryWriter.prototype.writeString);
  }
  f = message.getCreds();
  if (f != null) {
    writer.writeMessage(
      6,
      f,
      types_pb.Credentials.serializeBinaryToWriter
    );
  }
};


/**
 * optional string provider = 1;
 * @return {string}
 */
proto.infra.ListSecurityGroupsRequest.prototype.getProvider = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListSecurityGroupsRequest} returns this
 */
proto.infra.ListSecurityGroupsRequest.prototype.setProvider = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string account_id = 2;
 * @return {string}
 */
proto.infra.ListSecurityGroupsRequest.prototype.getAccountId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListSecurityGroupsRequest} returns this
 */
proto.infra.ListSecurityGroupsRequest.prototype.setAccountId = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string region = 3;
 * @return {string}
 */
proto.infra.ListSecurityGroupsRequest.prototype.getRegion = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListSecurityGroupsRequest} returns this
 */
proto.infra.ListSecurityGroupsRequest.prototype.setRegion = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string vpc_id = 4;
 * @return {string}
 */
proto.infra.ListSecurityGroupsRequest.prototype.getVpcId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListSecurityGroupsRequest} returns this
 */
proto.infra.ListSecurityGroupsRequest.prototype.setVpcId = function(value) {
  return jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * map<string, string> labels = 5;
 * @param {boolean=} opt_noLazyCreate Do not create the map if
 * empty, instead returning `undefined`
 * @return {!jspb.Map<string,string>}
 */
proto.infra.ListSecurityGroupsRequest.prototype.getLabelsMap = function(opt_noLazyCreate) {
  return /** @type {!jspb.Map<string,string>} */ (
      jspb.Message.getMapField(this, 5, opt_noLazyCreate,
      null));
};


/**
 * Clears values from the map. The map will be non-null.
 * @return {!proto.infra.ListSecurityGroupsRequest} returns this
 */
proto.infra.ListSecurityGroupsRequest.prototype.clearLabelsMap = function() {
  this.getLabelsMap().clear();
  return this;
};


/**
 * optional Credentials creds = 6;
 * @return {?proto.infra.Credentials}
 */
proto.infra.ListSecurityGroupsRequest.prototype.getCreds = function() {
  return /** @type{?proto.infra.Credentials} */ (
    jspb.Message.getWrapperField(this, types_pb.Credentials, 6));
};


/**
 * @param {?proto.infra.Credentials|undefined} value
 * @return {!proto.infra.ListSecurityGroupsRequest} returns this
*/
proto.infra.ListSecurityGroupsRequest.prototype.setCreds = function(value) {
  return jspb.Message.setWrapperField(this, 6, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.infra.ListSecurityGroupsRequest} returns this
 */
proto.infra.ListSecurityGroupsRequest.prototype.clearCreds = function() {
  return this.setCreds(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.infra.ListSecurityGroupsRequest.prototype.hasCreds = function() {
  return jspb.Message.getField(this, 6) != null;
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.infra.ListSecurityGroupsResponse.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.infra.ListSecurityGroupsResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.infra.ListSecurityGroupsResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.infra.ListSecurityGroupsResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListSecurityGroupsResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    securityGroupsList: jspb.Message.toObjectList(msg.getSecurityGroupsList(),
    types_pb.SecurityGroup.toObject, includeInstance),
    lastSyncTime: jspb.Message.getFieldWithDefault(msg, 2, ""),
    err: (f = msg.getErr()) && types_pb.Error.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.infra.ListSecurityGroupsResponse}
 */
proto.infra.ListSecurityGroupsResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.infra.ListSecurityGroupsResponse;
  return proto.infra.ListSecurityGroupsResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.infra.ListSecurityGroupsResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.infra.ListSecurityGroupsResponse}
 */
proto.infra.ListSecurityGroupsResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new types_pb.SecurityGroup;
      reader.readMessage(value,types_pb.SecurityGroup.deserializeBinaryFromReader);
      msg.addSecurityGroups(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setLastSyncTime(value);
      break;
    case 3:
      var value = new types_pb.Error;
      reader.readMessage(value,types_pb.Error.deserializeBinaryFromReader);
      msg.setErr(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.infra.ListSecurityGroupsResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.infra.ListSecurityGroupsResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.infra.ListSecurityGroupsResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListSecurityGroupsResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getSecurityGroupsList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      types_pb.SecurityGroup.serializeBinaryToWriter
    );
  }
  f = message.getLastSyncTime();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getErr();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      types_pb.Error.serializeBinaryToWriter
    );
  }
};


/**
 * repeated SecurityGroup security_groups = 1;
 * @return {!Array<!proto.infra.SecurityGroup>}
 */
proto.infra.ListSecurityGroupsResponse.prototype.getSecurityGroupsList = function() {
  return /** @type{!Array<!proto.infra.SecurityGroup>} */ (
    jspb.Message.getRepeatedWrapperField(this, types_pb.SecurityGroup, 1));
};


/**
 * @param {!Array<!proto.infra.SecurityGroup>} value
 * @return {!proto.infra.ListSecurityGroupsResponse} returns this
*/
proto.infra.ListSecurityGroupsResponse.prototype.setSecurityGroupsList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.infra.SecurityGroup=} opt_value
 * @param {number=} opt_index
 * @return {!proto.infra.SecurityGroup}
 */
proto.infra.ListSecurityGroupsResponse.prototype.addSecurityGroups = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.infra.SecurityGroup, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.infra.ListSecurityGroupsResponse} returns this
 */
proto.infra.ListSecurityGroupsResponse.prototype.clearSecurityGroupsList = function() {
  return this.setSecurityGroupsList([]);
};


/**
 * optional string last_sync_time = 2;
 * @return {string}
 */
proto.infra.ListSecurityGroupsResponse.prototype.getLastSyncTime = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListSecurityGroupsResponse} returns this
 */
proto.infra.ListSecurityGroupsResponse.prototype.setLastSyncTime = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional Error err = 3;
 * @return {?proto.infra.Error}
 */
proto.infra.ListSecurityGroupsResponse.prototype.getErr = function() {
  return /** @type{?proto.infra.Error} */ (
    jspb.Message.getWrapperField(this, types_pb.Error, 3));
};


/**
 * @param {?proto.infra.Error|undefined} value
 * @return {!proto.infra.ListSecurityGroupsResponse} returns this
*/
proto.infra.ListSecurityGroupsResponse.prototype.setErr = function(value) {
  return jspb.Message.setWrapperField(this, 3, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.infra.ListSecurityGroupsResponse} returns this
 */
proto.infra.ListSecurityGroupsResponse.prototype.clearErr = function() {
  return this.setErr(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.infra.ListSecurityGroupsResponse.prototype.hasErr = function() {
  return jspb.Message.getField(this, 3) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.infra.ListRouteTablesRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.infra.ListRouteTablesRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.infra.ListRouteTablesRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListRouteTablesRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    provider: jspb.Message.getFieldWithDefault(msg, 1, ""),
    accountId: jspb.Message.getFieldWithDefault(msg, 2, ""),
    region: jspb.Message.getFieldWithDefault(msg, 3, ""),
    vpcId: jspb.Message.getFieldWithDefault(msg, 4, ""),
    labelsMap: (f = msg.getLabelsMap()) ? f.toObject(includeInstance, undefined) : [],
    creds: (f = msg.getCreds()) && types_pb.Credentials.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.infra.ListRouteTablesRequest}
 */
proto.infra.ListRouteTablesRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.infra.ListRouteTablesRequest;
  return proto.infra.ListRouteTablesRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.infra.ListRouteTablesRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.infra.ListRouteTablesRequest}
 */
proto.infra.ListRouteTablesRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setProvider(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setAccountId(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setRegion(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setVpcId(value);
      break;
    case 5:
      var value = msg.getLabelsMap();
      reader.readMessage(value, function(message, reader) {
        jspb.Map.deserializeBinary(message, reader, jspb.BinaryReader.prototype.readString, jspb.BinaryReader.prototype.readString, null, "", "");
         });
      break;
    case 6:
      var value = new types_pb.Credentials;
      reader.readMessage(value,types_pb.Credentials.deserializeBinaryFromReader);
      msg.setCreds(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.infra.ListRouteTablesRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.infra.ListRouteTablesRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.infra.ListRouteTablesRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListRouteTablesRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getProvider();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getAccountId();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getRegion();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getVpcId();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getLabelsMap(true);
  if (f && f.getLength() > 0) {
    f.serializeBinary(5, writer, jspb.BinaryWriter.prototype.writeString, jspb.BinaryWriter.prototype.writeString);
  }
  f = message.getCreds();
  if (f != null) {
    writer.writeMessage(
      6,
      f,
      types_pb.Credentials.serializeBinaryToWriter
    );
  }
};


/**
 * optional string provider = 1;
 * @return {string}
 */
proto.infra.ListRouteTablesRequest.prototype.getProvider = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListRouteTablesRequest} returns this
 */
proto.infra.ListRouteTablesRequest.prototype.setProvider = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string account_id = 2;
 * @return {string}
 */
proto.infra.ListRouteTablesRequest.prototype.getAccountId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListRouteTablesRequest} returns this
 */
proto.infra.ListRouteTablesRequest.prototype.setAccountId = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string region = 3;
 * @return {string}
 */
proto.infra.ListRouteTablesRequest.prototype.getRegion = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListRouteTablesRequest} returns this
 */
proto.infra.ListRouteTablesRequest.prototype.setRegion = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string vpc_id = 4;
 * @return {string}
 */
proto.infra.ListRouteTablesRequest.prototype.getVpcId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListRouteTablesRequest} returns this
 */
proto.infra.ListRouteTablesRequest.prototype.setVpcId = function(value) {
  return jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * map<string, string> labels = 5;
 * @param {boolean=} opt_noLazyCreate Do not create the map if
 * empty, instead returning `undefined`
 * @return {!jspb.Map<string,string>}
 */
proto.infra.ListRouteTablesRequest.prototype.getLabelsMap = function(opt_noLazyCreate) {
  return /** @type {!jspb.Map<string,string>} */ (
      jspb.Message.getMapField(this, 5, opt_noLazyCreate,
      null));
};


/**
 * Clears values from the map. The map will be non-null.
 * @return {!proto.infra.ListRouteTablesRequest} returns this
 */
proto.infra.ListRouteTablesRequest.prototype.clearLabelsMap = function() {
  this.getLabelsMap().clear();
  return this;
};


/**
 * optional Credentials creds = 6;
 * @return {?proto.infra.Credentials}
 */
proto.infra.ListRouteTablesRequest.prototype.getCreds = function() {
  return /** @type{?proto.infra.Credentials} */ (
    jspb.Message.getWrapperField(this, types_pb.Credentials, 6));
};


/**
 * @param {?proto.infra.Credentials|undefined} value
 * @return {!proto.infra.ListRouteTablesRequest} returns this
*/
proto.infra.ListRouteTablesRequest.prototype.setCreds = function(value) {
  return jspb.Message.setWrapperField(this, 6, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.infra.ListRouteTablesRequest} returns this
 */
proto.infra.ListRouteTablesRequest.prototype.clearCreds = function() {
  return this.setCreds(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.infra.ListRouteTablesRequest.prototype.hasCreds = function() {
  return jspb.Message.getField(this, 6) != null;
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.infra.ListRouteTablesResponse.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.infra.ListRouteTablesResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.infra.ListRouteTablesResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.infra.ListRouteTablesResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListRouteTablesResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    routeTablesList: jspb.Message.toObjectList(msg.getRouteTablesList(),
    types_pb.RouteTable.toObject, includeInstance),
    lastSyncTime: jspb.Message.getFieldWithDefault(msg, 2, ""),
    err: (f = msg.getErr()) && types_pb.Error.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.infra.ListRouteTablesResponse}
 */
proto.infra.ListRouteTablesResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.infra.ListRouteTablesResponse;
  return proto.infra.ListRouteTablesResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.infra.ListRouteTablesResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.infra.ListRouteTablesResponse}
 */
proto.infra.ListRouteTablesResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new types_pb.RouteTable;
      reader.readMessage(value,types_pb.RouteTable.deserializeBinaryFromReader);
      msg.addRouteTables(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setLastSyncTime(value);
      break;
    case 3:
      var value = new types_pb.Error;
      reader.readMessage(value,types_pb.Error.deserializeBinaryFromReader);
      msg.setErr(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.infra.ListRouteTablesResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.infra.ListRouteTablesResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.infra.ListRouteTablesResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListRouteTablesResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getRouteTablesList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      types_pb.RouteTable.serializeBinaryToWriter
    );
  }
  f = message.getLastSyncTime();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getErr();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      types_pb.Error.serializeBinaryToWriter
    );
  }
};


/**
 * repeated RouteTable route_tables = 1;
 * @return {!Array<!proto.infra.RouteTable>}
 */
proto.infra.ListRouteTablesResponse.prototype.getRouteTablesList = function() {
  return /** @type{!Array<!proto.infra.RouteTable>} */ (
    jspb.Message.getRepeatedWrapperField(this, types_pb.RouteTable, 1));
};


/**
 * @param {!Array<!proto.infra.RouteTable>} value
 * @return {!proto.infra.ListRouteTablesResponse} returns this
*/
proto.infra.ListRouteTablesResponse.prototype.setRouteTablesList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.infra.RouteTable=} opt_value
 * @param {number=} opt_index
 * @return {!proto.infra.RouteTable}
 */
proto.infra.ListRouteTablesResponse.prototype.addRouteTables = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.infra.RouteTable, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.infra.ListRouteTablesResponse} returns this
 */
proto.infra.ListRouteTablesResponse.prototype.clearRouteTablesList = function() {
  return this.setRouteTablesList([]);
};


/**
 * optional string last_sync_time = 2;
 * @return {string}
 */
proto.infra.ListRouteTablesResponse.prototype.getLastSyncTime = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListRouteTablesResponse} returns this
 */
proto.infra.ListRouteTablesResponse.prototype.setLastSyncTime = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional Error err = 3;
 * @return {?proto.infra.Error}
 */
proto.infra.ListRouteTablesResponse.prototype.getErr = function() {
  return /** @type{?proto.infra.Error} */ (
    jspb.Message.getWrapperField(this, types_pb.Error, 3));
};


/**
 * @param {?proto.infra.Error|undefined} value
 * @return {!proto.infra.ListRouteTablesResponse} returns this
*/
proto.infra.ListRouteTablesResponse.prototype.setErr = function(value) {
  return jspb.Message.setWrapperField(this, 3, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.infra.ListRouteTablesResponse} returns this
 */
proto.infra.ListRouteTablesResponse.prototype.clearErr = function() {
  return this.setErr(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.infra.ListRouteTablesResponse.prototype.hasErr = function() {
  return jspb.Message.getField(this, 3) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.infra.ListNATGatewaysRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.infra.ListNATGatewaysRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.infra.ListNATGatewaysRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListNATGatewaysRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    provider: jspb.Message.getFieldWithDefault(msg, 1, ""),
    accountId: jspb.Message.getFieldWithDefault(msg, 2, ""),
    region: jspb.Message.getFieldWithDefault(msg, 3, ""),
    vpcId: jspb.Message.getFieldWithDefault(msg, 4, ""),
    labelsMap: (f = msg.getLabelsMap()) ? f.toObject(includeInstance, undefined) : [],
    creds: (f = msg.getCreds()) && types_pb.Credentials.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.infra.ListNATGatewaysRequest}
 */
proto.infra.ListNATGatewaysRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.infra.ListNATGatewaysRequest;
  return proto.infra.ListNATGatewaysRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.infra.ListNATGatewaysRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.infra.ListNATGatewaysRequest}
 */
proto.infra.ListNATGatewaysRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setProvider(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setAccountId(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setRegion(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setVpcId(value);
      break;
    case 5:
      var value = msg.getLabelsMap();
      reader.readMessage(value, function(message, reader) {
        jspb.Map.deserializeBinary(message, reader, jspb.BinaryReader.prototype.readString, jspb.BinaryReader.prototype.readString, null, "", "");
         });
      break;
    case 6:
      var value = new types_pb.Credentials;
      reader.readMessage(value,types_pb.Credentials.deserializeBinaryFromReader);
      msg.setCreds(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.infra.ListNATGatewaysRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.infra.ListNATGatewaysRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.infra.ListNATGatewaysRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListNATGatewaysRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getProvider();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getAccountId();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getRegion();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getVpcId();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getLabelsMap(true);
  if (f && f.getLength() > 0) {
    f.serializeBinary(5, writer, jspb.BinaryWriter.prototype.writeString, jspb.BinaryWriter.prototype.writeString);
  }
  f = message.getCreds();
  if (f != null) {
    writer.writeMessage(
      6,
      f,
      types_pb.Credentials.serializeBinaryToWriter
    );
  }
};


/**
 * optional string provider = 1;
 * @return {string}
 */
proto.infra.ListNATGatewaysRequest.prototype.getProvider = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListNATGatewaysRequest} returns this
 */
proto.infra.ListNATGatewaysRequest.prototype.setProvider = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string account_id = 2;
 * @return {string}
 */
proto.infra.ListNATGatewaysRequest.prototype.getAccountId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListNATGatewaysRequest} returns this
 */
proto.infra.ListNATGatewaysRequest.prototype.setAccountId = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string region = 3;
 * @return {string}
 */
proto.infra.ListNATGatewaysRequest.prototype.getRegion = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListNATGatewaysRequest} returns this
 */
proto.infra.ListNATGatewaysRequest.prototype.setRegion = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string vpc_id = 4;
 * @return {string}
 */
proto.infra.ListNATGatewaysRequest.prototype.getVpcId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListNATGatewaysRequest} returns this
 */
proto.infra.ListNATGatewaysRequest.prototype.setVpcId = function(value) {
  return jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * map<string, string> labels = 5;
 * @param {boolean=} opt_noLazyCreate Do not create the map if
 * empty, instead returning `undefined`
 * @return {!jspb.Map<string,string>}
 */
proto.infra.ListNATGatewaysRequest.prototype.getLabelsMap = function(opt_noLazyCreate) {
  return /** @type {!jspb.Map<string,string>} */ (
      jspb.Message.getMapField(this, 5, opt_noLazyCreate,
      null));
};


/**
 * Clears values from the map. The map will be non-null.
 * @return {!proto.infra.ListNATGatewaysRequest} returns this
 */
proto.infra.ListNATGatewaysRequest.prototype.clearLabelsMap = function() {
  this.getLabelsMap().clear();
  return this;
};


/**
 * optional Credentials creds = 6;
 * @return {?proto.infra.Credentials}
 */
proto.infra.ListNATGatewaysRequest.prototype.getCreds = function() {
  return /** @type{?proto.infra.Credentials} */ (
    jspb.Message.getWrapperField(this, types_pb.Credentials, 6));
};


/**
 * @param {?proto.infra.Credentials|undefined} value
 * @return {!proto.infra.ListNATGatewaysRequest} returns this
*/
proto.infra.ListNATGatewaysRequest.prototype.setCreds = function(value) {
  return jspb.Message.setWrapperField(this, 6, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.infra.ListNATGatewaysRequest} returns this
 */
proto.infra.ListNATGatewaysRequest.prototype.clearCreds = function() {
  return this.setCreds(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.infra.ListNATGatewaysRequest.prototype.hasCreds = function() {
  return jspb.Message.getField(this, 6) != null;
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.infra.ListNATGatewaysResponse.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.infra.ListNATGatewaysResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.infra.ListNATGatewaysResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.infra.ListNATGatewaysResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListNATGatewaysResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    natGatewaysList: jspb.Message.toObjectList(msg.getNatGatewaysList(),
    types_pb.NATGateway.toObject, includeInstance),
    lastSyncTime: jspb.Message.getFieldWithDefault(msg, 2, ""),
    err: (f = msg.getErr()) && types_pb.Error.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.infra.ListNATGatewaysResponse}
 */
proto.infra.ListNATGatewaysResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.infra.ListNATGatewaysResponse;
  return proto.infra.ListNATGatewaysResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.infra.ListNATGatewaysResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.infra.ListNATGatewaysResponse}
 */
proto.infra.ListNATGatewaysResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new types_pb.NATGateway;
      reader.readMessage(value,types_pb.NATGateway.deserializeBinaryFromReader);
      msg.addNatGateways(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setLastSyncTime(value);
      break;
    case 3:
      var value = new types_pb.Error;
      reader.readMessage(value,types_pb.Error.deserializeBinaryFromReader);
      msg.setErr(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.infra.ListNATGatewaysResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.infra.ListNATGatewaysResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.infra.ListNATGatewaysResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListNATGatewaysResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getNatGatewaysList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      types_pb.NATGateway.serializeBinaryToWriter
    );
  }
  f = message.getLastSyncTime();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getErr();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      types_pb.Error.serializeBinaryToWriter
    );
  }
};


/**
 * repeated NATGateway nat_gateways = 1;
 * @return {!Array<!proto.infra.NATGateway>}
 */
proto.infra.ListNATGatewaysResponse.prototype.getNatGatewaysList = function() {
  return /** @type{!Array<!proto.infra.NATGateway>} */ (
    jspb.Message.getRepeatedWrapperField(this, types_pb.NATGateway, 1));
};


/**
 * @param {!Array<!proto.infra.NATGateway>} value
 * @return {!proto.infra.ListNATGatewaysResponse} returns this
*/
proto.infra.ListNATGatewaysResponse.prototype.setNatGatewaysList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.infra.NATGateway=} opt_value
 * @param {number=} opt_index
 * @return {!proto.infra.NATGateway}
 */
proto.infra.ListNATGatewaysResponse.prototype.addNatGateways = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.infra.NATGateway, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.infra.ListNATGatewaysResponse} returns this
 */
proto.infra.ListNATGatewaysResponse.prototype.clearNatGatewaysList = function() {
  return this.setNatGatewaysList([]);
};


/**
 * optional string last_sync_time = 2;
 * @return {string}
 */
proto.infra.ListNATGatewaysResponse.prototype.getLastSyncTime = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListNATGatewaysResponse} returns this
 */
proto.infra.ListNATGatewaysResponse.prototype.setLastSyncTime = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional Error err = 3;
 * @return {?proto.infra.Error}
 */
proto.infra.ListNATGatewaysResponse.prototype.getErr = function() {
  return /** @type{?proto.infra.Error} */ (
    jspb.Message.getWrapperField(this, types_pb.Error, 3));
};


/**
 * @param {?proto.infra.Error|undefined} value
 * @return {!proto.infra.ListNATGatewaysResponse} returns this
*/
proto.infra.ListNATGatewaysResponse.prototype.setErr = function(value) {
  return jspb.Message.setWrapperField(this, 3, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.infra.ListNATGatewaysResponse} returns this
 */
proto.infra.ListNATGatewaysResponse.prototype.clearErr = function() {
  return this.setErr(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.infra.ListNATGatewaysResponse.prototype.hasErr = function() {
  return jspb.Message.getField(this, 3) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.infra.ListRoutersRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.infra.ListRoutersRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.infra.ListRoutersRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListRoutersRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    provider: jspb.Message.getFieldWithDefault(msg, 1, ""),
    accountId: jspb.Message.getFieldWithDefault(msg, 2, ""),
    region: jspb.Message.getFieldWithDefault(msg, 3, ""),
    vpcId: jspb.Message.getFieldWithDefault(msg, 4, ""),
    labelsMap: (f = msg.getLabelsMap()) ? f.toObject(includeInstance, undefined) : [],
    creds: (f = msg.getCreds()) && types_pb.Credentials.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.infra.ListRoutersRequest}
 */
proto.infra.ListRoutersRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.infra.ListRoutersRequest;
  return proto.infra.ListRoutersRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.infra.ListRoutersRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.infra.ListRoutersRequest}
 */
proto.infra.ListRoutersRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setProvider(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setAccountId(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setRegion(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setVpcId(value);
      break;
    case 5:
      var value = msg.getLabelsMap();
      reader.readMessage(value, function(message, reader) {
        jspb.Map.deserializeBinary(message, reader, jspb.BinaryReader.prototype.readString, jspb.BinaryReader.prototype.readString, null, "", "");
         });
      break;
    case 6:
      var value = new types_pb.Credentials;
      reader.readMessage(value,types_pb.Credentials.deserializeBinaryFromReader);
      msg.setCreds(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.infra.ListRoutersRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.infra.ListRoutersRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.infra.ListRoutersRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListRoutersRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getProvider();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getAccountId();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getRegion();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getVpcId();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getLabelsMap(true);
  if (f && f.getLength() > 0) {
    f.serializeBinary(5, writer, jspb.BinaryWriter.prototype.writeString, jspb.BinaryWriter.prototype.writeString);
  }
  f = message.getCreds();
  if (f != null) {
    writer.writeMessage(
      6,
      f,
      types_pb.Credentials.serializeBinaryToWriter
    );
  }
};


/**
 * optional string provider = 1;
 * @return {string}
 */
proto.infra.ListRoutersRequest.prototype.getProvider = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListRoutersRequest} returns this
 */
proto.infra.ListRoutersRequest.prototype.setProvider = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string account_id = 2;
 * @return {string}
 */
proto.infra.ListRoutersRequest.prototype.getAccountId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListRoutersRequest} returns this
 */
proto.infra.ListRoutersRequest.prototype.setAccountId = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string region = 3;
 * @return {string}
 */
proto.infra.ListRoutersRequest.prototype.getRegion = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListRoutersRequest} returns this
 */
proto.infra.ListRoutersRequest.prototype.setRegion = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string vpc_id = 4;
 * @return {string}
 */
proto.infra.ListRoutersRequest.prototype.getVpcId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListRoutersRequest} returns this
 */
proto.infra.ListRoutersRequest.prototype.setVpcId = function(value) {
  return jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * map<string, string> labels = 5;
 * @param {boolean=} opt_noLazyCreate Do not create the map if
 * empty, instead returning `undefined`
 * @return {!jspb.Map<string,string>}
 */
proto.infra.ListRoutersRequest.prototype.getLabelsMap = function(opt_noLazyCreate) {
  return /** @type {!jspb.Map<string,string>} */ (
      jspb.Message.getMapField(this, 5, opt_noLazyCreate,
      null));
};


/**
 * Clears values from the map. The map will be non-null.
 * @return {!proto.infra.ListRoutersRequest} returns this
 */
proto.infra.ListRoutersRequest.prototype.clearLabelsMap = function() {
  this.getLabelsMap().clear();
  return this;
};


/**
 * optional Credentials creds = 6;
 * @return {?proto.infra.Credentials}
 */
proto.infra.ListRoutersRequest.prototype.getCreds = function() {
  return /** @type{?proto.infra.Credentials} */ (
    jspb.Message.getWrapperField(this, types_pb.Credentials, 6));
};


/**
 * @param {?proto.infra.Credentials|undefined} value
 * @return {!proto.infra.ListRoutersRequest} returns this
*/
proto.infra.ListRoutersRequest.prototype.setCreds = function(value) {
  return jspb.Message.setWrapperField(this, 6, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.infra.ListRoutersRequest} returns this
 */
proto.infra.ListRoutersRequest.prototype.clearCreds = function() {
  return this.setCreds(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.infra.ListRoutersRequest.prototype.hasCreds = function() {
  return jspb.Message.getField(this, 6) != null;
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.infra.ListRoutersResponse.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.infra.ListRoutersResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.infra.ListRoutersResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.infra.ListRoutersResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListRoutersResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    routersList: jspb.Message.toObjectList(msg.getRoutersList(),
    types_pb.Router.toObject, includeInstance),
    lastSyncTime: jspb.Message.getFieldWithDefault(msg, 2, ""),
    err: (f = msg.getErr()) && types_pb.Error.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.infra.ListRoutersResponse}
 */
proto.infra.ListRoutersResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.infra.ListRoutersResponse;
  return proto.infra.ListRoutersResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.infra.ListRoutersResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.infra.ListRoutersResponse}
 */
proto.infra.ListRoutersResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new types_pb.Router;
      reader.readMessage(value,types_pb.Router.deserializeBinaryFromReader);
      msg.addRouters(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setLastSyncTime(value);
      break;
    case 3:
      var value = new types_pb.Error;
      reader.readMessage(value,types_pb.Error.deserializeBinaryFromReader);
      msg.setErr(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.infra.ListRoutersResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.infra.ListRoutersResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.infra.ListRoutersResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListRoutersResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getRoutersList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      types_pb.Router.serializeBinaryToWriter
    );
  }
  f = message.getLastSyncTime();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getErr();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      types_pb.Error.serializeBinaryToWriter
    );
  }
};


/**
 * repeated Router routers = 1;
 * @return {!Array<!proto.infra.Router>}
 */
proto.infra.ListRoutersResponse.prototype.getRoutersList = function() {
  return /** @type{!Array<!proto.infra.Router>} */ (
    jspb.Message.getRepeatedWrapperField(this, types_pb.Router, 1));
};


/**
 * @param {!Array<!proto.infra.Router>} value
 * @return {!proto.infra.ListRoutersResponse} returns this
*/
proto.infra.ListRoutersResponse.prototype.setRoutersList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.infra.Router=} opt_value
 * @param {number=} opt_index
 * @return {!proto.infra.Router}
 */
proto.infra.ListRoutersResponse.prototype.addRouters = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.infra.Router, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.infra.ListRoutersResponse} returns this
 */
proto.infra.ListRoutersResponse.prototype.clearRoutersList = function() {
  return this.setRoutersList([]);
};


/**
 * optional string last_sync_time = 2;
 * @return {string}
 */
proto.infra.ListRoutersResponse.prototype.getLastSyncTime = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListRoutersResponse} returns this
 */
proto.infra.ListRoutersResponse.prototype.setLastSyncTime = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional Error err = 3;
 * @return {?proto.infra.Error}
 */
proto.infra.ListRoutersResponse.prototype.getErr = function() {
  return /** @type{?proto.infra.Error} */ (
    jspb.Message.getWrapperField(this, types_pb.Error, 3));
};


/**
 * @param {?proto.infra.Error|undefined} value
 * @return {!proto.infra.ListRoutersResponse} returns this
*/
proto.infra.ListRoutersResponse.prototype.setErr = function(value) {
  return jspb.Message.setWrapperField(this, 3, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.infra.ListRoutersResponse} returns this
 */
proto.infra.ListRoutersResponse.prototype.clearErr = function() {
  return this.setErr(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.infra.ListRoutersResponse.prototype.hasErr = function() {
  return jspb.Message.getField(this, 3) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.infra.ListInternetGatewaysRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.infra.ListInternetGatewaysRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.infra.ListInternetGatewaysRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListInternetGatewaysRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    provider: jspb.Message.getFieldWithDefault(msg, 1, ""),
    accountId: jspb.Message.getFieldWithDefault(msg, 2, ""),
    region: jspb.Message.getFieldWithDefault(msg, 3, ""),
    vpcId: jspb.Message.getFieldWithDefault(msg, 4, ""),
    labelsMap: (f = msg.getLabelsMap()) ? f.toObject(includeInstance, undefined) : [],
    creds: (f = msg.getCreds()) && types_pb.Credentials.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.infra.ListInternetGatewaysRequest}
 */
proto.infra.ListInternetGatewaysRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.infra.ListInternetGatewaysRequest;
  return proto.infra.ListInternetGatewaysRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.infra.ListInternetGatewaysRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.infra.ListInternetGatewaysRequest}
 */
proto.infra.ListInternetGatewaysRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setProvider(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setAccountId(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setRegion(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setVpcId(value);
      break;
    case 5:
      var value = msg.getLabelsMap();
      reader.readMessage(value, function(message, reader) {
        jspb.Map.deserializeBinary(message, reader, jspb.BinaryReader.prototype.readString, jspb.BinaryReader.prototype.readString, null, "", "");
         });
      break;
    case 6:
      var value = new types_pb.Credentials;
      reader.readMessage(value,types_pb.Credentials.deserializeBinaryFromReader);
      msg.setCreds(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.infra.ListInternetGatewaysRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.infra.ListInternetGatewaysRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.infra.ListInternetGatewaysRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListInternetGatewaysRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getProvider();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getAccountId();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getRegion();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getVpcId();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getLabelsMap(true);
  if (f && f.getLength() > 0) {
    f.serializeBinary(5, writer, jspb.BinaryWriter.prototype.writeString, jspb.BinaryWriter.prototype.writeString);
  }
  f = message.getCreds();
  if (f != null) {
    writer.writeMessage(
      6,
      f,
      types_pb.Credentials.serializeBinaryToWriter
    );
  }
};


/**
 * optional string provider = 1;
 * @return {string}
 */
proto.infra.ListInternetGatewaysRequest.prototype.getProvider = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListInternetGatewaysRequest} returns this
 */
proto.infra.ListInternetGatewaysRequest.prototype.setProvider = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string account_id = 2;
 * @return {string}
 */
proto.infra.ListInternetGatewaysRequest.prototype.getAccountId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListInternetGatewaysRequest} returns this
 */
proto.infra.ListInternetGatewaysRequest.prototype.setAccountId = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string region = 3;
 * @return {string}
 */
proto.infra.ListInternetGatewaysRequest.prototype.getRegion = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListInternetGatewaysRequest} returns this
 */
proto.infra.ListInternetGatewaysRequest.prototype.setRegion = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string vpc_id = 4;
 * @return {string}
 */
proto.infra.ListInternetGatewaysRequest.prototype.getVpcId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListInternetGatewaysRequest} returns this
 */
proto.infra.ListInternetGatewaysRequest.prototype.setVpcId = function(value) {
  return jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * map<string, string> labels = 5;
 * @param {boolean=} opt_noLazyCreate Do not create the map if
 * empty, instead returning `undefined`
 * @return {!jspb.Map<string,string>}
 */
proto.infra.ListInternetGatewaysRequest.prototype.getLabelsMap = function(opt_noLazyCreate) {
  return /** @type {!jspb.Map<string,string>} */ (
      jspb.Message.getMapField(this, 5, opt_noLazyCreate,
      null));
};


/**
 * Clears values from the map. The map will be non-null.
 * @return {!proto.infra.ListInternetGatewaysRequest} returns this
 */
proto.infra.ListInternetGatewaysRequest.prototype.clearLabelsMap = function() {
  this.getLabelsMap().clear();
  return this;
};


/**
 * optional Credentials creds = 6;
 * @return {?proto.infra.Credentials}
 */
proto.infra.ListInternetGatewaysRequest.prototype.getCreds = function() {
  return /** @type{?proto.infra.Credentials} */ (
    jspb.Message.getWrapperField(this, types_pb.Credentials, 6));
};


/**
 * @param {?proto.infra.Credentials|undefined} value
 * @return {!proto.infra.ListInternetGatewaysRequest} returns this
*/
proto.infra.ListInternetGatewaysRequest.prototype.setCreds = function(value) {
  return jspb.Message.setWrapperField(this, 6, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.infra.ListInternetGatewaysRequest} returns this
 */
proto.infra.ListInternetGatewaysRequest.prototype.clearCreds = function() {
  return this.setCreds(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.infra.ListInternetGatewaysRequest.prototype.hasCreds = function() {
  return jspb.Message.getField(this, 6) != null;
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.infra.ListInternetGatewaysResponse.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.infra.ListInternetGatewaysResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.infra.ListInternetGatewaysResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.infra.ListInternetGatewaysResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListInternetGatewaysResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    igwsList: jspb.Message.toObjectList(msg.getIgwsList(),
    types_pb.IGW.toObject, includeInstance),
    lastSyncTime: jspb.Message.getFieldWithDefault(msg, 2, ""),
    err: (f = msg.getErr()) && types_pb.Error.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.infra.ListInternetGatewaysResponse}
 */
proto.infra.ListInternetGatewaysResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.infra.ListInternetGatewaysResponse;
  return proto.infra.ListInternetGatewaysResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.infra.ListInternetGatewaysResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.infra.ListInternetGatewaysResponse}
 */
proto.infra.ListInternetGatewaysResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new types_pb.IGW;
      reader.readMessage(value,types_pb.IGW.deserializeBinaryFromReader);
      msg.addIgws(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setLastSyncTime(value);
      break;
    case 3:
      var value = new types_pb.Error;
      reader.readMessage(value,types_pb.Error.deserializeBinaryFromReader);
      msg.setErr(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.infra.ListInternetGatewaysResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.infra.ListInternetGatewaysResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.infra.ListInternetGatewaysResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListInternetGatewaysResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getIgwsList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      types_pb.IGW.serializeBinaryToWriter
    );
  }
  f = message.getLastSyncTime();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getErr();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      types_pb.Error.serializeBinaryToWriter
    );
  }
};


/**
 * repeated IGW igws = 1;
 * @return {!Array<!proto.infra.IGW>}
 */
proto.infra.ListInternetGatewaysResponse.prototype.getIgwsList = function() {
  return /** @type{!Array<!proto.infra.IGW>} */ (
    jspb.Message.getRepeatedWrapperField(this, types_pb.IGW, 1));
};


/**
 * @param {!Array<!proto.infra.IGW>} value
 * @return {!proto.infra.ListInternetGatewaysResponse} returns this
*/
proto.infra.ListInternetGatewaysResponse.prototype.setIgwsList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.infra.IGW=} opt_value
 * @param {number=} opt_index
 * @return {!proto.infra.IGW}
 */
proto.infra.ListInternetGatewaysResponse.prototype.addIgws = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.infra.IGW, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.infra.ListInternetGatewaysResponse} returns this
 */
proto.infra.ListInternetGatewaysResponse.prototype.clearIgwsList = function() {
  return this.setIgwsList([]);
};


/**
 * optional string last_sync_time = 2;
 * @return {string}
 */
proto.infra.ListInternetGatewaysResponse.prototype.getLastSyncTime = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListInternetGatewaysResponse} returns this
 */
proto.infra.ListInternetGatewaysResponse.prototype.setLastSyncTime = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional Error err = 3;
 * @return {?proto.infra.Error}
 */
proto.infra.ListInternetGatewaysResponse.prototype.getErr = function() {
  return /** @type{?proto.infra.Error} */ (
    jspb.Message.getWrapperField(this, types_pb.Error, 3));
};


/**
 * @param {?proto.infra.Error|undefined} value
 * @return {!proto.infra.ListInternetGatewaysResponse} returns this
*/
proto.infra.ListInternetGatewaysResponse.prototype.setErr = function(value) {
  return jspb.Message.setWrapperField(this, 3, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.infra.ListInternetGatewaysResponse} returns this
 */
proto.infra.ListInternetGatewaysResponse.prototype.clearErr = function() {
  return this.setErr(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.infra.ListInternetGatewaysResponse.prototype.hasErr = function() {
  return jspb.Message.getField(this, 3) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.infra.ListVPCEndpointsRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.infra.ListVPCEndpointsRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.infra.ListVPCEndpointsRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListVPCEndpointsRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    provider: jspb.Message.getFieldWithDefault(msg, 1, ""),
    accountId: jspb.Message.getFieldWithDefault(msg, 2, ""),
    region: jspb.Message.getFieldWithDefault(msg, 3, ""),
    vpcId: jspb.Message.getFieldWithDefault(msg, 4, ""),
    labelsMap: (f = msg.getLabelsMap()) ? f.toObject(includeInstance, undefined) : [],
    creds: (f = msg.getCreds()) && types_pb.Credentials.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.infra.ListVPCEndpointsRequest}
 */
proto.infra.ListVPCEndpointsRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.infra.ListVPCEndpointsRequest;
  return proto.infra.ListVPCEndpointsRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.infra.ListVPCEndpointsRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.infra.ListVPCEndpointsRequest}
 */
proto.infra.ListVPCEndpointsRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setProvider(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setAccountId(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setRegion(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setVpcId(value);
      break;
    case 5:
      var value = msg.getLabelsMap();
      reader.readMessage(value, function(message, reader) {
        jspb.Map.deserializeBinary(message, reader, jspb.BinaryReader.prototype.readString, jspb.BinaryReader.prototype.readString, null, "", "");
         });
      break;
    case 6:
      var value = new types_pb.Credentials;
      reader.readMessage(value,types_pb.Credentials.deserializeBinaryFromReader);
      msg.setCreds(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.infra.ListVPCEndpointsRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.infra.ListVPCEndpointsRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.infra.ListVPCEndpointsRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListVPCEndpointsRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getProvider();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getAccountId();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getRegion();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getVpcId();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getLabelsMap(true);
  if (f && f.getLength() > 0) {
    f.serializeBinary(5, writer, jspb.BinaryWriter.prototype.writeString, jspb.BinaryWriter.prototype.writeString);
  }
  f = message.getCreds();
  if (f != null) {
    writer.writeMessage(
      6,
      f,
      types_pb.Credentials.serializeBinaryToWriter
    );
  }
};


/**
 * optional string provider = 1;
 * @return {string}
 */
proto.infra.ListVPCEndpointsRequest.prototype.getProvider = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListVPCEndpointsRequest} returns this
 */
proto.infra.ListVPCEndpointsRequest.prototype.setProvider = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string account_id = 2;
 * @return {string}
 */
proto.infra.ListVPCEndpointsRequest.prototype.getAccountId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListVPCEndpointsRequest} returns this
 */
proto.infra.ListVPCEndpointsRequest.prototype.setAccountId = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string region = 3;
 * @return {string}
 */
proto.infra.ListVPCEndpointsRequest.prototype.getRegion = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListVPCEndpointsRequest} returns this
 */
proto.infra.ListVPCEndpointsRequest.prototype.setRegion = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string vpc_id = 4;
 * @return {string}
 */
proto.infra.ListVPCEndpointsRequest.prototype.getVpcId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListVPCEndpointsRequest} returns this
 */
proto.infra.ListVPCEndpointsRequest.prototype.setVpcId = function(value) {
  return jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * map<string, string> labels = 5;
 * @param {boolean=} opt_noLazyCreate Do not create the map if
 * empty, instead returning `undefined`
 * @return {!jspb.Map<string,string>}
 */
proto.infra.ListVPCEndpointsRequest.prototype.getLabelsMap = function(opt_noLazyCreate) {
  return /** @type {!jspb.Map<string,string>} */ (
      jspb.Message.getMapField(this, 5, opt_noLazyCreate,
      null));
};


/**
 * Clears values from the map. The map will be non-null.
 * @return {!proto.infra.ListVPCEndpointsRequest} returns this
 */
proto.infra.ListVPCEndpointsRequest.prototype.clearLabelsMap = function() {
  this.getLabelsMap().clear();
  return this;
};


/**
 * optional Credentials creds = 6;
 * @return {?proto.infra.Credentials}
 */
proto.infra.ListVPCEndpointsRequest.prototype.getCreds = function() {
  return /** @type{?proto.infra.Credentials} */ (
    jspb.Message.getWrapperField(this, types_pb.Credentials, 6));
};


/**
 * @param {?proto.infra.Credentials|undefined} value
 * @return {!proto.infra.ListVPCEndpointsRequest} returns this
*/
proto.infra.ListVPCEndpointsRequest.prototype.setCreds = function(value) {
  return jspb.Message.setWrapperField(this, 6, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.infra.ListVPCEndpointsRequest} returns this
 */
proto.infra.ListVPCEndpointsRequest.prototype.clearCreds = function() {
  return this.setCreds(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.infra.ListVPCEndpointsRequest.prototype.hasCreds = function() {
  return jspb.Message.getField(this, 6) != null;
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.infra.ListVPCEndpointsResponse.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.infra.ListVPCEndpointsResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.infra.ListVPCEndpointsResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.infra.ListVPCEndpointsResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListVPCEndpointsResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    vepsList: jspb.Message.toObjectList(msg.getVepsList(),
    types_pb.VPCEndpoint.toObject, includeInstance),
    lastSyncTime: jspb.Message.getFieldWithDefault(msg, 2, ""),
    err: (f = msg.getErr()) && types_pb.Error.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.infra.ListVPCEndpointsResponse}
 */
proto.infra.ListVPCEndpointsResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.infra.ListVPCEndpointsResponse;
  return proto.infra.ListVPCEndpointsResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.infra.ListVPCEndpointsResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.infra.ListVPCEndpointsResponse}
 */
proto.infra.ListVPCEndpointsResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new types_pb.VPCEndpoint;
      reader.readMessage(value,types_pb.VPCEndpoint.deserializeBinaryFromReader);
      msg.addVeps(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setLastSyncTime(value);
      break;
    case 3:
      var value = new types_pb.Error;
      reader.readMessage(value,types_pb.Error.deserializeBinaryFromReader);
      msg.setErr(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.infra.ListVPCEndpointsResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.infra.ListVPCEndpointsResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.infra.ListVPCEndpointsResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListVPCEndpointsResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getVepsList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      types_pb.VPCEndpoint.serializeBinaryToWriter
    );
  }
  f = message.getLastSyncTime();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getErr();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      types_pb.Error.serializeBinaryToWriter
    );
  }
};


/**
 * repeated VPCEndpoint veps = 1;
 * @return {!Array<!proto.infra.VPCEndpoint>}
 */
proto.infra.ListVPCEndpointsResponse.prototype.getVepsList = function() {
  return /** @type{!Array<!proto.infra.VPCEndpoint>} */ (
    jspb.Message.getRepeatedWrapperField(this, types_pb.VPCEndpoint, 1));
};


/**
 * @param {!Array<!proto.infra.VPCEndpoint>} value
 * @return {!proto.infra.ListVPCEndpointsResponse} returns this
*/
proto.infra.ListVPCEndpointsResponse.prototype.setVepsList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.infra.VPCEndpoint=} opt_value
 * @param {number=} opt_index
 * @return {!proto.infra.VPCEndpoint}
 */
proto.infra.ListVPCEndpointsResponse.prototype.addVeps = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.infra.VPCEndpoint, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.infra.ListVPCEndpointsResponse} returns this
 */
proto.infra.ListVPCEndpointsResponse.prototype.clearVepsList = function() {
  return this.setVepsList([]);
};


/**
 * optional string last_sync_time = 2;
 * @return {string}
 */
proto.infra.ListVPCEndpointsResponse.prototype.getLastSyncTime = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListVPCEndpointsResponse} returns this
 */
proto.infra.ListVPCEndpointsResponse.prototype.setLastSyncTime = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional Error err = 3;
 * @return {?proto.infra.Error}
 */
proto.infra.ListVPCEndpointsResponse.prototype.getErr = function() {
  return /** @type{?proto.infra.Error} */ (
    jspb.Message.getWrapperField(this, types_pb.Error, 3));
};


/**
 * @param {?proto.infra.Error|undefined} value
 * @return {!proto.infra.ListVPCEndpointsResponse} returns this
*/
proto.infra.ListVPCEndpointsResponse.prototype.setErr = function(value) {
  return jspb.Message.setWrapperField(this, 3, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.infra.ListVPCEndpointsResponse} returns this
 */
proto.infra.ListVPCEndpointsResponse.prototype.clearErr = function() {
  return this.setErr(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.infra.ListVPCEndpointsResponse.prototype.hasErr = function() {
  return jspb.Message.getField(this, 3) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.infra.ListPublicIPsRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.infra.ListPublicIPsRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.infra.ListPublicIPsRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListPublicIPsRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    provider: jspb.Message.getFieldWithDefault(msg, 1, ""),
    accountId: jspb.Message.getFieldWithDefault(msg, 2, ""),
    region: jspb.Message.getFieldWithDefault(msg, 3, ""),
    vpcId: jspb.Message.getFieldWithDefault(msg, 4, ""),
    labelsMap: (f = msg.getLabelsMap()) ? f.toObject(includeInstance, undefined) : [],
    creds: (f = msg.getCreds()) && types_pb.Credentials.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.infra.ListPublicIPsRequest}
 */
proto.infra.ListPublicIPsRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.infra.ListPublicIPsRequest;
  return proto.infra.ListPublicIPsRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.infra.ListPublicIPsRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.infra.ListPublicIPsRequest}
 */
proto.infra.ListPublicIPsRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setProvider(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setAccountId(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setRegion(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setVpcId(value);
      break;
    case 5:
      var value = msg.getLabelsMap();
      reader.readMessage(value, function(message, reader) {
        jspb.Map.deserializeBinary(message, reader, jspb.BinaryReader.prototype.readString, jspb.BinaryReader.prototype.readString, null, "", "");
         });
      break;
    case 6:
      var value = new types_pb.Credentials;
      reader.readMessage(value,types_pb.Credentials.deserializeBinaryFromReader);
      msg.setCreds(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.infra.ListPublicIPsRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.infra.ListPublicIPsRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.infra.ListPublicIPsRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListPublicIPsRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getProvider();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getAccountId();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getRegion();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getVpcId();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getLabelsMap(true);
  if (f && f.getLength() > 0) {
    f.serializeBinary(5, writer, jspb.BinaryWriter.prototype.writeString, jspb.BinaryWriter.prototype.writeString);
  }
  f = message.getCreds();
  if (f != null) {
    writer.writeMessage(
      6,
      f,
      types_pb.Credentials.serializeBinaryToWriter
    );
  }
};


/**
 * optional string provider = 1;
 * @return {string}
 */
proto.infra.ListPublicIPsRequest.prototype.getProvider = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListPublicIPsRequest} returns this
 */
proto.infra.ListPublicIPsRequest.prototype.setProvider = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string account_id = 2;
 * @return {string}
 */
proto.infra.ListPublicIPsRequest.prototype.getAccountId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListPublicIPsRequest} returns this
 */
proto.infra.ListPublicIPsRequest.prototype.setAccountId = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string region = 3;
 * @return {string}
 */
proto.infra.ListPublicIPsRequest.prototype.getRegion = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListPublicIPsRequest} returns this
 */
proto.infra.ListPublicIPsRequest.prototype.setRegion = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string vpc_id = 4;
 * @return {string}
 */
proto.infra.ListPublicIPsRequest.prototype.getVpcId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListPublicIPsRequest} returns this
 */
proto.infra.ListPublicIPsRequest.prototype.setVpcId = function(value) {
  return jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * map<string, string> labels = 5;
 * @param {boolean=} opt_noLazyCreate Do not create the map if
 * empty, instead returning `undefined`
 * @return {!jspb.Map<string,string>}
 */
proto.infra.ListPublicIPsRequest.prototype.getLabelsMap = function(opt_noLazyCreate) {
  return /** @type {!jspb.Map<string,string>} */ (
      jspb.Message.getMapField(this, 5, opt_noLazyCreate,
      null));
};


/**
 * Clears values from the map. The map will be non-null.
 * @return {!proto.infra.ListPublicIPsRequest} returns this
 */
proto.infra.ListPublicIPsRequest.prototype.clearLabelsMap = function() {
  this.getLabelsMap().clear();
  return this;
};


/**
 * optional Credentials creds = 6;
 * @return {?proto.infra.Credentials}
 */
proto.infra.ListPublicIPsRequest.prototype.getCreds = function() {
  return /** @type{?proto.infra.Credentials} */ (
    jspb.Message.getWrapperField(this, types_pb.Credentials, 6));
};


/**
 * @param {?proto.infra.Credentials|undefined} value
 * @return {!proto.infra.ListPublicIPsRequest} returns this
*/
proto.infra.ListPublicIPsRequest.prototype.setCreds = function(value) {
  return jspb.Message.setWrapperField(this, 6, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.infra.ListPublicIPsRequest} returns this
 */
proto.infra.ListPublicIPsRequest.prototype.clearCreds = function() {
  return this.setCreds(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.infra.ListPublicIPsRequest.prototype.hasCreds = function() {
  return jspb.Message.getField(this, 6) != null;
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.infra.ListPublicIPsResponse.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.infra.ListPublicIPsResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.infra.ListPublicIPsResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.infra.ListPublicIPsResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListPublicIPsResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    publicIpsList: jspb.Message.toObjectList(msg.getPublicIpsList(),
    types_pb.PublicIP.toObject, includeInstance),
    lastSyncTime: jspb.Message.getFieldWithDefault(msg, 2, ""),
    err: (f = msg.getErr()) && types_pb.Error.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.infra.ListPublicIPsResponse}
 */
proto.infra.ListPublicIPsResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.infra.ListPublicIPsResponse;
  return proto.infra.ListPublicIPsResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.infra.ListPublicIPsResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.infra.ListPublicIPsResponse}
 */
proto.infra.ListPublicIPsResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new types_pb.PublicIP;
      reader.readMessage(value,types_pb.PublicIP.deserializeBinaryFromReader);
      msg.addPublicIps(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setLastSyncTime(value);
      break;
    case 3:
      var value = new types_pb.Error;
      reader.readMessage(value,types_pb.Error.deserializeBinaryFromReader);
      msg.setErr(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.infra.ListPublicIPsResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.infra.ListPublicIPsResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.infra.ListPublicIPsResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListPublicIPsResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getPublicIpsList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      types_pb.PublicIP.serializeBinaryToWriter
    );
  }
  f = message.getLastSyncTime();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getErr();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      types_pb.Error.serializeBinaryToWriter
    );
  }
};


/**
 * repeated PublicIP public_ips = 1;
 * @return {!Array<!proto.infra.PublicIP>}
 */
proto.infra.ListPublicIPsResponse.prototype.getPublicIpsList = function() {
  return /** @type{!Array<!proto.infra.PublicIP>} */ (
    jspb.Message.getRepeatedWrapperField(this, types_pb.PublicIP, 1));
};


/**
 * @param {!Array<!proto.infra.PublicIP>} value
 * @return {!proto.infra.ListPublicIPsResponse} returns this
*/
proto.infra.ListPublicIPsResponse.prototype.setPublicIpsList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.infra.PublicIP=} opt_value
 * @param {number=} opt_index
 * @return {!proto.infra.PublicIP}
 */
proto.infra.ListPublicIPsResponse.prototype.addPublicIps = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.infra.PublicIP, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.infra.ListPublicIPsResponse} returns this
 */
proto.infra.ListPublicIPsResponse.prototype.clearPublicIpsList = function() {
  return this.setPublicIpsList([]);
};


/**
 * optional string last_sync_time = 2;
 * @return {string}
 */
proto.infra.ListPublicIPsResponse.prototype.getLastSyncTime = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListPublicIPsResponse} returns this
 */
proto.infra.ListPublicIPsResponse.prototype.setLastSyncTime = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional Error err = 3;
 * @return {?proto.infra.Error}
 */
proto.infra.ListPublicIPsResponse.prototype.getErr = function() {
  return /** @type{?proto.infra.Error} */ (
    jspb.Message.getWrapperField(this, types_pb.Error, 3));
};


/**
 * @param {?proto.infra.Error|undefined} value
 * @return {!proto.infra.ListPublicIPsResponse} returns this
*/
proto.infra.ListPublicIPsResponse.prototype.setErr = function(value) {
  return jspb.Message.setWrapperField(this, 3, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.infra.ListPublicIPsResponse} returns this
 */
proto.infra.ListPublicIPsResponse.prototype.clearErr = function() {
  return this.setErr(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.infra.ListPublicIPsResponse.prototype.hasErr = function() {
  return jspb.Message.getField(this, 3) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.infra.ListLBsRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.infra.ListLBsRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.infra.ListLBsRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListLBsRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    provider: jspb.Message.getFieldWithDefault(msg, 1, ""),
    accountId: jspb.Message.getFieldWithDefault(msg, 2, ""),
    region: jspb.Message.getFieldWithDefault(msg, 3, ""),
    vpcId: jspb.Message.getFieldWithDefault(msg, 4, ""),
    labelsMap: (f = msg.getLabelsMap()) ? f.toObject(includeInstance, undefined) : [],
    creds: (f = msg.getCreds()) && types_pb.Credentials.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.infra.ListLBsRequest}
 */
proto.infra.ListLBsRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.infra.ListLBsRequest;
  return proto.infra.ListLBsRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.infra.ListLBsRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.infra.ListLBsRequest}
 */
proto.infra.ListLBsRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setProvider(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setAccountId(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setRegion(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setVpcId(value);
      break;
    case 5:
      var value = msg.getLabelsMap();
      reader.readMessage(value, function(message, reader) {
        jspb.Map.deserializeBinary(message, reader, jspb.BinaryReader.prototype.readString, jspb.BinaryReader.prototype.readString, null, "", "");
         });
      break;
    case 6:
      var value = new types_pb.Credentials;
      reader.readMessage(value,types_pb.Credentials.deserializeBinaryFromReader);
      msg.setCreds(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.infra.ListLBsRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.infra.ListLBsRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.infra.ListLBsRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListLBsRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getProvider();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getAccountId();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getRegion();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getVpcId();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getLabelsMap(true);
  if (f && f.getLength() > 0) {
    f.serializeBinary(5, writer, jspb.BinaryWriter.prototype.writeString, jspb.BinaryWriter.prototype.writeString);
  }
  f = message.getCreds();
  if (f != null) {
    writer.writeMessage(
      6,
      f,
      types_pb.Credentials.serializeBinaryToWriter
    );
  }
};


/**
 * optional string provider = 1;
 * @return {string}
 */
proto.infra.ListLBsRequest.prototype.getProvider = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListLBsRequest} returns this
 */
proto.infra.ListLBsRequest.prototype.setProvider = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string account_id = 2;
 * @return {string}
 */
proto.infra.ListLBsRequest.prototype.getAccountId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListLBsRequest} returns this
 */
proto.infra.ListLBsRequest.prototype.setAccountId = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string region = 3;
 * @return {string}
 */
proto.infra.ListLBsRequest.prototype.getRegion = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListLBsRequest} returns this
 */
proto.infra.ListLBsRequest.prototype.setRegion = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string vpc_id = 4;
 * @return {string}
 */
proto.infra.ListLBsRequest.prototype.getVpcId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListLBsRequest} returns this
 */
proto.infra.ListLBsRequest.prototype.setVpcId = function(value) {
  return jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * map<string, string> labels = 5;
 * @param {boolean=} opt_noLazyCreate Do not create the map if
 * empty, instead returning `undefined`
 * @return {!jspb.Map<string,string>}
 */
proto.infra.ListLBsRequest.prototype.getLabelsMap = function(opt_noLazyCreate) {
  return /** @type {!jspb.Map<string,string>} */ (
      jspb.Message.getMapField(this, 5, opt_noLazyCreate,
      null));
};


/**
 * Clears values from the map. The map will be non-null.
 * @return {!proto.infra.ListLBsRequest} returns this
 */
proto.infra.ListLBsRequest.prototype.clearLabelsMap = function() {
  this.getLabelsMap().clear();
  return this;
};


/**
 * optional Credentials creds = 6;
 * @return {?proto.infra.Credentials}
 */
proto.infra.ListLBsRequest.prototype.getCreds = function() {
  return /** @type{?proto.infra.Credentials} */ (
    jspb.Message.getWrapperField(this, types_pb.Credentials, 6));
};


/**
 * @param {?proto.infra.Credentials|undefined} value
 * @return {!proto.infra.ListLBsRequest} returns this
*/
proto.infra.ListLBsRequest.prototype.setCreds = function(value) {
  return jspb.Message.setWrapperField(this, 6, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.infra.ListLBsRequest} returns this
 */
proto.infra.ListLBsRequest.prototype.clearCreds = function() {
  return this.setCreds(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.infra.ListLBsRequest.prototype.hasCreds = function() {
  return jspb.Message.getField(this, 6) != null;
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.infra.ListLBsResponse.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.infra.ListLBsResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.infra.ListLBsResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.infra.ListLBsResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListLBsResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    lbsList: jspb.Message.toObjectList(msg.getLbsList(),
    types_pb.LB.toObject, includeInstance),
    lastSyncTime: jspb.Message.getFieldWithDefault(msg, 2, ""),
    err: (f = msg.getErr()) && types_pb.Error.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.infra.ListLBsResponse}
 */
proto.infra.ListLBsResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.infra.ListLBsResponse;
  return proto.infra.ListLBsResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.infra.ListLBsResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.infra.ListLBsResponse}
 */
proto.infra.ListLBsResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new types_pb.LB;
      reader.readMessage(value,types_pb.LB.deserializeBinaryFromReader);
      msg.addLbs(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setLastSyncTime(value);
      break;
    case 3:
      var value = new types_pb.Error;
      reader.readMessage(value,types_pb.Error.deserializeBinaryFromReader);
      msg.setErr(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.infra.ListLBsResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.infra.ListLBsResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.infra.ListLBsResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListLBsResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getLbsList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      types_pb.LB.serializeBinaryToWriter
    );
  }
  f = message.getLastSyncTime();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getErr();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      types_pb.Error.serializeBinaryToWriter
    );
  }
};


/**
 * repeated LB lbs = 1;
 * @return {!Array<!proto.infra.LB>}
 */
proto.infra.ListLBsResponse.prototype.getLbsList = function() {
  return /** @type{!Array<!proto.infra.LB>} */ (
    jspb.Message.getRepeatedWrapperField(this, types_pb.LB, 1));
};


/**
 * @param {!Array<!proto.infra.LB>} value
 * @return {!proto.infra.ListLBsResponse} returns this
*/
proto.infra.ListLBsResponse.prototype.setLbsList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.infra.LB=} opt_value
 * @param {number=} opt_index
 * @return {!proto.infra.LB}
 */
proto.infra.ListLBsResponse.prototype.addLbs = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.infra.LB, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.infra.ListLBsResponse} returns this
 */
proto.infra.ListLBsResponse.prototype.clearLbsList = function() {
  return this.setLbsList([]);
};


/**
 * optional string last_sync_time = 2;
 * @return {string}
 */
proto.infra.ListLBsResponse.prototype.getLastSyncTime = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListLBsResponse} returns this
 */
proto.infra.ListLBsResponse.prototype.setLastSyncTime = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional Error err = 3;
 * @return {?proto.infra.Error}
 */
proto.infra.ListLBsResponse.prototype.getErr = function() {
  return /** @type{?proto.infra.Error} */ (
    jspb.Message.getWrapperField(this, types_pb.Error, 3));
};


/**
 * @param {?proto.infra.Error|undefined} value
 * @return {!proto.infra.ListLBsResponse} returns this
*/
proto.infra.ListLBsResponse.prototype.setErr = function(value) {
  return jspb.Message.setWrapperField(this, 3, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.infra.ListLBsResponse} returns this
 */
proto.infra.ListLBsResponse.prototype.clearErr = function() {
  return this.setErr(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.infra.ListLBsResponse.prototype.hasErr = function() {
  return jspb.Message.getField(this, 3) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.infra.GetSubnetRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.infra.GetSubnetRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.infra.GetSubnetRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.GetSubnetRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    provider: jspb.Message.getFieldWithDefault(msg, 1, ""),
    accountId: jspb.Message.getFieldWithDefault(msg, 2, ""),
    region: jspb.Message.getFieldWithDefault(msg, 3, ""),
    vpcId: jspb.Message.getFieldWithDefault(msg, 4, ""),
    id: jspb.Message.getFieldWithDefault(msg, 5, ""),
    labelsMap: (f = msg.getLabelsMap()) ? f.toObject(includeInstance, undefined) : [],
    creds: (f = msg.getCreds()) && types_pb.Credentials.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.infra.GetSubnetRequest}
 */
proto.infra.GetSubnetRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.infra.GetSubnetRequest;
  return proto.infra.GetSubnetRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.infra.GetSubnetRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.infra.GetSubnetRequest}
 */
proto.infra.GetSubnetRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setProvider(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setAccountId(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setRegion(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setVpcId(value);
      break;
    case 5:
      var value = /** @type {string} */ (reader.readString());
      msg.setId(value);
      break;
    case 6:
      var value = msg.getLabelsMap();
      reader.readMessage(value, function(message, reader) {
        jspb.Map.deserializeBinary(message, reader, jspb.BinaryReader.prototype.readString, jspb.BinaryReader.prototype.readString, null, "", "");
         });
      break;
    case 7:
      var value = new types_pb.Credentials;
      reader.readMessage(value,types_pb.Credentials.deserializeBinaryFromReader);
      msg.setCreds(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.infra.GetSubnetRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.infra.GetSubnetRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.infra.GetSubnetRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.GetSubnetRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getProvider();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getAccountId();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getRegion();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getVpcId();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getId();
  if (f.length > 0) {
    writer.writeString(
      5,
      f
    );
  }
  f = message.getLabelsMap(true);
  if (f && f.getLength() > 0) {
    f.serializeBinary(6, writer, jspb.BinaryWriter.prototype.writeString, jspb.BinaryWriter.prototype.writeString);
  }
  f = message.getCreds();
  if (f != null) {
    writer.writeMessage(
      7,
      f,
      types_pb.Credentials.serializeBinaryToWriter
    );
  }
};


/**
 * optional string provider = 1;
 * @return {string}
 */
proto.infra.GetSubnetRequest.prototype.getProvider = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.GetSubnetRequest} returns this
 */
proto.infra.GetSubnetRequest.prototype.setProvider = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string account_id = 2;
 * @return {string}
 */
proto.infra.GetSubnetRequest.prototype.getAccountId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.GetSubnetRequest} returns this
 */
proto.infra.GetSubnetRequest.prototype.setAccountId = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string region = 3;
 * @return {string}
 */
proto.infra.GetSubnetRequest.prototype.getRegion = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.GetSubnetRequest} returns this
 */
proto.infra.GetSubnetRequest.prototype.setRegion = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string vpc_id = 4;
 * @return {string}
 */
proto.infra.GetSubnetRequest.prototype.getVpcId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.GetSubnetRequest} returns this
 */
proto.infra.GetSubnetRequest.prototype.setVpcId = function(value) {
  return jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * optional string id = 5;
 * @return {string}
 */
proto.infra.GetSubnetRequest.prototype.getId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 5, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.GetSubnetRequest} returns this
 */
proto.infra.GetSubnetRequest.prototype.setId = function(value) {
  return jspb.Message.setProto3StringField(this, 5, value);
};


/**
 * map<string, string> labels = 6;
 * @param {boolean=} opt_noLazyCreate Do not create the map if
 * empty, instead returning `undefined`
 * @return {!jspb.Map<string,string>}
 */
proto.infra.GetSubnetRequest.prototype.getLabelsMap = function(opt_noLazyCreate) {
  return /** @type {!jspb.Map<string,string>} */ (
      jspb.Message.getMapField(this, 6, opt_noLazyCreate,
      null));
};


/**
 * Clears values from the map. The map will be non-null.
 * @return {!proto.infra.GetSubnetRequest} returns this
 */
proto.infra.GetSubnetRequest.prototype.clearLabelsMap = function() {
  this.getLabelsMap().clear();
  return this;
};


/**
 * optional Credentials creds = 7;
 * @return {?proto.infra.Credentials}
 */
proto.infra.GetSubnetRequest.prototype.getCreds = function() {
  return /** @type{?proto.infra.Credentials} */ (
    jspb.Message.getWrapperField(this, types_pb.Credentials, 7));
};


/**
 * @param {?proto.infra.Credentials|undefined} value
 * @return {!proto.infra.GetSubnetRequest} returns this
*/
proto.infra.GetSubnetRequest.prototype.setCreds = function(value) {
  return jspb.Message.setWrapperField(this, 7, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.infra.GetSubnetRequest} returns this
 */
proto.infra.GetSubnetRequest.prototype.clearCreds = function() {
  return this.setCreds(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.infra.GetSubnetRequest.prototype.hasCreds = function() {
  return jspb.Message.getField(this, 7) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.infra.GetSubnetResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.infra.GetSubnetResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.infra.GetSubnetResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.GetSubnetResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    subnet: (f = msg.getSubnet()) && types_pb.Subnet.toObject(includeInstance, f),
    err: (f = msg.getErr()) && types_pb.Error.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.infra.GetSubnetResponse}
 */
proto.infra.GetSubnetResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.infra.GetSubnetResponse;
  return proto.infra.GetSubnetResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.infra.GetSubnetResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.infra.GetSubnetResponse}
 */
proto.infra.GetSubnetResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new types_pb.Subnet;
      reader.readMessage(value,types_pb.Subnet.deserializeBinaryFromReader);
      msg.setSubnet(value);
      break;
    case 3:
      var value = new types_pb.Error;
      reader.readMessage(value,types_pb.Error.deserializeBinaryFromReader);
      msg.setErr(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.infra.GetSubnetResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.infra.GetSubnetResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.infra.GetSubnetResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.GetSubnetResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getSubnet();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      types_pb.Subnet.serializeBinaryToWriter
    );
  }
  f = message.getErr();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      types_pb.Error.serializeBinaryToWriter
    );
  }
};


/**
 * optional Subnet subnet = 1;
 * @return {?proto.infra.Subnet}
 */
proto.infra.GetSubnetResponse.prototype.getSubnet = function() {
  return /** @type{?proto.infra.Subnet} */ (
    jspb.Message.getWrapperField(this, types_pb.Subnet, 1));
};


/**
 * @param {?proto.infra.Subnet|undefined} value
 * @return {!proto.infra.GetSubnetResponse} returns this
*/
proto.infra.GetSubnetResponse.prototype.setSubnet = function(value) {
  return jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.infra.GetSubnetResponse} returns this
 */
proto.infra.GetSubnetResponse.prototype.clearSubnet = function() {
  return this.setSubnet(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.infra.GetSubnetResponse.prototype.hasSubnet = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional Error err = 3;
 * @return {?proto.infra.Error}
 */
proto.infra.GetSubnetResponse.prototype.getErr = function() {
  return /** @type{?proto.infra.Error} */ (
    jspb.Message.getWrapperField(this, types_pb.Error, 3));
};


/**
 * @param {?proto.infra.Error|undefined} value
 * @return {!proto.infra.GetSubnetResponse} returns this
*/
proto.infra.GetSubnetResponse.prototype.setErr = function(value) {
  return jspb.Message.setWrapperField(this, 3, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.infra.GetSubnetResponse} returns this
 */
proto.infra.GetSubnetResponse.prototype.clearErr = function() {
  return this.setErr(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.infra.GetSubnetResponse.prototype.hasErr = function() {
  return jspb.Message.getField(this, 3) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.infra.ListSubnetsRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.infra.ListSubnetsRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.infra.ListSubnetsRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListSubnetsRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    provider: jspb.Message.getFieldWithDefault(msg, 1, ""),
    accountId: jspb.Message.getFieldWithDefault(msg, 2, ""),
    region: jspb.Message.getFieldWithDefault(msg, 3, ""),
    vpcId: jspb.Message.getFieldWithDefault(msg, 4, ""),
    zone: jspb.Message.getFieldWithDefault(msg, 5, ""),
    cidr: jspb.Message.getFieldWithDefault(msg, 6, ""),
    labelsMap: (f = msg.getLabelsMap()) ? f.toObject(includeInstance, undefined) : [],
    creds: (f = msg.getCreds()) && types_pb.Credentials.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.infra.ListSubnetsRequest}
 */
proto.infra.ListSubnetsRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.infra.ListSubnetsRequest;
  return proto.infra.ListSubnetsRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.infra.ListSubnetsRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.infra.ListSubnetsRequest}
 */
proto.infra.ListSubnetsRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setProvider(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setAccountId(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setRegion(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setVpcId(value);
      break;
    case 5:
      var value = /** @type {string} */ (reader.readString());
      msg.setZone(value);
      break;
    case 6:
      var value = /** @type {string} */ (reader.readString());
      msg.setCidr(value);
      break;
    case 7:
      var value = msg.getLabelsMap();
      reader.readMessage(value, function(message, reader) {
        jspb.Map.deserializeBinary(message, reader, jspb.BinaryReader.prototype.readString, jspb.BinaryReader.prototype.readString, null, "", "");
         });
      break;
    case 8:
      var value = new types_pb.Credentials;
      reader.readMessage(value,types_pb.Credentials.deserializeBinaryFromReader);
      msg.setCreds(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.infra.ListSubnetsRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.infra.ListSubnetsRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.infra.ListSubnetsRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListSubnetsRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getProvider();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getAccountId();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getRegion();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getVpcId();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getZone();
  if (f.length > 0) {
    writer.writeString(
      5,
      f
    );
  }
  f = message.getCidr();
  if (f.length > 0) {
    writer.writeString(
      6,
      f
    );
  }
  f = message.getLabelsMap(true);
  if (f && f.getLength() > 0) {
    f.serializeBinary(7, writer, jspb.BinaryWriter.prototype.writeString, jspb.BinaryWriter.prototype.writeString);
  }
  f = message.getCreds();
  if (f != null) {
    writer.writeMessage(
      8,
      f,
      types_pb.Credentials.serializeBinaryToWriter
    );
  }
};


/**
 * optional string provider = 1;
 * @return {string}
 */
proto.infra.ListSubnetsRequest.prototype.getProvider = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListSubnetsRequest} returns this
 */
proto.infra.ListSubnetsRequest.prototype.setProvider = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string account_id = 2;
 * @return {string}
 */
proto.infra.ListSubnetsRequest.prototype.getAccountId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListSubnetsRequest} returns this
 */
proto.infra.ListSubnetsRequest.prototype.setAccountId = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string region = 3;
 * @return {string}
 */
proto.infra.ListSubnetsRequest.prototype.getRegion = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListSubnetsRequest} returns this
 */
proto.infra.ListSubnetsRequest.prototype.setRegion = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string vpc_id = 4;
 * @return {string}
 */
proto.infra.ListSubnetsRequest.prototype.getVpcId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListSubnetsRequest} returns this
 */
proto.infra.ListSubnetsRequest.prototype.setVpcId = function(value) {
  return jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * optional string zone = 5;
 * @return {string}
 */
proto.infra.ListSubnetsRequest.prototype.getZone = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 5, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListSubnetsRequest} returns this
 */
proto.infra.ListSubnetsRequest.prototype.setZone = function(value) {
  return jspb.Message.setProto3StringField(this, 5, value);
};


/**
 * optional string cidr = 6;
 * @return {string}
 */
proto.infra.ListSubnetsRequest.prototype.getCidr = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 6, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListSubnetsRequest} returns this
 */
proto.infra.ListSubnetsRequest.prototype.setCidr = function(value) {
  return jspb.Message.setProto3StringField(this, 6, value);
};


/**
 * map<string, string> labels = 7;
 * @param {boolean=} opt_noLazyCreate Do not create the map if
 * empty, instead returning `undefined`
 * @return {!jspb.Map<string,string>}
 */
proto.infra.ListSubnetsRequest.prototype.getLabelsMap = function(opt_noLazyCreate) {
  return /** @type {!jspb.Map<string,string>} */ (
      jspb.Message.getMapField(this, 7, opt_noLazyCreate,
      null));
};


/**
 * Clears values from the map. The map will be non-null.
 * @return {!proto.infra.ListSubnetsRequest} returns this
 */
proto.infra.ListSubnetsRequest.prototype.clearLabelsMap = function() {
  this.getLabelsMap().clear();
  return this;
};


/**
 * optional Credentials creds = 8;
 * @return {?proto.infra.Credentials}
 */
proto.infra.ListSubnetsRequest.prototype.getCreds = function() {
  return /** @type{?proto.infra.Credentials} */ (
    jspb.Message.getWrapperField(this, types_pb.Credentials, 8));
};


/**
 * @param {?proto.infra.Credentials|undefined} value
 * @return {!proto.infra.ListSubnetsRequest} returns this
*/
proto.infra.ListSubnetsRequest.prototype.setCreds = function(value) {
  return jspb.Message.setWrapperField(this, 8, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.infra.ListSubnetsRequest} returns this
 */
proto.infra.ListSubnetsRequest.prototype.clearCreds = function() {
  return this.setCreds(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.infra.ListSubnetsRequest.prototype.hasCreds = function() {
  return jspb.Message.getField(this, 8) != null;
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.infra.ListSubnetsResponse.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.infra.ListSubnetsResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.infra.ListSubnetsResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.infra.ListSubnetsResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListSubnetsResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    subnetsList: jspb.Message.toObjectList(msg.getSubnetsList(),
    types_pb.Subnet.toObject, includeInstance),
    lastSyncTime: jspb.Message.getFieldWithDefault(msg, 2, ""),
    err: (f = msg.getErr()) && types_pb.Error.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.infra.ListSubnetsResponse}
 */
proto.infra.ListSubnetsResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.infra.ListSubnetsResponse;
  return proto.infra.ListSubnetsResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.infra.ListSubnetsResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.infra.ListSubnetsResponse}
 */
proto.infra.ListSubnetsResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new types_pb.Subnet;
      reader.readMessage(value,types_pb.Subnet.deserializeBinaryFromReader);
      msg.addSubnets(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setLastSyncTime(value);
      break;
    case 3:
      var value = new types_pb.Error;
      reader.readMessage(value,types_pb.Error.deserializeBinaryFromReader);
      msg.setErr(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.infra.ListSubnetsResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.infra.ListSubnetsResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.infra.ListSubnetsResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListSubnetsResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getSubnetsList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      types_pb.Subnet.serializeBinaryToWriter
    );
  }
  f = message.getLastSyncTime();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getErr();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      types_pb.Error.serializeBinaryToWriter
    );
  }
};


/**
 * repeated Subnet subnets = 1;
 * @return {!Array<!proto.infra.Subnet>}
 */
proto.infra.ListSubnetsResponse.prototype.getSubnetsList = function() {
  return /** @type{!Array<!proto.infra.Subnet>} */ (
    jspb.Message.getRepeatedWrapperField(this, types_pb.Subnet, 1));
};


/**
 * @param {!Array<!proto.infra.Subnet>} value
 * @return {!proto.infra.ListSubnetsResponse} returns this
*/
proto.infra.ListSubnetsResponse.prototype.setSubnetsList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.infra.Subnet=} opt_value
 * @param {number=} opt_index
 * @return {!proto.infra.Subnet}
 */
proto.infra.ListSubnetsResponse.prototype.addSubnets = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.infra.Subnet, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.infra.ListSubnetsResponse} returns this
 */
proto.infra.ListSubnetsResponse.prototype.clearSubnetsList = function() {
  return this.setSubnetsList([]);
};


/**
 * optional string last_sync_time = 2;
 * @return {string}
 */
proto.infra.ListSubnetsResponse.prototype.getLastSyncTime = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListSubnetsResponse} returns this
 */
proto.infra.ListSubnetsResponse.prototype.setLastSyncTime = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional Error err = 3;
 * @return {?proto.infra.Error}
 */
proto.infra.ListSubnetsResponse.prototype.getErr = function() {
  return /** @type{?proto.infra.Error} */ (
    jspb.Message.getWrapperField(this, types_pb.Error, 3));
};


/**
 * @param {?proto.infra.Error|undefined} value
 * @return {!proto.infra.ListSubnetsResponse} returns this
*/
proto.infra.ListSubnetsResponse.prototype.setErr = function(value) {
  return jspb.Message.setWrapperField(this, 3, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.infra.ListSubnetsResponse} returns this
 */
proto.infra.ListSubnetsResponse.prototype.clearErr = function() {
  return this.setErr(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.infra.ListSubnetsResponse.prototype.hasErr = function() {
  return jspb.Message.getField(this, 3) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.infra.ListNetworkInterfacesRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.infra.ListNetworkInterfacesRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.infra.ListNetworkInterfacesRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListNetworkInterfacesRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    provider: jspb.Message.getFieldWithDefault(msg, 1, ""),
    accountId: jspb.Message.getFieldWithDefault(msg, 2, ""),
    region: jspb.Message.getFieldWithDefault(msg, 3, ""),
    vpcId: jspb.Message.getFieldWithDefault(msg, 4, ""),
    labelsMap: (f = msg.getLabelsMap()) ? f.toObject(includeInstance, undefined) : [],
    creds: (f = msg.getCreds()) && types_pb.Credentials.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.infra.ListNetworkInterfacesRequest}
 */
proto.infra.ListNetworkInterfacesRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.infra.ListNetworkInterfacesRequest;
  return proto.infra.ListNetworkInterfacesRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.infra.ListNetworkInterfacesRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.infra.ListNetworkInterfacesRequest}
 */
proto.infra.ListNetworkInterfacesRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setProvider(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setAccountId(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setRegion(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setVpcId(value);
      break;
    case 5:
      var value = msg.getLabelsMap();
      reader.readMessage(value, function(message, reader) {
        jspb.Map.deserializeBinary(message, reader, jspb.BinaryReader.prototype.readString, jspb.BinaryReader.prototype.readString, null, "", "");
         });
      break;
    case 6:
      var value = new types_pb.Credentials;
      reader.readMessage(value,types_pb.Credentials.deserializeBinaryFromReader);
      msg.setCreds(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.infra.ListNetworkInterfacesRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.infra.ListNetworkInterfacesRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.infra.ListNetworkInterfacesRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListNetworkInterfacesRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getProvider();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getAccountId();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getRegion();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getVpcId();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getLabelsMap(true);
  if (f && f.getLength() > 0) {
    f.serializeBinary(5, writer, jspb.BinaryWriter.prototype.writeString, jspb.BinaryWriter.prototype.writeString);
  }
  f = message.getCreds();
  if (f != null) {
    writer.writeMessage(
      6,
      f,
      types_pb.Credentials.serializeBinaryToWriter
    );
  }
};


/**
 * optional string provider = 1;
 * @return {string}
 */
proto.infra.ListNetworkInterfacesRequest.prototype.getProvider = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListNetworkInterfacesRequest} returns this
 */
proto.infra.ListNetworkInterfacesRequest.prototype.setProvider = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string account_id = 2;
 * @return {string}
 */
proto.infra.ListNetworkInterfacesRequest.prototype.getAccountId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListNetworkInterfacesRequest} returns this
 */
proto.infra.ListNetworkInterfacesRequest.prototype.setAccountId = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string region = 3;
 * @return {string}
 */
proto.infra.ListNetworkInterfacesRequest.prototype.getRegion = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListNetworkInterfacesRequest} returns this
 */
proto.infra.ListNetworkInterfacesRequest.prototype.setRegion = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string vpc_id = 4;
 * @return {string}
 */
proto.infra.ListNetworkInterfacesRequest.prototype.getVpcId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListNetworkInterfacesRequest} returns this
 */
proto.infra.ListNetworkInterfacesRequest.prototype.setVpcId = function(value) {
  return jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * map<string, string> labels = 5;
 * @param {boolean=} opt_noLazyCreate Do not create the map if
 * empty, instead returning `undefined`
 * @return {!jspb.Map<string,string>}
 */
proto.infra.ListNetworkInterfacesRequest.prototype.getLabelsMap = function(opt_noLazyCreate) {
  return /** @type {!jspb.Map<string,string>} */ (
      jspb.Message.getMapField(this, 5, opt_noLazyCreate,
      null));
};


/**
 * Clears values from the map. The map will be non-null.
 * @return {!proto.infra.ListNetworkInterfacesRequest} returns this
 */
proto.infra.ListNetworkInterfacesRequest.prototype.clearLabelsMap = function() {
  this.getLabelsMap().clear();
  return this;
};


/**
 * optional Credentials creds = 6;
 * @return {?proto.infra.Credentials}
 */
proto.infra.ListNetworkInterfacesRequest.prototype.getCreds = function() {
  return /** @type{?proto.infra.Credentials} */ (
    jspb.Message.getWrapperField(this, types_pb.Credentials, 6));
};


/**
 * @param {?proto.infra.Credentials|undefined} value
 * @return {!proto.infra.ListNetworkInterfacesRequest} returns this
*/
proto.infra.ListNetworkInterfacesRequest.prototype.setCreds = function(value) {
  return jspb.Message.setWrapperField(this, 6, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.infra.ListNetworkInterfacesRequest} returns this
 */
proto.infra.ListNetworkInterfacesRequest.prototype.clearCreds = function() {
  return this.setCreds(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.infra.ListNetworkInterfacesRequest.prototype.hasCreds = function() {
  return jspb.Message.getField(this, 6) != null;
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.infra.ListNetworkInterfacesResponse.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.infra.ListNetworkInterfacesResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.infra.ListNetworkInterfacesResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.infra.ListNetworkInterfacesResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListNetworkInterfacesResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    networkInterfacesList: jspb.Message.toObjectList(msg.getNetworkInterfacesList(),
    types_pb.NetworkInterface.toObject, includeInstance),
    lastSyncTime: jspb.Message.getFieldWithDefault(msg, 2, ""),
    err: (f = msg.getErr()) && types_pb.Error.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.infra.ListNetworkInterfacesResponse}
 */
proto.infra.ListNetworkInterfacesResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.infra.ListNetworkInterfacesResponse;
  return proto.infra.ListNetworkInterfacesResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.infra.ListNetworkInterfacesResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.infra.ListNetworkInterfacesResponse}
 */
proto.infra.ListNetworkInterfacesResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new types_pb.NetworkInterface;
      reader.readMessage(value,types_pb.NetworkInterface.deserializeBinaryFromReader);
      msg.addNetworkInterfaces(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setLastSyncTime(value);
      break;
    case 3:
      var value = new types_pb.Error;
      reader.readMessage(value,types_pb.Error.deserializeBinaryFromReader);
      msg.setErr(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.infra.ListNetworkInterfacesResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.infra.ListNetworkInterfacesResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.infra.ListNetworkInterfacesResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListNetworkInterfacesResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getNetworkInterfacesList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      types_pb.NetworkInterface.serializeBinaryToWriter
    );
  }
  f = message.getLastSyncTime();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getErr();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      types_pb.Error.serializeBinaryToWriter
    );
  }
};


/**
 * repeated NetworkInterface network_interfaces = 1;
 * @return {!Array<!proto.infra.NetworkInterface>}
 */
proto.infra.ListNetworkInterfacesResponse.prototype.getNetworkInterfacesList = function() {
  return /** @type{!Array<!proto.infra.NetworkInterface>} */ (
    jspb.Message.getRepeatedWrapperField(this, types_pb.NetworkInterface, 1));
};


/**
 * @param {!Array<!proto.infra.NetworkInterface>} value
 * @return {!proto.infra.ListNetworkInterfacesResponse} returns this
*/
proto.infra.ListNetworkInterfacesResponse.prototype.setNetworkInterfacesList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.infra.NetworkInterface=} opt_value
 * @param {number=} opt_index
 * @return {!proto.infra.NetworkInterface}
 */
proto.infra.ListNetworkInterfacesResponse.prototype.addNetworkInterfaces = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.infra.NetworkInterface, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.infra.ListNetworkInterfacesResponse} returns this
 */
proto.infra.ListNetworkInterfacesResponse.prototype.clearNetworkInterfacesList = function() {
  return this.setNetworkInterfacesList([]);
};


/**
 * optional string last_sync_time = 2;
 * @return {string}
 */
proto.infra.ListNetworkInterfacesResponse.prototype.getLastSyncTime = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListNetworkInterfacesResponse} returns this
 */
proto.infra.ListNetworkInterfacesResponse.prototype.setLastSyncTime = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional Error err = 3;
 * @return {?proto.infra.Error}
 */
proto.infra.ListNetworkInterfacesResponse.prototype.getErr = function() {
  return /** @type{?proto.infra.Error} */ (
    jspb.Message.getWrapperField(this, types_pb.Error, 3));
};


/**
 * @param {?proto.infra.Error|undefined} value
 * @return {!proto.infra.ListNetworkInterfacesResponse} returns this
*/
proto.infra.ListNetworkInterfacesResponse.prototype.setErr = function(value) {
  return jspb.Message.setWrapperField(this, 3, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.infra.ListNetworkInterfacesResponse} returns this
 */
proto.infra.ListNetworkInterfacesResponse.prototype.clearErr = function() {
  return this.setErr(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.infra.ListNetworkInterfacesResponse.prototype.hasErr = function() {
  return jspb.Message.getField(this, 3) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.infra.ListKeyPairsRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.infra.ListKeyPairsRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.infra.ListKeyPairsRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListKeyPairsRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    provider: jspb.Message.getFieldWithDefault(msg, 1, ""),
    accountId: jspb.Message.getFieldWithDefault(msg, 2, ""),
    region: jspb.Message.getFieldWithDefault(msg, 3, ""),
    vpcId: jspb.Message.getFieldWithDefault(msg, 4, ""),
    labelsMap: (f = msg.getLabelsMap()) ? f.toObject(includeInstance, undefined) : [],
    creds: (f = msg.getCreds()) && types_pb.Credentials.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.infra.ListKeyPairsRequest}
 */
proto.infra.ListKeyPairsRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.infra.ListKeyPairsRequest;
  return proto.infra.ListKeyPairsRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.infra.ListKeyPairsRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.infra.ListKeyPairsRequest}
 */
proto.infra.ListKeyPairsRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setProvider(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setAccountId(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setRegion(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setVpcId(value);
      break;
    case 5:
      var value = msg.getLabelsMap();
      reader.readMessage(value, function(message, reader) {
        jspb.Map.deserializeBinary(message, reader, jspb.BinaryReader.prototype.readString, jspb.BinaryReader.prototype.readString, null, "", "");
         });
      break;
    case 6:
      var value = new types_pb.Credentials;
      reader.readMessage(value,types_pb.Credentials.deserializeBinaryFromReader);
      msg.setCreds(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.infra.ListKeyPairsRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.infra.ListKeyPairsRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.infra.ListKeyPairsRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListKeyPairsRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getProvider();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getAccountId();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getRegion();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getVpcId();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getLabelsMap(true);
  if (f && f.getLength() > 0) {
    f.serializeBinary(5, writer, jspb.BinaryWriter.prototype.writeString, jspb.BinaryWriter.prototype.writeString);
  }
  f = message.getCreds();
  if (f != null) {
    writer.writeMessage(
      6,
      f,
      types_pb.Credentials.serializeBinaryToWriter
    );
  }
};


/**
 * optional string provider = 1;
 * @return {string}
 */
proto.infra.ListKeyPairsRequest.prototype.getProvider = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListKeyPairsRequest} returns this
 */
proto.infra.ListKeyPairsRequest.prototype.setProvider = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string account_id = 2;
 * @return {string}
 */
proto.infra.ListKeyPairsRequest.prototype.getAccountId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListKeyPairsRequest} returns this
 */
proto.infra.ListKeyPairsRequest.prototype.setAccountId = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string region = 3;
 * @return {string}
 */
proto.infra.ListKeyPairsRequest.prototype.getRegion = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListKeyPairsRequest} returns this
 */
proto.infra.ListKeyPairsRequest.prototype.setRegion = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string vpc_id = 4;
 * @return {string}
 */
proto.infra.ListKeyPairsRequest.prototype.getVpcId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListKeyPairsRequest} returns this
 */
proto.infra.ListKeyPairsRequest.prototype.setVpcId = function(value) {
  return jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * map<string, string> labels = 5;
 * @param {boolean=} opt_noLazyCreate Do not create the map if
 * empty, instead returning `undefined`
 * @return {!jspb.Map<string,string>}
 */
proto.infra.ListKeyPairsRequest.prototype.getLabelsMap = function(opt_noLazyCreate) {
  return /** @type {!jspb.Map<string,string>} */ (
      jspb.Message.getMapField(this, 5, opt_noLazyCreate,
      null));
};


/**
 * Clears values from the map. The map will be non-null.
 * @return {!proto.infra.ListKeyPairsRequest} returns this
 */
proto.infra.ListKeyPairsRequest.prototype.clearLabelsMap = function() {
  this.getLabelsMap().clear();
  return this;
};


/**
 * optional Credentials creds = 6;
 * @return {?proto.infra.Credentials}
 */
proto.infra.ListKeyPairsRequest.prototype.getCreds = function() {
  return /** @type{?proto.infra.Credentials} */ (
    jspb.Message.getWrapperField(this, types_pb.Credentials, 6));
};


/**
 * @param {?proto.infra.Credentials|undefined} value
 * @return {!proto.infra.ListKeyPairsRequest} returns this
*/
proto.infra.ListKeyPairsRequest.prototype.setCreds = function(value) {
  return jspb.Message.setWrapperField(this, 6, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.infra.ListKeyPairsRequest} returns this
 */
proto.infra.ListKeyPairsRequest.prototype.clearCreds = function() {
  return this.setCreds(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.infra.ListKeyPairsRequest.prototype.hasCreds = function() {
  return jspb.Message.getField(this, 6) != null;
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.infra.ListKeyPairsResponse.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.infra.ListKeyPairsResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.infra.ListKeyPairsResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.infra.ListKeyPairsResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListKeyPairsResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    keyPairsList: jspb.Message.toObjectList(msg.getKeyPairsList(),
    types_pb.KeyPair.toObject, includeInstance),
    lastSyncTime: jspb.Message.getFieldWithDefault(msg, 2, ""),
    err: (f = msg.getErr()) && types_pb.Error.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.infra.ListKeyPairsResponse}
 */
proto.infra.ListKeyPairsResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.infra.ListKeyPairsResponse;
  return proto.infra.ListKeyPairsResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.infra.ListKeyPairsResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.infra.ListKeyPairsResponse}
 */
proto.infra.ListKeyPairsResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new types_pb.KeyPair;
      reader.readMessage(value,types_pb.KeyPair.deserializeBinaryFromReader);
      msg.addKeyPairs(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setLastSyncTime(value);
      break;
    case 3:
      var value = new types_pb.Error;
      reader.readMessage(value,types_pb.Error.deserializeBinaryFromReader);
      msg.setErr(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.infra.ListKeyPairsResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.infra.ListKeyPairsResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.infra.ListKeyPairsResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListKeyPairsResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getKeyPairsList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      types_pb.KeyPair.serializeBinaryToWriter
    );
  }
  f = message.getLastSyncTime();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getErr();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      types_pb.Error.serializeBinaryToWriter
    );
  }
};


/**
 * repeated KeyPair key_pairs = 1;
 * @return {!Array<!proto.infra.KeyPair>}
 */
proto.infra.ListKeyPairsResponse.prototype.getKeyPairsList = function() {
  return /** @type{!Array<!proto.infra.KeyPair>} */ (
    jspb.Message.getRepeatedWrapperField(this, types_pb.KeyPair, 1));
};


/**
 * @param {!Array<!proto.infra.KeyPair>} value
 * @return {!proto.infra.ListKeyPairsResponse} returns this
*/
proto.infra.ListKeyPairsResponse.prototype.setKeyPairsList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.infra.KeyPair=} opt_value
 * @param {number=} opt_index
 * @return {!proto.infra.KeyPair}
 */
proto.infra.ListKeyPairsResponse.prototype.addKeyPairs = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.infra.KeyPair, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.infra.ListKeyPairsResponse} returns this
 */
proto.infra.ListKeyPairsResponse.prototype.clearKeyPairsList = function() {
  return this.setKeyPairsList([]);
};


/**
 * optional string last_sync_time = 2;
 * @return {string}
 */
proto.infra.ListKeyPairsResponse.prototype.getLastSyncTime = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListKeyPairsResponse} returns this
 */
proto.infra.ListKeyPairsResponse.prototype.setLastSyncTime = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional Error err = 3;
 * @return {?proto.infra.Error}
 */
proto.infra.ListKeyPairsResponse.prototype.getErr = function() {
  return /** @type{?proto.infra.Error} */ (
    jspb.Message.getWrapperField(this, types_pb.Error, 3));
};


/**
 * @param {?proto.infra.Error|undefined} value
 * @return {!proto.infra.ListKeyPairsResponse} returns this
*/
proto.infra.ListKeyPairsResponse.prototype.setErr = function(value) {
  return jspb.Message.setWrapperField(this, 3, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.infra.ListKeyPairsResponse} returns this
 */
proto.infra.ListKeyPairsResponse.prototype.clearErr = function() {
  return this.setErr(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.infra.ListKeyPairsResponse.prototype.hasErr = function() {
  return jspb.Message.getField(this, 3) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.infra.ListVPNConcentratorsRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.infra.ListVPNConcentratorsRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.infra.ListVPNConcentratorsRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListVPNConcentratorsRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    provider: jspb.Message.getFieldWithDefault(msg, 1, ""),
    accountId: jspb.Message.getFieldWithDefault(msg, 2, ""),
    region: jspb.Message.getFieldWithDefault(msg, 3, ""),
    labelsMap: (f = msg.getLabelsMap()) ? f.toObject(includeInstance, undefined) : [],
    creds: (f = msg.getCreds()) && types_pb.Credentials.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.infra.ListVPNConcentratorsRequest}
 */
proto.infra.ListVPNConcentratorsRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.infra.ListVPNConcentratorsRequest;
  return proto.infra.ListVPNConcentratorsRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.infra.ListVPNConcentratorsRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.infra.ListVPNConcentratorsRequest}
 */
proto.infra.ListVPNConcentratorsRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setProvider(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setAccountId(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setRegion(value);
      break;
    case 4:
      var value = msg.getLabelsMap();
      reader.readMessage(value, function(message, reader) {
        jspb.Map.deserializeBinary(message, reader, jspb.BinaryReader.prototype.readString, jspb.BinaryReader.prototype.readString, null, "", "");
         });
      break;
    case 5:
      var value = new types_pb.Credentials;
      reader.readMessage(value,types_pb.Credentials.deserializeBinaryFromReader);
      msg.setCreds(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.infra.ListVPNConcentratorsRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.infra.ListVPNConcentratorsRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.infra.ListVPNConcentratorsRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListVPNConcentratorsRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getProvider();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getAccountId();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getRegion();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getLabelsMap(true);
  if (f && f.getLength() > 0) {
    f.serializeBinary(4, writer, jspb.BinaryWriter.prototype.writeString, jspb.BinaryWriter.prototype.writeString);
  }
  f = message.getCreds();
  if (f != null) {
    writer.writeMessage(
      5,
      f,
      types_pb.Credentials.serializeBinaryToWriter
    );
  }
};


/**
 * optional string provider = 1;
 * @return {string}
 */
proto.infra.ListVPNConcentratorsRequest.prototype.getProvider = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListVPNConcentratorsRequest} returns this
 */
proto.infra.ListVPNConcentratorsRequest.prototype.setProvider = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string account_id = 2;
 * @return {string}
 */
proto.infra.ListVPNConcentratorsRequest.prototype.getAccountId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListVPNConcentratorsRequest} returns this
 */
proto.infra.ListVPNConcentratorsRequest.prototype.setAccountId = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string region = 3;
 * @return {string}
 */
proto.infra.ListVPNConcentratorsRequest.prototype.getRegion = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListVPNConcentratorsRequest} returns this
 */
proto.infra.ListVPNConcentratorsRequest.prototype.setRegion = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * map<string, string> labels = 4;
 * @param {boolean=} opt_noLazyCreate Do not create the map if
 * empty, instead returning `undefined`
 * @return {!jspb.Map<string,string>}
 */
proto.infra.ListVPNConcentratorsRequest.prototype.getLabelsMap = function(opt_noLazyCreate) {
  return /** @type {!jspb.Map<string,string>} */ (
      jspb.Message.getMapField(this, 4, opt_noLazyCreate,
      null));
};


/**
 * Clears values from the map. The map will be non-null.
 * @return {!proto.infra.ListVPNConcentratorsRequest} returns this
 */
proto.infra.ListVPNConcentratorsRequest.prototype.clearLabelsMap = function() {
  this.getLabelsMap().clear();
  return this;
};


/**
 * optional Credentials creds = 5;
 * @return {?proto.infra.Credentials}
 */
proto.infra.ListVPNConcentratorsRequest.prototype.getCreds = function() {
  return /** @type{?proto.infra.Credentials} */ (
    jspb.Message.getWrapperField(this, types_pb.Credentials, 5));
};


/**
 * @param {?proto.infra.Credentials|undefined} value
 * @return {!proto.infra.ListVPNConcentratorsRequest} returns this
*/
proto.infra.ListVPNConcentratorsRequest.prototype.setCreds = function(value) {
  return jspb.Message.setWrapperField(this, 5, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.infra.ListVPNConcentratorsRequest} returns this
 */
proto.infra.ListVPNConcentratorsRequest.prototype.clearCreds = function() {
  return this.setCreds(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.infra.ListVPNConcentratorsRequest.prototype.hasCreds = function() {
  return jspb.Message.getField(this, 5) != null;
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.infra.ListVPNConcentratorsResponse.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.infra.ListVPNConcentratorsResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.infra.ListVPNConcentratorsResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.infra.ListVPNConcentratorsResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListVPNConcentratorsResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    vpnConcentratorsList: jspb.Message.toObjectList(msg.getVpnConcentratorsList(),
    types_pb.VPNConcentrator.toObject, includeInstance),
    lastSyncTime: jspb.Message.getFieldWithDefault(msg, 2, ""),
    err: (f = msg.getErr()) && types_pb.Error.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.infra.ListVPNConcentratorsResponse}
 */
proto.infra.ListVPNConcentratorsResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.infra.ListVPNConcentratorsResponse;
  return proto.infra.ListVPNConcentratorsResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.infra.ListVPNConcentratorsResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.infra.ListVPNConcentratorsResponse}
 */
proto.infra.ListVPNConcentratorsResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new types_pb.VPNConcentrator;
      reader.readMessage(value,types_pb.VPNConcentrator.deserializeBinaryFromReader);
      msg.addVpnConcentrators(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setLastSyncTime(value);
      break;
    case 3:
      var value = new types_pb.Error;
      reader.readMessage(value,types_pb.Error.deserializeBinaryFromReader);
      msg.setErr(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.infra.ListVPNConcentratorsResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.infra.ListVPNConcentratorsResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.infra.ListVPNConcentratorsResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListVPNConcentratorsResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getVpnConcentratorsList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      types_pb.VPNConcentrator.serializeBinaryToWriter
    );
  }
  f = message.getLastSyncTime();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getErr();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      types_pb.Error.serializeBinaryToWriter
    );
  }
};


/**
 * repeated VPNConcentrator vpn_concentrators = 1;
 * @return {!Array<!proto.infra.VPNConcentrator>}
 */
proto.infra.ListVPNConcentratorsResponse.prototype.getVpnConcentratorsList = function() {
  return /** @type{!Array<!proto.infra.VPNConcentrator>} */ (
    jspb.Message.getRepeatedWrapperField(this, types_pb.VPNConcentrator, 1));
};


/**
 * @param {!Array<!proto.infra.VPNConcentrator>} value
 * @return {!proto.infra.ListVPNConcentratorsResponse} returns this
*/
proto.infra.ListVPNConcentratorsResponse.prototype.setVpnConcentratorsList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.infra.VPNConcentrator=} opt_value
 * @param {number=} opt_index
 * @return {!proto.infra.VPNConcentrator}
 */
proto.infra.ListVPNConcentratorsResponse.prototype.addVpnConcentrators = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.infra.VPNConcentrator, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.infra.ListVPNConcentratorsResponse} returns this
 */
proto.infra.ListVPNConcentratorsResponse.prototype.clearVpnConcentratorsList = function() {
  return this.setVpnConcentratorsList([]);
};


/**
 * optional string last_sync_time = 2;
 * @return {string}
 */
proto.infra.ListVPNConcentratorsResponse.prototype.getLastSyncTime = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListVPNConcentratorsResponse} returns this
 */
proto.infra.ListVPNConcentratorsResponse.prototype.setLastSyncTime = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional Error err = 3;
 * @return {?proto.infra.Error}
 */
proto.infra.ListVPNConcentratorsResponse.prototype.getErr = function() {
  return /** @type{?proto.infra.Error} */ (
    jspb.Message.getWrapperField(this, types_pb.Error, 3));
};


/**
 * @param {?proto.infra.Error|undefined} value
 * @return {!proto.infra.ListVPNConcentratorsResponse} returns this
*/
proto.infra.ListVPNConcentratorsResponse.prototype.setErr = function(value) {
  return jspb.Message.setWrapperField(this, 3, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.infra.ListVPNConcentratorsResponse} returns this
 */
proto.infra.ListVPNConcentratorsResponse.prototype.clearErr = function() {
  return this.setErr(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.infra.ListVPNConcentratorsResponse.prototype.hasErr = function() {
  return jspb.Message.getField(this, 3) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.infra.GetVPCIDForCIDRRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.infra.GetVPCIDForCIDRRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.infra.GetVPCIDForCIDRRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.GetVPCIDForCIDRRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    provider: jspb.Message.getFieldWithDefault(msg, 1, ""),
    accountId: jspb.Message.getFieldWithDefault(msg, 2, ""),
    region: jspb.Message.getFieldWithDefault(msg, 3, ""),
    cidr: jspb.Message.getFieldWithDefault(msg, 4, ""),
    labelsMap: (f = msg.getLabelsMap()) ? f.toObject(includeInstance, undefined) : [],
    creds: (f = msg.getCreds()) && types_pb.Credentials.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.infra.GetVPCIDForCIDRRequest}
 */
proto.infra.GetVPCIDForCIDRRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.infra.GetVPCIDForCIDRRequest;
  return proto.infra.GetVPCIDForCIDRRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.infra.GetVPCIDForCIDRRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.infra.GetVPCIDForCIDRRequest}
 */
proto.infra.GetVPCIDForCIDRRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setProvider(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setAccountId(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setRegion(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setCidr(value);
      break;
    case 5:
      var value = msg.getLabelsMap();
      reader.readMessage(value, function(message, reader) {
        jspb.Map.deserializeBinary(message, reader, jspb.BinaryReader.prototype.readString, jspb.BinaryReader.prototype.readString, null, "", "");
         });
      break;
    case 6:
      var value = new types_pb.Credentials;
      reader.readMessage(value,types_pb.Credentials.deserializeBinaryFromReader);
      msg.setCreds(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.infra.GetVPCIDForCIDRRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.infra.GetVPCIDForCIDRRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.infra.GetVPCIDForCIDRRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.GetVPCIDForCIDRRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getProvider();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getAccountId();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getRegion();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getCidr();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getLabelsMap(true);
  if (f && f.getLength() > 0) {
    f.serializeBinary(5, writer, jspb.BinaryWriter.prototype.writeString, jspb.BinaryWriter.prototype.writeString);
  }
  f = message.getCreds();
  if (f != null) {
    writer.writeMessage(
      6,
      f,
      types_pb.Credentials.serializeBinaryToWriter
    );
  }
};


/**
 * optional string provider = 1;
 * @return {string}
 */
proto.infra.GetVPCIDForCIDRRequest.prototype.getProvider = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.GetVPCIDForCIDRRequest} returns this
 */
proto.infra.GetVPCIDForCIDRRequest.prototype.setProvider = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string account_id = 2;
 * @return {string}
 */
proto.infra.GetVPCIDForCIDRRequest.prototype.getAccountId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.GetVPCIDForCIDRRequest} returns this
 */
proto.infra.GetVPCIDForCIDRRequest.prototype.setAccountId = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string region = 3;
 * @return {string}
 */
proto.infra.GetVPCIDForCIDRRequest.prototype.getRegion = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.GetVPCIDForCIDRRequest} returns this
 */
proto.infra.GetVPCIDForCIDRRequest.prototype.setRegion = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string cidr = 4;
 * @return {string}
 */
proto.infra.GetVPCIDForCIDRRequest.prototype.getCidr = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.GetVPCIDForCIDRRequest} returns this
 */
proto.infra.GetVPCIDForCIDRRequest.prototype.setCidr = function(value) {
  return jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * map<string, string> labels = 5;
 * @param {boolean=} opt_noLazyCreate Do not create the map if
 * empty, instead returning `undefined`
 * @return {!jspb.Map<string,string>}
 */
proto.infra.GetVPCIDForCIDRRequest.prototype.getLabelsMap = function(opt_noLazyCreate) {
  return /** @type {!jspb.Map<string,string>} */ (
      jspb.Message.getMapField(this, 5, opt_noLazyCreate,
      null));
};


/**
 * Clears values from the map. The map will be non-null.
 * @return {!proto.infra.GetVPCIDForCIDRRequest} returns this
 */
proto.infra.GetVPCIDForCIDRRequest.prototype.clearLabelsMap = function() {
  this.getLabelsMap().clear();
  return this;
};


/**
 * optional Credentials creds = 6;
 * @return {?proto.infra.Credentials}
 */
proto.infra.GetVPCIDForCIDRRequest.prototype.getCreds = function() {
  return /** @type{?proto.infra.Credentials} */ (
    jspb.Message.getWrapperField(this, types_pb.Credentials, 6));
};


/**
 * @param {?proto.infra.Credentials|undefined} value
 * @return {!proto.infra.GetVPCIDForCIDRRequest} returns this
*/
proto.infra.GetVPCIDForCIDRRequest.prototype.setCreds = function(value) {
  return jspb.Message.setWrapperField(this, 6, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.infra.GetVPCIDForCIDRRequest} returns this
 */
proto.infra.GetVPCIDForCIDRRequest.prototype.clearCreds = function() {
  return this.setCreds(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.infra.GetVPCIDForCIDRRequest.prototype.hasCreds = function() {
  return jspb.Message.getField(this, 6) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.infra.GetVPCIDForCIDRResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.infra.GetVPCIDForCIDRResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.infra.GetVPCIDForCIDRResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.GetVPCIDForCIDRResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    vpcId: jspb.Message.getFieldWithDefault(msg, 1, ""),
    err: (f = msg.getErr()) && types_pb.Error.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.infra.GetVPCIDForCIDRResponse}
 */
proto.infra.GetVPCIDForCIDRResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.infra.GetVPCIDForCIDRResponse;
  return proto.infra.GetVPCIDForCIDRResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.infra.GetVPCIDForCIDRResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.infra.GetVPCIDForCIDRResponse}
 */
proto.infra.GetVPCIDForCIDRResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setVpcId(value);
      break;
    case 3:
      var value = new types_pb.Error;
      reader.readMessage(value,types_pb.Error.deserializeBinaryFromReader);
      msg.setErr(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.infra.GetVPCIDForCIDRResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.infra.GetVPCIDForCIDRResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.infra.GetVPCIDForCIDRResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.GetVPCIDForCIDRResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getVpcId();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getErr();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      types_pb.Error.serializeBinaryToWriter
    );
  }
};


/**
 * optional string vpc_id = 1;
 * @return {string}
 */
proto.infra.GetVPCIDForCIDRResponse.prototype.getVpcId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.GetVPCIDForCIDRResponse} returns this
 */
proto.infra.GetVPCIDForCIDRResponse.prototype.setVpcId = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional Error err = 3;
 * @return {?proto.infra.Error}
 */
proto.infra.GetVPCIDForCIDRResponse.prototype.getErr = function() {
  return /** @type{?proto.infra.Error} */ (
    jspb.Message.getWrapperField(this, types_pb.Error, 3));
};


/**
 * @param {?proto.infra.Error|undefined} value
 * @return {!proto.infra.GetVPCIDForCIDRResponse} returns this
*/
proto.infra.GetVPCIDForCIDRResponse.prototype.setErr = function(value) {
  return jspb.Message.setWrapperField(this, 3, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.infra.GetVPCIDForCIDRResponse} returns this
 */
proto.infra.GetVPCIDForCIDRResponse.prototype.clearErr = function() {
  return this.setErr(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.infra.GetVPCIDForCIDRResponse.prototype.hasErr = function() {
  return jspb.Message.getField(this, 3) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.infra.GetCIDRsForLabelsRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.infra.GetCIDRsForLabelsRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.infra.GetCIDRsForLabelsRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.GetCIDRsForLabelsRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    provider: jspb.Message.getFieldWithDefault(msg, 1, ""),
    accountId: jspb.Message.getFieldWithDefault(msg, 2, ""),
    region: jspb.Message.getFieldWithDefault(msg, 3, ""),
    labelsMap: (f = msg.getLabelsMap()) ? f.toObject(includeInstance, undefined) : [],
    creds: (f = msg.getCreds()) && types_pb.Credentials.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.infra.GetCIDRsForLabelsRequest}
 */
proto.infra.GetCIDRsForLabelsRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.infra.GetCIDRsForLabelsRequest;
  return proto.infra.GetCIDRsForLabelsRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.infra.GetCIDRsForLabelsRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.infra.GetCIDRsForLabelsRequest}
 */
proto.infra.GetCIDRsForLabelsRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setProvider(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setAccountId(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setRegion(value);
      break;
    case 4:
      var value = msg.getLabelsMap();
      reader.readMessage(value, function(message, reader) {
        jspb.Map.deserializeBinary(message, reader, jspb.BinaryReader.prototype.readString, jspb.BinaryReader.prototype.readString, null, "", "");
         });
      break;
    case 5:
      var value = new types_pb.Credentials;
      reader.readMessage(value,types_pb.Credentials.deserializeBinaryFromReader);
      msg.setCreds(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.infra.GetCIDRsForLabelsRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.infra.GetCIDRsForLabelsRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.infra.GetCIDRsForLabelsRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.GetCIDRsForLabelsRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getProvider();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getAccountId();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getRegion();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getLabelsMap(true);
  if (f && f.getLength() > 0) {
    f.serializeBinary(4, writer, jspb.BinaryWriter.prototype.writeString, jspb.BinaryWriter.prototype.writeString);
  }
  f = message.getCreds();
  if (f != null) {
    writer.writeMessage(
      5,
      f,
      types_pb.Credentials.serializeBinaryToWriter
    );
  }
};


/**
 * optional string provider = 1;
 * @return {string}
 */
proto.infra.GetCIDRsForLabelsRequest.prototype.getProvider = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.GetCIDRsForLabelsRequest} returns this
 */
proto.infra.GetCIDRsForLabelsRequest.prototype.setProvider = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string account_id = 2;
 * @return {string}
 */
proto.infra.GetCIDRsForLabelsRequest.prototype.getAccountId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.GetCIDRsForLabelsRequest} returns this
 */
proto.infra.GetCIDRsForLabelsRequest.prototype.setAccountId = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string region = 3;
 * @return {string}
 */
proto.infra.GetCIDRsForLabelsRequest.prototype.getRegion = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.GetCIDRsForLabelsRequest} returns this
 */
proto.infra.GetCIDRsForLabelsRequest.prototype.setRegion = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * map<string, string> labels = 4;
 * @param {boolean=} opt_noLazyCreate Do not create the map if
 * empty, instead returning `undefined`
 * @return {!jspb.Map<string,string>}
 */
proto.infra.GetCIDRsForLabelsRequest.prototype.getLabelsMap = function(opt_noLazyCreate) {
  return /** @type {!jspb.Map<string,string>} */ (
      jspb.Message.getMapField(this, 4, opt_noLazyCreate,
      null));
};


/**
 * Clears values from the map. The map will be non-null.
 * @return {!proto.infra.GetCIDRsForLabelsRequest} returns this
 */
proto.infra.GetCIDRsForLabelsRequest.prototype.clearLabelsMap = function() {
  this.getLabelsMap().clear();
  return this;
};


/**
 * optional Credentials creds = 5;
 * @return {?proto.infra.Credentials}
 */
proto.infra.GetCIDRsForLabelsRequest.prototype.getCreds = function() {
  return /** @type{?proto.infra.Credentials} */ (
    jspb.Message.getWrapperField(this, types_pb.Credentials, 5));
};


/**
 * @param {?proto.infra.Credentials|undefined} value
 * @return {!proto.infra.GetCIDRsForLabelsRequest} returns this
*/
proto.infra.GetCIDRsForLabelsRequest.prototype.setCreds = function(value) {
  return jspb.Message.setWrapperField(this, 5, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.infra.GetCIDRsForLabelsRequest} returns this
 */
proto.infra.GetCIDRsForLabelsRequest.prototype.clearCreds = function() {
  return this.setCreds(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.infra.GetCIDRsForLabelsRequest.prototype.hasCreds = function() {
  return jspb.Message.getField(this, 5) != null;
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.infra.GetCIDRsForLabelsResponse.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.infra.GetCIDRsForLabelsResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.infra.GetCIDRsForLabelsResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.infra.GetCIDRsForLabelsResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.GetCIDRsForLabelsResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    cidrsList: (f = jspb.Message.getRepeatedField(msg, 1)) == null ? undefined : f,
    err: (f = msg.getErr()) && types_pb.Error.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.infra.GetCIDRsForLabelsResponse}
 */
proto.infra.GetCIDRsForLabelsResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.infra.GetCIDRsForLabelsResponse;
  return proto.infra.GetCIDRsForLabelsResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.infra.GetCIDRsForLabelsResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.infra.GetCIDRsForLabelsResponse}
 */
proto.infra.GetCIDRsForLabelsResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.addCidrs(value);
      break;
    case 3:
      var value = new types_pb.Error;
      reader.readMessage(value,types_pb.Error.deserializeBinaryFromReader);
      msg.setErr(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.infra.GetCIDRsForLabelsResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.infra.GetCIDRsForLabelsResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.infra.GetCIDRsForLabelsResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.GetCIDRsForLabelsResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getCidrsList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      1,
      f
    );
  }
  f = message.getErr();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      types_pb.Error.serializeBinaryToWriter
    );
  }
};


/**
 * repeated string cidrs = 1;
 * @return {!Array<string>}
 */
proto.infra.GetCIDRsForLabelsResponse.prototype.getCidrsList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 1));
};


/**
 * @param {!Array<string>} value
 * @return {!proto.infra.GetCIDRsForLabelsResponse} returns this
 */
proto.infra.GetCIDRsForLabelsResponse.prototype.setCidrsList = function(value) {
  return jspb.Message.setField(this, 1, value || []);
};


/**
 * @param {string} value
 * @param {number=} opt_index
 * @return {!proto.infra.GetCIDRsForLabelsResponse} returns this
 */
proto.infra.GetCIDRsForLabelsResponse.prototype.addCidrs = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 1, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.infra.GetCIDRsForLabelsResponse} returns this
 */
proto.infra.GetCIDRsForLabelsResponse.prototype.clearCidrsList = function() {
  return this.setCidrsList([]);
};


/**
 * optional Error err = 3;
 * @return {?proto.infra.Error}
 */
proto.infra.GetCIDRsForLabelsResponse.prototype.getErr = function() {
  return /** @type{?proto.infra.Error} */ (
    jspb.Message.getWrapperField(this, types_pb.Error, 3));
};


/**
 * @param {?proto.infra.Error|undefined} value
 * @return {!proto.infra.GetCIDRsForLabelsResponse} returns this
*/
proto.infra.GetCIDRsForLabelsResponse.prototype.setErr = function(value) {
  return jspb.Message.setWrapperField(this, 3, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.infra.GetCIDRsForLabelsResponse} returns this
 */
proto.infra.GetCIDRsForLabelsResponse.prototype.clearErr = function() {
  return this.setErr(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.infra.GetCIDRsForLabelsResponse.prototype.hasErr = function() {
  return jspb.Message.getField(this, 3) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.infra.GetIPsForLabelsRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.infra.GetIPsForLabelsRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.infra.GetIPsForLabelsRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.GetIPsForLabelsRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    provider: jspb.Message.getFieldWithDefault(msg, 1, ""),
    accountId: jspb.Message.getFieldWithDefault(msg, 2, ""),
    region: jspb.Message.getFieldWithDefault(msg, 3, ""),
    labelsMap: (f = msg.getLabelsMap()) ? f.toObject(includeInstance, undefined) : [],
    creds: (f = msg.getCreds()) && types_pb.Credentials.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.infra.GetIPsForLabelsRequest}
 */
proto.infra.GetIPsForLabelsRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.infra.GetIPsForLabelsRequest;
  return proto.infra.GetIPsForLabelsRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.infra.GetIPsForLabelsRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.infra.GetIPsForLabelsRequest}
 */
proto.infra.GetIPsForLabelsRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setProvider(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setAccountId(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setRegion(value);
      break;
    case 4:
      var value = msg.getLabelsMap();
      reader.readMessage(value, function(message, reader) {
        jspb.Map.deserializeBinary(message, reader, jspb.BinaryReader.prototype.readString, jspb.BinaryReader.prototype.readString, null, "", "");
         });
      break;
    case 5:
      var value = new types_pb.Credentials;
      reader.readMessage(value,types_pb.Credentials.deserializeBinaryFromReader);
      msg.setCreds(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.infra.GetIPsForLabelsRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.infra.GetIPsForLabelsRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.infra.GetIPsForLabelsRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.GetIPsForLabelsRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getProvider();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getAccountId();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getRegion();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getLabelsMap(true);
  if (f && f.getLength() > 0) {
    f.serializeBinary(4, writer, jspb.BinaryWriter.prototype.writeString, jspb.BinaryWriter.prototype.writeString);
  }
  f = message.getCreds();
  if (f != null) {
    writer.writeMessage(
      5,
      f,
      types_pb.Credentials.serializeBinaryToWriter
    );
  }
};


/**
 * optional string provider = 1;
 * @return {string}
 */
proto.infra.GetIPsForLabelsRequest.prototype.getProvider = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.GetIPsForLabelsRequest} returns this
 */
proto.infra.GetIPsForLabelsRequest.prototype.setProvider = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string account_id = 2;
 * @return {string}
 */
proto.infra.GetIPsForLabelsRequest.prototype.getAccountId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.GetIPsForLabelsRequest} returns this
 */
proto.infra.GetIPsForLabelsRequest.prototype.setAccountId = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string region = 3;
 * @return {string}
 */
proto.infra.GetIPsForLabelsRequest.prototype.getRegion = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.GetIPsForLabelsRequest} returns this
 */
proto.infra.GetIPsForLabelsRequest.prototype.setRegion = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * map<string, string> labels = 4;
 * @param {boolean=} opt_noLazyCreate Do not create the map if
 * empty, instead returning `undefined`
 * @return {!jspb.Map<string,string>}
 */
proto.infra.GetIPsForLabelsRequest.prototype.getLabelsMap = function(opt_noLazyCreate) {
  return /** @type {!jspb.Map<string,string>} */ (
      jspb.Message.getMapField(this, 4, opt_noLazyCreate,
      null));
};


/**
 * Clears values from the map. The map will be non-null.
 * @return {!proto.infra.GetIPsForLabelsRequest} returns this
 */
proto.infra.GetIPsForLabelsRequest.prototype.clearLabelsMap = function() {
  this.getLabelsMap().clear();
  return this;
};


/**
 * optional Credentials creds = 5;
 * @return {?proto.infra.Credentials}
 */
proto.infra.GetIPsForLabelsRequest.prototype.getCreds = function() {
  return /** @type{?proto.infra.Credentials} */ (
    jspb.Message.getWrapperField(this, types_pb.Credentials, 5));
};


/**
 * @param {?proto.infra.Credentials|undefined} value
 * @return {!proto.infra.GetIPsForLabelsRequest} returns this
*/
proto.infra.GetIPsForLabelsRequest.prototype.setCreds = function(value) {
  return jspb.Message.setWrapperField(this, 5, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.infra.GetIPsForLabelsRequest} returns this
 */
proto.infra.GetIPsForLabelsRequest.prototype.clearCreds = function() {
  return this.setCreds(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.infra.GetIPsForLabelsRequest.prototype.hasCreds = function() {
  return jspb.Message.getField(this, 5) != null;
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.infra.GetIPsForLabelsResponse.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.infra.GetIPsForLabelsResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.infra.GetIPsForLabelsResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.infra.GetIPsForLabelsResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.GetIPsForLabelsResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    ipsList: (f = jspb.Message.getRepeatedField(msg, 1)) == null ? undefined : f,
    err: (f = msg.getErr()) && types_pb.Error.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.infra.GetIPsForLabelsResponse}
 */
proto.infra.GetIPsForLabelsResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.infra.GetIPsForLabelsResponse;
  return proto.infra.GetIPsForLabelsResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.infra.GetIPsForLabelsResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.infra.GetIPsForLabelsResponse}
 */
proto.infra.GetIPsForLabelsResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.addIps(value);
      break;
    case 3:
      var value = new types_pb.Error;
      reader.readMessage(value,types_pb.Error.deserializeBinaryFromReader);
      msg.setErr(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.infra.GetIPsForLabelsResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.infra.GetIPsForLabelsResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.infra.GetIPsForLabelsResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.GetIPsForLabelsResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getIpsList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      1,
      f
    );
  }
  f = message.getErr();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      types_pb.Error.serializeBinaryToWriter
    );
  }
};


/**
 * repeated string ips = 1;
 * @return {!Array<string>}
 */
proto.infra.GetIPsForLabelsResponse.prototype.getIpsList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 1));
};


/**
 * @param {!Array<string>} value
 * @return {!proto.infra.GetIPsForLabelsResponse} returns this
 */
proto.infra.GetIPsForLabelsResponse.prototype.setIpsList = function(value) {
  return jspb.Message.setField(this, 1, value || []);
};


/**
 * @param {string} value
 * @param {number=} opt_index
 * @return {!proto.infra.GetIPsForLabelsResponse} returns this
 */
proto.infra.GetIPsForLabelsResponse.prototype.addIps = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 1, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.infra.GetIPsForLabelsResponse} returns this
 */
proto.infra.GetIPsForLabelsResponse.prototype.clearIpsList = function() {
  return this.setIpsList([]);
};


/**
 * optional Error err = 3;
 * @return {?proto.infra.Error}
 */
proto.infra.GetIPsForLabelsResponse.prototype.getErr = function() {
  return /** @type{?proto.infra.Error} */ (
    jspb.Message.getWrapperField(this, types_pb.Error, 3));
};


/**
 * @param {?proto.infra.Error|undefined} value
 * @return {!proto.infra.GetIPsForLabelsResponse} returns this
*/
proto.infra.GetIPsForLabelsResponse.prototype.setErr = function(value) {
  return jspb.Message.setWrapperField(this, 3, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.infra.GetIPsForLabelsResponse} returns this
 */
proto.infra.GetIPsForLabelsResponse.prototype.clearErr = function() {
  return this.setErr(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.infra.GetIPsForLabelsResponse.prototype.hasErr = function() {
  return jspb.Message.getField(this, 3) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.infra.GetInstancesForLabelsRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.infra.GetInstancesForLabelsRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.infra.GetInstancesForLabelsRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.GetInstancesForLabelsRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    provider: jspb.Message.getFieldWithDefault(msg, 1, ""),
    accountId: jspb.Message.getFieldWithDefault(msg, 2, ""),
    region: jspb.Message.getFieldWithDefault(msg, 3, ""),
    vpcId: jspb.Message.getFieldWithDefault(msg, 4, ""),
    labelsMap: (f = msg.getLabelsMap()) ? f.toObject(includeInstance, undefined) : [],
    creds: (f = msg.getCreds()) && types_pb.Credentials.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.infra.GetInstancesForLabelsRequest}
 */
proto.infra.GetInstancesForLabelsRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.infra.GetInstancesForLabelsRequest;
  return proto.infra.GetInstancesForLabelsRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.infra.GetInstancesForLabelsRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.infra.GetInstancesForLabelsRequest}
 */
proto.infra.GetInstancesForLabelsRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setProvider(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setAccountId(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setRegion(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setVpcId(value);
      break;
    case 5:
      var value = msg.getLabelsMap();
      reader.readMessage(value, function(message, reader) {
        jspb.Map.deserializeBinary(message, reader, jspb.BinaryReader.prototype.readString, jspb.BinaryReader.prototype.readString, null, "", "");
         });
      break;
    case 6:
      var value = new types_pb.Credentials;
      reader.readMessage(value,types_pb.Credentials.deserializeBinaryFromReader);
      msg.setCreds(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.infra.GetInstancesForLabelsRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.infra.GetInstancesForLabelsRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.infra.GetInstancesForLabelsRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.GetInstancesForLabelsRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getProvider();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getAccountId();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getRegion();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getVpcId();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getLabelsMap(true);
  if (f && f.getLength() > 0) {
    f.serializeBinary(5, writer, jspb.BinaryWriter.prototype.writeString, jspb.BinaryWriter.prototype.writeString);
  }
  f = message.getCreds();
  if (f != null) {
    writer.writeMessage(
      6,
      f,
      types_pb.Credentials.serializeBinaryToWriter
    );
  }
};


/**
 * optional string provider = 1;
 * @return {string}
 */
proto.infra.GetInstancesForLabelsRequest.prototype.getProvider = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.GetInstancesForLabelsRequest} returns this
 */
proto.infra.GetInstancesForLabelsRequest.prototype.setProvider = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string account_id = 2;
 * @return {string}
 */
proto.infra.GetInstancesForLabelsRequest.prototype.getAccountId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.GetInstancesForLabelsRequest} returns this
 */
proto.infra.GetInstancesForLabelsRequest.prototype.setAccountId = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string region = 3;
 * @return {string}
 */
proto.infra.GetInstancesForLabelsRequest.prototype.getRegion = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.GetInstancesForLabelsRequest} returns this
 */
proto.infra.GetInstancesForLabelsRequest.prototype.setRegion = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string vpc_id = 4;
 * @return {string}
 */
proto.infra.GetInstancesForLabelsRequest.prototype.getVpcId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.GetInstancesForLabelsRequest} returns this
 */
proto.infra.GetInstancesForLabelsRequest.prototype.setVpcId = function(value) {
  return jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * map<string, string> labels = 5;
 * @param {boolean=} opt_noLazyCreate Do not create the map if
 * empty, instead returning `undefined`
 * @return {!jspb.Map<string,string>}
 */
proto.infra.GetInstancesForLabelsRequest.prototype.getLabelsMap = function(opt_noLazyCreate) {
  return /** @type {!jspb.Map<string,string>} */ (
      jspb.Message.getMapField(this, 5, opt_noLazyCreate,
      null));
};


/**
 * Clears values from the map. The map will be non-null.
 * @return {!proto.infra.GetInstancesForLabelsRequest} returns this
 */
proto.infra.GetInstancesForLabelsRequest.prototype.clearLabelsMap = function() {
  this.getLabelsMap().clear();
  return this;
};


/**
 * optional Credentials creds = 6;
 * @return {?proto.infra.Credentials}
 */
proto.infra.GetInstancesForLabelsRequest.prototype.getCreds = function() {
  return /** @type{?proto.infra.Credentials} */ (
    jspb.Message.getWrapperField(this, types_pb.Credentials, 6));
};


/**
 * @param {?proto.infra.Credentials|undefined} value
 * @return {!proto.infra.GetInstancesForLabelsRequest} returns this
*/
proto.infra.GetInstancesForLabelsRequest.prototype.setCreds = function(value) {
  return jspb.Message.setWrapperField(this, 6, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.infra.GetInstancesForLabelsRequest} returns this
 */
proto.infra.GetInstancesForLabelsRequest.prototype.clearCreds = function() {
  return this.setCreds(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.infra.GetInstancesForLabelsRequest.prototype.hasCreds = function() {
  return jspb.Message.getField(this, 6) != null;
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.infra.GetInstancesForLabelsResponse.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.infra.GetInstancesForLabelsResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.infra.GetInstancesForLabelsResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.infra.GetInstancesForLabelsResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.GetInstancesForLabelsResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    instancesList: jspb.Message.toObjectList(msg.getInstancesList(),
    types_pb.Instance.toObject, includeInstance),
    err: (f = msg.getErr()) && types_pb.Error.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.infra.GetInstancesForLabelsResponse}
 */
proto.infra.GetInstancesForLabelsResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.infra.GetInstancesForLabelsResponse;
  return proto.infra.GetInstancesForLabelsResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.infra.GetInstancesForLabelsResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.infra.GetInstancesForLabelsResponse}
 */
proto.infra.GetInstancesForLabelsResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new types_pb.Instance;
      reader.readMessage(value,types_pb.Instance.deserializeBinaryFromReader);
      msg.addInstances(value);
      break;
    case 3:
      var value = new types_pb.Error;
      reader.readMessage(value,types_pb.Error.deserializeBinaryFromReader);
      msg.setErr(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.infra.GetInstancesForLabelsResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.infra.GetInstancesForLabelsResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.infra.GetInstancesForLabelsResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.GetInstancesForLabelsResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getInstancesList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      types_pb.Instance.serializeBinaryToWriter
    );
  }
  f = message.getErr();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      types_pb.Error.serializeBinaryToWriter
    );
  }
};


/**
 * repeated Instance instances = 1;
 * @return {!Array<!proto.infra.Instance>}
 */
proto.infra.GetInstancesForLabelsResponse.prototype.getInstancesList = function() {
  return /** @type{!Array<!proto.infra.Instance>} */ (
    jspb.Message.getRepeatedWrapperField(this, types_pb.Instance, 1));
};


/**
 * @param {!Array<!proto.infra.Instance>} value
 * @return {!proto.infra.GetInstancesForLabelsResponse} returns this
*/
proto.infra.GetInstancesForLabelsResponse.prototype.setInstancesList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.infra.Instance=} opt_value
 * @param {number=} opt_index
 * @return {!proto.infra.Instance}
 */
proto.infra.GetInstancesForLabelsResponse.prototype.addInstances = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.infra.Instance, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.infra.GetInstancesForLabelsResponse} returns this
 */
proto.infra.GetInstancesForLabelsResponse.prototype.clearInstancesList = function() {
  return this.setInstancesList([]);
};


/**
 * optional Error err = 3;
 * @return {?proto.infra.Error}
 */
proto.infra.GetInstancesForLabelsResponse.prototype.getErr = function() {
  return /** @type{?proto.infra.Error} */ (
    jspb.Message.getWrapperField(this, types_pb.Error, 3));
};


/**
 * @param {?proto.infra.Error|undefined} value
 * @return {!proto.infra.GetInstancesForLabelsResponse} returns this
*/
proto.infra.GetInstancesForLabelsResponse.prototype.setErr = function(value) {
  return jspb.Message.setWrapperField(this, 3, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.infra.GetInstancesForLabelsResponse} returns this
 */
proto.infra.GetInstancesForLabelsResponse.prototype.clearErr = function() {
  return this.setErr(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.infra.GetInstancesForLabelsResponse.prototype.hasErr = function() {
  return jspb.Message.getField(this, 3) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.infra.GetVPCIDWithTagRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.infra.GetVPCIDWithTagRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.infra.GetVPCIDWithTagRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.GetVPCIDWithTagRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    provider: jspb.Message.getFieldWithDefault(msg, 1, ""),
    accountId: jspb.Message.getFieldWithDefault(msg, 2, ""),
    region: jspb.Message.getFieldWithDefault(msg, 3, ""),
    key: jspb.Message.getFieldWithDefault(msg, 4, ""),
    value: jspb.Message.getFieldWithDefault(msg, 5, ""),
    labelsMap: (f = msg.getLabelsMap()) ? f.toObject(includeInstance, undefined) : [],
    creds: (f = msg.getCreds()) && types_pb.Credentials.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.infra.GetVPCIDWithTagRequest}
 */
proto.infra.GetVPCIDWithTagRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.infra.GetVPCIDWithTagRequest;
  return proto.infra.GetVPCIDWithTagRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.infra.GetVPCIDWithTagRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.infra.GetVPCIDWithTagRequest}
 */
proto.infra.GetVPCIDWithTagRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setProvider(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setAccountId(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setRegion(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setKey(value);
      break;
    case 5:
      var value = /** @type {string} */ (reader.readString());
      msg.setValue(value);
      break;
    case 6:
      var value = msg.getLabelsMap();
      reader.readMessage(value, function(message, reader) {
        jspb.Map.deserializeBinary(message, reader, jspb.BinaryReader.prototype.readString, jspb.BinaryReader.prototype.readString, null, "", "");
         });
      break;
    case 7:
      var value = new types_pb.Credentials;
      reader.readMessage(value,types_pb.Credentials.deserializeBinaryFromReader);
      msg.setCreds(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.infra.GetVPCIDWithTagRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.infra.GetVPCIDWithTagRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.infra.GetVPCIDWithTagRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.GetVPCIDWithTagRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getProvider();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getAccountId();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getRegion();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getKey();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getValue();
  if (f.length > 0) {
    writer.writeString(
      5,
      f
    );
  }
  f = message.getLabelsMap(true);
  if (f && f.getLength() > 0) {
    f.serializeBinary(6, writer, jspb.BinaryWriter.prototype.writeString, jspb.BinaryWriter.prototype.writeString);
  }
  f = message.getCreds();
  if (f != null) {
    writer.writeMessage(
      7,
      f,
      types_pb.Credentials.serializeBinaryToWriter
    );
  }
};


/**
 * optional string provider = 1;
 * @return {string}
 */
proto.infra.GetVPCIDWithTagRequest.prototype.getProvider = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.GetVPCIDWithTagRequest} returns this
 */
proto.infra.GetVPCIDWithTagRequest.prototype.setProvider = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string account_id = 2;
 * @return {string}
 */
proto.infra.GetVPCIDWithTagRequest.prototype.getAccountId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.GetVPCIDWithTagRequest} returns this
 */
proto.infra.GetVPCIDWithTagRequest.prototype.setAccountId = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string region = 3;
 * @return {string}
 */
proto.infra.GetVPCIDWithTagRequest.prototype.getRegion = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.GetVPCIDWithTagRequest} returns this
 */
proto.infra.GetVPCIDWithTagRequest.prototype.setRegion = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string key = 4;
 * @return {string}
 */
proto.infra.GetVPCIDWithTagRequest.prototype.getKey = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.GetVPCIDWithTagRequest} returns this
 */
proto.infra.GetVPCIDWithTagRequest.prototype.setKey = function(value) {
  return jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * optional string value = 5;
 * @return {string}
 */
proto.infra.GetVPCIDWithTagRequest.prototype.getValue = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 5, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.GetVPCIDWithTagRequest} returns this
 */
proto.infra.GetVPCIDWithTagRequest.prototype.setValue = function(value) {
  return jspb.Message.setProto3StringField(this, 5, value);
};


/**
 * map<string, string> labels = 6;
 * @param {boolean=} opt_noLazyCreate Do not create the map if
 * empty, instead returning `undefined`
 * @return {!jspb.Map<string,string>}
 */
proto.infra.GetVPCIDWithTagRequest.prototype.getLabelsMap = function(opt_noLazyCreate) {
  return /** @type {!jspb.Map<string,string>} */ (
      jspb.Message.getMapField(this, 6, opt_noLazyCreate,
      null));
};


/**
 * Clears values from the map. The map will be non-null.
 * @return {!proto.infra.GetVPCIDWithTagRequest} returns this
 */
proto.infra.GetVPCIDWithTagRequest.prototype.clearLabelsMap = function() {
  this.getLabelsMap().clear();
  return this;
};


/**
 * optional Credentials creds = 7;
 * @return {?proto.infra.Credentials}
 */
proto.infra.GetVPCIDWithTagRequest.prototype.getCreds = function() {
  return /** @type{?proto.infra.Credentials} */ (
    jspb.Message.getWrapperField(this, types_pb.Credentials, 7));
};


/**
 * @param {?proto.infra.Credentials|undefined} value
 * @return {!proto.infra.GetVPCIDWithTagRequest} returns this
*/
proto.infra.GetVPCIDWithTagRequest.prototype.setCreds = function(value) {
  return jspb.Message.setWrapperField(this, 7, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.infra.GetVPCIDWithTagRequest} returns this
 */
proto.infra.GetVPCIDWithTagRequest.prototype.clearCreds = function() {
  return this.setCreds(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.infra.GetVPCIDWithTagRequest.prototype.hasCreds = function() {
  return jspb.Message.getField(this, 7) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.infra.GetVPCIDWithTagResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.infra.GetVPCIDWithTagResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.infra.GetVPCIDWithTagResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.GetVPCIDWithTagResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    vpcId: jspb.Message.getFieldWithDefault(msg, 1, ""),
    err: (f = msg.getErr()) && types_pb.Error.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.infra.GetVPCIDWithTagResponse}
 */
proto.infra.GetVPCIDWithTagResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.infra.GetVPCIDWithTagResponse;
  return proto.infra.GetVPCIDWithTagResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.infra.GetVPCIDWithTagResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.infra.GetVPCIDWithTagResponse}
 */
proto.infra.GetVPCIDWithTagResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setVpcId(value);
      break;
    case 3:
      var value = new types_pb.Error;
      reader.readMessage(value,types_pb.Error.deserializeBinaryFromReader);
      msg.setErr(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.infra.GetVPCIDWithTagResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.infra.GetVPCIDWithTagResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.infra.GetVPCIDWithTagResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.GetVPCIDWithTagResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getVpcId();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getErr();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      types_pb.Error.serializeBinaryToWriter
    );
  }
};


/**
 * optional string vpc_id = 1;
 * @return {string}
 */
proto.infra.GetVPCIDWithTagResponse.prototype.getVpcId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.GetVPCIDWithTagResponse} returns this
 */
proto.infra.GetVPCIDWithTagResponse.prototype.setVpcId = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional Error err = 3;
 * @return {?proto.infra.Error}
 */
proto.infra.GetVPCIDWithTagResponse.prototype.getErr = function() {
  return /** @type{?proto.infra.Error} */ (
    jspb.Message.getWrapperField(this, types_pb.Error, 3));
};


/**
 * @param {?proto.infra.Error|undefined} value
 * @return {!proto.infra.GetVPCIDWithTagResponse} returns this
*/
proto.infra.GetVPCIDWithTagResponse.prototype.setErr = function(value) {
  return jspb.Message.setWrapperField(this, 3, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.infra.GetVPCIDWithTagResponse} returns this
 */
proto.infra.GetVPCIDWithTagResponse.prototype.clearErr = function() {
  return this.setErr(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.infra.GetVPCIDWithTagResponse.prototype.hasErr = function() {
  return jspb.Message.getField(this, 3) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.infra.ListCloudClustersRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.infra.ListCloudClustersRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.infra.ListCloudClustersRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListCloudClustersRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    provider: jspb.Message.getFieldWithDefault(msg, 1, ""),
    accountId: jspb.Message.getFieldWithDefault(msg, 2, ""),
    region: jspb.Message.getFieldWithDefault(msg, 3, ""),
    vpcId: jspb.Message.getFieldWithDefault(msg, 4, ""),
    labelsMap: (f = msg.getLabelsMap()) ? f.toObject(includeInstance, undefined) : [],
    creds: (f = msg.getCreds()) && types_pb.Credentials.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.infra.ListCloudClustersRequest}
 */
proto.infra.ListCloudClustersRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.infra.ListCloudClustersRequest;
  return proto.infra.ListCloudClustersRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.infra.ListCloudClustersRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.infra.ListCloudClustersRequest}
 */
proto.infra.ListCloudClustersRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setProvider(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setAccountId(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setRegion(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setVpcId(value);
      break;
    case 5:
      var value = msg.getLabelsMap();
      reader.readMessage(value, function(message, reader) {
        jspb.Map.deserializeBinary(message, reader, jspb.BinaryReader.prototype.readString, jspb.BinaryReader.prototype.readString, null, "", "");
         });
      break;
    case 6:
      var value = new types_pb.Credentials;
      reader.readMessage(value,types_pb.Credentials.deserializeBinaryFromReader);
      msg.setCreds(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.infra.ListCloudClustersRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.infra.ListCloudClustersRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.infra.ListCloudClustersRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListCloudClustersRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getProvider();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getAccountId();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getRegion();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getVpcId();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getLabelsMap(true);
  if (f && f.getLength() > 0) {
    f.serializeBinary(5, writer, jspb.BinaryWriter.prototype.writeString, jspb.BinaryWriter.prototype.writeString);
  }
  f = message.getCreds();
  if (f != null) {
    writer.writeMessage(
      6,
      f,
      types_pb.Credentials.serializeBinaryToWriter
    );
  }
};


/**
 * optional string provider = 1;
 * @return {string}
 */
proto.infra.ListCloudClustersRequest.prototype.getProvider = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListCloudClustersRequest} returns this
 */
proto.infra.ListCloudClustersRequest.prototype.setProvider = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string account_id = 2;
 * @return {string}
 */
proto.infra.ListCloudClustersRequest.prototype.getAccountId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListCloudClustersRequest} returns this
 */
proto.infra.ListCloudClustersRequest.prototype.setAccountId = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string region = 3;
 * @return {string}
 */
proto.infra.ListCloudClustersRequest.prototype.getRegion = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListCloudClustersRequest} returns this
 */
proto.infra.ListCloudClustersRequest.prototype.setRegion = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string vpc_id = 4;
 * @return {string}
 */
proto.infra.ListCloudClustersRequest.prototype.getVpcId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListCloudClustersRequest} returns this
 */
proto.infra.ListCloudClustersRequest.prototype.setVpcId = function(value) {
  return jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * map<string, string> labels = 5;
 * @param {boolean=} opt_noLazyCreate Do not create the map if
 * empty, instead returning `undefined`
 * @return {!jspb.Map<string,string>}
 */
proto.infra.ListCloudClustersRequest.prototype.getLabelsMap = function(opt_noLazyCreate) {
  return /** @type {!jspb.Map<string,string>} */ (
      jspb.Message.getMapField(this, 5, opt_noLazyCreate,
      null));
};


/**
 * Clears values from the map. The map will be non-null.
 * @return {!proto.infra.ListCloudClustersRequest} returns this
 */
proto.infra.ListCloudClustersRequest.prototype.clearLabelsMap = function() {
  this.getLabelsMap().clear();
  return this;
};


/**
 * optional Credentials creds = 6;
 * @return {?proto.infra.Credentials}
 */
proto.infra.ListCloudClustersRequest.prototype.getCreds = function() {
  return /** @type{?proto.infra.Credentials} */ (
    jspb.Message.getWrapperField(this, types_pb.Credentials, 6));
};


/**
 * @param {?proto.infra.Credentials|undefined} value
 * @return {!proto.infra.ListCloudClustersRequest} returns this
*/
proto.infra.ListCloudClustersRequest.prototype.setCreds = function(value) {
  return jspb.Message.setWrapperField(this, 6, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.infra.ListCloudClustersRequest} returns this
 */
proto.infra.ListCloudClustersRequest.prototype.clearCreds = function() {
  return this.setCreds(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.infra.ListCloudClustersRequest.prototype.hasCreds = function() {
  return jspb.Message.getField(this, 6) != null;
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.infra.ListCloudClustersResponse.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.infra.ListCloudClustersResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.infra.ListCloudClustersResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.infra.ListCloudClustersResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListCloudClustersResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    clustersList: jspb.Message.toObjectList(msg.getClustersList(),
    types_k8s_pb.Cluster.toObject, includeInstance),
    lastSyncTime: jspb.Message.getFieldWithDefault(msg, 2, ""),
    err: (f = msg.getErr()) && types_pb.Error.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.infra.ListCloudClustersResponse}
 */
proto.infra.ListCloudClustersResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.infra.ListCloudClustersResponse;
  return proto.infra.ListCloudClustersResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.infra.ListCloudClustersResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.infra.ListCloudClustersResponse}
 */
proto.infra.ListCloudClustersResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new types_k8s_pb.Cluster;
      reader.readMessage(value,types_k8s_pb.Cluster.deserializeBinaryFromReader);
      msg.addClusters(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setLastSyncTime(value);
      break;
    case 3:
      var value = new types_pb.Error;
      reader.readMessage(value,types_pb.Error.deserializeBinaryFromReader);
      msg.setErr(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.infra.ListCloudClustersResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.infra.ListCloudClustersResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.infra.ListCloudClustersResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListCloudClustersResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getClustersList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      types_k8s_pb.Cluster.serializeBinaryToWriter
    );
  }
  f = message.getLastSyncTime();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getErr();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      types_pb.Error.serializeBinaryToWriter
    );
  }
};


/**
 * repeated Cluster clusters = 1;
 * @return {!Array<!proto.infra.Cluster>}
 */
proto.infra.ListCloudClustersResponse.prototype.getClustersList = function() {
  return /** @type{!Array<!proto.infra.Cluster>} */ (
    jspb.Message.getRepeatedWrapperField(this, types_k8s_pb.Cluster, 1));
};


/**
 * @param {!Array<!proto.infra.Cluster>} value
 * @return {!proto.infra.ListCloudClustersResponse} returns this
*/
proto.infra.ListCloudClustersResponse.prototype.setClustersList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.infra.Cluster=} opt_value
 * @param {number=} opt_index
 * @return {!proto.infra.Cluster}
 */
proto.infra.ListCloudClustersResponse.prototype.addClusters = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.infra.Cluster, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.infra.ListCloudClustersResponse} returns this
 */
proto.infra.ListCloudClustersResponse.prototype.clearClustersList = function() {
  return this.setClustersList([]);
};


/**
 * optional string last_sync_time = 2;
 * @return {string}
 */
proto.infra.ListCloudClustersResponse.prototype.getLastSyncTime = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListCloudClustersResponse} returns this
 */
proto.infra.ListCloudClustersResponse.prototype.setLastSyncTime = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional Error err = 3;
 * @return {?proto.infra.Error}
 */
proto.infra.ListCloudClustersResponse.prototype.getErr = function() {
  return /** @type{?proto.infra.Error} */ (
    jspb.Message.getWrapperField(this, types_pb.Error, 3));
};


/**
 * @param {?proto.infra.Error|undefined} value
 * @return {!proto.infra.ListCloudClustersResponse} returns this
*/
proto.infra.ListCloudClustersResponse.prototype.setErr = function(value) {
  return jspb.Message.setWrapperField(this, 3, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.infra.ListCloudClustersResponse} returns this
 */
proto.infra.ListCloudClustersResponse.prototype.clearErr = function() {
  return this.setErr(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.infra.ListCloudClustersResponse.prototype.hasErr = function() {
  return jspb.Message.getField(this, 3) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.infra.SummaryRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.infra.SummaryRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.infra.SummaryRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.SummaryRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    provider: jspb.Message.getFieldWithDefault(msg, 1, ""),
    accountId: jspb.Message.getFieldWithDefault(msg, 2, ""),
    region: jspb.Message.getFieldWithDefault(msg, 3, ""),
    vpcId: jspb.Message.getFieldWithDefault(msg, 4, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.infra.SummaryRequest}
 */
proto.infra.SummaryRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.infra.SummaryRequest;
  return proto.infra.SummaryRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.infra.SummaryRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.infra.SummaryRequest}
 */
proto.infra.SummaryRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setProvider(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setAccountId(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setRegion(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setVpcId(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.infra.SummaryRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.infra.SummaryRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.infra.SummaryRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.SummaryRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getProvider();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getAccountId();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getRegion();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getVpcId();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
};


/**
 * optional string provider = 1;
 * @return {string}
 */
proto.infra.SummaryRequest.prototype.getProvider = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.SummaryRequest} returns this
 */
proto.infra.SummaryRequest.prototype.setProvider = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string account_id = 2;
 * @return {string}
 */
proto.infra.SummaryRequest.prototype.getAccountId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.SummaryRequest} returns this
 */
proto.infra.SummaryRequest.prototype.setAccountId = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string region = 3;
 * @return {string}
 */
proto.infra.SummaryRequest.prototype.getRegion = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.SummaryRequest} returns this
 */
proto.infra.SummaryRequest.prototype.setRegion = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string vpc_id = 4;
 * @return {string}
 */
proto.infra.SummaryRequest.prototype.getVpcId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.SummaryRequest} returns this
 */
proto.infra.SummaryRequest.prototype.setVpcId = function(value) {
  return jspb.Message.setProto3StringField(this, 4, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.infra.Counters.prototype.toObject = function(opt_includeInstance) {
  return proto.infra.Counters.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.infra.Counters} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.Counters.toObject = function(includeInstance, msg) {
  var f, obj = {
    accounts: jspb.Message.getFieldWithDefault(msg, 1, 0),
    vpc: jspb.Message.getFieldWithDefault(msg, 2, 0),
    subnets: jspb.Message.getFieldWithDefault(msg, 3, 0),
    routeTables: jspb.Message.getFieldWithDefault(msg, 4, 0),
    instances: jspb.Message.getFieldWithDefault(msg, 5, 0),
    clusters: jspb.Message.getFieldWithDefault(msg, 6, 0),
    pods: jspb.Message.getFieldWithDefault(msg, 7, 0),
    services: jspb.Message.getFieldWithDefault(msg, 8, 0),
    namespaces: jspb.Message.getFieldWithDefault(msg, 9, 0),
    acls: jspb.Message.getFieldWithDefault(msg, 10, 0),
    securityGroups: jspb.Message.getFieldWithDefault(msg, 11, 0),
    natGateways: jspb.Message.getFieldWithDefault(msg, 12, 0),
    routers: jspb.Message.getFieldWithDefault(msg, 13, 0),
    igws: jspb.Message.getFieldWithDefault(msg, 14, 0),
    vpcEndpoints: jspb.Message.getFieldWithDefault(msg, 15, 0),
    publicIps: jspb.Message.getFieldWithDefault(msg, 16, 0),
    internetGateways: jspb.Message.getFieldWithDefault(msg, 17, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.infra.Counters}
 */
proto.infra.Counters.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.infra.Counters;
  return proto.infra.Counters.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.infra.Counters} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.infra.Counters}
 */
proto.infra.Counters.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setAccounts(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setVpc(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setSubnets(value);
      break;
    case 4:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setRouteTables(value);
      break;
    case 5:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setInstances(value);
      break;
    case 6:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setClusters(value);
      break;
    case 7:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setPods(value);
      break;
    case 8:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setServices(value);
      break;
    case 9:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setNamespaces(value);
      break;
    case 10:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setAcls(value);
      break;
    case 11:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setSecurityGroups(value);
      break;
    case 12:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setNatGateways(value);
      break;
    case 13:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setRouters(value);
      break;
    case 14:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setIgws(value);
      break;
    case 15:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setVpcEndpoints(value);
      break;
    case 16:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setPublicIps(value);
      break;
    case 17:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setInternetGateways(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.infra.Counters.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.infra.Counters.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.infra.Counters} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.Counters.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getAccounts();
  if (f !== 0) {
    writer.writeInt32(
      1,
      f
    );
  }
  f = message.getVpc();
  if (f !== 0) {
    writer.writeInt32(
      2,
      f
    );
  }
  f = message.getSubnets();
  if (f !== 0) {
    writer.writeInt32(
      3,
      f
    );
  }
  f = message.getRouteTables();
  if (f !== 0) {
    writer.writeInt32(
      4,
      f
    );
  }
  f = message.getInstances();
  if (f !== 0) {
    writer.writeInt32(
      5,
      f
    );
  }
  f = message.getClusters();
  if (f !== 0) {
    writer.writeInt32(
      6,
      f
    );
  }
  f = message.getPods();
  if (f !== 0) {
    writer.writeInt32(
      7,
      f
    );
  }
  f = message.getServices();
  if (f !== 0) {
    writer.writeInt32(
      8,
      f
    );
  }
  f = message.getNamespaces();
  if (f !== 0) {
    writer.writeInt32(
      9,
      f
    );
  }
  f = message.getAcls();
  if (f !== 0) {
    writer.writeInt32(
      10,
      f
    );
  }
  f = message.getSecurityGroups();
  if (f !== 0) {
    writer.writeInt32(
      11,
      f
    );
  }
  f = message.getNatGateways();
  if (f !== 0) {
    writer.writeInt32(
      12,
      f
    );
  }
  f = message.getRouters();
  if (f !== 0) {
    writer.writeInt32(
      13,
      f
    );
  }
  f = message.getIgws();
  if (f !== 0) {
    writer.writeInt32(
      14,
      f
    );
  }
  f = message.getVpcEndpoints();
  if (f !== 0) {
    writer.writeInt32(
      15,
      f
    );
  }
  f = message.getPublicIps();
  if (f !== 0) {
    writer.writeInt32(
      16,
      f
    );
  }
  f = message.getInternetGateways();
  if (f !== 0) {
    writer.writeInt32(
      17,
      f
    );
  }
};


/**
 * optional int32 accounts = 1;
 * @return {number}
 */
proto.infra.Counters.prototype.getAccounts = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {number} value
 * @return {!proto.infra.Counters} returns this
 */
proto.infra.Counters.prototype.setAccounts = function(value) {
  return jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional int32 vpc = 2;
 * @return {number}
 */
proto.infra.Counters.prototype.getVpc = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/**
 * @param {number} value
 * @return {!proto.infra.Counters} returns this
 */
proto.infra.Counters.prototype.setVpc = function(value) {
  return jspb.Message.setProto3IntField(this, 2, value);
};


/**
 * optional int32 subnets = 3;
 * @return {number}
 */
proto.infra.Counters.prototype.getSubnets = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/**
 * @param {number} value
 * @return {!proto.infra.Counters} returns this
 */
proto.infra.Counters.prototype.setSubnets = function(value) {
  return jspb.Message.setProto3IntField(this, 3, value);
};


/**
 * optional int32 route_tables = 4;
 * @return {number}
 */
proto.infra.Counters.prototype.getRouteTables = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 4, 0));
};


/**
 * @param {number} value
 * @return {!proto.infra.Counters} returns this
 */
proto.infra.Counters.prototype.setRouteTables = function(value) {
  return jspb.Message.setProto3IntField(this, 4, value);
};


/**
 * optional int32 instances = 5;
 * @return {number}
 */
proto.infra.Counters.prototype.getInstances = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 5, 0));
};


/**
 * @param {number} value
 * @return {!proto.infra.Counters} returns this
 */
proto.infra.Counters.prototype.setInstances = function(value) {
  return jspb.Message.setProto3IntField(this, 5, value);
};


/**
 * optional int32 clusters = 6;
 * @return {number}
 */
proto.infra.Counters.prototype.getClusters = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 6, 0));
};


/**
 * @param {number} value
 * @return {!proto.infra.Counters} returns this
 */
proto.infra.Counters.prototype.setClusters = function(value) {
  return jspb.Message.setProto3IntField(this, 6, value);
};


/**
 * optional int32 pods = 7;
 * @return {number}
 */
proto.infra.Counters.prototype.getPods = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 7, 0));
};


/**
 * @param {number} value
 * @return {!proto.infra.Counters} returns this
 */
proto.infra.Counters.prototype.setPods = function(value) {
  return jspb.Message.setProto3IntField(this, 7, value);
};


/**
 * optional int32 services = 8;
 * @return {number}
 */
proto.infra.Counters.prototype.getServices = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 8, 0));
};


/**
 * @param {number} value
 * @return {!proto.infra.Counters} returns this
 */
proto.infra.Counters.prototype.setServices = function(value) {
  return jspb.Message.setProto3IntField(this, 8, value);
};


/**
 * optional int32 namespaces = 9;
 * @return {number}
 */
proto.infra.Counters.prototype.getNamespaces = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 9, 0));
};


/**
 * @param {number} value
 * @return {!proto.infra.Counters} returns this
 */
proto.infra.Counters.prototype.setNamespaces = function(value) {
  return jspb.Message.setProto3IntField(this, 9, value);
};


/**
 * optional int32 acls = 10;
 * @return {number}
 */
proto.infra.Counters.prototype.getAcls = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 10, 0));
};


/**
 * @param {number} value
 * @return {!proto.infra.Counters} returns this
 */
proto.infra.Counters.prototype.setAcls = function(value) {
  return jspb.Message.setProto3IntField(this, 10, value);
};


/**
 * optional int32 security_groups = 11;
 * @return {number}
 */
proto.infra.Counters.prototype.getSecurityGroups = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 11, 0));
};


/**
 * @param {number} value
 * @return {!proto.infra.Counters} returns this
 */
proto.infra.Counters.prototype.setSecurityGroups = function(value) {
  return jspb.Message.setProto3IntField(this, 11, value);
};


/**
 * optional int32 nat_gateways = 12;
 * @return {number}
 */
proto.infra.Counters.prototype.getNatGateways = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 12, 0));
};


/**
 * @param {number} value
 * @return {!proto.infra.Counters} returns this
 */
proto.infra.Counters.prototype.setNatGateways = function(value) {
  return jspb.Message.setProto3IntField(this, 12, value);
};


/**
 * optional int32 routers = 13;
 * @return {number}
 */
proto.infra.Counters.prototype.getRouters = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 13, 0));
};


/**
 * @param {number} value
 * @return {!proto.infra.Counters} returns this
 */
proto.infra.Counters.prototype.setRouters = function(value) {
  return jspb.Message.setProto3IntField(this, 13, value);
};


/**
 * optional int32 igws = 14;
 * @return {number}
 */
proto.infra.Counters.prototype.getIgws = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 14, 0));
};


/**
 * @param {number} value
 * @return {!proto.infra.Counters} returns this
 */
proto.infra.Counters.prototype.setIgws = function(value) {
  return jspb.Message.setProto3IntField(this, 14, value);
};


/**
 * optional int32 vpc_endpoints = 15;
 * @return {number}
 */
proto.infra.Counters.prototype.getVpcEndpoints = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 15, 0));
};


/**
 * @param {number} value
 * @return {!proto.infra.Counters} returns this
 */
proto.infra.Counters.prototype.setVpcEndpoints = function(value) {
  return jspb.Message.setProto3IntField(this, 15, value);
};


/**
 * optional int32 public_ips = 16;
 * @return {number}
 */
proto.infra.Counters.prototype.getPublicIps = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 16, 0));
};


/**
 * @param {number} value
 * @return {!proto.infra.Counters} returns this
 */
proto.infra.Counters.prototype.setPublicIps = function(value) {
  return jspb.Message.setProto3IntField(this, 16, value);
};


/**
 * optional int32 internet_gateways = 17;
 * @return {number}
 */
proto.infra.Counters.prototype.getInternetGateways = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 17, 0));
};


/**
 * @param {number} value
 * @return {!proto.infra.Counters} returns this
 */
proto.infra.Counters.prototype.setInternetGateways = function(value) {
  return jspb.Message.setProto3IntField(this, 17, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.infra.StatusSummary.prototype.toObject = function(opt_includeInstance) {
  return proto.infra.StatusSummary.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.infra.StatusSummary} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.StatusSummary.toObject = function(includeInstance, msg) {
  var f, obj = {
    vmStatusMap: (f = msg.getVmStatusMap()) ? f.toObject(includeInstance, undefined) : [],
    podStatusMap: (f = msg.getPodStatusMap()) ? f.toObject(includeInstance, undefined) : [],
    vmTypesMap: (f = msg.getVmTypesMap()) ? f.toObject(includeInstance, undefined) : []
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.infra.StatusSummary}
 */
proto.infra.StatusSummary.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.infra.StatusSummary;
  return proto.infra.StatusSummary.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.infra.StatusSummary} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.infra.StatusSummary}
 */
proto.infra.StatusSummary.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = msg.getVmStatusMap();
      reader.readMessage(value, function(message, reader) {
        jspb.Map.deserializeBinary(message, reader, jspb.BinaryReader.prototype.readString, jspb.BinaryReader.prototype.readInt32, null, "", 0);
         });
      break;
    case 2:
      var value = msg.getPodStatusMap();
      reader.readMessage(value, function(message, reader) {
        jspb.Map.deserializeBinary(message, reader, jspb.BinaryReader.prototype.readString, jspb.BinaryReader.prototype.readInt32, null, "", 0);
         });
      break;
    case 3:
      var value = msg.getVmTypesMap();
      reader.readMessage(value, function(message, reader) {
        jspb.Map.deserializeBinary(message, reader, jspb.BinaryReader.prototype.readString, jspb.BinaryReader.prototype.readInt32, null, "", 0);
         });
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.infra.StatusSummary.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.infra.StatusSummary.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.infra.StatusSummary} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.StatusSummary.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getVmStatusMap(true);
  if (f && f.getLength() > 0) {
    f.serializeBinary(1, writer, jspb.BinaryWriter.prototype.writeString, jspb.BinaryWriter.prototype.writeInt32);
  }
  f = message.getPodStatusMap(true);
  if (f && f.getLength() > 0) {
    f.serializeBinary(2, writer, jspb.BinaryWriter.prototype.writeString, jspb.BinaryWriter.prototype.writeInt32);
  }
  f = message.getVmTypesMap(true);
  if (f && f.getLength() > 0) {
    f.serializeBinary(3, writer, jspb.BinaryWriter.prototype.writeString, jspb.BinaryWriter.prototype.writeInt32);
  }
};


/**
 * map<string, int32> vm_status = 1;
 * @param {boolean=} opt_noLazyCreate Do not create the map if
 * empty, instead returning `undefined`
 * @return {!jspb.Map<string,number>}
 */
proto.infra.StatusSummary.prototype.getVmStatusMap = function(opt_noLazyCreate) {
  return /** @type {!jspb.Map<string,number>} */ (
      jspb.Message.getMapField(this, 1, opt_noLazyCreate,
      null));
};


/**
 * Clears values from the map. The map will be non-null.
 * @return {!proto.infra.StatusSummary} returns this
 */
proto.infra.StatusSummary.prototype.clearVmStatusMap = function() {
  this.getVmStatusMap().clear();
  return this;
};


/**
 * map<string, int32> pod_status = 2;
 * @param {boolean=} opt_noLazyCreate Do not create the map if
 * empty, instead returning `undefined`
 * @return {!jspb.Map<string,number>}
 */
proto.infra.StatusSummary.prototype.getPodStatusMap = function(opt_noLazyCreate) {
  return /** @type {!jspb.Map<string,number>} */ (
      jspb.Message.getMapField(this, 2, opt_noLazyCreate,
      null));
};


/**
 * Clears values from the map. The map will be non-null.
 * @return {!proto.infra.StatusSummary} returns this
 */
proto.infra.StatusSummary.prototype.clearPodStatusMap = function() {
  this.getPodStatusMap().clear();
  return this;
};


/**
 * map<string, int32> vm_types = 3;
 * @param {boolean=} opt_noLazyCreate Do not create the map if
 * empty, instead returning `undefined`
 * @return {!jspb.Map<string,number>}
 */
proto.infra.StatusSummary.prototype.getVmTypesMap = function(opt_noLazyCreate) {
  return /** @type {!jspb.Map<string,number>} */ (
      jspb.Message.getMapField(this, 3, opt_noLazyCreate,
      null));
};


/**
 * Clears values from the map. The map will be non-null.
 * @return {!proto.infra.StatusSummary} returns this
 */
proto.infra.StatusSummary.prototype.clearVmTypesMap = function() {
  this.getVmTypesMap().clear();
  return this;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.infra.SummaryResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.infra.SummaryResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.infra.SummaryResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.SummaryResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    count: (f = msg.getCount()) && proto.infra.Counters.toObject(includeInstance, f),
    statuses: (f = msg.getStatuses()) && proto.infra.StatusSummary.toObject(includeInstance, f),
    err: (f = msg.getErr()) && types_pb.Error.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.infra.SummaryResponse}
 */
proto.infra.SummaryResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.infra.SummaryResponse;
  return proto.infra.SummaryResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.infra.SummaryResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.infra.SummaryResponse}
 */
proto.infra.SummaryResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.infra.Counters;
      reader.readMessage(value,proto.infra.Counters.deserializeBinaryFromReader);
      msg.setCount(value);
      break;
    case 2:
      var value = new proto.infra.StatusSummary;
      reader.readMessage(value,proto.infra.StatusSummary.deserializeBinaryFromReader);
      msg.setStatuses(value);
      break;
    case 3:
      var value = new types_pb.Error;
      reader.readMessage(value,types_pb.Error.deserializeBinaryFromReader);
      msg.setErr(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.infra.SummaryResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.infra.SummaryResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.infra.SummaryResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.SummaryResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getCount();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      proto.infra.Counters.serializeBinaryToWriter
    );
  }
  f = message.getStatuses();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      proto.infra.StatusSummary.serializeBinaryToWriter
    );
  }
  f = message.getErr();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      types_pb.Error.serializeBinaryToWriter
    );
  }
};


/**
 * optional Counters count = 1;
 * @return {?proto.infra.Counters}
 */
proto.infra.SummaryResponse.prototype.getCount = function() {
  return /** @type{?proto.infra.Counters} */ (
    jspb.Message.getWrapperField(this, proto.infra.Counters, 1));
};


/**
 * @param {?proto.infra.Counters|undefined} value
 * @return {!proto.infra.SummaryResponse} returns this
*/
proto.infra.SummaryResponse.prototype.setCount = function(value) {
  return jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.infra.SummaryResponse} returns this
 */
proto.infra.SummaryResponse.prototype.clearCount = function() {
  return this.setCount(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.infra.SummaryResponse.prototype.hasCount = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional StatusSummary statuses = 2;
 * @return {?proto.infra.StatusSummary}
 */
proto.infra.SummaryResponse.prototype.getStatuses = function() {
  return /** @type{?proto.infra.StatusSummary} */ (
    jspb.Message.getWrapperField(this, proto.infra.StatusSummary, 2));
};


/**
 * @param {?proto.infra.StatusSummary|undefined} value
 * @return {!proto.infra.SummaryResponse} returns this
*/
proto.infra.SummaryResponse.prototype.setStatuses = function(value) {
  return jspb.Message.setWrapperField(this, 2, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.infra.SummaryResponse} returns this
 */
proto.infra.SummaryResponse.prototype.clearStatuses = function() {
  return this.setStatuses(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.infra.SummaryResponse.prototype.hasStatuses = function() {
  return jspb.Message.getField(this, 2) != null;
};


/**
 * optional Error err = 3;
 * @return {?proto.infra.Error}
 */
proto.infra.SummaryResponse.prototype.getErr = function() {
  return /** @type{?proto.infra.Error} */ (
    jspb.Message.getWrapperField(this, types_pb.Error, 3));
};


/**
 * @param {?proto.infra.Error|undefined} value
 * @return {!proto.infra.SummaryResponse} returns this
*/
proto.infra.SummaryResponse.prototype.setErr = function(value) {
  return jspb.Message.setWrapperField(this, 3, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.infra.SummaryResponse} returns this
 */
proto.infra.SummaryResponse.prototype.clearErr = function() {
  return this.setErr(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.infra.SummaryResponse.prototype.hasErr = function() {
  return jspb.Message.getField(this, 3) != null;
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.infra.SearchResourcesRequest.repeatedFields_ = [19];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.infra.SearchResourcesRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.infra.SearchResourcesRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.infra.SearchResourcesRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.SearchResourcesRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    filterProvider: jspb.Message.getFieldWithDefault(msg, 1, ""),
    filterAccountId: jspb.Message.getFieldWithDefault(msg, 2, ""),
    filterRegion: jspb.Message.getFieldWithDefault(msg, 3, ""),
    filterVpcId: jspb.Message.getFieldWithDefault(msg, 4, ""),
    filterZone: jspb.Message.getFieldWithDefault(msg, 5, ""),
    searchLabelsMap: (f = msg.getSearchLabelsMap()) ? f.toObject(includeInstance, undefined) : [],
    searchName: jspb.Message.getFieldWithDefault(msg, 7, ""),
    searchId: jspb.Message.getFieldWithDefault(msg, 8, ""),
    searchStatus: jspb.Message.getFieldWithDefault(msg, 9, ""),
    searchCreationTimeStart: (f = msg.getSearchCreationTimeStart()) && google_protobuf_timestamp_pb.Timestamp.toObject(includeInstance, f),
    searchCreationTimeEnd: (f = msg.getSearchCreationTimeEnd()) && google_protobuf_timestamp_pb.Timestamp.toObject(includeInstance, f),
    searchTerminationTimeStart: (f = msg.getSearchTerminationTimeStart()) && google_protobuf_timestamp_pb.Timestamp.toObject(includeInstance, f),
    searchTerminationTimeEnd: (f = msg.getSearchTerminationTimeEnd()) && google_protobuf_timestamp_pb.Timestamp.toObject(includeInstance, f),
    pageSize: jspb.Message.getFieldWithDefault(msg, 14, 0),
    pageNumber: jspb.Message.getFieldWithDefault(msg, 15, 0),
    sortBy: jspb.Message.getFieldWithDefault(msg, 16, ""),
    sortDescending: jspb.Message.getBooleanFieldWithDefault(msg, 17, false),
    fieldMask: (f = msg.getFieldMask()) && google_protobuf_field_mask_pb.FieldMask.toObject(includeInstance, f),
    resourceTypesList: (f = jspb.Message.getRepeatedField(msg, 19)) == null ? undefined : f
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.infra.SearchResourcesRequest}
 */
proto.infra.SearchResourcesRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.infra.SearchResourcesRequest;
  return proto.infra.SearchResourcesRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.infra.SearchResourcesRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.infra.SearchResourcesRequest}
 */
proto.infra.SearchResourcesRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setFilterProvider(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setFilterAccountId(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setFilterRegion(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setFilterVpcId(value);
      break;
    case 5:
      var value = /** @type {string} */ (reader.readString());
      msg.setFilterZone(value);
      break;
    case 6:
      var value = msg.getSearchLabelsMap();
      reader.readMessage(value, function(message, reader) {
        jspb.Map.deserializeBinary(message, reader, jspb.BinaryReader.prototype.readString, jspb.BinaryReader.prototype.readString, null, "", "");
         });
      break;
    case 7:
      var value = /** @type {string} */ (reader.readString());
      msg.setSearchName(value);
      break;
    case 8:
      var value = /** @type {string} */ (reader.readString());
      msg.setSearchId(value);
      break;
    case 9:
      var value = /** @type {string} */ (reader.readString());
      msg.setSearchStatus(value);
      break;
    case 10:
      var value = new google_protobuf_timestamp_pb.Timestamp;
      reader.readMessage(value,google_protobuf_timestamp_pb.Timestamp.deserializeBinaryFromReader);
      msg.setSearchCreationTimeStart(value);
      break;
    case 11:
      var value = new google_protobuf_timestamp_pb.Timestamp;
      reader.readMessage(value,google_protobuf_timestamp_pb.Timestamp.deserializeBinaryFromReader);
      msg.setSearchCreationTimeEnd(value);
      break;
    case 12:
      var value = new google_protobuf_timestamp_pb.Timestamp;
      reader.readMessage(value,google_protobuf_timestamp_pb.Timestamp.deserializeBinaryFromReader);
      msg.setSearchTerminationTimeStart(value);
      break;
    case 13:
      var value = new google_protobuf_timestamp_pb.Timestamp;
      reader.readMessage(value,google_protobuf_timestamp_pb.Timestamp.deserializeBinaryFromReader);
      msg.setSearchTerminationTimeEnd(value);
      break;
    case 14:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setPageSize(value);
      break;
    case 15:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setPageNumber(value);
      break;
    case 16:
      var value = /** @type {string} */ (reader.readString());
      msg.setSortBy(value);
      break;
    case 17:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setSortDescending(value);
      break;
    case 18:
      var value = new google_protobuf_field_mask_pb.FieldMask;
      reader.readMessage(value,google_protobuf_field_mask_pb.FieldMask.deserializeBinaryFromReader);
      msg.setFieldMask(value);
      break;
    case 19:
      var value = /** @type {string} */ (reader.readString());
      msg.addResourceTypes(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.infra.SearchResourcesRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.infra.SearchResourcesRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.infra.SearchResourcesRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.SearchResourcesRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getFilterProvider();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getFilterAccountId();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getFilterRegion();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getFilterVpcId();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getFilterZone();
  if (f.length > 0) {
    writer.writeString(
      5,
      f
    );
  }
  f = message.getSearchLabelsMap(true);
  if (f && f.getLength() > 0) {
    f.serializeBinary(6, writer, jspb.BinaryWriter.prototype.writeString, jspb.BinaryWriter.prototype.writeString);
  }
  f = message.getSearchName();
  if (f.length > 0) {
    writer.writeString(
      7,
      f
    );
  }
  f = message.getSearchId();
  if (f.length > 0) {
    writer.writeString(
      8,
      f
    );
  }
  f = message.getSearchStatus();
  if (f.length > 0) {
    writer.writeString(
      9,
      f
    );
  }
  f = message.getSearchCreationTimeStart();
  if (f != null) {
    writer.writeMessage(
      10,
      f,
      google_protobuf_timestamp_pb.Timestamp.serializeBinaryToWriter
    );
  }
  f = message.getSearchCreationTimeEnd();
  if (f != null) {
    writer.writeMessage(
      11,
      f,
      google_protobuf_timestamp_pb.Timestamp.serializeBinaryToWriter
    );
  }
  f = message.getSearchTerminationTimeStart();
  if (f != null) {
    writer.writeMessage(
      12,
      f,
      google_protobuf_timestamp_pb.Timestamp.serializeBinaryToWriter
    );
  }
  f = message.getSearchTerminationTimeEnd();
  if (f != null) {
    writer.writeMessage(
      13,
      f,
      google_protobuf_timestamp_pb.Timestamp.serializeBinaryToWriter
    );
  }
  f = message.getPageSize();
  if (f !== 0) {
    writer.writeInt32(
      14,
      f
    );
  }
  f = message.getPageNumber();
  if (f !== 0) {
    writer.writeInt32(
      15,
      f
    );
  }
  f = message.getSortBy();
  if (f.length > 0) {
    writer.writeString(
      16,
      f
    );
  }
  f = message.getSortDescending();
  if (f) {
    writer.writeBool(
      17,
      f
    );
  }
  f = message.getFieldMask();
  if (f != null) {
    writer.writeMessage(
      18,
      f,
      google_protobuf_field_mask_pb.FieldMask.serializeBinaryToWriter
    );
  }
  f = message.getResourceTypesList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      19,
      f
    );
  }
};


/**
 * optional string filter_provider = 1;
 * @return {string}
 */
proto.infra.SearchResourcesRequest.prototype.getFilterProvider = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.SearchResourcesRequest} returns this
 */
proto.infra.SearchResourcesRequest.prototype.setFilterProvider = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string filter_account_id = 2;
 * @return {string}
 */
proto.infra.SearchResourcesRequest.prototype.getFilterAccountId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.SearchResourcesRequest} returns this
 */
proto.infra.SearchResourcesRequest.prototype.setFilterAccountId = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string filter_region = 3;
 * @return {string}
 */
proto.infra.SearchResourcesRequest.prototype.getFilterRegion = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.SearchResourcesRequest} returns this
 */
proto.infra.SearchResourcesRequest.prototype.setFilterRegion = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string filter_vpc_id = 4;
 * @return {string}
 */
proto.infra.SearchResourcesRequest.prototype.getFilterVpcId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.SearchResourcesRequest} returns this
 */
proto.infra.SearchResourcesRequest.prototype.setFilterVpcId = function(value) {
  return jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * optional string filter_zone = 5;
 * @return {string}
 */
proto.infra.SearchResourcesRequest.prototype.getFilterZone = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 5, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.SearchResourcesRequest} returns this
 */
proto.infra.SearchResourcesRequest.prototype.setFilterZone = function(value) {
  return jspb.Message.setProto3StringField(this, 5, value);
};


/**
 * map<string, string> search_labels = 6;
 * @param {boolean=} opt_noLazyCreate Do not create the map if
 * empty, instead returning `undefined`
 * @return {!jspb.Map<string,string>}
 */
proto.infra.SearchResourcesRequest.prototype.getSearchLabelsMap = function(opt_noLazyCreate) {
  return /** @type {!jspb.Map<string,string>} */ (
      jspb.Message.getMapField(this, 6, opt_noLazyCreate,
      null));
};


/**
 * Clears values from the map. The map will be non-null.
 * @return {!proto.infra.SearchResourcesRequest} returns this
 */
proto.infra.SearchResourcesRequest.prototype.clearSearchLabelsMap = function() {
  this.getSearchLabelsMap().clear();
  return this;
};


/**
 * optional string search_name = 7;
 * @return {string}
 */
proto.infra.SearchResourcesRequest.prototype.getSearchName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 7, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.SearchResourcesRequest} returns this
 */
proto.infra.SearchResourcesRequest.prototype.setSearchName = function(value) {
  return jspb.Message.setProto3StringField(this, 7, value);
};


/**
 * optional string search_id = 8;
 * @return {string}
 */
proto.infra.SearchResourcesRequest.prototype.getSearchId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 8, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.SearchResourcesRequest} returns this
 */
proto.infra.SearchResourcesRequest.prototype.setSearchId = function(value) {
  return jspb.Message.setProto3StringField(this, 8, value);
};


/**
 * optional string search_status = 9;
 * @return {string}
 */
proto.infra.SearchResourcesRequest.prototype.getSearchStatus = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 9, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.SearchResourcesRequest} returns this
 */
proto.infra.SearchResourcesRequest.prototype.setSearchStatus = function(value) {
  return jspb.Message.setProto3StringField(this, 9, value);
};


/**
 * optional google.protobuf.Timestamp search_creation_time_start = 10;
 * @return {?proto.google.protobuf.Timestamp}
 */
proto.infra.SearchResourcesRequest.prototype.getSearchCreationTimeStart = function() {
  return /** @type{?proto.google.protobuf.Timestamp} */ (
    jspb.Message.getWrapperField(this, google_protobuf_timestamp_pb.Timestamp, 10));
};


/**
 * @param {?proto.google.protobuf.Timestamp|undefined} value
 * @return {!proto.infra.SearchResourcesRequest} returns this
*/
proto.infra.SearchResourcesRequest.prototype.setSearchCreationTimeStart = function(value) {
  return jspb.Message.setWrapperField(this, 10, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.infra.SearchResourcesRequest} returns this
 */
proto.infra.SearchResourcesRequest.prototype.clearSearchCreationTimeStart = function() {
  return this.setSearchCreationTimeStart(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.infra.SearchResourcesRequest.prototype.hasSearchCreationTimeStart = function() {
  return jspb.Message.getField(this, 10) != null;
};


/**
 * optional google.protobuf.Timestamp search_creation_time_end = 11;
 * @return {?proto.google.protobuf.Timestamp}
 */
proto.infra.SearchResourcesRequest.prototype.getSearchCreationTimeEnd = function() {
  return /** @type{?proto.google.protobuf.Timestamp} */ (
    jspb.Message.getWrapperField(this, google_protobuf_timestamp_pb.Timestamp, 11));
};


/**
 * @param {?proto.google.protobuf.Timestamp|undefined} value
 * @return {!proto.infra.SearchResourcesRequest} returns this
*/
proto.infra.SearchResourcesRequest.prototype.setSearchCreationTimeEnd = function(value) {
  return jspb.Message.setWrapperField(this, 11, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.infra.SearchResourcesRequest} returns this
 */
proto.infra.SearchResourcesRequest.prototype.clearSearchCreationTimeEnd = function() {
  return this.setSearchCreationTimeEnd(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.infra.SearchResourcesRequest.prototype.hasSearchCreationTimeEnd = function() {
  return jspb.Message.getField(this, 11) != null;
};


/**
 * optional google.protobuf.Timestamp search_termination_time_start = 12;
 * @return {?proto.google.protobuf.Timestamp}
 */
proto.infra.SearchResourcesRequest.prototype.getSearchTerminationTimeStart = function() {
  return /** @type{?proto.google.protobuf.Timestamp} */ (
    jspb.Message.getWrapperField(this, google_protobuf_timestamp_pb.Timestamp, 12));
};


/**
 * @param {?proto.google.protobuf.Timestamp|undefined} value
 * @return {!proto.infra.SearchResourcesRequest} returns this
*/
proto.infra.SearchResourcesRequest.prototype.setSearchTerminationTimeStart = function(value) {
  return jspb.Message.setWrapperField(this, 12, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.infra.SearchResourcesRequest} returns this
 */
proto.infra.SearchResourcesRequest.prototype.clearSearchTerminationTimeStart = function() {
  return this.setSearchTerminationTimeStart(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.infra.SearchResourcesRequest.prototype.hasSearchTerminationTimeStart = function() {
  return jspb.Message.getField(this, 12) != null;
};


/**
 * optional google.protobuf.Timestamp search_termination_time_end = 13;
 * @return {?proto.google.protobuf.Timestamp}
 */
proto.infra.SearchResourcesRequest.prototype.getSearchTerminationTimeEnd = function() {
  return /** @type{?proto.google.protobuf.Timestamp} */ (
    jspb.Message.getWrapperField(this, google_protobuf_timestamp_pb.Timestamp, 13));
};


/**
 * @param {?proto.google.protobuf.Timestamp|undefined} value
 * @return {!proto.infra.SearchResourcesRequest} returns this
*/
proto.infra.SearchResourcesRequest.prototype.setSearchTerminationTimeEnd = function(value) {
  return jspb.Message.setWrapperField(this, 13, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.infra.SearchResourcesRequest} returns this
 */
proto.infra.SearchResourcesRequest.prototype.clearSearchTerminationTimeEnd = function() {
  return this.setSearchTerminationTimeEnd(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.infra.SearchResourcesRequest.prototype.hasSearchTerminationTimeEnd = function() {
  return jspb.Message.getField(this, 13) != null;
};


/**
 * optional int32 page_size = 14;
 * @return {number}
 */
proto.infra.SearchResourcesRequest.prototype.getPageSize = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 14, 0));
};


/**
 * @param {number} value
 * @return {!proto.infra.SearchResourcesRequest} returns this
 */
proto.infra.SearchResourcesRequest.prototype.setPageSize = function(value) {
  return jspb.Message.setProto3IntField(this, 14, value);
};


/**
 * optional int32 page_number = 15;
 * @return {number}
 */
proto.infra.SearchResourcesRequest.prototype.getPageNumber = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 15, 0));
};


/**
 * @param {number} value
 * @return {!proto.infra.SearchResourcesRequest} returns this
 */
proto.infra.SearchResourcesRequest.prototype.setPageNumber = function(value) {
  return jspb.Message.setProto3IntField(this, 15, value);
};


/**
 * optional string sort_by = 16;
 * @return {string}
 */
proto.infra.SearchResourcesRequest.prototype.getSortBy = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 16, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.SearchResourcesRequest} returns this
 */
proto.infra.SearchResourcesRequest.prototype.setSortBy = function(value) {
  return jspb.Message.setProto3StringField(this, 16, value);
};


/**
 * optional bool sort_descending = 17;
 * @return {boolean}
 */
proto.infra.SearchResourcesRequest.prototype.getSortDescending = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 17, false));
};


/**
 * @param {boolean} value
 * @return {!proto.infra.SearchResourcesRequest} returns this
 */
proto.infra.SearchResourcesRequest.prototype.setSortDescending = function(value) {
  return jspb.Message.setProto3BooleanField(this, 17, value);
};


/**
 * optional google.protobuf.FieldMask field_mask = 18;
 * @return {?proto.google.protobuf.FieldMask}
 */
proto.infra.SearchResourcesRequest.prototype.getFieldMask = function() {
  return /** @type{?proto.google.protobuf.FieldMask} */ (
    jspb.Message.getWrapperField(this, google_protobuf_field_mask_pb.FieldMask, 18));
};


/**
 * @param {?proto.google.protobuf.FieldMask|undefined} value
 * @return {!proto.infra.SearchResourcesRequest} returns this
*/
proto.infra.SearchResourcesRequest.prototype.setFieldMask = function(value) {
  return jspb.Message.setWrapperField(this, 18, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.infra.SearchResourcesRequest} returns this
 */
proto.infra.SearchResourcesRequest.prototype.clearFieldMask = function() {
  return this.setFieldMask(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.infra.SearchResourcesRequest.prototype.hasFieldMask = function() {
  return jspb.Message.getField(this, 18) != null;
};


/**
 * repeated string resource_types = 19;
 * @return {!Array<string>}
 */
proto.infra.SearchResourcesRequest.prototype.getResourceTypesList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 19));
};


/**
 * @param {!Array<string>} value
 * @return {!proto.infra.SearchResourcesRequest} returns this
 */
proto.infra.SearchResourcesRequest.prototype.setResourceTypesList = function(value) {
  return jspb.Message.setField(this, 19, value || []);
};


/**
 * @param {string} value
 * @param {number=} opt_index
 * @return {!proto.infra.SearchResourcesRequest} returns this
 */
proto.infra.SearchResourcesRequest.prototype.addResourceTypes = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 19, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.infra.SearchResourcesRequest} returns this
 */
proto.infra.SearchResourcesRequest.prototype.clearResourceTypesList = function() {
  return this.setResourceTypesList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.infra.SearchResourcesResponse.repeatedFields_ = [1,2,3,4,5,6,7,8,9,10,11,12];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.infra.SearchResourcesResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.infra.SearchResourcesResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.infra.SearchResourcesResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.SearchResourcesResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    instancesList: jspb.Message.toObjectList(msg.getInstancesList(),
    types_pb.Instance.toObject, includeInstance),
    vpcsList: jspb.Message.toObjectList(msg.getVpcsList(),
    types_pb.VPC.toObject, includeInstance),
    subnetsList: jspb.Message.toObjectList(msg.getSubnetsList(),
    types_pb.Subnet.toObject, includeInstance),
    aclsList: jspb.Message.toObjectList(msg.getAclsList(),
    types_pb.ACL.toObject, includeInstance),
    securityGroupsList: jspb.Message.toObjectList(msg.getSecurityGroupsList(),
    types_pb.SecurityGroup.toObject, includeInstance),
    routeTablesList: jspb.Message.toObjectList(msg.getRouteTablesList(),
    types_pb.RouteTable.toObject, includeInstance),
    natGatewaysList: jspb.Message.toObjectList(msg.getNatGatewaysList(),
    types_pb.NATGateway.toObject, includeInstance),
    routersList: jspb.Message.toObjectList(msg.getRoutersList(),
    types_pb.Router.toObject, includeInstance),
    igwsList: jspb.Message.toObjectList(msg.getIgwsList(),
    types_pb.IGW.toObject, includeInstance),
    vpcEndpointsList: jspb.Message.toObjectList(msg.getVpcEndpointsList(),
    types_pb.VPCEndpoint.toObject, includeInstance),
    publicIpsList: jspb.Message.toObjectList(msg.getPublicIpsList(),
    types_pb.PublicIP.toObject, includeInstance),
    clustersList: jspb.Message.toObjectList(msg.getClustersList(),
    types_k8s_pb.Cluster.toObject, includeInstance),
    totalResults: jspb.Message.getFieldWithDefault(msg, 13, 0),
    totalPages: jspb.Message.getFieldWithDefault(msg, 14, 0),
    currentPage: jspb.Message.getFieldWithDefault(msg, 15, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.infra.SearchResourcesResponse}
 */
proto.infra.SearchResourcesResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.infra.SearchResourcesResponse;
  return proto.infra.SearchResourcesResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.infra.SearchResourcesResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.infra.SearchResourcesResponse}
 */
proto.infra.SearchResourcesResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new types_pb.Instance;
      reader.readMessage(value,types_pb.Instance.deserializeBinaryFromReader);
      msg.addInstances(value);
      break;
    case 2:
      var value = new types_pb.VPC;
      reader.readMessage(value,types_pb.VPC.deserializeBinaryFromReader);
      msg.addVpcs(value);
      break;
    case 3:
      var value = new types_pb.Subnet;
      reader.readMessage(value,types_pb.Subnet.deserializeBinaryFromReader);
      msg.addSubnets(value);
      break;
    case 4:
      var value = new types_pb.ACL;
      reader.readMessage(value,types_pb.ACL.deserializeBinaryFromReader);
      msg.addAcls(value);
      break;
    case 5:
      var value = new types_pb.SecurityGroup;
      reader.readMessage(value,types_pb.SecurityGroup.deserializeBinaryFromReader);
      msg.addSecurityGroups(value);
      break;
    case 6:
      var value = new types_pb.RouteTable;
      reader.readMessage(value,types_pb.RouteTable.deserializeBinaryFromReader);
      msg.addRouteTables(value);
      break;
    case 7:
      var value = new types_pb.NATGateway;
      reader.readMessage(value,types_pb.NATGateway.deserializeBinaryFromReader);
      msg.addNatGateways(value);
      break;
    case 8:
      var value = new types_pb.Router;
      reader.readMessage(value,types_pb.Router.deserializeBinaryFromReader);
      msg.addRouters(value);
      break;
    case 9:
      var value = new types_pb.IGW;
      reader.readMessage(value,types_pb.IGW.deserializeBinaryFromReader);
      msg.addIgws(value);
      break;
    case 10:
      var value = new types_pb.VPCEndpoint;
      reader.readMessage(value,types_pb.VPCEndpoint.deserializeBinaryFromReader);
      msg.addVpcEndpoints(value);
      break;
    case 11:
      var value = new types_pb.PublicIP;
      reader.readMessage(value,types_pb.PublicIP.deserializeBinaryFromReader);
      msg.addPublicIps(value);
      break;
    case 12:
      var value = new types_k8s_pb.Cluster;
      reader.readMessage(value,types_k8s_pb.Cluster.deserializeBinaryFromReader);
      msg.addClusters(value);
      break;
    case 13:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setTotalResults(value);
      break;
    case 14:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setTotalPages(value);
      break;
    case 15:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setCurrentPage(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.infra.SearchResourcesResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.infra.SearchResourcesResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.infra.SearchResourcesResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.SearchResourcesResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getInstancesList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      types_pb.Instance.serializeBinaryToWriter
    );
  }
  f = message.getVpcsList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      2,
      f,
      types_pb.VPC.serializeBinaryToWriter
    );
  }
  f = message.getSubnetsList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      3,
      f,
      types_pb.Subnet.serializeBinaryToWriter
    );
  }
  f = message.getAclsList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      4,
      f,
      types_pb.ACL.serializeBinaryToWriter
    );
  }
  f = message.getSecurityGroupsList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      5,
      f,
      types_pb.SecurityGroup.serializeBinaryToWriter
    );
  }
  f = message.getRouteTablesList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      6,
      f,
      types_pb.RouteTable.serializeBinaryToWriter
    );
  }
  f = message.getNatGatewaysList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      7,
      f,
      types_pb.NATGateway.serializeBinaryToWriter
    );
  }
  f = message.getRoutersList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      8,
      f,
      types_pb.Router.serializeBinaryToWriter
    );
  }
  f = message.getIgwsList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      9,
      f,
      types_pb.IGW.serializeBinaryToWriter
    );
  }
  f = message.getVpcEndpointsList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      10,
      f,
      types_pb.VPCEndpoint.serializeBinaryToWriter
    );
  }
  f = message.getPublicIpsList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      11,
      f,
      types_pb.PublicIP.serializeBinaryToWriter
    );
  }
  f = message.getClustersList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      12,
      f,
      types_k8s_pb.Cluster.serializeBinaryToWriter
    );
  }
  f = message.getTotalResults();
  if (f !== 0) {
    writer.writeInt32(
      13,
      f
    );
  }
  f = message.getTotalPages();
  if (f !== 0) {
    writer.writeInt32(
      14,
      f
    );
  }
  f = message.getCurrentPage();
  if (f !== 0) {
    writer.writeInt32(
      15,
      f
    );
  }
};


/**
 * repeated Instance instances = 1;
 * @return {!Array<!proto.infra.Instance>}
 */
proto.infra.SearchResourcesResponse.prototype.getInstancesList = function() {
  return /** @type{!Array<!proto.infra.Instance>} */ (
    jspb.Message.getRepeatedWrapperField(this, types_pb.Instance, 1));
};


/**
 * @param {!Array<!proto.infra.Instance>} value
 * @return {!proto.infra.SearchResourcesResponse} returns this
*/
proto.infra.SearchResourcesResponse.prototype.setInstancesList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.infra.Instance=} opt_value
 * @param {number=} opt_index
 * @return {!proto.infra.Instance}
 */
proto.infra.SearchResourcesResponse.prototype.addInstances = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.infra.Instance, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.infra.SearchResourcesResponse} returns this
 */
proto.infra.SearchResourcesResponse.prototype.clearInstancesList = function() {
  return this.setInstancesList([]);
};


/**
 * repeated VPC vpcs = 2;
 * @return {!Array<!proto.infra.VPC>}
 */
proto.infra.SearchResourcesResponse.prototype.getVpcsList = function() {
  return /** @type{!Array<!proto.infra.VPC>} */ (
    jspb.Message.getRepeatedWrapperField(this, types_pb.VPC, 2));
};


/**
 * @param {!Array<!proto.infra.VPC>} value
 * @return {!proto.infra.SearchResourcesResponse} returns this
*/
proto.infra.SearchResourcesResponse.prototype.setVpcsList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 2, value);
};


/**
 * @param {!proto.infra.VPC=} opt_value
 * @param {number=} opt_index
 * @return {!proto.infra.VPC}
 */
proto.infra.SearchResourcesResponse.prototype.addVpcs = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 2, opt_value, proto.infra.VPC, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.infra.SearchResourcesResponse} returns this
 */
proto.infra.SearchResourcesResponse.prototype.clearVpcsList = function() {
  return this.setVpcsList([]);
};


/**
 * repeated Subnet subnets = 3;
 * @return {!Array<!proto.infra.Subnet>}
 */
proto.infra.SearchResourcesResponse.prototype.getSubnetsList = function() {
  return /** @type{!Array<!proto.infra.Subnet>} */ (
    jspb.Message.getRepeatedWrapperField(this, types_pb.Subnet, 3));
};


/**
 * @param {!Array<!proto.infra.Subnet>} value
 * @return {!proto.infra.SearchResourcesResponse} returns this
*/
proto.infra.SearchResourcesResponse.prototype.setSubnetsList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 3, value);
};


/**
 * @param {!proto.infra.Subnet=} opt_value
 * @param {number=} opt_index
 * @return {!proto.infra.Subnet}
 */
proto.infra.SearchResourcesResponse.prototype.addSubnets = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 3, opt_value, proto.infra.Subnet, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.infra.SearchResourcesResponse} returns this
 */
proto.infra.SearchResourcesResponse.prototype.clearSubnetsList = function() {
  return this.setSubnetsList([]);
};


/**
 * repeated ACL acls = 4;
 * @return {!Array<!proto.infra.ACL>}
 */
proto.infra.SearchResourcesResponse.prototype.getAclsList = function() {
  return /** @type{!Array<!proto.infra.ACL>} */ (
    jspb.Message.getRepeatedWrapperField(this, types_pb.ACL, 4));
};


/**
 * @param {!Array<!proto.infra.ACL>} value
 * @return {!proto.infra.SearchResourcesResponse} returns this
*/
proto.infra.SearchResourcesResponse.prototype.setAclsList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 4, value);
};


/**
 * @param {!proto.infra.ACL=} opt_value
 * @param {number=} opt_index
 * @return {!proto.infra.ACL}
 */
proto.infra.SearchResourcesResponse.prototype.addAcls = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 4, opt_value, proto.infra.ACL, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.infra.SearchResourcesResponse} returns this
 */
proto.infra.SearchResourcesResponse.prototype.clearAclsList = function() {
  return this.setAclsList([]);
};


/**
 * repeated SecurityGroup security_groups = 5;
 * @return {!Array<!proto.infra.SecurityGroup>}
 */
proto.infra.SearchResourcesResponse.prototype.getSecurityGroupsList = function() {
  return /** @type{!Array<!proto.infra.SecurityGroup>} */ (
    jspb.Message.getRepeatedWrapperField(this, types_pb.SecurityGroup, 5));
};


/**
 * @param {!Array<!proto.infra.SecurityGroup>} value
 * @return {!proto.infra.SearchResourcesResponse} returns this
*/
proto.infra.SearchResourcesResponse.prototype.setSecurityGroupsList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 5, value);
};


/**
 * @param {!proto.infra.SecurityGroup=} opt_value
 * @param {number=} opt_index
 * @return {!proto.infra.SecurityGroup}
 */
proto.infra.SearchResourcesResponse.prototype.addSecurityGroups = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 5, opt_value, proto.infra.SecurityGroup, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.infra.SearchResourcesResponse} returns this
 */
proto.infra.SearchResourcesResponse.prototype.clearSecurityGroupsList = function() {
  return this.setSecurityGroupsList([]);
};


/**
 * repeated RouteTable route_tables = 6;
 * @return {!Array<!proto.infra.RouteTable>}
 */
proto.infra.SearchResourcesResponse.prototype.getRouteTablesList = function() {
  return /** @type{!Array<!proto.infra.RouteTable>} */ (
    jspb.Message.getRepeatedWrapperField(this, types_pb.RouteTable, 6));
};


/**
 * @param {!Array<!proto.infra.RouteTable>} value
 * @return {!proto.infra.SearchResourcesResponse} returns this
*/
proto.infra.SearchResourcesResponse.prototype.setRouteTablesList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 6, value);
};


/**
 * @param {!proto.infra.RouteTable=} opt_value
 * @param {number=} opt_index
 * @return {!proto.infra.RouteTable}
 */
proto.infra.SearchResourcesResponse.prototype.addRouteTables = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 6, opt_value, proto.infra.RouteTable, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.infra.SearchResourcesResponse} returns this
 */
proto.infra.SearchResourcesResponse.prototype.clearRouteTablesList = function() {
  return this.setRouteTablesList([]);
};


/**
 * repeated NATGateway nat_gateways = 7;
 * @return {!Array<!proto.infra.NATGateway>}
 */
proto.infra.SearchResourcesResponse.prototype.getNatGatewaysList = function() {
  return /** @type{!Array<!proto.infra.NATGateway>} */ (
    jspb.Message.getRepeatedWrapperField(this, types_pb.NATGateway, 7));
};


/**
 * @param {!Array<!proto.infra.NATGateway>} value
 * @return {!proto.infra.SearchResourcesResponse} returns this
*/
proto.infra.SearchResourcesResponse.prototype.setNatGatewaysList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 7, value);
};


/**
 * @param {!proto.infra.NATGateway=} opt_value
 * @param {number=} opt_index
 * @return {!proto.infra.NATGateway}
 */
proto.infra.SearchResourcesResponse.prototype.addNatGateways = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 7, opt_value, proto.infra.NATGateway, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.infra.SearchResourcesResponse} returns this
 */
proto.infra.SearchResourcesResponse.prototype.clearNatGatewaysList = function() {
  return this.setNatGatewaysList([]);
};


/**
 * repeated Router routers = 8;
 * @return {!Array<!proto.infra.Router>}
 */
proto.infra.SearchResourcesResponse.prototype.getRoutersList = function() {
  return /** @type{!Array<!proto.infra.Router>} */ (
    jspb.Message.getRepeatedWrapperField(this, types_pb.Router, 8));
};


/**
 * @param {!Array<!proto.infra.Router>} value
 * @return {!proto.infra.SearchResourcesResponse} returns this
*/
proto.infra.SearchResourcesResponse.prototype.setRoutersList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 8, value);
};


/**
 * @param {!proto.infra.Router=} opt_value
 * @param {number=} opt_index
 * @return {!proto.infra.Router}
 */
proto.infra.SearchResourcesResponse.prototype.addRouters = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 8, opt_value, proto.infra.Router, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.infra.SearchResourcesResponse} returns this
 */
proto.infra.SearchResourcesResponse.prototype.clearRoutersList = function() {
  return this.setRoutersList([]);
};


/**
 * repeated IGW igws = 9;
 * @return {!Array<!proto.infra.IGW>}
 */
proto.infra.SearchResourcesResponse.prototype.getIgwsList = function() {
  return /** @type{!Array<!proto.infra.IGW>} */ (
    jspb.Message.getRepeatedWrapperField(this, types_pb.IGW, 9));
};


/**
 * @param {!Array<!proto.infra.IGW>} value
 * @return {!proto.infra.SearchResourcesResponse} returns this
*/
proto.infra.SearchResourcesResponse.prototype.setIgwsList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 9, value);
};


/**
 * @param {!proto.infra.IGW=} opt_value
 * @param {number=} opt_index
 * @return {!proto.infra.IGW}
 */
proto.infra.SearchResourcesResponse.prototype.addIgws = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 9, opt_value, proto.infra.IGW, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.infra.SearchResourcesResponse} returns this
 */
proto.infra.SearchResourcesResponse.prototype.clearIgwsList = function() {
  return this.setIgwsList([]);
};


/**
 * repeated VPCEndpoint vpc_endpoints = 10;
 * @return {!Array<!proto.infra.VPCEndpoint>}
 */
proto.infra.SearchResourcesResponse.prototype.getVpcEndpointsList = function() {
  return /** @type{!Array<!proto.infra.VPCEndpoint>} */ (
    jspb.Message.getRepeatedWrapperField(this, types_pb.VPCEndpoint, 10));
};


/**
 * @param {!Array<!proto.infra.VPCEndpoint>} value
 * @return {!proto.infra.SearchResourcesResponse} returns this
*/
proto.infra.SearchResourcesResponse.prototype.setVpcEndpointsList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 10, value);
};


/**
 * @param {!proto.infra.VPCEndpoint=} opt_value
 * @param {number=} opt_index
 * @return {!proto.infra.VPCEndpoint}
 */
proto.infra.SearchResourcesResponse.prototype.addVpcEndpoints = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 10, opt_value, proto.infra.VPCEndpoint, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.infra.SearchResourcesResponse} returns this
 */
proto.infra.SearchResourcesResponse.prototype.clearVpcEndpointsList = function() {
  return this.setVpcEndpointsList([]);
};


/**
 * repeated PublicIP public_ips = 11;
 * @return {!Array<!proto.infra.PublicIP>}
 */
proto.infra.SearchResourcesResponse.prototype.getPublicIpsList = function() {
  return /** @type{!Array<!proto.infra.PublicIP>} */ (
    jspb.Message.getRepeatedWrapperField(this, types_pb.PublicIP, 11));
};


/**
 * @param {!Array<!proto.infra.PublicIP>} value
 * @return {!proto.infra.SearchResourcesResponse} returns this
*/
proto.infra.SearchResourcesResponse.prototype.setPublicIpsList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 11, value);
};


/**
 * @param {!proto.infra.PublicIP=} opt_value
 * @param {number=} opt_index
 * @return {!proto.infra.PublicIP}
 */
proto.infra.SearchResourcesResponse.prototype.addPublicIps = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 11, opt_value, proto.infra.PublicIP, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.infra.SearchResourcesResponse} returns this
 */
proto.infra.SearchResourcesResponse.prototype.clearPublicIpsList = function() {
  return this.setPublicIpsList([]);
};


/**
 * repeated Cluster clusters = 12;
 * @return {!Array<!proto.infra.Cluster>}
 */
proto.infra.SearchResourcesResponse.prototype.getClustersList = function() {
  return /** @type{!Array<!proto.infra.Cluster>} */ (
    jspb.Message.getRepeatedWrapperField(this, types_k8s_pb.Cluster, 12));
};


/**
 * @param {!Array<!proto.infra.Cluster>} value
 * @return {!proto.infra.SearchResourcesResponse} returns this
*/
proto.infra.SearchResourcesResponse.prototype.setClustersList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 12, value);
};


/**
 * @param {!proto.infra.Cluster=} opt_value
 * @param {number=} opt_index
 * @return {!proto.infra.Cluster}
 */
proto.infra.SearchResourcesResponse.prototype.addClusters = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 12, opt_value, proto.infra.Cluster, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.infra.SearchResourcesResponse} returns this
 */
proto.infra.SearchResourcesResponse.prototype.clearClustersList = function() {
  return this.setClustersList([]);
};


/**
 * optional int32 total_results = 13;
 * @return {number}
 */
proto.infra.SearchResourcesResponse.prototype.getTotalResults = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 13, 0));
};


/**
 * @param {number} value
 * @return {!proto.infra.SearchResourcesResponse} returns this
 */
proto.infra.SearchResourcesResponse.prototype.setTotalResults = function(value) {
  return jspb.Message.setProto3IntField(this, 13, value);
};


/**
 * optional int32 total_pages = 14;
 * @return {number}
 */
proto.infra.SearchResourcesResponse.prototype.getTotalPages = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 14, 0));
};


/**
 * @param {number} value
 * @return {!proto.infra.SearchResourcesResponse} returns this
 */
proto.infra.SearchResourcesResponse.prototype.setTotalPages = function(value) {
  return jspb.Message.setProto3IntField(this, 14, value);
};


/**
 * optional int32 current_page = 15;
 * @return {number}
 */
proto.infra.SearchResourcesResponse.prototype.getCurrentPage = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 15, 0));
};


/**
 * @param {number} value
 * @return {!proto.infra.SearchResourcesResponse} returns this
 */
proto.infra.SearchResourcesResponse.prototype.setCurrentPage = function(value) {
  return jspb.Message.setProto3IntField(this, 15, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.infra.GetVpcConnectivityGraphRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.infra.GetVpcConnectivityGraphRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.infra.GetVpcConnectivityGraphRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.GetVpcConnectivityGraphRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    provider: jspb.Message.getFieldWithDefault(msg, 1, ""),
    accountId: jspb.Message.getFieldWithDefault(msg, 2, ""),
    region: jspb.Message.getFieldWithDefault(msg, 3, ""),
    vpcId: jspb.Message.getFieldWithDefault(msg, 4, ""),
    creds: (f = msg.getCreds()) && types_pb.Credentials.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.infra.GetVpcConnectivityGraphRequest}
 */
proto.infra.GetVpcConnectivityGraphRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.infra.GetVpcConnectivityGraphRequest;
  return proto.infra.GetVpcConnectivityGraphRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.infra.GetVpcConnectivityGraphRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.infra.GetVpcConnectivityGraphRequest}
 */
proto.infra.GetVpcConnectivityGraphRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setProvider(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setAccountId(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setRegion(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setVpcId(value);
      break;
    case 7:
      var value = new types_pb.Credentials;
      reader.readMessage(value,types_pb.Credentials.deserializeBinaryFromReader);
      msg.setCreds(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.infra.GetVpcConnectivityGraphRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.infra.GetVpcConnectivityGraphRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.infra.GetVpcConnectivityGraphRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.GetVpcConnectivityGraphRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getProvider();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getAccountId();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getRegion();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getVpcId();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getCreds();
  if (f != null) {
    writer.writeMessage(
      7,
      f,
      types_pb.Credentials.serializeBinaryToWriter
    );
  }
};


/**
 * optional string provider = 1;
 * @return {string}
 */
proto.infra.GetVpcConnectivityGraphRequest.prototype.getProvider = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.GetVpcConnectivityGraphRequest} returns this
 */
proto.infra.GetVpcConnectivityGraphRequest.prototype.setProvider = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string account_id = 2;
 * @return {string}
 */
proto.infra.GetVpcConnectivityGraphRequest.prototype.getAccountId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.GetVpcConnectivityGraphRequest} returns this
 */
proto.infra.GetVpcConnectivityGraphRequest.prototype.setAccountId = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string region = 3;
 * @return {string}
 */
proto.infra.GetVpcConnectivityGraphRequest.prototype.getRegion = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.GetVpcConnectivityGraphRequest} returns this
 */
proto.infra.GetVpcConnectivityGraphRequest.prototype.setRegion = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string vpc_id = 4;
 * @return {string}
 */
proto.infra.GetVpcConnectivityGraphRequest.prototype.getVpcId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.GetVpcConnectivityGraphRequest} returns this
 */
proto.infra.GetVpcConnectivityGraphRequest.prototype.setVpcId = function(value) {
  return jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * optional Credentials creds = 7;
 * @return {?proto.infra.Credentials}
 */
proto.infra.GetVpcConnectivityGraphRequest.prototype.getCreds = function() {
  return /** @type{?proto.infra.Credentials} */ (
    jspb.Message.getWrapperField(this, types_pb.Credentials, 7));
};


/**
 * @param {?proto.infra.Credentials|undefined} value
 * @return {!proto.infra.GetVpcConnectivityGraphRequest} returns this
*/
proto.infra.GetVpcConnectivityGraphRequest.prototype.setCreds = function(value) {
  return jspb.Message.setWrapperField(this, 7, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.infra.GetVpcConnectivityGraphRequest} returns this
 */
proto.infra.GetVpcConnectivityGraphRequest.prototype.clearCreds = function() {
  return this.setCreds(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.infra.GetVpcConnectivityGraphRequest.prototype.hasCreds = function() {
  return jspb.Message.getField(this, 7) != null;
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.infra.GetVpcConnectivityGraphResponse.repeatedFields_ = [1,2];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.infra.GetVpcConnectivityGraphResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.infra.GetVpcConnectivityGraphResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.infra.GetVpcConnectivityGraphResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.GetVpcConnectivityGraphResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    nodesList: jspb.Message.toObjectList(msg.getNodesList(),
    types_pb.VpcGraphNode.toObject, includeInstance),
    edgesList: jspb.Message.toObjectList(msg.getEdgesList(),
    types_pb.VpcGraphEdge.toObject, includeInstance),
    err: (f = msg.getErr()) && types_pb.Error.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.infra.GetVpcConnectivityGraphResponse}
 */
proto.infra.GetVpcConnectivityGraphResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.infra.GetVpcConnectivityGraphResponse;
  return proto.infra.GetVpcConnectivityGraphResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.infra.GetVpcConnectivityGraphResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.infra.GetVpcConnectivityGraphResponse}
 */
proto.infra.GetVpcConnectivityGraphResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new types_pb.VpcGraphNode;
      reader.readMessage(value,types_pb.VpcGraphNode.deserializeBinaryFromReader);
      msg.addNodes(value);
      break;
    case 2:
      var value = new types_pb.VpcGraphEdge;
      reader.readMessage(value,types_pb.VpcGraphEdge.deserializeBinaryFromReader);
      msg.addEdges(value);
      break;
    case 3:
      var value = new types_pb.Error;
      reader.readMessage(value,types_pb.Error.deserializeBinaryFromReader);
      msg.setErr(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.infra.GetVpcConnectivityGraphResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.infra.GetVpcConnectivityGraphResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.infra.GetVpcConnectivityGraphResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.GetVpcConnectivityGraphResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getNodesList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      types_pb.VpcGraphNode.serializeBinaryToWriter
    );
  }
  f = message.getEdgesList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      2,
      f,
      types_pb.VpcGraphEdge.serializeBinaryToWriter
    );
  }
  f = message.getErr();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      types_pb.Error.serializeBinaryToWriter
    );
  }
};


/**
 * repeated VpcGraphNode nodes = 1;
 * @return {!Array<!proto.infra.VpcGraphNode>}
 */
proto.infra.GetVpcConnectivityGraphResponse.prototype.getNodesList = function() {
  return /** @type{!Array<!proto.infra.VpcGraphNode>} */ (
    jspb.Message.getRepeatedWrapperField(this, types_pb.VpcGraphNode, 1));
};


/**
 * @param {!Array<!proto.infra.VpcGraphNode>} value
 * @return {!proto.infra.GetVpcConnectivityGraphResponse} returns this
*/
proto.infra.GetVpcConnectivityGraphResponse.prototype.setNodesList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.infra.VpcGraphNode=} opt_value
 * @param {number=} opt_index
 * @return {!proto.infra.VpcGraphNode}
 */
proto.infra.GetVpcConnectivityGraphResponse.prototype.addNodes = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.infra.VpcGraphNode, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.infra.GetVpcConnectivityGraphResponse} returns this
 */
proto.infra.GetVpcConnectivityGraphResponse.prototype.clearNodesList = function() {
  return this.setNodesList([]);
};


/**
 * repeated VpcGraphEdge edges = 2;
 * @return {!Array<!proto.infra.VpcGraphEdge>}
 */
proto.infra.GetVpcConnectivityGraphResponse.prototype.getEdgesList = function() {
  return /** @type{!Array<!proto.infra.VpcGraphEdge>} */ (
    jspb.Message.getRepeatedWrapperField(this, types_pb.VpcGraphEdge, 2));
};


/**
 * @param {!Array<!proto.infra.VpcGraphEdge>} value
 * @return {!proto.infra.GetVpcConnectivityGraphResponse} returns this
*/
proto.infra.GetVpcConnectivityGraphResponse.prototype.setEdgesList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 2, value);
};


/**
 * @param {!proto.infra.VpcGraphEdge=} opt_value
 * @param {number=} opt_index
 * @return {!proto.infra.VpcGraphEdge}
 */
proto.infra.GetVpcConnectivityGraphResponse.prototype.addEdges = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 2, opt_value, proto.infra.VpcGraphEdge, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.infra.GetVpcConnectivityGraphResponse} returns this
 */
proto.infra.GetVpcConnectivityGraphResponse.prototype.clearEdgesList = function() {
  return this.setEdgesList([]);
};


/**
 * optional Error err = 3;
 * @return {?proto.infra.Error}
 */
proto.infra.GetVpcConnectivityGraphResponse.prototype.getErr = function() {
  return /** @type{?proto.infra.Error} */ (
    jspb.Message.getWrapperField(this, types_pb.Error, 3));
};


/**
 * @param {?proto.infra.Error|undefined} value
 * @return {!proto.infra.GetVpcConnectivityGraphResponse} returns this
*/
proto.infra.GetVpcConnectivityGraphResponse.prototype.setErr = function(value) {
  return jspb.Message.setWrapperField(this, 3, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.infra.GetVpcConnectivityGraphResponse} returns this
 */
proto.infra.GetVpcConnectivityGraphResponse.prototype.clearErr = function() {
  return this.setErr(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.infra.GetVpcConnectivityGraphResponse.prototype.hasErr = function() {
  return jspb.Message.getField(this, 3) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.infra.ListVpcGraphNodesRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.infra.ListVpcGraphNodesRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.infra.ListVpcGraphNodesRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListVpcGraphNodesRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    provider: jspb.Message.getFieldWithDefault(msg, 1, ""),
    accountId: jspb.Message.getFieldWithDefault(msg, 2, ""),
    region: jspb.Message.getFieldWithDefault(msg, 3, ""),
    vpcId: jspb.Message.getFieldWithDefault(msg, 4, ""),
    creds: (f = msg.getCreds()) && types_pb.Credentials.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.infra.ListVpcGraphNodesRequest}
 */
proto.infra.ListVpcGraphNodesRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.infra.ListVpcGraphNodesRequest;
  return proto.infra.ListVpcGraphNodesRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.infra.ListVpcGraphNodesRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.infra.ListVpcGraphNodesRequest}
 */
proto.infra.ListVpcGraphNodesRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setProvider(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setAccountId(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setRegion(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setVpcId(value);
      break;
    case 6:
      var value = new types_pb.Credentials;
      reader.readMessage(value,types_pb.Credentials.deserializeBinaryFromReader);
      msg.setCreds(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.infra.ListVpcGraphNodesRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.infra.ListVpcGraphNodesRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.infra.ListVpcGraphNodesRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListVpcGraphNodesRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getProvider();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getAccountId();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getRegion();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getVpcId();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getCreds();
  if (f != null) {
    writer.writeMessage(
      6,
      f,
      types_pb.Credentials.serializeBinaryToWriter
    );
  }
};


/**
 * optional string provider = 1;
 * @return {string}
 */
proto.infra.ListVpcGraphNodesRequest.prototype.getProvider = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListVpcGraphNodesRequest} returns this
 */
proto.infra.ListVpcGraphNodesRequest.prototype.setProvider = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string account_id = 2;
 * @return {string}
 */
proto.infra.ListVpcGraphNodesRequest.prototype.getAccountId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListVpcGraphNodesRequest} returns this
 */
proto.infra.ListVpcGraphNodesRequest.prototype.setAccountId = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string region = 3;
 * @return {string}
 */
proto.infra.ListVpcGraphNodesRequest.prototype.getRegion = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListVpcGraphNodesRequest} returns this
 */
proto.infra.ListVpcGraphNodesRequest.prototype.setRegion = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string vpc_id = 4;
 * @return {string}
 */
proto.infra.ListVpcGraphNodesRequest.prototype.getVpcId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListVpcGraphNodesRequest} returns this
 */
proto.infra.ListVpcGraphNodesRequest.prototype.setVpcId = function(value) {
  return jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * optional Credentials creds = 6;
 * @return {?proto.infra.Credentials}
 */
proto.infra.ListVpcGraphNodesRequest.prototype.getCreds = function() {
  return /** @type{?proto.infra.Credentials} */ (
    jspb.Message.getWrapperField(this, types_pb.Credentials, 6));
};


/**
 * @param {?proto.infra.Credentials|undefined} value
 * @return {!proto.infra.ListVpcGraphNodesRequest} returns this
*/
proto.infra.ListVpcGraphNodesRequest.prototype.setCreds = function(value) {
  return jspb.Message.setWrapperField(this, 6, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.infra.ListVpcGraphNodesRequest} returns this
 */
proto.infra.ListVpcGraphNodesRequest.prototype.clearCreds = function() {
  return this.setCreds(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.infra.ListVpcGraphNodesRequest.prototype.hasCreds = function() {
  return jspb.Message.getField(this, 6) != null;
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.infra.ListVpcGraphNodesResponse.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.infra.ListVpcGraphNodesResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.infra.ListVpcGraphNodesResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.infra.ListVpcGraphNodesResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListVpcGraphNodesResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    nodesList: jspb.Message.toObjectList(msg.getNodesList(),
    types_pb.VpcGraphNode.toObject, includeInstance),
    err: (f = msg.getErr()) && types_pb.Error.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.infra.ListVpcGraphNodesResponse}
 */
proto.infra.ListVpcGraphNodesResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.infra.ListVpcGraphNodesResponse;
  return proto.infra.ListVpcGraphNodesResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.infra.ListVpcGraphNodesResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.infra.ListVpcGraphNodesResponse}
 */
proto.infra.ListVpcGraphNodesResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new types_pb.VpcGraphNode;
      reader.readMessage(value,types_pb.VpcGraphNode.deserializeBinaryFromReader);
      msg.addNodes(value);
      break;
    case 2:
      var value = new types_pb.Error;
      reader.readMessage(value,types_pb.Error.deserializeBinaryFromReader);
      msg.setErr(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.infra.ListVpcGraphNodesResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.infra.ListVpcGraphNodesResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.infra.ListVpcGraphNodesResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListVpcGraphNodesResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getNodesList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      types_pb.VpcGraphNode.serializeBinaryToWriter
    );
  }
  f = message.getErr();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      types_pb.Error.serializeBinaryToWriter
    );
  }
};


/**
 * repeated VpcGraphNode nodes = 1;
 * @return {!Array<!proto.infra.VpcGraphNode>}
 */
proto.infra.ListVpcGraphNodesResponse.prototype.getNodesList = function() {
  return /** @type{!Array<!proto.infra.VpcGraphNode>} */ (
    jspb.Message.getRepeatedWrapperField(this, types_pb.VpcGraphNode, 1));
};


/**
 * @param {!Array<!proto.infra.VpcGraphNode>} value
 * @return {!proto.infra.ListVpcGraphNodesResponse} returns this
*/
proto.infra.ListVpcGraphNodesResponse.prototype.setNodesList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.infra.VpcGraphNode=} opt_value
 * @param {number=} opt_index
 * @return {!proto.infra.VpcGraphNode}
 */
proto.infra.ListVpcGraphNodesResponse.prototype.addNodes = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.infra.VpcGraphNode, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.infra.ListVpcGraphNodesResponse} returns this
 */
proto.infra.ListVpcGraphNodesResponse.prototype.clearNodesList = function() {
  return this.setNodesList([]);
};


/**
 * optional Error err = 2;
 * @return {?proto.infra.Error}
 */
proto.infra.ListVpcGraphNodesResponse.prototype.getErr = function() {
  return /** @type{?proto.infra.Error} */ (
    jspb.Message.getWrapperField(this, types_pb.Error, 2));
};


/**
 * @param {?proto.infra.Error|undefined} value
 * @return {!proto.infra.ListVpcGraphNodesResponse} returns this
*/
proto.infra.ListVpcGraphNodesResponse.prototype.setErr = function(value) {
  return jspb.Message.setWrapperField(this, 2, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.infra.ListVpcGraphNodesResponse} returns this
 */
proto.infra.ListVpcGraphNodesResponse.prototype.clearErr = function() {
  return this.setErr(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.infra.ListVpcGraphNodesResponse.prototype.hasErr = function() {
  return jspb.Message.getField(this, 2) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.infra.ListVpcGraphEdgesRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.infra.ListVpcGraphEdgesRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.infra.ListVpcGraphEdgesRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListVpcGraphEdgesRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    provider: jspb.Message.getFieldWithDefault(msg, 1, ""),
    accountId: jspb.Message.getFieldWithDefault(msg, 2, ""),
    region: jspb.Message.getFieldWithDefault(msg, 3, ""),
    vpcId: jspb.Message.getFieldWithDefault(msg, 4, ""),
    creds: (f = msg.getCreds()) && types_pb.Credentials.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.infra.ListVpcGraphEdgesRequest}
 */
proto.infra.ListVpcGraphEdgesRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.infra.ListVpcGraphEdgesRequest;
  return proto.infra.ListVpcGraphEdgesRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.infra.ListVpcGraphEdgesRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.infra.ListVpcGraphEdgesRequest}
 */
proto.infra.ListVpcGraphEdgesRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setProvider(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setAccountId(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setRegion(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setVpcId(value);
      break;
    case 6:
      var value = new types_pb.Credentials;
      reader.readMessage(value,types_pb.Credentials.deserializeBinaryFromReader);
      msg.setCreds(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.infra.ListVpcGraphEdgesRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.infra.ListVpcGraphEdgesRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.infra.ListVpcGraphEdgesRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListVpcGraphEdgesRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getProvider();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getAccountId();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getRegion();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getVpcId();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getCreds();
  if (f != null) {
    writer.writeMessage(
      6,
      f,
      types_pb.Credentials.serializeBinaryToWriter
    );
  }
};


/**
 * optional string provider = 1;
 * @return {string}
 */
proto.infra.ListVpcGraphEdgesRequest.prototype.getProvider = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListVpcGraphEdgesRequest} returns this
 */
proto.infra.ListVpcGraphEdgesRequest.prototype.setProvider = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string account_id = 2;
 * @return {string}
 */
proto.infra.ListVpcGraphEdgesRequest.prototype.getAccountId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListVpcGraphEdgesRequest} returns this
 */
proto.infra.ListVpcGraphEdgesRequest.prototype.setAccountId = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string region = 3;
 * @return {string}
 */
proto.infra.ListVpcGraphEdgesRequest.prototype.getRegion = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListVpcGraphEdgesRequest} returns this
 */
proto.infra.ListVpcGraphEdgesRequest.prototype.setRegion = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string vpc_id = 4;
 * @return {string}
 */
proto.infra.ListVpcGraphEdgesRequest.prototype.getVpcId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * @param {string} value
 * @return {!proto.infra.ListVpcGraphEdgesRequest} returns this
 */
proto.infra.ListVpcGraphEdgesRequest.prototype.setVpcId = function(value) {
  return jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * optional Credentials creds = 6;
 * @return {?proto.infra.Credentials}
 */
proto.infra.ListVpcGraphEdgesRequest.prototype.getCreds = function() {
  return /** @type{?proto.infra.Credentials} */ (
    jspb.Message.getWrapperField(this, types_pb.Credentials, 6));
};


/**
 * @param {?proto.infra.Credentials|undefined} value
 * @return {!proto.infra.ListVpcGraphEdgesRequest} returns this
*/
proto.infra.ListVpcGraphEdgesRequest.prototype.setCreds = function(value) {
  return jspb.Message.setWrapperField(this, 6, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.infra.ListVpcGraphEdgesRequest} returns this
 */
proto.infra.ListVpcGraphEdgesRequest.prototype.clearCreds = function() {
  return this.setCreds(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.infra.ListVpcGraphEdgesRequest.prototype.hasCreds = function() {
  return jspb.Message.getField(this, 6) != null;
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.infra.ListVpcGraphEdgesResponse.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.infra.ListVpcGraphEdgesResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.infra.ListVpcGraphEdgesResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.infra.ListVpcGraphEdgesResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListVpcGraphEdgesResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    edgesList: jspb.Message.toObjectList(msg.getEdgesList(),
    types_pb.VpcGraphEdge.toObject, includeInstance),
    err: (f = msg.getErr()) && types_pb.Error.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.infra.ListVpcGraphEdgesResponse}
 */
proto.infra.ListVpcGraphEdgesResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.infra.ListVpcGraphEdgesResponse;
  return proto.infra.ListVpcGraphEdgesResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.infra.ListVpcGraphEdgesResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.infra.ListVpcGraphEdgesResponse}
 */
proto.infra.ListVpcGraphEdgesResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new types_pb.VpcGraphEdge;
      reader.readMessage(value,types_pb.VpcGraphEdge.deserializeBinaryFromReader);
      msg.addEdges(value);
      break;
    case 2:
      var value = new types_pb.Error;
      reader.readMessage(value,types_pb.Error.deserializeBinaryFromReader);
      msg.setErr(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.infra.ListVpcGraphEdgesResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.infra.ListVpcGraphEdgesResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.infra.ListVpcGraphEdgesResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infra.ListVpcGraphEdgesResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getEdgesList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      types_pb.VpcGraphEdge.serializeBinaryToWriter
    );
  }
  f = message.getErr();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      types_pb.Error.serializeBinaryToWriter
    );
  }
};


/**
 * repeated VpcGraphEdge edges = 1;
 * @return {!Array<!proto.infra.VpcGraphEdge>}
 */
proto.infra.ListVpcGraphEdgesResponse.prototype.getEdgesList = function() {
  return /** @type{!Array<!proto.infra.VpcGraphEdge>} */ (
    jspb.Message.getRepeatedWrapperField(this, types_pb.VpcGraphEdge, 1));
};


/**
 * @param {!Array<!proto.infra.VpcGraphEdge>} value
 * @return {!proto.infra.ListVpcGraphEdgesResponse} returns this
*/
proto.infra.ListVpcGraphEdgesResponse.prototype.setEdgesList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.infra.VpcGraphEdge=} opt_value
 * @param {number=} opt_index
 * @return {!proto.infra.VpcGraphEdge}
 */
proto.infra.ListVpcGraphEdgesResponse.prototype.addEdges = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.infra.VpcGraphEdge, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.infra.ListVpcGraphEdgesResponse} returns this
 */
proto.infra.ListVpcGraphEdgesResponse.prototype.clearEdgesList = function() {
  return this.setEdgesList([]);
};


/**
 * optional Error err = 2;
 * @return {?proto.infra.Error}
 */
proto.infra.ListVpcGraphEdgesResponse.prototype.getErr = function() {
  return /** @type{?proto.infra.Error} */ (
    jspb.Message.getWrapperField(this, types_pb.Error, 2));
};


/**
 * @param {?proto.infra.Error|undefined} value
 * @return {!proto.infra.ListVpcGraphEdgesResponse} returns this
*/
proto.infra.ListVpcGraphEdgesResponse.prototype.setErr = function(value) {
  return jspb.Message.setWrapperField(this, 2, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.infra.ListVpcGraphEdgesResponse} returns this
 */
proto.infra.ListVpcGraphEdgesResponse.prototype.clearErr = function() {
  return this.setErr(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.infra.ListVpcGraphEdgesResponse.prototype.hasErr = function() {
  return jspb.Message.getField(this, 2) != null;
};


goog.object.extend(exports, proto.infra);
