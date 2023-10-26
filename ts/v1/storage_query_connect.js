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
// @generated from file v1/storage_query.proto (package v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { PublishRequest, PublishResponse, QueryRequest, QueryResponse, SubscribeRequest, SubscriptionEvent } from "./storage_query_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * StorageQueryService is the service for querying information about the mesh state.
 *
 * @generated from service v1.StorageQueryService
 */
export const StorageQueryService = {
  typeName: "v1.StorageQueryService",
  methods: {
    /**
     * Query is used to query the mesh for information.
     *
     * @generated from rpc v1.StorageQueryService.Query
     */
    query: {
      name: "Query",
      I: QueryRequest,
      O: QueryResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Publish is used to publish events to the mesh database. A restricted set
     * of keys are allowed to be published to. This is only available on nodes 
     * that are able to provide storage.
     *
     * @generated from rpc v1.StorageQueryService.Publish
     */
    publish: {
      name: "Publish",
      I: PublishRequest,
      O: PublishResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Subscribe is used to subscribe to events at a particular prefix. This is
     * only available on nodes that are able to provide storage.
     *
     * @generated from rpc v1.StorageQueryService.Subscribe
     */
    subscribe: {
      name: "Subscribe",
      I: SubscribeRequest,
      O: SubscriptionEvent,
      kind: MethodKind.ServerStreaming,
    },
  }
};

