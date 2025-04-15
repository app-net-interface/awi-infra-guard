import * as jspb from 'google-protobuf'

import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb';


export class Error extends jspb.Message {
  getCode(): number;
  setCode(value: number): Error;

  getErrormessage(): string;
  setErrormessage(value: string): Error;

  getServerity(): string;
  setServerity(value: string): Error;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Error.AsObject;
  static toObject(includeInstance: boolean, msg: Error): Error.AsObject;
  static serializeBinaryToWriter(message: Error, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Error;
  static deserializeBinaryFromReader(message: Error, reader: jspb.BinaryReader): Error;
}

export namespace Error {
  export type AsObject = {
    code: number,
    errormessage: string,
    serverity: string,
  }
}

export class AwsRole extends jspb.Message {
  getRoleArn(): string;
  setRoleArn(value: string): AwsRole;

  getRoleSessionName(): string;
  setRoleSessionName(value: string): AwsRole;

  getDurationSeconds(): number;
  setDurationSeconds(value: number): AwsRole;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AwsRole.AsObject;
  static toObject(includeInstance: boolean, msg: AwsRole): AwsRole.AsObject;
  static serializeBinaryToWriter(message: AwsRole, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AwsRole;
  static deserializeBinaryFromReader(message: AwsRole, reader: jspb.BinaryReader): AwsRole;
}

export namespace AwsRole {
  export type AsObject = {
    roleArn: string,
    roleSessionName: string,
    durationSeconds: number,
  }
}

export class GcpRole extends jspb.Message {
  getServiceAccount(): string;
  setServiceAccount(value: string): GcpRole;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GcpRole.AsObject;
  static toObject(includeInstance: boolean, msg: GcpRole): GcpRole.AsObject;
  static serializeBinaryToWriter(message: GcpRole, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GcpRole;
  static deserializeBinaryFromReader(message: GcpRole, reader: jspb.BinaryReader): GcpRole;
}

export namespace GcpRole {
  export type AsObject = {
    serviceAccount: string,
  }
}

export class AzureRole extends jspb.Message {
  getManagedIdentity(): string;
  setManagedIdentity(value: string): AzureRole;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AzureRole.AsObject;
  static toObject(includeInstance: boolean, msg: AzureRole): AzureRole.AsObject;
  static serializeBinaryToWriter(message: AzureRole, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AzureRole;
  static deserializeBinaryFromReader(message: AzureRole, reader: jspb.BinaryReader): AzureRole;
}

export namespace AzureRole {
  export type AsObject = {
    managedIdentity: string,
  }
}

export class RoleBasedAuth extends jspb.Message {
  getAwsRole(): AwsRole | undefined;
  setAwsRole(value?: AwsRole): RoleBasedAuth;
  hasAwsRole(): boolean;
  clearAwsRole(): RoleBasedAuth;

  getGcpRole(): GcpRole | undefined;
  setGcpRole(value?: GcpRole): RoleBasedAuth;
  hasGcpRole(): boolean;
  clearGcpRole(): RoleBasedAuth;

  getAzureRole(): AzureRole | undefined;
  setAzureRole(value?: AzureRole): RoleBasedAuth;
  hasAzureRole(): boolean;
  clearAzureRole(): RoleBasedAuth;

  getRoleCase(): RoleBasedAuth.RoleCase;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RoleBasedAuth.AsObject;
  static toObject(includeInstance: boolean, msg: RoleBasedAuth): RoleBasedAuth.AsObject;
  static serializeBinaryToWriter(message: RoleBasedAuth, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RoleBasedAuth;
  static deserializeBinaryFromReader(message: RoleBasedAuth, reader: jspb.BinaryReader): RoleBasedAuth;
}

export namespace RoleBasedAuth {
  export type AsObject = {
    awsRole?: AwsRole.AsObject,
    gcpRole?: GcpRole.AsObject,
    azureRole?: AzureRole.AsObject,
  }

  export enum RoleCase { 
    ROLE_NOT_SET = 0,
    AWS_ROLE = 1,
    GCP_ROLE = 2,
    AZURE_ROLE = 3,
  }
}

export class AwsUserAuth extends jspb.Message {
  getAccessKeyId(): string;
  setAccessKeyId(value: string): AwsUserAuth;

  getSecretAccessKey(): string;
  setSecretAccessKey(value: string): AwsUserAuth;

  getSessionToken(): string;
  setSessionToken(value: string): AwsUserAuth;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AwsUserAuth.AsObject;
  static toObject(includeInstance: boolean, msg: AwsUserAuth): AwsUserAuth.AsObject;
  static serializeBinaryToWriter(message: AwsUserAuth, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AwsUserAuth;
  static deserializeBinaryFromReader(message: AwsUserAuth, reader: jspb.BinaryReader): AwsUserAuth;
}

export namespace AwsUserAuth {
  export type AsObject = {
    accessKeyId: string,
    secretAccessKey: string,
    sessionToken: string,
  }
}

export class GcpUserAuth extends jspb.Message {
  getApiKey(): string;
  setApiKey(value: string): GcpUserAuth;

  getJsonKey(): string;
  setJsonKey(value: string): GcpUserAuth;

  getAuthMethodCase(): GcpUserAuth.AuthMethodCase;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GcpUserAuth.AsObject;
  static toObject(includeInstance: boolean, msg: GcpUserAuth): GcpUserAuth.AsObject;
  static serializeBinaryToWriter(message: GcpUserAuth, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GcpUserAuth;
  static deserializeBinaryFromReader(message: GcpUserAuth, reader: jspb.BinaryReader): GcpUserAuth;
}

export namespace GcpUserAuth {
  export type AsObject = {
    apiKey: string,
    jsonKey: string,
  }

  export enum AuthMethodCase { 
    AUTH_METHOD_NOT_SET = 0,
    API_KEY = 1,
    JSON_KEY = 2,
  }
}

export class AzureUserAuth extends jspb.Message {
  getClientId(): string;
  setClientId(value: string): AzureUserAuth;

  getClientSecret(): string;
  setClientSecret(value: string): AzureUserAuth;

  getTenantId(): string;
  setTenantId(value: string): AzureUserAuth;

  getCertificatePath(): string;
  setCertificatePath(value: string): AzureUserAuth;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AzureUserAuth.AsObject;
  static toObject(includeInstance: boolean, msg: AzureUserAuth): AzureUserAuth.AsObject;
  static serializeBinaryToWriter(message: AzureUserAuth, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AzureUserAuth;
  static deserializeBinaryFromReader(message: AzureUserAuth, reader: jspb.BinaryReader): AzureUserAuth;
}

export namespace AzureUserAuth {
  export type AsObject = {
    clientId: string,
    clientSecret: string,
    tenantId: string,
    certificatePath: string,
  }
}

export class UserBasedAuth extends jspb.Message {
  getAwsUserAuth(): AwsUserAuth | undefined;
  setAwsUserAuth(value?: AwsUserAuth): UserBasedAuth;
  hasAwsUserAuth(): boolean;
  clearAwsUserAuth(): UserBasedAuth;

  getGcpUserAuth(): GcpUserAuth | undefined;
  setGcpUserAuth(value?: GcpUserAuth): UserBasedAuth;
  hasGcpUserAuth(): boolean;
  clearGcpUserAuth(): UserBasedAuth;

  getAzureUserAuth(): AzureUserAuth | undefined;
  setAzureUserAuth(value?: AzureUserAuth): UserBasedAuth;
  hasAzureUserAuth(): boolean;
  clearAzureUserAuth(): UserBasedAuth;

