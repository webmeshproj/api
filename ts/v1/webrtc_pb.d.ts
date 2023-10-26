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

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * StartDataChannelRequest is a request to start a data channel.
 * The answer and candidate fields are populated after the offer 
 * is received.
 *
 * @generated from message v1.StartDataChannelRequest
 */
export declare class StartDataChannelRequest extends Message<StartDataChannelRequest> {
  /**
   * node_id is the ID of the node to send the data to.
   *
   * @generated from field: string node_id = 1;
   */
  nodeId: string;

  /**
   * proto is the protocol of the traffic.
   *
   * @generated from field: string proto = 2;
   */
  proto: string;

  /**
   * dst is the destination address of the traffic.
   *
   * @generated from field: string dst = 3;
   */
  dst: string;

  /**
   * port is the destination port of the traffic. A port of 0 coupled
   * with the udp protocol indicates forwarding to the WireGuard interface.
   *
   * @generated from field: uint32 port = 4;
   */
  port: number;

  /**
   * answer is the answer to the offer.
   *
   * @generated from field: string answer = 5;
   */
  answer: string;

  /**
   * candidate is an ICE candidate.
   *
   * @generated from field: string candidate = 6;
   */
  candidate: string;

  constructor(data?: PartialMessage<StartDataChannelRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "v1.StartDataChannelRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): StartDataChannelRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): StartDataChannelRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): StartDataChannelRequest;

  static equals(a: StartDataChannelRequest | PlainMessage<StartDataChannelRequest> | undefined, b: StartDataChannelRequest | PlainMessage<StartDataChannelRequest> | undefined): boolean;
}

/**
 * DataChannelOffer is an offer for a data channel. Candidates
 * are sent after the offer is sent.
 *
 * @generated from message v1.DataChannelOffer
 */
export declare class DataChannelOffer extends Message<DataChannelOffer> {
  /**
   * offer is the offer.
   *
   * @generated from field: string offer = 1;
   */
  offer: string;

  /**
   * stun_servers is the list of STUN servers to use.
   *
   * @generated from field: repeated string stunServers = 2;
   */
  stunServers: string[];

  /**
   * candidate is an ICE candidate.
   *
   * @generated from field: string candidate = 3;
   */
  candidate: string;

  constructor(data?: PartialMessage<DataChannelOffer>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "v1.DataChannelOffer";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DataChannelOffer;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DataChannelOffer;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DataChannelOffer;

  static equals(a: DataChannelOffer | PlainMessage<DataChannelOffer> | undefined, b: DataChannelOffer | PlainMessage<DataChannelOffer> | undefined): boolean;
}

