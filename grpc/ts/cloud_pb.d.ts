import * as jspb from 'google-protobuf'

import * as types_pb from './types_pb';


export class ListAccountsRequest extends jspb.Message {
  getProvider(): string;
  setProvider(value: string): ListAccountsRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListAccountsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListAccountsRequest): ListAccountsRequest.AsObject;
  static serializeBinaryToWriter(message: ListAccountsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListAccountsRequest;
  static deserializeBinaryFromReader(message: ListAccountsRequest, reader: jspb.BinaryReader): ListAccountsRequest;
}

export namespace ListAccountsRequest {
  export type AsObject = {
    provider: string,
  }
}

export class ListAccountsResponse extends jspb.Message {
  getAccountsList(): Array<types_pb.Account>;
  setAccountsList(value: Array<types_pb.Account>): ListAccountsResponse;
  clearAccountsList(): ListAccountsResponse;
  addAccounts(value?: types_pb.Account, index?: number): types_pb.Account;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListAccountsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListAccountsResponse): ListAccountsResponse.AsObject;
  static serializeBinaryToWriter(message: ListAccountsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListAccountsResponse;
  static deserializeBinaryFromReader(message: ListAccountsResponse, reader: jspb.BinaryReader): ListAccountsResponse;
}

export namespace ListAccountsResponse {
  export type AsObject = {
    accountsList: Array<types_pb.Account.AsObject>,
  }
}

export class ListVPCRequest extends jspb.Message {
  getProvider(): string;
  setProvider(value: string): ListVPCRequest;

  getRegion(): string;
  setRegion(value: string): ListVPCRequest;

  getLabelsMap(): jspb.Map<string, string>;
  clearLabelsMap(): ListVPCRequest;

  getAccountId(): string;
  setAccountId(value: string): ListVPCRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListVPCRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListVPCRequest): ListVPCRequest.AsObject;
  static serializeBinaryToWriter(message: ListVPCRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListVPCRequest;
  static deserializeBinaryFromReader(message: ListVPCRequest, reader: jspb.BinaryReader): ListVPCRequest;
}

export namespace ListVPCRequest {
  export type AsObject = {
    provider: string,
    region: string,
    labelsMap: Array<[string, string]>,
    accountId: string,
  }
}

export class ListVPCResponse extends jspb.Message {
  getVpcsList(): Array<types_pb.VPC>;
  setVpcsList(value: Array<types_pb.VPC>): ListVPCResponse;
  clearVpcsList(): ListVPCResponse;
  addVpcs(value?: types_pb.VPC, index?: number): types_pb.VPC;

  getLastSyncTime(): string;
  setLastSyncTime(value: string): ListVPCResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListVPCResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListVPCResponse): ListVPCResponse.AsObject;
  static serializeBinaryToWriter(message: ListVPCResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListVPCResponse;
  static deserializeBinaryFromReader(message: ListVPCResponse, reader: jspb.BinaryReader): ListVPCResponse;
}

export namespace ListVPCResponse {
  export type AsObject = {
    vpcsList: Array<types_pb.VPC.AsObject>,
    lastSyncTime: string,
  }
}

export class ListInstancesRequest extends jspb.Message {
  getProvider(): string;
  setProvider(value: string): ListInstancesRequest;

  getVpcId(): string;
  setVpcId(value: string): ListInstancesRequest;

  getZone(): string;
  setZone(value: string): ListInstancesRequest;

  getLabelsMap(): jspb.Map<string, string>;
  clearLabelsMap(): ListInstancesRequest;

  getRegion(): string;
  setRegion(value: string): ListInstancesRequest;

  getAccountId(): string;
  setAccountId(value: string): ListInstancesRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListInstancesRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListInstancesRequest): ListInstancesRequest.AsObject;
  static serializeBinaryToWriter(message: ListInstancesRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListInstancesRequest;
  static deserializeBinaryFromReader(message: ListInstancesRequest, reader: jspb.BinaryReader): ListInstancesRequest;
}

export namespace ListInstancesRequest {
  export type AsObject = {
    provider: string,
    vpcId: string,
    zone: string,
    labelsMap: Array<[string, string]>,
    region: string,
    accountId: string,
  }
}

