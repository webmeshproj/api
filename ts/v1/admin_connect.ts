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
// @generated from file v1/admin.proto (package v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { Group, Groups, Role, RoleBinding, RoleBindings, Roles } from "./rbac_pb.js";
import { Empty, MethodKind } from "@bufbuild/protobuf";
import { NetworkACL, NetworkACLs, Route, Routes } from "./network_acls_pb.js";
import { MeshEdge, MeshEdges } from "./mesh_pb.js";

/**
 * Admin is the service that provides cluster admin operations. Most methods 
 * require the leader to be contacted.
 *
 * @generated from service v1.Admin
 */
export const Admin = {
  typeName: "v1.Admin",
  methods: {
    /**
     * PutRole creates or updates a role.
     *
     * @generated from rpc v1.Admin.PutRole
     */
    putRole: {
      name: "PutRole",
      I: Role,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * DeleteRole deletes a role.
     *
     * @generated from rpc v1.Admin.DeleteRole
     */
    deleteRole: {
      name: "DeleteRole",
      I: Role,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * GetRole gets a role.
     *
     * @generated from rpc v1.Admin.GetRole
     */
    getRole: {
      name: "GetRole",
      I: Role,
      O: Role,
      kind: MethodKind.Unary,
    },
    /**
     * ListRoles gets all roles.
     *
     * @generated from rpc v1.Admin.ListRoles
     */
    listRoles: {
      name: "ListRoles",
      I: Empty,
      O: Roles,
      kind: MethodKind.Unary,
    },
    /**
     * PutRoleBinding creates or updates a role binding.
     *
     * @generated from rpc v1.Admin.PutRoleBinding
     */
    putRoleBinding: {
      name: "PutRoleBinding",
      I: RoleBinding,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * DeleteRoleBinding deletes a role binding.
     *
     * @generated from rpc v1.Admin.DeleteRoleBinding
     */
    deleteRoleBinding: {
      name: "DeleteRoleBinding",
      I: RoleBinding,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * GetRoleBinding gets a role binding.
     *
     * @generated from rpc v1.Admin.GetRoleBinding
     */
    getRoleBinding: {
      name: "GetRoleBinding",
      I: RoleBinding,
      O: RoleBinding,
      kind: MethodKind.Unary,
    },
    /**
     * ListRoleBindings gets all role bindings.
     *
     * @generated from rpc v1.Admin.ListRoleBindings
     */
    listRoleBindings: {
      name: "ListRoleBindings",
      I: Empty,
      O: RoleBindings,
      kind: MethodKind.Unary,
    },
    /**
     * PutGroup creates or updates a group.
     *
     * @generated from rpc v1.Admin.PutGroup
     */
    putGroup: {
      name: "PutGroup",
      I: Group,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * DeleteGroup deletes a group.
     *
     * @generated from rpc v1.Admin.DeleteGroup
     */
    deleteGroup: {
      name: "DeleteGroup",
      I: Group,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * GetGroup gets a group.
     *
     * @generated from rpc v1.Admin.GetGroup
     */
    getGroup: {
      name: "GetGroup",
      I: Group,
      O: Group,
      kind: MethodKind.Unary,
    },
    /**
     * ListGroups gets all groups.
     *
     * @generated from rpc v1.Admin.ListGroups
     */
    listGroups: {
      name: "ListGroups",
      I: Empty,
      O: Groups,
      kind: MethodKind.Unary,
    },
    /**
     * PutNetworkACL creates or updates a network ACL.
     *
     * @generated from rpc v1.Admin.PutNetworkACL
     */
    putNetworkACL: {
      name: "PutNetworkACL",
      I: NetworkACL,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * DeleteNetworkACL deletes a network ACL.
     *
     * @generated from rpc v1.Admin.DeleteNetworkACL
     */
    deleteNetworkACL: {
      name: "DeleteNetworkACL",
      I: NetworkACL,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * GetNetworkACL gets a network ACL.
     *
     * @generated from rpc v1.Admin.GetNetworkACL
     */
    getNetworkACL: {
      name: "GetNetworkACL",
      I: NetworkACL,
      O: NetworkACL,
      kind: MethodKind.Unary,
    },
    /**
     * ListNetworkACLs gets all network ACLs.
     *
     * @generated from rpc v1.Admin.ListNetworkACLs
     */
    listNetworkACLs: {
      name: "ListNetworkACLs",
      I: Empty,
      O: NetworkACLs,
      kind: MethodKind.Unary,
    },
    /**
     * PutRoute creates or updates a route.
     *
     * @generated from rpc v1.Admin.PutRoute
     */
    putRoute: {
      name: "PutRoute",
      I: Route,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * DeleteRoute deletes a route.
     *
     * @generated from rpc v1.Admin.DeleteRoute
     */
    deleteRoute: {
      name: "DeleteRoute",
      I: Route,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * GetRoute gets a route.
     *
     * @generated from rpc v1.Admin.GetRoute
     */
    getRoute: {
      name: "GetRoute",
      I: Route,
      O: Route,
      kind: MethodKind.Unary,
    },
    /**
     * ListRoutes gets all routes.
     *
     * @generated from rpc v1.Admin.ListRoutes
     */
    listRoutes: {
      name: "ListRoutes",
      I: Empty,
      O: Routes,
      kind: MethodKind.Unary,
    },
    /**
     * PutEdge creates or updates an edge between two nodes.
     *
     * @generated from rpc v1.Admin.PutEdge
     */
    putEdge: {
      name: "PutEdge",
      I: MeshEdge,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * DeleteEdge deletes an edge between two nodes.
     *
     * @generated from rpc v1.Admin.DeleteEdge
     */
    deleteEdge: {
      name: "DeleteEdge",
      I: MeshEdge,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * GetEdge gets an edge between two nodes.
     *
     * @generated from rpc v1.Admin.GetEdge
     */
    getEdge: {
      name: "GetEdge",
      I: MeshEdge,
      O: MeshEdge,
      kind: MethodKind.Unary,
    },
    /**
     * ListEdges gets all current edges.
     *
     * @generated from rpc v1.Admin.ListEdges
     */
    listEdges: {
      name: "ListEdges",
      I: Empty,
      O: MeshEdges,
      kind: MethodKind.Unary,
    },
  }
} as const;