  getUserAuthCase(): UserBasedAuth.UserAuthCase;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UserBasedAuth.AsObject;
  static toObject(includeInstance: boolean, msg: UserBasedAuth): UserBasedAuth.AsObject;
  static serializeBinaryToWriter(message: UserBasedAuth, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UserBasedAuth;
  static deserializeBinaryFromReader(message: UserBasedAuth, reader: jspb.BinaryReader): UserBasedAuth;
}

export namespace UserBasedAuth {
  export type AsObject = {
    awsUserAuth?: AwsUserAuth.AsObject,
    gcpUserAuth?: GcpUserAuth.AsObject,
    azureUserAuth?: AzureUserAuth.AsObject,
  }

  export enum UserAuthCase { 
    USER_AUTH_NOT_SET = 0,
    AWS_USER_AUTH = 1,
    GCP_USER_AUTH = 2,
    AZURE_USER_AUTH = 3,
  }
}

export class Credentials extends jspb.Message {
  getRoleBasedAuth(): RoleBasedAuth | undefined;
  setRoleBasedAuth(value?: RoleBasedAuth): Credentials;
  hasRoleBasedAuth(): boolean;
  clearRoleBasedAuth(): Credentials;

  getUserBasedAuth(): UserBasedAuth | undefined;
  setUserBasedAuth(value?: UserBasedAuth): Credentials;
  hasUserBasedAuth(): boolean;
  clearUserBasedAuth(): Credentials;

  getAuthCase(): Credentials.AuthCase;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Credentials.AsObject;
  static toObject(includeInstance: boolean, msg: Credentials): Credentials.AsObject;
  static serializeBinaryToWriter(message: Credentials, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Credentials;
  static deserializeBinaryFromReader(message: Credentials, reader: jspb.BinaryReader): Credentials;
}

export namespace Credentials {
  export type AsObject = {
    roleBasedAuth?: RoleBasedAuth.AsObject,
    userBasedAuth?: UserBasedAuth.AsObject,
  }

  export enum AuthCase { 
    AUTH_NOT_SET = 0,
    ROLE_BASED_AUTH = 1,
    USER_BASED_AUTH = 2,
  }
}

export class Account extends jspb.Message {
  getProvider(): string;
  setProvider(value: string): Account;

  getId(): string;
  setId(value: string): Account;

  getName(): string;
  setName(value: string): Account;

  getLastSyncTime(): string;
  setLastSyncTime(value: string): Account;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Account.AsObject;
  static toObject(includeInstance: boolean, msg: Account): Account.AsObject;
  static serializeBinaryToWriter(message: Account, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Account;
  static deserializeBinaryFromReader(message: Account, reader: jspb.BinaryReader): Account;
}

export namespace Account {
  export type AsObject = {
    provider: string,
    id: string,
    name: string,
    lastSyncTime: string,
  }
}

export class Region extends jspb.Message {
  getProvider(): string;
  setProvider(value: string): Region;

  getId(): string;
  setId(value: string): Region;

  getName(): string;
  setName(value: string): Region;

  getLastSyncTime(): string;
  setLastSyncTime(value: string): Region;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Region.AsObject;
  static toObject(includeInstance: boolean, msg: Region): Region.AsObject;
  static serializeBinaryToWriter(message: Region, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Region;
  static deserializeBinaryFromReader(message: Region, reader: jspb.BinaryReader): Region;
}

export namespace Region {
  export type AsObject = {
    provider: string,
    id: string,
    name: string,
    lastSyncTime: string,
  }
}

export class VPC extends jspb.Message {
  getId(): string;
  setId(value: string): VPC;

  getName(): string;
  setName(value: string): VPC;

  getRegion(): string;
  setRegion(value: string): VPC;

  getLabelsMap(): jspb.Map<string, string>;
  clearLabelsMap(): VPC;

  getIpv4Cidr(): string;
  setIpv4Cidr(value: string): VPC;

  getIpv6Cidr(): string;
  setIpv6Cidr(value: string): VPC;

  getProject(): string;
  setProject(value: string): VPC;

  getProvider(): string;
  setProvider(value: string): VPC;

  getAccountId(): string;
  setAccountId(value: string): VPC;

  getLastSyncTime(): string;
  setLastSyncTime(value: string): VPC;

  getSelfLink(): string;
  setSelfLink(value: string): VPC;

  getCreatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreatedAt(value?: google_protobuf_timestamp_pb.Timestamp): VPC;
  hasCreatedAt(): boolean;
  clearCreatedAt(): VPC;

  getUpdatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setUpdatedAt(value?: google_protobuf_timestamp_pb.Timestamp): VPC;
  hasUpdatedAt(): boolean;
  clearUpdatedAt(): VPC;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): VPC.AsObject;
  static toObject(includeInstance: boolean, msg: VPC): VPC.AsObject;
  static serializeBinaryToWriter(message: VPC, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): VPC;
  static deserializeBinaryFromReader(message: VPC, reader: jspb.BinaryReader): VPC;
}

export namespace VPC {
  export type AsObject = {
    id: string,
    name: string,
    region: string,
    labelsMap: Array<[string, string]>,
    ipv4Cidr: string,
    ipv6Cidr: string,
    project: string,
    provider: string,
    accountId: string,
    lastSyncTime: string,
    selfLink: string,
    createdAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    updatedAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }
}

export class VPCIndex extends jspb.Message {
  getVpcId(): string;
  setVpcId(value: string): VPCIndex;

  getInstanceIdsList(): Array<string>;
  setInstanceIdsList(value: Array<string>): VPCIndex;
  clearInstanceIdsList(): VPCIndex;
  addInstanceIds(value: string, index?: number): VPCIndex;

  getAclIdsList(): Array<string>;
  setAclIdsList(value: Array<string>): VPCIndex;
  clearAclIdsList(): VPCIndex;
  addAclIds(value: string, index?: number): VPCIndex;

  getSecurityGroupIdsList(): Array<string>;
  setSecurityGroupIdsList(value: Array<string>): VPCIndex;
  clearSecurityGroupIdsList(): VPCIndex;
  addSecurityGroupIds(value: string, index?: number): VPCIndex;

  getNatGatewayIdsList(): Array<string>;
  setNatGatewayIdsList(value: Array<string>): VPCIndex;
  clearNatGatewayIdsList(): VPCIndex;
  addNatGatewayIds(value: string, index?: number): VPCIndex;

  getVpcEndpointIdsList(): Array<string>;
  setVpcEndpointIdsList(value: Array<string>): VPCIndex;
  clearVpcEndpointIdsList(): VPCIndex;
  addVpcEndpointIds(value: string, index?: number): VPCIndex;

  getLbIdsList(): Array<string>;
  setLbIdsList(value: Array<string>): VPCIndex;
  clearLbIdsList(): VPCIndex;
  addLbIds(value: string, index?: number): VPCIndex;

  getRouterIdsList(): Array<string>;
  setRouterIdsList(value: Array<string>): VPCIndex;
  clearRouterIdsList(): VPCIndex;
  addRouterIds(value: string, index?: number): VPCIndex;

  getIgwIdsList(): Array<string>;
  setIgwIdsList(value: Array<string>): VPCIndex;
  clearIgwIdsList(): VPCIndex;
  addIgwIds(value: string, index?: number): VPCIndex;

  getSubnetIdsList(): Array<string>;
  setSubnetIdsList(value: Array<string>): VPCIndex;
  clearSubnetIdsList(): VPCIndex;
  addSubnetIds(value: string, index?: number): VPCIndex;

  getRouteTableIdsList(): Array<string>;
  setRouteTableIdsList(value: Array<string>): VPCIndex;
  clearRouteTableIdsList(): VPCIndex;
  addRouteTableIds(value: string, index?: number): VPCIndex;

