import * as jspb from 'google-protobuf'

import * as types_pb from './types_pb';


export class ListNamespacesRequest extends jspb.Message {
  getClusterName(): string;
  setClusterName(value: string): ListNamespacesRequest;

  getLabelsMap(): jspb.Map<string, string>;
  clearLabelsMap(): ListNamespacesRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListNamespacesRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListNamespacesRequest): ListNamespacesRequest.AsObject;
  static serializeBinaryToWriter(message: ListNamespacesRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListNamespacesRequest;
  static deserializeBinaryFromReader(message: ListNamespacesRequest, reader: jspb.BinaryReader): ListNamespacesRequest;
}

export namespace ListNamespacesRequest {
  export type AsObject = {
    clusterName: string,
    labelsMap: Array<[string, string]>,
  }
}

export class ListNamespacesResponse extends jspb.Message {
  getNamespacesList(): Array<types_pb.Namespace>;
  setNamespacesList(value: Array<types_pb.Namespace>): ListNamespacesResponse;
  clearNamespacesList(): ListNamespacesResponse;
  addNamespaces(value?: types_pb.Namespace, index?: number): types_pb.Namespace;

  getLastSyncTime(): string;
  setLastSyncTime(value: string): ListNamespacesResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListNamespacesResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListNamespacesResponse): ListNamespacesResponse.AsObject;
  static serializeBinaryToWriter(message: ListNamespacesResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListNamespacesResponse;
  static deserializeBinaryFromReader(message: ListNamespacesResponse, reader: jspb.BinaryReader): ListNamespacesResponse;
}

export namespace ListNamespacesResponse {
  export type AsObject = {
    namespacesList: Array<types_pb.Namespace.AsObject>,
    lastSyncTime: string,
  }
}

export class ListNodesRequest extends jspb.Message {
  getClusterName(): string;
  setClusterName(value: string): ListNodesRequest;

  getLabelsMap(): jspb.Map<string, string>;
  clearLabelsMap(): ListNodesRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListNodesRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListNodesRequest): ListNodesRequest.AsObject;
  static serializeBinaryToWriter(message: ListNodesRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListNodesRequest;
  static deserializeBinaryFromReader(message: ListNodesRequest, reader: jspb.BinaryReader): ListNodesRequest;
}

export namespace ListNodesRequest {
  export type AsObject = {
    clusterName: string,
    labelsMap: Array<[string, string]>,
  }
}

export class ListNodesResponse extends jspb.Message {
  getNodesList(): Array<types_pb.Node>;
  setNodesList(value: Array<types_pb.Node>): ListNodesResponse;
  clearNodesList(): ListNodesResponse;
  addNodes(value?: types_pb.Node, index?: number): types_pb.Node;

  getLastSyncTime(): string;
  setLastSyncTime(value: string): ListNodesResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListNodesResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListNodesResponse): ListNodesResponse.AsObject;
  static serializeBinaryToWriter(message: ListNodesResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListNodesResponse;
  static deserializeBinaryFromReader(message: ListNodesResponse, reader: jspb.BinaryReader): ListNodesResponse;
}

export namespace ListNodesResponse {
  export type AsObject = {
    nodesList: Array<types_pb.Node.AsObject>,
    lastSyncTime: string,
  }
}

export class ListPodsRequest extends jspb.Message {
  getClusterName(): string;
  setClusterName(value: string): ListPodsRequest;

  getLabelsMap(): jspb.Map<string, string>;
  clearLabelsMap(): ListPodsRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListPodsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListPodsRequest): ListPodsRequest.AsObject;
  static serializeBinaryToWriter(message: ListPodsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListPodsRequest;
  static deserializeBinaryFromReader(message: ListPodsRequest, reader: jspb.BinaryReader): ListPodsRequest;
}

export namespace ListPodsRequest {
  export type AsObject = {
    clusterName: string,
    labelsMap: Array<[string, string]>,
  }
}

export class ListPodsResponse extends jspb.Message {
  getPodsList(): Array<types_pb.Pod>;
  setPodsList(value: Array<types_pb.Pod>): ListPodsResponse;
  clearPodsList(): ListPodsResponse;
  addPods(value?: types_pb.Pod, index?: number): types_pb.Pod;

  getLastSyncTime(): string;
  setLastSyncTime(value: string): ListPodsResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListPodsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListPodsResponse): ListPodsResponse.AsObject;
  static serializeBinaryToWriter(message: ListPodsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListPodsResponse;
  static deserializeBinaryFromReader(message: ListPodsResponse, reader: jspb.BinaryReader): ListPodsResponse;
}

export namespace ListPodsResponse {
  export type AsObject = {
    podsList: Array<types_pb.Pod.AsObject>,
    lastSyncTime: string,
  }
}

export class ListServicesRequest extends jspb.Message {
  getClusterName(): string;
  setClusterName(value: string): ListServicesRequest;

  getLabelsMap(): jspb.Map<string, string>;
  clearLabelsMap(): ListServicesRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListServicesRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListServicesRequest): ListServicesRequest.AsObject;
  static serializeBinaryToWriter(message: ListServicesRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListServicesRequest;
  static deserializeBinaryFromReader(message: ListServicesRequest, reader: jspb.BinaryReader): ListServicesRequest;
}