export class ListInstancesResponse extends jspb.Message {
  getInstancesList(): Array<types_pb.Instance>;
  setInstancesList(value: Array<types_pb.Instance>): ListInstancesResponse;
  clearInstancesList(): ListInstancesResponse;
  addInstances(value?: types_pb.Instance, index?: number): types_pb.Instance;

  getLastSyncTime(): string;
  setLastSyncTime(value: string): ListInstancesResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListInstancesResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListInstancesResponse): ListInstancesResponse.AsObject;
  static serializeBinaryToWriter(message: ListInstancesResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListInstancesResponse;
  static deserializeBinaryFromReader(message: ListInstancesResponse, reader: jspb.BinaryReader): ListInstancesResponse;
}

export namespace ListInstancesResponse {
  export type AsObject = {
    instancesList: Array<types_pb.Instance.AsObject>,
    lastSyncTime: string,
  }
}

export class ListACLsRequest extends jspb.Message {
  getProvider(): string;
  setProvider(value: string): ListACLsRequest;

  getVpcId(): string;
  setVpcId(value: string): ListACLsRequest;

  getRegion(): string;
  setRegion(value: string): ListACLsRequest;

  getAccountId(): string;
  setAccountId(value: string): ListACLsRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListACLsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListACLsRequest): ListACLsRequest.AsObject;
  static serializeBinaryToWriter(message: ListACLsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListACLsRequest;
  static deserializeBinaryFromReader(message: ListACLsRequest, reader: jspb.BinaryReader): ListACLsRequest;
}

export namespace ListACLsRequest {
  export type AsObject = {
    provider: string,
    vpcId: string,
    region: string,
    accountId: string,
  }
}

export class ListACLsResponse extends jspb.Message {
  getAclsList(): Array<types_pb.ACL>;
  setAclsList(value: Array<types_pb.ACL>): ListACLsResponse;
  clearAclsList(): ListACLsResponse;
  addAcls(value?: types_pb.ACL, index?: number): types_pb.ACL;

  getLastSyncTime(): string;
  setLastSyncTime(value: string): ListACLsResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListACLsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListACLsResponse): ListACLsResponse.AsObject;
  static serializeBinaryToWriter(message: ListACLsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListACLsResponse;
  static deserializeBinaryFromReader(message: ListACLsResponse, reader: jspb.BinaryReader): ListACLsResponse;
}

export namespace ListACLsResponse {
  export type AsObject = {
    aclsList: Array<types_pb.ACL.AsObject>,
    lastSyncTime: string,
  }
}

export class ListSecurityGroupsRequest extends jspb.Message {
  getProvider(): string;
  setProvider(value: string): ListSecurityGroupsRequest;

  getVpcId(): string;
  setVpcId(value: string): ListSecurityGroupsRequest;

  getRegion(): string;
  setRegion(value: string): ListSecurityGroupsRequest;

  getAccountId(): string;
  setAccountId(value: string): ListSecurityGroupsRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListSecurityGroupsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListSecurityGroupsRequest): ListSecurityGroupsRequest.AsObject;
  static serializeBinaryToWriter(message: ListSecurityGroupsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListSecurityGroupsRequest;
  static deserializeBinaryFromReader(message: ListSecurityGroupsRequest, reader: jspb.BinaryReader): ListSecurityGroupsRequest;
}

export namespace ListSecurityGroupsRequest {
  export type AsObject = {
    provider: string,
    vpcId: string,
    region: string,
    accountId: string,
  }
}

export class ListSecurityGroupsResponse extends jspb.Message {
  getSecurityGroupsList(): Array<types_pb.SecurityGroup>;
  setSecurityGroupsList(value: Array<types_pb.SecurityGroup>): ListSecurityGroupsResponse;
  clearSecurityGroupsList(): ListSecurityGroupsResponse;
  addSecurityGroups(value?: types_pb.SecurityGroup, index?: number): types_pb.SecurityGroup;

