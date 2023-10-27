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

// @generated by protoc-gen-es v1.4.1
// @generated from file v1/members.proto (package v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";
import { ClusterStatus, FeaturePort, MeshNode } from "./node_pb.js";

/**
 * ConnectProtocol is a type of protocol for establishing a connection into a mesh.
 *
 * @generated from enum v1.ConnectProtocol
 */
export const ConnectProtocol = proto3.makeEnum(
  "v1.ConnectProtocol",
  [
    {no: 0, name: "CONNECT_NATIVE"},
    {no: 1, name: "CONNECT_ICE"},
    {no: 2, name: "CONNECT_LIBP2P"},
  ],
);

/**
 * JoinRequest is a request to join the cluster.
 *
 * @generated from message v1.JoinRequest
 */
export const JoinRequest = proto3.makeMessageType(
  "v1.JoinRequest",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "publicKey", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "primaryEndpoint", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "wireguardEndpoints", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 6, name: "zoneAwarenessID", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 7, name: "assignIPv4", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 8, name: "preferStorageIPv6", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 9, name: "asVoter", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 10, name: "asObserver", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 11, name: "routes", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 12, name: "directPeers", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "enum", T: proto3.getEnumType(ConnectProtocol)} },
    { no: 13, name: "features", kind: "message", T: FeaturePort, repeated: true },
    { no: 14, name: "multiaddrs", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ],
);

/**
 * JoinResponse is a response to a join request.
 *
 * @generated from message v1.JoinResponse
 */
export const JoinResponse = proto3.makeMessageType(
  "v1.JoinResponse",
  () => [
    { no: 1, name: "addressIPv4", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "addressIPv6", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "networkIPv4", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "networkIPv6", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "peers", kind: "message", T: WireGuardPeer, repeated: true },
    { no: 6, name: "iceServers", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 7, name: "dnsServers", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 8, name: "meshDomain", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * UpdateRequest contains most of the same fields as JoinRequest, but is
 * used to update the state of a node in the cluster.
 *
 * @generated from message v1.UpdateRequest
 */
export const UpdateRequest = proto3.makeMessageType(
  "v1.UpdateRequest",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "publicKey", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "primaryEndpoint", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "wireguardEndpoints", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 5, name: "zoneAwarenessID", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "asVoter", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 7, name: "routes", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 8, name: "features", kind: "message", T: FeaturePort, repeated: true },
    { no: 9, name: "multiaddrs", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ],
);

/**
 * UpdateResponse is a response to an update request. It is currently empty.
 *
 * @generated from message v1.UpdateResponse
 */
export const UpdateResponse = proto3.makeMessageType(
  "v1.UpdateResponse",
  [],
);

/**
 * WireGuardPeer is a peer in the Wireguard network.
 *
 * @generated from message v1.WireGuardPeer
 */
export const WireGuardPeer = proto3.makeMessageType(
  "v1.WireGuardPeer",
  () => [
    { no: 1, name: "node", kind: "message", T: MeshNode },
    { no: 2, name: "allowedIPs", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 3, name: "allowedRoutes", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 4, name: "proto", kind: "enum", T: proto3.getEnumType(ConnectProtocol) },
  ],
);

/**
 * LeaveRequest is a request to leave the cluster.
 *
 * @generated from message v1.LeaveRequest
 */
export const LeaveRequest = proto3.makeMessageType(
  "v1.LeaveRequest",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * LeaveResponse is a response to a leave request. It is currently empty.
 *
 * @generated from message v1.LeaveResponse
 */
export const LeaveResponse = proto3.makeMessageType(
  "v1.LeaveResponse",
  [],
);

/**
 * StorageConsensusRequest is a request to get the current Storage configuration.
 *
 * @generated from message v1.StorageConsensusRequest
 */
export const StorageConsensusRequest = proto3.makeMessageType(
  "v1.StorageConsensusRequest",
  [],
);

/**
 * StorageConsensusResponse is a response to a Storage consensus request.
 *
 * @generated from message v1.StorageConsensusResponse
 */
export const StorageConsensusResponse = proto3.makeMessageType(
  "v1.StorageConsensusResponse",
  () => [
    { no: 1, name: "servers", kind: "message", T: StorageServer, repeated: true },
  ],
);

/**
 * StorageServer is a server in the Storage configuration.
 *
 * @generated from message v1.StorageServer
 */
export const StorageServer = proto3.makeMessageType(
  "v1.StorageServer",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "suffrage", kind: "enum", T: proto3.getEnumType(ClusterStatus) },
    { no: 3, name: "publicKey", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "address", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * SubscribePeersRequest is a request to subscribe to peer updates.
 *
 * @generated from message v1.SubscribePeersRequest
 */
export const SubscribePeersRequest = proto3.makeMessageType(
  "v1.SubscribePeersRequest",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * PeerConfigurations is a stream of peer configurations.
 *
 * @generated from message v1.PeerConfigurations
 */
export const PeerConfigurations = proto3.makeMessageType(
  "v1.PeerConfigurations",
  () => [
    { no: 5, name: "peers", kind: "message", T: WireGuardPeer, repeated: true },
    { no: 6, name: "iceServers", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 7, name: "dnsServers", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ],
);

