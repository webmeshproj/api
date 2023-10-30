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
// @generated from file v1/plugin.proto (package v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage, Struct } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import type { MeshNode } from "./node_pb.js";

/**
 * PluginConfiguration is the message containing the configuration of a plugin.
 *
 * @generated from message v1.PluginConfiguration
 */
export declare class PluginConfiguration extends Message<PluginConfiguration> {
  /**
   * Config is the configuration for the plugin. This will be specific
   * for each plugin.
   *
   * @generated from field: google.protobuf.Struct config = 1;
   */
  config?: Struct;

  /**
   * NodeConfig is the configuration of the node and the network that it is a part of.
   *
   * @generated from field: v1.NodeConfiguration nodeConfig = 2;
   */
  nodeConfig?: NodeConfiguration;

  constructor(data?: PartialMessage<PluginConfiguration>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "v1.PluginConfiguration";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): PluginConfiguration;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): PluginConfiguration;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): PluginConfiguration;

  static equals(a: PluginConfiguration | PlainMessage<PluginConfiguration> | undefined, b: PluginConfiguration | PlainMessage<PluginConfiguration> | undefined): boolean;
}

/**
 * NodeConfiguration is the message containing the configuration of the
 * node and the network that it is a part of.
 *
 * @generated from message v1.NodeConfiguration
 */
export declare class NodeConfiguration extends Message<NodeConfiguration> {
  /**
   * ID is the ID of the node.
   *
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * NetworkIPv4 is the IPv4 network that the node is a part of.
   *
   * @generated from field: string networkIPv4 = 2;
   */
  networkIPv4: string;

  /**
   * NetworkIPv6 is the IPv6 network that the node is a part of.
   *
   * @generated from field: string networkIPv6 = 3;
   */
  networkIPv6: string;

  /**
   * AddressIPv4 is the IPv4 address of the node.
   *
   * @generated from field: string addressIPv4 = 4;
   */
  addressIPv4: string;

  /**
   * AddressIPv6 is the IPv6 address of the node.
   *
   * @generated from field: string addressIPv6 = 5;
   */
  addressIPv6: string;

  /**
   * Domain is the domain of the network.
   *
   * @generated from field: string domain = 6;
   */
  domain: string;

  /**
   * PrivateKey is the private key of the node.
   *
   * @generated from field: bytes privateKey = 7;
   */
  privateKey: Uint8Array;

  constructor(data?: PartialMessage<NodeConfiguration>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "v1.NodeConfiguration";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): NodeConfiguration;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): NodeConfiguration;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): NodeConfiguration;

  static equals(a: NodeConfiguration | PlainMessage<NodeConfiguration> | undefined, b: NodeConfiguration | PlainMessage<NodeConfiguration> | undefined): boolean;
}

/**
 * PluginInfo is the information of a plugin.
 *
 * @generated from message v1.PluginInfo
 */
export declare class PluginInfo extends Message<PluginInfo> {
  /**
   * Name is the name of the plugin.
   *
   * @generated from field: string name = 1;
   */
  name: string;

  /**
   * Version is the version of the plugin.
   *
   * @generated from field: string version = 2;
   */
  version: string;

  /**
   * Description is the description of the plugin.
   *
   * @generated from field: string description = 3;
   */
  description: string;

  /**
   * Capabilities is the capabilities of the plugin.
   *
   * @generated from field: repeated v1.PluginInfo.PluginCapability capabilities = 5;
   */
  capabilities: PluginInfo_PluginCapability[];

  constructor(data?: PartialMessage<PluginInfo>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "v1.PluginInfo";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): PluginInfo;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): PluginInfo;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): PluginInfo;

  static equals(a: PluginInfo | PlainMessage<PluginInfo> | undefined, b: PluginInfo | PlainMessage<PluginInfo> | undefined): boolean;
}

/**
 * PluginCapability is the capabilities of a plugin.
 *
 * @generated from enum v1.PluginInfo.PluginCapability
 */
export declare enum PluginInfo_PluginCapability {
  /**
   * UNKNOWN is the default value of PluginCapability.
   *
   * @generated from enum value: UNKNOWN = 0;
   */
  UNKNOWN = 0,

  /**
   * STORAGE_PROVIDER indicates that the plugin can provide storage and underlying consistency.
   *
   * @generated from enum value: STORAGE_PROVIDER = 1;
   */
  STORAGE_PROVIDER = 1,

  /**
   * AUTH indicates that the plugin is an auth plugin.
   *
   * @generated from enum value: AUTH = 2;
   */
  AUTH = 2,

  /**
   * WATCH indicates that the plugin wants to receive watch events.
   *
   * @generated from enum value: WATCH = 3;
   */
  WATCH = 3,

  /**
   * IPAMV4 indicates that the plugin is an IPv4 IPAM plugin.
   *
   * @generated from enum value: IPAMV4 = 4;
   */
  IPAMV4 = 4,

  /**
   * STORAGE_QUERIER indicates a plugin that wants to interact with storage.
   *
   * @generated from enum value: STORAGE_QUERIER = 5;
   */
  STORAGE_QUERIER = 5,
}

/**
 * AuthenticationRequest is the message containing an authentication request.
 *
 * @generated from message v1.AuthenticationRequest
 */
export declare class AuthenticationRequest extends Message<AuthenticationRequest> {
  /**
   * Headers are the headers of the request.
   *
   * @generated from field: map<string, string> headers = 1;
   */
  headers: { [key: string]: string };