  getLastSyncTime(): string;
  setLastSyncTime(value: string): ListSecurityGroupsResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListSecurityGroupsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListSecurityGroupsResponse): ListSecurityGroupsResponse.AsObject;
  static serializeBinaryToWriter(message: ListSecurityGroupsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListSecurityGroupsResponse;
  static deserializeBinaryFromReader(message: ListSecurityGroupsResponse, reader: jspb.BinaryReader): ListSecurityGroupsResponse;
}

export namespace ListSecurityGroupsResponse {
  export type AsObject = {
    securityGroupsList: Array<types_pb.SecurityGroup.AsObject>,
    lastSyncTime: string,
  }
}

export class ListRouteTablesRequest extends jspb.Message {
  getProvider(): string;
  setProvider(value: string): ListRouteTablesRequest;

  getVpcId(): string;
  setVpcId(value: string): ListRouteTablesRequest;

  getRegion(): string;
  setRegion(value: string): ListRouteTablesRequest;

  getAccountId(): string;
  setAccountId(value: string): ListRouteTablesRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListRouteTablesRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListRouteTablesRequest): ListRouteTablesRequest.AsObject;
  static serializeBinaryToWriter(message: ListRouteTablesRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListRouteTablesRequest;
  static deserializeBinaryFromReader(message: ListRouteTablesRequest, reader: jspb.BinaryReader): ListRouteTablesRequest;
}

export namespace ListRouteTablesRequest {
  export type AsObject = {
    provider: string,
    vpcId: string,
    region: string,
    accountId: string,
  }
}

export class ListRouteTablesResponse extends jspb.Message {
  getRouteTablesList(): Array<types_pb.RouteTable>;
  setRouteTablesList(value: Array<types_pb.RouteTable>): ListRouteTablesResponse;
  clearRouteTablesList(): ListRouteTablesResponse;
  addRouteTables(value?: types_pb.RouteTable, index?: number): types_pb.RouteTable;

  getLastSyncTime(): string;
  setLastSyncTime(value: string): ListRouteTablesResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListRouteTablesResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListRouteTablesResponse): ListRouteTablesResponse.AsObject;
  static serializeBinaryToWriter(message: ListRouteTablesResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListRouteTablesResponse;
  static deserializeBinaryFromReader(message: ListRouteTablesResponse, reader: jspb.BinaryReader): ListRouteTablesResponse;
}

export namespace ListRouteTablesResponse {
  export type AsObject = {
    routeTablesList: Array<types_pb.RouteTable.AsObject>,
    lastSyncTime: string,
  }
}

export class GetSubnetRequest extends jspb.Message {
  getProvider(): string;
  setProvider(value: string): GetSubnetRequest;

  getVpcId(): string;
  setVpcId(value: string): GetSubnetRequest;

  getId(): string;
  setId(value: string): GetSubnetRequest;

  getRegion(): string;
  setRegion(value: string): GetSubnetRequest;

  getAccountId(): string;
  setAccountId(value: string): GetSubnetRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetSubnetRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetSubnetRequest): GetSubnetRequest.AsObject;
  static serializeBinaryToWriter(message: GetSubnetRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetSubnetRequest;
  static deserializeBinaryFromReader(message: GetSubnetRequest, reader: jspb.BinaryReader): GetSubnetRequest;
}

export namespace GetSubnetRequest {
  export type AsObject = {
    provider: string,
    vpcId: string,
    id: string,
    region: string,
    accountId: string,
  }
}

export class GetSubnetResponse extends jspb.Message {
  getSubnet(): types_pb.Subnet | undefined;
  setSubnet(value?: types_pb.Subnet): GetSubnetResponse;
  hasSubnet(): boolean;
  clearSubnet(): GetSubnetResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetSubnetResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetSubnetResponse): GetSubnetResponse.AsObject;
  static serializeBinaryToWriter(message: GetSubnetResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetSubnetResponse;
  static deserializeBinaryFromReader(message: GetSubnetResponse, reader: jspb.BinaryReader): GetSubnetResponse;
}

export namespace GetSubnetResponse {
  export type AsObject = {
    subnet?: types_pb.Subnet.AsObject,
  }
}

export class ListSubnetsRequest extends jspb.Message {
  getProvider(): string;
  setProvider(value: string): ListSubnetsRequest;

