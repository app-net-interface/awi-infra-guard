import * as jspb from 'google-protobuf'

import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb';
import * as google_protobuf_field_mask_pb from 'google-protobuf/google/protobuf/field_mask_pb';
import * as types_pb from './types_pb';
import * as types_k8s_pb from './types_k8s_pb';


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

  getErr(): types_pb.Error | undefined;
  setErr(value?: types_pb.Error): ListAccountsResponse;
  hasErr(): boolean;
  clearErr(): ListAccountsResponse;

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
    err?: types_pb.Error.AsObject,
  }
}

export class ListRegionsRequest extends jspb.Message {
  getProvider(): string;
  setProvider(value: string): ListRegionsRequest;

  getAccountId(): string;
  setAccountId(value: string): ListRegionsRequest;

  getCreds(): types_pb.Credentials | undefined;
  setCreds(value?: types_pb.Credentials): ListRegionsRequest;
  hasCreds(): boolean;
  clearCreds(): ListRegionsRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListRegionsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListRegionsRequest): ListRegionsRequest.AsObject;
  static serializeBinaryToWriter(message: ListRegionsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListRegionsRequest;
  static deserializeBinaryFromReader(message: ListRegionsRequest, reader: jspb.BinaryReader): ListRegionsRequest;
}

export namespace ListRegionsRequest {
  export type AsObject = {
    provider: string,
    accountId: string,
    creds?: types_pb.Credentials.AsObject,
  }

  export enum CredsCase { 
    _CREDS_NOT_SET = 0,
    CREDS = 3,
  }
}

export class ListRegionsResponse extends jspb.Message {
  getRegionsList(): Array<types_pb.Region>;
  setRegionsList(value: Array<types_pb.Region>): ListRegionsResponse;
  clearRegionsList(): ListRegionsResponse;
  addRegions(value?: types_pb.Region, index?: number): types_pb.Region;

  getErr(): types_pb.Error | undefined;
  setErr(value?: types_pb.Error): ListRegionsResponse;
  hasErr(): boolean;
  clearErr(): ListRegionsResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListRegionsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListRegionsResponse): ListRegionsResponse.AsObject;
  static serializeBinaryToWriter(message: ListRegionsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListRegionsResponse;
  static deserializeBinaryFromReader(message: ListRegionsResponse, reader: jspb.BinaryReader): ListRegionsResponse;
}

export namespace ListRegionsResponse {
  export type AsObject = {
    regionsList: Array<types_pb.Region.AsObject>,
    err?: types_pb.Error.AsObject,
  }
}

export class ListVPCRequest extends jspb.Message {
  getProvider(): string;
  setProvider(value: string): ListVPCRequest;

  getAccountId(): string;
  setAccountId(value: string): ListVPCRequest;

  getRegion(): string;
  setRegion(value: string): ListVPCRequest;

  getLabelsMap(): jspb.Map<string, string>;
  clearLabelsMap(): ListVPCRequest;

  getCreds(): types_pb.Credentials | undefined;
  setCreds(value?: types_pb.Credentials): ListVPCRequest;
  hasCreds(): boolean;
  clearCreds(): ListVPCRequest;

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
    accountId: string,
    region: string,
    labelsMap: Array<[string, string]>,
    creds?: types_pb.Credentials.AsObject,
  }

  export enum CredsCase { 
    _CREDS_NOT_SET = 0,
    CREDS = 5,
  }
}

export class ListVPCResponse extends jspb.Message {
  getVpcsList(): Array<types_pb.VPC>;
  setVpcsList(value: Array<types_pb.VPC>): ListVPCResponse;
  clearVpcsList(): ListVPCResponse;
  addVpcs(value?: types_pb.VPC, index?: number): types_pb.VPC;

  getLastSyncTime(): string;
  setLastSyncTime(value: string): ListVPCResponse;

  getErr(): types_pb.Error | undefined;
  setErr(value?: types_pb.Error): ListVPCResponse;
  hasErr(): boolean;
  clearErr(): ListVPCResponse;

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
    err?: types_pb.Error.AsObject,
  }
}

export class ListInstancesRequest extends jspb.Message {
  getProvider(): string;
  setProvider(value: string): ListInstancesRequest;

  getAccountId(): string;
  setAccountId(value: string): ListInstancesRequest;

  getRegion(): string;
  setRegion(value: string): ListInstancesRequest;

  getVpcId(): string;
  setVpcId(value: string): ListInstancesRequest;

  getZone(): string;
  setZone(value: string): ListInstancesRequest;

  getLabelsMap(): jspb.Map<string, string>;
  clearLabelsMap(): ListInstancesRequest;

  getCreds(): types_pb.Credentials | undefined;
  setCreds(value?: types_pb.Credentials): ListInstancesRequest;
  hasCreds(): boolean;
  clearCreds(): ListInstancesRequest;

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
    accountId: string,
    region: string,
    vpcId: string,
    zone: string,
    labelsMap: Array<[string, string]>,
    creds?: types_pb.Credentials.AsObject,
  }

  export enum CredsCase { 
    _CREDS_NOT_SET = 0,
    CREDS = 7,
  }
}

export class ListInstancesResponse extends jspb.Message {
  getInstancesList(): Array<types_pb.Instance>;
  setInstancesList(value: Array<types_pb.Instance>): ListInstancesResponse;
  clearInstancesList(): ListInstancesResponse;
  addInstances(value?: types_pb.Instance, index?: number): types_pb.Instance;

  getLastSyncTime(): string;
  setLastSyncTime(value: string): ListInstancesResponse;

  getErr(): types_pb.Error | undefined;
  setErr(value?: types_pb.Error): ListInstancesResponse;
  hasErr(): boolean;
  clearErr(): ListInstancesResponse;

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
    err?: types_pb.Error.AsObject,
  }
}

export class ListACLsRequest extends jspb.Message {
  getProvider(): string;
  setProvider(value: string): ListACLsRequest;

  getAccountId(): string;
  setAccountId(value: string): ListACLsRequest;

  getRegion(): string;
  setRegion(value: string): ListACLsRequest;

  getVpcId(): string;
  setVpcId(value: string): ListACLsRequest;

  getLabelsMap(): jspb.Map<string, string>;
  clearLabelsMap(): ListACLsRequest;

  getCreds(): types_pb.Credentials | undefined;
  setCreds(value?: types_pb.Credentials): ListACLsRequest;
  hasCreds(): boolean;
  clearCreds(): ListACLsRequest;

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
    accountId: string,
    region: string,
    vpcId: string,
    labelsMap: Array<[string, string]>,
    creds?: types_pb.Credentials.AsObject,
  }

  export enum CredsCase { 
    _CREDS_NOT_SET = 0,
    CREDS = 6,
  }
}

export class ListACLsResponse extends jspb.Message {
  getAclsList(): Array<types_pb.ACL>;
  setAclsList(value: Array<types_pb.ACL>): ListACLsResponse;
  clearAclsList(): ListACLsResponse;
  addAcls(value?: types_pb.ACL, index?: number): types_pb.ACL;

  getLastSyncTime(): string;
  setLastSyncTime(value: string): ListACLsResponse;

  getErr(): types_pb.Error | undefined;
  setErr(value?: types_pb.Error): ListACLsResponse;
  hasErr(): boolean;
  clearErr(): ListACLsResponse;

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
    err?: types_pb.Error.AsObject,
  }
}

export class ListSecurityGroupsRequest extends jspb.Message {
  getProvider(): string;
  setProvider(value: string): ListSecurityGroupsRequest;

