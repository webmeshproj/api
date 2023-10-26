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
// @generated from file v1/raft.proto (package v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { Duration, proto3 } from "@bufbuild/protobuf";

/**
 * RaftCommandType is the type of command being sent to the
 * Raft log.
 *
 * @generated from enum v1.RaftCommandType
 */
export const RaftCommandType = proto3.makeEnum(
  "v1.RaftCommandType",
  [
    {no: 0, name: "UNKNOWN"},
    {no: 1, name: "PUT"},
    {no: 2, name: "DELETE"},
  ],
);

/**
 * RaftLogEntry is the data of an entry in the Raft log.
 *
 * @generated from message v1.RaftLogEntry
 */
export const RaftLogEntry = proto3.makeMessageType(
  "v1.RaftLogEntry",
  () => [
    { no: 1, name: "type", kind: "enum", T: proto3.getEnumType(RaftCommandType) },
    { no: 2, name: "key", kind: "scalar", T: 12 /* ScalarType.BYTES */ },
    { no: 3, name: "value", kind: "scalar", T: 12 /* ScalarType.BYTES */ },
    { no: 4, name: "ttl", kind: "message", T: Duration },
  ],
);

/**
 * RaftApplyResponse is the response to an apply request. It
 * contains the result of applying the log entry.
 *
 * @generated from message v1.RaftApplyResponse
 */
export const RaftApplyResponse = proto3.makeMessageType(
  "v1.RaftApplyResponse",
  () => [
    { no: 1, name: "time", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "error", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * RaftSnapshot is the data of a snapshot.
 *
 * @generated from message v1.RaftSnapshot
 */
export const RaftSnapshot = proto3.makeMessageType(
  "v1.RaftSnapshot",
  () => [
    { no: 1, name: "kv", kind: "message", T: RaftDataItem, repeated: true },
  ],
);

/**
 * RaftDataItem represents a value in the Raft data store.
 *
 * @generated from message v1.RaftDataItem
 */
export const RaftDataItem = proto3.makeMessageType(
  "v1.RaftDataItem",
  () => [
    { no: 1, name: "key", kind: "scalar", T: 12 /* ScalarType.BYTES */ },
    { no: 2, name: "value", kind: "scalar", T: 12 /* ScalarType.BYTES */ },
    { no: 3, name: "ttl", kind: "message", T: Duration },
  ],
);

