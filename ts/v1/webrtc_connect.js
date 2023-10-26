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
// @generated from file v1/webrtc.proto (package v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { DataChannelOffer, StartDataChannelRequest } from "./webrtc_pb.js";
import { MethodKind } from "@bufbuild/protobuf";
import { WebRTCSignal } from "./node_pb.js";

/**
 * WebRTC is a service for negotiating WebRTC connections to nodes in the mesh.
 * It is typically run alongside a TURN server, however the server can be configured
 * to use public STUN servers instead.
 *
 * @generated from service v1.WebRTC
 */
export const WebRTC = {
  typeName: "v1.WebRTC",
  methods: {
    /**
     * StartDataChannel requests a new WebRTC connection to a node.
     * The client speaks first with the request containing the node ID
     * and where forwarded packets should be sent. The server responds
     * with an offer and STUN servers to be used to establish a WebRTC connection.
     * The client should then respond with an answer to the offer that matches the
     * spec of the DataChannel.CHANNELS enum value.
     *
     * After the offer is accepted, the stream can be used to exchange ICE candidates
     * to speed up the connection process.
     *
     * @generated from rpc v1.WebRTC.StartDataChannel
     */
    startDataChannel: {
      name: "StartDataChannel",
      I: StartDataChannelRequest,
      O: DataChannelOffer,
      kind: MethodKind.BiDiStreaming,
    },
    /**
     * StartSignalChannel starts a signaling channel to a remote node. This can be used to 
     * negotiate WebRTC connections both inside and outside of the mesh. Messages on the wire
     * are proxied to the remote node.
     *
     * @generated from rpc v1.WebRTC.StartSignalChannel
     */
    startSignalChannel: {
      name: "StartSignalChannel",
      I: WebRTCSignal,
      O: WebRTCSignal,
      kind: MethodKind.BiDiStreaming,
    },
  }
};

