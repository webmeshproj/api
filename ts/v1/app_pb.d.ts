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
// @generated from file v1/app.proto (package v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import type { InterfaceMetrics, MeshNode } from "./node_pb.js";

/**
 * ConnectRequest is sent by an application to a daemon to establish a connection to a mesh.
 *
 * @generated from message v1.ConnectRequest
 */
export declare class ConnectRequest extends Message<ConnectRequest> {
  /**
   * ID is the unique identifier of this connection. If not provided
   * one will be generated.
   *
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * AuthMethod is the type of authentication to use.
   *
   * @generated from field: v1.ConnectRequest.AuthMethod authMethod = 2;
   */
  authMethod: ConnectRequest_AuthMethod;

  /**
   * AuthCredentials are additional credentials as required by the authType. 
   *
   * @generated from field: map<string, bytes> authCredentials = 3;
   */
  authCredentials: { [key: string]: Uint8Array };

  /**
   * AddrType is the type of join addresses in the addrs list.
   *
   * @generated from field: v1.ConnectRequest.AddrType addrType = 4;
   */
  addrType: ConnectRequest_AddrType;

  /**
   * Addrs are the join addresses to use to connect to the mesh.
   *
   * @generated from field: repeated string addrs = 5;
   */
  addrs: string[];

  /**
   * Networking is the networking configuration to use.
   *
   * @generated from field: v1.MeshConnNetworking networking = 6;
   */
  networking?: MeshConnNetworking;

  /**
   * Services are the services to expose to other nodes on the mesh.
   *
   * @generated from field: v1.MeshConnServices services = 7;
   */
  services?: MeshConnServices;

  /**
   * Bootstrap are options for bootstrapping a new mesh.
   *
   * @generated from field: v1.MeshConnBootstrap bootstrap = 8;
   */
  bootstrap?: MeshConnBootstrap;

  /**
   * TLS are TLS configurations for the mesh connection.
   *
   * @generated from field: v1.MeshConnTLS tls = 9;
   */
  tls?: MeshConnTLS;

  constructor(data?: PartialMessage<ConnectRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "v1.ConnectRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ConnectRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ConnectRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ConnectRequest;

  static equals(a: ConnectRequest | PlainMessage<ConnectRequest> | undefined, b: ConnectRequest | PlainMessage<ConnectRequest> | undefined): boolean;
}

/**
 * AddrType is the type of join addresses included in the request.
 *
 * @generated from enum v1.ConnectRequest.AddrType
 */
export declare enum ConnectRequest_AddrType {
  /**
   * ADDR is used to join a mesh using an IP or DNS address.
   *
   * @generated from enum value: ADDR = 0;
   */
  ADDR = 0,

  /**
   * MULTIADDR is used to join a mesh using a multiaddr.
   *
   * @generated from enum value: MULTIADDR = 1;
   */
  MULTIADDR = 1,

  /**
   * RENDEZVOUS is used to join a mesh using a rendezvous string.
   *
   * @generated from enum value: RENDEZVOUS = 2;
   */
  RENDEZVOUS = 2,
}

/**
 * AuthMethod are types of RPC credentials to supply to the connection.
 *
 * @generated from enum v1.ConnectRequest.AuthMethod
 */
export declare enum ConnectRequest_AuthMethod {
  /**
   * NO_AUTH is used to indicate that no authentication is required.
   *
   * @generated from enum value: NO_AUTH = 0;
   */
  NO_AUTH = 0,

  /**
   * BASIC is used to indicate that basic authentication is required.
   *
   * @generated from enum value: BASIC = 1;
   */
  BASIC = 1,

  /**
   * LDAP is used to indicate that LDAP authentication is required.
   *
   * @generated from enum value: LDAP = 2;
   */
  LDAP = 2,

  /**
   * ID is used to indicate that an identity is required.
   *
   * @generated from enum value: ID = 3;
   */
  ID = 3,

