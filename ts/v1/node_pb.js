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

// @generated by protoc-gen-es v1.4.0
// @generated from file v1/node.proto (package v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3, Timestamp } from "@bufbuild/protobuf";

/**
 * ClusterStatus is the status of the node in the cluster.
 *
 * @generated from enum v1.ClusterStatus
 */
export const ClusterStatus = proto3.makeEnum(
  "v1.ClusterStatus",
  [
    {no: 0, name: "CLUSTER_STATUS_UNKNOWN"},
    {no: 1, name: "CLUSTER_LEADER"},
    {no: 2, name: "CLUSTER_VOTER"},
    {no: 3, name: "CLUSTER_OBSERVER"},
    {no: 4, name: "CLUSTER_NODE"},
  ],
);

/**
 * Feature is a list of features supported by a node.
 *
 * @generated from enum v1.Feature
 */
export const Feature = proto3.makeEnum(
  "v1.Feature",
  [
    {no: 0, name: "FEATURE_NONE"},
    {no: 1, name: "NODES"},
    {no: 2, name: "LEADER_PROXY"},
    {no: 3, name: "MESH_API"},
    {no: 4, name: "ADMIN_API"},
    {no: 5, name: "MEMBERSHIP"},
    {no: 6, name: "METRICS"},
    {no: 7, name: "ICE_NEGOTIATION"},
    {no: 8, name: "TURN_SERVER"},
    {no: 9, name: "MESH_DNS"},
    {no: 10, name: "FORWARD_MESH_DNS"},
    {no: 11, name: "STORAGE_QUERIER"},
    {no: 12, name: "STORAGE_PROVIDER"},
    {no: 13, name: "REGISTRAR"},
  ],
);

/**
 * EdgeAttribute are pre-defined edge attributes. They should
 * be used as their string values.
 *
 * @generated from enum v1.EdgeAttribute
 */
export const EdgeAttribute = proto3.makeEnum(
  "v1.EdgeAttribute",
  [
    {no: 0, name: "EDGE_ATTRIBUTE_UNKNOWN", localName: "UNKNOWN"},
    {no: 1, name: "EDGE_ATTRIBUTE_NATIVE", localName: "NATIVE"},
    {no: 2, name: "EDGE_ATTRIBUTE_ICE", localName: "ICE"},
    {no: 3, name: "EDGE_ATTRIBUTE_LIBP2P", localName: "LIBP2P"},
  ],
);

/**
 * DataChannel are the data channels used when communicating over ICE
 * with a node.
 *
 * @generated from enum v1.DataChannel
 */
export const DataChannel = proto3.makeEnum(
  "v1.DataChannel",
  [
    {no: 0, name: "CHANNELS"},
    {no: 1, name: "CONNECTIONS"},
  ],
);

/**
 * FeaturePort describes a feature and the port it is advertised on.
 *
 * @generated from message v1.FeaturePort
 */
export const FeaturePort = proto3.makeMessageType(
  "v1.FeaturePort",
  () => [
    { no: 1, name: "feature", kind: "enum", T: proto3.getEnumType(Feature) },
    { no: 2, name: "port", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ],
);

/**
 * MeshNode is a node that has been registered with the mesh.
 *
 * @generated from message v1.MeshNode
 */
export const MeshNode = proto3.makeMessageType(
  "v1.MeshNode",
  () => [
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
  ],
);

/**
 * NodeList is a list of nodes.
 *
 * @generated from message v1.NodeList
 */
export const NodeList = proto3.makeMessageType(
  "v1.NodeList",
  () => [
    { no: 1, name: "nodes", kind: "message", T: MeshNode, repeated: true },
  ],
);

/**
 * GetStatusRequest is a request to get the status of a node.
 *
 * @generated from message v1.GetStatusRequest
 */
export const GetStatusRequest = proto3.makeMessageType(
  "v1.GetStatusRequest",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * Status represents the status of a node.
 *
 * @generated from message v1.Status
 */
export const Status = proto3.makeMessageType(
  "v1.Status",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "description", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "version", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "commit", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "build_date", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "uptime", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 7, name: "startedAt", kind: "message", T: Timestamp },
    { no: 8, name: "features", kind: "message", T: FeaturePort, repeated: true },
    { no: 9, name: "clusterStatus", kind: "enum", T: proto3.getEnumType(ClusterStatus) },
    { no: 10, name: "currentLeader", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 11, name: "interfaceMetrics", kind: "message", T: InterfaceMetrics },
  ],
);

/**
 * DataChannelNegotiation is the message for communicating data channels to nodes.
 *
 * @generated from message v1.DataChannelNegotiation
 */
export const DataChannelNegotiation = proto3.makeMessageType(
  "v1.DataChannelNegotiation",
  () => [
    { no: 1, name: "proto", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "src", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "dst", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "port", kind: "scalar", T: 13 /* ScalarType.UINT32 */ },
    { no: 5, name: "offer", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "answer", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 7, name: "candidate", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 8, name: "stunServers", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ],
);

/**
 * InterfaceMetrics is the metrics for the WireGuard interface on a node.
 *
 * @generated from message v1.InterfaceMetrics
 */
export const InterfaceMetrics = proto3.makeMessageType(
  "v1.InterfaceMetrics",
  () => [
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
  ],
);

/**
 * PeerMetrics are the metrics for a node's peer.
 *
 * @generated from message v1.PeerMetrics
 */
export const PeerMetrics = proto3.makeMessageType(
  "v1.PeerMetrics",
  () => [
    { no: 1, name: "publicKey", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "endpoint", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "persistentKeepAlive", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "lastHandshakeTime", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "allowedIPs", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 6, name: "protocolVersion", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 7, name: "receiveBytes", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
    { no: 8, name: "transmitBytes", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
  ],
);

/**
 * WebRTCSignal is a signal sent to a remote peer over the WebRTC API.
 *
 * @generated from message v1.WebRTCSignal
 */
export const WebRTCSignal = proto3.makeMessageType(
  "v1.WebRTCSignal",
  () => [
    { no: 1, name: "nodeID", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "candidate", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "description", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);
