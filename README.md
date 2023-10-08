# Protocol Documentation

## Table of Contents

<div id="toc-container">

- [v1/node.proto](#v1%2fnode.proto)
  - [<span class="badge">M</span>DataChannelNegotiation](#v1.DataChannelNegotiation)
  - [<span class="badge">M</span>FeaturePort](#v1.FeaturePort)
  - [<span class="badge">M</span>GetStatusRequest](#v1.GetStatusRequest)
  - [<span class="badge">M</span>InterfaceMetrics](#v1.InterfaceMetrics)
  - [<span class="badge">M</span>MeshNode](#v1.MeshNode)
  - [<span class="badge">M</span>NodeList](#v1.NodeList)
  - [<span class="badge">M</span>PeerMetrics](#v1.PeerMetrics)
  - [<span class="badge">M</span>Status](#v1.Status)
  - [<span class="badge">M</span>WebRTCSignal](#v1.WebRTCSignal)
  - [<span class="badge">E</span>ClusterStatus](#v1.ClusterStatus)
  - [<span class="badge">E</span>DataChannel](#v1.DataChannel)
  - [<span class="badge">E</span>EdgeAttribute](#v1.EdgeAttribute)
  - [<span class="badge">E</span>Feature](#v1.Feature)
  - [<span class="badge">S</span>Node](#v1.Node)
- [v1/mesh.proto](#v1%2fmesh.proto)
  - [<span class="badge">M</span>GetNodeRequest](#v1.GetNodeRequest)
  - [<span class="badge">M</span>MeshEdge](#v1.MeshEdge)
  - [<span class="badge">M</span>MeshEdge.AttributesEntry](#v1.MeshEdge.AttributesEntry)
  - [<span class="badge">M</span>MeshEdges](#v1.MeshEdges)
  - [<span class="badge">M</span>MeshGraph](#v1.MeshGraph)
  - [<span class="badge">S</span>Mesh](#v1.Mesh)
- [v1/network_acls.proto](#v1%2fnetwork_acls.proto)
  - [<span class="badge">M</span>NetworkACL](#v1.NetworkACL)
  - [<span class="badge">M</span>NetworkACLs](#v1.NetworkACLs)
  - [<span class="badge">M</span>NetworkAction](#v1.NetworkAction)
  - [<span class="badge">M</span>Route](#v1.Route)
  - [<span class="badge">M</span>Routes](#v1.Routes)
  - [<span class="badge">E</span>ACLAction](#v1.ACLAction)
- [v1/rbac.proto](#v1%2frbac.proto)
  - [<span class="badge">M</span>Group](#v1.Group)
  - [<span class="badge">M</span>Groups](#v1.Groups)
  - [<span class="badge">M</span>RBACAction](#v1.RBACAction)
  - [<span class="badge">M</span>Role](#v1.Role)
  - [<span class="badge">M</span>RoleBinding](#v1.RoleBinding)
  - [<span class="badge">M</span>RoleBindings](#v1.RoleBindings)
  - [<span class="badge">M</span>Roles](#v1.Roles)
  - [<span class="badge">M</span>Rule](#v1.Rule)
  - [<span class="badge">M</span>Subject](#v1.Subject)
  - [<span class="badge">E</span>RuleResource](#v1.RuleResource)
  - [<span class="badge">E</span>RuleVerb](#v1.RuleVerb)
  - [<span class="badge">E</span>SubjectType](#v1.SubjectType)
- [v1/admin.proto](#v1%2fadmin.proto)
  - [<span class="badge">S</span>Admin](#v1.Admin)
- [v1/storage_query.proto](#v1%2fstorage_query.proto)
  - [<span class="badge">M</span>NetworkState](#v1.NetworkState)
  - [<span class="badge">M</span>PublishRequest](#v1.PublishRequest)
  - [<span class="badge">M</span>PublishResponse](#v1.PublishResponse)
  - [<span class="badge">M</span>QueryRequest](#v1.QueryRequest)
  - [<span class="badge">M</span>QueryResponse](#v1.QueryResponse)
  - [<span class="badge">M</span>SubscribeRequest](#v1.SubscribeRequest)
  - [<span class="badge">M</span>SubscriptionEvent](#v1.SubscriptionEvent)
  - [<span class="badge">E</span>QueryRequest.QueryCommand](#v1.QueryRequest.QueryCommand)
  - [<span class="badge">E</span>QueryRequest.QueryType](#v1.QueryRequest.QueryType)
  - [<span class="badge">S</span>StorageQueryService](#v1.StorageQueryService)
- [v1/app.proto](#v1%2fapp.proto)
  - [<span class="badge">M</span>AnnounceDHTRequest](#v1.AnnounceDHTRequest)
  - [<span class="badge">M</span>AnnounceDHTResponse](#v1.AnnounceDHTResponse)
  - [<span class="badge">M</span>ConnectRequest](#v1.ConnectRequest)
  - [<span class="badge">M</span>ConnectResponse](#v1.ConnectResponse)
  - [<span class="badge">M</span>DisconnectRequest](#v1.DisconnectRequest)
  - [<span class="badge">M</span>DisconnectResponse](#v1.DisconnectResponse)
  - [<span class="badge">M</span>LeaveDHTRequest](#v1.LeaveDHTRequest)
  - [<span class="badge">M</span>LeaveDHTResponse](#v1.LeaveDHTResponse)
  - [<span class="badge">M</span>MetricsRequest](#v1.MetricsRequest)
  - [<span class="badge">M</span>MetricsResponse](#v1.MetricsResponse)
  - [<span class="badge">M</span>MetricsResponse.InterfacesEntry](#v1.MetricsResponse.InterfacesEntry)
  - [<span class="badge">M</span>StatusRequest](#v1.StatusRequest)
  - [<span class="badge">M</span>StatusResponse](#v1.StatusResponse)
  - [<span class="badge">E</span>StatusResponse.ConnectionStatus](#v1.StatusResponse.ConnectionStatus)
  - [<span class="badge">S</span>AppDaemon](#v1.AppDaemon)
- [v1/raft.proto](#v1%2fraft.proto)
  - [<span class="badge">M</span>RaftApplyResponse](#v1.RaftApplyResponse)
  - [<span class="badge">M</span>RaftDataItem](#v1.RaftDataItem)
  - [<span class="badge">M</span>RaftLogEntry](#v1.RaftLogEntry)
  - [<span class="badge">M</span>RaftSnapshot](#v1.RaftSnapshot)
  - [<span class="badge">E</span>RaftCommandType](#v1.RaftCommandType)
- [v1/members.proto](#v1%2fmembers.proto)
  - [<span class="badge">M</span>JoinRequest](#v1.JoinRequest)
  - [<span class="badge">M</span>JoinRequest.DirectPeersEntry](#v1.JoinRequest.DirectPeersEntry)
  - [<span class="badge">M</span>JoinResponse](#v1.JoinResponse)
  - [<span class="badge">M</span>LeaveRequest](#v1.LeaveRequest)
  - [<span class="badge">M</span>LeaveResponse](#v1.LeaveResponse)
  - [<span class="badge">M</span>PeerConfigurations](#v1.PeerConfigurations)
  - [<span class="badge">M</span>StorageConfigurationRequest](#v1.StorageConfigurationRequest)
  - [<span class="badge">M</span>StorageConfigurationResponse](#v1.StorageConfigurationResponse)
  - [<span class="badge">M</span>StorageServer](#v1.StorageServer)
  - [<span class="badge">M</span>SubscribePeersRequest](#v1.SubscribePeersRequest)
  - [<span class="badge">M</span>UpdateRequest](#v1.UpdateRequest)
  - [<span class="badge">M</span>UpdateResponse](#v1.UpdateResponse)
  - [<span class="badge">M</span>WireGuardPeer](#v1.WireGuardPeer)
  - [<span class="badge">E</span>ConnectProtocol](#v1.ConnectProtocol)
  - [<span class="badge">S</span>Membership](#v1.Membership)
- [v1/plugin.proto](#v1%2fplugin.proto)
  - [<span class="badge">M</span>AllocateIPRequest](#v1.AllocateIPRequest)
  - [<span class="badge">M</span>AllocatedIP](#v1.AllocatedIP)
  - [<span class="badge">M</span>AuthenticationRequest](#v1.AuthenticationRequest)
  - [<span class="badge">M</span>AuthenticationRequest.HeadersEntry](#v1.AuthenticationRequest.HeadersEntry)
  - [<span class="badge">M</span>AuthenticationResponse](#v1.AuthenticationResponse)
  - [<span class="badge">M</span>Event](#v1.Event)
  - [<span class="badge">M</span>PluginConfiguration](#v1.PluginConfiguration)
  - [<span class="badge">M</span>PluginInfo](#v1.PluginInfo)
  - [<span class="badge">M</span>ReleaseIPRequest](#v1.ReleaseIPRequest)
  - [<span class="badge">E</span>Event.WatchEvent](#v1.Event.WatchEvent)
  - [<span class="badge">E</span>PluginInfo.PluginCapability](#v1.PluginInfo.PluginCapability)
  - [<span class="badge">S</span>AuthPlugin](#v1.AuthPlugin)
  - [<span class="badge">S</span>IPAMPlugin](#v1.IPAMPlugin)
  - [<span class="badge">S</span>Plugin](#v1.Plugin)
  - [<span class="badge">S</span>StorageQuerierPlugin](#v1.StorageQuerierPlugin)
  - [<span class="badge">S</span>WatchPlugin](#v1.WatchPlugin)
- [v1/storage_provider.proto](#v1%2fstorage_provider.proto)
  - [<span class="badge">M</span>AddObserverResponse](#v1.AddObserverResponse)
  - [<span class="badge">M</span>AddVoterResponse](#v1.AddVoterResponse)
  - [<span class="badge">M</span>BootstrapRequest](#v1.BootstrapRequest)
  - [<span class="badge">M</span>BootstrapResponse](#v1.BootstrapResponse)
  - [<span class="badge">M</span>DeleteValueRequest](#v1.DeleteValueRequest)
  - [<span class="badge">M</span>DeleteValueResponse](#v1.DeleteValueResponse)
  - [<span class="badge">M</span>DemoteVoterResponse](#v1.DemoteVoterResponse)
  - [<span class="badge">M</span>GetLeaderRequest](#v1.GetLeaderRequest)
  - [<span class="badge">M</span>GetPeersRequest](#v1.GetPeersRequest)
  - [<span class="badge">M</span>GetValueRequest](#v1.GetValueRequest)
  - [<span class="badge">M</span>GetValueResponse](#v1.GetValueResponse)
  - [<span class="badge">M</span>ListKeysRequest](#v1.ListKeysRequest)
  - [<span class="badge">M</span>ListKeysResponse](#v1.ListKeysResponse)
  - [<span class="badge">M</span>ListValuesRequest](#v1.ListValuesRequest)
  - [<span class="badge">M</span>ListValuesResponse](#v1.ListValuesResponse)
  - [<span class="badge">M</span>PrefixEvent](#v1.PrefixEvent)
  - [<span class="badge">M</span>PutValueRequest](#v1.PutValueRequest)
  - [<span class="badge">M</span>PutValueResponse](#v1.PutValueResponse)
  - [<span class="badge">M</span>RemoveServerResponse](#v1.RemoveServerResponse)
  - [<span class="badge">M</span>StoragePeer](#v1.StoragePeer)
  - [<span class="badge">M</span>StoragePeers](#v1.StoragePeers)
  - [<span class="badge">M</span>StorageStatus](#v1.StorageStatus)
  - [<span class="badge">M</span>StorageStatusRequest](#v1.StorageStatusRequest)
  - [<span class="badge">M</span>StorageValue](#v1.StorageValue)
  - [<span class="badge">M</span>SubscribePrefixRequest](#v1.SubscribePrefixRequest)
  - [<span class="badge">E</span>PrefixEvent.EventType](#v1.PrefixEvent.EventType)
  - [<span class="badge">S</span>StorageProviderPlugin](#v1.StorageProviderPlugin)
- [v1/webrtc.proto](#v1%2fwebrtc.proto)
  - [<span class="badge">M</span>DataChannelOffer](#v1.DataChannelOffer)
  - [<span class="badge">M</span>StartDataChannelRequest](#v1.StartDataChannelRequest)
  - [<span class="badge">S</span>WebRTC](#v1.WebRTC)
- [Scalar Value Types](#scalar-value-types)

</div>

<div class="file-heading">

## v1/node.proto

[Top](#title)

</div>

### DataChannelNegotiation

DataChannelNegotiation is the message for communicating data channels to
nodes.

| Field       | Type              | Label    | Description                                                         |
|-------------|-------------------|----------|---------------------------------------------------------------------|
| proto       | [string](#string) |          | proto is the protocol of the traffic.                               |
| src         | [string](#string) |          | src is the address of the client that initiated the request.        |
| dst         | [string](#string) |          | dst is the destination address of the traffic.                      |
| port        | [uint32](#uint32) |          | port is the destination port of the traffic.                        |
| offer       | [string](#string) |          | offer is the offer for the node to use as its local description.    |
| answer      | [string](#string) |          | answer is the answer for the node to use as its remote description. |
| candidate   | [string](#string) |          | candidate is an ICE candidate.                                      |
| stunServers | [string](#string) | repeated | stun_servers is the list of STUN servers to use.                    |

### FeaturePort

FeaturePort describes a feature and the port it is advertised on.

| Field   | Type                   | Label | Description                                    |
|---------|------------------------|-------|------------------------------------------------|
| feature | [Feature](#v1.Feature) |       | feature is the feature.                        |
| port    | [int32](#int32)        |       | port is the port the feature is advertised on. |

### GetStatusRequest

GetStatusRequest is a request to get the status of a node.

| Field | Type              | Label | Description                                                                   |
|-------|-------------------|-------|-------------------------------------------------------------------------------|
| id    | [string](#string) |       | id is the ID of the node. If unset, the status of the local node is returned. |

### InterfaceMetrics

InterfaceMetrics is the metrics for the WireGuard interface on a node.

| Field              | Type                           | Label    | Description                                                    |
|--------------------|--------------------------------|----------|----------------------------------------------------------------|
| deviceName         | [string](#string)              |          | device_name is the name of the device.                         |
| publicKey          | [string](#string)              |          | public_key is the public key of the node.                      |
| addressV4          | [string](#string)              |          | address_v4 is the IPv4 address of the node.                    |
| addressV6          | [string](#string)              |          | address_v6 is the IPv6 address of the node.                    |
| type               | [string](#string)              |          | type is the type of interface being used for wireguard.        |
| listenPort         | [int32](#int32)                |          | listen_port is the port wireguard is listening on.             |
| totalReceiveBytes  | [uint64](#uint64)              |          | total_receive_bytes is the total number of bytes received.     |
| totalTransmitBytes | [uint64](#uint64)              |          | total_transmit_bytes is the total number of bytes transmitted. |
| numPeers           | [int32](#int32)                |          | num_peers is the number of peers connected to the node.        |
| peers              | [PeerMetrics](#v1.PeerMetrics) | repeated | peers are the per-peer statistics.                             |

### MeshNode

MeshNode is a node that has been registered with the mesh.

| Field              | Type                                                    | Label    | Description                                                           |
|--------------------|---------------------------------------------------------|----------|-----------------------------------------------------------------------|
| id                 | [string](#string)                                       |          | id is the ID of the node.                                             |
| publicKey          | [string](#string)                                       |          | public_key is the public key of the node.                             |
| primaryEndpoint    | [string](#string)                                       |          | primary_endpoint is the primary endpoint of the node.                 |
| wireguardEndpoints | [string](#string)                                       | repeated | wireguard_endpoints is a list of WireGuard endpoints for the node.    |
| zoneAwarenessID    | [string](#string)                                       |          | zone_awareness_id is the zone awareness ID of the node.               |
| privateIPv4        | [string](#string)                                       |          | private_ipv4 is the private IPv4 address of the node.                 |
| privateIPv6        | [string](#string)                                       |          | private_ipv6 is the private IPv6 address of the node.                 |
| features           | [FeaturePort](#v1.FeaturePort)                          | repeated | features are a list of features and the ports they are advertised on. |
| multiaddrs         | [string](#string)                                       | repeated | multiaddrs are the multiaddrs of the node.                            |
| joinedAt           | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |          | joined_at is the time the node joined the cluster.                    |

### NodeList

NodeList is a list of nodes.

| Field | Type                     | Label    | Description                 |
|-------|--------------------------|----------|-----------------------------|
| nodes | [MeshNode](#v1.MeshNode) | repeated | nodes is the list of nodes. |

### PeerMetrics

PeerMetrics are the metrics for a node's peer.

| Field               | Type              | Label    | Description                                                                         |
|---------------------|-------------------|----------|-------------------------------------------------------------------------------------|
| publicKey           | [string](#string) |          | public_key is the public key of the peer.                                           |
| endpoint            | [string](#string) |          | endpoint is the connected endpoint of the peer.                                     |
| persistentKeepAlive | [string](#string) |          | persistent_keep_alive is the persistent keep alive interval for the peer.           |
| lastHandshakeTime   | [string](#string) |          | last_handshake_time is the last handshake time for the peer.                        |
| allowedIPs          | [string](#string) | repeated | allowed_ips is the list of allowed IPs for the peer.                                |
| protocolVersion     | [int64](#int64)   |          | protocol_version is the version of the wireguard protocol negotiated with the peer. |
| receiveBytes        | [uint64](#uint64) |          | receive_bytes is the bytes received from the peer.                                  |
| transmitBytes       | [uint64](#uint64) |          | transmit_bytes is the bytes transmitted to the peer.                                |

### Status

Status represents the status of a node.

| Field            | Type                                                    | Label    | Description                                                  |
|------------------|---------------------------------------------------------|----------|--------------------------------------------------------------|
| id               | [string](#string)                                       |          | id is the ID of the node.                                    |
| version          | [string](#string)                                       |          | version is the version of the node.                          |
| commit           | [string](#string)                                       |          | commit is the commit of the node.                            |
| build_date       | [string](#string)                                       |          | build_date is the build date of the node.                    |
| uptime           | [string](#string)                                       |          | uptime is the uptime of the node.                            |
| startedAt        | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |          | started_at is the time the node started.                     |
| features         | [FeaturePort](#v1.FeaturePort)                          | repeated | features is the list of features currently enabled.          |
| clusterStatus    | [ClusterStatus](#v1.ClusterStatus)                      |          | cluster_status is the status of the node in the cluster.     |
| currentLeader    | [string](#string)                                       |          | current_leader is the current leader of the cluster.         |
| interfaceMetrics | [InterfaceMetrics](#v1.InterfaceMetrics)                |          | interface_metrics are the metrics for the node's interfaces. |

### WebRTCSignal

WebRTCSignal is a signal sent to a remote peer over the WebRTC API.

| Field       | Type              | Label | Description                                                                                                                                                                                                        |
|-------------|-------------------|-------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| nodeID      | [string](#string) |       | node_id is the ID of the node to send the signal to. This is set by the original sender. On the node that receives the ReceiveSignalChannel request, this will be set to the ID of the node that sent the request. |
| candidate   | [string](#string) |       | candidate is an ICE candidate.                                                                                                                                                                                     |
| description | [string](#string) |       | description is a session description.                                                                                                                                                                              |

### ClusterStatus

ClusterStatus is the status of the node in the cluster.

| Name                   | Number | Description                                                                       |
|------------------------|--------|-----------------------------------------------------------------------------------|
| CLUSTER_STATUS_UNKNOWN | 0      | CLUSTER_STATUS_UNKNOWN is the default status.                                     |
| CLUSTER_LEADER         | 1      | CLUSTER_LEADER is the status for the leader node.                                 |
| CLUSTER_VOTER          | 2      | CLUSTER_VOTER is the status for a voter node.                                     |
| CLUSTER_OBSERVER       | 3      | CLUSTER_OBSERVER is the status for a non-voter node.                              |
| CLUSTER_NODE           | 4      | CLUSTER_NODE is the status of a node that is not a part of the storage consensus. |

### DataChannel

DataChannel are the data channels used when communicating over ICE

with a node.

| Name        | Number | Description                                                                                                                                                                                                             |
|-------------|--------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| CHANNELS    | 0      | CHANNELS is the data channel used for negotiating new channels. This is the first channel that is opened. The ID of the channel should be 0.                                                                            |
| CONNECTIONS | 1      | CONNECTIONS is the data channel used for negotiating new connections. This is a channel that is opened for each incoming connection from a client. The ID should start at 0 and be incremented for each new connection. |

### EdgeAttribute

EdgeAttribute are pre-defined edge attributes. They should

be used as their string values.

| Name                   | Number | Description                                          |
|------------------------|--------|------------------------------------------------------|
| EDGE_ATTRIBUTE_UNKNOWN | 0      | EDGE_ATTRIBUTE_UNKNOWN is an unknown edge attribute. |
| EDGE_ATTRIBUTE_NATIVE  | 1      | EDGE_ATTRIBUTE_NATIVE is a native edge attribute.    |
| EDGE_ATTRIBUTE_ICE     | 2      | EDGE_ATTRIBUTE_ICE is an ICE edge attribute.         |
| EDGE_ATTRIBUTE_LIBP2P  | 3      | EDGE_ATTRIBUTE_LIBP2P is a libp2p edge attribute.    |

### Feature

Feature is a list of features supported by a node.

| Name             | Number | Description                                                                                           |
|------------------|--------|-------------------------------------------------------------------------------------------------------|
| FEATURE_NONE     | 0      | FEATURE_NONE is the default feature set.                                                              |
| NODES            | 1      | NODES is the feature for nodes. This is always supported.                                             |
| LEADER_PROXY     | 2      | LEADER_PROXY is the feature for leader proxying.                                                      |
| MESH_API         | 3      | MESH_API is the feature for the mesh API. This will be deprecated in favor of the MEMBERSHIP feature. |
| ADMIN_API        | 4      | ADMIN_API is the feature for the admin API.                                                           |
| MEMBERSHIP       | 5      | MEMBERSHIP is the feature for membership. This is always supported on storage-providing members.      |
| METRICS          | 6      | METRICS is the feature for exposing metrics.                                                          |
| ICE_NEGOTIATION  | 7      | ICE_NEGOTIATION is the feature for ICE negotiation.                                                   |
| TURN_SERVER      | 8      | TURN_SERVER is the feature for TURN server.                                                           |
| MESH_DNS         | 9      | MESH_DNS is the feature for mesh DNS.                                                                 |
| FORWARD_MESH_DNS | 10     | FORWARD_MESH_DNS is the feature for forwarding mesh DNS lookups to other meshes.                      |
| STORAGE_QUERIER  | 11     | STORAGE_QUERIER is the feature for querying, publishing, and subscribing to mesh state.               |
| STORAGE_PROVIDER | 12     | STORAGE_PROVIDER is the feature for being able to provide distributed storage.                        |

### Node

Node is the service exposed on every node in the mesh to communicate
network

information amongst themselves and facilitate inbound/outbound
connections.

| Method Name          | Request Type                                                | Response Type                                               | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         |
|----------------------|-------------------------------------------------------------|-------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| GetStatus            | [GetStatusRequest](#v1.GetStatusRequest)                    | [Status](#v1.Status)                                        | GetStatus gets the status of a node in the cluster. If the node is not able to return the status of the ID requested, it should return an error.                                                                                                                                                                                                                                                                                                                                                                    |
| NegotiateDataChannel | [DataChannelNegotiation](#v1.DataChannelNegotiation) stream | [DataChannelNegotiation](#v1.DataChannelNegotiation) stream | NegotiateDataChannel is used to negotiate a WebRTC connection between a webmesh client and a node in the cluster. The handling server will send the target node the source address, the destination for traffic, and STUN/TURN servers to use for the negotiation. The node responds with an offer to be forwarded to the client. When the handler receives an answer from the client, it forwards it to the node. Once the node receives the answer, the stream can optionally be used to exchange ICE candidates. |
| ReceiveSignalChannel | [WebRTCSignal](#v1.WebRTCSignal) stream                     | [WebRTCSignal](#v1.WebRTCSignal) stream                     | ReceiveSignalChannel is used to receive a request to start a WebRTC connection between a remote node and this node. The node should wait for the client to send an offer, and then respond with an answer. Once the node receives the answer, the stream can optionally be used to exchange ICE candidates.                                                                                                                                                                                                         |

<div class="file-heading">

## v1/mesh.proto

[Top](#title)

</div>

### GetNodeRequest

GetNodeRequest is a request to get a node.

| Field | Type              | Label | Description               |
|-------|-------------------|-------|---------------------------|
| id    | [string](#string) |       | id is the ID of the node. |

### MeshEdge

MeshEdge is an edge between two nodes.

| Field      | Type                                                     | Label    | Description                                      |
|------------|----------------------------------------------------------|----------|--------------------------------------------------|
| source     | [string](#string)                                        |          | source is the source node.                       |
| target     | [string](#string)                                        |          | target is the target node.                       |
| weight     | [int32](#int32)                                          |          | weight is the weight of the edge.                |
| attributes | [MeshEdge.AttributesEntry](#v1.MeshEdge.AttributesEntry) | repeated | attributes is a list of attributes for the edge. |

### MeshEdge.AttributesEntry

| Field | Type              | Label | Description |
|-------|-------------------|-------|-------------|
| key   | [string](#string) |       |             |
| value | [string](#string) |       |             |

### MeshEdges

MeshEdges is a list of edges.

| Field | Type                     | Label    | Description                 |
|-------|--------------------------|----------|-----------------------------|
| items | [MeshEdge](#v1.MeshEdge) | repeated | items is the list of edges. |

### MeshGraph

MeshGraph is a graph of nodes.

| Field | Type                     | Label    | Description                                 |
|-------|--------------------------|----------|---------------------------------------------|
| nodes | [string](#string)        | repeated | nodes is the list of nodes.                 |
| edges | [MeshEdge](#v1.MeshEdge) | repeated | edges is the list of edges.                 |
| dot   | [string](#string)        |          | dot is the DOT representation of the graph. |

### Mesh

Mesh is a service that can optionally be exposed by a node. It provides
methods for

interfacing with the webmesh from the outside.

| Method Name  | Request Type                                     | Response Type              | Description                                                                                                |
|--------------|--------------------------------------------------|----------------------------|------------------------------------------------------------------------------------------------------------|
| GetNode      | [GetNodeRequest](#v1.GetNodeRequest)             | [MeshNode](#v1.MeshNode)   | GetNode gets a node by ID.                                                                                 |
| ListNodes    | [.google.protobuf.Empty](#google.protobuf.Empty) | [NodeList](#v1.NodeList)   | ListNodes lists all nodes.                                                                                 |
| GetMeshGraph | [.google.protobuf.Empty](#google.protobuf.Empty) | [MeshGraph](#v1.MeshGraph) | GetMeshGraph fetches the mesh graph. It returns a list of nodes, edges, and a rendering in the dot format. |

<div class="file-heading">

## v1/network_acls.proto

[Top](#title)

</div>

### NetworkACL

NetworkACL is a network ACL.

| Field            | Type                       | Label    | Description                                                                                                                                                                                                                                                                                                                                                                                                                                          |
|------------------|----------------------------|----------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| name             | [string](#string)          |          | name is the name of the ACL.                                                                                                                                                                                                                                                                                                                                                                                                                         |
| priority         | [int32](#int32)            |          | priority is the priority of the ACL. ACLs with higher priority are evaluated first.                                                                                                                                                                                                                                                                                                                                                                  |
| action           | [ACLAction](#v1.ACLAction) |          | action is the action to take when a request matches the ACL.                                                                                                                                                                                                                                                                                                                                                                                         |
| sourceNodes      | [string](#string)          | repeated | source_nodes is a list of source nodes to match against. If empty, all nodes are matched. Groups can be specified with the prefix "group:". If one or more of the nodes is '\*', all nodes are matched.                                                                                                                                                                                                                                              |
| destinationNodes | [string](#string)          | repeated | destination_nodes is a list of destination nodes to match against. If empty, all nodes are matched. Groups can be specified with the prefix "group:". If one or more of the nodes is '\*', all nodes are matched.                                                                                                                                                                                                                                    |
| sourceCIDRs      | [string](#string)          | repeated | source_cidrs is a list of source CIDRs to match against. If empty, all CIDRs are matched. If one or more of the CIDRs is '\*', all CIDRs are matched.                                                                                                                                                                                                                                                                                                |
| destinationCIDRs | [string](#string)          | repeated | destination_cidrs is a list of destination CIDRs to match against. If empty, all CIDRs are matched. If one or more of the CIDRs is '\*', all CIDRs are matched. // protocols is a list of protocols to match against. If empty, all protocols are matched. // Protocols can be specified by name or number. repeated string protocols = 8; // ports is a list of ports to match against. If empty, all ports are matched. repeated uint32 ports = 9; |

### NetworkACLs

NetworkACLs is a list of network ACLs.

| Field | Type                         | Label    | Description                        |
|-------|------------------------------|----------|------------------------------------|
| items | [NetworkACL](#v1.NetworkACL) | repeated | items is the list of network ACLs. |

### NetworkAction

NetworkAction is an action that can be performed on a network resource.
It is used

by implementations to evaluate network ACLs.

| Field   | Type              | Label | Description                                     |
|---------|-------------------|-------|-------------------------------------------------|
| srcNode | [string](#string) |       | src_node is the source node of the action.      |
| srcCIDR | [string](#string) |       | src_cidr is the source CIDR of the action.      |
| dstNode | [string](#string) |       | dst_node is the destination node of the action. |
| dstCIDR | [string](#string) |       | dst_cidr is the destination CIDR of the action. |

### Route

Route is a route that is broadcasted by one or more nodes.

| Field            | Type              | Label    | Description                                                                                    |
|------------------|-------------------|----------|------------------------------------------------------------------------------------------------|
| name             | [string](#string) |          | name is the name of the route.                                                                 |
| node             | [string](#string) |          | node is the node that broadcasts the route. A group can be specified with the prefix "group:". |
| destinationCIDRs | [string](#string) | repeated | destination_cidrs are the destination CIDRs of the route.                                      |
| nextHopNode      | [string](#string) |          | nextHopNode is an optional node that is used as the next hop for the route.                    |

### Routes

Routes is a list of routes.

| Field | Type               | Label    | Description                  |
|-------|--------------------|----------|------------------------------|
| items | [Route](#v1.Route) | repeated | items is the list of routes. |

### ACLAction

ACLAction is the action to take when a request matches an ACL.

| Name           | Number | Description                                                                       |
|----------------|--------|-----------------------------------------------------------------------------------|
| ACTION_UNKNOWN | 0      | ACTION_UNKNOWN is the default action for ACLs. It is synonymous with ACTION_DENY. |
| ACTION_ACCEPT  | 1      | ACTION_ACCEPT allows the request to proceed.                                      |
| ACTION_DENY    | 2      | ACTION_DENY denies the request.                                                   |

<div class="file-heading">

## v1/rbac.proto

[Top](#title)

</div>

### Group

Group is a group of subjects.

| Field    | Type                   | Label    | Description                                    |
|----------|------------------------|----------|------------------------------------------------|
| name     | [string](#string)      |          | name is the name of the group.                 |
| subjects | [Subject](#v1.Subject) | repeated | subjects is the list of subjects in the group. |

### Groups

Groups is a list of groups.

| Field | Type               | Label    | Description                  |
|-------|--------------------|----------|------------------------------|
| items | [Group](#v1.Group) | repeated | items is the list of groups. |

### RBACAction

RBACAction is an action that can be performed on a resource. It is used
by implementations

to evaluate rules.

| Field        | Type                             | Label | Description                                                                 |
|--------------|----------------------------------|-------|-----------------------------------------------------------------------------|
| resource     | [RuleResource](#v1.RuleResource) |       | resource is the resource on which the action is performed.                  |
| resourceName | [string](#string)                |       | resource_name is the name of the resource on which the action is performed. |
| verb         | [RuleVerb](#v1.RuleVerb)         |       | verb is the verb that is performed on the resource.                         |

### Role

Role is a role that can be assigned to a subject.

| Field | Type              | Label    | Description                                        |
|-------|-------------------|----------|----------------------------------------------------|
| name  | [string](#string) |          | name is the name of the role.                      |
| rules | [Rule](#v1.Rule)  | repeated | rules is the list of rules that apply to the role. |

### RoleBinding

RoleBinding is a binding of a role to one or more subjects.

| Field    | Type                   | Label    | Description                                                    |
|----------|------------------------|----------|----------------------------------------------------------------|
| name     | [string](#string)      |          | name is the name of the role binding.                          |
| role     | [string](#string)      |          | role is the name of the role to which the binding applies.     |
| subjects | [Subject](#v1.Subject) | repeated | subjects is the list of subjects to which the binding applies. |

### RoleBindings

RoleBindings is a list of role bindings.

| Field | Type                           | Label    | Description                         |
|-------|--------------------------------|----------|-------------------------------------|
| items | [RoleBinding](#v1.RoleBinding) | repeated | items is the list of role bindings. |

### Roles

Roles is a list of roles.

| Field | Type             | Label    | Description                 |
|-------|------------------|----------|-----------------------------|
| items | [Role](#v1.Role) | repeated | items is the list of roles. |

### Rule

Rule is a rule that applies to a resource.

| Field         | Type                             | Label    | Description                                                             |
|---------------|----------------------------------|----------|-------------------------------------------------------------------------|
| resources     | [RuleResource](#v1.RuleResource) | repeated | resources is the resources to which the rule applies.                   |
| resourceNames | [string](#string)                | repeated | resource_names is the list of resource names to which the rule applies. |
| verbs         | [RuleVerb](#v1.RuleVerb)         | repeated | verbs is the list of verbs that apply to the resource.                  |

### Subject

Subject is a subject to which a role can be bound.

| Field | Type                           | Label | Description                      |
|-------|--------------------------------|-------|----------------------------------|
| name  | [string](#string)              |       | name is the name of the subject. |
| type  | [SubjectType](#v1.SubjectType) |       | type is the type of the subject. |

### RuleResource

RuleResource is the resource type for a rule.

| Name                   | Number | Description                                                                                                       |
|------------------------|--------|-------------------------------------------------------------------------------------------------------------------|
| RESOURCE_UNKNOWN       | 0      | RESOURCE_UNKNOWN is an unknown resource.                                                                          |
| RESOURCE_VOTES         | 1      | RESOURCE_VOTES is the resource for voting in storage elections. The only verb evaluated for this resource is PUT. |
| RESOURCE_ROLES         | 2      | RESOURCE_ROLES is the resource for managing roles.                                                                |
| RESOURCE_ROLE_BINDINGS | 3      | RESOURCE_ROLE_BINDINGS is the resource for managing role bindings.                                                |
| RESOURCE_GROUPS        | 4      | RESOURCE_GROUPS is the resource for managing groups.                                                              |
| RESOURCE_NETWORK_ACLS  | 5      | RESOURCE_NETWORK_ACLS is the resource for managing network ACLs.                                                  |
| RESOURCE_ROUTES        | 6      | RESOURCE_ROUTES is the resource for managing routes.                                                              |
| RESOURCE_DATA_CHANNELS | 7      | RESOURCE_DATA_CHANNELS is the resource for creating data channels.                                                |
| RESOURCE_EDGES         | 8      | RESOURCE_EDGES is the resource for managing edges between nodes.                                                  |
| RESOURCE_OBSERVERS     | 9      | RESOURCE_OBSERVERS is the resource for managing observers. The only verb evaluated for this resource is PUT.      |
| RESOURCE_PUBSUB        | 10     | RESOURCE_PUBSUB is the resource for managing pubsub topics.                                                       |
| RESOURCE_ALL           | 999    | RESOURCE_ALL is a wildcard resource that matches all resources.                                                   |

### RuleVerb

RuleVerb is the verb type for a rule.

| Name         | Number | Description                                               |
|--------------|--------|-----------------------------------------------------------|
| VERB_UNKNOWN | 0      | VERB_UNKNOWN is an unknown verb.                          |
| VERB_PUT     | 1      | VERB_PUT is the verb for creating or updating a resource. |
| VERB_GET     | 2      | VERB_GET is the verb for getting a resource.              |
| VERB_DELETE  | 3      | VERB_DELETE is the verb for deleting a resource.          |
| VERB_ALL     | 999    | VERB_ALL is a wildcard verb that matches all verbs.       |

### SubjectType

SubjectType is the type of a subject.

| Name            | Number | Description                                                                                                                            |
|-----------------|--------|----------------------------------------------------------------------------------------------------------------------------------------|
| SUBJECT_UNKNOWN | 0      | SUBJECT_UNKNOWN is an unknown subject type.                                                                                            |
| SUBJECT_NODE    | 1      | SUBJECT_NODE is a subject type for a node.                                                                                             |
| SUBJECT_USER    | 2      | SUBJECT_USER is a subject type for a user.                                                                                             |
| SUBJECT_GROUP   | 3      | SUBJECT_GROUP is a subject type for a group.                                                                                           |
| SUBJECT_ALL     | 999    | SUBJECT_ALL is a wildcard subject type that matches all subject types. It can be used with a subject named '\*' to match all subjects. |

<div class="file-heading">

## v1/admin.proto

[Top](#title)

</div>

### Admin

Admin is the service that provides cluster admin operations. Most
methods

require the leader to be contacted.

| Method Name       | Request Type                                     | Response Type                                    | Description                                           |
|-------------------|--------------------------------------------------|--------------------------------------------------|-------------------------------------------------------|
| PutRole           | [Role](#v1.Role)                                 | [.google.protobuf.Empty](#google.protobuf.Empty) | PutRole creates or updates a role.                    |
| DeleteRole        | [Role](#v1.Role)                                 | [.google.protobuf.Empty](#google.protobuf.Empty) | DeleteRole deletes a role.                            |
| GetRole           | [Role](#v1.Role)                                 | [Role](#v1.Role)                                 | GetRole gets a role.                                  |
| ListRoles         | [.google.protobuf.Empty](#google.protobuf.Empty) | [Roles](#v1.Roles)                               | ListRoles gets all roles.                             |
| PutRoleBinding    | [RoleBinding](#v1.RoleBinding)                   | [.google.protobuf.Empty](#google.protobuf.Empty) | PutRoleBinding creates or updates a role binding.     |
| DeleteRoleBinding | [RoleBinding](#v1.RoleBinding)                   | [.google.protobuf.Empty](#google.protobuf.Empty) | DeleteRoleBinding deletes a role binding.             |
| GetRoleBinding    | [RoleBinding](#v1.RoleBinding)                   | [RoleBinding](#v1.RoleBinding)                   | GetRoleBinding gets a role binding.                   |
| ListRoleBindings  | [.google.protobuf.Empty](#google.protobuf.Empty) | [RoleBindings](#v1.RoleBindings)                 | ListRoleBindings gets all role bindings.              |
| PutGroup          | [Group](#v1.Group)                               | [.google.protobuf.Empty](#google.protobuf.Empty) | PutGroup creates or updates a group.                  |
| DeleteGroup       | [Group](#v1.Group)                               | [.google.protobuf.Empty](#google.protobuf.Empty) | DeleteGroup deletes a group.                          |
| GetGroup          | [Group](#v1.Group)                               | [Group](#v1.Group)                               | GetGroup gets a group.                                |
| ListGroups        | [.google.protobuf.Empty](#google.protobuf.Empty) | [Groups](#v1.Groups)                             | ListGroups gets all groups.                           |
| PutNetworkACL     | [NetworkACL](#v1.NetworkACL)                     | [.google.protobuf.Empty](#google.protobuf.Empty) | PutNetworkACL creates or updates a network ACL.       |
| DeleteNetworkACL  | [NetworkACL](#v1.NetworkACL)                     | [.google.protobuf.Empty](#google.protobuf.Empty) | DeleteNetworkACL deletes a network ACL.               |
| GetNetworkACL     | [NetworkACL](#v1.NetworkACL)                     | [NetworkACL](#v1.NetworkACL)                     | GetNetworkACL gets a network ACL.                     |
| ListNetworkACLs   | [.google.protobuf.Empty](#google.protobuf.Empty) | [NetworkACLs](#v1.NetworkACLs)                   | ListNetworkACLs gets all network ACLs.                |
| PutRoute          | [Route](#v1.Route)                               | [.google.protobuf.Empty](#google.protobuf.Empty) | PutRoute creates or updates a route.                  |
| DeleteRoute       | [Route](#v1.Route)                               | [.google.protobuf.Empty](#google.protobuf.Empty) | DeleteRoute deletes a route.                          |
| GetRoute          | [Route](#v1.Route)                               | [Route](#v1.Route)                               | GetRoute gets a route.                                |
| ListRoutes        | [.google.protobuf.Empty](#google.protobuf.Empty) | [Routes](#v1.Routes)                             | ListRoutes gets all routes.                           |
| PutEdge           | [MeshEdge](#v1.MeshEdge)                         | [.google.protobuf.Empty](#google.protobuf.Empty) | PutEdge creates or updates an edge between two nodes. |
| DeleteEdge        | [MeshEdge](#v1.MeshEdge)                         | [.google.protobuf.Empty](#google.protobuf.Empty) | DeleteEdge deletes an edge between two nodes.         |
| GetEdge           | [MeshEdge](#v1.MeshEdge)                         | [MeshEdge](#v1.MeshEdge)                         | GetEdge gets an edge between two nodes.               |
| ListEdges         | [.google.protobuf.Empty](#google.protobuf.Empty) | [MeshEdges](#v1.MeshEdges)                       | ListEdges gets all current edges.                     |

<div class="file-heading">

## v1/storage_query.proto

[Top](#title)

</div>

### NetworkState

NetworkState represents the full network state as returned by

a network state query.

| Field     | Type              | Label | Description |
|-----------|-------------------|-------|-------------|
| networkV4 | [string](#string) |       |             |
| networkV6 | [string](#string) |       |             |
| domain    | [string](#string) |       |             |

### PublishRequest

PublishRequest is sent by the application to the node to publish events.

This currently only supports database events.

| Field | Type                                                  | Label | Description                                                             |
|-------|-------------------------------------------------------|-------|-------------------------------------------------------------------------|
| key   | [bytes](#bytes)                                       |       | key is the key of the event.                                            |
| value | [bytes](#bytes)                                       |       | value is the value of the event. This will be the raw value of the key. |
| ttl   | [google.protobuf.Duration](#google.protobuf.Duration) |       | ttl is the time for the event to live in the database.                  |

### PublishResponse

PublishResponse is the response to a publish request. This is currently

empty.

### QueryRequest

QueryRequest is sent by the application to the node to query the mesh
for

information.

| Field   | Type                                                       | Label | Description                                                                                                                                                                                                                      |
|---------|------------------------------------------------------------|-------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| command | [QueryRequest.QueryCommand](#v1.QueryRequest.QueryCommand) |       | command is the command of the query.                                                                                                                                                                                             |
| type    | [QueryRequest.QueryType](#v1.QueryRequest.QueryType)       |       | type is the type of the query.                                                                                                                                                                                                   |
| query   | [string](#string)                                          |       | query is the string of the query. This follows the format of a label selector and is only applicable for certain queries. For get queries this will usually be an ID. For list queries this will usually be one or more filters. |

### QueryResponse

QueryResponse is the message containing a mesh query result.

| Field | Type            | Label    | Description                                                                                                  |
|-------|-----------------|----------|--------------------------------------------------------------------------------------------------------------|
| items | [bytes](#bytes) | repeated | items contain the results of the query. These will be protobuf json-encoded objects of the given query type. |

### SubscribeRequest

SubscribeRequest is sent by the application to the node to subscribe to

events. This currently only supports database events.

| Field  | Type            | Label | Description                                         |
|--------|-----------------|-------|-----------------------------------------------------|
| prefix | [bytes](#bytes) |       | prefix is the prefix of the events to subscribe to. |

### SubscriptionEvent

SubscriptionEvent is a message containing a subscription event.

| Field | Type            | Label | Description                                                             |
|-------|-----------------|-------|-------------------------------------------------------------------------|
| key   | [bytes](#bytes) |       | key is the key of the event.                                            |
| value | [bytes](#bytes) |       | value is the value of the event. This will be the raw value of the key. |

### QueryRequest.QueryCommand

QueryCommand is the type of the query.

| Name | Number | Description                                               |
|------|--------|-----------------------------------------------------------|
| GET  | 0      | GET is the command to get a value.                        |
| LIST | 1      | LIST is the command to list keys with an optional prefix. |

### QueryRequest.QueryType

QueryType is the type of object being queried.

| Name          | Number | Description                                           |
|---------------|--------|-------------------------------------------------------|
| VALUE         | 0      | VALUE represents a raw value query at a supplied key. |
| KEYS          | 1      | KEYS is the type for querying keys.                   |
| PEERS         | 2      | PEERS is the type for querying peers.                 |
| EDGES         | 3      | EDGES is the type for querying edges.                 |
| ROUTES        | 4      | ROUTES is the type for querying routes.               |
| ACLS          | 5      | ACLS is the type for querying ACLs.                   |
| ROLES         | 6      | ROLES is the type for querying roles.                 |
| ROLEBINDINGS  | 7      | ROLEBINDINGS is the type for querying role bindings.  |
| GROUPS        | 8      | GROUPS is the type for querying groups.               |
| NETWORK_STATE | 9      | NETWORK_STATE is the type for querying network state. |

### StorageQueryService

StorageQueryService is the service for querying information about the
mesh state.

| Method Name | Request Type                             | Response Type                                     | Description                                                                                                                                                                        |
|-------------|------------------------------------------|---------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Query       | [QueryRequest](#v1.QueryRequest)         | [QueryResponse](#v1.QueryResponse)                | Query is used to query the mesh for information.                                                                                                                                   |
| Publish     | [PublishRequest](#v1.PublishRequest)     | [PublishResponse](#v1.PublishResponse)            | Publish is used to publish events to the mesh database. A restricted set of keys are allowed to be published to. This is only available on nodes that are able to provide storage. |
| Subscribe   | [SubscribeRequest](#v1.SubscribeRequest) | [SubscriptionEvent](#v1.SubscriptionEvent) stream | Subscribe is used to subscribe to events at a particular prefix. This is only available on nodes that are able to provide storage.                                                 |

<div class="file-heading">

## v1/app.proto

[Top](#title)

</div>

### AnnounceDHTRequest

AnnounceDHTRequest is sent by the application to the node to announce
the

node's presence on the Kademlia DHT for other nodes to discover.

| Field            | Type              | Label    | Description                                                                                                                                          |
|------------------|-------------------|----------|------------------------------------------------------------------------------------------------------------------------------------------------------|
| bootstrapServers | [string](#string) | repeated | Bootstrap servers are optional bootstrap servers to use for bootstrapping the DHT. If not provided, the node will use the default bootstrap servers. |
| psk              | [string](#string) |          | PSK is the pre-shared key to use for the DHT.                                                                                                        |

### AnnounceDHTResponse

AnnounceDHTResponse is returned by the AnnounceDHT RPC.

### ConnectRequest

ConnectRequest is sent by the application to the node to establish a

connection to a mesh. This message will eventually contain unique

identifiers to allow creating connections to multiple meshes.

| Field            | Type                                              | Label | Description                                                                                                    |
|------------------|---------------------------------------------------|-------|----------------------------------------------------------------------------------------------------------------|
| config           | [google.protobuf.Struct](#google.protobuf.Struct) |       | Config is used to override any defaults configured on the node.                                                |
| disableBootstrap | [bool](#bool)                                     |       | Disable bootstrap tells a node that is otherwise configured to bootstrap to not bootstrap for this connection. |
| joinPSK          | [string](#string)                                 |       | Join PSK is the pre-shared key to use for joining the mesh.                                                    |

### ConnectResponse

ConnectResponse is returned by the Connect RPC.

| Field      | Type              | Label | Description                                   |
|------------|-------------------|-------|-----------------------------------------------|
| nodeID     | [string](#string) |       | node id is the unique identifier of the node. |
| meshDomain | [string](#string) |       | mesh domain is the domain of the mesh.        |
| ipv4       | [string](#string) |       | ipv4 is the IPv4 address of the node.         |
| ipv6       | [string](#string) |       | ipv6 is the IPv6 address of the node.         |

### DisconnectRequest

DisconnectRequest is sent by the application to the node to disconnect

from a mesh. This message will eventually contain unique identifiers

for allowing the application to disconnect from a specific mesh.

### DisconnectResponse

DisconnectResponse is returned by the Disconnect RPC.

### LeaveDHTRequest

LeaveDHTRequest is sent by the application to the node to leave the
Kademlia

DHT.

| Field | Type              | Label | Description                                              |
|-------|-------------------|-------|----------------------------------------------------------|
| psk   | [string](#string) |       | PSK is the pre-shared key that was used to join the DHT. |

### LeaveDHTResponse

LeaveDHTResponse is returned by the LeaveDHT RPC.

### MetricsRequest

MetricsRequest is sent by the application to the node to retrieve
interface

metrics. It is intentionally empty for now, but can eventually be used
to

query specific interfaces/metrics.

### MetricsResponse

MetricsResponse is a message containing interface metrics.

| Field      | Type                                                                   | Label    | Description                                        |
|------------|------------------------------------------------------------------------|----------|----------------------------------------------------|
| interfaces | [MetricsResponse.InterfacesEntry](#v1.MetricsResponse.InterfacesEntry) | repeated | interfaces is a map of interface names to metrics. |

### MetricsResponse.InterfacesEntry

| Field | Type                                     | Label | Description |
|-------|------------------------------------------|-------|-------------|
| key   | [string](#string)                        |       |             |
| value | [InterfaceMetrics](#v1.InterfaceMetrics) |       |             |

### StatusRequest

StatusRequest is sent by the application to the node to retrieve the
status

of the node.

### StatusResponse

StatusResponse is a message containing the status of the node.

| Field             | Type                                                                   | Label | Description                                                               |
|-------------------|------------------------------------------------------------------------|-------|---------------------------------------------------------------------------|
| connection_status | [StatusResponse.ConnectionStatus](#v1.StatusResponse.ConnectionStatus) |       | connection status is the status of the connection.                        |
| node              | [MeshNode](#v1.MeshNode)                                               |       | node is the node status. This is only populated if the node is connected. |

### StatusResponse.ConnectionStatus

| Name         | Number | Description                                                                   |
|--------------|--------|-------------------------------------------------------------------------------|
| DISCONNECTED | 0      | DISCONNECTED indicates that the node is not connected to a mesh.              |
| CONNECTING   | 1      | CONNECTING indicates that the node is in the process of connecting to a mesh. |
| CONNECTED    | 2      | CONNECTED indicates that the node is connected to a mesh.                     |

### AppDaemon

AppDaemon is exposed by nodes running in the app-daemon mode. This mode

allows the node to run in an idle state and be controlled by the

application. The application can send commands to the node to execute

tasks and receive responses.

| Method Name | Request Type                                 | Response Type                                     | Description                                                                                                                                         |
|-------------|----------------------------------------------|---------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------|
| Connect     | [ConnectRequest](#v1.ConnectRequest)         | [ConnectResponse](#v1.ConnectResponse)            | Connect is used to establish a connection between the node and a mesh. The provided struct is used to override any defaults configured on the node. |
| Disconnect  | [DisconnectRequest](#v1.DisconnectRequest)   | [DisconnectResponse](#v1.DisconnectResponse)      | Disconnect is used to disconnect the node from the mesh.                                                                                            |
| Query       | [QueryRequest](#v1.QueryRequest)             | [QueryResponse](#v1.QueryResponse) stream         | Query is used to query the mesh for information.                                                                                                    |
| Metrics     | [MetricsRequest](#v1.MetricsRequest)         | [MetricsResponse](#v1.MetricsResponse)            | Metrics is used to retrieve interface metrics from the node.                                                                                        |
| Status      | [StatusRequest](#v1.StatusRequest)           | [StatusResponse](#v1.StatusResponse)              | Status is used to retrieve the status of the node.                                                                                                  |
| Subscribe   | [SubscribeRequest](#v1.SubscribeRequest)     | [SubscriptionEvent](#v1.SubscriptionEvent) stream | Subscribe is used to subscribe to events in the mesh database.                                                                                      |
| Publish     | [PublishRequest](#v1.PublishRequest)         | [PublishResponse](#v1.PublishResponse)            | Publish is used to publish events to the mesh database. A restricted set of keys are allowed to be published to.                                    |
| AnnounceDHT | [AnnounceDHTRequest](#v1.AnnounceDHTRequest) | [AnnounceDHTResponse](#v1.AnnounceDHTResponse)    | AnnounceDHT is used to announce the node's presence on the Kademlia DHT for other nodes to discover.                                                |
| LeaveDHT    | [LeaveDHTRequest](#v1.LeaveDHTRequest)       | [LeaveDHTResponse](#v1.LeaveDHTResponse)          | LeaveDHT is used to leave the Kademlia DHT.                                                                                                         |

<div class="file-heading">

## v1/raft.proto

[Top](#title)

</div>

### RaftApplyResponse

RaftApplyResponse is the response to an apply request. It

contains the result of applying the log entry.

| Field | Type              | Label | Description                                            |
|-------|-------------------|-------|--------------------------------------------------------|
| time  | [string](#string) |       | time is the total time it took to apply the log entry. |
| error | [string](#string) |       | error is an error that occurred during the apply.      |

### RaftDataItem

RaftDataItem represents a value in the Raft data store.

| Field | Type                                                  | Label | Description                               |
|-------|-------------------------------------------------------|-------|-------------------------------------------|
| key   | [bytes](#bytes)                                       |       | key is the key of the data item.          |
| value | [bytes](#bytes)                                       |       | value is the value of the data item.      |
| ttl   | [google.protobuf.Duration](#google.protobuf.Duration) |       | ttl is the time to live of the data item. |

### RaftLogEntry

RaftLogEntry is the data of an entry in the Raft log.

| Field | Type                                                  | Label | Description                               |
|-------|-------------------------------------------------------|-------|-------------------------------------------|
| type  | [RaftCommandType](#v1.RaftCommandType)                |       | type is the type of the log entry.        |
| key   | [bytes](#bytes)                                       |       | key is the key of the log entry.          |
| value | [bytes](#bytes)                                       |       | value is the value of the log entry.      |
| ttl   | [google.protobuf.Duration](#google.protobuf.Duration) |       | ttl is the time to live of the log entry. |

### RaftSnapshot

RaftSnapshot is the data of a snapshot.

| Field | Type                             | Label    | Description |
|-------|----------------------------------|----------|-------------|
| kv    | [RaftDataItem](#v1.RaftDataItem) | repeated |             |

### RaftCommandType

RaftCommandType is the type of command being sent to the

Raft log.

| Name    | Number | Description                                          |
|---------|--------|------------------------------------------------------|
| UNKNOWN | 0      | UNKNOWN is the unknown command type.                 |
| PUT     | 1      | PUT is the command for putting a key/value pair.     |
| DELETE  | 2      | DELETE is the command for deleting a key/value pair. |

<div class="file-heading">

## v1/members.proto

[Top](#title)

</div>

### JoinRequest

JoinRequest is a request to join the cluster.

| Field              | Type                                                             | Label    | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               |
|--------------------|------------------------------------------------------------------|----------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| id                 | [string](#string)                                                |          | id is the ID of the node.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
| publicKey          | [string](#string)                                                |          | public_key is the public key of the node to broadcast to peers.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| primaryEndpoint    | [string](#string)                                                |          | primary_endpoint is a routable address for the node. If left unset, the node is assumed to be behind a NAT and not directly accessible.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |
| wireguardEndpoints | [string](#string)                                                | repeated | wireguard_endpoints is a list of WireGuard endpoints for the node.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        |
| zoneAwarenessID    | [string](#string)                                                |          | zone_awareness_id is the zone awareness ID of the node.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |
| assignIPv4         | [bool](#bool)                                                    |          | assign_ipv4 is whether an IPv4 address should be assigned to the node.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    |
| preferStorageIPv6  | [bool](#bool)                                                    |          | prefer_storage_ipv6 is whether IPv6 should be preferred over IPv4 for storage communication. This is only used if assign_ipv4 is true.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    |
| asVoter            | [bool](#bool)                                                    |          | as_voter is whether the node should receive a vote in elections. The request will be denied if the node is not allowed to vote.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| asObserver         | [bool](#bool)                                                    |          | as_observer is whether the node should be added as an observer. They will receive updates to the storage, but not be able to vote in elections.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| routes             | [string](#string)                                                | repeated | routes is a list of routes to advertise to peers. The request will be denied if the node is not allowed to put routes.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    |
| directPeers        | [JoinRequest.DirectPeersEntry](#v1.JoinRequest.DirectPeersEntry) | repeated | direct_peers is a map of extra peers that should be connected to directly over relays. The provided edge attribute is the callers preference of how the relay should be created. The request will be denied if the node is not allowed to put data channels or edges. The default joining behavior creates direct links between the caller and the joiner. If the caller has a primary endpoint, the joiner will link the caller to all other nodes with a primary endpoint. If the caller has a zone awareness ID, the joiner will link the caller to all other nodes with the same zone awareness ID that also have a primary endpoint. |
| features           | [FeaturePort](#v1.FeaturePort)                                   | repeated | features is a list of features supported by the node that should be advertised to peers and the port they are available on.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               |
| multiaddrs         | [string](#string)                                                | repeated | multiaddrs are libp2p multiaddresses this node is listening on.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |

### JoinRequest.DirectPeersEntry

| Field | Type                                   | Label | Description |
|-------|----------------------------------------|-------|-------------|
| key   | [string](#string)                      |       |             |
| value | [ConnectProtocol](#v1.ConnectProtocol) |       |             |

### JoinResponse

JoinResponse is a response to a join request.

| Field       | Type                               | Label    | Description                                                                                                                                                                                                                       |
|-------------|------------------------------------|----------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| addressIPv4 | [string](#string)                  |          | address_ipv4 is the private IPv4 wireguard address of the node in CIDR format representing the network. This is only set if assign_ipv4 was set in the request or no network_ipv6 was provided.                                   |
| addressIPv6 | [string](#string)                  |          | address_ipv6 is the IPv6 network assigned to the node.                                                                                                                                                                            |
| networkIPv4 | [string](#string)                  |          | network_ipv4 is the IPv4 network of the Mesh.                                                                                                                                                                                     |
| networkIPv6 | [string](#string)                  |          | network_ipv6 is the IPv6 network of the Mesh.                                                                                                                                                                                     |
| peers       | [WireGuardPeer](#v1.WireGuardPeer) | repeated | peers is a list of wireguard peers to connect to.                                                                                                                                                                                 |
| iceServers  | [string](#string)                  | repeated | ice_servers is a list of public nodes that can be used to negotiate ICE connections if required. This may only be populated when one of the peers has the ICE flag set. This must be set if the requestor specifies direct_peers. |
| dnsServers  | [string](#string)                  | repeated | dns_servers is a list of peers offering DNS services.                                                                                                                                                                             |
| meshDomain  | [string](#string)                  |          | mesh_domain is the domain of the mesh.                                                                                                                                                                                            |

### LeaveRequest

LeaveRequest is a request to leave the cluster.

| Field | Type              | Label | Description               |
|-------|-------------------|-------|---------------------------|
| id    | [string](#string) |       | id is the ID of the node. |

### LeaveResponse

LeaveResponse is a response to a leave request. It is currently empty.

### PeerConfigurations

PeerConfigurations is a stream of peer configurations.

| Field      | Type                               | Label    | Description                                                                                                                                                             |
|------------|------------------------------------|----------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| peers      | [WireGuardPeer](#v1.WireGuardPeer) | repeated | peers is a list of wireguard peers to connect to.                                                                                                                       |
| iceServers | [string](#string)                  | repeated | ice_servers is a list of public nodes that can be used to negotiate ICE connections if required. This may only be populated when one of the peers has the ICE flag set. |
| dnsServers | [string](#string)                  | repeated | dns_servers is a list of peers offering DNS services.                                                                                                                   |

### StorageConfigurationRequest

StorageConfigurationRequest is a request to get the current Storage
configuration.

### StorageConfigurationResponse

StorageConfigurationResponse is a response to a Storage configuration
request.

| Field   | Type                               | Label    | Description                                                  |
|---------|------------------------------------|----------|--------------------------------------------------------------|
| servers | [StorageServer](#v1.StorageServer) | repeated | servers is the list of servers in the storage configuration. |

### StorageServer

StorageServer is a server in the Storage configuration.

| Field    | Type                               | Label | Description                                |
|----------|------------------------------------|-------|--------------------------------------------|
| id       | [string](#string)                  |       | ID is the ID of the server.                |
| suffrage | [ClusterStatus](#v1.ClusterStatus) |       | Suffrage is the suffrage of the server.    |
| address  | [string](#string)                  |       | Address is the mesh address of the server. |

### SubscribePeersRequest

SubscribePeersRequest is a request to subscribe to peer updates.

| Field | Type              | Label | Description               |
|-------|-------------------|-------|---------------------------|
| id    | [string](#string) |       | id is the ID of the node. |

### UpdateRequest

UpdateRequest contains most of the same fields as JoinRequest, but is

used to update the state of a node in the cluster.

| Field              | Type                           | Label    | Description                                                                                                                             |
|--------------------|--------------------------------|----------|-----------------------------------------------------------------------------------------------------------------------------------------|
| id                 | [string](#string)              |          | id is the ID of the node.                                                                                                               |
| publicKey          | [string](#string)              |          | public_key is the public key of the node to broadcast to peers.                                                                         |
| primaryEndpoint    | [string](#string)              |          | primary_endpoint is a routable address for the node. If left unset, the node is assumed to be behind a NAT and not directly accessible. |
| wireguardEndpoints | [string](#string)              | repeated | wireguard_endpoints is a list of WireGuard endpoints for the node.                                                                      |
| zoneAwarenessID    | [string](#string)              |          | zone_awareness_id is the zone awareness ID of the node.                                                                                 |
| asVoter            | [bool](#bool)                  |          | as_voter is whether the node should receive a vote in elections. The request will be denied if the node is not allowed to vote.         |
| routes             | [string](#string)              | repeated | routes is a list of routes to advertise to peers. The request will be denied if the node is not allowed to put routes.                  |
| features           | [FeaturePort](#v1.FeaturePort) | repeated | features is a list of features supported by the node that should be advertised to peers and the port they are available on.             |
| multiaddrs         | [string](#string)              | repeated | multiaddrs are libp2p multiaddresses this node is listening on.                                                                         |

### UpdateResponse

UpdateResponse is a response to an update request. It is currently
empty.

### WireGuardPeer

WireGuardPeer is a peer in the Wireguard network.

| Field         | Type                                   | Label    | Description                                                 |
|---------------|----------------------------------------|----------|-------------------------------------------------------------|
| node          | [MeshNode](#v1.MeshNode)               |          | Node is information about this node.                        |
| allowedIPs    | [string](#string)                      | repeated | allowed_ips is the list of allowed IPs for the peer.        |
| allowedRoutes | [string](#string)                      | repeated | allowed_routes is the list of allowed routes for the peer.  |
| proto         | [ConnectProtocol](#v1.ConnectProtocol) |          | proto indicates the protocol to use to connect to the peer. |

### ConnectProtocol

ConnectProtocol is a type of protocol for establishing a connection into
a mesh.

| Name           | Number | Description                                                                                             |
|----------------|--------|---------------------------------------------------------------------------------------------------------|
| CONNECT_NATIVE | 0      | CONNECT_NATIVE indicates that the node should connect to other nodes via the native webmesh mechanisms. |
| CONNECT_ICE    | 1      | CONNECT_ICE indicates that the node should connect to other nodes via ICE.                              |
| CONNECT_LIBP2P | 2      | CONNECT_LIBP2P indicates that the node should connect to other nodes via libp2p.                        |

### Membership

The membership service is exposed on storage-providing nodes to allow
nodes to join

and leave the cluster. This service is meant to be made available
publicly

to allow people in from the outside.

| Method Name             | Request Type                                                   | Response Type                                                    | Description                                                                                                                                                                                                                                                                                                                                |
|-------------------------|----------------------------------------------------------------|------------------------------------------------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Join                    | [JoinRequest](#v1.JoinRequest)                                 | [JoinResponse](#v1.JoinResponse)                                 | Join is used to join a node to the mesh.                                                                                                                                                                                                                                                                                                   |
| Update                  | [UpdateRequest](#v1.UpdateRequest)                             | [UpdateResponse](#v1.UpdateResponse)                             | Update is used by a node to update its state in the mesh. The node will be updated in the mesh and will be able to query the mesh state or vote in elections. Only non-empty fields will be updated. It is almost semantically equivalent to a join request with the same ID, but redefined to avoid confusion and to allow for expansion. |
| Leave                   | [LeaveRequest](#v1.LeaveRequest)                               | [LeaveResponse](#v1.LeaveResponse)                               | Leave is used to remove a node from the mesh. The node will be removed from the mesh and will no longer be able to query the mesh state or vote in elections.                                                                                                                                                                              |
| SubscribePeers          | [SubscribePeersRequest](#v1.SubscribePeersRequest)             | [PeerConfigurations](#v1.PeerConfigurations) stream              | SubscribePeers subscribes to the peer configuration for the given node. The node will receive updates to the peer configuration as it changes.                                                                                                                                                                                             |
| Apply                   | [RaftLogEntry](#v1.RaftLogEntry)                               | [RaftApplyResponse](#v1.RaftApplyResponse)                       | Apply is used by voting nodes to request a log entry be applied to the state machine. This is only available on the leader, and can only be called by nodes that are allowed to vote. This is only used by the built-in raft storage implementation.                                                                                       |
| GetStorageConfiguration | [StorageConfigurationRequest](#v1.StorageConfigurationRequest) | [StorageConfigurationResponse](#v1.StorageConfigurationResponse) | GetStorageConfiguration returns the current Storage configuration.                                                                                                                                                                                                                                                                         |

<div class="file-heading">

## v1/plugin.proto

[Top](#title)

</div>

### AllocateIPRequest

AllocateIPRequest is the message containing an IP allocation request.

| Field  | Type              | Label | Description                                                |
|--------|-------------------|-------|------------------------------------------------------------|
| nodeID | [string](#string) |       | node_id is the node that the IP should be allocated for.   |
| subnet | [string](#string) |       | subnet is the subnet that the IP should be allocated from. |

### AllocatedIP

AllocatedIP is the message containing an allocated IP.

| Field | Type              | Label | Description                                                     |
|-------|-------------------|-------|-----------------------------------------------------------------|
| ip    | [string](#string) |       | ip is the allocated IP. It should be returned in CIDR notation. |

### AuthenticationRequest

AuthenticationRequest is the message containing an authentication
request.

| Field        | Type                                                                         | Label    | Description                                                   |
|--------------|------------------------------------------------------------------------------|----------|---------------------------------------------------------------|
| headers      | [AuthenticationRequest.HeadersEntry](#v1.AuthenticationRequest.HeadersEntry) | repeated | headers are the headers of the request.                       |
| certificates | [bytes](#bytes)                                                              | repeated | certificates are the DER encoded certificates of the request. |

### AuthenticationRequest.HeadersEntry

| Field | Type              | Label | Description |
|-------|-------------------|-------|-------------|
| key   | [string](#string) |       |             |
| value | [string](#string) |       |             |

### AuthenticationResponse

AuthenticationResponse is the message containing an authentication
response.

| Field | Type              | Label | Description                             |
|-------|-------------------|-------|-----------------------------------------|
| id    | [string](#string) |       | id is the id of the authenticated user. |

### Event

Event is the message containing a watch event.

| Field | Type                                     | Label | Description                               |
|-------|------------------------------------------|-------|-------------------------------------------|
| type  | [Event.WatchEvent](#v1.Event.WatchEvent) |       | type is the type of the watch event.      |
| node  | [MeshNode](#v1.MeshNode)                 |       | node is the node that the event is about. |

### PluginConfiguration

PluginConfiguration is the message containing the configuration of a
plugin.

| Field  | Type                                              | Label | Description                                |
|--------|---------------------------------------------------|-------|--------------------------------------------|
| config | [google.protobuf.Struct](#google.protobuf.Struct) |       | Config is the configuration of the plugin. |

### PluginInfo

PluginInfo is the information of a plugin.

| Field        | Type                                                           | Label    | Description                                     |
|--------------|----------------------------------------------------------------|----------|-------------------------------------------------|
| name         | [string](#string)                                              |          | Name is the name of the plugin.                 |
| version      | [string](#string)                                              |          | Version is the version of the plugin.           |
| description  | [string](#string)                                              |          | Description is the description of the plugin.   |
| capabilities | [PluginInfo.PluginCapability](#v1.PluginInfo.PluginCapability) | repeated | Capabilities is the capabilities of the plugin. |

### ReleaseIPRequest

ReleaseIPRequest is the message containing an IP release request.

| Field  | Type              | Label | Description                                             |
|--------|-------------------|-------|---------------------------------------------------------|
| nodeID | [string](#string) |       | node_id is the node that the IP should be released for. |
| ip     | [string](#string) |       | ip is the IP that should be released.                   |

### Event.WatchEvent

WatchEvent is the type of a watch event.

| Name          | Number | Description                                                         |
|---------------|--------|---------------------------------------------------------------------|
| UNKNOWN       | 0      | UNKNOWN is the default value of WatchEvent.                         |
| NODE_JOIN     | 1      | NODE_JOIN indicates that a node has joined the cluster.             |
| NODE_LEAVE    | 2      | NODE_LEAVE indicates that a node has left the cluster.              |
| LEADER_CHANGE | 3      | LEADER_CHANGE indicates that the leader of the cluster has changed. |

### PluginInfo.PluginCapability

PluginCapability is the capabilities of a plugin.

| Name             | Number | Description                                                                                |
|------------------|--------|--------------------------------------------------------------------------------------------|
| UNKNOWN          | 0      | UNKNOWN is the default value of PluginCapability.                                          |
| STORAGE_PROVIDER | 1      | STORAGE_PROVIDER indicates that the plugin can provide storage and underlying consistency. |
| AUTH             | 2      | AUTH indicates that the plugin is an auth plugin.                                          |
| WATCH            | 3      | WATCH indicates that the plugin wants to receive watch events.                             |
| IPAMV4           | 4      | IPAMV4 indicates that the plugin is an IPv4 IPAM plugin.                                   |
| STORAGE_QUERIER  | 5      | STORAGE_QUERIER indicates a plugin that wants to interact with storage.                    |

### AuthPlugin

AuthPlugin is the service definition for a Webmesh auth plugin.

| Method Name  | Request Type                                       | Response Type                                        | Description                           |
|--------------|----------------------------------------------------|------------------------------------------------------|---------------------------------------|
| Authenticate | [AuthenticationRequest](#v1.AuthenticationRequest) | [AuthenticationResponse](#v1.AuthenticationResponse) | Authenticate authenticates a request. |

### IPAMPlugin

IPAMPlugin is the service definition for a Webmesh IPAM plugin.

| Method Name | Request Type                               | Response Type                                    | Description                          |
|-------------|--------------------------------------------|--------------------------------------------------|--------------------------------------|
| Allocate    | [AllocateIPRequest](#v1.AllocateIPRequest) | [AllocatedIP](#v1.AllocatedIP)                   | Allocate allocates an IP for a node. |
| Release     | [ReleaseIPRequest](#v1.ReleaseIPRequest)   | [.google.protobuf.Empty](#google.protobuf.Empty) | Release releases an IP for a node.   |

### Plugin

Plugin is the general service definition for a Webmesh plugin.

It must be implemented by all plugins.

| Method Name | Request Type                                     | Response Type                                    | Description                                                           |
|-------------|--------------------------------------------------|--------------------------------------------------|-----------------------------------------------------------------------|
| GetInfo     | [.google.protobuf.Empty](#google.protobuf.Empty) | [PluginInfo](#v1.PluginInfo)                     | GetInfo returns the information for the plugin.                       |
| Configure   | [PluginConfiguration](#v1.PluginConfiguration)   | [.google.protobuf.Empty](#google.protobuf.Empty) | Configure starts and configures the plugin.                           |
| Close       | [.google.protobuf.Empty](#google.protobuf.Empty) | [.google.protobuf.Empty](#google.protobuf.Empty) | Close closes the plugin. It is called when the node is shutting down. |

### StorageQuerierPlugin

StorageQuerierPlugin is the service definition for a Webmesh storage
querier plugin.

| Method Name   | Request Type                              | Response Type                           | Description                                                                                                                                                                                                                                                                                                                 |
|---------------|-------------------------------------------|-----------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| InjectQuerier | [QueryResponse](#v1.QueryResponse) stream | [QueryRequest](#v1.QueryRequest) stream | InjectQuerier is a stream opened by the node to faciliate read operations against the mesh state. The signature is misleading, but it is required to be able to stream the query results back to the node. The node will open a stream to the plugin and send a PluginQueryResult message for every query that is received. |

### WatchPlugin

WatchPlugin is the service definition for a Webmesh watch plugin.

| Method Name | Request Type       | Response Type                                    | Description                 |
|-------------|--------------------|--------------------------------------------------|-----------------------------|
| Emit        | [Event](#v1.Event) | [.google.protobuf.Empty](#google.protobuf.Empty) | Emit handles a watch event. |

<div class="file-heading">

## v1/storage_provider.proto

[Top](#title)

</div>

### AddObserverResponse

AddObserverResponse is the response object for the AddObserver RPC.

### AddVoterResponse

AddVoterResponse is the response object for the AddVoter RPC.

### BootstrapRequest

BootstrapRequest is the request object for the Bootstrap RPC.

### BootstrapResponse

BootstrapResponse is the response object for the Bootstrap RPC.

| Field  | Type                               | Label | Description                                              |
|--------|------------------------------------|-------|----------------------------------------------------------|
| status | [StorageStatus](#v1.StorageStatus) |       | Status is the status of the storage after the bootstrap. |

### DeleteValueRequest

DeleteValueRequest is the request object for the DeleteValue RPC.

| Field | Type            | Label | Description               |
|-------|-----------------|-------|---------------------------|
| key   | [bytes](#bytes) |       | Key is the key to delete. |

### DeleteValueResponse

DeleteValueResponse is the response object for the DeleteValue RPC.

### DemoteVoterResponse

DemoteVoterResponse is the response object for the DemoteVoter RPC.

### GetLeaderRequest

GetLeaderRequest is the request object for the GetLeader RPC.

### GetPeersRequest

GetPeersRequest is the request object for the GetPeers RPC.

### GetValueRequest

GetValueRequest is the request object for the GetValue RPC.

| Field | Type            | Label | Description                          |
|-------|-----------------|-------|--------------------------------------|
| key   | [bytes](#bytes) |       | Key is the key to get the value for. |

### GetValueResponse

GetValueResponse is the response object for the GetValue RPC.

| Field | Type                             | Label | Description                    |
|-------|----------------------------------|-------|--------------------------------|
| value | [StorageValue](#v1.StorageValue) |       | Value is the value of the key. |

### ListKeysRequest

ListKeysRequest is the request object for the ListValues RPC.

| Field  | Type            | Label | Description                              |
|--------|-----------------|-------|------------------------------------------|
| prefix | [bytes](#bytes) |       | Prefix is the prefix to list values for. |

### ListKeysResponse

ListKeysResponse is the response object for the ListValues RPC.

| Field | Type            | Label    | Description                                    |
|-------|-----------------|----------|------------------------------------------------|
| keys  | [bytes](#bytes) | repeated | Keys is the list of value keys for the prefix. |

### ListValuesRequest

ListValuesRequest is the request object for the ListValues RPC.

| Field  | Type            | Label | Description                              |
|--------|-----------------|-------|------------------------------------------|
| prefix | [bytes](#bytes) |       | Prefix is the prefix to list values for. |

### ListValuesResponse

ListValuesResponse is the response object for the ListValues RPC.

| Field  | Type                             | Label    | Description                                  |
|--------|----------------------------------|----------|----------------------------------------------|
| values | [StorageValue](#v1.StorageValue) | repeated | Values is the list of values for the prefix. |

### PrefixEvent

PrefixEvent is an event that is emitted when a value is added or removed

from the storage for a prefix.

| Field     | Type                                               | Label | Description                                   |
|-----------|----------------------------------------------------|-------|-----------------------------------------------|
| prefix    | [bytes](#bytes)                                    |       | Prefix is the prefix that the event is for.   |
| value     | [StorageValue](#v1.StorageValue)                   |       | Value is the value that was added or removed. |
| eventType | [PrefixEvent.EventType](#v1.PrefixEvent.EventType) |       | EventType is the type of event.               |

### PutValueRequest

PutValueRequest is the request object for the PutValue RPC.

| Field | Type                                                  | Label | Description                            |
|-------|-------------------------------------------------------|-------|----------------------------------------|
| value | [StorageValue](#v1.StorageValue)                      |       | Value is the value to put.             |
| ttl   | [google.protobuf.Duration](#google.protobuf.Duration) |       | TTL is the time to live for the value. |

### PutValueResponse

PutValueResponse is the response object for the PutValue RPC.

### RemoveServerResponse

RemoveServerResponse is the response object for the RemoveServer RPC.

### StoragePeer

StoragePeer is a server that is currently recognized by the storage
plugin.

| Field         | Type                               | Label | Description                                                                                                                                          |
|---------------|------------------------------------|-------|------------------------------------------------------------------------------------------------------------------------------------------------------|
| id            | [string](#string)                  |       | ID is the id of the server.                                                                                                                          |
| publicKey     | [string](#string)                  |       | public_key is the encoded public key of the server. This is not required for demotion or removal RPCs. Not all implementations need to support this. |
| address       | [string](#string)                  |       | Address is the address of the server. This is not required for demotion or removal RPCs.                                                             |
| clusterStatus | [ClusterStatus](#v1.ClusterStatus) |       | ClusterStatus is the status of the server. This is only applicable during a GetStatus RPC.                                                           |

### StoragePeers

StoragePeers is a list of servers that are currently recognized by the
storage plugin.

| Field | Type                           | Label    | Description                                                                                |
|-------|--------------------------------|----------|--------------------------------------------------------------------------------------------|
| peers | [StoragePeer](#v1.StoragePeer) | repeated | Peers is the list of servers that are currently recognized as peers by the storage plugin. |

### StorageStatus

StorageStatus is the response object for the StorageStatus RPC.

| Field         | Type                               | Label    | Description                                                                                                                        |
|---------------|------------------------------------|----------|------------------------------------------------------------------------------------------------------------------------------------|
| isWritable    | [bool](#bool)                      |          | IsWritable is true if the storage can currently be written to.                                                                     |
| clusterStatus | [ClusterStatus](#v1.ClusterStatus) |          | ClusterStatus is the status of the storage. The definitions applied to each status are implementation specific.                    |
| peers         | [StoragePeer](#v1.StoragePeer)     | repeated | Peers is the list of servers that are currently recognized as peers by the storage plugin. This should include the current server. |
| message       | [string](#string)                  |          | message is an implementation specific message that can be used to provide additional information about the storage status.         |

### StorageStatusRequest

StorageStatusRequest is the request object for the StorageStatus RPC.

### StorageValue

StorageValue is a value stored in the storage.

| Field | Type            | Label | Description                    |
|-------|-----------------|-------|--------------------------------|
| key   | [bytes](#bytes) |       | Key is the key of the value.   |
| value | [bytes](#bytes) |       | Value is the value of the key. |

### SubscribePrefixRequest

SubscribePrefixRequest is the request object for the SubscribePrefix
RPC.

| Field  | Type            | Label | Description                           |
|--------|-----------------|-------|---------------------------------------|
| prefix | [bytes](#bytes) |       | Prefix is the prefix to subscribe to. |

### PrefixEvent.EventType

| Name             | Number | Description                                                        |
|------------------|--------|--------------------------------------------------------------------|
| EventTypeUnknown | 0      | EventTypeUnknown is an unknown event type.                         |
| EventTypeUpdated | 1      | EventTypeUpdated is an event for when a value is added or updated. |
| EventTypeRemoved | 2      | EventTypeRemoved is an event for when a value is removed.          |

### StorageProviderPlugin

StorageProviderPlugin is the service definition for a Webmesh storage
provider.

| Method Name     | Request Type                                         | Response Type                                    | Description                                                                                                                                                                                                                                                                                                                                                              |
|-----------------|------------------------------------------------------|--------------------------------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| GetStatus       | [StorageStatusRequest](#v1.StorageStatusRequest)     | [StorageStatus](#v1.StorageStatus)               | GetStatus returns the status of the storage.                                                                                                                                                                                                                                                                                                                             |
| Bootstrap       | [BootstrapRequest](#v1.BootstrapRequest)             | [BootstrapResponse](#v1.BootstrapResponse)       | Bootstrap is called when the storage is created for the first time. It is assumed that this node has been elected as the leader of the cluster. FailedPrecondition should be returned if the storage is already bootstrapped.                                                                                                                                            |
| AddVoter        | [StoragePeer](#v1.StoragePeer)                       | [AddVoterResponse](#v1.AddVoterResponse)         | AddVoter adds a voter to the storage. The underlying implementation should ensure that the voter is added to the storage and that the storage is in a consistent state before returning.                                                                                                                                                                                 |
| AddObserver     | [StoragePeer](#v1.StoragePeer)                       | [AddObserverResponse](#v1.AddObserverResponse)   | AddObserver adds an observer to the storage. The underlying implementation should ensure that the observer is added to the storage and that the storage is in a consistent state before returning. If observers are not supported the underlying implementation can silently ignore this RPC, but it should keep track of the observer in the GetStatus RPC if possible. |
| DemoteVoter     | [StoragePeer](#v1.StoragePeer)                       | [DemoteVoterResponse](#v1.DemoteVoterResponse)   | DemoteVoter demotes a voter to an observer. The underlying implementation should ensure that the voter is demoted and that the storage is in a consistent state before returning. If observers are not supported the underlying implementation can silently ignore this RPC, but it should keep track of the observer in the GetStatus RPC if possible.                  |
| RemovePeer      | [StoragePeer](#v1.StoragePeer)                       | [RemoveServerResponse](#v1.RemoveServerResponse) | RemovePeer removes a peer from the storage. The underlying implementation should ensure that the server is removed and that the storage is in a consistent state before returning.                                                                                                                                                                                       |
| GetLeader       | [GetLeaderRequest](#v1.GetLeaderRequest)             | [StoragePeer](#v1.StoragePeer)                   | GetLeader returns the leader of the storage. Leader may be loosely defined by the implementation, but must be a node that can reliably be used to mutate the storage.                                                                                                                                                                                                    |
| GetPeers        | [GetPeersRequest](#v1.GetPeersRequest)               | [StoragePeers](#v1.StoragePeers)                 | GetPeers returns all peers of the storage. Peer status may be loosely defined by the implementation, but must correlate to nodes that can reliably be used to mutate the storage.                                                                                                                                                                                        |
| GetValue        | [GetValueRequest](#v1.GetValueRequest)               | [GetValueResponse](#v1.GetValueResponse)         | GetValue returns the value for a key.                                                                                                                                                                                                                                                                                                                                    |
| PutValue        | [PutValueRequest](#v1.PutValueRequest)               | [PutValueResponse](#v1.PutValueResponse)         | PutValue puts a value for a key.                                                                                                                                                                                                                                                                                                                                         |
| DeleteValue     | [DeleteValueRequest](#v1.DeleteValueRequest)         | [DeleteValueResponse](#v1.DeleteValueResponse)   | DeleteValue deletes a value for a key.                                                                                                                                                                                                                                                                                                                                   |
| ListKeys        | [ListKeysRequest](#v1.ListKeysRequest)               | [ListKeysResponse](#v1.ListKeysResponse)         | ListKeys lists all keys for a prefix.                                                                                                                                                                                                                                                                                                                                    |
| ListValues      | [ListValuesRequest](#v1.ListValuesRequest)           | [ListValuesResponse](#v1.ListValuesResponse)     | ListValues lists all values for a prefix.                                                                                                                                                                                                                                                                                                                                |
| SubscribePrefix | [SubscribePrefixRequest](#v1.SubscribePrefixRequest) | [PrefixEvent](#v1.PrefixEvent) stream            | SubscribePrefix subscribes to events for a prefix.                                                                                                                                                                                                                                                                                                                       |

<div class="file-heading">

## v1/webrtc.proto

[Top](#title)

</div>

### DataChannelOffer

DataChannelOffer is an offer for a data channel. Candidates

are sent after the offer is sent.

| Field       | Type              | Label    | Description                                      |
|-------------|-------------------|----------|--------------------------------------------------|
| offer       | [string](#string) |          | offer is the offer.                              |
| stunServers | [string](#string) | repeated | stun_servers is the list of STUN servers to use. |
| candidate   | [string](#string) |          | candidate is an ICE candidate.                   |

### StartDataChannelRequest

StartDataChannelRequest is a request to start a data channel.

The answer and candidate fields are populated after the offer

is received.

| Field     | Type              | Label | Description                                                                                                                             |
|-----------|-------------------|-------|-----------------------------------------------------------------------------------------------------------------------------------------|
| node_id   | [string](#string) |       | node_id is the ID of the node to send the data to.                                                                                      |
| proto     | [string](#string) |       | proto is the protocol of the traffic.                                                                                                   |
| dst       | [string](#string) |       | dst is the destination address of the traffic.                                                                                          |
| port      | [uint32](#uint32) |       | port is the destination port of the traffic. A port of 0 coupled with the udp protocol indicates forwarding to the WireGuard interface. |
| answer    | [string](#string) |       | answer is the answer to the offer.                                                                                                      |
| candidate | [string](#string) |       | candidate is an ICE candidate.                                                                                                          |

### WebRTC

WebRTC is a service for negotiating WebRTC connections to nodes in the
mesh.

It is typically run alongside a TURN server, however the server can be
configured

to use public STUN servers instead.

| Method Name        | Request Type                                                  | Response Type                                   | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          |
|--------------------|---------------------------------------------------------------|-------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| StartDataChannel   | [StartDataChannelRequest](#v1.StartDataChannelRequest) stream | [DataChannelOffer](#v1.DataChannelOffer) stream | StartDataChannel requests a new WebRTC connection to a node. The client speaks first with the request containing the node ID and where forwarded packets should be sent. The server responds with an offer and STUN servers to be used to establish a WebRTC connection. The client should then respond with an answer to the offer that matches the spec of the DataChannel.CHANNELS enum value. After the offer is accepted, the stream can be used to exchange ICE candidates to speed up the connection process. |
| StartSignalChannel | [WebRTCSignal](#v1.WebRTCSignal) stream                       | [WebRTCSignal](#v1.WebRTCSignal) stream         | StartSignalChannel starts a signaling channel to a remote node. This can be used to negotiate WebRTC connections both inside and outside of the mesh. Messages on the wire are proxied to the remote node.                                                                                                                                                                                                                                                                                                           |

## Scalar Value Types

| .proto Type | Notes                                                                                                                                           | C++    | Java       | Python      | Go       | C#         | PHP            | Ruby                           |
|-------------|-------------------------------------------------------------------------------------------------------------------------------------------------|--------|------------|-------------|----------|------------|----------------|--------------------------------|
| double      |                                                                                                                                                 | double | double     | float       | float64  | double     | float          | Float                          |
| float       |                                                                                                                                                 | float  | float      | float       | float32  | float      | float          | Float                          |
| int32       | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32  | int        | int         | int32    | int        | integer        | Bignum or Fixnum (as required) |
| int64       | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64  | long       | int/long    | int64    | long       | integer/string | Bignum                         |
| uint32      | Uses variable-length encoding.                                                                                                                  | uint32 | int        | int/long    | uint32   | uint       | integer        | Bignum or Fixnum (as required) |
| uint64      | Uses variable-length encoding.                                                                                                                  | uint64 | long       | int/long    | uint64   | ulong      | integer/string | Bignum or Fixnum (as required) |
| sint32      | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s.                            | int32  | int        | int         | int32    | int        | integer        | Bignum or Fixnum (as required) |
| sint64      | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s.                            | int64  | long       | int/long    | int64    | long       | integer/string | Bignum                         |
| fixed32     | Always four bytes. More efficient than uint32 if values are often greater than 2^28.                                                            | uint32 | int        | int         | uint32   | uint       | integer        | Bignum or Fixnum (as required) |
| fixed64     | Always eight bytes. More efficient than uint64 if values are often greater than 2^56.                                                           | uint64 | long       | int/long    | uint64   | ulong      | integer/string | Bignum                         |
| sfixed32    | Always four bytes.                                                                                                                              | int32  | int        | int         | int32    | int        | integer        | Bignum or Fixnum (as required) |
| sfixed64    | Always eight bytes.                                                                                                                             | int64  | long       | int/long    | int64    | long       | integer/string | Bignum                         |
| bool        |                                                                                                                                                 | bool   | boolean    | boolean     | bool     | bool       | boolean        | TrueClass/FalseClass           |
| string      | A string must always contain UTF-8 encoded or 7-bit ASCII text.                                                                                 | string | String     | str/unicode | string   | string     | string         | String (UTF-8)                 |
| bytes       | May contain any arbitrary sequence of bytes.                                                                                                    | string | ByteString | str         | \[\]byte | ByteString | string         | String (ASCII-8BIT)            |