  /**
   * MTLS is used to indicate that mutual TLS authentication is required.
   * The TLS object should be used to configure the TLS connection.
   *
   * @generated from enum value: MTLS = 4;
   */
  MTLS = 4,
}

/**
 * AuthHeader is an enumeration of headers that coorespond to the AuthMethod.
 * They are used to pass authentication credentials to the daemon. Enums 
 * cannot be used as map keys, so their string values are used instead.
 *
 * @generated from enum v1.ConnectRequest.AuthHeader
 */
export declare enum ConnectRequest_AuthHeader {
  /**
   * BASIC_USERNAME is the username for basic authentication.
   *
   * @generated from enum value: BASIC_USERNAME = 0;
   */
  BASIC_USERNAME = 0,

  /**
   * BASIC_PASSWORD is the password for basic authentication.
   *
   * @generated from enum value: BASIC_PASSWORD = 1;
   */
  BASIC_PASSWORD = 1,

  /**
   * LDAP_USERNAME is the username for LDAP authentication.
   *
   * @generated from enum value: LDAP_USERNAME = 2;
   */
  LDAP_USERNAME = 2,

  /**
   * LDAP_PASSWORD is the password for LDAP authentication.
   *
   * @generated from enum value: LDAP_PASSWORD = 3;
   */
  LDAP_PASSWORD = 3,

  /**
   * ADDRS_ENVELOPE is the header for a signed envelope containing
   * the join addresses to use to connect to the mesh.
   *
   * @generated from enum value: ADDRS_ENVELOPE = 4;
   */
  ADDRS_ENVELOPE = 4,
}

/**
 * MeshConnNetworking are configurations for networking on a mesh.
 *
 * @generated from message v1.MeshConnNetworking
 */
export declare class MeshConnNetworking extends Message<MeshConnNetworking> {
  /**
   * UseDNS indicates whether or not to use the DNS servers of the mesh.
   *
   * @generated from field: bool useDNS = 1;
   */
  useDNS: boolean;

  constructor(data?: PartialMessage<MeshConnNetworking>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "v1.MeshConnNetworking";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MeshConnNetworking;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MeshConnNetworking;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MeshConnNetworking;

  static equals(a: MeshConnNetworking | PlainMessage<MeshConnNetworking> | undefined, b: MeshConnNetworking | PlainMessage<MeshConnNetworking> | undefined): boolean;
}

/**
 * MeshConnServices are configurations for exposing services to other nodes on a mesh.
 *
 * @generated from message v1.MeshConnServices
 */
export declare class MeshConnServices extends Message<MeshConnServices> {
  /**
   * Enabled indicates whether or not to expose services to other nodes.
   *
   * @generated from field: bool enabled = 1;
   */
  enabled: boolean;

  constructor(data?: PartialMessage<MeshConnServices>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "v1.MeshConnServices";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MeshConnServices;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MeshConnServices;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MeshConnServices;

  static equals(a: MeshConnServices | PlainMessage<MeshConnServices> | undefined, b: MeshConnServices | PlainMessage<MeshConnServices> | undefined): boolean;
}

/**
 * MeshConnBootstrap are configurations for bootstrapping a new mesh.
 *
 * @generated from message v1.MeshConnBootstrap
 */
export declare class MeshConnBootstrap extends Message<MeshConnBootstrap> {
  /**
   * Enabled indicates whether or not to bootstrap a new mesh.
   *
   * @generated from field: bool enabled = 1;
   */
  enabled: boolean;

  constructor(data?: PartialMessage<MeshConnBootstrap>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "v1.MeshConnBootstrap";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MeshConnBootstrap;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MeshConnBootstrap;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MeshConnBootstrap;

  static equals(a: MeshConnBootstrap | PlainMessage<MeshConnBootstrap> | undefined, b: MeshConnBootstrap | PlainMessage<MeshConnBootstrap> | undefined): boolean;
}

/**
 * MeshhConnTLS are TLS configurations for a mesh connection.
 *
 * @generated from message v1.MeshConnTLS
 */
export declare class MeshConnTLS extends Message<MeshConnTLS> {
  /**
   * Enabled indicates whether or not to use TLS.
   *
   * @generated from field: bool enabled = 1;
   */
  enabled: boolean;