  getAccountId(): string;
  setAccountId(value: string): ListSecurityGroupsRequest;

  getRegion(): string;
  setRegion(value: string): ListSecurityGroupsRequest;

  getVpcId(): string;
  setVpcId(value: string): ListSecurityGroupsRequest;

  getLabelsMap(): jspb.Map<string, string>;
  clearLabelsMap(): ListSecurityGroupsRequest;

  getCreds(): types_pb.Credentials | undefined;
  setCreds(value?: types_pb.Credentials): ListSecurityGroupsRequest;
  hasCreds(): boolean;
  clearCreds(): ListSecurityGroupsRequest;

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
    accountId: string,
    region: string,
    vpcId: string,
    labelsMap: Array<[string, string]>,
    creds?: types_pb.Credentials.AsObject,
  }

  export enum CredsCase { 
    _CREDS_NOT_SET = 0,
    CREDS = 6,
  }
}

export class ListSecurityGroupsResponse extends jspb.Message {
  getSecurityGroupsList(): Array<types_pb.SecurityGroup>;
  setSecurityGroupsList(value: Array<types_pb.SecurityGroup>): ListSecurityGroupsResponse;
  clearSecurityGroupsList(): ListSecurityGroupsResponse;
  addSecurityGroups(value?: types_pb.SecurityGroup, index?: number): types_pb.SecurityGroup;

  getLastSyncTime(): string;
  setLastSyncTime(value: string): ListSecurityGroupsResponse;

  getErr(): types_pb.Error | undefined;
  setErr(value?: types_pb.Error): ListSecurityGroupsResponse;
  hasErr(): boolean;
  clearErr(): ListSecurityGroupsResponse;

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
    err?: types_pb.Error.AsObject,
  }
}

export class ListRouteTablesRequest extends jspb.Message {
  getProvider(): string;
  setProvider(value: string): ListRouteTablesRequest;

  getAccountId(): string;
  setAccountId(value: string): ListRouteTablesRequest;

  getRegion(): string;
  setRegion(value: string): ListRouteTablesRequest;

  getVpcId(): string;
  setVpcId(value: string): ListRouteTablesRequest;

  getLabelsMap(): jspb.Map<string, string>;
  clearLabelsMap(): ListRouteTablesRequest;

  getCreds(): types_pb.Credentials | undefined;
  setCreds(value?: types_pb.Credentials): ListRouteTablesRequest;
  hasCreds(): boolean;
  clearCreds(): ListRouteTablesRequest;

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
    accountId: string,
    region: string,
    vpcId: string,
    labelsMap: Array<[string, string]>,
    creds?: types_pb.Credentials.AsObject,
  }

  export enum CredsCase { 
    _CREDS_NOT_SET = 0,
    CREDS = 6,
  }
}

export class ListRouteTablesResponse extends jspb.Message {
  getRouteTablesList(): Array<types_pb.RouteTable>;
  setRouteTablesList(value: Array<types_pb.RouteTable>): ListRouteTablesResponse;
  clearRouteTablesList(): ListRouteTablesResponse;
  addRouteTables(value?: types_pb.RouteTable, index?: number): types_pb.RouteTable;

  getLastSyncTime(): string;
  setLastSyncTime(value: string): ListRouteTablesResponse;

  getErr(): types_pb.Error | undefined;
  setErr(value?: types_pb.Error): ListRouteTablesResponse;
  hasErr(): boolean;
  clearErr(): ListRouteTablesResponse;

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
    err?: types_pb.Error.AsObject,
  }
}

export class ListNATGatewaysRequest extends jspb.Message {
  getProvider(): string;
  setProvider(value: string): ListNATGatewaysRequest;

  getAccountId(): string;
  setAccountId(value: string): ListNATGatewaysRequest;

  getRegion(): string;
  setRegion(value: string): ListNATGatewaysRequest;

  getVpcId(): string;
  setVpcId(value: string): ListNATGatewaysRequest;

  getLabelsMap(): jspb.Map<string, string>;
  clearLabelsMap(): ListNATGatewaysRequest;

  getCreds(): types_pb.Credentials | undefined;
  setCreds(value?: types_pb.Credentials): ListNATGatewaysRequest;
  hasCreds(): boolean;
  clearCreds(): ListNATGatewaysRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListNATGatewaysRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListNATGatewaysRequest): ListNATGatewaysRequest.AsObject;
  static serializeBinaryToWriter(message: ListNATGatewaysRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListNATGatewaysRequest;
  static deserializeBinaryFromReader(message: ListNATGatewaysRequest, reader: jspb.BinaryReader): ListNATGatewaysRequest;
}

export namespace ListNATGatewaysRequest {
  export type AsObject = {
    provider: string,
    accountId: string,
    region: string,
    vpcId: string,
    labelsMap: Array<[string, string]>,
    creds?: types_pb.Credentials.AsObject,
  }

  export enum CredsCase { 
    _CREDS_NOT_SET = 0,
    CREDS = 6,
  }
}

export class ListNATGatewaysResponse extends jspb.Message {
  getNatGatewaysList(): Array<types_pb.NATGateway>;
  setNatGatewaysList(value: Array<types_pb.NATGateway>): ListNATGatewaysResponse;
  clearNatGatewaysList(): ListNATGatewaysResponse;
  addNatGateways(value?: types_pb.NATGateway, index?: number): types_pb.NATGateway;

  getLastSyncTime(): string;
  setLastSyncTime(value: string): ListNATGatewaysResponse;

  getErr(): types_pb.Error | undefined;
  setErr(value?: types_pb.Error): ListNATGatewaysResponse;
  hasErr(): boolean;
  clearErr(): ListNATGatewaysResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListNATGatewaysResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListNATGatewaysResponse): ListNATGatewaysResponse.AsObject;
  static serializeBinaryToWriter(message: ListNATGatewaysResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListNATGatewaysResponse;
  static deserializeBinaryFromReader(message: ListNATGatewaysResponse, reader: jspb.BinaryReader): ListNATGatewaysResponse;
}

export namespace ListNATGatewaysResponse {
  export type AsObject = {
    natGatewaysList: Array<types_pb.NATGateway.AsObject>,
    lastSyncTime: string,
    err?: types_pb.Error.AsObject,
  }
}

export class ListRoutersRequest extends jspb.Message {
  getProvider(): string;
  setProvider(value: string): ListRoutersRequest;

  getAccountId(): string;
  setAccountId(value: string): ListRoutersRequest;

  getRegion(): string;
  setRegion(value: string): ListRoutersRequest;

  getVpcId(): string;
  setVpcId(value: string): ListRoutersRequest;

  getLabelsMap(): jspb.Map<string, string>;
  clearLabelsMap(): ListRoutersRequest;

  getCreds(): types_pb.Credentials | undefined;
  setCreds(value?: types_pb.Credentials): ListRoutersRequest;
  hasCreds(): boolean;
  clearCreds(): ListRoutersRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListRoutersRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListRoutersRequest): ListRoutersRequest.AsObject;
  static serializeBinaryToWriter(message: ListRoutersRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListRoutersRequest;
  static deserializeBinaryFromReader(message: ListRoutersRequest, reader: jspb.BinaryReader): ListRoutersRequest;
}

export namespace ListRoutersRequest {
  export type AsObject = {
    provider: string,
    accountId: string,
    region: string,
    vpcId: string,
    labelsMap: Array<[string, string]>,
    creds?: types_pb.Credentials.AsObject,
  }

