import * as jspb from 'google-protobuf'

import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb';


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

