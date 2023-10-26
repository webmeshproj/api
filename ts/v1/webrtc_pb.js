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
// @generated from file v1/webrtc.proto (package v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";

/**
 * StartDataChannelRequest is a request to start a data channel.
 * The answer and candidate fields are populated after the offer 
 * is received.
 *
 * @generated from message v1.StartDataChannelRequest
 */
export const StartDataChannelRequest = proto3.makeMessageType(
  "v1.StartDataChannelRequest",
  () => [
    { no: 1, name: "node_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "proto", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "dst", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "port", kind: "scalar", T: 13 /* ScalarType.UINT32 */ },
    { no: 5, name: "answer", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "candidate", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * DataChannelOffer is an offer for a data channel. Candidates
 * are sent after the offer is sent.
 *
 * @generated from message v1.DataChannelOffer
 */
export const DataChannelOffer = proto3.makeMessageType(
  "v1.DataChannelOffer",
  () => [
    { no: 1, name: "offer", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "stunServers", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 3, name: "candidate", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

