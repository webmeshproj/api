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
// @generated from file v1/members.proto (package v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import type { ClusterStatus, FeaturePort, MeshNode } from "./node_pb.js";

/**
 * ConnectProtocol is a type of protocol for establishing a connection into a mesh.
 *
 * @generated from enum v1.ConnectProtocol
 */
export declare enum ConnectProtocol {
  /**
   * CONNECT_NATIVE indicates that the node should connect to other nodes via the native
   * webmesh mechanisms.
   *
   * @generated from enum value: CONNECT_NATIVE = 0;
   */
  CONNECT_NATIVE = 0,

  /**
   * CONNECT_ICE indicates that the node should connect to other nodes via ICE.
   *
   * @generated from enum value: CONNECT_ICE = 1;
   */
  CONNECT_ICE = 1,

  /**
   * CONNECT_LIBP2P indicates that the node should connect to other nodes via libp2p.
   *
   * @generated from enum value: CONNECT_LIBP2P = 2;
   */
  CONNECT_LIBP2P = 2,
}

/**
 * JoinRequest is a request to join the cluster.
 *
 * @generated from message v1.JoinRequest
 */
export declare class JoinRequest extends Message<JoinRequest> {
  /**
   * id is the ID of the node.
   *
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * public_key is the public key of the node to broadcast to peers.
   *
   * @generated from field: string publicKey = 2;
   */
  publicKey: string;

  /**
   * primary_endpoint is a routable address for the node. If left unset, 
   * the node is assumed to be behind a NAT and not directly accessible.
   *
   * @generated from field: string primaryEndpoint = 4;
   */
  primaryEndpoint: string;

  /**
   * wireguard_endpoints is a list of WireGuard endpoints for the node.
   *
   * @generated from field: repeated string wireguardEndpoints = 5;
   */
  wireguardEndpoints: string[];

  /**
   * zone_awareness_id is the zone awareness ID of the node.
   *
   * @generated from field: string zoneAwarenessID = 6;
   */
  zoneAwarenessID: string;

  /**
   * assign_ipv4 is whether an IPv4 address should be assigned to the node.
   *
   * @generated from field: bool assignIPv4 = 7;
   */
  assignIPv4: boolean;

  /**
   * prefer_storage_ipv6 is whether IPv6 should be preferred over IPv4 for storage communication.
   * This is only used if assign_ipv4 is true.
   *
   * @generated from field: bool preferStorageIPv6 = 8;
   */
  preferStorageIPv6: boolean;

  /**
   * as_voter is whether the node should receive a vote in elections. The request
   * will be denied if the node is not allowed to vote.
   *
   * @generated from field: bool asVoter = 9;
   */
  asVoter: boolean;

  /**
   * as_observer is whether the node should be added as an observer. They will receive
   * updates to the storage, but not be able to vote in elections.
   *
   * @generated from field: bool asObserver = 10;
   */
  asObserver: boolean;

  /**
   * routes is a list of routes to advertise to peers. The request will be denied
   * if the node is not allowed to put routes.
   *
   * @generated from field: repeated string routes = 11;
   */
  routes: string[];

  /**
   * direct_peers is a map of extra peers that should be connected to directly over relays. 
   * The provided edge attribute is the callers preference of how the relay should be created.
   * The request will be denied if the node is not allowed to put data channels or edges.
   * The default joining behavior creates direct links between the caller and the joiner.
   * If the caller has a primary endpoint, the joiner will link the caller to all
   * other nodes with a primary endpoint. If the caller has a zone awareness ID,
   * the joiner will link the caller to all other nodes with the same zone awareness ID
   * that also have a primary endpoint.
   *
   * @generated from field: map<string, v1.ConnectProtocol> directPeers = 12;
   */
  directPeers: { [key: string]: ConnectProtocol };

  /**
   * features is a list of features supported by the node that should be advertised to peers
   * and the port they are available on.
   *
   * @generated from field: repeated v1.FeaturePort features = 13;
   */
  features: FeaturePort[];

