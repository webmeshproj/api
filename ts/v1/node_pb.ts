//
//Copyright 2023 Avi Zimmerman <avi.zimmerman@gmail.com>
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// @generated by protoc-gen-es v1.4.2 with parameter "target=ts"
// @generated from file v1/node.proto (package v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, protoInt64, Timestamp } from "@bufbuild/protobuf";

/**
 * ClusterStatus is the status of the node in the cluster.
 *
 * @generated from enum v1.ClusterStatus
 */
export enum ClusterStatus {
  /**
   * CLUSTER_STATUS_UNKNOWN is the default status.
   *
   * @generated from enum value: CLUSTER_STATUS_UNKNOWN = 0;
   */
  CLUSTER_STATUS_UNKNOWN = 0,

  /**
   * CLUSTER_LEADER is the status for the leader node.
   *
   * @generated from enum value: CLUSTER_LEADER = 1;
   */
  CLUSTER_LEADER = 1,

  /**
   * CLUSTER_VOTER is the status for a voter node.
   *
   * @generated from enum value: CLUSTER_VOTER = 2;
   */
  CLUSTER_VOTER = 2,

  /**
   * CLUSTER_OBSERVER is the status for a non-voter node.
   *
   * @generated from enum value: CLUSTER_OBSERVER = 3;
   */
  CLUSTER_OBSERVER = 3,

  /**
   * CLUSTER_NODE is the status of a node that is not a part of the storage consensus.
   *
   * @generated from enum value: CLUSTER_NODE = 4;
   */
  CLUSTER_NODE = 4,
}
// Retrieve enum metadata with: proto3.getEnumType(ClusterStatus)
proto3.util.setEnumType(ClusterStatus, "v1.ClusterStatus", [
  { no: 0, name: "CLUSTER_STATUS_UNKNOWN" },
  { no: 1, name: "CLUSTER_LEADER" },
  { no: 2, name: "CLUSTER_VOTER" },
  { no: 3, name: "CLUSTER_OBSERVER" },
  { no: 4, name: "CLUSTER_NODE" },
]);

/**
 * Feature is a list of features supported by a node.
 *
 * @generated from enum v1.Feature
 */
export enum Feature {
  /**
   * FEATURE_NONE is the default feature set.
   *
   * @generated from enum value: FEATURE_NONE = 0;
   */
  FEATURE_NONE = 0,

  /**
   * NODES is the feature for nodes. This is always supported.
   *
   * @generated from enum value: NODES = 1;
   */
  NODES = 1,

  /**
   * LEADER_PROXY is the feature for leader proxying.
   *
   * @generated from enum value: LEADER_PROXY = 2;
   */
  LEADER_PROXY = 2,

  /**
   * MESH_API is the feature for the mesh API.
   * This will be deprecated in favor of the MEMBERSHIP feature.
   *
   * @generated from enum value: MESH_API = 3;
   */
  MESH_API = 3,

  /**
   * ADMIN_API is the feature for the admin API.
   *
   * @generated from enum value: ADMIN_API = 4;
   */
  ADMIN_API = 4,

  /**
   * MEMBERSHIP is the feature for membership. This is always supported on storage-providing members.
   *
   * @generated from enum value: MEMBERSHIP = 5;
   */
  MEMBERSHIP = 5,

  /**
   * METRICS is the feature for exposing metrics.
   *
   * @generated from enum value: METRICS = 6;
   */
  METRICS = 6,

  /**
   * ICE_NEGOTIATION is the feature for ICE negotiation.
   *
   * @generated from enum value: ICE_NEGOTIATION = 7;
   */
  ICE_NEGOTIATION = 7,

  /**
   * TURN_SERVER is the feature for TURN server.
   *
   * @generated from enum value: TURN_SERVER = 8;
   */
  TURN_SERVER = 8,

  /**
   * MESH_DNS is the feature for mesh DNS.
   *
   * @generated from enum value: MESH_DNS = 9;
   */
  MESH_DNS = 9,

  /**
   * FORWARD_MESH_DNS is the feature for forwarding mesh DNS lookups to other meshes.
   *
   * @generated from enum value: FORWARD_MESH_DNS = 10;
   */
  FORWARD_MESH_DNS = 10,

