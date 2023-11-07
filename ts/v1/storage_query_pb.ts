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
// @generated from file v1/storage_query.proto (package v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Duration, Message, proto3 } from "@bufbuild/protobuf";

/**
 * NetworkState represents the full network state as returned by
 * a network state query.
 *
 * @generated from message v1.NetworkState
 */
export class NetworkState extends Message<NetworkState> {
  /**
   * @generated from field: string networkV4 = 1;
   */
  networkV4 = "";

  /**
   * @generated from field: string networkV6 = 2;
   */
  networkV6 = "";

  /**
   * @generated from field: string domain = 3;
   */
  domain = "";

  constructor(data?: PartialMessage<NetworkState>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "v1.NetworkState";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "networkV4", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "networkV6", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "domain", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): NetworkState {
    return new NetworkState().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): NetworkState {
    return new NetworkState().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): NetworkState {
    return new NetworkState().fromJsonString(jsonString, options);
  }

  static equals(a: NetworkState | PlainMessage<NetworkState> | undefined, b: NetworkState | PlainMessage<NetworkState> | undefined): boolean {
    return proto3.util.equals(NetworkState, a, b);
  }
}

/**
 * QueryRequest is sent by the application to the node to query the mesh for
 * information.
 *
 * @generated from message v1.QueryRequest
 */
export class QueryRequest extends Message<QueryRequest> {
  /**
   * Command is the command of the query.
   *
   * @generated from field: v1.QueryRequest.QueryCommand command = 1;
   */
  command = QueryRequest_QueryCommand.GET;

  /**
   * Type is the type of resource for the query.
   *
   * @generated from field: v1.QueryRequest.QueryType type = 2;
   */
  type = QueryRequest_QueryType.VALUE;

  /**
   * Query is the string of the query. This follows the format of a comma-separted
   * label selector and is only applicable for certain queries. For get queries this
   * will usually be an ID. For list queries this will usually be one or more filters.
   * On put or delete queries, this should be an ID.
   *
   * @generated from field: string query = 3;
   */
  query = "";

  /**
   * Item is an item to put. This is only applicable for PUT queries. It should be a
   * protobuf-JSON encoded object of the given query type.
   *
   * @generated from field: bytes item = 4;
   */
  item = new Uint8Array(0);

  constructor(data?: PartialMessage<QueryRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "v1.QueryRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "command", kind: "enum", T: proto3.getEnumType(QueryRequest_QueryCommand) },
    { no: 2, name: "type", kind: "enum", T: proto3.getEnumType(QueryRequest_QueryType) },
    { no: 3, name: "query", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "item", kind: "scalar", T: 12 /* ScalarType.BYTES */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): QueryRequest {
    return new QueryRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): QueryRequest {
    return new QueryRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): QueryRequest {
    return new QueryRequest().fromJsonString(jsonString, options);
  }

  static equals(a: QueryRequest | PlainMessage<QueryRequest> | undefined, b: QueryRequest | PlainMessage<QueryRequest> | undefined): boolean {
    return proto3.util.equals(QueryRequest, a, b);
  }
}

/**
 * QueryCommand is the type of the query.
 *
 * @generated from enum v1.QueryRequest.QueryCommand
 */
export enum QueryRequest_QueryCommand {
  /**
   * GET is the command to get a value.
   *
   * @generated from enum value: GET = 0;
   */
  GET = 0,

  /**
   * LIST is the command to list keys with an optional prefix.
   *
   * @generated from enum value: LIST = 1;
   */
  LIST = 1,

  /**
   * PUT is the command to put a value.
   *
   * @generated from enum value: PUT = 2;
   */
  PUT = 2,

  /**
   * DELETE is the command to delete a value.
   *
   * @generated from enum value: DELETE = 3;
   */
  DELETE = 3,
}
// Retrieve enum metadata with: proto3.getEnumType(QueryRequest_QueryCommand)
proto3.util.setEnumType(QueryRequest_QueryCommand, "v1.QueryRequest.QueryCommand", [
  { no: 0, name: "GET" },
  { no: 1, name: "LIST" },
  { no: 2, name: "PUT" },
  { no: 3, name: "DELETE" },
]);

/**
 * QueryType is the type of object being queried.
 *
 * @generated from enum v1.QueryRequest.QueryType
 */
export enum QueryRequest_QueryType {
  /**
   * VALUE represents a raw value query at a supplied key.
   *
   * @generated from enum value: VALUE = 0;
   */
  VALUE = 0,