  getNetworkInterfaceIdsList(): Array<string>;
  setNetworkInterfaceIdsList(value: Array<string>): VPCIndex;
  clearNetworkInterfaceIdsList(): VPCIndex;
  addNetworkInterfaceIds(value: string, index?: number): VPCIndex;

  getKeyPairIdsList(): Array<string>;
  setKeyPairIdsList(value: Array<string>): VPCIndex;
  clearKeyPairIdsList(): VPCIndex;
  addKeyPairIds(value: string, index?: number): VPCIndex;

  getVpnConcentratorIdsList(): Array<string>;
  setVpnConcentratorIdsList(value: Array<string>): VPCIndex;
  clearVpnConcentratorIdsList(): VPCIndex;
  addVpnConcentratorIds(value: string, index?: number): VPCIndex;

  getPublicIpIdsList(): Array<string>;
  setPublicIpIdsList(value: Array<string>): VPCIndex;
  clearPublicIpIdsList(): VPCIndex;
  addPublicIpIds(value: string, index?: number): VPCIndex;

  getClusterIdsList(): Array<string>;
  setClusterIdsList(value: Array<string>): VPCIndex;
  clearClusterIdsList(): VPCIndex;
  addClusterIds(value: string, index?: number): VPCIndex;

  getLastSyncTime(): string;
  setLastSyncTime(value: string): VPCIndex;

  getProvider(): string;
  setProvider(value: string): VPCIndex;

  getAccountId(): string;
  setAccountId(value: string): VPCIndex;

  getRegion(): string;
  setRegion(value: string): VPCIndex;

  getCreatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreatedAt(value?: google_protobuf_timestamp_pb.Timestamp): VPCIndex;
  hasCreatedAt(): boolean;
  clearCreatedAt(): VPCIndex;

  getUpdatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setUpdatedAt(value?: google_protobuf_timestamp_pb.Timestamp): VPCIndex;
  hasUpdatedAt(): boolean;
  clearUpdatedAt(): VPCIndex;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): VPCIndex.AsObject;
  static toObject(includeInstance: boolean, msg: VPCIndex): VPCIndex.AsObject;
  static serializeBinaryToWriter(message: VPCIndex, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): VPCIndex;
  static deserializeBinaryFromReader(message: VPCIndex, reader: jspb.BinaryReader): VPCIndex;
}

export namespace VPCIndex {
  export type AsObject = {
    vpcId: string,
    instanceIdsList: Array<string>,
    aclIdsList: Array<string>,
    securityGroupIdsList: Array<string>,
    natGatewayIdsList: Array<string>,
    vpcEndpointIdsList: Array<string>,
    lbIdsList: Array<string>,
    routerIdsList: Array<string>,
    igwIdsList: Array<string>,
    subnetIdsList: Array<string>,
    routeTableIdsList: Array<string>,
    networkInterfaceIdsList: Array<string>,
    keyPairIdsList: Array<string>,
    vpnConcentratorIdsList: Array<string>,
    publicIpIdsList: Array<string>,
    clusterIdsList: Array<string>,
    lastSyncTime: string,
    provider: string,
    accountId: string,
    region: string,
    createdAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    updatedAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }
}

export class Instance extends jspb.Message {
  getProvider(): string;
  setProvider(value: string): Instance;

  getAccountId(): string;
  setAccountId(value: string): Instance;

  getId(): string;
  setId(value: string): Instance;

  getName(): string;
  setName(value: string): Instance;

  getSubnetid(): string;
  setSubnetid(value: string): Instance;

  getProject(): string;
  setProject(value: string): Instance;

  getVpcid(): string;
  setVpcid(value: string): Instance;

  getRegion(): string;
  setRegion(value: string): Instance;

  getZone(): string;
  setZone(value: string): Instance;

  getPublicip(): string;
  setPublicip(value: string): Instance;

  getPrivateip(): string;
  setPrivateip(value: string): Instance;

  getSecuritygroupidsList(): Array<string>;
  setSecuritygroupidsList(value: Array<string>): Instance;
  clearSecuritygroupidsList(): Instance;
  addSecuritygroupids(value: string, index?: number): Instance;

  getInterfaceidsList(): Array<string>;
  setInterfaceidsList(value: Array<string>): Instance;
  clearInterfaceidsList(): Instance;
  addInterfaceids(value: string, index?: number): Instance;

  getState(): string;
  setState(value: string): Instance;

  getType(): string;
  setType(value: string): Instance;

  getLabelsMap(): jspb.Map<string, string>;
  clearLabelsMap(): Instance;

  getCreatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreatedAt(value?: google_protobuf_timestamp_pb.Timestamp): Instance;
  hasCreatedAt(): boolean;
  clearCreatedAt(): Instance;

  getUpdatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setUpdatedAt(value?: google_protobuf_timestamp_pb.Timestamp): Instance;
  hasUpdatedAt(): boolean;
  clearUpdatedAt(): Instance;

  getLastSyncTime(): string;
  setLastSyncTime(value: string): Instance;

  getSelfLink(): string;
  setSelfLink(value: string): Instance;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Instance.AsObject;
  static toObject(includeInstance: boolean, msg: Instance): Instance.AsObject;
  static serializeBinaryToWriter(message: Instance, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Instance;
  static deserializeBinaryFromReader(message: Instance, reader: jspb.BinaryReader): Instance;
}

export namespace Instance {
  export type AsObject = {
    provider: string,
    accountId: string,
    id: string,
    name: string,
    subnetid: string,
    project: string,
    vpcid: string,
    region: string,
    zone: string,
    publicip: string,
    privateip: string,
    securitygroupidsList: Array<string>,
    interfaceidsList: Array<string>,
    state: string,
    type: string,
    labelsMap: Array<[string, string]>,
    createdAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    updatedAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    lastSyncTime: string,
    selfLink: string,
  }
}

export class Subnet extends jspb.Message {
  getId(): string;
  setId(value: string): Subnet;

  getName(): string;
  setName(value: string): Subnet;

  getCidrBlock(): string;
  setCidrBlock(value: string): Subnet;

  getVpcId(): string;
  setVpcId(value: string): Subnet;

  getZone(): string;
  setZone(value: string): Subnet;

  getRegion(): string;
  setRegion(value: string): Subnet;

  getLabelsMap(): jspb.Map<string, string>;
  clearLabelsMap(): Subnet;

  getProject(): string;
  setProject(value: string): Subnet;

  getProvider(): string;
  setProvider(value: string): Subnet;

  getAccountId(): string;
  setAccountId(value: string): Subnet;

  getLastSyncTime(): string;
  setLastSyncTime(value: string): Subnet;

  getSelfLink(): string;
  setSelfLink(value: string): Subnet;

  getCreatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreatedAt(value?: google_protobuf_timestamp_pb.Timestamp): Subnet;
  hasCreatedAt(): boolean;
  clearCreatedAt(): Subnet;

  getUpdatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setUpdatedAt(value?: google_protobuf_timestamp_pb.Timestamp): Subnet;
  hasUpdatedAt(): boolean;
  clearUpdatedAt(): Subnet;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Subnet.AsObject;
  static toObject(includeInstance: boolean, msg: Subnet): Subnet.AsObject;
  static serializeBinaryToWriter(message: Subnet, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Subnet;
  static deserializeBinaryFromReader(message: Subnet, reader: jspb.BinaryReader): Subnet;
}

export namespace Subnet {
  export type AsObject = {
    id: string,
    name: string,
    cidrBlock: string,
    vpcId: string,
    zone: string,
    region: string,
    labelsMap: Array<[string, string]>,
    project: string,
    provider: string,
    accountId: string,
    lastSyncTime: string,
    selfLink: string,
    createdAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    updatedAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }
}

export class ACL extends jspb.Message {
  getProvider(): string;
  setProvider(value: string): ACL;