export namespace ListServicesRequest {
  export type AsObject = {
    clusterName: string,
    labelsMap: Array<[string, string]>,
  }
}

export class ListServicesResponse extends jspb.Message {
  getServicesList(): Array<types_pb.K8sService>;
  setServicesList(value: Array<types_pb.K8sService>): ListServicesResponse;
  clearServicesList(): ListServicesResponse;
  addServices(value?: types_pb.K8sService, index?: number): types_pb.K8sService;

  getLastSyncTime(): string;
  setLastSyncTime(value: string): ListServicesResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListServicesResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListServicesResponse): ListServicesResponse.AsObject;
  static serializeBinaryToWriter(message: ListServicesResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListServicesResponse;
  static deserializeBinaryFromReader(message: ListServicesResponse, reader: jspb.BinaryReader): ListServicesResponse;
}

export namespace ListServicesResponse {
  export type AsObject = {
    servicesList: Array<types_pb.K8sService.AsObject>,
    lastSyncTime: string,
  }
}

export class ListClustersRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListClustersRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListClustersRequest): ListClustersRequest.AsObject;
  static serializeBinaryToWriter(message: ListClustersRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListClustersRequest;
  static deserializeBinaryFromReader(message: ListClustersRequest, reader: jspb.BinaryReader): ListClustersRequest;
}

export namespace ListClustersRequest {
  export type AsObject = {
  }
}

export class ListClustersResponse extends jspb.Message {
  getClustersList(): Array<types_pb.Cluster>;
  setClustersList(value: Array<types_pb.Cluster>): ListClustersResponse;
  clearClustersList(): ListClustersResponse;
  addClusters(value?: types_pb.Cluster, index?: number): types_pb.Cluster;

  getLastSyncTime(): string;
  setLastSyncTime(value: string): ListClustersResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListClustersResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListClustersResponse): ListClustersResponse.AsObject;
  static serializeBinaryToWriter(message: ListClustersResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListClustersResponse;
  static deserializeBinaryFromReader(message: ListClustersResponse, reader: jspb.BinaryReader): ListClustersResponse;
}

export namespace ListClustersResponse {
  export type AsObject = {
    clustersList: Array<types_pb.Cluster.AsObject>,
    lastSyncTime: string,
  }
}

export class ListPodsCIDRsRequest extends jspb.Message {
  getClusterName(): string;
  setClusterName(value: string): ListPodsCIDRsRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListPodsCIDRsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListPodsCIDRsRequest): ListPodsCIDRsRequest.AsObject;
  static serializeBinaryToWriter(message: ListPodsCIDRsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListPodsCIDRsRequest;
  static deserializeBinaryFromReader(message: ListPodsCIDRsRequest, reader: jspb.BinaryReader): ListPodsCIDRsRequest;
}

export namespace ListPodsCIDRsRequest {
  export type AsObject = {
    clusterName: string,
  }
}

export class ListPodsCIDRsResponse extends jspb.Message {
  getCidrsList(): Array<string>;
  setCidrsList(value: Array<string>): ListPodsCIDRsResponse;
  clearCidrsList(): ListPodsCIDRsResponse;
  addCidrs(value: string, index?: number): ListPodsCIDRsResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListPodsCIDRsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListPodsCIDRsResponse): ListPodsCIDRsResponse.AsObject;
  static serializeBinaryToWriter(message: ListPodsCIDRsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListPodsCIDRsResponse;
  static deserializeBinaryFromReader(message: ListPodsCIDRsResponse, reader: jspb.BinaryReader): ListPodsCIDRsResponse;
}

export namespace ListPodsCIDRsResponse {
  export type AsObject = {
    cidrsList: Array<string>,
  }
}

export class ListServicesCIDRsRequest extends jspb.Message {
  getClusterName(): string;
  setClusterName(value: string): ListServicesCIDRsRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListServicesCIDRsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListServicesCIDRsRequest): ListServicesCIDRsRequest.AsObject;
  static serializeBinaryToWriter(message: ListServicesCIDRsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListServicesCIDRsRequest;
  static deserializeBinaryFromReader(message: ListServicesCIDRsRequest, reader: jspb.BinaryReader): ListServicesCIDRsRequest;
}

export namespace ListServicesCIDRsRequest {
  export type AsObject = {
    clusterName: string,
  }
}

export class ListServicesCIDRsResponse extends jspb.Message {
  getCidr(): string;
  setCidr(value: string): ListServicesCIDRsResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListServicesCIDRsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListServicesCIDRsResponse): ListServicesCIDRsResponse.AsObject;
  static serializeBinaryToWriter(message: ListServicesCIDRsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListServicesCIDRsResponse;
  static deserializeBinaryFromReader(message: ListServicesCIDRsResponse, reader: jspb.BinaryReader): ListServicesCIDRsResponse;
}

export namespace ListServicesCIDRsResponse {
  export type AsObject = {
    cidr: string,
  }
}

