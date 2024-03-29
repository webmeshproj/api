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
// @generated from file v1/registrar.proto (package v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, Timestamp } from "@bufbuild/protobuf";

/**
 * RegisterRequest is the request object for the Register RPC.
 *
 * @generated from message v1.RegisterRequest
 */
export class RegisterRequest extends Message<RegisterRequest> {
  /**
   * The encoded public key to register.
   *
   * @generated from field: string publicKey = 1;
   */
  publicKey = "";

  /**
   * An alias to associate with the public key. This can be used to lookup
   * the public key later.
   *
   * @generated from field: string alias = 2;
   */
  alias = "";

  /**
   * Expiry is the time at which the public key and its associated aliases
   * should be removed from the registrar. If not provided, a default value
   * of 1 day from the time of registration will be used.
   *
   * @generated from field: google.protobuf.Timestamp expiry = 3;
   */
  expiry?: Timestamp;

  constructor(data?: PartialMessage<RegisterRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "v1.RegisterRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "publicKey", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "alias", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "expiry", kind: "message", T: Timestamp },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RegisterRequest {
    return new RegisterRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RegisterRequest {
    return new RegisterRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RegisterRequest {
    return new RegisterRequest().fromJsonString(jsonString, options);
  }

  static equals(a: RegisterRequest | PlainMessage<RegisterRequest> | undefined, b: RegisterRequest | PlainMessage<RegisterRequest> | undefined): boolean {
    return proto3.util.equals(RegisterRequest, a, b);
  }
}

/**
 * RegisterResponse is the response object for the Register RPC.
 *
 * @generated from message v1.RegisterResponse
 */
export class RegisterResponse extends Message<RegisterResponse> {
  /**
   * ID of the public key that was registered.
   *
   * @generated from field: string id = 1;
   */
  id = "";

  constructor(data?: PartialMessage<RegisterResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "v1.RegisterResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RegisterResponse {
    return new RegisterResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RegisterResponse {
    return new RegisterResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RegisterResponse {
    return new RegisterResponse().fromJsonString(jsonString, options);
  }

  static equals(a: RegisterResponse | PlainMessage<RegisterResponse> | undefined, b: RegisterResponse | PlainMessage<RegisterResponse> | undefined): boolean {
    return proto3.util.equals(RegisterResponse, a, b);
  }
}

/**
 * LookupRequest is the request object for the Lookup RPC. One of the fields
 * must be provided.
 *
 * @generated from message v1.LookupRequest
 */
export class LookupRequest extends Message<LookupRequest> {
  /**
   * The ID derived from the public key to lookup.
   *
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * The public key to lookup.
   *
   * @generated from field: string publicKey = 2;
   */
  publicKey = "";

  /**
   * The alias of the public key to lookup.
   *
   * @generated from field: string alias = 3;
   */
  alias = "";

  constructor(data?: PartialMessage<LookupRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "v1.LookupRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "publicKey", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "alias", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): LookupRequest {
    return new LookupRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): LookupRequest {
    return new LookupRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): LookupRequest {
    return new LookupRequest().fromJsonString(jsonString, options);
  }

  static equals(a: LookupRequest | PlainMessage<LookupRequest> | undefined, b: LookupRequest | PlainMessage<LookupRequest> | undefined): boolean {
    return proto3.util.equals(LookupRequest, a, b);
  }
}

/**
 * LookupResponse is the response object for the Lookup RPC.
 *
 * @generated from message v1.LookupResponse
 */
export class LookupResponse extends Message<LookupResponse> {
  /**
   * The ID of the public key that was looked up.
   *
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * The encoded public key that was looked up.
   *
   * @generated from field: string publicKey = 2;
   */
  publicKey = "";

  /**
   * Any alias associated with the public key.
   *
   * @generated from field: string alias = 3;
   */
  alias = "";

  constructor(data?: PartialMessage<LookupResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "v1.LookupResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "publicKey", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "alias", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): LookupResponse {
    return new LookupResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): LookupResponse {
    return new LookupResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): LookupResponse {
    return new LookupResponse().fromJsonString(jsonString, options);
  }

  static equals(a: LookupResponse | PlainMessage<LookupResponse> | undefined, b: LookupResponse | PlainMessage<LookupResponse> | undefined): boolean {
    return proto3.util.equals(LookupResponse, a, b);
  }
}