  getId(): string;
  setId(value: string): ACL;

  getName(): string;
  setName(value: string): ACL;

  getVpcId(): string;
  setVpcId(value: string): ACL;

  getRegion(): string;
  setRegion(value: string): ACL;

  getAccountId(): string;
  setAccountId(value: string): ACL;

  getRulesList(): Array<ACL.ACLRule>;
  setRulesList(value: Array<ACL.ACLRule>): ACL;
  clearRulesList(): ACL;
  addRules(value?: ACL.ACLRule, index?: number): ACL.ACLRule;

  getLabelsMap(): jspb.Map<string, string>;
  clearLabelsMap(): ACL;

  getProject(): string;
  setProject(value: string): ACL;

  getLastSyncTime(): string;
  setLastSyncTime(value: string): ACL;

  getSelfLink(): string;
  setSelfLink(value: string): ACL;

  getCreatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreatedAt(value?: google_protobuf_timestamp_pb.Timestamp): ACL;
  hasCreatedAt(): boolean;
  clearCreatedAt(): ACL;

  getUpdatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setUpdatedAt(value?: google_protobuf_timestamp_pb.Timestamp): ACL;
  hasUpdatedAt(): boolean;
  clearUpdatedAt(): ACL;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ACL.AsObject;
  static toObject(includeInstance: boolean, msg: ACL): ACL.AsObject;
  static serializeBinaryToWriter(message: ACL, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ACL;
  static deserializeBinaryFromReader(message: ACL, reader: jspb.BinaryReader): ACL;
}

export namespace ACL {
  export type AsObject = {
    provider: string,
    id: string,
    name: string,
    vpcId: string,
    region: string,
    accountId: string,
    rulesList: Array<ACL.ACLRule.AsObject>,
    labelsMap: Array<[string, string]>,
    project: string,
    lastSyncTime: string,
    selfLink: string,
    createdAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    updatedAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }

  export class ACLRule extends jspb.Message {
    getNumber(): number;
    setNumber(value: number): ACLRule;

    getProtocol(): string;
    setProtocol(value: string): ACLRule;

    getPortRange(): string;
    setPortRange(value: string): ACLRule;

    getSourceRangesList(): Array<string>;
    setSourceRangesList(value: Array<string>): ACLRule;
    clearSourceRangesList(): ACLRule;
    addSourceRanges(value: string, index?: number): ACLRule;

    getDestinationRangesList(): Array<string>;
    setDestinationRangesList(value: Array<string>): ACLRule;
    clearDestinationRangesList(): ACLRule;
    addDestinationRanges(value: string, index?: number): ACLRule;

    getAction(): string;
    setAction(value: string): ACLRule;

    getDirection(): string;
    setDirection(value: string): ACLRule;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ACLRule.AsObject;
    static toObject(includeInstance: boolean, msg: ACLRule): ACLRule.AsObject;
    static serializeBinaryToWriter(message: ACLRule, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ACLRule;
    static deserializeBinaryFromReader(message: ACLRule, reader: jspb.BinaryReader): ACLRule;
  }

  export namespace ACLRule {
    export type AsObject = {
      number: number,
      protocol: string,
      portRange: string,
      sourceRangesList: Array<string>,
      destinationRangesList: Array<string>,
      action: string,
      direction: string,
    }
  }

}

export class SecurityGroup extends jspb.Message {
  getProvider(): string;
  setProvider(value: string): SecurityGroup;

  getId(): string;
  setId(value: string): SecurityGroup;

  getName(): string;
  setName(value: string): SecurityGroup;

  getVpcId(): string;
  setVpcId(value: string): SecurityGroup;

  getRegion(): string;
  setRegion(value: string): SecurityGroup;

  getAccountId(): string;
  setAccountId(value: string): SecurityGroup;

  getRulesList(): Array<SecurityGroup.SecurityGroupRule>;
  setRulesList(value: Array<SecurityGroup.SecurityGroupRule>): SecurityGroup;
  clearRulesList(): SecurityGroup;
  addRules(value?: SecurityGroup.SecurityGroupRule, index?: number): SecurityGroup.SecurityGroupRule;

  getLabelsMap(): jspb.Map<string, string>;
  clearLabelsMap(): SecurityGroup;

  getProject(): string;
  setProject(value: string): SecurityGroup;

  getLastSyncTime(): string;
  setLastSyncTime(value: string): SecurityGroup;

  getSelfLink(): string;
  setSelfLink(value: string): SecurityGroup;

  getAttachedRunningInstancesList(): Array<string>;
  setAttachedRunningInstancesList(value: Array<string>): SecurityGroup;
  clearAttachedRunningInstancesList(): SecurityGroup;
  addAttachedRunningInstances(value: string, index?: number): SecurityGroup;

  getCreatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreatedAt(value?: google_protobuf_timestamp_pb.Timestamp): SecurityGroup;
  hasCreatedAt(): boolean;
  clearCreatedAt(): SecurityGroup;

  getUpdatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setUpdatedAt(value?: google_protobuf_timestamp_pb.Timestamp): SecurityGroup;
  hasUpdatedAt(): boolean;
  clearUpdatedAt(): SecurityGroup;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SecurityGroup.AsObject;
  static toObject(includeInstance: boolean, msg: SecurityGroup): SecurityGroup.AsObject;
  static serializeBinaryToWriter(message: SecurityGroup, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SecurityGroup;
  static deserializeBinaryFromReader(message: SecurityGroup, reader: jspb.BinaryReader): SecurityGroup;
}

export namespace SecurityGroup {
  export type AsObject = {
    provider: string,
    id: string,
    name: string,
    vpcId: string,
    region: string,
    accountId: string,
    rulesList: Array<SecurityGroup.SecurityGroupRule.AsObject>,
    labelsMap: Array<[string, string]>,
    project: string,
    lastSyncTime: string,
    selfLink: string,
    attachedRunningInstancesList: Array<string>,
    createdAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    updatedAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }

  export class SecurityGroupRule extends jspb.Message {
    getProtocol(): string;
    setProtocol(value: string): SecurityGroupRule;

    getPortRange(): string;
    setPortRange(value: string): SecurityGroupRule;

    getSourceList(): Array<string>;
    setSourceList(value: Array<string>): SecurityGroupRule;
    clearSourceList(): SecurityGroupRule;
    addSource(value: string, index?: number): SecurityGroupRule;

    getDirection(): string;
    setDirection(value: string): SecurityGroupRule;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): SecurityGroupRule.AsObject;
    static toObject(includeInstance: boolean, msg: SecurityGroupRule): SecurityGroupRule.AsObject;
    static serializeBinaryToWriter(message: SecurityGroupRule, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): SecurityGroupRule;
    static deserializeBinaryFromReader(message: SecurityGroupRule, reader: jspb.BinaryReader): SecurityGroupRule;
  }

  export namespace SecurityGroupRule {
    export type AsObject = {
      protocol: string,
      portRange: string,
      sourceList: Array<string>,
      direction: string,
    }
  }

}

export class RouteTable extends jspb.Message {
  getProvider(): string;
  setProvider(value: string): RouteTable;

  getId(): string;
  setId(value: string): RouteTable;

  getName(): string;
  setName(value: string): RouteTable;

  getVpcId(): string;
  setVpcId(value: string): RouteTable;

  getRegion(): string;
  setRegion(value: string): RouteTable;

  getAccountId(): string;
  setAccountId(value: string): RouteTable;

