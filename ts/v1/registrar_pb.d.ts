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
// @generated from file v1/registrar.proto (package v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage, Timestamp } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * RegisterRequest is the request object for the Register RPC.
 *
 * @generated from message v1.RegisterRequest
 */
export declare class RegisterRequest extends Message<RegisterRequest> {
  /**
   * The encoded public key to register.
   *
   * @generated from field: string publicKey = 1;
   */
  publicKey: string;

  /**
   * An alias to associate with the public key. This can be used to lookup
   * the public key later.
   *
   * @generated from field: string alias = 2;
   */
  alias: string;

  /**
   * Expiry is the time at which the public key and its associated aliases
   * should be removed from the registrar. If not provided, a default value
   * of 1 day from the time of registration will be used.
   *
   * @generated from field: google.protobuf.Timestamp expiry = 3;
   */
  expiry?: Timestamp;

  constructor(data?: PartialMessage<RegisterRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "v1.RegisterRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RegisterRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RegisterRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RegisterRequest;

  static equals(a: RegisterRequest | PlainMessage<RegisterRequest> | undefined, b: RegisterRequest | PlainMessage<RegisterRequest> | undefined): boolean;
}

/**
 * RegisterResponse is the response object for the Register RPC.
 *
 * @generated from message v1.RegisterResponse
 */
export declare class RegisterResponse extends Message<RegisterResponse> {
  /**
   * ID of the public key that was registered.
   *
   * @generated from field: string id = 1;
   */
  id: string;

  constructor(data?: PartialMessage<RegisterResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "v1.RegisterResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RegisterResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RegisterResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RegisterResponse;

  static equals(a: RegisterResponse | PlainMessage<RegisterResponse> | undefined, b: RegisterResponse | PlainMessage<RegisterResponse> | undefined): boolean;
}

/**
 * LookupRequest is the request object for the Lookup RPC. One of the fields
 * must be provided.
 *
 * @generated from message v1.LookupRequest
 */
export declare class LookupRequest extends Message<LookupRequest> {
  /**
   * The ID derived from the public key to lookup.
   *
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * The public key to lookup.
   *
   * @generated from field: string publicKey = 2;
   */
  publicKey: string;

  /**
   * The alias of the public key to lookup.
   *
   * @generated from field: string alias = 3;
   */
  alias: string;

  constructor(data?: PartialMessage<LookupRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "v1.LookupRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): LookupRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): LookupRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): LookupRequest;

  static equals(a: LookupRequest | PlainMessage<LookupRequest> | undefined, b: LookupRequest | PlainMessage<LookupRequest> | undefined): boolean;
}

/**
 * LookupResponse is the response object for the Lookup RPC.
 *
 * @generated from message v1.LookupResponse
 */
export declare class LookupResponse extends Message<LookupResponse> {
  /**
   * The ID of the public key that was looked up.
   *
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * The encoded public key that was looked up.
   *
   * @generated from field: string publicKey = 2;
   */
  publicKey: string;

  /**
   * Any alias associated with the public key.
   *
   * @generated from field: string alias = 3;
   */
  alias: string;

  constructor(data?: PartialMessage<LookupResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "v1.LookupResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): LookupResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): LookupResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): LookupResponse;

  static equals(a: LookupResponse | PlainMessage<LookupResponse> | undefined, b: LookupResponse | PlainMessage<LookupResponse> | undefined): boolean;
}