  export enum CredsCase { 
    _CREDS_NOT_SET = 0,
    CREDS = 6,
  }
}

export class ListRoutersResponse extends jspb.Message {
  getRoutersList(): Array<types_pb.Router>;
  setRoutersList(value: Array<types_pb.Router>): ListRoutersResponse;
  clearRoutersList(): ListRoutersResponse;
  addRouters(value?: types_pb.Router, index?: number): types_pb.Router;

  getLastSyncTime(): string;
  setLastSyncTime(value: string): ListRoutersResponse;

  getErr(): types_pb.Error | undefined;
  setErr(value?: types_pb.Error): ListRoutersResponse;
  hasErr(): boolean;
  clearErr(): ListRoutersResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListRoutersResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListRoutersResponse): ListRoutersResponse.AsObject;
  static serializeBinaryToWriter(message: ListRoutersResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListRoutersResponse;
  static deserializeBinaryFromReader(message: ListRoutersResponse, reader: jspb.BinaryReader): ListRoutersResponse;
}

export namespace ListRoutersResponse {
  export type AsObject = {
    routersList: Array<types_pb.Router.AsObject>,
    lastSyncTime: string,
    err?: types_pb.Error.AsObject,
  }
}

export class ListInternetGatewaysRequest extends jspb.Message {
  getProvider(): string;
  setProvider(value: string): ListInternetGatewaysRequest;

  getAccountId(): string;
  setAccountId(value: string): ListInternetGatewaysRequest;

  getRegion(): string;
  setRegion(value: string): ListInternetGatewaysRequest;

  getVpcId(): string;
  setVpcId(value: string): ListInternetGatewaysRequest;

  getLabelsMap(): jspb.Map<string, string>;
  clearLabelsMap(): ListInternetGatewaysRequest;

  getCreds(): types_pb.Credentials | undefined;
  setCreds(value?: types_pb.Credentials): ListInternetGatewaysRequest;
  hasCreds(): boolean;
  clearCreds(): ListInternetGatewaysRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListInternetGatewaysRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListInternetGatewaysRequest): ListInternetGatewaysRequest.AsObject;
  static serializeBinaryToWriter(message: ListInternetGatewaysRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListInternetGatewaysRequest;
  static deserializeBinaryFromReader(message: ListInternetGatewaysRequest, reader: jspb.BinaryReader): ListInternetGatewaysRequest;
}

export namespace ListInternetGatewaysRequest {
  export type AsObject = {
    provider: string,
    accountId: string,
    region: string,
    vpcId: string,
    labelsMap: Array<[string, string]>,
    creds?: types_pb.Credentials.AsObject,
  }

  export enum CredsCase { 
    _CREDS_NOT_SET = 0,
    CREDS = 6,
  }
}

export class ListInternetGatewaysResponse extends jspb.Message {
  getIgwsList(): Array<types_pb.IGW>;
  setIgwsList(value: Array<types_pb.IGW>): ListInternetGatewaysResponse;
  clearIgwsList(): ListInternetGatewaysResponse;
  addIgws(value?: types_pb.IGW, index?: number): types_pb.IGW;

  getLastSyncTime(): string;
  setLastSyncTime(value: string): ListInternetGatewaysResponse;

  getErr(): types_pb.Error | undefined;
  setErr(value?: types_pb.Error): ListInternetGatewaysResponse;
  hasErr(): boolean;
  clearErr(): ListInternetGatewaysResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListInternetGatewaysResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListInternetGatewaysResponse): ListInternetGatewaysResponse.AsObject;
  static serializeBinaryToWriter(message: ListInternetGatewaysResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListInternetGatewaysResponse;
  static deserializeBinaryFromReader(message: ListInternetGatewaysResponse, reader: jspb.BinaryReader): ListInternetGatewaysResponse;
}

export namespace ListInternetGatewaysResponse {
  export type AsObject = {
    igwsList: Array<types_pb.IGW.AsObject>,
    lastSyncTime: string,
    err?: types_pb.Error.AsObject,
  }
}

export class ListVPCEndpointsRequest extends jspb.Message {
  getProvider(): string;
  setProvider(value: string): ListVPCEndpointsRequest;

  getAccountId(): string;
  setAccountId(value: string): ListVPCEndpointsRequest;

  getRegion(): string;
  setRegion(value: string): ListVPCEndpointsRequest;

  getVpcId(): string;
  setVpcId(value: string): ListVPCEndpointsRequest;

  getLabelsMap(): jspb.Map<string, string>;
  clearLabelsMap(): ListVPCEndpointsRequest;

  getCreds(): types_pb.Credentials | undefined;
  setCreds(value?: types_pb.Credentials): ListVPCEndpointsRequest;
  hasCreds(): boolean;
  clearCreds(): ListVPCEndpointsRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListVPCEndpointsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListVPCEndpointsRequest): ListVPCEndpointsRequest.AsObject;
  static serializeBinaryToWriter(message: ListVPCEndpointsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListVPCEndpointsRequest;
  static deserializeBinaryFromReader(message: ListVPCEndpointsRequest, reader: jspb.BinaryReader): ListVPCEndpointsRequest;
}

export namespace ListVPCEndpointsRequest {
  export type AsObject = {
    provider: string,
    accountId: string,
    region: string,
    vpcId: string,
    labelsMap: Array<[string, string]>,
    creds?: types_pb.Credentials.AsObject,
  }

  export enum CredsCase { 
    _CREDS_NOT_SET = 0,
    CREDS = 6,
  }
}

export class ListVPCEndpointsResponse extends jspb.Message {
  getVepsList(): Array<types_pb.VPCEndpoint>;
  setVepsList(value: Array<types_pb.VPCEndpoint>): ListVPCEndpointsResponse;
  clearVepsList(): ListVPCEndpointsResponse;
  addVeps(value?: types_pb.VPCEndpoint, index?: number): types_pb.VPCEndpoint;

  getLastSyncTime(): string;
  setLastSyncTime(value: string): ListVPCEndpointsResponse;

  getErr(): types_pb.Error | undefined;
  setErr(value?: types_pb.Error): ListVPCEndpointsResponse;
  hasErr(): boolean;
  clearErr(): ListVPCEndpointsResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListVPCEndpointsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListVPCEndpointsResponse): ListVPCEndpointsResponse.AsObject;
  static serializeBinaryToWriter(message: ListVPCEndpointsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListVPCEndpointsResponse;
  static deserializeBinaryFromReader(message: ListVPCEndpointsResponse, reader: jspb.BinaryReader): ListVPCEndpointsResponse;
}

export namespace ListVPCEndpointsResponse {
  export type AsObject = {
    vepsList: Array<types_pb.VPCEndpoint.AsObject>,
    lastSyncTime: string,
    err?: types_pb.Error.AsObject,
  }
}

export class ListPublicIPsRequest extends jspb.Message {
  getProvider(): string;
  setProvider(value: string): ListPublicIPsRequest;

  getAccountId(): string;
  setAccountId(value: string): ListPublicIPsRequest;

  getRegion(): string;
  setRegion(value: string): ListPublicIPsRequest;

  getVpcId(): string;
  setVpcId(value: string): ListPublicIPsRequest;

  getLabelsMap(): jspb.Map<string, string>;
  clearLabelsMap(): ListPublicIPsRequest;

