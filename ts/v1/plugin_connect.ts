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
// @generated from file v1/plugin.proto (package v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { Empty, MethodKind } from "@bufbuild/protobuf";
import { AllocatedIP, AllocateIPRequest, AuthenticationRequest, AuthenticationResponse, Event, PluginConfiguration, PluginInfo, ReleaseIPRequest } from "./plugin_pb.js";
import { QueryRequest, QueryResponse } from "./storage_query_pb.js";

/**
 * Plugin is the general service definition for a Webmesh plugin.
 * It must be implemented by all plugins.
 *
 * @generated from service v1.Plugin
 */
export const Plugin = {
  typeName: "v1.Plugin",
  methods: {
    /**
     * GetInfo returns the information for the plugin.
     *
     * @generated from rpc v1.Plugin.GetInfo
     */
    getInfo: {
      name: "GetInfo",
      I: Empty,
      O: PluginInfo,
      kind: MethodKind.Unary,
    },
    /**
     * Configure starts and configures the plugin.
     *
     * @generated from rpc v1.Plugin.Configure
     */
    configure: {
      name: "Configure",
      I: PluginConfiguration,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * Close closes the plugin. It is called when the node is shutting down.
     *
     * @generated from rpc v1.Plugin.Close
     */
    close: {
      name: "Close",
      I: Empty,
      O: Empty,
      kind: MethodKind.Unary,
    },
  }
} as const;

/**
 * AuthPlugin is the service definition for a Webmesh auth plugin.
 *
 * @generated from service v1.AuthPlugin
 */
export const AuthPlugin = {
  typeName: "v1.AuthPlugin",
  methods: {
    /**
     * Authenticate authenticates a request.
     *
     * @generated from rpc v1.AuthPlugin.Authenticate
     */
    authenticate: {
      name: "Authenticate",
      I: AuthenticationRequest,
      O: AuthenticationResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

/**
 * WatchPlugin is the service definition for a Webmesh watch plugin.
 *
 * @generated from service v1.WatchPlugin
 */
export const WatchPlugin = {
  typeName: "v1.WatchPlugin",
  methods: {
    /**
     * Emit handles a watch event.
     *
     * @generated from rpc v1.WatchPlugin.Emit
     */
    emit: {
      name: "Emit",
      I: Event,
      O: Empty,
      kind: MethodKind.Unary,
    },
  }
} as const;

/**
 * IPAMPlugin is the service definition for a Webmesh IPAM plugin.
 *
 * @generated from service v1.IPAMPlugin
 */
export const IPAMPlugin = {
  typeName: "v1.IPAMPlugin",
  methods: {
    /**
     * Allocate allocates an IP for a node.
     *
     * @generated from rpc v1.IPAMPlugin.Allocate
     */
    allocate: {
      name: "Allocate",
      I: AllocateIPRequest,
      O: AllocatedIP,
      kind: MethodKind.Unary,
    },
    /**
     * Release releases an IP for a node.
     *
     * @generated from rpc v1.IPAMPlugin.Release
     */
    release: {
      name: "Release",
      I: ReleaseIPRequest,
      O: Empty,
      kind: MethodKind.Unary,
    },
  }
} as const;

/**
 * StorageQuerierPlugin is the service definition for a Webmesh storage querier plugin.
 *
 * @generated from service v1.StorageQuerierPlugin
 */
export const StorageQuerierPlugin = {
  typeName: "v1.StorageQuerierPlugin",
  methods: {
    /**
     * InjectQuerier is a stream opened by the node to faciliate read operations
     * against the mesh state. The signature is misleading, but it is required to be 
     * able to stream the query results back to the node. The node will open a stream 
     * to the plugin and send a PluginQueryResult message for every query that is 
     * received.
     *
     * @generated from rpc v1.StorageQuerierPlugin.InjectQuerier
     */
    injectQuerier: {
      name: "InjectQuerier",
      I: QueryResponse,
      O: QueryRequest,
      kind: MethodKind.BiDiStreaming,
    },
  }
} as const;

