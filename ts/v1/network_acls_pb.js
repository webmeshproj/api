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
// @generated from file v1/network_acls.proto (package v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";

/**
 * ACLAction is the action to take when a request matches an ACL.
 *
 * @generated from enum v1.ACLAction
 */
export const ACLAction = proto3.makeEnum(
  "v1.ACLAction",
  [
    {no: 0, name: "ACTION_UNKNOWN"},
    {no: 1, name: "ACTION_ACCEPT"},
    {no: 2, name: "ACTION_DENY"},
  ],
);

/**
 * NetworkACL is a network ACL.
 *
 * @generated from message v1.NetworkACL
 */
export const NetworkACL = proto3.makeMessageType(
  "v1.NetworkACL",
  () => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "priority", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "action", kind: "enum", T: proto3.getEnumType(ACLAction) },
    { no: 4, name: "sourceNodes", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 5, name: "destinationNodes", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 6, name: "sourceCIDRs", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 7, name: "destinationCIDRs", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ],
);

/**
 * NetworkACLs is a list of network ACLs.
 *
 * @generated from message v1.NetworkACLs
 */
export const NetworkACLs = proto3.makeMessageType(
  "v1.NetworkACLs",
  () => [
    { no: 1, name: "items", kind: "message", T: NetworkACL, repeated: true },
  ],
);

/**
 * Route is a route that is broadcasted by one or more nodes.
 *
 * @generated from message v1.Route
 */
export const Route = proto3.makeMessageType(
  "v1.Route",
  () => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "node", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "destinationCIDRs", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 4, name: "nextHopNode", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * Routes is a list of routes.
 *
 * @generated from message v1.Routes
 */
export const Routes = proto3.makeMessageType(
  "v1.Routes",
  () => [
    { no: 1, name: "items", kind: "message", T: Route, repeated: true },
  ],
);

/**
 * NetworkAction is an action that can be performed on a network resource. It is used
 * by implementations to evaluate network ACLs.
 *
 * @generated from message v1.NetworkAction
 */
export const NetworkAction = proto3.makeMessageType(
  "v1.NetworkAction",
  () => [
    { no: 1, name: "srcNode", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "srcCIDR", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "dstNode", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "dstCIDR", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);
