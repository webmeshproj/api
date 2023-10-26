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

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * ACLAction is the action to take when a request matches an ACL.
 *
 * @generated from enum v1.ACLAction
 */
export declare enum ACLAction {
  /**
   * ACTION_UNKNOWN is the default action for ACLs. It is synonymous with ACTION_DENY.
   *
   * @generated from enum value: ACTION_UNKNOWN = 0;
   */
  ACTION_UNKNOWN = 0,

  /**
   * ACTION_ACCEPT allows the request to proceed.
   *
   * @generated from enum value: ACTION_ACCEPT = 1;
   */
  ACTION_ACCEPT = 1,

  /**
   * ACTION_DENY denies the request.
   *
   * @generated from enum value: ACTION_DENY = 2;
   */
  ACTION_DENY = 2,
}

/**
 * NetworkACL is a network ACL.
 *
 * @generated from message v1.NetworkACL
 */
export declare class NetworkACL extends Message<NetworkACL> {
  /**
   * name is the name of the ACL.
   *
   * @generated from field: string name = 1;
   */
  name: string;

  /**
   * priority is the priority of the ACL. ACLs with higher priority are evaluated first.
   *
   * @generated from field: int32 priority = 2;
   */
  priority: number;

  /**
   * action is the action to take when a request matches the ACL.
   *
   * @generated from field: v1.ACLAction action = 3;
   */
  action: ACLAction;

  /**
   * source_nodes is a list of source nodes to match against. If empty, all nodes are matched. Groups
   * can be specified with the prefix "group:". If one or more of the nodes is '*', all nodes are matched.
   *
   * @generated from field: repeated string sourceNodes = 4;
   */
  sourceNodes: string[];

  /**
   * destination_nodes is a list of destination nodes to match against. If empty, all nodes are matched.
   * Groups can be specified with the prefix "group:". If one or more of the nodes is '*', all nodes are matched.
   *
   * @generated from field: repeated string destinationNodes = 5;
   */
  destinationNodes: string[];

  /**
   * source_cidrs is a list of source CIDRs to match against. If empty, all CIDRs are matched.
   * If one or more of the CIDRs is '*', all CIDRs are matched.
   *
   * @generated from field: repeated string sourceCIDRs = 6;
   */
  sourceCIDRs: string[];

  /**
   * destination_cidrs is a list of destination CIDRs to match against. If empty, all CIDRs are matched.
   * If one or more of the CIDRs is '*', all CIDRs are matched.
   *
   * // protocols is a list of protocols to match against. If empty, all protocols are matched.
   * // Protocols can be specified by name or number.
   * repeated string protocols = 8;
   * // ports is a list of ports to match against. If empty, all ports are matched.
   * repeated uint32 ports = 9;
   *
   * @generated from field: repeated string destinationCIDRs = 7;
   */
  destinationCIDRs: string[];

  constructor(data?: PartialMessage<NetworkACL>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "v1.NetworkACL";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): NetworkACL;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): NetworkACL;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): NetworkACL;

  static equals(a: NetworkACL | PlainMessage<NetworkACL> | undefined, b: NetworkACL | PlainMessage<NetworkACL> | undefined): boolean;
}

/**
 * NetworkACLs is a list of network ACLs.
 *
 * @generated from message v1.NetworkACLs
 */
export declare class NetworkACLs extends Message<NetworkACLs> {
  /**
   * items is the list of network ACLs.
   *
   * @generated from field: repeated v1.NetworkACL items = 1;
   */
  items: NetworkACL[];

  constructor(data?: PartialMessage<NetworkACLs>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "v1.NetworkACLs";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): NetworkACLs;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): NetworkACLs;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): NetworkACLs;

  static equals(a: NetworkACLs | PlainMessage<NetworkACLs> | undefined, b: NetworkACLs | PlainMessage<NetworkACLs> | undefined): boolean;
}

/**
 * Route is a route that is broadcasted by one or more nodes.
 *
 * @generated from message v1.Route
 */
export declare class Route extends Message<Route> {
  /**
   * name is the name of the route.
   *
   * @generated from field: string name = 1;
   */
  name: string;

  /**
   * node is the node that broadcasts the route. A group can be specified with the prefix "group:".
   *
   * @generated from field: string node = 2;
   */
  node: string;

  /**
   * destination_cidrs are the destination CIDRs of the route.
   *
   * @generated from field: repeated string destinationCIDRs = 3;
   */
  destinationCIDRs: string[];

  /**
   * nextHopNode is an optional node that is used as the next hop for the route.
   *
   * @generated from field: string nextHopNode = 4;
   */
  nextHopNode: string;

  constructor(data?: PartialMessage<Route>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "v1.Route";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Route;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Route;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Route;

  static equals(a: Route | PlainMessage<Route> | undefined, b: Route | PlainMessage<Route> | undefined): boolean;
}

/**
 * Routes is a list of routes.
 *
 * @generated from message v1.Routes
 */
export declare class Routes extends Message<Routes> {
  /**
   * items is the list of routes.
   *
   * @generated from field: repeated v1.Route items = 1;
   */
  items: Route[];

  constructor(data?: PartialMessage<Routes>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "v1.Routes";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Routes;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Routes;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Routes;

  static equals(a: Routes | PlainMessage<Routes> | undefined, b: Routes | PlainMessage<Routes> | undefined): boolean;
}

/**
 * NetworkAction is an action that can be performed on a network resource. It is used
 * by implementations to evaluate network ACLs.
 *
 * @generated from message v1.NetworkAction
 */
export declare class NetworkAction extends Message<NetworkAction> {
  /**
   * src_node is the source node of the action.
   *
   * @generated from field: string srcNode = 1;
   */
  srcNode: string;

  /**
   * src_cidr is the source CIDR of the action.
   *
   * @generated from field: string srcCIDR = 2;
   */
  srcCIDR: string;

  /**
   * dst_node is the destination node of the action.
   *
   * @generated from field: string dstNode = 3;
   */
  dstNode: string;

  /**
   * dst_cidr is the destination CIDR of the action.
   *
   * @generated from field: string dstCIDR = 4;
   */
  dstCIDR: string;

  constructor(data?: PartialMessage<NetworkAction>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "v1.NetworkAction";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): NetworkAction;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): NetworkAction;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): NetworkAction;

  static equals(a: NetworkAction | PlainMessage<NetworkAction> | undefined, b: NetworkAction | PlainMessage<NetworkAction> | undefined): boolean;
}