  /**
   * STORAGE_QUERIER is the feature for querying, publishing, and subscribing to mesh state.
   *
   * @generated from enum value: STORAGE_QUERIER = 11;
   */
  STORAGE_QUERIER = 11,

  /**
   * STORAGE_PROVIDER is the feature for being able to provide distributed storage.
   *
   * @generated from enum value: STORAGE_PROVIDER = 12;
   */
  STORAGE_PROVIDER = 12,

  /**
   * REGISTRAR is the feature for being able to register aliases to node IDs and/or public keys.
   *
   * @generated from enum value: REGISTRAR = 13;
   */
  REGISTRAR = 13,
}
// Retrieve enum metadata with: proto3.getEnumType(Feature)
proto3.util.setEnumType(Feature, "v1.Feature", [
  { no: 0, name: "FEATURE_NONE" },
  { no: 1, name: "NODES" },
  { no: 2, name: "LEADER_PROXY" },
  { no: 3, name: "MESH_API" },
  { no: 4, name: "ADMIN_API" },
  { no: 5, name: "MEMBERSHIP" },
  { no: 6, name: "METRICS" },
  { no: 7, name: "ICE_NEGOTIATION" },
  { no: 8, name: "TURN_SERVER" },
  { no: 9, name: "MESH_DNS" },
  { no: 10, name: "FORWARD_MESH_DNS" },
  { no: 11, name: "STORAGE_QUERIER" },
  { no: 12, name: "STORAGE_PROVIDER" },
  { no: 13, name: "REGISTRAR" },
]);

/**
 * EdgeAttribute are pre-defined edge attributes. They should
 * be used as their string values.
 *
 * @generated from enum v1.EdgeAttribute
 */
export enum EdgeAttribute {
  /**
   * EDGE_ATTRIBUTE_UNKNOWN is an unknown edge attribute.
   *
   * @generated from enum value: EDGE_ATTRIBUTE_UNKNOWN = 0;
   */
  UNKNOWN = 0,

  /**
   * EDGE_ATTRIBUTE_NATIVE is a native edge attribute.
   *
   * @generated from enum value: EDGE_ATTRIBUTE_NATIVE = 1;
   */
  NATIVE = 1,

  /**
   * EDGE_ATTRIBUTE_ICE is an ICE edge attribute.
   *
   * @generated from enum value: EDGE_ATTRIBUTE_ICE = 2;
   */
  ICE = 2,

  /**
   * EDGE_ATTRIBUTE_LIBP2P is a libp2p edge attribute.
   *
   * @generated from enum value: EDGE_ATTRIBUTE_LIBP2P = 3;
   */
  LIBP2P = 3,
}
// Retrieve enum metadata with: proto3.getEnumType(EdgeAttribute)
proto3.util.setEnumType(EdgeAttribute, "v1.EdgeAttribute", [
  { no: 0, name: "EDGE_ATTRIBUTE_UNKNOWN" },
  { no: 1, name: "EDGE_ATTRIBUTE_NATIVE" },
  { no: 2, name: "EDGE_ATTRIBUTE_ICE" },
  { no: 3, name: "EDGE_ATTRIBUTE_LIBP2P" },
]);

/**
 * DataChannel are the data channels used when communicating over ICE
 * with a node.
 *
 * @generated from enum v1.DataChannel
 */
export enum DataChannel {
  /**
   * CHANNELS is the data channel used for negotiating new channels.
   * This is the first channel that is opened. The ID of the channel
   * should be 0.
   *
   * @generated from enum value: CHANNELS = 0;
   */
  CHANNELS = 0,

  /**
   * CONNECTIONS is the data channel used for negotiating new connections.
   * This is a channel that is opened for each incoming connection from a
   * client. The ID should start at 0 and be incremented for each new connection.
   *
   * @generated from enum value: CONNECTIONS = 1;
   */
  CONNECTIONS = 1,
}
// Retrieve enum metadata with: proto3.getEnumType(DataChannel)
proto3.util.setEnumType(DataChannel, "v1.DataChannel", [
  { no: 0, name: "CHANNELS" },
  { no: 1, name: "CONNECTIONS" },
]);

