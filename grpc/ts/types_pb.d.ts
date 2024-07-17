import * as jspb from 'google-protobuf'

import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb';


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

export class Cluster extends jspb.Message {
  getName(): string;
  setName(value: string): Cluster;

  getFullName(): string;
  setFullName(value: string): Cluster;

  getArn(): string;
  setArn(value: string): Cluster;

  getVpcId(): string;
  setVpcId(value: string): Cluster;

  getRegion(): string;
  setRegion(value: string): Cluster;

  getProject(): string;
  setProject(value: string): Cluster;

  getLabelsMap(): jspb.Map<string, string>;
  clearLabelsMap(): Cluster;

  getProvider(): string;
  setProvider(value: string): Cluster;

  getAccountId(): string;
  setAccountId(value: string): Cluster;

  getId(): string;
  setId(value: string): Cluster;

  getLastSyncTime(): string;
  setLastSyncTime(value: string): Cluster;

  getSelfLink(): string;
  setSelfLink(value: string): Cluster;

  getCreatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreatedAt(value?: google_protobuf_timestamp_pb.Timestamp): Cluster;
  hasCreatedAt(): boolean;
  clearCreatedAt(): Cluster;

  getUpdatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setUpdatedAt(value?: google_protobuf_timestamp_pb.Timestamp): Cluster;
  hasUpdatedAt(): boolean;
  clearUpdatedAt(): Cluster;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Cluster.AsObject;
  static toObject(includeInstance: boolean, msg: Cluster): Cluster.AsObject;
  static serializeBinaryToWriter(message: Cluster, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Cluster;
  static deserializeBinaryFromReader(message: Cluster, reader: jspb.BinaryReader): Cluster;
}

export namespace Cluster {
  export type AsObject = {
    name: string,
    fullName: string,
    arn: string,
    vpcId: string,
    region: string,
    project: string,
    labelsMap: Array<[string, string]>,
    provider: string,
    accountId: string,
    id: string,
    lastSyncTime: string,
    selfLink: string,
    createdAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    updatedAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }
}

export class Node extends jspb.Message {
  getCluster(): string;
  setCluster(value: string): Node;

  getName(): string;
  setName(value: string): Node;

  getNamespace(): string;
  setNamespace(value: string): Node;

  getAddressesList(): Array<string>;
  setAddressesList(value: Array<string>): Node;
  clearAddressesList(): Node;
  addAddresses(value: string, index?: number): Node;

  getProject(): string;
  setProject(value: string): Node;

  getLastSyncTime(): string;
  setLastSyncTime(value: string): Node;

  getSelfLink(): string;
  setSelfLink(value: string): Node;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Node.AsObject;
  static toObject(includeInstance: boolean, msg: Node): Node.AsObject;
  static serializeBinaryToWriter(message: Node, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Node;
  static deserializeBinaryFromReader(message: Node, reader: jspb.BinaryReader): Node;
}

export namespace Node {
  export type AsObject = {
    cluster: string,
    name: string,
    namespace: string,
    addressesList: Array<string>,
    project: string,
    lastSyncTime: string,
    selfLink: string,
  }
}

export class Namespace extends jspb.Message {
  getCluster(): string;
  setCluster(value: string): Namespace;

  getName(): string;
  setName(value: string): Namespace;

  getLabelsMap(): jspb.Map<string, string>;
  clearLabelsMap(): Namespace;

  getProject(): string;
  setProject(value: string): Namespace;

  getLastSyncTime(): string;
  setLastSyncTime(value: string): Namespace;

  getSelfLink(): string;
  setSelfLink(value: string): Namespace;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Namespace.AsObject;
  static toObject(includeInstance: boolean, msg: Namespace): Namespace.AsObject;
  static serializeBinaryToWriter(message: Namespace, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Namespace;
  static deserializeBinaryFromReader(message: Namespace, reader: jspb.BinaryReader): Namespace;
}

export namespace Namespace {
  export type AsObject = {
    cluster: string,
    name: string,
    labelsMap: Array<[string, string]>,
    project: string,
    lastSyncTime: string,
    selfLink: string,
  }
}

export class Pod extends jspb.Message {
  getCluster(): string;
  setCluster(value: string): Pod;

  getNamespace(): string;
  setNamespace(value: string): Pod;

  getName(): string;
  setName(value: string): Pod;

  getIp(): string;
  setIp(value: string): Pod;

  getLabelsMap(): jspb.Map<string, string>;
  clearLabelsMap(): Pod;

  getState(): string;
  setState(value: string): Pod;

  getProject(): string;
  setProject(value: string): Pod;

  getLastSyncTime(): string;
  setLastSyncTime(value: string): Pod;

  getSelfLink(): string;
  setSelfLink(value: string): Pod;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Pod.AsObject;
  static toObject(includeInstance: boolean, msg: Pod): Pod.AsObject;
  static serializeBinaryToWriter(message: Pod, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Pod;
  static deserializeBinaryFromReader(message: Pod, reader: jspb.BinaryReader): Pod;
}

export namespace Pod {
  export type AsObject = {
    cluster: string,
    namespace: string,
    name: string,
    ip: string,
    labelsMap: Array<[string, string]>,
    state: string,
    project: string,
    lastSyncTime: string,
    selfLink: string,
  }
}

export class K8sService extends jspb.Message {
  getCluster(): string;
  setCluster(value: string): K8sService;

  getNamespace(): string;
  setNamespace(value: string): K8sService;

  getName(): string;
  setName(value: string): K8sService;

  getIngressesList(): Array<K8sService.Ingress>;
  setIngressesList(value: Array<K8sService.Ingress>): K8sService;
  clearIngressesList(): K8sService;
  addIngresses(value?: K8sService.Ingress, index?: number): K8sService.Ingress;

  getLabelsMap(): jspb.Map<string, string>;
  clearLabelsMap(): K8sService;

  getType(): string;
  setType(value: string): K8sService;

  getProject(): string;
  setProject(value: string): K8sService;

  getLastSyncTime(): string;
  setLastSyncTime(value: string): K8sService;

  getSelfLink(): string;
  setSelfLink(value: string): K8sService;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): K8sService.AsObject;
  static toObject(includeInstance: boolean, msg: K8sService): K8sService.AsObject;
  static serializeBinaryToWriter(message: K8sService, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): K8sService;
  static deserializeBinaryFromReader(message: K8sService, reader: jspb.BinaryReader): K8sService;
}

export namespace K8sService {
  export type AsObject = {
    cluster: string,
    namespace: string,
    name: string,
    ingressesList: Array<K8sService.Ingress.AsObject>,
    labelsMap: Array<[string, string]>,
    type: string,
    project: string,
    lastSyncTime: string,
    selfLink: string,
  }

  export class Ingress extends jspb.Message {
    getHostname(): string;
    setHostname(value: string): Ingress;

    getIp(): string;
    setIp(value: string): Ingress;

    getPortsList(): Array<string>;
    setPortsList(value: Array<string>): Ingress;
    clearPortsList(): Ingress;
    addPorts(value: string, index?: number): Ingress;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Ingress.AsObject;
    static toObject(includeInstance: boolean, msg: Ingress): Ingress.AsObject;
    static serializeBinaryToWriter(message: Ingress, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Ingress;
    static deserializeBinaryFromReader(message: Ingress, reader: jspb.BinaryReader): Ingress;
  }

  export namespace Ingress {
    export type AsObject = {
      hostname: string,
      ip: string,
      portsList: Array<string>,
    }
  }

}

