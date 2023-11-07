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

// @generated by protoc-gen-es v1.4.2 with parameter "target=ts"
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
export class StartDataChannelRequest extends Message<StartDataChannelRequest> {
  /**
   * NodeID is the ID of the node to send the data to.
   *
   * @generated from field: string nodeID = 1;
   */
  nodeID = "";

  /**
   * Proto is the protocol of the traffic.
   *
   * @generated from field: string proto = 2;
   */
  proto = "";

  /**
   * Dst is the destination address of the traffic.
   *
   * @generated from field: string dst = 3;
   */
  dst = "";

  /**
   * Port is the destination port of the traffic. A port of 0 coupled
   * with the udp protocol indicates forwarding to the WireGuard interface.
   *
   * @generated from field: uint32 port = 4;
   */
  port = 0;

  /**
   * Answer is the answer to the offer.
   *
   * @generated from field: string answer = 5;
   */
  answer = "";

  /**
   * Candidate is an ICE candidate.
   *
   * @generated from field: string candidate = 6;
   */
  candidate = "";

  constructor(data?: PartialMessage<StartDataChannelRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "v1.StartDataChannelRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "nodeID", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "proto", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "dst", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "port", kind: "scalar", T: 13 /* ScalarType.UINT32 */ },
    { no: 5, name: "answer", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "candidate", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): StartDataChannelRequest {
    return new StartDataChannelRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): StartDataChannelRequest {
    return new StartDataChannelRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): StartDataChannelRequest {
    return new StartDataChannelRequest().fromJsonString(jsonString, options);
  }

  static equals(a: StartDataChannelRequest | PlainMessage<StartDataChannelRequest> | undefined, b: StartDataChannelRequest | PlainMessage<StartDataChannelRequest> | undefined): boolean {
    return proto3.util.equals(StartDataChannelRequest, a, b);
  }
}

/**
 * DataChannelOffer is an offer for a data channel. Candidates
 * are sent after the offer is sent.
 *
 * @generated from message v1.DataChannelOffer
 */
export class DataChannelOffer extends Message<DataChannelOffer> {
  /**
   * Offer is the offer.
   *
   * @generated from field: string offer = 1;
   */
  offer = "";

  /**
   * STUNServers is the list of STUN servers to use.
   *
   * @generated from field: repeated string stunServers = 2;
   */
  stunServers: string[] = [];

  /**
   * Candidate is an ICE candidate.
   *
   * @generated from field: string candidate = 3;
   */
  candidate = "";

  constructor(data?: PartialMessage<DataChannelOffer>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "v1.DataChannelOffer";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "offer", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "stunServers", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 3, name: "candidate", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DataChannelOffer {
    return new DataChannelOffer().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DataChannelOffer {
    return new DataChannelOffer().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DataChannelOffer {
    return new DataChannelOffer().fromJsonString(jsonString, options);
  }

  static equals(a: DataChannelOffer | PlainMessage<DataChannelOffer> | undefined, b: DataChannelOffer | PlainMessage<DataChannelOffer> | undefined): boolean {
    return proto3.util.equals(DataChannelOffer, a, b);
  }
}