  getRoutesList(): Array<RouteTable.Route>;
  setRoutesList(value: Array<RouteTable.Route>): RouteTable;
  clearRoutesList(): RouteTable;
  addRoutes(value?: RouteTable.Route, index?: number): RouteTable.Route;

  getSubnetIdsList(): Array<string>;
  setSubnetIdsList(value: Array<string>): RouteTable;
  clearSubnetIdsList(): RouteTable;
  addSubnetIds(value: string, index?: number): RouteTable;

  getGatewayIdsList(): Array<string>;
  setGatewayIdsList(value: Array<string>): RouteTable;
  clearGatewayIdsList(): RouteTable;
  addGatewayIds(value: string, index?: number): RouteTable;

  getLabelsMap(): jspb.Map<string, string>;
  clearLabelsMap(): RouteTable;

  getProject(): string;
  setProject(value: string): RouteTable;

  getLastSyncTime(): string;
  setLastSyncTime(value: string): RouteTable;

  getSelfLink(): string;
  setSelfLink(value: string): RouteTable;

  getCreatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreatedAt(value?: google_protobuf_timestamp_pb.Timestamp): RouteTable;
  hasCreatedAt(): boolean;
  clearCreatedAt(): RouteTable;

  getUpdatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setUpdatedAt(value?: google_protobuf_timestamp_pb.Timestamp): RouteTable;
  hasUpdatedAt(): boolean;
  clearUpdatedAt(): RouteTable;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RouteTable.AsObject;
  static toObject(includeInstance: boolean, msg: RouteTable): RouteTable.AsObject;
  static serializeBinaryToWriter(message: RouteTable, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RouteTable;
  static deserializeBinaryFromReader(message: RouteTable, reader: jspb.BinaryReader): RouteTable;
}

export namespace RouteTable {
  export type AsObject = {
    provider: string,
    id: string,
    name: string,
    vpcId: string,
    region: string,
    accountId: string,
    routesList: Array<RouteTable.Route.AsObject>,
    subnetIdsList: Array<string>,
    gatewayIdsList: Array<string>,
    labelsMap: Array<[string, string]>,
    project: string,
    lastSyncTime: string,
    selfLink: string,
    createdAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    updatedAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }

  export class Route extends jspb.Message {
    getDestination(): string;
    setDestination(value: string): Route;

    getTarget(): string;
    setTarget(value: string): Route;

    getStatus(): string;
    setStatus(value: string): Route;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Route.AsObject;
    static toObject(includeInstance: boolean, msg: Route): Route.AsObject;
    static serializeBinaryToWriter(message: Route, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Route;
    static deserializeBinaryFromReader(message: Route, reader: jspb.BinaryReader): Route;
  }

  export namespace Route {
    export type AsObject = {
      destination: string,
      target: string,
      status: string,
    }
  }

}

export class Router extends jspb.Message {
  getId(): string;
  setId(value: string): Router;

  getName(): string;
  setName(value: string): Router;

  getProvider(): string;
  setProvider(value: string): Router;

  getRegion(): string;
  setRegion(value: string): Router;

  getVpcId(): string;
  setVpcId(value: string): Router;

  getState(): string;
  setState(value: string): Router;

  getAsn(): number;
  setAsn(value: number): Router;

  getAdvertisedRange(): string;
  setAdvertisedRange(value: string): Router;

  getAdvertisedGroup(): string;
  setAdvertisedGroup(value: string): Router;

  getVpnType(): string;
  setVpnType(value: string): Router;

  getSubnetId(): string;
  setSubnetId(value: string): Router;

  getLabelsMap(): jspb.Map<string, string>;
  clearLabelsMap(): Router;

  getCreatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreatedAt(value?: google_protobuf_timestamp_pb.Timestamp): Router;
  hasCreatedAt(): boolean;
  clearCreatedAt(): Router;

  getUpdatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setUpdatedAt(value?: google_protobuf_timestamp_pb.Timestamp): Router;
  hasUpdatedAt(): boolean;
  clearUpdatedAt(): Router;

  getAccountId(): string;
  setAccountId(value: string): Router;

  getProject(): string;
  setProject(value: string): Router;

  getLastSyncTime(): string;
  setLastSyncTime(value: string): Router;

  getAdditionalPropertiesMap(): jspb.Map<string, string>;
  clearAdditionalPropertiesMap(): Router;

  getSelfLink(): string;
  setSelfLink(value: string): Router;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Router.AsObject;
  static toObject(includeInstance: boolean, msg: Router): Router.AsObject;
  static serializeBinaryToWriter(message: Router, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Router;
  static deserializeBinaryFromReader(message: Router, reader: jspb.BinaryReader): Router;
}

export namespace Router {
  export type AsObject = {
    id: string,
    name: string,
    provider: string,
    region: string,
    vpcId: string,
    state: string,
    asn: number,
    advertisedRange: string,
    advertisedGroup: string,
    vpnType: string,
    subnetId: string,
    labelsMap: Array<[string, string]>,
    createdAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    updatedAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    accountId: string,
    project: string,
    lastSyncTime: string,
    additionalPropertiesMap: Array<[string, string]>,
    selfLink: string,
  }
}

export class NATGateway extends jspb.Message {
  getId(): string;
  setId(value: string): NATGateway;

  getName(): string;
  setName(value: string): NATGateway;

  getProvider(): string;
  setProvider(value: string): NATGateway;

  getAccountId(): string;
  setAccountId(value: string): NATGateway;

  getVpcId(): string;
  setVpcId(value: string): NATGateway;

  getRegion(): string;
  setRegion(value: string): NATGateway;

  getState(): string;
  setState(value: string): NATGateway;

  getPublicIp(): string;
  setPublicIp(value: string): NATGateway;

  getPrivateIp(): string;
  setPrivateIp(value: string): NATGateway;

  getSubnetId(): string;
  setSubnetId(value: string): NATGateway;

  getProject(): string;
  setProject(value: string): NATGateway;

  getLabelsMap(): jspb.Map<string, string>;
  clearLabelsMap(): NATGateway;

  getCreatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreatedAt(value?: google_protobuf_timestamp_pb.Timestamp): NATGateway;
  hasCreatedAt(): boolean;
  clearCreatedAt(): NATGateway;

  getUpdatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setUpdatedAt(value?: google_protobuf_timestamp_pb.Timestamp): NATGateway;
  hasUpdatedAt(): boolean;
  clearUpdatedAt(): NATGateway;

  getLastSyncTime(): string;
  setLastSyncTime(value: string): NATGateway;

  getAdditionalPropertiesMap(): jspb.Map<string, string>;
  clearAdditionalPropertiesMap(): NATGateway;

  getSelfLink(): string;
  setSelfLink(value: string): NATGateway;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): NATGateway.AsObject;
  static toObject(includeInstance: boolean, msg: NATGateway): NATGateway.AsObject;
  static serializeBinaryToWriter(message: NATGateway, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): NATGateway;
  static deserializeBinaryFromReader(message: NATGateway, reader: jspb.BinaryReader): NATGateway;
}

export namespace NATGateway {
  export type AsObject = {
    id: string,
    name: string,
    provider: string,
    accountId: string,
    vpcId: string,
    region: string,
    state: string,
    publicIp: string,
    privateIp: string,
    subnetId: string,
    project: string,
    labelsMap: Array<[string, string]>,
    createdAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    updatedAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    lastSyncTime: string,
    additionalPropertiesMap: Array<[string, string]>,
    selfLink: string,
  }
}

export class IGW extends jspb.Message {
  getId(): string;
  setId(value: string): IGW;

  getName(): string;
  setName(value: string): IGW;

  getProvider(): string;
  setProvider(value: string): IGW;