/**
 * FeaturePort describes a feature and the port it is advertised on.
 *
 * @generated from message v1.FeaturePort
 */
export class FeaturePort extends Message<FeaturePort> {
  /**
   * Feature is the feature advertised on the port.
   *
   * @generated from field: v1.Feature feature = 1;
   */
  feature = Feature.FEATURE_NONE;

  /**
   * Port is the port the feature is advertised on.
   *
   * @generated from field: int32 port = 2;
   */
  port = 0;

  constructor(data?: PartialMessage<FeaturePort>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "v1.FeaturePort";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "feature", kind: "enum", T: proto3.getEnumType(Feature) },
    { no: 2, name: "port", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): FeaturePort {
    return new FeaturePort().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): FeaturePort {
    return new FeaturePort().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): FeaturePort {
    return new FeaturePort().fromJsonString(jsonString, options);
  }

  static equals(a: FeaturePort | PlainMessage<FeaturePort> | undefined, b: FeaturePort | PlainMessage<FeaturePort> | undefined): boolean {
    return proto3.util.equals(FeaturePort, a, b);
  }
}

/**
 * MeshNode is a node that has been registered with the mesh.
 *
 * @generated from message v1.MeshNode
 */
export class MeshNode extends Message<MeshNode> {
  /**
   * ID is the ID of the node.
   *
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * PublicKey is the public key of the node.
   *
   * @generated from field: string publicKey = 2;
   */
  publicKey = "";

  /**
   * PrimaryEndpoint is the primary endpoint of the node.
   *
   * @generated from field: string primaryEndpoint = 4;
   */
  primaryEndpoint = "";

  /**
   * WireguardEndpoints is a list of WireGuard endpoints for the node.
   *
   * @generated from field: repeated string wireguardEndpoints = 5;
   */
  wireguardEndpoints: string[] = [];

  /**
   * ZoneAwarenessID is the zone awareness ID of the node.
   *
   * @generated from field: string zoneAwarenessID = 6;
   */
  zoneAwarenessID = "";

  /**
   * PrivateIPv4 is the private IPv4 address of the node.
   *
   * @generated from field: string privateIPv4 = 7;
   */
  privateIPv4 = "";

  /**
   * PrivateIPv6 is the private IPv6 address of the node.
   *
   * @generated from field: string privateIPv6 = 8;
   */
  privateIPv6 = "";

  /**
   * Features are a list of features and the ports they are advertised on.
   *
   * @generated from field: repeated v1.FeaturePort features = 9;
   */
  features: FeaturePort[] = [];

  /**
   * Multiaddrs are the multiaddrs of the node.
   *
   * @generated from field: repeated string multiaddrs = 10;
   */
  multiaddrs: string[] = [];

  /**
   * JoinedAt is the time the node joined the cluster.
   *
   * @generated from field: google.protobuf.Timestamp joinedAt = 11;
   */
  joinedAt?: Timestamp;