  /**
   * multiaddrs are libp2p multiaddresses this node is listening on.
   *
   * @generated from field: repeated string multiaddrs = 14;
   */
  multiaddrs: string[];

  constructor(data?: PartialMessage<JoinRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "v1.JoinRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): JoinRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): JoinRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): JoinRequest;

  static equals(a: JoinRequest | PlainMessage<JoinRequest> | undefined, b: JoinRequest | PlainMessage<JoinRequest> | undefined): boolean;
}

/**
 * JoinResponse is a response to a join request.
 *
 * @generated from message v1.JoinResponse
 */
export declare class JoinResponse extends Message<JoinResponse> {
  /**
   * address_ipv4 is the private IPv4 wireguard address of the node
   * in CIDR format representing the network. This is only set if
   * assign_ipv4 was set in the request or no network_ipv6 was provided.
   *
   * @generated from field: string addressIPv4 = 1;
   */
  addressIPv4: string;

  /**
   * address_ipv6 is the IPv6 network assigned to the node.
   *
   * @generated from field: string addressIPv6 = 2;
   */
  addressIPv6: string;

  /**
   * network_ipv4 is the IPv4 network of the Mesh.
   *
   * @generated from field: string networkIPv4 = 3;
   */
  networkIPv4: string;

  /**
   * network_ipv6 is the IPv6 network of the Mesh.
   *
   * @generated from field: string networkIPv6 = 4;
   */
  networkIPv6: string;

  /**
   * peers is a list of wireguard peers to connect to.
   *
   * @generated from field: repeated v1.WireGuardPeer peers = 5;
   */
  peers: WireGuardPeer[];

  /**
   * ice_servers is a list of public nodes that can be used to negotiate
   * ICE connections if required. This may only be populated when one of
   * the peers has the ICE flag set. This must be set if the requestor
   * specifies direct_peers.
   *
   * @generated from field: repeated string iceServers = 6;
   */
  iceServers: string[];

  /**
   * dns_servers is a list of peers offering DNS services.
   *
   * @generated from field: repeated string dnsServers = 7;
   */
  dnsServers: string[];

  /**
   * mesh_domain is the domain of the mesh.
   *
   * @generated from field: string meshDomain = 8;
   */
  meshDomain: string;

  constructor(data?: PartialMessage<JoinResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "v1.JoinResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): JoinResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): JoinResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): JoinResponse;

  static equals(a: JoinResponse | PlainMessage<JoinResponse> | undefined, b: JoinResponse | PlainMessage<JoinResponse> | undefined): boolean;
}

/**
 * UpdateRequest contains most of the same fields as JoinRequest, but is
 * used to update the state of a node in the cluster.
 *
 * @generated from message v1.UpdateRequest
 */
export declare class UpdateRequest extends Message<UpdateRequest> {
  /**
   * id is the ID of the node.
   *
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * public_key is the public key of the node to broadcast to peers.
   *
   * @generated from field: string publicKey = 2;
   */
  publicKey: string;

  /**
   * primary_endpoint is a routable address for the node. If left unset, 
   * the node is assumed to be behind a NAT and not directly accessible.
   *
   * @generated from field: string primaryEndpoint = 3;
   */
  primaryEndpoint: string;

  /**
   * wireguard_endpoints is a list of WireGuard endpoints for the node.
   *
   * @generated from field: repeated string wireguardEndpoints = 4;
   */
  wireguardEndpoints: string[];

  /**
   * zone_awareness_id is the zone awareness ID of the node.
   *
   * @generated from field: string zoneAwarenessID = 5;
   */
  zoneAwarenessID: string;

  /**
   * as_voter is whether the node should receive a vote in elections. The request
   * will be denied if the node is not allowed to vote.
   *
   * @generated from field: bool asVoter = 6;
   */
  asVoter: boolean;