  getVpcId(): string;
  setVpcId(value: string): ListSubnetsRequest;

  getZone(): string;
  setZone(value: string): ListSubnetsRequest;

  getCidr(): string;
  setCidr(value: string): ListSubnetsRequest;

  getLabelsMap(): jspb.Map<string, string>;
  clearLabelsMap(): ListSubnetsRequest;

  getRegion(): string;
  setRegion(value: string): ListSubnetsRequest;

  getAccountId(): string;
  setAccountId(value: string): ListSubnetsRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListSubnetsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListSubnetsRequest): ListSubnetsRequest.AsObject;
  static serializeBinaryToWriter(message: ListSubnetsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListSubnetsRequest;
  static deserializeBinaryFromReader(message: ListSubnetsRequest, reader: jspb.BinaryReader): ListSubnetsRequest;
}

export namespace ListSubnetsRequest {
  export type AsObject = {
    provider: string,
    vpcId: string,
    zone: string,
    cidr: string,
    labelsMap: Array<[string, string]>,
    region: string,
    accountId: string,
  }
}

export class ListSubnetsResponse extends jspb.Message {
  getSubnetsList(): Array<types_pb.Subnet>;
  setSubnetsList(value: Array<types_pb.Subnet>): ListSubnetsResponse;
  clearSubnetsList(): ListSubnetsResponse;
  addSubnets(value?: types_pb.Subnet, index?: number): types_pb.Subnet;

  getLastSyncTime(): string;
  setLastSyncTime(value: string): ListSubnetsResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListSubnetsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListSubnetsResponse): ListSubnetsResponse.AsObject;
  static serializeBinaryToWriter(message: ListSubnetsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListSubnetsResponse;
  static deserializeBinaryFromReader(message: ListSubnetsResponse, reader: jspb.BinaryReader): ListSubnetsResponse;
}

export namespace ListSubnetsResponse {
  export type AsObject = {
    subnetsList: Array<types_pb.Subnet.AsObject>,
    lastSyncTime: string,
  }
}

export class GetVPCIDForCIDRRequest extends jspb.Message {
  getProvider(): string;
  setProvider(value: string): GetVPCIDForCIDRRequest;

  getCidr(): string;
  setCidr(value: string): GetVPCIDForCIDRRequest;

  getRegion(): string;
  setRegion(value: string): GetVPCIDForCIDRRequest;

  getAccountId(): string;
  setAccountId(value: string): GetVPCIDForCIDRRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetVPCIDForCIDRRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetVPCIDForCIDRRequest): GetVPCIDForCIDRRequest.AsObject;
  static serializeBinaryToWriter(message: GetVPCIDForCIDRRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetVPCIDForCIDRRequest;
  static deserializeBinaryFromReader(message: GetVPCIDForCIDRRequest, reader: jspb.BinaryReader): GetVPCIDForCIDRRequest;
}

export namespace GetVPCIDForCIDRRequest {
  export type AsObject = {
    provider: string,
    cidr: string,
    region: string,
    accountId: string,
  }
}

export class GetVPCIDForCIDRResponse extends jspb.Message {
  getVpcId(): string;
  setVpcId(value: string): GetVPCIDForCIDRResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetVPCIDForCIDRResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetVPCIDForCIDRResponse): GetVPCIDForCIDRResponse.AsObject;
  static serializeBinaryToWriter(message: GetVPCIDForCIDRResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetVPCIDForCIDRResponse;
  static deserializeBinaryFromReader(message: GetVPCIDForCIDRResponse, reader: jspb.BinaryReader): GetVPCIDForCIDRResponse;
}

export namespace GetVPCIDForCIDRResponse {
  export type AsObject = {
    vpcId: string,
  }
}

export class GetCIDRsForLabelsRequest extends jspb.Message {
  getProvider(): string;
  setProvider(value: string): GetCIDRsForLabelsRequest;

  getLabelsMap(): jspb.Map<string, string>;
  clearLabelsMap(): GetCIDRsForLabelsRequest;

  getRegion(): string;
  setRegion(value: string): GetCIDRsForLabelsRequest;