  getAccountId(): string;
  setAccountId(value: string): IGW;

  getAttachedVpcId(): string;
  setAttachedVpcId(value: string): IGW;

  getRegion(): string;
  setRegion(value: string): IGW;

  getState(): string;
  setState(value: string): IGW;

  getProject(): string;
  setProject(value: string): IGW;

  getLabelsMap(): jspb.Map<string, string>;
  clearLabelsMap(): IGW;

  getCreatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreatedAt(value?: google_protobuf_timestamp_pb.Timestamp): IGW;
  hasCreatedAt(): boolean;
  clearCreatedAt(): IGW;

  getUpdatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setUpdatedAt(value?: google_protobuf_timestamp_pb.Timestamp): IGW;
  hasUpdatedAt(): boolean;
  clearUpdatedAt(): IGW;

  getLastSyncTime(): string;
  setLastSyncTime(value: string): IGW;

  getSelfLink(): string;
  setSelfLink(value: string): IGW;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): IGW.AsObject;
  static toObject(includeInstance: boolean, msg: IGW): IGW.AsObject;
  static serializeBinaryToWriter(message: IGW, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): IGW;
  static deserializeBinaryFromReader(message: IGW, reader: jspb.BinaryReader): IGW;
}

export namespace IGW {
  export type AsObject = {
    id: string,
    name: string,
    provider: string,
    accountId: string,
    attachedVpcId: string,
    region: string,
    state: string,
    project: string,
    labelsMap: Array<[string, string]>,
    createdAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    updatedAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    lastSyncTime: string,
    selfLink: string,
  }
}

export class VPCEndpoint extends jspb.Message {
  getId(): string;
  setId(value: string): VPCEndpoint;

  getName(): string;
  setName(value: string): VPCEndpoint;

  getProvider(): string;
  setProvider(value: string): VPCEndpoint;

  getAccountId(): string;
  setAccountId(value: string): VPCEndpoint;

  getVpcId(): string;
  setVpcId(value: string): VPCEndpoint;

  getRegion(): string;
  setRegion(value: string): VPCEndpoint;

  getState(): string;
  setState(value: string): VPCEndpoint;

  getType(): string;
  setType(value: string): VPCEndpoint;

  getProject(): string;
  setProject(value: string): VPCEndpoint;

  getServiceName(): string;
  setServiceName(value: string): VPCEndpoint;

  getRouteTableIds(): string;
  setRouteTableIds(value: string): VPCEndpoint;

  getSubnetIds(): string;
  setSubnetIds(value: string): VPCEndpoint;

  getLabelsMap(): jspb.Map<string, string>;
  clearLabelsMap(): VPCEndpoint;

  getCreatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreatedAt(value?: google_protobuf_timestamp_pb.Timestamp): VPCEndpoint;
  hasCreatedAt(): boolean;
  clearCreatedAt(): VPCEndpoint;

  getUpdatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setUpdatedAt(value?: google_protobuf_timestamp_pb.Timestamp): VPCEndpoint;
  hasUpdatedAt(): boolean;
  clearUpdatedAt(): VPCEndpoint;

  getLastSyncTime(): string;
  setLastSyncTime(value: string): VPCEndpoint;

  getSelfLink(): string;
  setSelfLink(value: string): VPCEndpoint;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): VPCEndpoint.AsObject;
  static toObject(includeInstance: boolean, msg: VPCEndpoint): VPCEndpoint.AsObject;
  static serializeBinaryToWriter(message: VPCEndpoint, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): VPCEndpoint;
  static deserializeBinaryFromReader(message: VPCEndpoint, reader: jspb.BinaryReader): VPCEndpoint;
}

export namespace VPCEndpoint {
  export type AsObject = {
    id: string,
    name: string,
    provider: string,
    accountId: string,
    vpcId: string,
    region: string,
    state: string,
    type: string,
    project: string,
    serviceName: string,
    routeTableIds: string,
    subnetIds: string,
    labelsMap: Array<[string, string]>,
    createdAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    updatedAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    lastSyncTime: string,
    selfLink: string,
  }
}

export class PublicIP extends jspb.Message {
  getId(): string;
  setId(value: string): PublicIP;

  getType(): string;
  setType(value: string): PublicIP;

  getProvider(): string;
  setProvider(value: string): PublicIP;

  getAccountId(): string;
  setAccountId(value: string): PublicIP;

  getVpcId(): string;
  setVpcId(value: string): PublicIP;

  getRegion(): string;
  setRegion(value: string): PublicIP;

  getPublicIp(): string;
  setPublicIp(value: string): PublicIP;

  getInstanceId(): string;
  setInstanceId(value: string): PublicIP;

  getPrivateIp(): string;
  setPrivateIp(value: string): PublicIP;

  getByoip(): string;
  setByoip(value: string): PublicIP;

  getProject(): string;
  setProject(value: string): PublicIP;

  getLabelsMap(): jspb.Map<string, string>;
  clearLabelsMap(): PublicIP;

  getCreatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreatedAt(value?: google_protobuf_timestamp_pb.Timestamp): PublicIP;
  hasCreatedAt(): boolean;
  clearCreatedAt(): PublicIP;

  getUpdatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setUpdatedAt(value?: google_protobuf_timestamp_pb.Timestamp): PublicIP;
  hasUpdatedAt(): boolean;
  clearUpdatedAt(): PublicIP;

  getLastSyncTime(): string;
  setLastSyncTime(value: string): PublicIP;

  getSelfLink(): string;
  setSelfLink(value: string): PublicIP;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PublicIP.AsObject;
  static toObject(includeInstance: boolean, msg: PublicIP): PublicIP.AsObject;
  static serializeBinaryToWriter(message: PublicIP, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PublicIP;
  static deserializeBinaryFromReader(message: PublicIP, reader: jspb.BinaryReader): PublicIP;
}

export namespace PublicIP {
  export type AsObject = {
    id: string,
    type: string,
    provider: string,
    accountId: string,
    vpcId: string,
    region: string,
    publicIp: string,
    instanceId: string,
    privateIp: string,
    byoip: string,
    project: string,
    labelsMap: Array<[string, string]>,
    createdAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    updatedAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    lastSyncTime: string,
    selfLink: string,
  }
}

export class LB extends jspb.Message {
  getId(): string;
  setId(value: string): LB;

  getName(): string;
  setName(value: string): LB;

  getLoadBalancerType(): string;
  setLoadBalancerType(value: string): LB;

  getScheme(): string;
  setScheme(value: string): LB;

  getVpcId(): string;
  setVpcId(value: string): LB;

  getDnsName(): string;
  setDnsName(value: string): LB;

  getInstanceIdsList(): Array<string>;
  setInstanceIdsList(value: Array<string>): LB;
  clearInstanceIdsList(): LB;
  addInstanceIds(value: string, index?: number): LB;

  getTargetGroupIdsList(): Array<string>;
  setTargetGroupIdsList(value: Array<string>): LB;
  clearTargetGroupIdsList(): LB;
  addTargetGroupIds(value: string, index?: number): LB;

  getListenersList(): Array<LB.Listener>;
  setListenersList(value: Array<LB.Listener>): LB;
  clearListenersList(): LB;
  addListeners(value?: LB.Listener, index?: number): LB.Listener;

  getCrossZoneLoadBalancing(): boolean;
  setCrossZoneLoadBalancing(value: boolean): LB;

  getAccessLogsEnabled(): boolean;
  setAccessLogsEnabled(value: boolean): LB;

  getLoggingBucket(): string;
  setLoggingBucket(value: string): LB;