  /**
   * KEYS is the type for querying keys.
   *
   * @generated from enum value: KEYS = 1;
   */
  KEYS = 1,

  /**
   * PEERS is the type for querying peers.
   *
   * @generated from enum value: PEERS = 2;
   */
  PEERS = 2,

  /**
   * EDGES is the type for querying edges.
   *
   * @generated from enum value: EDGES = 3;
   */
  EDGES = 3,

  /**
   * ROUTES is the type for querying routes.
   *
   * @generated from enum value: ROUTES = 4;
   */
  ROUTES = 4,

  /**
   * ACLS is the type for querying ACLs.
   *
   * @generated from enum value: ACLS = 5;
   */
  ACLS = 5,

  /**
   * ROLES is the type for querying roles.
   *
   * @generated from enum value: ROLES = 6;
   */
  ROLES = 6,

  /**
   * ROLEBINDINGS is the type for querying role bindings.
   *
   * @generated from enum value: ROLEBINDINGS = 7;
   */
  ROLEBINDINGS = 7,

  /**
   * GROUPS is the type for querying groups.
   *
   * @generated from enum value: GROUPS = 8;
   */
  GROUPS = 8,

  /**
   * NETWORK_STATE is the type for querying network configuration.
   *
   * @generated from enum value: NETWORK_STATE = 9;
   */
  NETWORK_STATE = 9,

  /**
   * RBAC_STATE is the type for querying RBAC configuration.
   * This will return a single item of true or false.
   *
   * @generated from enum value: RBAC_STATE = 10;
   */
  RBAC_STATE = 10,
}
// Retrieve enum metadata with: proto3.getEnumType(QueryRequest_QueryType)
proto3.util.setEnumType(QueryRequest_QueryType, "v1.QueryRequest.QueryType", [
  { no: 0, name: "VALUE" },
  { no: 1, name: "KEYS" },
  { no: 2, name: "PEERS" },
  { no: 3, name: "EDGES" },
  { no: 4, name: "ROUTES" },
  { no: 5, name: "ACLS" },
  { no: 6, name: "ROLES" },
  { no: 7, name: "ROLEBINDINGS" },
  { no: 8, name: "GROUPS" },
  { no: 9, name: "NETWORK_STATE" },
  { no: 10, name: "RBAC_STATE" },
]);

/**
 * QueryResponse is the message containing a mesh query result.
 *
 * @generated from message v1.QueryResponse
 */
export class QueryResponse extends Message<QueryResponse> {
  /**
   * Items contain the results of the query. These will be protobuf
   * json-encoded objects of the given query type.
   *
   * @generated from field: repeated bytes items = 1;
   */
  items: Uint8Array[] = [];

  /**
   * Error is an error that happened during the query. This will always
   * be populated on errors, but single-flight queries will return
   * a coded error instead.
   *
   * @generated from field: string error = 4;
   */
  error = "";

  constructor(data?: PartialMessage<QueryResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "v1.QueryResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "items", kind: "scalar", T: 12 /* ScalarType.BYTES */, repeated: true },
    { no: 4, name: "error", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): QueryResponse {
    return new QueryResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): QueryResponse {
    return new QueryResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): QueryResponse {
    return new QueryResponse().fromJsonString(jsonString, options);
  }

  static equals(a: QueryResponse | PlainMessage<QueryResponse> | undefined, b: QueryResponse | PlainMessage<QueryResponse> | undefined): boolean {
    return proto3.util.equals(QueryResponse, a, b);
  }
}

/**
 * SubscribeRequest is sent by the application to the node to subscribe to
 * events. This currently only supports database events.
 *
 * @generated from message v1.SubscribeRequest
 */
export class SubscribeRequest extends Message<SubscribeRequest> {
  /**
   * Prefix is the prefix of the events to subscribe to.
   *
   * @generated from field: bytes prefix = 1;
   */
  prefix = new Uint8Array(0);