  getAccountId(): string;
  setAccountId(value: string): GetCIDRsForLabelsRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetCIDRsForLabelsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetCIDRsForLabelsRequest): GetCIDRsForLabelsRequest.AsObject;
  static serializeBinaryToWriter(message: GetCIDRsForLabelsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetCIDRsForLabelsRequest;
  static deserializeBinaryFromReader(message: GetCIDRsForLabelsRequest, reader: jspb.BinaryReader): GetCIDRsForLabelsRequest;
}

export namespace GetCIDRsForLabelsRequest {
  export type AsObject = {
    provider: string,
    labelsMap: Array<[string, string]>,
    region: string,
    accountId: string,
  }
}

export class GetCIDRsForLabelsResponse extends jspb.Message {
  getCidrsList(): Array<string>;
  setCidrsList(value: Array<string>): GetCIDRsForLabelsResponse;
  clearCidrsList(): GetCIDRsForLabelsResponse;
  addCidrs(value: string, index?: number): GetCIDRsForLabelsResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetCIDRsForLabelsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetCIDRsForLabelsResponse): GetCIDRsForLabelsResponse.AsObject;
  static serializeBinaryToWriter(message: GetCIDRsForLabelsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetCIDRsForLabelsResponse;
  static deserializeBinaryFromReader(message: GetCIDRsForLabelsResponse, reader: jspb.BinaryReader): GetCIDRsForLabelsResponse;
}

export namespace GetCIDRsForLabelsResponse {
  export type AsObject = {
    cidrsList: Array<string>,
  }
}

export class GetIPsForLabelsRequest extends jspb.Message {
  getProvider(): string;
  setProvider(value: string): GetIPsForLabelsRequest;

  getLabelsMap(): jspb.Map<string, string>;
  clearLabelsMap(): GetIPsForLabelsRequest;

  getRegion(): string;
  setRegion(value: string): GetIPsForLabelsRequest;

  getAccountId(): string;
  setAccountId(value: string): GetIPsForLabelsRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetIPsForLabelsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetIPsForLabelsRequest): GetIPsForLabelsRequest.AsObject;
  static serializeBinaryToWriter(message: GetIPsForLabelsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetIPsForLabelsRequest;
  static deserializeBinaryFromReader(message: GetIPsForLabelsRequest, reader: jspb.BinaryReader): GetIPsForLabelsRequest;
}

export namespace GetIPsForLabelsRequest {
  export type AsObject = {
    provider: string,
    labelsMap: Array<[string, string]>,
    region: string,
    accountId: string,
  }
}

export class GetIPsForLabelsResponse extends jspb.Message {
  getIpsList(): Array<string>;
  setIpsList(value: Array<string>): GetIPsForLabelsResponse;
  clearIpsList(): GetIPsForLabelsResponse;
  addIps(value: string, index?: number): GetIPsForLabelsResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetIPsForLabelsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetIPsForLabelsResponse): GetIPsForLabelsResponse.AsObject;
  static serializeBinaryToWriter(message: GetIPsForLabelsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetIPsForLabelsResponse;
  static deserializeBinaryFromReader(message: GetIPsForLabelsResponse, reader: jspb.BinaryReader): GetIPsForLabelsResponse;
}

export namespace GetIPsForLabelsResponse {
  export type AsObject = {
    ipsList: Array<string>,
  }
}

export class GetInstancesForLabelsRequest extends jspb.Message {
  getProvider(): string;
  setProvider(value: string): GetInstancesForLabelsRequest;

  getVpcId(): string;
  setVpcId(value: string): GetInstancesForLabelsRequest;

  getLabelsMap(): jspb.Map<string, string>;
  clearLabelsMap(): GetInstancesForLabelsRequest;

  getRegion(): string;
  setRegion(value: string): GetInstancesForLabelsRequest;

  getAccountId(): string;
  setAccountId(value: string): GetInstancesForLabelsRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetInstancesForLabelsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetInstancesForLabelsRequest): GetInstancesForLabelsRequest.AsObject;
  static serializeBinaryToWriter(message: GetInstancesForLabelsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetInstancesForLabelsRequest;
  static deserializeBinaryFromReader(message: GetInstancesForLabelsRequest, reader: jspb.BinaryReader): GetInstancesForLabelsRequest;
}

