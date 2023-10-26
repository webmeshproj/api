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
// @generated from file v1/registrar.proto (package v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3, Timestamp } from "@bufbuild/protobuf";

/**
 * RegisterRequest is the request object for the Register RPC.
 *
 * @generated from message v1.RegisterRequest
 */
export const RegisterRequest = proto3.makeMessageType(
  "v1.RegisterRequest",
  () => [
    { no: 1, name: "publicKey", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "alias", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "expiry", kind: "message", T: Timestamp },
  ],
);

/**
 * RegisterResponse is the response object for the Register RPC.
 *
 * @generated from message v1.RegisterResponse
 */
export const RegisterResponse = proto3.makeMessageType(
  "v1.RegisterResponse",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * LookupRequest is the request object for the Lookup RPC. One of the fields
 * must be provided.
 *
 * @generated from message v1.LookupRequest
 */
export const LookupRequest = proto3.makeMessageType(
  "v1.LookupRequest",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "publicKey", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "alias", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * LookupResponse is the response object for the Lookup RPC.
 *
 * @generated from message v1.LookupResponse
 */
export const LookupResponse = proto3.makeMessageType(
  "v1.LookupResponse",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "publicKey", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "alias", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