  getIpAddressesList(): Array<string>;
  setIpAddressesList(value: Array<string>): LB;
  clearIpAddressesList(): LB;
  addIpAddresses(value: string, index?: number): LB;

  getIpAddressType(): string;
  setIpAddressType(value: string): LB;

  getRegion(): string;
  setRegion(value: string): LB;

  getZone(): string;
  setZone(value: string): LB;

  getLabelsMap(): jspb.Map<string, string>;
  clearLabelsMap(): LB;

  getProject(): string;
  setProject(value: string): LB;

  getProvider(): string;
  setProvider(value: string): LB;

  getAccountId(): string;
  setAccountId(value: string): LB;

  getLastSyncTime(): string;
  setLastSyncTime(value: string): LB;

  getSelfLink(): string;
  setSelfLink(value: string): LB;

  getCreatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreatedAt(value?: google_protobuf_timestamp_pb.Timestamp): LB;
  hasCreatedAt(): boolean;
  clearCreatedAt(): LB;

  getUpdatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setUpdatedAt(value?: google_protobuf_timestamp_pb.Timestamp): LB;
  hasUpdatedAt(): boolean;
  clearUpdatedAt(): LB;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): LB.AsObject;
  static toObject(includeInstance: boolean, msg: LB): LB.AsObject;
  static serializeBinaryToWriter(message: LB, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): LB;
  static deserializeBinaryFromReader(message: LB, reader: jspb.BinaryReader): LB;
}

export namespace LB {
  export type AsObject = {
    id: string,
    name: string,
    loadBalancerType: string,
    scheme: string,
    vpcId: string,
    dnsName: string,
    instanceIdsList: Array<string>,
    targetGroupIdsList: Array<string>,
    listenersList: Array<LB.Listener.AsObject>,
    crossZoneLoadBalancing: boolean,
    accessLogsEnabled: boolean,
    loggingBucket: string,
    ipAddressesList: Array<string>,
    ipAddressType: string,
    region: string,
    zone: string,
    labelsMap: Array<[string, string]>,
    project: string,
    provider: string,
    accountId: string,
    lastSyncTime: string,
    selfLink: string,
    createdAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    updatedAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }

  export class Listener extends jspb.Message {
    getListenerId(): string;
    setListenerId(value: string): Listener;

    getProtocol(): string;
    setProtocol(value: string): Listener;

    getPort(): number;
    setPort(value: number): Listener;

    getTargetGroupId(): string;
    setTargetGroupId(value: string): Listener;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Listener.AsObject;
    static toObject(includeInstance: boolean, msg: Listener): Listener.AsObject;
    static serializeBinaryToWriter(message: Listener, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Listener;
    static deserializeBinaryFromReader(message: Listener, reader: jspb.BinaryReader): Listener;
  }

  export namespace Listener {
    export type AsObject = {
      listenerId: string,
      protocol: string,
      port: number,
      targetGroupId: string,
    }
  }

}

export class NetworkInterface extends jspb.Message {
  getId(): string;
  setId(value: string): NetworkInterface;

  getName(): string;
  setName(value: string): NetworkInterface;

  getSubnetId(): string;
  setSubnetId(value: string): NetworkInterface;

  getVpcId(): string;
  setVpcId(value: string): NetworkInterface;

  getInstanceId(): string;
  setInstanceId(value: string): NetworkInterface;

  getMacAddress(): string;
  setMacAddress(value: string): NetworkInterface;

  getPublicIp(): string;
  setPublicIp(value: string): NetworkInterface;

  getPrivateIpsList(): Array<string>;
  setPrivateIpsList(value: Array<string>): NetworkInterface;
  clearPrivateIpsList(): NetworkInterface;
  addPrivateIps(value: string, index?: number): NetworkInterface;

  getSecondaryPrivateIpsList(): Array<string>;
  setSecondaryPrivateIpsList(value: Array<string>): NetworkInterface;
  clearSecondaryPrivateIpsList(): NetworkInterface;
  addSecondaryPrivateIps(value: string, index?: number): NetworkInterface;

  getStatus(): string;
  setStatus(value: string): NetworkInterface;

  getAttachment(): NetworkInterface.Attachment | undefined;
  setAttachment(value?: NetworkInterface.Attachment): NetworkInterface;
  hasAttachment(): boolean;
  clearAttachment(): NetworkInterface;

  getDnsName(): string;
  setDnsName(value: string): NetworkInterface;

  getDnsServersList(): Array<string>;
  setDnsServersList(value: Array<string>): NetworkInterface;
  clearDnsServersList(): NetworkInterface;
  addDnsServers(value: string, index?: number): NetworkInterface;

  getSecurityGroupIdsList(): Array<string>;
  setSecurityGroupIdsList(value: Array<string>): NetworkInterface;
  clearSecurityGroupIdsList(): NetworkInterface;
  addSecurityGroupIds(value: string, index?: number): NetworkInterface;

  getIpForwarding(): boolean;
  setIpForwarding(value: boolean): NetworkInterface;

  getMtu(): number;
  setMtu(value: number): NetworkInterface;

  getBandwidth(): string;
  setBandwidth(value: string): NetworkInterface;

  getRegion(): string;
  setRegion(value: string): NetworkInterface;

  getZone(): string;
  setZone(value: string): NetworkInterface;

  getLabelsMap(): jspb.Map<string, string>;
  clearLabelsMap(): NetworkInterface;

  getProject(): string;
  setProject(value: string): NetworkInterface;

  getProvider(): string;
  setProvider(value: string): NetworkInterface;

  getAccountId(): string;
  setAccountId(value: string): NetworkInterface;

  getLastSyncTime(): string;
  setLastSyncTime(value: string): NetworkInterface;

  getSelfLink(): string;
  setSelfLink(value: string): NetworkInterface;

  getCreatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreatedAt(value?: google_protobuf_timestamp_pb.Timestamp): NetworkInterface;
  hasCreatedAt(): boolean;
  clearCreatedAt(): NetworkInterface;

  getUpdatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setUpdatedAt(value?: google_protobuf_timestamp_pb.Timestamp): NetworkInterface;
  hasUpdatedAt(): boolean;
  clearUpdatedAt(): NetworkInterface;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): NetworkInterface.AsObject;
  static toObject(includeInstance: boolean, msg: NetworkInterface): NetworkInterface.AsObject;
  static serializeBinaryToWriter(message: NetworkInterface, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): NetworkInterface;
  static deserializeBinaryFromReader(message: NetworkInterface, reader: jspb.BinaryReader): NetworkInterface;
}

export namespace NetworkInterface {
  export type AsObject = {
    id: string,
    name: string,
    subnetId: string,
    vpcId: string,
    instanceId: string,
    macAddress: string,
    publicIp: string,
    privateIpsList: Array<string>,
    secondaryPrivateIpsList: Array<string>,
    status: string,
    attachment?: NetworkInterface.Attachment.AsObject,
    dnsName: string,
    dnsServersList: Array<string>,
    securityGroupIdsList: Array<string>,
    ipForwarding: boolean,
    mtu: number,
    bandwidth: string,
    region: string,
    zone: string,
    labelsMap: Array<[string, string]>,
    project: string,
    provider: string,
    accountId: string,
    lastSyncTime: string,
    selfLink: string,
    createdAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    updatedAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }

  export class Attachment extends jspb.Message {
    getInstanceId(): string;
    setInstanceId(value: string): Attachment;

    getAttachmentId(): string;
    setAttachmentId(value: string): Attachment;

    getDeviceIndex(): string;
    setDeviceIndex(value: string): Attachment;

    getAttachTime(): google_protobuf_timestamp_pb.Timestamp | undefined;
    setAttachTime(value?: google_protobuf_timestamp_pb.Timestamp): Attachment;
    hasAttachTime(): boolean;
    clearAttachTime(): Attachment;