export namespace GetInstancesForLabelsRequest {
  export type AsObject = {
    provider: string,
    vpcId: string,
    labelsMap: Array<[string, string]>,
    region: string,
    accountId: string,
  }
}

export class GetInstancesForLabelsResponse extends jspb.Message {
  getInstancesList(): Array<types_pb.Instance>;
  setInstancesList(value: Array<types_pb.Instance>): GetInstancesForLabelsResponse;
  clearInstancesList(): GetInstancesForLabelsResponse;
  addInstances(value?: types_pb.Instance, index?: number): types_pb.Instance;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetInstancesForLabelsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetInstancesForLabelsResponse): GetInstancesForLabelsResponse.AsObject;
  static serializeBinaryToWriter(message: GetInstancesForLabelsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetInstancesForLabelsResponse;
  static deserializeBinaryFromReader(message: GetInstancesForLabelsResponse, reader: jspb.BinaryReader): GetInstancesForLabelsResponse;
}

export namespace GetInstancesForLabelsResponse {
  export type AsObject = {
    instancesList: Array<types_pb.Instance.AsObject>,
  }
}

export class GetVPCIDWithTagRequest extends jspb.Message {
  getProvider(): string;
  setProvider(value: string): GetVPCIDWithTagRequest;

  getKey(): string;
  setKey(value: string): GetVPCIDWithTagRequest;

  getValue(): string;
  setValue(value: string): GetVPCIDWithTagRequest;

  getRegion(): string;
  setRegion(value: string): GetVPCIDWithTagRequest;

  getAccountId(): string;
  setAccountId(value: string): GetVPCIDWithTagRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetVPCIDWithTagRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetVPCIDWithTagRequest): GetVPCIDWithTagRequest.AsObject;
  static serializeBinaryToWriter(message: GetVPCIDWithTagRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetVPCIDWithTagRequest;
  static deserializeBinaryFromReader(message: GetVPCIDWithTagRequest, reader: jspb.BinaryReader): GetVPCIDWithTagRequest;
}

export namespace GetVPCIDWithTagRequest {
  export type AsObject = {
    provider: string,
    key: string,
    value: string,
    region: string,
    accountId: string,
  }
}

export class GetVPCIDWithTagResponse extends jspb.Message {
  getVpcId(): string;
  setVpcId(value: string): GetVPCIDWithTagResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetVPCIDWithTagResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetVPCIDWithTagResponse): GetVPCIDWithTagResponse.AsObject;
  static serializeBinaryToWriter(message: GetVPCIDWithTagResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetVPCIDWithTagResponse;
  static deserializeBinaryFromReader(message: GetVPCIDWithTagResponse, reader: jspb.BinaryReader): GetVPCIDWithTagResponse;
}

export namespace GetVPCIDWithTagResponse {
  export type AsObject = {
    vpcId: string,
  }
}

export class ListCloudClustersRequest extends jspb.Message {
  getProvider(): string;
  setProvider(value: string): ListCloudClustersRequest;

  getRegion(): string;
  setRegion(value: string): ListCloudClustersRequest;

  getVpcId(): string;
  setVpcId(value: string): ListCloudClustersRequest;

  getLabelsMap(): jspb.Map<string, string>;
  clearLabelsMap(): ListCloudClustersRequest;

  getAccountId(): string;
  setAccountId(value: string): ListCloudClustersRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListCloudClustersRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListCloudClustersRequest): ListCloudClustersRequest.AsObject;
  static serializeBinaryToWriter(message: ListCloudClustersRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListCloudClustersRequest;
  static deserializeBinaryFromReader(message: ListCloudClustersRequest, reader: jspb.BinaryReader): ListCloudClustersRequest;
}

export namespace ListCloudClustersRequest {
  export type AsObject = {
    provider: string,
    region: string,
    vpcId: string,
    labelsMap: Array<[string, string]>,
    accountId: string,
  }
}

export class ListCloudClustersResponse extends jspb.Message {
  getClustersList(): Array<types_pb.Cluster>;
  setClustersList(value: Array<types_pb.Cluster>): ListCloudClustersResponse;
  clearClustersList(): ListCloudClustersResponse;
  addClusters(value?: types_pb.Cluster, index?: number): types_pb.Cluster;