  /**
   * CACert is a PEM-encoded CA certificate to use for TLS.
   *
   * @generated from field: bytes caCertData = 2;
   */
  caCertData: Uint8Array;

  /**
   * CertData is a PEM-encoded certificate to use for TLS.
   *
   * @generated from field: bytes certData = 3;
   */
  certData: Uint8Array;

  /**
   * KeyData is a PEM-encoded private key to use for TLS.
   *
   * @generated from field: bytes keyData = 4;
   */
  keyData: Uint8Array;

  /**
   * VerifyChainOnly indicates whether or not to only verify the
   * certificate chain.
   *
   * @generated from field: bool verifyChainOnly = 5;
   */
  verifyChainOnly: boolean;

  /**
   * SkipVerify indicates whether or not to skip verification of the
   * server certificate.
   *
   * @generated from field: bool skipVerify = 6;
   */
  skipVerify: boolean;

  constructor(data?: PartialMessage<MeshConnTLS>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "v1.MeshConnTLS";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MeshConnTLS;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MeshConnTLS;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MeshConnTLS;

  static equals(a: MeshConnTLS | PlainMessage<MeshConnTLS> | undefined, b: MeshConnTLS | PlainMessage<MeshConnTLS> | undefined): boolean;
}

/**
 * ConnectResponse is returned by the Connect RPC.
 *
 * @generated from message v1.ConnectResponse
 */
export declare class ConnectResponse extends Message<ConnectResponse> {
  /**
   * ID is the unique identifier of this connection.
   *
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * Node id is the unique identifier of the node.
   *
   * @generated from field: string nodeID = 2;
   */
  nodeID: string;

  /**
   * Mesh domain is the domain of the mesh.
   *
   * @generated from field: string meshDomain = 3;
   */
  meshDomain: string;

  /**
   * IPv4 is the IPv4 address of the node.
   *
   * @generated from field: string ipv4 = 4;
   */
  ipv4: string;

  /**
   * IPv6 is the IPv6 address of the node.
   *
   * @generated from field: string ipv6 = 5;
   */
  ipv6: string;

  constructor(data?: PartialMessage<ConnectResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "v1.ConnectResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ConnectResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ConnectResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ConnectResponse;

  static equals(a: ConnectResponse | PlainMessage<ConnectResponse> | undefined, b: ConnectResponse | PlainMessage<ConnectResponse> | undefined): boolean;
}

/**
 * DisconnectRequest is sent by an application to a daemon to disconnect
 * from a mesh. This message will eventually contain unique identifiers
 * for allowing the application to disconnect from a specific mesh.
 *
 * @generated from message v1.DisconnectRequest
 */
export declare class DisconnectRequest extends Message<DisconnectRequest> {
  /**
   * ID is the unique identifier of this connection.
   *
   * @generated from field: string id = 1;
   */
  id: string;

  constructor(data?: PartialMessage<DisconnectRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "v1.DisconnectRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DisconnectRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DisconnectRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DisconnectRequest;

  static equals(a: DisconnectRequest | PlainMessage<DisconnectRequest> | undefined, b: DisconnectRequest | PlainMessage<DisconnectRequest> | undefined): boolean;
}

/**
 * DisconnectResponse is returned by the Disconnect RPC.
 *
 * @generated from message v1.DisconnectResponse
 */
export declare class DisconnectResponse extends Message<DisconnectResponse> {
  constructor(data?: PartialMessage<DisconnectResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "v1.DisconnectResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DisconnectResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DisconnectResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DisconnectResponse;

  static equals(a: DisconnectResponse | PlainMessage<DisconnectResponse> | undefined, b: DisconnectResponse | PlainMessage<DisconnectResponse> | undefined): boolean;
}

/**
 * MetricsRequest is sent by the application to a daemon to retrieve interface
 * metrics for a mesh connection.
 *
 * @generated from message v1.MetricsRequest
 */
export declare class MetricsRequest extends Message<MetricsRequest> {
  /**
   * IDs are the unique identifiers of the connections to retrieve metrics for.
   * If not provided, metrics for all connections will be returned.
   *
   * @generated from field: repeated string ids = 1;
   */
  ids: string[];