    getDeleteOnTermination(): boolean;
    setDeleteOnTermination(value: boolean): Attachment;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Attachment.AsObject;
    static toObject(includeInstance: boolean, msg: Attachment): Attachment.AsObject;
    static serializeBinaryToWriter(message: Attachment, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Attachment;
    static deserializeBinaryFromReader(message: Attachment, reader: jspb.BinaryReader): Attachment;
  }

  export namespace Attachment {
    export type AsObject = {
      instanceId: string,
      attachmentId: string,
      deviceIndex: string,
      attachTime?: google_protobuf_timestamp_pb.Timestamp.AsObject,
      deleteOnTermination: boolean,
    }
  }

}

export class VPNConcentrator extends jspb.Message {
  getId(): string;
  setId(value: string): VPNConcentrator;

  getName(): string;
  setName(value: string): VPNConcentrator;

  getProvider(): string;
  setProvider(value: string): VPNConcentrator;

  getAccountId(): string;
  setAccountId(value: string): VPNConcentrator;

  getRegion(): string;
  setRegion(value: string): VPNConcentrator;

  getState(): string;
  setState(value: string): VPNConcentrator;

  getType(): string;
  setType(value: string): VPNConcentrator;

  getVpcId(): string;
  setVpcId(value: string): VPNConcentrator;

  getAsn(): number;
  setAsn(value: number): VPNConcentrator;

  getProject(): string;
  setProject(value: string): VPNConcentrator;

  getLastSyncTime(): string;
  setLastSyncTime(value: string): VPNConcentrator;

  getSelfLink(): string;
  setSelfLink(value: string): VPNConcentrator;

  getLabelsMap(): jspb.Map<string, string>;
  clearLabelsMap(): VPNConcentrator;

  getCreatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreatedAt(value?: google_protobuf_timestamp_pb.Timestamp): VPNConcentrator;
  hasCreatedAt(): boolean;
  clearCreatedAt(): VPNConcentrator;

  getUpdatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setUpdatedAt(value?: google_protobuf_timestamp_pb.Timestamp): VPNConcentrator;
  hasUpdatedAt(): boolean;
  clearUpdatedAt(): VPNConcentrator;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): VPNConcentrator.AsObject;
  static toObject(includeInstance: boolean, msg: VPNConcentrator): VPNConcentrator.AsObject;
  static serializeBinaryToWriter(message: VPNConcentrator, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): VPNConcentrator;
  static deserializeBinaryFromReader(message: VPNConcentrator, reader: jspb.BinaryReader): VPNConcentrator;
}

export namespace VPNConcentrator {
  export type AsObject = {
    id: string,
    name: string,
    provider: string,
    accountId: string,
    region: string,
    state: string,
    type: string,
    vpcId: string,
    asn: number,
    project: string,
    lastSyncTime: string,
    selfLink: string,
    labelsMap: Array<[string, string]>,
    createdAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    updatedAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }
}

export class KeyPair extends jspb.Message {
  getId(): string;
  setId(value: string): KeyPair;

  getName(): string;
  setName(value: string): KeyPair;

  getKeyPairType(): string;
  setKeyPairType(value: string): KeyPair;

  getPublicKey(): string;
  setPublicKey(value: string): KeyPair;

  getPrivateKeyFingerprint(): string;
  setPrivateKeyFingerprint(value: string): KeyPair;

  getInstanceIdsList(): Array<string>;
  setInstanceIdsList(value: Array<string>): KeyPair;
  clearInstanceIdsList(): KeyPair;
  addInstanceIds(value: string, index?: number): KeyPair;

  getRegion(): string;
  setRegion(value: string): KeyPair;

  getProject(): string;
  setProject(value: string): KeyPair;

  getProvider(): string;
  setProvider(value: string): KeyPair;

  getAccountId(): string;
  setAccountId(value: string): KeyPair;

  getLabelsMap(): jspb.Map<string, string>;
  clearLabelsMap(): KeyPair;

  getLastSyncTime(): string;
  setLastSyncTime(value: string): KeyPair;

  getSelfLink(): string;
  setSelfLink(value: string): KeyPair;

  getCreatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreatedAt(value?: google_protobuf_timestamp_pb.Timestamp): KeyPair;
  hasCreatedAt(): boolean;
  clearCreatedAt(): KeyPair;

  getUpdatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setUpdatedAt(value?: google_protobuf_timestamp_pb.Timestamp): KeyPair;
  hasUpdatedAt(): boolean;
  clearUpdatedAt(): KeyPair;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): KeyPair.AsObject;
  static toObject(includeInstance: boolean, msg: KeyPair): KeyPair.AsObject;
  static serializeBinaryToWriter(message: KeyPair, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): KeyPair;
  static deserializeBinaryFromReader(message: KeyPair, reader: jspb.BinaryReader): KeyPair;
}

export namespace KeyPair {
  export type AsObject = {
    id: string,
    name: string,
    keyPairType: string,
    publicKey: string,
    privateKeyFingerprint: string,
    instanceIdsList: Array<string>,
    region: string,
    project: string,
    provider: string,
    accountId: string,
    labelsMap: Array<[string, string]>,
    lastSyncTime: string,
    selfLink: string,
    createdAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    updatedAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }
}

export class VpcGraphNode extends jspb.Message {
  getId(): string;
  setId(value: string): VpcGraphNode;

  getResourceType(): string;
  setResourceType(value: string): VpcGraphNode;

  getName(): string;
  setName(value: string): VpcGraphNode;

  getPropertiesMap(): jspb.Map<string, string>;
  clearPropertiesMap(): VpcGraphNode;

  getProvider(): string;
  setProvider(value: string): VpcGraphNode;

  getAccountId(): string;
  setAccountId(value: string): VpcGraphNode;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): VpcGraphNode.AsObject;
  static toObject(includeInstance: boolean, msg: VpcGraphNode): VpcGraphNode.AsObject;
  static serializeBinaryToWriter(message: VpcGraphNode, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): VpcGraphNode;
  static deserializeBinaryFromReader(message: VpcGraphNode, reader: jspb.BinaryReader): VpcGraphNode;
}

export namespace VpcGraphNode {
  export type AsObject = {
    id: string,
    resourceType: string,
    name: string,
    propertiesMap: Array<[string, string]>,
    provider: string,
    accountId: string,
  }
}

export class VpcGraphEdge extends jspb.Message {
  getSourceNodeId(): string;
  setSourceNodeId(value: string): VpcGraphEdge;

  getTargetNodeId(): string;
  setTargetNodeId(value: string): VpcGraphEdge;

  getRelationshipType(): string;
  setRelationshipType(value: string): VpcGraphEdge;

  getProvider(): string;
  setProvider(value: string): VpcGraphEdge;

  getAccountId(): string;
  setAccountId(value: string): VpcGraphEdge;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): VpcGraphEdge.AsObject;
  static toObject(includeInstance: boolean, msg: VpcGraphEdge): VpcGraphEdge.AsObject;
  static serializeBinaryToWriter(message: VpcGraphEdge, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): VpcGraphEdge;
  static deserializeBinaryFromReader(message: VpcGraphEdge, reader: jspb.BinaryReader): VpcGraphEdge;
}

export namespace VpcGraphEdge {
  export type AsObject = {
    sourceNodeId: string,
    targetNodeId: string,
    relationshipType: string,
    provider: string,
    accountId: string,
  }
}

export enum LoadBalancerType { 
  ALB = 0,
  NLB = 1,
  ELB = 2,
  GLB = 3,
  UNKNOWN = 4,
}