  constructor(data?: PartialMessage<SubscribeRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "v1.SubscribeRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "prefix", kind: "scalar", T: 12 /* ScalarType.BYTES */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SubscribeRequest {
    return new SubscribeRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SubscribeRequest {
    return new SubscribeRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SubscribeRequest {
    return new SubscribeRequest().fromJsonString(jsonString, options);
  }

  static equals(a: SubscribeRequest | PlainMessage<SubscribeRequest> | undefined, b: SubscribeRequest | PlainMessage<SubscribeRequest> | undefined): boolean {
    return proto3.util.equals(SubscribeRequest, a, b);
  }
}

/**
 * SubscriptionEvent is a message containing a subscription event.
 *
 * @generated from message v1.SubscriptionEvent
 */
export class SubscriptionEvent extends Message<SubscriptionEvent> {
  /**
   * Key is the key of the event.
   *
   * @generated from field: bytes key = 1;
   */
  key = new Uint8Array(0);

  /**
   * Value is the value of the event. This will be the raw value of the key.
   *
   * @generated from field: bytes value = 2;
   */
  value = new Uint8Array(0);

  constructor(data?: PartialMessage<SubscriptionEvent>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "v1.SubscriptionEvent";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "key", kind: "scalar", T: 12 /* ScalarType.BYTES */ },
    { no: 2, name: "value", kind: "scalar", T: 12 /* ScalarType.BYTES */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SubscriptionEvent {
    return new SubscriptionEvent().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SubscriptionEvent {
    return new SubscriptionEvent().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SubscriptionEvent {
    return new SubscriptionEvent().fromJsonString(jsonString, options);
  }

  static equals(a: SubscriptionEvent | PlainMessage<SubscriptionEvent> | undefined, b: SubscriptionEvent | PlainMessage<SubscriptionEvent> | undefined): boolean {
    return proto3.util.equals(SubscriptionEvent, a, b);
  }
}

/**
 * PublishRequest is sent by the application to the node to publish events.
 * This currently only supports database events.
 *
 * @generated from message v1.PublishRequest
 */
export class PublishRequest extends Message<PublishRequest> {
  /**
   * Key is the key of the event.
   *
   * @generated from field: bytes key = 1;
   */
  key = new Uint8Array(0);

  /**
   * Value is the value of the event. This will be the raw value of the key.
   *
   * @generated from field: bytes value = 2;
   */
  value = new Uint8Array(0);

  /**
   * TTL is the time for the event to live in the database.
   *
   * @generated from field: google.protobuf.Duration ttl = 3;
   */
  ttl?: Duration;

  constructor(data?: PartialMessage<PublishRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "v1.PublishRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "key", kind: "scalar", T: 12 /* ScalarType.BYTES */ },
    { no: 2, name: "value", kind: "scalar", T: 12 /* ScalarType.BYTES */ },
    { no: 3, name: "ttl", kind: "message", T: Duration },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): PublishRequest {
    return new PublishRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): PublishRequest {
    return new PublishRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): PublishRequest {
    return new PublishRequest().fromJsonString(jsonString, options);
  }

  static equals(a: PublishRequest | PlainMessage<PublishRequest> | undefined, b: PublishRequest | PlainMessage<PublishRequest> | undefined): boolean {
    return proto3.util.equals(PublishRequest, a, b);
  }
}

/**
 * PublishResponse is the response to a publish request. This is currently
 * empty.
 *
 * @generated from message v1.PublishResponse
 */
export class PublishResponse extends Message<PublishResponse> {
  constructor(data?: PartialMessage<PublishResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "v1.PublishResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): PublishResponse {
    return new PublishResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): PublishResponse {
    return new PublishResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): PublishResponse {
    return new PublishResponse().fromJsonString(jsonString, options);
  }

  static equals(a: PublishResponse | PlainMessage<PublishResponse> | undefined, b: PublishResponse | PlainMessage<PublishResponse> | undefined): boolean {
    return proto3.util.equals(PublishResponse, a, b);
  }
}