  /**
   * routes is a list of routes to advertise to peers. The request will be denied
   * if the node is not allowed to put routes.
   *
   * @generated from field: repeated string routes = 7;
   */
  routes: string[];

  /**
   * features is a list of features supported by the node that should be advertised to peers
   * and the port they are available on.
   *
   * @generated from field: repeated v1.FeaturePort features = 8;
   */
  features: FeaturePort[];

  /**
   * multiaddrs are libp2p multiaddresses this node is listening on.
   *
   * @generated from field: repeated string multiaddrs = 9;
   */
  multiaddrs: string[];

  constructor(data?: PartialMessage<UpdateRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "v1.UpdateRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateRequest;

  static equals(a: UpdateRequest | PlainMessage<UpdateRequest> | undefined, b: UpdateRequest | PlainMessage<UpdateRequest> | undefined): boolean;
}

/**
 * UpdateResponse is a response to an update request. It is currently empty.
 *
 * @generated from message v1.UpdateResponse
 */
export declare class UpdateResponse extends Message<UpdateResponse> {
  constructor(data?: PartialMessage<UpdateResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "v1.UpdateResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateResponse;

  static equals(a: UpdateResponse | PlainMessage<UpdateResponse> | undefined, b: UpdateResponse | PlainMessage<UpdateResponse> | undefined): boolean;
}

/**
 * WireGuardPeer is a peer in the Wireguard network.
 *
 * @generated from message v1.WireGuardPeer
 */
export declare class WireGuardPeer extends Message<WireGuardPeer> {
  /**
   * Node is information about this node.
   *
   * @generated from field: v1.MeshNode node = 1;
   */
  node?: MeshNode;

  /**
   * allowed_ips is the list of allowed IPs for the peer.
   *
   * @generated from field: repeated string allowedIPs = 2;
   */
  allowedIPs: string[];

  /**
   * allowed_routes is the list of allowed routes for the peer.
   *
   * @generated from field: repeated string allowedRoutes = 3;
   */
  allowedRoutes: string[];

  /**
   * proto indicates the protocol to use to connect to the peer.
   *
   * @generated from field: v1.ConnectProtocol proto = 4;
   */
  proto: ConnectProtocol;

  constructor(data?: PartialMessage<WireGuardPeer>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "v1.WireGuardPeer";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): WireGuardPeer;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): WireGuardPeer;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): WireGuardPeer;

  static equals(a: WireGuardPeer | PlainMessage<WireGuardPeer> | undefined, b: WireGuardPeer | PlainMessage<WireGuardPeer> | undefined): boolean;
}

/**
 * LeaveRequest is a request to leave the cluster.
 *
 * @generated from message v1.LeaveRequest
 */
export declare class LeaveRequest extends Message<LeaveRequest> {
  /**
   * id is the ID of the node.
   *
   * @generated from field: string id = 1;
   */
  id: string;

  constructor(data?: PartialMessage<LeaveRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "v1.LeaveRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): LeaveRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): LeaveRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): LeaveRequest;

  static equals(a: LeaveRequest | PlainMessage<LeaveRequest> | undefined, b: LeaveRequest | PlainMessage<LeaveRequest> | undefined): boolean;
}

/**
 * LeaveResponse is a response to a leave request. It is currently empty.
 *
 * @generated from message v1.LeaveResponse
 */
export declare class LeaveResponse extends Message<LeaveResponse> {
  constructor(data?: PartialMessage<LeaveResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "v1.LeaveResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): LeaveResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): LeaveResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): LeaveResponse;

  static equals(a: LeaveResponse | PlainMessage<LeaveResponse> | undefined, b: LeaveResponse | PlainMessage<LeaveResponse> | undefined): boolean;
}

/**
 * StorageConsensusRequest is a request to get the current Storage configuration.
 *
 * @generated from message v1.StorageConsensusRequest
 */
export declare class StorageConsensusRequest extends Message<StorageConsensusRequest> {
  constructor(data?: PartialMessage<StorageConsensusRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "v1.StorageConsensusRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): StorageConsensusRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): StorageConsensusRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): StorageConsensusRequest;

  static equals(a: StorageConsensusRequest | PlainMessage<StorageConsensusRequest> | undefined, b: StorageConsensusRequest | PlainMessage<StorageConsensusRequest> | undefined): boolean;
}

/**
 * StorageConsensusResponse is a response to a Storage consensus request.
 *
 * @generated from message v1.StorageConsensusResponse
 */
export declare class StorageConsensusResponse extends Message<StorageConsensusResponse> {
  /**
   * servers is the list of servers in the storage configuration.
   *
   * @generated from field: repeated v1.StorageServer servers = 1;
   */
  servers: StorageServer[];

