/**
 * Copyright (c) 2024 Cisco Systems, Inc. and its affiliates
 * All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http:www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

import * as jspb from 'google-protobuf'

import * as types_pb from './types_pb';


export class Ports extends jspb.Message {
  getPortsList(): Array<string>;
  setPortsList(value: Array<string>): Ports;
  clearPortsList(): Ports;
  addPorts(value: string, index?: number): Ports;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Ports.AsObject;
  static toObject(includeInstance: boolean, msg: Ports): Ports.AsObject;
  static serializeBinaryToWriter(message: Ports, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Ports;
  static deserializeBinaryFromReader(message: Ports, reader: jspb.BinaryReader): Ports;
}

export namespace Ports {
  export type AsObject = {
    portsList: Array<string>,
  }
}

export class AddInboundAllowRuleInVPCRequest extends jspb.Message {
  getProvider(): string;
  setProvider(value: string): AddInboundAllowRuleInVPCRequest;

  getDestinationVpcId(): string;
  setDestinationVpcId(value: string): AddInboundAllowRuleInVPCRequest;

  getCidrsToAllowList(): Array<string>;
  setCidrsToAllowList(value: Array<string>): AddInboundAllowRuleInVPCRequest;
  clearCidrsToAllowList(): AddInboundAllowRuleInVPCRequest;
  addCidrsToAllow(value: string, index?: number): AddInboundAllowRuleInVPCRequest;

  getRuleName(): string;
  setRuleName(value: string): AddInboundAllowRuleInVPCRequest;

  getTagsMap(): jspb.Map<string, string>;
  clearTagsMap(): AddInboundAllowRuleInVPCRequest;

  getRegion(): string;
  setRegion(value: string): AddInboundAllowRuleInVPCRequest;

  getAccountId(): string;
  setAccountId(value: string): AddInboundAllowRuleInVPCRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AddInboundAllowRuleInVPCRequest.AsObject;
  static toObject(includeInstance: boolean, msg: AddInboundAllowRuleInVPCRequest): AddInboundAllowRuleInVPCRequest.AsObject;
  static serializeBinaryToWriter(message: AddInboundAllowRuleInVPCRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AddInboundAllowRuleInVPCRequest;
  static deserializeBinaryFromReader(message: AddInboundAllowRuleInVPCRequest, reader: jspb.BinaryReader): AddInboundAllowRuleInVPCRequest;
}

export namespace AddInboundAllowRuleInVPCRequest {
  export type AsObject = {
    provider: string,
    destinationVpcId: string,
    cidrsToAllowList: Array<string>,
    ruleName: string,
    tagsMap: Array<[string, string]>,
    region: string,
    accountId: string,
  }
}

export class AddInboundAllowRuleInVPCResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AddInboundAllowRuleInVPCResponse.AsObject;
  static toObject(includeInstance: boolean, msg: AddInboundAllowRuleInVPCResponse): AddInboundAllowRuleInVPCResponse.AsObject;
  static serializeBinaryToWriter(message: AddInboundAllowRuleInVPCResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AddInboundAllowRuleInVPCResponse;
  static deserializeBinaryFromReader(message: AddInboundAllowRuleInVPCResponse, reader: jspb.BinaryReader): AddInboundAllowRuleInVPCResponse;
}

export namespace AddInboundAllowRuleInVPCResponse {
  export type AsObject = {
  }
}

export class AddInboundAllowRuleByLabelsMatchRequest extends jspb.Message {
  getProvider(): string;
  setProvider(value: string): AddInboundAllowRuleByLabelsMatchRequest;

  getVpcId(): string;
  setVpcId(value: string): AddInboundAllowRuleByLabelsMatchRequest;

  getRuleName(): string;
  setRuleName(value: string): AddInboundAllowRuleByLabelsMatchRequest;

  getLabelsMap(): jspb.Map<string, string>;
  clearLabelsMap(): AddInboundAllowRuleByLabelsMatchRequest;

  getCidrsToAllowList(): Array<string>;
  setCidrsToAllowList(value: Array<string>): AddInboundAllowRuleByLabelsMatchRequest;
  clearCidrsToAllowList(): AddInboundAllowRuleByLabelsMatchRequest;
  addCidrsToAllow(value: string, index?: number): AddInboundAllowRuleByLabelsMatchRequest;

  getProtocolsAndPortsMap(): jspb.Map<string, Ports>;
  clearProtocolsAndPortsMap(): AddInboundAllowRuleByLabelsMatchRequest;

  getRegion(): string;
  setRegion(value: string): AddInboundAllowRuleByLabelsMatchRequest;

  getAccountId(): string;
  setAccountId(value: string): AddInboundAllowRuleByLabelsMatchRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AddInboundAllowRuleByLabelsMatchRequest.AsObject;
  static toObject(includeInstance: boolean, msg: AddInboundAllowRuleByLabelsMatchRequest): AddInboundAllowRuleByLabelsMatchRequest.AsObject;
  static serializeBinaryToWriter(message: AddInboundAllowRuleByLabelsMatchRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AddInboundAllowRuleByLabelsMatchRequest;
  static deserializeBinaryFromReader(message: AddInboundAllowRuleByLabelsMatchRequest, reader: jspb.BinaryReader): AddInboundAllowRuleByLabelsMatchRequest;
}

export namespace AddInboundAllowRuleByLabelsMatchRequest {
  export type AsObject = {
    provider: string,
    vpcId: string,
    ruleName: string,
    labelsMap: Array<[string, string]>,
    cidrsToAllowList: Array<string>,
    protocolsAndPortsMap: Array<[string, Ports.AsObject]>,
    region: string,
    accountId: string,
  }
}

export class AddInboundAllowRuleByLabelsMatchResponse extends jspb.Message {
  getRuleId(): string;
  setRuleId(value: string): AddInboundAllowRuleByLabelsMatchResponse;

  getInstancesList(): Array<types_pb.Instance>;
  setInstancesList(value: Array<types_pb.Instance>): AddInboundAllowRuleByLabelsMatchResponse;
  clearInstancesList(): AddInboundAllowRuleByLabelsMatchResponse;
  addInstances(value?: types_pb.Instance, index?: number): types_pb.Instance;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AddInboundAllowRuleByLabelsMatchResponse.AsObject;
  static toObject(includeInstance: boolean, msg: AddInboundAllowRuleByLabelsMatchResponse): AddInboundAllowRuleByLabelsMatchResponse.AsObject;
  static serializeBinaryToWriter(message: AddInboundAllowRuleByLabelsMatchResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AddInboundAllowRuleByLabelsMatchResponse;
  static deserializeBinaryFromReader(message: AddInboundAllowRuleByLabelsMatchResponse, reader: jspb.BinaryReader): AddInboundAllowRuleByLabelsMatchResponse;
}

export namespace AddInboundAllowRuleByLabelsMatchResponse {
  export type AsObject = {
    ruleId: string,
    instancesList: Array<types_pb.Instance.AsObject>,
  }
}

export class AddInboundAllowRuleBySubnetMatchRequest extends jspb.Message {
  getProvider(): string;
  setProvider(value: string): AddInboundAllowRuleBySubnetMatchRequest;

  getVpcId(): string;
  setVpcId(value: string): AddInboundAllowRuleBySubnetMatchRequest;

  getRuleName(): string;
  setRuleName(value: string): AddInboundAllowRuleBySubnetMatchRequest;

  getSubnetCidrsList(): Array<string>;
  setSubnetCidrsList(value: Array<string>): AddInboundAllowRuleBySubnetMatchRequest;
  clearSubnetCidrsList(): AddInboundAllowRuleBySubnetMatchRequest;
  addSubnetCidrs(value: string, index?: number): AddInboundAllowRuleBySubnetMatchRequest;

  getCidrsToAllowList(): Array<string>;
  setCidrsToAllowList(value: Array<string>): AddInboundAllowRuleBySubnetMatchRequest;
  clearCidrsToAllowList(): AddInboundAllowRuleBySubnetMatchRequest;
  addCidrsToAllow(value: string, index?: number): AddInboundAllowRuleBySubnetMatchRequest;

  getProtocolsAndPortsMap(): jspb.Map<string, Ports>;
  clearProtocolsAndPortsMap(): AddInboundAllowRuleBySubnetMatchRequest;

  getRegion(): string;
  setRegion(value: string): AddInboundAllowRuleBySubnetMatchRequest;

  getAccountId(): string;
  setAccountId(value: string): AddInboundAllowRuleBySubnetMatchRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AddInboundAllowRuleBySubnetMatchRequest.AsObject;
  static toObject(includeInstance: boolean, msg: AddInboundAllowRuleBySubnetMatchRequest): AddInboundAllowRuleBySubnetMatchRequest.AsObject;
  static serializeBinaryToWriter(message: AddInboundAllowRuleBySubnetMatchRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AddInboundAllowRuleBySubnetMatchRequest;
  static deserializeBinaryFromReader(message: AddInboundAllowRuleBySubnetMatchRequest, reader: jspb.BinaryReader): AddInboundAllowRuleBySubnetMatchRequest;
}

export namespace AddInboundAllowRuleBySubnetMatchRequest {
  export type AsObject = {
    provider: string,
    vpcId: string,
    ruleName: string,
    subnetCidrsList: Array<string>,
    cidrsToAllowList: Array<string>,
    protocolsAndPortsMap: Array<[string, Ports.AsObject]>,
    region: string,
    accountId: string,
  }
}

export class AddInboundAllowRuleBySubnetMatchResponse extends jspb.Message {
  getRuleId(): string;
  setRuleId(value: string): AddInboundAllowRuleBySubnetMatchResponse;

  getInstancesList(): Array<types_pb.Instance>;
  setInstancesList(value: Array<types_pb.Instance>): AddInboundAllowRuleBySubnetMatchResponse;
  clearInstancesList(): AddInboundAllowRuleBySubnetMatchResponse;
  addInstances(value?: types_pb.Instance, index?: number): types_pb.Instance;

  getSubnetsList(): Array<types_pb.Subnet>;
  setSubnetsList(value: Array<types_pb.Subnet>): AddInboundAllowRuleBySubnetMatchResponse;
  clearSubnetsList(): AddInboundAllowRuleBySubnetMatchResponse;
  addSubnets(value?: types_pb.Subnet, index?: number): types_pb.Subnet;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AddInboundAllowRuleBySubnetMatchResponse.AsObject;
  static toObject(includeInstance: boolean, msg: AddInboundAllowRuleBySubnetMatchResponse): AddInboundAllowRuleBySubnetMatchResponse.AsObject;
  static serializeBinaryToWriter(message: AddInboundAllowRuleBySubnetMatchResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AddInboundAllowRuleBySubnetMatchResponse;
  static deserializeBinaryFromReader(message: AddInboundAllowRuleBySubnetMatchResponse, reader: jspb.BinaryReader): AddInboundAllowRuleBySubnetMatchResponse;
}

export namespace AddInboundAllowRuleBySubnetMatchResponse {
  export type AsObject = {
    ruleId: string,
    instancesList: Array<types_pb.Instance.AsObject>,
    subnetsList: Array<types_pb.Subnet.AsObject>,
  }
}

export class AddInboundAllowRuleByInstanceIPMatchRequest extends jspb.Message {
  getProvider(): string;
  setProvider(value: string): AddInboundAllowRuleByInstanceIPMatchRequest;

  getVpcId(): string;
  setVpcId(value: string): AddInboundAllowRuleByInstanceIPMatchRequest;

  getRuleName(): string;
  setRuleName(value: string): AddInboundAllowRuleByInstanceIPMatchRequest;

  getInstancesIpsList(): Array<string>;
  setInstancesIpsList(value: Array<string>): AddInboundAllowRuleByInstanceIPMatchRequest;
  clearInstancesIpsList(): AddInboundAllowRuleByInstanceIPMatchRequest;
  addInstancesIps(value: string, index?: number): AddInboundAllowRuleByInstanceIPMatchRequest;

  getCidrsToAllowList(): Array<string>;
  setCidrsToAllowList(value: Array<string>): AddInboundAllowRuleByInstanceIPMatchRequest;
  clearCidrsToAllowList(): AddInboundAllowRuleByInstanceIPMatchRequest;
  addCidrsToAllow(value: string, index?: number): AddInboundAllowRuleByInstanceIPMatchRequest;

  getProtocolsAndPortsMap(): jspb.Map<string, Ports>;
  clearProtocolsAndPortsMap(): AddInboundAllowRuleByInstanceIPMatchRequest;

  getRegion(): string;
  setRegion(value: string): AddInboundAllowRuleByInstanceIPMatchRequest;

  getAccountId(): string;
  setAccountId(value: string): AddInboundAllowRuleByInstanceIPMatchRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AddInboundAllowRuleByInstanceIPMatchRequest.AsObject;
  static toObject(includeInstance: boolean, msg: AddInboundAllowRuleByInstanceIPMatchRequest): AddInboundAllowRuleByInstanceIPMatchRequest.AsObject;
  static serializeBinaryToWriter(message: AddInboundAllowRuleByInstanceIPMatchRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AddInboundAllowRuleByInstanceIPMatchRequest;
  static deserializeBinaryFromReader(message: AddInboundAllowRuleByInstanceIPMatchRequest, reader: jspb.BinaryReader): AddInboundAllowRuleByInstanceIPMatchRequest;
}

export namespace AddInboundAllowRuleByInstanceIPMatchRequest {
  export type AsObject = {
    provider: string,
    vpcId: string,
    ruleName: string,
    instancesIpsList: Array<string>,
    cidrsToAllowList: Array<string>,
    protocolsAndPortsMap: Array<[string, Ports.AsObject]>,
    region: string,
    accountId: string,
  }
}

export class AddInboundAllowRuleByInstanceIPMatchResponse extends jspb.Message {
  getRuleId(): string;
  setRuleId(value: string): AddInboundAllowRuleByInstanceIPMatchResponse;

  getInstancesList(): Array<types_pb.Instance>;
  setInstancesList(value: Array<types_pb.Instance>): AddInboundAllowRuleByInstanceIPMatchResponse;
  clearInstancesList(): AddInboundAllowRuleByInstanceIPMatchResponse;
  addInstances(value?: types_pb.Instance, index?: number): types_pb.Instance;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AddInboundAllowRuleByInstanceIPMatchResponse.AsObject;
  static toObject(includeInstance: boolean, msg: AddInboundAllowRuleByInstanceIPMatchResponse): AddInboundAllowRuleByInstanceIPMatchResponse.AsObject;
  static serializeBinaryToWriter(message: AddInboundAllowRuleByInstanceIPMatchResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AddInboundAllowRuleByInstanceIPMatchResponse;
  static deserializeBinaryFromReader(message: AddInboundAllowRuleByInstanceIPMatchResponse, reader: jspb.BinaryReader): AddInboundAllowRuleByInstanceIPMatchResponse;
}

export namespace AddInboundAllowRuleByInstanceIPMatchResponse {
  export type AsObject = {
    ruleId: string,
    instancesList: Array<types_pb.Instance.AsObject>,
  }
}

export class AddInboundAllowRuleForLoadBalancerByDNSRequest extends jspb.Message {
  getProvider(): string;
  setProvider(value: string): AddInboundAllowRuleForLoadBalancerByDNSRequest;

  getLoadBalancerDns(): string;
  setLoadBalancerDns(value: string): AddInboundAllowRuleForLoadBalancerByDNSRequest;

  getVpcId(): string;
  setVpcId(value: string): AddInboundAllowRuleForLoadBalancerByDNSRequest;

  getRuleName(): string;
  setRuleName(value: string): AddInboundAllowRuleForLoadBalancerByDNSRequest;

  getCidrsToAllowList(): Array<string>;
  setCidrsToAllowList(value: Array<string>): AddInboundAllowRuleForLoadBalancerByDNSRequest;
  clearCidrsToAllowList(): AddInboundAllowRuleForLoadBalancerByDNSRequest;
  addCidrsToAllow(value: string, index?: number): AddInboundAllowRuleForLoadBalancerByDNSRequest;

  getProtocolsAndPortsMap(): jspb.Map<string, Ports>;
  clearProtocolsAndPortsMap(): AddInboundAllowRuleForLoadBalancerByDNSRequest;

  getRegion(): string;
  setRegion(value: string): AddInboundAllowRuleForLoadBalancerByDNSRequest;

  getAccountId(): string;
  setAccountId(value: string): AddInboundAllowRuleForLoadBalancerByDNSRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AddInboundAllowRuleForLoadBalancerByDNSRequest.AsObject;
  static toObject(includeInstance: boolean, msg: AddInboundAllowRuleForLoadBalancerByDNSRequest): AddInboundAllowRuleForLoadBalancerByDNSRequest.AsObject;
  static serializeBinaryToWriter(message: AddInboundAllowRuleForLoadBalancerByDNSRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AddInboundAllowRuleForLoadBalancerByDNSRequest;
  static deserializeBinaryFromReader(message: AddInboundAllowRuleForLoadBalancerByDNSRequest, reader: jspb.BinaryReader): AddInboundAllowRuleForLoadBalancerByDNSRequest;
}

export namespace AddInboundAllowRuleForLoadBalancerByDNSRequest {
  export type AsObject = {
    provider: string,
    loadBalancerDns: string,
    vpcId: string,
    ruleName: string,
    cidrsToAllowList: Array<string>,
    protocolsAndPortsMap: Array<[string, Ports.AsObject]>,
    region: string,
    accountId: string,
  }
}

export class AddInboundAllowRuleForLoadBalancerByDNSResponse extends jspb.Message {
  getLoadBalancerId(): string;
  setLoadBalancerId(value: string): AddInboundAllowRuleForLoadBalancerByDNSResponse;

  getRuleId(): string;
  setRuleId(value: string): AddInboundAllowRuleForLoadBalancerByDNSResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AddInboundAllowRuleForLoadBalancerByDNSResponse.AsObject;
  static toObject(includeInstance: boolean, msg: AddInboundAllowRuleForLoadBalancerByDNSResponse): AddInboundAllowRuleForLoadBalancerByDNSResponse.AsObject;
  static serializeBinaryToWriter(message: AddInboundAllowRuleForLoadBalancerByDNSResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AddInboundAllowRuleForLoadBalancerByDNSResponse;
  static deserializeBinaryFromReader(message: AddInboundAllowRuleForLoadBalancerByDNSResponse, reader: jspb.BinaryReader): AddInboundAllowRuleForLoadBalancerByDNSResponse;
}

export namespace AddInboundAllowRuleForLoadBalancerByDNSResponse {
  export type AsObject = {
    loadBalancerId: string,
    ruleId: string,
  }
}

export class RemoveInboundAllowRuleFromVPCByNameRequest extends jspb.Message {
  getProvider(): string;
  setProvider(value: string): RemoveInboundAllowRuleFromVPCByNameRequest;

  getVpcId(): string;
  setVpcId(value: string): RemoveInboundAllowRuleFromVPCByNameRequest;

  getRuleName(): string;
  setRuleName(value: string): RemoveInboundAllowRuleFromVPCByNameRequest;

  getRegion(): string;
  setRegion(value: string): RemoveInboundAllowRuleFromVPCByNameRequest;

  getAccountId(): string;
  setAccountId(value: string): RemoveInboundAllowRuleFromVPCByNameRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RemoveInboundAllowRuleFromVPCByNameRequest.AsObject;
  static toObject(includeInstance: boolean, msg: RemoveInboundAllowRuleFromVPCByNameRequest): RemoveInboundAllowRuleFromVPCByNameRequest.AsObject;
  static serializeBinaryToWriter(message: RemoveInboundAllowRuleFromVPCByNameRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RemoveInboundAllowRuleFromVPCByNameRequest;
  static deserializeBinaryFromReader(message: RemoveInboundAllowRuleFromVPCByNameRequest, reader: jspb.BinaryReader): RemoveInboundAllowRuleFromVPCByNameRequest;
}

export namespace RemoveInboundAllowRuleFromVPCByNameRequest {
  export type AsObject = {
    provider: string,
    vpcId: string,
    ruleName: string,
    region: string,
    accountId: string,
  }
}

export class RemoveInboundAllowRuleFromVPCByNameResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RemoveInboundAllowRuleFromVPCByNameResponse.AsObject;
  static toObject(includeInstance: boolean, msg: RemoveInboundAllowRuleFromVPCByNameResponse): RemoveInboundAllowRuleFromVPCByNameResponse.AsObject;
  static serializeBinaryToWriter(message: RemoveInboundAllowRuleFromVPCByNameResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RemoveInboundAllowRuleFromVPCByNameResponse;
  static deserializeBinaryFromReader(message: RemoveInboundAllowRuleFromVPCByNameResponse, reader: jspb.BinaryReader): RemoveInboundAllowRuleFromVPCByNameResponse;
}

export namespace RemoveInboundAllowRuleFromVPCByNameResponse {
  export type AsObject = {
  }
}

export class RemoveInboundAllowRulesFromVPCByIdRequest extends jspb.Message {
  getProvider(): string;
  setProvider(value: string): RemoveInboundAllowRulesFromVPCByIdRequest;

  getVpcId(): string;
  setVpcId(value: string): RemoveInboundAllowRulesFromVPCByIdRequest;

  getInstanceIdsList(): Array<string>;
  setInstanceIdsList(value: Array<string>): RemoveInboundAllowRulesFromVPCByIdRequest;
  clearInstanceIdsList(): RemoveInboundAllowRulesFromVPCByIdRequest;
  addInstanceIds(value: string, index?: number): RemoveInboundAllowRulesFromVPCByIdRequest;

  getLoadBalancerIdsList(): Array<string>;
  setLoadBalancerIdsList(value: Array<string>): RemoveInboundAllowRulesFromVPCByIdRequest;
  clearLoadBalancerIdsList(): RemoveInboundAllowRulesFromVPCByIdRequest;
  addLoadBalancerIds(value: string, index?: number): RemoveInboundAllowRulesFromVPCByIdRequest;

  getRuleId(): string;
  setRuleId(value: string): RemoveInboundAllowRulesFromVPCByIdRequest;

  getRegion(): string;
  setRegion(value: string): RemoveInboundAllowRulesFromVPCByIdRequest;

  getAccountId(): string;
  setAccountId(value: string): RemoveInboundAllowRulesFromVPCByIdRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RemoveInboundAllowRulesFromVPCByIdRequest.AsObject;
  static toObject(includeInstance: boolean, msg: RemoveInboundAllowRulesFromVPCByIdRequest): RemoveInboundAllowRulesFromVPCByIdRequest.AsObject;
  static serializeBinaryToWriter(message: RemoveInboundAllowRulesFromVPCByIdRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RemoveInboundAllowRulesFromVPCByIdRequest;
  static deserializeBinaryFromReader(message: RemoveInboundAllowRulesFromVPCByIdRequest, reader: jspb.BinaryReader): RemoveInboundAllowRulesFromVPCByIdRequest;
}

export namespace RemoveInboundAllowRulesFromVPCByIdRequest {
  export type AsObject = {
    provider: string,
    vpcId: string,
    instanceIdsList: Array<string>,
    loadBalancerIdsList: Array<string>,
    ruleId: string,
    region: string,
    accountId: string,
  }
}

export class RemoveInboundAllowRulesFromVPCByIdResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RemoveInboundAllowRulesFromVPCByIdResponse.AsObject;
  static toObject(includeInstance: boolean, msg: RemoveInboundAllowRulesFromVPCByIdResponse): RemoveInboundAllowRulesFromVPCByIdResponse.AsObject;
  static serializeBinaryToWriter(message: RemoveInboundAllowRulesFromVPCByIdResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RemoveInboundAllowRulesFromVPCByIdResponse;
  static deserializeBinaryFromReader(message: RemoveInboundAllowRulesFromVPCByIdResponse, reader: jspb.BinaryReader): RemoveInboundAllowRulesFromVPCByIdResponse;
}

export namespace RemoveInboundAllowRulesFromVPCByIdResponse {
  export type AsObject = {
  }
}

export class RemoveInboundAllowRuleRulesByTagsRequest extends jspb.Message {
  getProvider(): string;
  setProvider(value: string): RemoveInboundAllowRuleRulesByTagsRequest;

  getVpcId(): string;
  setVpcId(value: string): RemoveInboundAllowRuleRulesByTagsRequest;

  getRuleName(): string;
  setRuleName(value: string): RemoveInboundAllowRuleRulesByTagsRequest;

  getTagsMap(): jspb.Map<string, string>;
  clearTagsMap(): RemoveInboundAllowRuleRulesByTagsRequest;

  getRegion(): string;
  setRegion(value: string): RemoveInboundAllowRuleRulesByTagsRequest;

  getAccountId(): string;
  setAccountId(value: string): RemoveInboundAllowRuleRulesByTagsRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RemoveInboundAllowRuleRulesByTagsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: RemoveInboundAllowRuleRulesByTagsRequest): RemoveInboundAllowRuleRulesByTagsRequest.AsObject;
  static serializeBinaryToWriter(message: RemoveInboundAllowRuleRulesByTagsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RemoveInboundAllowRuleRulesByTagsRequest;
  static deserializeBinaryFromReader(message: RemoveInboundAllowRuleRulesByTagsRequest, reader: jspb.BinaryReader): RemoveInboundAllowRuleRulesByTagsRequest;
}

export namespace RemoveInboundAllowRuleRulesByTagsRequest {
  export type AsObject = {
    provider: string,
    vpcId: string,
    ruleName: string,
    tagsMap: Array<[string, string]>,
    region: string,
    accountId: string,
  }
}

export class RemoveInboundAllowRuleRulesByTagsResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RemoveInboundAllowRuleRulesByTagsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: RemoveInboundAllowRuleRulesByTagsResponse): RemoveInboundAllowRuleRulesByTagsResponse.AsObject;
  static serializeBinaryToWriter(message: RemoveInboundAllowRuleRulesByTagsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RemoveInboundAllowRuleRulesByTagsResponse;
  static deserializeBinaryFromReader(message: RemoveInboundAllowRuleRulesByTagsResponse, reader: jspb.BinaryReader): RemoveInboundAllowRuleRulesByTagsResponse;
}

export namespace RemoveInboundAllowRuleRulesByTagsResponse {
  export type AsObject = {
  }
}

export class RefreshInboundAllowRuleRequest extends jspb.Message {
  getProvider(): string;
  setProvider(value: string): RefreshInboundAllowRuleRequest;

  getRuleId(): string;
  setRuleId(value: string): RefreshInboundAllowRuleRequest;

  getCidrsToAddList(): Array<string>;
  setCidrsToAddList(value: Array<string>): RefreshInboundAllowRuleRequest;
  clearCidrsToAddList(): RefreshInboundAllowRuleRequest;
  addCidrsToAdd(value: string, index?: number): RefreshInboundAllowRuleRequest;

  getCidrsToRemoveList(): Array<string>;
  setCidrsToRemoveList(value: Array<string>): RefreshInboundAllowRuleRequest;
  clearCidrsToRemoveList(): RefreshInboundAllowRuleRequest;
  addCidrsToRemove(value: string, index?: number): RefreshInboundAllowRuleRequest;

  getDestinationLabelsMap(): jspb.Map<string, string>;
  clearDestinationLabelsMap(): RefreshInboundAllowRuleRequest;

  getDestinationPrefixesList(): Array<string>;
  setDestinationPrefixesList(value: Array<string>): RefreshInboundAllowRuleRequest;
  clearDestinationPrefixesList(): RefreshInboundAllowRuleRequest;
  addDestinationPrefixes(value: string, index?: number): RefreshInboundAllowRuleRequest;

  getDestinationVpcId(): string;
  setDestinationVpcId(value: string): RefreshInboundAllowRuleRequest;

  getProtocolsAndPortsMap(): jspb.Map<string, Ports>;
  clearProtocolsAndPortsMap(): RefreshInboundAllowRuleRequest;

  getRegion(): string;
  setRegion(value: string): RefreshInboundAllowRuleRequest;

  getAccountId(): string;
  setAccountId(value: string): RefreshInboundAllowRuleRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RefreshInboundAllowRuleRequest.AsObject;
  static toObject(includeInstance: boolean, msg: RefreshInboundAllowRuleRequest): RefreshInboundAllowRuleRequest.AsObject;
  static serializeBinaryToWriter(message: RefreshInboundAllowRuleRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RefreshInboundAllowRuleRequest;
  static deserializeBinaryFromReader(message: RefreshInboundAllowRuleRequest, reader: jspb.BinaryReader): RefreshInboundAllowRuleRequest;
}

export namespace RefreshInboundAllowRuleRequest {
  export type AsObject = {
    provider: string,
    ruleId: string,
    cidrsToAddList: Array<string>,
    cidrsToRemoveList: Array<string>,
    destinationLabelsMap: Array<[string, string]>,
    destinationPrefixesList: Array<string>,
    destinationVpcId: string,
    protocolsAndPortsMap: Array<[string, Ports.AsObject]>,
    region: string,
    accountId: string,
  }
}

export class RefreshInboundAllowRuleResponse extends jspb.Message {
  getInstancesList(): Array<types_pb.Instance>;
  setInstancesList(value: Array<types_pb.Instance>): RefreshInboundAllowRuleResponse;
  clearInstancesList(): RefreshInboundAllowRuleResponse;
  addInstances(value?: types_pb.Instance, index?: number): types_pb.Instance;

  getSubnetsList(): Array<types_pb.Subnet>;
  setSubnetsList(value: Array<types_pb.Subnet>): RefreshInboundAllowRuleResponse;
  clearSubnetsList(): RefreshInboundAllowRuleResponse;
  addSubnets(value?: types_pb.Subnet, index?: number): types_pb.Subnet;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RefreshInboundAllowRuleResponse.AsObject;
  static toObject(includeInstance: boolean, msg: RefreshInboundAllowRuleResponse): RefreshInboundAllowRuleResponse.AsObject;
  static serializeBinaryToWriter(message: RefreshInboundAllowRuleResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RefreshInboundAllowRuleResponse;
  static deserializeBinaryFromReader(message: RefreshInboundAllowRuleResponse, reader: jspb.BinaryReader): RefreshInboundAllowRuleResponse;
}

export namespace RefreshInboundAllowRuleResponse {
  export type AsObject = {
    instancesList: Array<types_pb.Instance.AsObject>,
    subnetsList: Array<types_pb.Subnet.AsObject>,
  }
}

