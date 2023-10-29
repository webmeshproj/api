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

import { proto3 } from "@bufbuild/protobuf";
import { InterfaceMetrics, MeshNode } from "./node_pb.js";
import { QueryRequest } from "./storage_query_pb.js";

/**
 * ConnectRequest is sent by an application to a daemon to establish a connection to a mesh.
 *
 * @generated from message v1.ConnectRequest
 */
export const ConnectRequest = proto3.makeMessageType(
  "v1.ConnectRequest",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "authMethod", kind: "enum", T: proto3.getEnumType(ConnectRequest_AuthMethod) },
    { no: 3, name: "authCredentials", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "scalar", T: 12 /* ScalarType.BYTES */} },
    { no: 4, name: "addrType", kind: "enum", T: proto3.getEnumType(ConnectRequest_AddrType) },
    { no: 5, name: "addrs", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 6, name: "networking", kind: "message", T: MeshConnNetworking },
    { no: 7, name: "services", kind: "message", T: MeshConnServices },
    { no: 8, name: "bootstrap", kind: "message", T: MeshConnBootstrap },
    { no: 9, name: "tls", kind: "message", T: MeshConnTLS },
  ],
);

/**
 * AddrType is the type of join addresses included in the request.
 *
 * @generated from enum v1.ConnectRequest.AddrType
 */
export const ConnectRequest_AddrType = proto3.makeEnum(
  "v1.ConnectRequest.AddrType",
  [
    {no: 0, name: "ADDR"},
    {no: 1, name: "MULTIADDR"},
    {no: 2, name: "RENDEZVOUS"},
  ],
);

/**
 * AuthMethod are types of RPC credentials to supply to the connection.
 *
 * @generated from enum v1.ConnectRequest.AuthMethod
 */
export const ConnectRequest_AuthMethod = proto3.makeEnum(
  "v1.ConnectRequest.AuthMethod",
  [
    {no: 0, name: "NO_AUTH"},
    {no: 1, name: "BASIC"},
    {no: 2, name: "LDAP"},
    {no: 3, name: "ID"},
    {no: 4, name: "MTLS"},
  ],
);

/**
 * AuthHeader is an enumeration of headers that coorespond to the AuthMethod.
 * They are used to pass authentication credentials to the daemon. Enums 
 * cannot be used as map keys, so their string values are used instead.
 *
 * @generated from enum v1.ConnectRequest.AuthHeader
 */
export const ConnectRequest_AuthHeader = proto3.makeEnum(
  "v1.ConnectRequest.AuthHeader",
  [
    {no: 0, name: "BASIC_USERNAME"},
    {no: 1, name: "BASIC_PASSWORD"},
    {no: 2, name: "LDAP_USERNAME"},
    {no: 3, name: "LDAP_PASSWORD"},
    {no: 4, name: "ADDRS_ENVELOPE"},
  ],
);

/**
 * MeshConnNetworking are configurations for networking on a mesh.
 *
 * @generated from message v1.MeshConnNetworking
 */