  /**
   * Certificates are the DER encoded certificates of the request.
   *
   * @generated from field: repeated bytes certificates = 2;
   */
  certificates: Uint8Array[];

  constructor(data?: PartialMessage<AuthenticationRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "v1.AuthenticationRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AuthenticationRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AuthenticationRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AuthenticationRequest;

  static equals(a: AuthenticationRequest | PlainMessage<AuthenticationRequest> | undefined, b: AuthenticationRequest | PlainMessage<AuthenticationRequest> | undefined): boolean;
}

/**
 * AuthenticationResponse is the message containing an authentication response.
 *
 * @generated from message v1.AuthenticationResponse
 */
export declare class AuthenticationResponse extends Message<AuthenticationResponse> {
  /**
   * ID is the id of the authenticated user.
   *
   * @generated from field: string id = 1;
   */
  id: string;

  constructor(data?: PartialMessage<AuthenticationResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "v1.AuthenticationResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AuthenticationResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AuthenticationResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AuthenticationResponse;

  static equals(a: AuthenticationResponse | PlainMessage<AuthenticationResponse> | undefined, b: AuthenticationResponse | PlainMessage<AuthenticationResponse> | undefined): boolean;
}

/**
 * Event is the message containing a watch event.
 *
 * @generated from message v1.Event
 */
export declare class Event extends Message<Event> {
  /**
   * Type is the type of the watch event.
   *
   * @generated from field: v1.Event.WatchEvent type = 1;
   */
  type: Event_WatchEvent;

  /**
   * Event is the data of the watch event.
   *
   * @generated from oneof v1.Event.event
   */
  event: {
    /**
     * Node is the node that the event is about.
     *
     * @generated from field: v1.MeshNode node = 2;
     */
    value: MeshNode;
    case: "node";
  } | { case: undefined; value?: undefined };

  constructor(data?: PartialMessage<Event>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "v1.Event";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Event;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Event;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Event;

  static equals(a: Event | PlainMessage<Event> | undefined, b: Event | PlainMessage<Event> | undefined): boolean;
}

/**
 * WatchEvent is the type of a watch event.
 *
 * @generated from enum v1.Event.WatchEvent
 */
export declare enum Event_WatchEvent {
  /**
   * UNKNOWN is the default value of WatchEvent.
   *
   * @generated from enum value: UNKNOWN = 0;
   */
  UNKNOWN = 0,

  /**
   * NODE_JOIN indicates that a node has joined the cluster.
   *
   * @generated from enum value: NODE_JOIN = 1;
   */
  NODE_JOIN = 1,

  /**
   * NODE_LEAVE indicates that a node has left the cluster.
   *
   * @generated from enum value: NODE_LEAVE = 2;
   */
  NODE_LEAVE = 2,

  /**
   * LEADER_CHANGE indicates that the leader of the cluster has changed.
   *
   * @generated from enum value: LEADER_CHANGE = 3;
   */
  LEADER_CHANGE = 3,
}

/**
 * AllocateIPRequest is the message containing an IP allocation request.
 *
 * @generated from message v1.AllocateIPRequest
 */
export declare class AllocateIPRequest extends Message<AllocateIPRequest> {
  /**
   * NodeID is the node that the IP should be allocated for.
   *
   * @generated from field: string nodeID = 1;
   */
  nodeID: string;

  /**
   * Subnet is the subnet that the IP should be allocated from.
   *
   * @generated from field: string subnet = 2;
   */
  subnet: string;

  constructor(data?: PartialMessage<AllocateIPRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "v1.AllocateIPRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AllocateIPRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AllocateIPRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AllocateIPRequest;

  static equals(a: AllocateIPRequest | PlainMessage<AllocateIPRequest> | undefined, b: AllocateIPRequest | PlainMessage<AllocateIPRequest> | undefined): boolean;
}

/**
 * AllocatedIP is the message containing an allocated IP.
 *
 * @generated from message v1.AllocatedIP
 */
export declare class AllocatedIP extends Message<AllocatedIP> {
  /**
   * IP is the allocated IP. It should be returned in CIDR notation.
   *
   * @generated from field: string ip = 1;
   */
  ip: string;

  constructor(data?: PartialMessage<AllocatedIP>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "v1.AllocatedIP";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AllocatedIP;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AllocatedIP;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AllocatedIP;

  static equals(a: AllocatedIP | PlainMessage<AllocatedIP> | undefined, b: AllocatedIP | PlainMessage<AllocatedIP> | undefined): boolean;
}

/**
 * ReleaseIPRequest is the message containing an IP release request.
 *
 * @generated from message v1.ReleaseIPRequest
 */
export declare class ReleaseIPRequest extends Message<ReleaseIPRequest> {
  /**
   * NodeID is the node that the IP should be released for.
   *
   * @generated from field: string nodeID = 1;
   */
  nodeID: string;

  /**
   * IP is the IP that should be released.
   *
   * @generated from field: string ip = 2;
   */
  ip: string;

  constructor(data?: PartialMessage<ReleaseIPRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "v1.ReleaseIPRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ReleaseIPRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ReleaseIPRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ReleaseIPRequest;

  static equals(a: ReleaseIPRequest | PlainMessage<ReleaseIPRequest> | undefined, b: ReleaseIPRequest | PlainMessage<ReleaseIPRequest> | undefined): boolean;
}

