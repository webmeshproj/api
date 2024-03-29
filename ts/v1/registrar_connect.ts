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

// @generated by protoc-gen-connect-es v1.1.2 with parameter "target=ts"
// @generated from file v1/registrar.proto (package v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { LookupRequest, LookupResponse, RegisterRequest, RegisterResponse } from "./registrar_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * The registrar service can be used as a means of providing off-network storage of public 
 * keys and other information. This is useful for (and should only be used with) public-key
 * derived ID authentication where one might want to register simpler aliases for a public key. 
 * This service could eventually evolve into a full key-server, but for now it is just a simple
 * registrar.
 *
 * @generated from service v1.Registrar
 */
export const Registrar = {
  typeName: "v1.Registrar",
  methods: {
    /**
     * Register a public key with the registrar. An alias can be provided to make it easier 
     * to lookup the public key later. If the alias is already in use, the request will fail.
     * This method can be used to change the alias of a public key by providing the same
     * public key with a different alias.
     *
     * @generated from rpc v1.Registrar.Register
     */
    register: {
      name: "Register",
      I: RegisterRequest,
      O: RegisterResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Lookup a public key by ID or alias. If the ID is not found, the
     * request will fail.
     *
     * @generated from rpc v1.Registrar.Lookup
     */
    lookup: {
      name: "Lookup",
      I: LookupRequest,
      O: LookupResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