  getCreds(): types_pb.Credentials | undefined;
  setCreds(value?: types_pb.Credentials): ListPublicIPsRequest;
  hasCreds(): boolean;
  clearCreds(): ListPublicIPsRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListPublicIPsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListPublicIPsRequest): ListPublicIPsRequest.AsObject;
  static serializeBinaryToWriter(message: ListPublicIPsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListPublicIPsRequest;
  static deserializeBinaryFromReader(message: ListPublicIPsRequest, reader: jspb.BinaryReader): ListPublicIPsRequest;
}

export namespace ListPublicIPsRequest {
  export type AsObject = {
    provider: string,
    accountId: string,
    region: string,
    vpcId: string,
    labelsMap: Array<[string, string]>,
    creds?: types_pb.Credentials.AsObject,
  }

  export enum CredsCase { 
    _CREDS_NOT_SET = 0,
    CREDS = 6,
  }
}

export class ListPublicIPsResponse extends jspb.Message {
  getPublicIpsList(): Array<types_pb.PublicIP>;
  setPublicIpsList(value: Array<types_pb.PublicIP>): ListPublicIPsResponse;
  clearPublicIpsList(): ListPublicIPsResponse;
  addPublicIps(value?: types_pb.PublicIP, index?: number): types_pb.PublicIP;

  getLastSyncTime(): string;
  setLastSyncTime(value: string): ListPublicIPsResponse;

  getErr(): types_pb.Error | undefined;
  setErr(value?: types_pb.Error): ListPublicIPsResponse;
  hasErr(): boolean;
  clearErr(): ListPublicIPsResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListPublicIPsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListPublicIPsResponse): ListPublicIPsResponse.AsObject;
  static serializeBinaryToWriter(message: ListPublicIPsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListPublicIPsResponse;
  static deserializeBinaryFromReader(message: ListPublicIPsResponse, reader: jspb.BinaryReader): ListPublicIPsResponse;
}

export namespace ListPublicIPsResponse {
  export type AsObject = {
    publicIpsList: Array<types_pb.PublicIP.AsObject>,
    lastSyncTime: string,
    err?: types_pb.Error.AsObject,
  }
}

export class ListLBsRequest extends jspb.Message {
  getProvider(): string;
  setProvider(value: string): ListLBsRequest;

  getAccountId(): string;
  setAccountId(value: string): ListLBsRequest;

  getRegion(): string;
  setRegion(value: string): ListLBsRequest;

  getVpcId(): string;
  setVpcId(value: string): ListLBsRequest;

  getLabelsMap(): jspb.Map<string, string>;
  clearLabelsMap(): ListLBsRequest;

  getCreds(): types_pb.Credentials | undefined;
  setCreds(value?: types_pb.Credentials): ListLBsRequest;
  hasCreds(): boolean;
  clearCreds(): ListLBsRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListLBsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListLBsRequest): ListLBsRequest.AsObject;
  static serializeBinaryToWriter(message: ListLBsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListLBsRequest;
  static deserializeBinaryFromReader(message: ListLBsRequest, reader: jspb.BinaryReader): ListLBsRequest;
}

export namespace ListLBsRequest {
  export type AsObject = {
    provider: string,
    accountId: string,
    region: string,
    vpcId: string,
    labelsMap: Array<[string, string]>,
    creds?: types_pb.Credentials.AsObject,
  }

  export enum CredsCase { 
    _CREDS_NOT_SET = 0,
    CREDS = 6,
  }
}

export class ListLBsResponse extends jspb.Message {
  getLbsList(): Array<types_pb.LB>;
  setLbsList(value: Array<types_pb.LB>): ListLBsResponse;
  clearLbsList(): ListLBsResponse;
  addLbs(value?: types_pb.LB, index?: number): types_pb.LB;

  getLastSyncTime(): string;
  setLastSyncTime(value: string): ListLBsResponse;

  getErr(): types_pb.Error | undefined;
  setErr(value?: types_pb.Error): ListLBsResponse;
  hasErr(): boolean;
  clearErr(): ListLBsResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListLBsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListLBsResponse): ListLBsResponse.AsObject;
  static serializeBinaryToWriter(message: ListLBsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListLBsResponse;
  static deserializeBinaryFromReader(message: ListLBsResponse, reader: jspb.BinaryReader): ListLBsResponse;
}

export namespace ListLBsResponse {
  export type AsObject = {
    lbsList: Array<types_pb.LB.AsObject>,
    lastSyncTime: string,
    err?: types_pb.Error.AsObject,
  }
}

export class GetSubnetRequest extends jspb.Message {
  getProvider(): string;
  setProvider(value: string): GetSubnetRequest;

  getAccountId(): string;
  setAccountId(value: string): GetSubnetRequest;

  getRegion(): string;
  setRegion(value: string): GetSubnetRequest;

  getVpcId(): string;
  setVpcId(value: string): GetSubnetRequest;

  getId(): string;
  setId(value: string): GetSubnetRequest;

  getLabelsMap(): jspb.Map<string, string>;
  clearLabelsMap(): GetSubnetRequest;

  getCreds(): types_pb.Credentials | undefined;
  setCreds(value?: types_pb.Credentials): GetSubnetRequest;
  hasCreds(): boolean;
  clearCreds(): GetSubnetRequest;

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
    accountId: string,
    region: string,
    vpcId: string,
    id: string,
    labelsMap: Array<[string, string]>,
    creds?: types_pb.Credentials.AsObject,
  }

  export enum CredsCase { 
    _CREDS_NOT_SET = 0,
    CREDS = 7,
  }
}

export class GetSubnetResponse extends jspb.Message {
  getSubnet(): types_pb.Subnet | undefined;
  setSubnet(value?: types_pb.Subnet): GetSubnetResponse;
  hasSubnet(): boolean;
  clearSubnet(): GetSubnetResponse;

  getErr(): types_pb.Error | undefined;
  setErr(value?: types_pb.Error): GetSubnetResponse;
  hasErr(): boolean;
  clearErr(): GetSubnetResponse;

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
    err?: types_pb.Error.AsObject,
  }
}

export class ListSubnetsRequest extends jspb.Message {
  getProvider(): string;
  setProvider(value: string): ListSubnetsRequest;

  getAccountId(): string;
  setAccountId(value: string): ListSubnetsRequest;

  getRegion(): string;
  setRegion(value: string): ListSubnetsRequest;

  getVpcId(): string;
  setVpcId(value: string): ListSubnetsRequest;

  getZone(): string;
  setZone(value: string): ListSubnetsRequest;

  getCidr(): string;
  setCidr(value: string): ListSubnetsRequest;

  getLabelsMap(): jspb.Map<string, string>;
  clearLabelsMap(): ListSubnetsRequest;

  getCreds(): types_pb.Credentials | undefined;
  setCreds(value?: types_pb.Credentials): ListSubnetsRequest;
  hasCreds(): boolean;
  clearCreds(): ListSubnetsRequest;

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
    accountId: string,
    region: string,
    vpcId: string,
    zone: string,
    cidr: string,
    labelsMap: Array<[string, string]>,
    creds?: types_pb.Credentials.AsObject,
  }

  export enum CredsCase { 
    _CREDS_NOT_SET = 0,
    CREDS = 8,
  }
}

export class ListSubnetsResponse extends jspb.Message {
  getSubnetsList(): Array<types_pb.Subnet>;
  setSubnetsList(value: Array<types_pb.Subnet>): ListSubnetsResponse;
  clearSubnetsList(): ListSubnetsResponse;
  addSubnets(value?: types_pb.Subnet, index?: number): types_pb.Subnet;

  getLastSyncTime(): string;
  setLastSyncTime(value: string): ListSubnetsResponse;