  constructor(data?: PartialMessage<StorageConsensusResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "v1.StorageConsensusResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): StorageConsensusResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): StorageConsensusResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): StorageConsensusResponse;

  static equals(a: StorageConsensusResponse | PlainMessage<StorageConsensusResponse> | undefined, b: StorageConsensusResponse | PlainMessage<StorageConsensusResponse> | undefined): boolean;
}

/**
 * StorageServer is a server in the Storage configuration.
 *
 * @generated from message v1.StorageServer
 */
export declare class StorageServer extends Message<StorageServer> {
  /**
   * ID is the ID of the server.
   *
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * Suffrage is the suffrage of the server.
   *
   * @generated from field: v1.ClusterStatus suffrage = 2;
   */
  suffrage: ClusterStatus;

  /**
   * PublicKey is the public key of this server. Not all storage providers track this field.
   *
   * @generated from field: string publicKey = 3;
   */
  publicKey: string;

  /**
   * Address is the mesh address of the server.
   *
   * @generated from field: string address = 4;
   */
  address: string;

  constructor(data?: PartialMessage<StorageServer>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "v1.StorageServer";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): StorageServer;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): StorageServer;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): StorageServer;

  static equals(a: StorageServer | PlainMessage<StorageServer> | undefined, b: StorageServer | PlainMessage<StorageServer> | undefined): boolean;
}

/**
 * SubscribePeersRequest is a request to subscribe to peer updates.
 *
 * @generated from message v1.SubscribePeersRequest
 */
export declare class SubscribePeersRequest extends Message<SubscribePeersRequest> {
  /**
   * id is the ID of the node.
   *
   * @generated from field: string id = 1;
   */
  id: string;

  constructor(data?: PartialMessage<SubscribePeersRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "v1.SubscribePeersRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SubscribePeersRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SubscribePeersRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SubscribePeersRequest;

  static equals(a: SubscribePeersRequest | PlainMessage<SubscribePeersRequest> | undefined, b: SubscribePeersRequest | PlainMessage<SubscribePeersRequest> | undefined): boolean;
}

/**
 * PeerConfigurations is a stream of peer configurations.
 *
 * @generated from message v1.PeerConfigurations
 */
export declare class PeerConfigurations extends Message<PeerConfigurations> {
  /**
   * peers is a list of wireguard peers to connect to.
   *
   * @generated from field: repeated v1.WireGuardPeer peers = 5;
   */
  peers: WireGuardPeer[];

  /**
   * ice_servers is a list of public nodes that can be used to negotiate
   * ICE connections if required. This may only be populated when one of
   * the peers has the ICE flag set.
   *
   * @generated from field: repeated string iceServers = 6;
   */
  iceServers: string[];

  /**
   * dns_servers is a list of peers offering DNS services.
   *
   * @generated from field: repeated string dnsServers = 7;
   */
  dnsServers: string[];

  constructor(data?: PartialMessage<PeerConfigurations>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "v1.PeerConfigurations";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): PeerConfigurations;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): PeerConfigurations;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): PeerConfigurations;

  static equals(a: PeerConfigurations | PlainMessage<PeerConfigurations> | undefined, b: PeerConfigurations | PlainMessage<PeerConfigurations> | undefined): boolean;
}
