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
// @generated from file v1/raft.proto (package v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, Duration, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * RaftCommandType is the type of command being sent to the
 * Raft log.
 *
 * @generated from enum v1.RaftCommandType
 */
export declare enum RaftCommandType {
  /**
   * UNKNOWN is the unknown command type.
   *
   * @generated from enum value: UNKNOWN = 0;
   */
  UNKNOWN = 0,

  /**
   * PUT is the command for putting a key/value pair.
   *
   * @generated from enum value: PUT = 1;
   */
  PUT = 1,

  /**
   * DELETE is the command for deleting a key/value pair.
   *
   * @generated from enum value: DELETE = 2;
   */
  DELETE = 2,
}

/**
 * RaftLogEntry is the data of an entry in the Raft log.
 *
 * @generated from message v1.RaftLogEntry
 */
export declare class RaftLogEntry extends Message<RaftLogEntry> {
  /**
   * type is the type of the log entry.
   *
   * @generated from field: v1.RaftCommandType type = 1;
   */
  type: RaftCommandType;

  /**
   * key is the key of the log entry.
   *
   * @generated from field: bytes key = 2;
   */
  key: Uint8Array;

  /**
   * value is the value of the log entry.
   *
   * @generated from field: bytes value = 3;
   */
  value: Uint8Array;

  /**
   * ttl is the time to live of the log entry.
   *
   * @generated from field: google.protobuf.Duration ttl = 4;
   */
  ttl?: Duration;

  constructor(data?: PartialMessage<RaftLogEntry>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "v1.RaftLogEntry";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RaftLogEntry;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RaftLogEntry;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RaftLogEntry;

  static equals(a: RaftLogEntry | PlainMessage<RaftLogEntry> | undefined, b: RaftLogEntry | PlainMessage<RaftLogEntry> | undefined): boolean;
}

/**
 * RaftApplyResponse is the response to an apply request. It
 * contains the result of applying the log entry.
 *
 * @generated from message v1.RaftApplyResponse
 */
export declare class RaftApplyResponse extends Message<RaftApplyResponse> {
  /**
   * time is the total time it took to apply the log entry.
   *
   * @generated from field: string time = 1;
   */
  time: string;

  /**
   * error is an error that occurred during the apply.
   *
   * @generated from field: string error = 2;
   */
  error: string;

  constructor(data?: PartialMessage<RaftApplyResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "v1.RaftApplyResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RaftApplyResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RaftApplyResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RaftApplyResponse;

  static equals(a: RaftApplyResponse | PlainMessage<RaftApplyResponse> | undefined, b: RaftApplyResponse | PlainMessage<RaftApplyResponse> | undefined): boolean;
}

/**
 * RaftSnapshot is the data of a snapshot.
 *
 * @generated from message v1.RaftSnapshot
 */
export declare class RaftSnapshot extends Message<RaftSnapshot> {
  /**
   * @generated from field: repeated v1.RaftDataItem kv = 1;
   */
  kv: RaftDataItem[];

  constructor(data?: PartialMessage<RaftSnapshot>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "v1.RaftSnapshot";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RaftSnapshot;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RaftSnapshot;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RaftSnapshot;

  static equals(a: RaftSnapshot | PlainMessage<RaftSnapshot> | undefined, b: RaftSnapshot | PlainMessage<RaftSnapshot> | undefined): boolean;
}

/**
 * RaftDataItem represents a value in the Raft data store.
 *
 * @generated from message v1.RaftDataItem
 */
export declare class RaftDataItem extends Message<RaftDataItem> {
  /**
   * key is the key of the data item.
   *
   * @generated from field: bytes key = 1;
   */
  key: Uint8Array;

  /**
   * value is the value of the data item.
   *
   * @generated from field: bytes value = 2;
   */
  value: Uint8Array;

  /**
   * ttl is the time to live of the data item.
   *
   * @generated from field: google.protobuf.Duration ttl = 3;
   */
  ttl?: Duration;

  constructor(data?: PartialMessage<RaftDataItem>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "v1.RaftDataItem";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RaftDataItem;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RaftDataItem;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RaftDataItem;

  static equals(a: RaftDataItem | PlainMessage<RaftDataItem> | undefined, b: RaftDataItem | PlainMessage<RaftDataItem> | undefined): boolean;
}