  getErr(): types_pb.Error | undefined;
  setErr(value?: types_pb.Error): ListSubnetsResponse;
  hasErr(): boolean;
  clearErr(): ListSubnetsResponse;

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
    err?: types_pb.Error.AsObject,
  }
}

export class ListNetworkInterfacesRequest extends jspb.Message {
  getProvider(): string;
  setProvider(value: string): ListNetworkInterfacesRequest;

  getAccountId(): string;
  setAccountId(value: string): ListNetworkInterfacesRequest;

  getRegion(): string;
  setRegion(value: string): ListNetworkInterfacesRequest;

  getVpcId(): string;
  setVpcId(value: string): ListNetworkInterfacesRequest;

  getLabelsMap(): jspb.Map<string, string>;
  clearLabelsMap(): ListNetworkInterfacesRequest;

  getCreds(): types_pb.Credentials | undefined;
  setCreds(value?: types_pb.Credentials): ListNetworkInterfacesRequest;
  hasCreds(): boolean;
  clearCreds(): ListNetworkInterfacesRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListNetworkInterfacesRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListNetworkInterfacesRequest): ListNetworkInterfacesRequest.AsObject;
  static serializeBinaryToWriter(message: ListNetworkInterfacesRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListNetworkInterfacesRequest;
  static deserializeBinaryFromReader(message: ListNetworkInterfacesRequest, reader: jspb.BinaryReader): ListNetworkInterfacesRequest;
}

export namespace ListNetworkInterfacesRequest {
  export type AsObject = {
    provider: string,
    accountId: string,
    region: string,
    vpcId: string,
    labelsMap: Array<[string, string]>,
    creds?: types_pb.Credentials.AsObject,
  }

  export enum CredsCase { 
    _CREDS_NOT_SET = 0,
    CREDS = 6,
  }
}

export class ListNetworkInterfacesResponse extends jspb.Message {
  getNetworkInterfacesList(): Array<types_pb.NetworkInterface>;
  setNetworkInterfacesList(value: Array<types_pb.NetworkInterface>): ListNetworkInterfacesResponse;
  clearNetworkInterfacesList(): ListNetworkInterfacesResponse;
  addNetworkInterfaces(value?: types_pb.NetworkInterface, index?: number): types_pb.NetworkInterface;

  getLastSyncTime(): string;
  setLastSyncTime(value: string): ListNetworkInterfacesResponse;

  getErr(): types_pb.Error | undefined;
  setErr(value?: types_pb.Error): ListNetworkInterfacesResponse;
  hasErr(): boolean;
  clearErr(): ListNetworkInterfacesResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListNetworkInterfacesResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListNetworkInterfacesResponse): ListNetworkInterfacesResponse.AsObject;
  static serializeBinaryToWriter(message: ListNetworkInterfacesResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListNetworkInterfacesResponse;
  static deserializeBinaryFromReader(message: ListNetworkInterfacesResponse, reader: jspb.BinaryReader): ListNetworkInterfacesResponse;
}

export namespace ListNetworkInterfacesResponse {
  export type AsObject = {
    networkInterfacesList: Array<types_pb.NetworkInterface.AsObject>,
    lastSyncTime: string,
    err?: types_pb.Error.AsObject,
  }
}

export class ListKeyPairsRequest extends jspb.Message {
  getProvider(): string;
  setProvider(value: string): ListKeyPairsRequest;

  getAccountId(): string;
  setAccountId(value: string): ListKeyPairsRequest;

  getRegion(): string;
  setRegion(value: string): ListKeyPairsRequest;

  getVpcId(): string;
  setVpcId(value: string): ListKeyPairsRequest;

  getLabelsMap(): jspb.Map<string, string>;
  clearLabelsMap(): ListKeyPairsRequest;

  getCreds(): types_pb.Credentials | undefined;
  setCreds(value?: types_pb.Credentials): ListKeyPairsRequest;
  hasCreds(): boolean;
  clearCreds(): ListKeyPairsRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListKeyPairsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListKeyPairsRequest): ListKeyPairsRequest.AsObject;
  static serializeBinaryToWriter(message: ListKeyPairsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListKeyPairsRequest;
  static deserializeBinaryFromReader(message: ListKeyPairsRequest, reader: jspb.BinaryReader): ListKeyPairsRequest;
}

export namespace ListKeyPairsRequest {
  export type AsObject = {
    provider: string,
    accountId: string,
    region: string,
    vpcId: string,
    labelsMap: Array<[string, string]>,
    creds?: types_pb.Credentials.AsObject,
  }

  export enum CredsCase { 
    _CREDS_NOT_SET = 0,
    CREDS = 6,
  }
}

export class ListKeyPairsResponse extends jspb.Message {
  getKeyPairsList(): Array<types_pb.KeyPair>;
  setKeyPairsList(value: Array<types_pb.KeyPair>): ListKeyPairsResponse;
  clearKeyPairsList(): ListKeyPairsResponse;
  addKeyPairs(value?: types_pb.KeyPair, index?: number): types_pb.KeyPair;

  getLastSyncTime(): string;
  setLastSyncTime(value: string): ListKeyPairsResponse;

  getErr(): types_pb.Error | undefined;
  setErr(value?: types_pb.Error): ListKeyPairsResponse;
  hasErr(): boolean;
  clearErr(): ListKeyPairsResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListKeyPairsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListKeyPairsResponse): ListKeyPairsResponse.AsObject;
  static serializeBinaryToWriter(message: ListKeyPairsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListKeyPairsResponse;
  static deserializeBinaryFromReader(message: ListKeyPairsResponse, reader: jspb.BinaryReader): ListKeyPairsResponse;
}

export namespace ListKeyPairsResponse {
  export type AsObject = {
    keyPairsList: Array<types_pb.KeyPair.AsObject>,
    lastSyncTime: string,
    err?: types_pb.Error.AsObject,
  }
}

export class ListVPNConcentratorsRequest extends jspb.Message {
  getProvider(): string;
  setProvider(value: string): ListVPNConcentratorsRequest;

  getAccountId(): string;
  setAccountId(value: string): ListVPNConcentratorsRequest;

  getRegion(): string;
  setRegion(value: string): ListVPNConcentratorsRequest;

  getLabelsMap(): jspb.Map<string, string>;
  clearLabelsMap(): ListVPNConcentratorsRequest;

  getCreds(): types_pb.Credentials | undefined;
  setCreds(value?: types_pb.Credentials): ListVPNConcentratorsRequest;
  hasCreds(): boolean;
  clearCreds(): ListVPNConcentratorsRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListVPNConcentratorsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListVPNConcentratorsRequest): ListVPNConcentratorsRequest.AsObject;
  static serializeBinaryToWriter(message: ListVPNConcentratorsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListVPNConcentratorsRequest;
  static deserializeBinaryFromReader(message: ListVPNConcentratorsRequest, reader: jspb.BinaryReader): ListVPNConcentratorsRequest;
}

export namespace ListVPNConcentratorsRequest {
  export type AsObject = {
    provider: string,
    accountId: string,
    region: string,
    labelsMap: Array<[string, string]>,
    creds?: types_pb.Credentials.AsObject,
  }

  export enum CredsCase { 
    _CREDS_NOT_SET = 0,
    CREDS = 5,
  }
}