  constructor(data?: PartialMessage<MeshNode>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "v1.MeshNode";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "publicKey", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "primaryEndpoint", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "wireguardEndpoints", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 6, name: "zoneAwarenessID", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 7, name: "privateIPv4", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 8, name: "privateIPv6", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 9, name: "features", kind: "message", T: FeaturePort, repeated: true },
    { no: 10, name: "multiaddrs", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 11, name: "joinedAt", kind: "message", T: Timestamp },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MeshNode {
    return new MeshNode().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MeshNode {
    return new MeshNode().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MeshNode {
    return new MeshNode().fromJsonString(jsonString, options);
  }

  static equals(a: MeshNode | PlainMessage<MeshNode> | undefined, b: MeshNode | PlainMessage<MeshNode> | undefined): boolean {
    return proto3.util.equals(MeshNode, a, b);
  }
}

/**
 * NodeList is a list of nodes.
 *
 * @generated from message v1.NodeList
 */
export class NodeList extends Message<NodeList> {
  /**
   * Nodes is the list of nodes.
   *
   * @generated from field: repeated v1.MeshNode nodes = 1;
   */
  nodes: MeshNode[] = [];

  constructor(data?: PartialMessage<NodeList>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "v1.NodeList";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "nodes", kind: "message", T: MeshNode, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): NodeList {
    return new NodeList().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): NodeList {
    return new NodeList().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): NodeList {
    return new NodeList().fromJsonString(jsonString, options);
  }

  static equals(a: NodeList | PlainMessage<NodeList> | undefined, b: NodeList | PlainMessage<NodeList> | undefined): boolean {
    return proto3.util.equals(NodeList, a, b);
  }
}

/**
 * GetStatusRequest is a request to get the status of a node.
 *
 * @generated from message v1.GetStatusRequest
 */
export class GetStatusRequest extends Message<GetStatusRequest> {
  /**
   * ID is the ID of the node. If unset, the status of the local node is returned.
   *
   * @generated from field: string id = 1;
   */
  id = "";

  constructor(data?: PartialMessage<GetStatusRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "v1.GetStatusRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetStatusRequest {
    return new GetStatusRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetStatusRequest {
    return new GetStatusRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetStatusRequest {
    return new GetStatusRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetStatusRequest | PlainMessage<GetStatusRequest> | undefined, b: GetStatusRequest | PlainMessage<GetStatusRequest> | undefined): boolean {
    return proto3.util.equals(GetStatusRequest, a, b);
  }
}

/**
 * Status represents the status of a node.
 *
 * @generated from message v1.Status
 */
export class Status extends Message<Status> {
  /**
   * ID is the ID of the node.
   *
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * Description is an optional description provided by the node.
   *
   * @generated from field: string description = 2;
   */
  description = "";

  /**
   * Version is the version of the node.
   *
   * @generated from field: string version = 3;
   */
  version = "";

  /**
   * GitCommit is the git commit of the node.
   *
   * @generated from field: string gitCommit = 4;
   */
  gitCommit = "";

  /**
   * BuildDate is the build date of the node.
   *
   * @generated from field: string buildDate = 5;
   */
  buildDate = "";

  /**
   * Uptime is the uptime of the node.
   *
   * @generated from field: string uptime = 6;
   */
  uptime = "";

  /**
   * StartedAt is the time the node started.
   *
   * @generated from field: google.protobuf.Timestamp startedAt = 7;
   */
  startedAt?: Timestamp;

  /**
   * Features is the list of features currently enabled.
   *
   * @generated from field: repeated v1.FeaturePort features = 8;
   */
  features: FeaturePort[] = [];

  /**
   * ClusterStatus is the status of the node in the cluster.
   *
   * @generated from field: v1.ClusterStatus clusterStatus = 9;
   */
  clusterStatus = ClusterStatus.CLUSTER_STATUS_UNKNOWN;

  /**
   * CurrentLeader is the current leader of the cluster.
   *
   * @generated from field: string currentLeader = 10;
   */
  currentLeader = "";

  /**
   * InterfaceMetrics are the metrics for the node's interfaces.
   *
   * @generated from field: v1.InterfaceMetrics interfaceMetrics = 11;
   */
  interfaceMetrics?: InterfaceMetrics;

  constructor(data?: PartialMessage<Status>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "v1.Status";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "description", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "version", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "gitCommit", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "buildDate", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "uptime", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 7, name: "startedAt", kind: "message", T: Timestamp },
    { no: 8, name: "features", kind: "message", T: FeaturePort, repeated: true },
    { no: 9, name: "clusterStatus", kind: "enum", T: proto3.getEnumType(ClusterStatus) },
    { no: 10, name: "currentLeader", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 11, name: "interfaceMetrics", kind: "message", T: InterfaceMetrics },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Status {
    return new Status().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Status {
    return new Status().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Status {
    return new Status().fromJsonString(jsonString, options);
  }

  static equals(a: Status | PlainMessage<Status> | undefined, b: Status | PlainMessage<Status> | undefined): boolean {
    return proto3.util.equals(Status, a, b);
  }
}

/**
 * DataChannelNegotiation is the message for communicating data channels to nodes.
 *
 * @generated from message v1.DataChannelNegotiation
 */
export class DataChannelNegotiation extends Message<DataChannelNegotiation> {
  /**
   * Proto is the protocol of the traffic.
   *
   * @generated from field: string proto = 1;
   */
  proto = "";

  /**
   * Src is the address of the client that initiated the request.
   *
   * @generated from field: string src = 2;
   */
  src = "";

  /**
   * Dst is the destination address of the traffic.
   *
   * @generated from field: string dst = 3;
   */
  dst = "";

  /**
   * Port is the destination port of the traffic.
   *
   * @generated from field: uint32 port = 4;
   */
  port = 0;

  /**
   * Offer is the offer for the node to use as its local description.
   *
   * @generated from field: string offer = 5;
   */
  offer = "";

  /**
   * Answer is the answer for the node to use as its remote description.
   *
   * @generated from field: string answer = 6;
   */
  answer = "";

  /**
   * Candidate is an ICE candidate.
   *
   * @generated from field: string candidate = 7;
   */
  candidate = "";

  /**
   * StunServers is the list of STUN servers to use.
   *
   * @generated from field: repeated string stunServers = 8;
   */
  stunServers: string[] = [];

  constructor(data?: PartialMessage<DataChannelNegotiation>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "v1.DataChannelNegotiation";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "proto", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "src", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "dst", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "port", kind: "scalar", T: 13 /* ScalarType.UINT32 */ },
    { no: 5, name: "offer", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "answer", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 7, name: "candidate", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 8, name: "stunServers", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DataChannelNegotiation {
    return new DataChannelNegotiation().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DataChannelNegotiation {
    return new DataChannelNegotiation().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DataChannelNegotiation {
    return new DataChannelNegotiation().fromJsonString(jsonString, options);
  }

  static equals(a: DataChannelNegotiation | PlainMessage<DataChannelNegotiation> | undefined, b: DataChannelNegotiation | PlainMessage<DataChannelNegotiation> | undefined): boolean {
    return proto3.util.equals(DataChannelNegotiation, a, b);
  }
}

/**
 * InterfaceMetrics is the metrics for the WireGuard interface on a node.
 *
 * @generated from message v1.InterfaceMetrics
 */
export class InterfaceMetrics extends Message<InterfaceMetrics> {
  /**
   * DeviceName is the name of the device.
   *
   * @generated from field: string deviceName = 1;
   */
  deviceName = "";

  /**
   * PublicKey is the public key of the node.
   *
   * @generated from field: string publicKey = 2;
   */
  publicKey = "";

  /**
   * AddressV4 is the IPv4 address of the node.
   *
   * @generated from field: string addressV4 = 3;
   */
  addressV4 = "";

  /**
   * AddressV6 is the IPv6 address of the node.
   *
   * @generated from field: string addressV6 = 4;
   */
  addressV6 = "";

  /**
   * Type is the type of interface being used for wireguard.
   *
   * @generated from field: string type = 5;
   */
  type = "";

  /**
   * ListenPort is the port wireguard is listening on.
   *
   * @generated from field: int32 listenPort = 6;
   */
  listenPort = 0;

  /**
   * TotalReceiveBytes is the total number of bytes received.
   *
   * @generated from field: uint64 totalReceiveBytes = 7;
   */
  totalReceiveBytes = protoInt64.zero;

  /**
   * TotalTransmitBytes is the total number of bytes transmitted.
   *
   * @generated from field: uint64 totalTransmitBytes = 8;
   */
  totalTransmitBytes = protoInt64.zero;

  /**
   * NumPeers is the number of peers connected to the node.
   *
   * @generated from field: int32 numPeers = 9;
   */
  numPeers = 0;

  /**
   * Peers are the per-peer statistics.
   *
   * @generated from field: repeated v1.PeerMetrics peers = 10;
   */
  peers: PeerMetrics[] = [];

  constructor(data?: PartialMessage<InterfaceMetrics>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "v1.InterfaceMetrics";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "deviceName", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "publicKey", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "addressV4", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "addressV6", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "type", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "listenPort", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 7, name: "totalReceiveBytes", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
    { no: 8, name: "totalTransmitBytes", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
    { no: 9, name: "numPeers", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 10, name: "peers", kind: "message", T: PeerMetrics, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): InterfaceMetrics {
    return new InterfaceMetrics().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): InterfaceMetrics {
    return new InterfaceMetrics().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): InterfaceMetrics {
    return new InterfaceMetrics().fromJsonString(jsonString, options);
  }

  static equals(a: InterfaceMetrics | PlainMessage<InterfaceMetrics> | undefined, b: InterfaceMetrics | PlainMessage<InterfaceMetrics> | undefined): boolean {
    return proto3.util.equals(InterfaceMetrics, a, b);
  }
}

/**
 * PeerMetrics are the metrics for a node's peer.
 *
 * @generated from message v1.PeerMetrics
 */
export class PeerMetrics extends Message<PeerMetrics> {
  /**
   * PublicKey is the public key of the peer.
   *
   * @generated from field: string publicKey = 1;
   */
  publicKey = "";

  /**
   * Endpoint is the connected endpoint of the peer.
   *
   * @generated from field: string endpoint = 2;
   */
  endpoint = "";

  /**
   * PersistentKeepAlive is the persistent keep alive interval for the peer.
   *
   * @generated from field: string persistentKeepAlive = 3;
   */
  persistentKeepAlive = "";

  /**
   * LastHandshakeTime is the last handshake time for the peer.
   *
   * @generated from field: string lastHandshakeTime = 4;
   */
  lastHandshakeTime = "";

  /**
   * AllowedIPs is the list of allowed IPs for the peer.
   *
   * @generated from field: repeated string allowedIPs = 5;
   */
  allowedIPs: string[] = [];

  /**
   * ProtocolVersion is the version of the wireguard protocol negotiated with the peer.
   *
   * @generated from field: int64 protocolVersion = 6;
   */
  protocolVersion = protoInt64.zero;

  /**
   * ReceiveBytes is the bytes received from the peer.
   *
   * @generated from field: uint64 receiveBytes = 7;
   */
  receiveBytes = protoInt64.zero;

  /**
   * TransmitBytes is the bytes transmitted to the peer.
   *
   * @generated from field: uint64 transmitBytes = 8;
   */
  transmitBytes = protoInt64.zero;

  constructor(data?: PartialMessage<PeerMetrics>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "v1.PeerMetrics";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "publicKey", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "endpoint", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "persistentKeepAlive", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "lastHandshakeTime", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "allowedIPs", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 6, name: "protocolVersion", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 7, name: "receiveBytes", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
    { no: 8, name: "transmitBytes", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): PeerMetrics {
    return new PeerMetrics().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): PeerMetrics {
    return new PeerMetrics().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): PeerMetrics {
    return new PeerMetrics().fromJsonString(jsonString, options);
  }

  static equals(a: PeerMetrics | PlainMessage<PeerMetrics> | undefined, b: PeerMetrics | PlainMessage<PeerMetrics> | undefined): boolean {
    return proto3.util.equals(PeerMetrics, a, b);
  }
}

/**
 * WebRTCSignal is a signal sent to a remote peer over the WebRTC API.
 *
 * @generated from message v1.WebRTCSignal
 */
export class WebRTCSignal extends Message<WebRTCSignal> {
  /**
   * NodeID is the ID of the node to send the signal to.
   * This is set by the original sender. On the node that
   * receives the ReceiveSignalChannel request, this will
   * be set to the ID of the node that sent the request.
   *
   * @generated from field: string nodeID = 1;
   */
  nodeID = "";

  /**
   * Candidate is an ICE candidate.
   *
   * @generated from field: string candidate = 2;
   */
  candidate = "";

  /**
   * Description is a session description.
   *
   * @generated from field: string description = 3;
   */
  description = "";

  constructor(data?: PartialMessage<WebRTCSignal>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "v1.WebRTCSignal";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "nodeID", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "candidate", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "description", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): WebRTCSignal {
    return new WebRTCSignal().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): WebRTCSignal {
    return new WebRTCSignal().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): WebRTCSignal {
    return new WebRTCSignal().fromJsonString(jsonString, options);
  }

  static equals(a: WebRTCSignal | PlainMessage<WebRTCSignal> | undefined, b: WebRTCSignal | PlainMessage<WebRTCSignal> | undefined): boolean {
    return proto3.util.equals(WebRTCSignal, a, b);
  }
}