export const MeshConnNetworking = proto3.makeMessageType(
  "v1.MeshConnNetworking",
  () => [
    { no: 1, name: "useDNS", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ],
);

/**
 * MeshConnServices are configurations for exposing services to other nodes on a mesh.
 *
 * @generated from message v1.MeshConnServices
 */
export const MeshConnServices = proto3.makeMessageType(
  "v1.MeshConnServices",
  () => [
    { no: 1, name: "enabled", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 2, name: "public", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ],
);

/**
 * MeshConnBootstrap are configurations for bootstrapping a new mesh.
 *
 * @generated from message v1.MeshConnBootstrap
 */
export const MeshConnBootstrap = proto3.makeMessageType(
  "v1.MeshConnBootstrap",
  () => [
    { no: 1, name: "enabled", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 2, name: "domain", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "ipv4Network", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "rbacEnabled", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 5, name: "defaultNetworkACL", kind: "enum", T: proto3.getEnumType(MeshConnBootstrap_DefaultNetworkACL) },
  ],
);

/**
 * @generated from enum v1.MeshConnBootstrap.DefaultNetworkACL
 */
export const MeshConnBootstrap_DefaultNetworkACL = proto3.makeEnum(
  "v1.MeshConnBootstrap.DefaultNetworkACL",
  [
    {no: 0, name: "ACCEPT"},
    {no: 1, name: "DROP"},
  ],
);

/**
 * MeshhConnTLS are TLS configurations for a mesh connection.
 *
 * @generated from message v1.MeshConnTLS
 */
export const MeshConnTLS = proto3.makeMessageType(
  "v1.MeshConnTLS",
  () => [
    { no: 1, name: "enabled", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 2, name: "caCertData", kind: "scalar", T: 12 /* ScalarType.BYTES */ },
    { no: 3, name: "certData", kind: "scalar", T: 12 /* ScalarType.BYTES */ },
    { no: 4, name: "keyData", kind: "scalar", T: 12 /* ScalarType.BYTES */ },
    { no: 5, name: "verifyChainOnly", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 6, name: "skipVerify", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ],
);

/**
 * ConnectResponse is returned by the Connect RPC.
 *
 * @generated from message v1.ConnectResponse
 */
export const ConnectResponse = proto3.makeMessageType(
  "v1.ConnectResponse",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "nodeID", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "meshDomain", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "ipv4Address", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "ipv6Address", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "ipv4Network", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 7, name: "ipv6Network", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * DisconnectRequest is sent by an application to a daemon to disconnect
 * from a mesh. This message will eventually contain unique identifiers
 * for allowing the application to disconnect from a specific mesh.
 *
 * @generated from message v1.DisconnectRequest
 */
export const DisconnectRequest = proto3.makeMessageType(
  "v1.DisconnectRequest",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * DisconnectResponse is returned by the Disconnect RPC.
 *
 * @generated from message v1.DisconnectResponse
 */
export const DisconnectResponse = proto3.makeMessageType(
  "v1.DisconnectResponse",
  [],
);

/**
 * MetricsRequest is sent by the application to a daemon to retrieve interface
 * metrics for a mesh connection.
 *
 * @generated from message v1.MetricsRequest
 */
export const MetricsRequest = proto3.makeMessageType(
  "v1.MetricsRequest",
  () => [
    { no: 1, name: "ids", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ],
);

/**
 * MetricsResponse is a message containing interface metrics.
 *
 * @generated from message v1.MetricsResponse
 */
export const MetricsResponse = proto3.makeMessageType(
  "v1.MetricsResponse",
  () => [
    { no: 1, name: "interfaces", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "message", T: InterfaceMetrics} },
  ],
);

/**
 * StatusRequest is sent by the application to a daemon to retrieve the status
 * of a mesh connection.
 *
 * @generated from message v1.StatusRequest
 */
export const StatusRequest = proto3.makeMessageType(
  "v1.StatusRequest",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * StatusResponse is a message containing the status of the node.
 *
 * @generated from message v1.StatusResponse
 */
export const StatusResponse = proto3.makeMessageType(
  "v1.StatusResponse",
  () => [
    { no: 1, name: "connectionStatus", kind: "enum", T: proto3.getEnumType(StatusResponse_ConnectionStatus) },
    { no: 2, name: "node", kind: "message", T: MeshNode },
  ],
);

/**
 * @generated from enum v1.StatusResponse.ConnectionStatus
 */
export const StatusResponse_ConnectionStatus = proto3.makeEnum(
  "v1.StatusResponse.ConnectionStatus",
  [
    {no: 0, name: "DISCONNECTED"},
    {no: 1, name: "CONNECTING"},
    {no: 2, name: "CONNECTED"},
  ],
);

/**
 * AppQueryRequest is sent by the application to a daemon to query a mesh's storage.
 *
 * @generated from message v1.AppQueryRequest
 */
export const AppQueryRequest = proto3.makeMessageType(
  "v1.AppQueryRequest",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "query", kind: "message", T: QueryRequest },
  ],
);