export class ListVPNConcentratorsResponse extends jspb.Message {
  getVpnConcentratorsList(): Array<types_pb.VPNConcentrator>;
  setVpnConcentratorsList(value: Array<types_pb.VPNConcentrator>): ListVPNConcentratorsResponse;
  clearVpnConcentratorsList(): ListVPNConcentratorsResponse;
  addVpnConcentrators(value?: types_pb.VPNConcentrator, index?: number): types_pb.VPNConcentrator;

  getLastSyncTime(): string;
  setLastSyncTime(value: string): ListVPNConcentratorsResponse;

  getErr(): types_pb.Error | undefined;
  setErr(value?: types_pb.Error): ListVPNConcentratorsResponse;
  hasErr(): boolean;
  clearErr(): ListVPNConcentratorsResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListVPNConcentratorsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListVPNConcentratorsResponse): ListVPNConcentratorsResponse.AsObject;
  static serializeBinaryToWriter(message: ListVPNConcentratorsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListVPNConcentratorsResponse;
  static deserializeBinaryFromReader(message: ListVPNConcentratorsResponse, reader: jspb.BinaryReader): ListVPNConcentratorsResponse;
}

export namespace ListVPNConcentratorsResponse {
  export type AsObject = {
    vpnConcentratorsList: Array<types_pb.VPNConcentrator.AsObject>,
    lastSyncTime: string,
    err?: types_pb.Error.AsObject,
  }
}

export class GetVPCIDForCIDRRequest extends jspb.Message {
  getProvider(): string;
  setProvider(value: string): GetVPCIDForCIDRRequest;

  getAccountId(): string;
  setAccountId(value: string): GetVPCIDForCIDRRequest;

  getRegion(): string;
  setRegion(value: string): GetVPCIDForCIDRRequest;

  getCidr(): string;
  setCidr(value: string): GetVPCIDForCIDRRequest;

  getLabelsMap(): jspb.Map<string, string>;
  clearLabelsMap(): GetVPCIDForCIDRRequest;

  getCreds(): types_pb.Credentials | undefined;
  setCreds(value?: types_pb.Credentials): GetVPCIDForCIDRRequest;
  hasCreds(): boolean;
  clearCreds(): GetVPCIDForCIDRRequest;

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
    accountId: string,
    region: string,
    cidr: string,
    labelsMap: Array<[string, string]>,
    creds?: types_pb.Credentials.AsObject,
  }

  export enum CredsCase { 
    _CREDS_NOT_SET = 0,
    CREDS = 6,
  }
}

export class GetVPCIDForCIDRResponse extends jspb.Message {
  getVpcId(): string;
  setVpcId(value: string): GetVPCIDForCIDRResponse;

  getErr(): types_pb.Error | undefined;
  setErr(value?: types_pb.Error): GetVPCIDForCIDRResponse;
  hasErr(): boolean;
  clearErr(): GetVPCIDForCIDRResponse;

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
    err?: types_pb.Error.AsObject,
  }
}

export class GetCIDRsForLabelsRequest extends jspb.Message {
  getProvider(): string;
  setProvider(value: string): GetCIDRsForLabelsRequest;

  getAccountId(): string;
  setAccountId(value: string): GetCIDRsForLabelsRequest;

  getRegion(): string;
  setRegion(value: string): GetCIDRsForLabelsRequest;

  getLabelsMap(): jspb.Map<string, string>;
  clearLabelsMap(): GetCIDRsForLabelsRequest;

  getCreds(): types_pb.Credentials | undefined;
  setCreds(value?: types_pb.Credentials): GetCIDRsForLabelsRequest;
  hasCreds(): boolean;
  clearCreds(): GetCIDRsForLabelsRequest;

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
    accountId: string,
    region: string,
    labelsMap: Array<[string, string]>,
    creds?: types_pb.Credentials.AsObject,
  }

  export enum CredsCase { 
    _CREDS_NOT_SET = 0,
    CREDS = 5,
  }
}

export class GetCIDRsForLabelsResponse extends jspb.Message {
  getCidrsList(): Array<string>;
  setCidrsList(value: Array<string>): GetCIDRsForLabelsResponse;
  clearCidrsList(): GetCIDRsForLabelsResponse;
  addCidrs(value: string, index?: number): GetCIDRsForLabelsResponse;

  getErr(): types_pb.Error | undefined;
  setErr(value?: types_pb.Error): GetCIDRsForLabelsResponse;
  hasErr(): boolean;
  clearErr(): GetCIDRsForLabelsResponse;

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
    err?: types_pb.Error.AsObject,
  }
}

export class GetIPsForLabelsRequest extends jspb.Message {
  getProvider(): string;
  setProvider(value: string): GetIPsForLabelsRequest;

  getAccountId(): string;
  setAccountId(value: string): GetIPsForLabelsRequest;

  getRegion(): string;
  setRegion(value: string): GetIPsForLabelsRequest;

  getLabelsMap(): jspb.Map<string, string>;
  clearLabelsMap(): GetIPsForLabelsRequest;

  getCreds(): types_pb.Credentials | undefined;
  setCreds(value?: types_pb.Credentials): GetIPsForLabelsRequest;
  hasCreds(): boolean;
  clearCreds(): GetIPsForLabelsRequest;

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
    accountId: string,
    region: string,
    labelsMap: Array<[string, string]>,
    creds?: types_pb.Credentials.AsObject,
  }

  export enum CredsCase { 
    _CREDS_NOT_SET = 0,
    CREDS = 5,
  }
}

export class GetIPsForLabelsResponse extends jspb.Message {
  getIpsList(): Array<string>;
  setIpsList(value: Array<string>): GetIPsForLabelsResponse;
  clearIpsList(): GetIPsForLabelsResponse;
  addIps(value: string, index?: number): GetIPsForLabelsResponse;

  getErr(): types_pb.Error | undefined;
  setErr(value?: types_pb.Error): GetIPsForLabelsResponse;
  hasErr(): boolean;
  clearErr(): GetIPsForLabelsResponse;

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
    err?: types_pb.Error.AsObject,
  }
}

export class GetInstancesForLabelsRequest extends jspb.Message {
  getProvider(): string;
  setProvider(value: string): GetInstancesForLabelsRequest;

  getAccountId(): string;
  setAccountId(value: string): GetInstancesForLabelsRequest;

  getRegion(): string;
  setRegion(value: string): GetInstancesForLabelsRequest;

  getVpcId(): string;
  setVpcId(value: string): GetInstancesForLabelsRequest;

  getLabelsMap(): jspb.Map<string, string>;
  clearLabelsMap(): GetInstancesForLabelsRequest;

  getCreds(): types_pb.Credentials | undefined;
  setCreds(value?: types_pb.Credentials): GetInstancesForLabelsRequest;
  hasCreds(): boolean;
  clearCreds(): GetInstancesForLabelsRequest;

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
    accountId: string,
    region: string,
    vpcId: string,
    labelsMap: Array<[string, string]>,
    creds?: types_pb.Credentials.AsObject,
  }

  export enum CredsCase { 
    _CREDS_NOT_SET = 0,
    CREDS = 6,
  }
}

export class GetInstancesForLabelsResponse extends jspb.Message {
  getInstancesList(): Array<types_pb.Instance>;
  setInstancesList(value: Array<types_pb.Instance>): GetInstancesForLabelsResponse;
  clearInstancesList(): GetInstancesForLabelsResponse;
  addInstances(value?: types_pb.Instance, index?: number): types_pb.Instance;

  getErr(): types_pb.Error | undefined;
  setErr(value?: types_pb.Error): GetInstancesForLabelsResponse;
  hasErr(): boolean;
  clearErr(): GetInstancesForLabelsResponse;

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
    err?: types_pb.Error.AsObject,
  }
}

