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
// @generated from file v1/app.proto (package v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { AppQueryRequest, ConnectRequest, ConnectResponse, DaemonStatus, DisconnectRequest, DisconnectResponse, DropConnectionRequest, DropConnectionResponse, GetConnectionRequest, GetConnectionResponse, ListConnectionsRequest, ListConnectionsResponse, MetricsRequest, MetricsResponse, PutConnectionRequest, PutConnectionResponse, StatusRequest } from "./app_pb.js";
import { MethodKind } from "@bufbuild/protobuf";
import { QueryResponse } from "./storage_query_pb.js";

/**
 * AppDaemon is exposed by nodes running in the daemon mode.
 * This mode allows the node to run in an idle state and be controlled by an application.
 * The application can send commands to the node to execute tasks and receive responses.
 *
 * @generated from service v1.AppDaemon
 */
export const AppDaemon = {
  typeName: "v1.AppDaemon",
  methods: {
    /**
     * PutConnection stores the parameters for a connection in the daemon.
     *
     * @generated from rpc v1.AppDaemon.PutConnection
     */
    putConnection: {
      name: "PutConnection",
      I: PutConnectionRequest,
      O: PutConnectionResponse,
      kind: MethodKind.Unary,
    },
    /**
     * GetConnection retrieves the parameters and current status of a connection in the daemon.
     *
     * @generated from rpc v1.AppDaemon.GetConnection
     */
    getConnection: {
      name: "GetConnection",
      I: GetConnectionRequest,
      O: GetConnectionResponse,
      kind: MethodKind.Unary,
    },
    /**
     * DropConnection deletes all data stored for a given mesh connection.
     *
     * @generated from rpc v1.AppDaemon.DropConnection
     */
    dropConnection: {
      name: "DropConnection",
      I: DropConnectionRequest,
      O: DropConnectionResponse,
      kind: MethodKind.Unary,
    },
    /**
     * ListConnections retrieves the parameters and current status of all connections in the daemon.
     *
     * @generated from rpc v1.AppDaemon.ListConnections
     */
    listConnections: {
      name: "ListConnections",
      I: ListConnectionsRequest,
      O: ListConnectionsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Connect is used to establish a connection between the node and a mesh.
     *
     * @generated from rpc v1.AppDaemon.Connect
     */
    connect: {
      name: "Connect",
      I: ConnectRequest,
      O: ConnectResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Disconnect is used to disconnect the node from a mesh.
     *
     * @generated from rpc v1.AppDaemon.Disconnect
     */
    disconnect: {
      name: "Disconnect",
      I: DisconnectRequest,
      O: DisconnectResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Metrics is used to retrieve interface metrics for one or more mesh connections.
     *
     * @generated from rpc v1.AppDaemon.Metrics
     */
    metrics: {
      name: "Metrics",
      I: MetricsRequest,
      O: MetricsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Query is used to query a mesh connection for information.
     *
     * @generated from rpc v1.AppDaemon.Query
     */
    query: {
      name: "Query",
      I: AppQueryRequest,
      O: QueryResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Status is used to retrieve the status of the daemon. This includes a map of known connections and their statuses.
     *
     * @generated from rpc v1.AppDaemon.Status
     */
    status: {
      name: "Status",
      I: StatusRequest,
      O: DaemonStatus,
      kind: MethodKind.Unary,
    },
  }
} as const;

