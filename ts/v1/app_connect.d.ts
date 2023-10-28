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

// @generated by protoc-gen-connect-es v1.1.2
// @generated from file v1/app.proto (package v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { AppPublishRequest, AppQueryRequest, AppSubscribeRequest, ConnectRequest, ConnectResponse, DisconnectRequest, DisconnectResponse, MetricsRequest, MetricsResponse, StatusRequest, StatusResponse } from "./app_pb.js";
import { MethodKind } from "@bufbuild/protobuf";
import { PublishResponse, QueryResponse, SubscriptionEvent } from "./storage_query_pb.js";

/**
 * AppDaemon is exposed by nodes running in the app-daemon mode. This mode
 * allows the node to run in an idle state and be controlled by an application.
 * The application can send commands to the node to execute tasks and receive 
 * responses.
 *
 * @generated from service v1.AppDaemon
 */
export declare const AppDaemon: {
  readonly typeName: "v1.AppDaemon",
  readonly methods: {
    /**
     * Connect is used to establish a connection between the node and a mesh.
     *
     * @generated from rpc v1.AppDaemon.Connect
     */
    readonly connect: {
      readonly name: "Connect",
      readonly I: typeof ConnectRequest,
      readonly O: typeof ConnectResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * Disconnect is used to disconnect the node from a mesh.
     *
     * @generated from rpc v1.AppDaemon.Disconnect
     */
    readonly disconnect: {
      readonly name: "Disconnect",
      readonly I: typeof DisconnectRequest,
      readonly O: typeof DisconnectResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * Metrics is used to retrieve interface metrics for a mesh connection.
     *
     * @generated from rpc v1.AppDaemon.Metrics
     */
    readonly metrics: {
      readonly name: "Metrics",
      readonly I: typeof MetricsRequest,
      readonly O: typeof MetricsResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * Status is used to retrieve the status a mesh connection.
     *
     * @generated from rpc v1.AppDaemon.Status
     */
    readonly status: {
      readonly name: "Status",
      readonly I: typeof StatusRequest,
      readonly O: typeof StatusResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * Query is used to query a mesh for information.
     *
     * @generated from rpc v1.AppDaemon.Query
     */
    readonly query: {
      readonly name: "Query",
      readonly I: typeof AppQueryRequest,
      readonly O: typeof QueryResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * Subscribe is used to subscribe to events in a mesh database.
     *
     * @generated from rpc v1.AppDaemon.Subscribe
     */
    readonly subscribe: {
      readonly name: "Subscribe",
      readonly I: typeof AppSubscribeRequest,
      readonly O: typeof SubscriptionEvent,
      readonly kind: MethodKind.ServerStreaming,
    },
    /**
     * Publish is used to publish events to a mesh database.
     *
     * @generated from rpc v1.AppDaemon.Publish
     */
    readonly publish: {
      readonly name: "Publish",
      readonly I: typeof AppPublishRequest,
      readonly O: typeof PublishResponse,
      readonly kind: MethodKind.Unary,
    },
  }
};