export class GetVPCIDWithTagRequest extends jspb.Message {
  getProvider(): string;
  setProvider(value: string): GetVPCIDWithTagRequest;

  getAccountId(): string;
  setAccountId(value: string): GetVPCIDWithTagRequest;

  getRegion(): string;
  setRegion(value: string): GetVPCIDWithTagRequest;

  getKey(): string;
  setKey(value: string): GetVPCIDWithTagRequest;

  getValue(): string;
  setValue(value: string): GetVPCIDWithTagRequest;

  getLabelsMap(): jspb.Map<string, string>;
  clearLabelsMap(): GetVPCIDWithTagRequest;

  getCreds(): types_pb.Credentials | undefined;
  setCreds(value?: types_pb.Credentials): GetVPCIDWithTagRequest;
  hasCreds(): boolean;
  clearCreds(): GetVPCIDWithTagRequest;

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
    accountId: string,
    region: string,
    key: string,
    value: string,
    labelsMap: Array<[string, string]>,
    creds?: types_pb.Credentials.AsObject,
  }

  export enum CredsCase { 
    _CREDS_NOT_SET = 0,
    CREDS = 7,
  }
}

export class GetVPCIDWithTagResponse extends jspb.Message {
  getVpcId(): string;
  setVpcId(value: string): GetVPCIDWithTagResponse;

  getErr(): types_pb.Error | undefined;
  setErr(value?: types_pb.Error): GetVPCIDWithTagResponse;
  hasErr(): boolean;
  clearErr(): GetVPCIDWithTagResponse;

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
    err?: types_pb.Error.AsObject,
  }
}

export class ListCloudClustersRequest extends jspb.Message {
  getProvider(): string;
  setProvider(value: string): ListCloudClustersRequest;

  getAccountId(): string;
  setAccountId(value: string): ListCloudClustersRequest;

  getRegion(): string;
  setRegion(value: string): ListCloudClustersRequest;

  getVpcId(): string;
  setVpcId(value: string): ListCloudClustersRequest;

  getLabelsMap(): jspb.Map<string, string>;
  clearLabelsMap(): ListCloudClustersRequest;

  getCreds(): types_pb.Credentials | undefined;
  setCreds(value?: types_pb.Credentials): ListCloudClustersRequest;
  hasCreds(): boolean;
  clearCreds(): ListCloudClustersRequest;

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
    accountId: string,
    region: string,
    vpcId: string,
    labelsMap: Array<[string, string]>,
    creds?: types_pb.Credentials.AsObject,
  }

  export enum CredsCase { 
    _CREDS_NOT_SET = 0,
    CREDS = 6,
  }
}

export class ListCloudClustersResponse extends jspb.Message {
  getClustersList(): Array<types_k8s_pb.Cluster>;
  setClustersList(value: Array<types_k8s_pb.Cluster>): ListCloudClustersResponse;
  clearClustersList(): ListCloudClustersResponse;
  addClusters(value?: types_k8s_pb.Cluster, index?: number): types_k8s_pb.Cluster;

  getLastSyncTime(): string;
  setLastSyncTime(value: string): ListCloudClustersResponse;

  getErr(): types_pb.Error | undefined;
  setErr(value?: types_pb.Error): ListCloudClustersResponse;
  hasErr(): boolean;
  clearErr(): ListCloudClustersResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListCloudClustersResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListCloudClustersResponse): ListCloudClustersResponse.AsObject;
  static serializeBinaryToWriter(message: ListCloudClustersResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListCloudClustersResponse;
  static deserializeBinaryFromReader(message: ListCloudClustersResponse, reader: jspb.BinaryReader): ListCloudClustersResponse;
}

export namespace ListCloudClustersResponse {
  export type AsObject = {
    clustersList: Array<types_k8s_pb.Cluster.AsObject>,
    lastSyncTime: string,
    err?: types_pb.Error.AsObject,
  }
}

export class SummaryRequest extends jspb.Message {
  getProvider(): string;
  setProvider(value: string): SummaryRequest;

  getAccountId(): string;
  setAccountId(value: string): SummaryRequest;

  getRegion(): string;
  setRegion(value: string): SummaryRequest;

  getVpcId(): string;
  setVpcId(value: string): SummaryRequest;

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
    accountId: string,
    region: string,
    vpcId: string,
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

  getNatGateways(): number;
  setNatGateways(value: number): Counters;

  getRouters(): number;
  setRouters(value: number): Counters;

  getIgws(): number;
  setIgws(value: number): Counters;

  getVpcEndpoints(): number;
  setVpcEndpoints(value: number): Counters;

  getPublicIps(): number;
  setPublicIps(value: number): Counters;

  getInternetGateways(): number;
  setInternetGateways(value: number): Counters;

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
    natGateways: number,
    routers: number,
    igws: number,
    vpcEndpoints: number,
    publicIps: number,
    internetGateways: number,
  }
}

export class StatusSummary extends jspb.Message {
  getVmStatusMap(): jspb.Map<string, number>;
  clearVmStatusMap(): StatusSummary;

  getPodStatusMap(): jspb.Map<string, number>;
  clearPodStatusMap(): StatusSummary;

  getVmTypesMap(): jspb.Map<string, number>;
  clearVmTypesMap(): StatusSummary;

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
    vmTypesMap: Array<[string, number]>,
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

  getErr(): types_pb.Error | undefined;
  setErr(value?: types_pb.Error): SummaryResponse;
  hasErr(): boolean;
  clearErr(): SummaryResponse;

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
    err?: types_pb.Error.AsObject,
  }
}

export class SearchResourcesRequest extends jspb.Message {
  getFilterProvider(): string;
  setFilterProvider(value: string): SearchResourcesRequest;

  getFilterAccountId(): string;
  setFilterAccountId(value: string): SearchResourcesRequest;

  getFilterRegion(): string;
  setFilterRegion(value: string): SearchResourcesRequest;

  getFilterVpcId(): string;
  setFilterVpcId(value: string): SearchResourcesRequest;

  getFilterZone(): string;
  setFilterZone(value: string): SearchResourcesRequest;

  getSearchLabelsMap(): jspb.Map<string, string>;
  clearSearchLabelsMap(): SearchResourcesRequest;

  getSearchName(): string;
  setSearchName(value: string): SearchResourcesRequest;

  getSearchId(): string;
  setSearchId(value: string): SearchResourcesRequest;

  getSearchStatus(): string;
  setSearchStatus(value: string): SearchResourcesRequest;