  getLastSyncTime(): string;
  setLastSyncTime(value: string): ListCloudClustersResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListCloudClustersResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListCloudClustersResponse): ListCloudClustersResponse.AsObject;
  static serializeBinaryToWriter(message: ListCloudClustersResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListCloudClustersResponse;
  static deserializeBinaryFromReader(message: ListCloudClustersResponse, reader: jspb.BinaryReader): ListCloudClustersResponse;
}

export namespace ListCloudClustersResponse {
  export type AsObject = {
    clustersList: Array<types_pb.Cluster.AsObject>,
    lastSyncTime: string,
  }
}

export class SummaryRequest extends jspb.Message {
  getProvider(): string;
  setProvider(value: string): SummaryRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SummaryRequest.AsObject;
  static toObject(includeInstance: boolean, msg: SummaryRequest): SummaryRequest.AsObject;
  static serializeBinaryToWriter(message: SummaryRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SummaryRequest;
  static deserializeBinaryFromReader(message: SummaryRequest, reader: jspb.BinaryReader): SummaryRequest;
}

export namespace SummaryRequest {
  export type AsObject = {
    provider: string,
  }
}

export class Counters extends jspb.Message {
  getAccounts(): number;
  setAccounts(value: number): Counters;

  getVpc(): number;
  setVpc(value: number): Counters;

  getSubnets(): number;
  setSubnets(value: number): Counters;

  getRouteTables(): number;
  setRouteTables(value: number): Counters;

  getInstances(): number;
  setInstances(value: number): Counters;

  getClusters(): number;
  setClusters(value: number): Counters;

  getPods(): number;
  setPods(value: number): Counters;

  getServices(): number;
  setServices(value: number): Counters;

  getNamespaces(): number;
  setNamespaces(value: number): Counters;

  getAcls(): number;
  setAcls(value: number): Counters;

  getSecurityGroups(): number;
  setSecurityGroups(value: number): Counters;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Counters.AsObject;
  static toObject(includeInstance: boolean, msg: Counters): Counters.AsObject;
  static serializeBinaryToWriter(message: Counters, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Counters;
  static deserializeBinaryFromReader(message: Counters, reader: jspb.BinaryReader): Counters;
}

export namespace Counters {
  export type AsObject = {
    accounts: number,
    vpc: number,
    subnets: number,
    routeTables: number,
    instances: number,
    clusters: number,
    pods: number,
    services: number,
    namespaces: number,
    acls: number,
    securityGroups: number,
  }
}

export class StatusSummary extends jspb.Message {
  getVmStatusMap(): jspb.Map<string, number>;
  clearVmStatusMap(): StatusSummary;

  getPodStatusMap(): jspb.Map<string, number>;
  clearPodStatusMap(): StatusSummary;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): StatusSummary.AsObject;
  static toObject(includeInstance: boolean, msg: StatusSummary): StatusSummary.AsObject;
  static serializeBinaryToWriter(message: StatusSummary, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): StatusSummary;
  static deserializeBinaryFromReader(message: StatusSummary, reader: jspb.BinaryReader): StatusSummary;
}

export namespace StatusSummary {
  export type AsObject = {
    vmStatusMap: Array<[string, number]>,
    podStatusMap: Array<[string, number]>,
  }
}

export class SummaryResponse extends jspb.Message {
  getCount(): Counters | undefined;
  setCount(value?: Counters): SummaryResponse;
  hasCount(): boolean;
  clearCount(): SummaryResponse;

  getStatuses(): StatusSummary | undefined;
  setStatuses(value?: StatusSummary): SummaryResponse;
  hasStatuses(): boolean;
  clearStatuses(): SummaryResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SummaryResponse.AsObject;
  static toObject(includeInstance: boolean, msg: SummaryResponse): SummaryResponse.AsObject;
  static serializeBinaryToWriter(message: SummaryResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SummaryResponse;
  static deserializeBinaryFromReader(message: SummaryResponse, reader: jspb.BinaryReader): SummaryResponse;
}

export namespace SummaryResponse {
  export type AsObject = {
    count?: Counters.AsObject,
    statuses?: StatusSummary.AsObject,
  }
}

