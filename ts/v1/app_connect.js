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

import { AnnounceDHTRequest, AnnounceDHTResponse, ConnectRequest, ConnectResponse, DisconnectRequest, DisconnectResponse, LeaveDHTRequest, LeaveDHTResponse, MetricsRequest, MetricsResponse, StatusRequest, StatusResponse } from "./app_pb.js";
import { MethodKind } from "@bufbuild/protobuf";
import { PublishRequest, PublishResponse, QueryRequest, QueryResponse, SubscribeRequest, SubscriptionEvent } from "./storage_query_pb.js";

/**
 * AppDaemon is exposed by nodes running in the app-daemon mode. This mode
 * allows the node to run in an idle state and be controlled by the
 * application. The application can send commands to the node to execute
 * tasks and receive responses.
 *
 * @generated from service v1.AppDaemon
 */
export const AppDaemon = {
  typeName: "v1.AppDaemon",
  methods: {
    /**
     * Connect is used to establish a connection between the node and a
     * mesh. The provided struct is used to override any defaults configured
     * on the node.
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
     * Disconnect is used to disconnect the node from the mesh.
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
     * Query is used to query the mesh for information.
     *
     * @generated from rpc v1.AppDaemon.Query
     */
    query: {
      name: "Query",
      I: QueryRequest,
      O: QueryResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Metrics is used to retrieve interface metrics from the node.
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
     * Status is used to retrieve the status of the node.
     *
     * @generated from rpc v1.AppDaemon.Status
     */
    status: {
      name: "Status",
      I: StatusRequest,
      O: StatusResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Subscribe is used to subscribe to events in the mesh database.
     *
     * @generated from rpc v1.AppDaemon.Subscribe
     */
    subscribe: {
      name: "Subscribe",
      I: SubscribeRequest,
      O: SubscriptionEvent,
      kind: MethodKind.ServerStreaming,
    },
    /**
     * Publish is used to publish events to the mesh database. A restricted set
     * of keys are allowed to be published to.
     *
     * @generated from rpc v1.AppDaemon.Publish
     */
    publish: {
      name: "Publish",
      I: PublishRequest,
      O: PublishResponse,
      kind: MethodKind.Unary,
    },
    /**
     * AnnounceDHT is used to announce the node's presence on the Kademlia DHT
     * for other nodes to discover.
     *
     * @generated from rpc v1.AppDaemon.AnnounceDHT
     */
    announceDHT: {
      name: "AnnounceDHT",
      I: AnnounceDHTRequest,
      O: AnnounceDHTResponse,
      kind: MethodKind.Unary,
    },
    /**
     * LeaveDHT is used to leave the Kademlia DHT.
     *
     * @generated from rpc v1.AppDaemon.LeaveDHT
     */
    leaveDHT: {
      name: "LeaveDHT",
      I: LeaveDHTRequest,
      O: LeaveDHTResponse,
      kind: MethodKind.Unary,
    },
  }
};