  constructor(data?: PartialMessage<MetricsRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "v1.MetricsRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MetricsRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MetricsRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MetricsRequest;

  static equals(a: MetricsRequest | PlainMessage<MetricsRequest> | undefined, b: MetricsRequest | PlainMessage<MetricsRequest> | undefined): boolean;
}

/**
 * MetricsResponse is a message containing interface metrics.
 *
 * @generated from message v1.MetricsResponse
 */
export declare class MetricsResponse extends Message<MetricsResponse> {
  /**
   * Interfaces is a map of network IDs to their interface metrics.
   *
   * @generated from field: map<string, v1.InterfaceMetrics> interfaces = 1;
   */
  interfaces: { [key: string]: InterfaceMetrics };

  constructor(data?: PartialMessage<MetricsResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "v1.MetricsResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MetricsResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MetricsResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MetricsResponse;

  static equals(a: MetricsResponse | PlainMessage<MetricsResponse> | undefined, b: MetricsResponse | PlainMessage<MetricsResponse> | undefined): boolean;
}

/**
 * StatusRequest is sent by the application to a daemon to retrieve the status
 * of a mesh connection.
 *
 * @generated from message v1.StatusRequest
 */
export declare class StatusRequest extends Message<StatusRequest> {
  /**
   * ID is the unique identifier of this connection.
   *
   * @generated from field: string id = 1;
   */
  id: string;

  constructor(data?: PartialMessage<StatusRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "v1.StatusRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): StatusRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): StatusRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): StatusRequest;

  static equals(a: StatusRequest | PlainMessage<StatusRequest> | undefined, b: StatusRequest | PlainMessage<StatusRequest> | undefined): boolean;
}

/**
 * StatusResponse is a message containing the status of the node.
 *
 * @generated from message v1.StatusResponse
 */
export declare class StatusResponse extends Message<StatusResponse> {
  /**
   * ConnectionStatus is the status of the connection.
   *
   * @generated from field: v1.StatusResponse.ConnectionStatus connectionStatus = 1;
   */
  connectionStatus: StatusResponse_ConnectionStatus;

  /**
   * Node is the node status. This is only populated if the node is connected.
   *
   * @generated from field: v1.MeshNode node = 2;
   */
  node?: MeshNode;

  constructor(data?: PartialMessage<StatusResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "v1.StatusResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): StatusResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): StatusResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): StatusResponse;

  static equals(a: StatusResponse | PlainMessage<StatusResponse> | undefined, b: StatusResponse | PlainMessage<StatusResponse> | undefined): boolean;
}

/**
 * @generated from enum v1.StatusResponse.ConnectionStatus
 */
export declare enum StatusResponse_ConnectionStatus {
  /**
   * DISCONNECTED indicates that the node is not connected to a mesh.
   *
   * @generated from enum value: DISCONNECTED = 0;
   */
  DISCONNECTED = 0,

  /**
   * CONNECTING indicates that the node is in the process of connecting to a mesh.
   *
   * @generated from enum value: CONNECTING = 1;
   */
  CONNECTING = 1,

  /**
   * CONNECTED indicates that the node is connected to a mesh.
   *
   * @generated from enum value: CONNECTED = 2;
   */
  CONNECTED = 2,
}