  getSearchCreationTimeStart(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setSearchCreationTimeStart(value?: google_protobuf_timestamp_pb.Timestamp): SearchResourcesRequest;
  hasSearchCreationTimeStart(): boolean;
  clearSearchCreationTimeStart(): SearchResourcesRequest;

  getSearchCreationTimeEnd(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setSearchCreationTimeEnd(value?: google_protobuf_timestamp_pb.Timestamp): SearchResourcesRequest;
  hasSearchCreationTimeEnd(): boolean;
  clearSearchCreationTimeEnd(): SearchResourcesRequest;

  getSearchTerminationTimeStart(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setSearchTerminationTimeStart(value?: google_protobuf_timestamp_pb.Timestamp): SearchResourcesRequest;
  hasSearchTerminationTimeStart(): boolean;
  clearSearchTerminationTimeStart(): SearchResourcesRequest;

  getSearchTerminationTimeEnd(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setSearchTerminationTimeEnd(value?: google_protobuf_timestamp_pb.Timestamp): SearchResourcesRequest;
  hasSearchTerminationTimeEnd(): boolean;
  clearSearchTerminationTimeEnd(): SearchResourcesRequest;

  getPageSize(): number;
  setPageSize(value: number): SearchResourcesRequest;

  getPageNumber(): number;
  setPageNumber(value: number): SearchResourcesRequest;

  getSortBy(): string;
  setSortBy(value: string): SearchResourcesRequest;

  getSortDescending(): boolean;
  setSortDescending(value: boolean): SearchResourcesRequest;

  getFieldMask(): google_protobuf_field_mask_pb.FieldMask | undefined;
  setFieldMask(value?: google_protobuf_field_mask_pb.FieldMask): SearchResourcesRequest;
  hasFieldMask(): boolean;
  clearFieldMask(): SearchResourcesRequest;

  getResourceTypesList(): Array<string>;
  setResourceTypesList(value: Array<string>): SearchResourcesRequest;
  clearResourceTypesList(): SearchResourcesRequest;
  addResourceTypes(value: string, index?: number): SearchResourcesRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SearchResourcesRequest.AsObject;
  static toObject(includeInstance: boolean, msg: SearchResourcesRequest): SearchResourcesRequest.AsObject;
  static serializeBinaryToWriter(message: SearchResourcesRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SearchResourcesRequest;
  static deserializeBinaryFromReader(message: SearchResourcesRequest, reader: jspb.BinaryReader): SearchResourcesRequest;
}

export namespace SearchResourcesRequest {
  export type AsObject = {
    filterProvider: string,
    filterAccountId: string,
    filterRegion: string,
    filterVpcId: string,
    filterZone: string,
    searchLabelsMap: Array<[string, string]>,
    searchName: string,
    searchId: string,
    searchStatus: string,
    searchCreationTimeStart?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    searchCreationTimeEnd?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    searchTerminationTimeStart?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    searchTerminationTimeEnd?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    pageSize: number,
    pageNumber: number,
    sortBy: string,
    sortDescending: boolean,
    fieldMask?: google_protobuf_field_mask_pb.FieldMask.AsObject,
    resourceTypesList: Array<string>,
  }
}

export class SearchResourcesResponse extends jspb.Message {
  getInstancesList(): Array<types_pb.Instance>;
  setInstancesList(value: Array<types_pb.Instance>): SearchResourcesResponse;
  clearInstancesList(): SearchResourcesResponse;
  addInstances(value?: types_pb.Instance, index?: number): types_pb.Instance;

  getVpcsList(): Array<types_pb.VPC>;
  setVpcsList(value: Array<types_pb.VPC>): SearchResourcesResponse;
  clearVpcsList(): SearchResourcesResponse;
  addVpcs(value?: types_pb.VPC, index?: number): types_pb.VPC;

  getSubnetsList(): Array<types_pb.Subnet>;
  setSubnetsList(value: Array<types_pb.Subnet>): SearchResourcesResponse;
  clearSubnetsList(): SearchResourcesResponse;
  addSubnets(value?: types_pb.Subnet, index?: number): types_pb.Subnet;

  getAclsList(): Array<types_pb.ACL>;
  setAclsList(value: Array<types_pb.ACL>): SearchResourcesResponse;
  clearAclsList(): SearchResourcesResponse;
  addAcls(value?: types_pb.ACL, index?: number): types_pb.ACL;

  getSecurityGroupsList(): Array<types_pb.SecurityGroup>;
  setSecurityGroupsList(value: Array<types_pb.SecurityGroup>): SearchResourcesResponse;
  clearSecurityGroupsList(): SearchResourcesResponse;
  addSecurityGroups(value?: types_pb.SecurityGroup, index?: number): types_pb.SecurityGroup;

  getRouteTablesList(): Array<types_pb.RouteTable>;
  setRouteTablesList(value: Array<types_pb.RouteTable>): SearchResourcesResponse;
  clearRouteTablesList(): SearchResourcesResponse;
  addRouteTables(value?: types_pb.RouteTable, index?: number): types_pb.RouteTable;

  getNatGatewaysList(): Array<types_pb.NATGateway>;
  setNatGatewaysList(value: Array<types_pb.NATGateway>): SearchResourcesResponse;
  clearNatGatewaysList(): SearchResourcesResponse;
  addNatGateways(value?: types_pb.NATGateway, index?: number): types_pb.NATGateway;

  getRoutersList(): Array<types_pb.Router>;
  setRoutersList(value: Array<types_pb.Router>): SearchResourcesResponse;
  clearRoutersList(): SearchResourcesResponse;
  addRouters(value?: types_pb.Router, index?: number): types_pb.Router;

  getIgwsList(): Array<types_pb.IGW>;
  setIgwsList(value: Array<types_pb.IGW>): SearchResourcesResponse;
  clearIgwsList(): SearchResourcesResponse;
  addIgws(value?: types_pb.IGW, index?: number): types_pb.IGW;

  getVpcEndpointsList(): Array<types_pb.VPCEndpoint>;
  setVpcEndpointsList(value: Array<types_pb.VPCEndpoint>): SearchResourcesResponse;
  clearVpcEndpointsList(): SearchResourcesResponse;
  addVpcEndpoints(value?: types_pb.VPCEndpoint, index?: number): types_pb.VPCEndpoint;

  getPublicIpsList(): Array<types_pb.PublicIP>;
  setPublicIpsList(value: Array<types_pb.PublicIP>): SearchResourcesResponse;
  clearPublicIpsList(): SearchResourcesResponse;
  addPublicIps(value?: types_pb.PublicIP, index?: number): types_pb.PublicIP;

  getClustersList(): Array<types_k8s_pb.Cluster>;
  setClustersList(value: Array<types_k8s_pb.Cluster>): SearchResourcesResponse;
  clearClustersList(): SearchResourcesResponse;
  addClusters(value?: types_k8s_pb.Cluster, index?: number): types_k8s_pb.Cluster;

  getTotalResults(): number;
  setTotalResults(value: number): SearchResourcesResponse;

  getTotalPages(): number;
  setTotalPages(value: number): SearchResourcesResponse;

  getCurrentPage(): number;
  setCurrentPage(value: number): SearchResourcesResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SearchResourcesResponse.AsObject;
  static toObject(includeInstance: boolean, msg: SearchResourcesResponse): SearchResourcesResponse.AsObject;
  static serializeBinaryToWriter(message: SearchResourcesResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SearchResourcesResponse;
  static deserializeBinaryFromReader(message: SearchResourcesResponse, reader: jspb.BinaryReader): SearchResourcesResponse;
}

export namespace SearchResourcesResponse {
  export type AsObject = {
    instancesList: Array<types_pb.Instance.AsObject>,
    vpcsList: Array<types_pb.VPC.AsObject>,
    subnetsList: Array<types_pb.Subnet.AsObject>,
    aclsList: Array<types_pb.ACL.AsObject>,
    securityGroupsList: Array<types_pb.SecurityGroup.AsObject>,
    routeTablesList: Array<types_pb.RouteTable.AsObject>,
    natGatewaysList: Array<types_pb.NATGateway.AsObject>,
    routersList: Array<types_pb.Router.AsObject>,
    igwsList: Array<types_pb.IGW.AsObject>,
    vpcEndpointsList: Array<types_pb.VPCEndpoint.AsObject>,
    publicIpsList: Array<types_pb.PublicIP.AsObject>,
    clustersList: Array<types_k8s_pb.Cluster.AsObject>,
    totalResults: number,
    totalPages: number,
    currentPage: number,
  }
}

