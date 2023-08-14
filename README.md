# Protocol Documentation

## Table of Contents

<div id="toc-container">

- [v1/raft.proto](#v1%2fraft.proto)
  - [<span class="badge">M</span>RaftApplyResponse](#v1.RaftApplyResponse)
  - [<span class="badge">M</span>RaftLogEntry](#v1.RaftLogEntry)
  - [<span class="badge">E</span>RaftCommandType](#v1.RaftCommandType)
- [v1/node.proto](#v1%2fnode.proto)
  - [<span class="badge">M</span>DataChannelNegotiation](#v1.DataChannelNegotiation)
  - [<span class="badge">M</span>GetStatusRequest](#v1.GetStatusRequest)
  - [<span class="badge">M</span>InterfaceMetrics](#v1.InterfaceMetrics)
  - [<span class="badge">M</span>JoinRequest](#v1.JoinRequest)
  - [<span class="badge">M</span>JoinResponse](#v1.JoinResponse)
  - [<span class="badge">M</span>LeaveRequest](#v1.LeaveRequest)
  - [<span class="badge">M</span>PeerMetrics](#v1.PeerMetrics)
  - [<span class="badge">M</span>SnapshotRequest](#v1.SnapshotRequest)
  - [<span class="badge">M</span>SnapshotResponse](#v1.SnapshotResponse)
  - [<span class="badge">M</span>Status](#v1.Status)
  - [<span class="badge">M</span>UpdateRequest](#v1.UpdateRequest)
  - [<span class="badge">M</span>UpdateResponse](#v1.UpdateResponse)
  - [<span class="badge">M</span>WireGuardPeer](#v1.WireGuardPeer)
  - [<span class="badge">E</span>ClusterStatus](#v1.ClusterStatus)
  - [<span class="badge">E</span>DataChannel](#v1.DataChannel)
  - [<span class="badge">E</span>Feature](#v1.Feature)
  - [<span class="badge">S</span>Node](#v1.Node)
- [v1/mesh.proto](#v1%2fmesh.proto)
  - [<span class="badge">M</span>GetNodeRequest](#v1.GetNodeRequest)
  - [<span class="badge">M</span>MeshEdge](#v1.MeshEdge)
  - [<span class="badge">M</span>MeshEdge.AttributesEntry](#v1.MeshEdge.AttributesEntry)
  - [<span class="badge">M</span>MeshEdges](#v1.MeshEdges)
  - [<span class="badge">M</span>MeshGraph](#v1.MeshGraph)
  - [<span class="badge">M</span>MeshNode](#v1.MeshNode)
  - [<span class="badge">M</span>NodeList](#v1.NodeList)
  - [<span class="badge">E</span>EdgeAttributes](#v1.EdgeAttributes)
  - [<span class="badge">S</span>Mesh](#v1.Mesh)
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
- [v1/network_acls.proto](#v1%2fnetwork_acls.proto)
  - [<span class="badge">M</span>NetworkACL](#v1.NetworkACL)
  - [<span class="badge">M</span>NetworkACLs](#v1.NetworkACLs)
  - [<span class="badge">M</span>NetworkAction](#v1.NetworkAction)
  - [<span class="badge">M</span>Route](#v1.Route)
  - [<span class="badge">M</span>Routes](#v1.Routes)
  - [<span class="badge">E</span>ACLAction](#v1.ACLAction)
- [v1/admin.proto](#v1%2fadmin.proto)
  - [<span class="badge">S</span>Admin](#v1.Admin)
- [v1/app.proto](#v1%2fapp.proto)
  - [<span class="badge">M</span>ConnectRequest](#v1.ConnectRequest)
  - [<span class="badge">M</span>ConnectResponse](#v1.ConnectResponse)
  - [<span class="badge">M</span>DisconnectRequest](#v1.DisconnectRequest)
  - [<span class="badge">M</span>DisconnectResponse](#v1.DisconnectResponse)
  - [<span class="badge">M</span>LeaveCampfireRequest](#v1.LeaveCampfireRequest)
  - [<span class="badge">M</span>LeaveCampfireResponse](#v1.LeaveCampfireResponse)
  - [<span class="badge">M</span>MetricsRequest](#v1.MetricsRequest)
  - [<span class="badge">M</span>MetricsResponse](#v1.MetricsResponse)
  - [<span class="badge">M</span>MetricsResponse.InterfacesEntry](#v1.MetricsResponse.InterfacesEntry)
  - [<span class="badge">M</span>PublishRequest](#v1.PublishRequest)
  - [<span class="badge">M</span>QueryRequest](#v1.QueryRequest)
  - [<span class="badge">M</span>QueryResponse](#v1.QueryResponse)
  - [<span class="badge">M</span>StartCampfireRequest](#v1.StartCampfireRequest)
  - [<span class="badge">M</span>StartCampfireResponse](#v1.StartCampfireResponse)
  - [<span class="badge">M</span>StatusRequest](#v1.StatusRequest)
  - [<span class="badge">M</span>StatusResponse](#v1.StatusResponse)
  - [<span class="badge">M</span>SubscribeRequest](#v1.SubscribeRequest)
  - [<span class="badge">M</span>SubscriptionEvent](#v1.SubscriptionEvent)
  - [<span class="badge">E</span>QueryRequest.QueryCommand](#v1.QueryRequest.QueryCommand)
  - [<span class="badge">E</span>StatusResponse.ConnectionStatus](#v1.StatusResponse.ConnectionStatus)
  - [<span class="badge">S</span>AppDaemon](#v1.AppDaemon)
- [v1/campfire.proto](#v1%2fcampfire.proto)
  - [<span class="badge">M</span>CampfireMessage](#v1.CampfireMessage)
  - [<span class="badge">E</span>CampfireMessage.MessageType](#v1.CampfireMessage.MessageType)
  - [<span class="badge">S</span>Campfire](#v1.Campfire)
- [v1/peer_discovery.proto](#v1%2fpeer_discovery.proto)
  - [<span class="badge">M</span>ListRaftPeersResponse](#v1.ListRaftPeersResponse)
  - [<span class="badge">M</span>RaftPeer](#v1.RaftPeer)
  - [<span class="badge">S</span>PeerDiscovery](#v1.PeerDiscovery)
- [v1/plugin.proto](#v1%2fplugin.proto)
  - [<span class="badge">M</span>AllocateIPRequest](#v1.AllocateIPRequest)
  - [<span class="badge">M</span>AllocatedIP](#v1.AllocatedIP)
  - [<span class="badge">M</span>AuthenticationRequest](#v1.AuthenticationRequest)
  - [<span class="badge">M</span>AuthenticationRequest.HeadersEntry](#v1.AuthenticationRequest.HeadersEntry)
  - [<span class="badge">M</span>AuthenticationResponse](#v1.AuthenticationResponse)
  - [<span class="badge">M</span>DataSnapshot](#v1.DataSnapshot)
  - [<span class="badge">M</span>Event](#v1.Event)
  - [<span class="badge">M</span>PluginConfiguration](#v1.PluginConfiguration)
  - [<span class="badge">M</span>PluginInfo](#v1.PluginInfo)
  - [<span class="badge">M</span>PluginQuery](#v1.PluginQuery)
  - [<span class="badge">M</span>PluginQueryResult](#v1.PluginQueryResult)
  - [<span class="badge">M</span>ReleaseIPRequest](#v1.ReleaseIPRequest)
  - [<span class="badge">M</span>StoreLogRequest](#v1.StoreLogRequest)
  - [<span class="badge">E</span>AllocateIPRequest.IPVersion](#v1.AllocateIPRequest.IPVersion)
  - [<span class="badge">E</span>PluginCapability](#v1.PluginCapability)
  - [<span class="badge">E</span>PluginQuery.QueryCommand](#v1.PluginQuery.QueryCommand)
  - [<span class="badge">E</span>WatchEvent](#v1.WatchEvent)
  - [<span class="badge">S</span>AuthPlugin](#v1.AuthPlugin)
  - [<span class="badge">S</span>IPAMPlugin](#v1.IPAMPlugin)
  - [<span class="badge">S</span>Plugin](#v1.Plugin)
  - [<span class="badge">S</span>StoragePlugin](#v1.StoragePlugin)
  - [<span class="badge">S</span>WatchPlugin](#v1.WatchPlugin)
- [v1/webrtc.proto](#v1%2fwebrtc.proto)
  - [<span class="badge">M</span>DataChannelOffer](#v1.DataChannelOffer)
  - [<span class="badge">M</span>StartDataChannelRequest](#v1.StartDataChannelRequest)
  - [<span class="badge">S</span>WebRTC](#v1.WebRTC)
- [Scalar Value Types](#scalar-value-types)

</div>

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

### RaftLogEntry

RaftLogEntry is the data of an entry in the Raft log.

| Field | Type                                                  | Label | Description                               |
|-------|-------------------------------------------------------|-------|-------------------------------------------|
| type  | [RaftCommandType](#v1.RaftCommandType)                |       | type is the type of the log entry.        |
| key   | [string](#string)                                     |       | key is the key of the log entry.          |
| value | [string](#string)                                     |       | value is the value of the log entry.      |
| ttl   | [google.protobuf.Duration](#google.protobuf.Duration) |       | ttl is the time to live of the log entry. |

### RaftCommandType

RaftCommandType is the type of command being sent to the

Raft log.

| Name    | Number | Description                                          |
|---------|--------|------------------------------------------------------|
| UNKNOWN | 0      | UNKNOWN is the unknown command type.                 |
| PUT     | 1      | PUT is the command for putting a key/value pair.     |
| DELETE  | 2      | DELETE is the command for deleting a key/value pair. |

<div class="file-heading">

## v1/node.proto

[Top](#title)

</div>

### DataChannelNegotiation

DataChannelNegotiation is the message for communicating data channels to
nodes.

| Field        | Type              | Label    | Description                                                         |
|--------------|-------------------|----------|---------------------------------------------------------------------|
| proto        | [string](#string) |          | proto is the protocol of the traffic.                               |
| src          | [string](#string) |          | src is the address of the client that initiated the request.        |
| dst          | [string](#string) |          | dst is the destination address of the traffic.                      |
| port         | [uint32](#uint32) |          | port is the destination port of the traffic.                        |
| offer        | [string](#string) |          | offer is the offer for the node to use as its local description.    |
| answer       | [string](#string) |          | answer is the answer for the node to use as its remote description. |
| candidate    | [string](#string) |          | candidate is an ICE candidate.                                      |
| stun_servers | [string](#string) | repeated | stun_servers is the list of STUN servers to use.                    |

### GetStatusRequest

GetStatusRequest is a request to get the status of a node.

| Field | Type              | Label | Description                                                                   |
|-------|-------------------|-------|-------------------------------------------------------------------------------|
| id    | [string](#string) |       | id is the ID of the node. If unset, the status of the local node is returned. |

### InterfaceMetrics

InterfaceMetrics is the metrics for the WireGuard interface on a node.

| Field                | Type                           | Label    | Description                                                    |
|----------------------|--------------------------------|----------|----------------------------------------------------------------|
| device_name          | [string](#string)              |          | device_name is the name of the device.                         |
| public_key           | [string](#string)              |          | public_key is the public key of the node.                      |
| address_v4           | [string](#string)              |          | address_v4 is the IPv4 address of the node.                    |
| address_v6           | [string](#string)              |          | address_v6 is the IPv6 address of the node.                    |
| type                 | [string](#string)              |          | type is the type of interface being used for wireguard.        |
| listen_port          | [int32](#int32)                |          | listen_port is the port wireguard is listening on.             |
| total_receive_bytes  | [uint64](#uint64)              |          | total_receive_bytes is the total number of bytes received.     |
| total_transmit_bytes | [uint64](#uint64)              |          | total_transmit_bytes is the total number of bytes transmitted. |
| num_peers            | [int32](#int32)                |          | num_peers is the number of peers connected to the node.        |
| peers                | [PeerMetrics](#v1.PeerMetrics) | repeated | peers are the per-peer statistics.                             |

### JoinRequest

JoinRequest is a request to join the cluster.

| Field               | Type                   | Label    | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    |
|---------------------|------------------------|----------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| id                  | [string](#string)      |          | id is the ID of the node.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      |
| public_key          | [string](#string)      |          | public_key is the public wireguard key of the node to broadcast to peers.                                                                                                                                                                                                                                                                                                                                                                                                                                                                      |
| raft_port           | [int32](#int32)        |          | raft_port is the Raft listen port of the node.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
| grpc_port           | [int32](#int32)        |          | grpc_port is the gRPC listen port of the node.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
| meshdns_port        | [int32](#int32)        |          | meshdns_port is the port the node wishes to advertise to offer DNS to other peers. a value of 0 indicates the node does not wish to offer DNS. ports are currently assumed to be UDP.                                                                                                                                                                                                                                                                                                                                                          |
| primary_endpoint    | [string](#string)      |          | primary_endpoint is a routable address for the node. If left unset, the node is assumed to be behind a NAT and not directly accessible.                                                                                                                                                                                                                                                                                                                                                                                                        |
| wireguard_endpoints | [string](#string)      | repeated | wireguard_endpoints is a list of WireGuard endpoints for the node.                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |
| zone_awareness_id   | [string](#string)      |          | zone_awareness_id is the zone awareness ID of the node.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        |
| assign_ipv4         | [bool](#bool)          |          | assign_ipv4 is whether an IPv4 address should be assigned to the node.                                                                                                                                                                                                                                                                                                                                                                                                                                                                         |
| prefer_raft_ipv6    | [bool](#bool)          |          | prefer_raft_ipv6 is whether IPv6 should be preferred over IPv4 for raft communication. This is only used if assign_ipv4 is true.                                                                                                                                                                                                                                                                                                                                                                                                               |
| as_voter            | [bool](#bool)          |          | as_voter is whether the node should receive a vote in elections. The request will be denied if the node is not allowed to vote.                                                                                                                                                                                                                                                                                                                                                                                                                |
| routes              | [string](#string)      | repeated | routes is a list of routes to advertise to peers. The request will be denied if the node is not allowed to put routes.                                                                                                                                                                                                                                                                                                                                                                                                                         |
| direct_peers        | [string](#string)      | repeated | direct_peers is a list of extra peers that should be connected to directly over ICE. The request will be denied if the node is not allowed to put data channels or edges. The default joining behavior creates non-ICE links between the caller and the joiner. If the caller has a primary endpoint, the joiner will link the caller to all other nodes with a primary endpoint. If the caller has a zone awareness ID, the joiner will link the caller to all other nodes with the same zone awareness ID that also have a primary endpoint. |
| features            | [Feature](#v1.Feature) | repeated | features is a list of features supported by the node that should be advertised to peers.                                                                                                                                                                                                                                                                                                                                                                                                                                                       |

### JoinResponse

JoinResponse is a response to a join request.

| Field        | Type                               | Label    | Description                                                                                                                                                                                                                                                    |
|--------------|------------------------------------|----------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| address_ipv4 | [string](#string)                  |          | address_ipv4 is the private IPv4 wireguard address of the node in CIDR format representing the network. This is only set if assign_ipv4 was set in the request or no network_ipv6 was provided. The bits are set to the network bits of the Mesh IPv4 network. |
| address_ipv6 | [string](#string)                  |          | address_ipv6 is the IPv6 network assigned to the node.                                                                                                                                                                                                         |
| network_ipv4 | [string](#string)                  |          | network_ipv4 is the IPv4 network of the Mesh.                                                                                                                                                                                                                  |
| network_ipv6 | [string](#string)                  |          | network_ipv6 is the IPv6 network of the Mesh.                                                                                                                                                                                                                  |
| peers        | [WireGuardPeer](#v1.WireGuardPeer) | repeated | peers is a list of wireguard peers to connect to.                                                                                                                                                                                                              |
| ice_servers  | [string](#string)                  | repeated | ice_servers is a list of public nodes that can be used to negotiate ICE connections if required. This may only be populated when one of the peers has the ICE flag set. This must be set if the requestor specifies direct_peers.                              |
| dns_servers  | [string](#string)                  | repeated | dns_servers is a list of peers offering DNS services.                                                                                                                                                                                                          |
| mesh_domain  | [string](#string)                  |          | mesh_domain is the domain of the mesh.                                                                                                                                                                                                                         |

### LeaveRequest

LeaveRequest is a request to leave the cluster.

| Field | Type              | Label | Description               |
|-------|-------------------|-------|---------------------------|
| id    | [string](#string) |       | id is the ID of the node. |

### PeerMetrics

PeerMetrics are the metrics for a node's peer.

| Field                 | Type              | Label    | Description                                                                         |
|-----------------------|-------------------|----------|-------------------------------------------------------------------------------------|
| public_key            | [string](#string) |          | public_key is the public key of the peer.                                           |
| endpoint              | [string](#string) |          | endpoint is the connected endpoint of the peer.                                     |
| persistent_keep_alive | [string](#string) |          | persistent_keep_alive is the persistent keep alive interval for the peer.           |
| last_handshake_time   | [string](#string) |          | last_handshake_time is the last handshake time for the peer.                        |
| allowed_ips           | [string](#string) | repeated | allowed_ips is the list of allowed IPs for the peer.                                |
| protocol_version      | [int64](#int64)   |          | protocol_version is the version of the wireguard protocol negotiated with the peer. |
| receive_bytes         | [uint64](#uint64) |          | receive_bytes is the bytes received from the peer.                                  |
| transmit_bytes        | [uint64](#uint64) |          | transmit_bytes is the bytes transmitted to the peer.                                |

### SnapshotRequest

SnapshotRequest is a request to create a snapshot. It is intentionally

empty for now as there are no options.

### SnapshotResponse

SnapshotResponse is a response to a snapshot request.

| Field          | Type              | Label | Description                                           |
|----------------|-------------------|-------|-------------------------------------------------------|
| last_log_index | [uint64](#uint64) |       | last_log_index is the last log index of the snapshot. |
| current_term   | [uint64](#uint64) |       | current_term is the current term of the snapshot.     |
| snapshot       | [bytes](#bytes)   |       | snapshot is the snapshot data.                        |

### Status

Status represents the status of a node.

| Field             | Type                                                    | Label    | Description                                                  |
|-------------------|---------------------------------------------------------|----------|--------------------------------------------------------------|
| id                | [string](#string)                                       |          | id is the ID of the node.                                    |
| version           | [string](#string)                                       |          | version is the version of the node.                          |
| commit            | [string](#string)                                       |          | commit is the commit of the node.                            |
| build_date        | [string](#string)                                       |          | build_date is the build date of the node.                    |
| uptime            | [string](#string)                                       |          | uptime is the uptime of the node.                            |
| started_at        | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |          | started_at is the time the node started.                     |
| features          | [Feature](#v1.Feature)                                  | repeated | features is the list of features currently enabled.          |
| cluster_status    | [ClusterStatus](#v1.ClusterStatus)                      |          | cluster_status is the status of the node in the cluster.     |
| current_leader    | [string](#string)                                       |          | current_leader is the current leader of the cluster.         |
| current_term      | [uint64](#uint64)                                       |          | current_term is the current term of the cluster.             |
| last_log_index    | [uint64](#uint64)                                       |          | last_log_index is the last log index of the cluster.         |
| last_applied      | [uint64](#uint64)                                       |          | last_applied is the last applied index of the cluster.       |
| interface_metrics | [InterfaceMetrics](#v1.InterfaceMetrics)                |          | interface_metrics are the metrics for the node's interfaces. |

### UpdateRequest

UpdateRequest contains most of the same fields as JoinRequest, but is

used to update the state of a node in the cluster.

| Field               | Type                   | Label    | Description                                                                                                                                                                           |
|---------------------|------------------------|----------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| id                  | [string](#string)      |          | id is the ID of the node.                                                                                                                                                             |
| public_key          | [string](#string)      |          | public_key is the public wireguard key of the node to broadcast to peers.                                                                                                             |
| raft_port           | [int32](#int32)        |          | raft_port is the Raft listen port of the node.                                                                                                                                        |
| grpc_port           | [int32](#int32)        |          | grpc_port is the gRPC listen port of the node.                                                                                                                                        |
| meshdns_port        | [int32](#int32)        |          | meshdns_port is the port the node wishes to advertise to offer DNS to other peers. a value of 0 indicates the node does not wish to offer DNS. ports are currently assumed to be UDP. |
| primary_endpoint    | [string](#string)      |          | primary_endpoint is a routable address for the node. If left unset, the node is assumed to be behind a NAT and not directly accessible.                                               |
| wireguard_endpoints | [string](#string)      | repeated | wireguard_endpoints is a list of WireGuard endpoints for the node.                                                                                                                    |
| zone_awareness_id   | [string](#string)      |          | zone_awareness_id is the zone awareness ID of the node.                                                                                                                               |
| as_voter            | [bool](#bool)          |          | as_voter is whether the node should receive a vote in elections. The request will be denied if the node is not allowed to vote.                                                       |
| routes              | [string](#string)      | repeated | routes is a list of routes to advertise to peers. The request will be denied if the node is not allowed to put routes.                                                                |
| features            | [Feature](#v1.Feature) | repeated | features is a list of features supported by the node that should be advertised to peers.                                                                                              |

### UpdateResponse

UpdateResponse is a response to an update request. It is currently
empty.

### WireGuardPeer

WireGuardPeer is a peer in the Wireguard network.

| Field               | Type              | Label    | Description                                                                       |
|---------------------|-------------------|----------|-----------------------------------------------------------------------------------|
| id                  | [string](#string) |          | id is the ID of the peer.                                                         |
| public_key          | [string](#string) |          | public_key is the public key of the peer.                                         |
| primary_endpoint    | [string](#string) |          | primary_endpoint is the primary endpoint of the peer.                             |
| wireguard_endpoints | [string](#string) | repeated | wireguard_endpoints are the WireGuard endpoints for the peer, if applicable.      |
| zone_awareness_id   | [string](#string) |          | zone_awareness_id is the zone awareness ID of the peer.                           |
| address_ipv4        | [string](#string) |          | address_ipv4 is the private IPv4 wireguard address of the peer.                   |
| address_ipv6        | [string](#string) |          | address_ipv6 is the private IPv6 wireguard address of the peer.                   |
| allowed_ips         | [string](#string) | repeated | allowed_ips is the list of allowed IPs for the peer.                              |
| allowed_routes      | [string](#string) | repeated | allowed_routes is the list of allowed routes for the peer.                        |
| ice                 | [bool](#bool)     |          | ice indicates whether the connection to this peer should be established over ICE. |

### ClusterStatus

ClusterStatus is the status of the node in the cluster.

| Name                   | Number | Description                                           |
|------------------------|--------|-------------------------------------------------------|
| CLUSTER_STATUS_UNKNOWN | 0      | CLUSTER_STATUS_UNKNOWN is the default status.         |
| CLUSTER_LEADER         | 1      | CLUSTER_LEADER is the status for the leader node.     |
| CLUSTER_VOTER          | 2      | CLUSTER_VOTER is the status for a voter node.         |
| CLUSTER_NON_VOTER      | 3      | CLUSTER_NON_VOTER is the status for a non-voter node. |

### DataChannel

DataChannel are the data channels used when communicating over ICE

with a node.

| Name        | Number | Description                                                                                                                                                                                                             |
|-------------|--------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| CHANNELS    | 0      | CHANNELS is the data channel used for negotiating new channels. This is the first channel that is opened. The ID of the channel should be 0.                                                                            |
| CONNECTIONS | 1      | CONNECTIONS is the data channel used for negotiating new connections. This is a channel that is opened for each incoming connection from a client. The ID should start at 0 and be incremented for each new connection. |

### Feature

Feature is a list of features supported by a node.

| Name             | Number | Description                                                                      |
|------------------|--------|----------------------------------------------------------------------------------|
| FEATURE_NONE     | 0      | FEATURE_NONE is the default feature set.                                         |
| NODES            | 1      | NODES is the feature for nodes. This is always supported.                        |
| LEADER_PROXY     | 2      | LEADER_PROXY is the feature for leader proxying.                                 |
| MESH_API         | 3      | MESH_API is the feature for the mesh API.                                        |
| ADMIN_API        | 4      | ADMIN_API is the feature for the admin API.                                      |
| PEER_DISCOVERY   | 5      | PEER_DISCOVERY is the feature for peer discovery.                                |
| METRICS          | 6      | METRICS is the feature for exposing metrics.                                     |
| ICE_NEGOTIATION  | 7      | ICE_NEGOTIATION is the feature for ICE negotiation.                              |
| TURN_SERVER      | 8      | TURN_SERVER is the feature for TURN server.                                      |
| MESH_DNS         | 9      | MESH_DNS is the feature for mesh DNS.                                            |
| FORWARD_MESH_DNS | 10     | FORWARD_MESH_DNS is the feature for forwarding mesh DNS lookups to other meshes. |

### Node

Node is the service exposed on every node in the mesh to communicate
network

information amongst themselves. Some methods are only available on the
currently

elected leader. This service can optionally be exposed on public
interfaces to allow

external users to query the mesh state, join as an observer, or
proxy/inspect traffic.

Nodes can optionally be configured to proxy requests to the leader. To
prefer the leader

handle the request when a non-leader can otherwise serve it, use the
"prefer-leader" header.

| Method Name          | Request Type                                                | Response Type                                               | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         |
|----------------------|-------------------------------------------------------------|-------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Join                 | [JoinRequest](#v1.JoinRequest)                              | [JoinResponse](#v1.JoinResponse)                            | Join is used to join a node to the mesh. The joining node will be added to the mesh as an observer, and will be able to query the mesh state, but will not be able to vote in elections. To join as a voter pass the as_voter flag.                                                                                                                                                                                                                                                                                 |
| Update               | [UpdateRequest](#v1.UpdateRequest)                          | [UpdateResponse](#v1.UpdateResponse)                        | Update is used by a node to update its state in the mesh. The node will be updated in the mesh and will be able to query the mesh state or vote in elections. Only non-empty fields will be updated. It is almost semantically equivalent to a join request with the same ID, but redefined to avoid confusion and to allow for expansion.                                                                                                                                                                          |
| Leave                | [LeaveRequest](#v1.LeaveRequest)                            | [.google.protobuf.Empty](#google.protobuf.Empty)            | Leave is used to remove a node from the mesh. The node will be removed from the mesh and will no longer be able to query the mesh state or vote in elections.                                                                                                                                                                                                                                                                                                                                                       |
| GetStatus            | [GetStatusRequest](#v1.GetStatusRequest)                    | [Status](#v1.Status)                                        | GetStatus gets the status of a node in the cluster.                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
| Snapshot             | [SnapshotRequest](#v1.SnapshotRequest)                      | [SnapshotResponse](#v1.SnapshotResponse)                    | Snapshot is used to create a snapshot of the current state of the mesh. The snapshot can be used to restore the mesh state.                                                                                                                                                                                                                                                                                                                                                                                         |
| Apply                | [RaftLogEntry](#v1.RaftLogEntry)                            | [RaftApplyResponse](#v1.RaftApplyResponse)                  | Apply is used by voting nodes to request a log entry be applied to the state machine. This is only available on the leader, and can only be called by nodes that are allowed to vote.                                                                                                                                                                                                                                                                                                                               |
| NegotiateDataChannel | [DataChannelNegotiation](#v1.DataChannelNegotiation) stream | [DataChannelNegotiation](#v1.DataChannelNegotiation) stream | NegotiateDataChannel is used to negotiate a WebRTC connection between a webmesh client and a node in the cluster. The handling server will send the target node the source address, the destination for traffic, and STUN/TURN servers to use for the negotiation. The node responds with an offer to be forwarded to the client. When the handler receives an answer from the client, it forwards it to the node. Once the node receives the answer, the stream can optionally be used to exchange ICE candidates. |

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

### MeshNode

MeshNode is a node that has been registered with the controller.

| Field               | Type                                                    | Label    | Description                                                                                                                                                      |
|---------------------|---------------------------------------------------------|----------|------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| id                  | [string](#string)                                       |          | id is the ID of the node.                                                                                                                                        |
| primary_endpoint    | [string](#string)                                       |          | primary_endpoint is the primary endpoint of the node.                                                                                                            |
| wireguard_endpoints | [string](#string)                                       | repeated | wireguard_endpoints is a list of WireGuard endpoints for the node.                                                                                               |
| zone_awareness_id   | [string](#string)                                       |          | zone_awareness_id is the zone awareness ID of the node.                                                                                                          |
| raft_port           | [int32](#int32)                                         |          | raft_port is the Raft listen port of the node.                                                                                                                   |
| grpc_port           | [int32](#int32)                                         |          | grpc_port is the gRPC listen port of the node.                                                                                                                   |
| meshdns_port        | [int32](#int32)                                         |          | meshdns_port is the port the node advertises to offer DNS to other peers. a value of 0 means the node does not offer DNS. ports are currently assumed to be UDP. |
| public_key          | [string](#string)                                       |          | public_key is the public key of the node.                                                                                                                        |
| private_ipv4        | [string](#string)                                       |          | private_ipv4 is the private IPv4 address of the node.                                                                                                            |
| private_ipv6        | [string](#string)                                       |          | private_ipv6 is the private IPv6 address of the node.                                                                                                            |
| features            | [Feature](#v1.Feature)                                  | repeated | features is a list of features supported by the node.                                                                                                            |
| updated_at          | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |          | updated_at is the last time the node joined the cluster.                                                                                                         |
| created_at          | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |          | created_at is the creation time for the node.                                                                                                                    |
| cluster_status      | [ClusterStatus](#v1.ClusterStatus)                      |          | cluster_status is the status of the node in the cluster.                                                                                                         |

### NodeList

NodeList is a list of nodes.

| Field | Type                     | Label    | Description                 |
|-------|--------------------------|----------|-----------------------------|
| nodes | [MeshNode](#v1.MeshNode) | repeated | nodes is the list of nodes. |

### EdgeAttributes

EdgeAttributes are pre-defined edge attributes. They should

be used as their string values.

| Name                   | Number | Description                                          |
|------------------------|--------|------------------------------------------------------|
| EDGE_ATTRIBUTE_UNKNOWN | 0      | EDGE_ATTRIBUTE_UNKNOWN is an unknown edge attribute. |
| EDGE_ATTRIBUTE_ICE     | 1      | EDGE_ATTRIBUTE_ICE is an ICE edge attribute.         |

### Mesh

Mesh is a service that can optionally be exposed by a node. It provides
methods for

interfacing with the webmesh from the outside. Some methods are only
available on the

leader. Nodes can enable the leader proxy to expose the leader's Mesh
service.

| Method Name  | Request Type                                     | Response Type              | Description                                                                                                |
|--------------|--------------------------------------------------|----------------------------|------------------------------------------------------------------------------------------------------------|
| GetNode      | [GetNodeRequest](#v1.GetNodeRequest)             | [MeshNode](#v1.MeshNode)   | GetNode gets a node by ID.                                                                                 |
| ListNodes    | [.google.protobuf.Empty](#google.protobuf.Empty) | [NodeList](#v1.NodeList)   | ListNodes lists all nodes.                                                                                 |
| GetMeshGraph | [.google.protobuf.Empty](#google.protobuf.Empty) | [MeshGraph](#v1.MeshGraph) | GetMeshGraph fetches the mesh graph. It returns a list of nodes, edges, and a rendering in the dot format. |

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

| Field         | Type                             | Label | Description                                                                 |
|---------------|----------------------------------|-------|-----------------------------------------------------------------------------|
| resource      | [RuleResource](#v1.RuleResource) |       | resource is the resource on which the action is performed.                  |
| resource_name | [string](#string)                |       | resource_name is the name of the resource on which the action is performed. |
| verb          | [RuleVerb](#v1.RuleVerb)         |       | verb is the verb that is performed on the resource.                         |

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

| Field          | Type                             | Label    | Description                                                             |
|----------------|----------------------------------|----------|-------------------------------------------------------------------------|
| resources      | [RuleResource](#v1.RuleResource) | repeated | resources is the resources to which the rule applies.                   |
| resource_names | [string](#string)                | repeated | resource_names is the list of resource names to which the rule applies. |
| verbs          | [RuleVerb](#v1.RuleVerb)         | repeated | verbs is the list of verbs that apply to the resource.                  |

### Subject

Subject is a subject to which a role can be bound.

| Field | Type                           | Label | Description                      |
|-------|--------------------------------|-------|----------------------------------|
| name  | [string](#string)              |       | name is the name of the subject. |
| type  | [SubjectType](#v1.SubjectType) |       | type is the type of the subject. |

### RuleResource

RuleResource is the resource type for a rule.

| Name                   | Number | Description                                                                                                    |
|------------------------|--------|----------------------------------------------------------------------------------------------------------------|
| RESOURCE_UNKNOWN       | 0      | RESOURCE_UNKNOWN is an unknown resource.                                                                       |
| RESOURCE_VOTES         | 1      | RESOURCE_VOTES is the resource for voting in raft elections. The only verb evaluated for this resource is PUT. |
| RESOURCE_ROLES         | 2      | RESOURCE_ROLES is the resource for managing roles.                                                             |
| RESOURCE_ROLE_BINDINGS | 3      | RESOURCE_ROLE_BINDINGS is the resource for managing role bindings.                                             |
| RESOURCE_GROUPS        | 4      | RESOURCE_GROUPS is the resource for managing groups.                                                           |
| RESOURCE_NETWORK_ACLS  | 5      | RESOURCE_NETWORK_ACLS is the resource for managing network ACLs.                                               |
| RESOURCE_ROUTES        | 6      | RESOURCE_ROUTES is the resource for managing routes.                                                           |
| RESOURCE_DATA_CHANNELS | 7      | RESOURCE_DATA_CHANNELS is the resource for creating data channels.                                             |
| RESOURCE_EDGES         | 8      | RESOURCE_EDGES is the resource for managing edges.                                                             |
| RESOURCE_ALL           | 999    | RESOURCE_ALL is a wildcard resource that matches all resources.                                                |

### RuleVerb

RuleVerb is the verb type for a rule.

| Name         | Number | Description                                                                                                                                                                                                                  |
|--------------|--------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| VERB_UNKNOWN | 0      | VERB_UNKNOWN is an unknown verb.                                                                                                                                                                                             |
| VERB_PUT     | 1      | VERB_PUT is the verb for creating or updating a resource.                                                                                                                                                                    |
| VERB_DELETE  | 3      | VERB_GET is the verb for getting a resource. It is currently not used by the system since all nodes have read access to all resources. Commented out for now. VERB_GET = 2; VERB_DELETE is the verb for deleting a resource. |
| VERB_ALL     | 999    | VERB_ALL is a wildcard verb that matches all verbs.                                                                                                                                                                          |

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

## v1/network_acls.proto

[Top](#title)

</div>

### NetworkACL

NetworkACL is a network ACL.

| Field             | Type                       | Label    | Description                                                                                                                                 |
|-------------------|----------------------------|----------|---------------------------------------------------------------------------------------------------------------------------------------------|
| name              | [string](#string)          |          | name is the name of the ACL.                                                                                                                |
| priority          | [int32](#int32)            |          | priority is the priority of the ACL. ACLs with higher priority are evaluated first.                                                         |
| action            | [ACLAction](#v1.ACLAction) |          | action is the action to take when a request matches the ACL.                                                                                |
| source_nodes      | [string](#string)          | repeated | source_nodes is a list of source nodes to match against. If empty, all nodes are matched. Groups can be specified with the prefix "group:". |
| destination_nodes | [string](#string)          | repeated | destination_nodes is a list of destination nodes to match against. If empty, all nodes are matched.                                         |
| source_cidrs      | [string](#string)          | repeated | source_cidrs is a list of source CIDRs to match against. If empty, all CIDRs are matched.                                                   |
| destination_cidrs | [string](#string)          | repeated | destination_cidrs is a list of destination CIDRs to match against. If empty, all CIDRs are matched.                                         |
| protocols         | [string](#string)          | repeated | protocols is a list of protocols to match against. If empty, all protocols are matched. Protocols can be specified by name or number.       |
| ports             | [uint32](#uint32)          | repeated | ports is a list of ports to match against. If empty, all ports are matched.                                                                 |

### NetworkACLs

NetworkACLs is a list of network ACLs.

| Field | Type                         | Label    | Description                        |
|-------|------------------------------|----------|------------------------------------|
| items | [NetworkACL](#v1.NetworkACL) | repeated | items is the list of network ACLs. |

### NetworkAction

NetworkAction is an action that can be performed on a network resource.
It is used

by implementations to evaluate network ACLs.

| Field    | Type              | Label | Description                                     |
|----------|-------------------|-------|-------------------------------------------------|
| src_node | [string](#string) |       | src_node is the source node of the action.      |
| src_cidr | [string](#string) |       | src_cidr is the source CIDR of the action.      |
| dst_node | [string](#string) |       | dst_node is the destination node of the action. |
| dst_cidr | [string](#string) |       | dst_cidr is the destination CIDR of the action. |
| protocol | [string](#string) |       | protocol is the protocol of the action.         |
| port     | [uint32](#uint32) |       | port is the port of the action.                 |

### Route

Route is a route that is broadcasted by one or more nodes.

| Field             | Type              | Label    | Description                                                                                    |
|-------------------|-------------------|----------|------------------------------------------------------------------------------------------------|
| name              | [string](#string) |          | name is the name of the route.                                                                 |
| node              | [string](#string) |          | node is the node that broadcasts the route. A group can be specified with the prefix "group:". |
| destination_cidrs | [string](#string) | repeated | destination_cidrs are the destination CIDRs of the route.                                      |
| next_hop_node     | [string](#string) |          | next_hop_node is an optional node that is used as the next hop for the route.                  |

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

## v1/app.proto

[Top](#title)

</div>

### ConnectRequest

ConnectRequest is sent by the application to the node to establish a

connection to a mesh. This message will eventually contain unique

identifiers to allow creating connections to multiple meshes.

| Field             | Type                                              | Label | Description                                                                                                                          |
|-------------------|---------------------------------------------------|-------|--------------------------------------------------------------------------------------------------------------------------------------|
| config            | [google.protobuf.Struct](#google.protobuf.Struct) |       | Config is used to override any defaults configured on the node.                                                                      |
| disable_bootstrap | [bool](#bool)                                     |       | Disable bootstrap tells a node that is otherwise configured to bootstrap to not bootstrap for this connection.                       |
| campfire_uri      | [string](#string)                                 |       | Campfire URI is the campfire URI to join. This implies that the node should join the mesh that the campfire is in and not bootstrap. |

### ConnectResponse

ConnectResponse is returned by the Connect RPC.

| Field       | Type              | Label | Description                                   |
|-------------|-------------------|-------|-----------------------------------------------|
| node_id     | [string](#string) |       | node id is the unique identifier of the node. |
| mesh_domain | [string](#string) |       | mesh domain is the domain of the mesh.        |
| ipv4        | [string](#string) |       | ipv4 is the IPv4 address of the node.         |
| ipv6        | [string](#string) |       | ipv6 is the IPv6 address of the node.         |

### DisconnectRequest

DisconnectRequest is sent by the application to the node to disconnect

from a mesh. This message will eventually contain unique identifiers

for allowing the application to disconnect from a specific mesh.

### DisconnectResponse

DisconnectResponse is returned by the Disconnect RPC.

### LeaveCampfireRequest

LeaveCampfire is sent by the application to the node to leave a
campfire.

| Field    | Type              | Label | Description                                          |
|----------|-------------------|-------|------------------------------------------------------|
| camp_url | [string](#string) |       | CampURL is the camp:// URL of the campfire to leave. |

### LeaveCampfireResponse

LeaveCampfireResponse is returned by the LeaveCampfire RPC.

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

### PublishRequest

PublishRequest is sent by the application to the node to publish events.

This currently only supports database events.

| Field | Type                                                  | Label | Description                                                             |
|-------|-------------------------------------------------------|-------|-------------------------------------------------------------------------|
| key   | [string](#string)                                     |       | key is the key of the event.                                            |
| value | [string](#string)                                     |       | value is the value of the event. This will be the raw value of the key. |
| ttl   | [google.protobuf.Duration](#google.protobuf.Duration) |       | ttl is the time for the event to live in the database.                  |

### QueryRequest

QueryRequest is sent by the application to the node to query the mesh
for

information.

| Field   | Type                                                       | Label | Description                              |
|---------|------------------------------------------------------------|-------|------------------------------------------|
| command | [QueryRequest.QueryCommand](#v1.QueryRequest.QueryCommand) |       | command is the command of the query.     |
| query   | [string](#string)                                          |       | query is the key or prefix of the query. |

### QueryResponse

QueryResponse is the message containing a mesh query result.

| Field | Type              | Label    | Description                                                                                                                                                            |
|-------|-------------------|----------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| key   | [string](#string) |          | key is the key of the query. For GET and ITER queries it will be the current key. For LIST queries it will be the prefix.                                              |
| value | [string](#string) | repeated | value is the value of the query. For GET and ITER queries it will be the value of the current key. For LIST queries it will be the list of keys that match the prefix. |
| error | [string](#string) |          | error is an error that occurred during the query. At the end of an ITER query it will be set to "EOF" to indicate that the iteration is complete.                      |

### StartCampfireRequest

StartCampfire is sent by the application to the node to start a
campfire.

If the URL is empty, the node will start a new campfire using it's own

TURN server.

| Field    | Type              | Label | Description                                 |
|----------|-------------------|-------|---------------------------------------------|
| camp_url | [string](#string) |       | CampURL is the camp:// URL of the campfire. |

### StartCampfireResponse

StartCampfireResponse is returned by the StartCampfire RPC. This should

eventually return a camp:// URL that can be used to invite other nodes

to the campfire.

| Field    | Type              | Label | Description                                 |
|----------|-------------------|-------|---------------------------------------------|
| camp_url | [string](#string) |       | CampURL is the camp:// URL of the campfire. |

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

### SubscribeRequest

SubscribeRequest is sent by the application to the node to subscribe to

events. This currently only supports database events.

| Field  | Type              | Label | Description                                         |
|--------|-------------------|-------|-----------------------------------------------------|
| prefix | [string](#string) |       | prefix is the prefix of the events to subscribe to. |

### SubscriptionEvent

SubscriptionEvent is a message containing a subscription event.

| Field | Type              | Label | Description                                                             |
|-------|-------------------|-------|-------------------------------------------------------------------------|
| key   | [string](#string) |       | key is the key of the event.                                            |
| value | [string](#string) |       | value is the value of the event. This will be the raw value of the key. |

### QueryRequest.QueryCommand

QueryCommand is the type of the query.

| Name | Number | Description                                                       |
|------|--------|-------------------------------------------------------------------|
| GET  | 0      | GET is the command to get a value.                                |
| LIST | 1      | LIST is the command to list keys with an optional prefix.         |
| ITER | 2      | ITER is the command to iterate over keys with an optional prefix. |

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

| Method Name   | Request Type                                     | Response Type                                      | Description                                                                                                                                         |
|---------------|--------------------------------------------------|----------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------|
| Connect       | [ConnectRequest](#v1.ConnectRequest)             | [ConnectResponse](#v1.ConnectResponse)             | Connect is used to establish a connection between the node and a mesh. The provided struct is used to override any defaults configured on the node. |
| Disconnect    | [DisconnectRequest](#v1.DisconnectRequest)       | [DisconnectResponse](#v1.DisconnectResponse)       | Disconnect is used to disconnect the node from the mesh.                                                                                            |
| StartCampfire | [StartCampfireRequest](#v1.StartCampfireRequest) | [StartCampfireResponse](#v1.StartCampfireResponse) | StartCampfire is used to start a campfire.                                                                                                          |
| LeaveCampfire | [LeaveCampfireRequest](#v1.LeaveCampfireRequest) | [LeaveCampfireResponse](#v1.LeaveCampfireResponse) | LeaveCampfire is used to leave a campfire.                                                                                                          |
| Query         | [QueryRequest](#v1.QueryRequest)                 | [QueryResponse](#v1.QueryResponse) stream          | Query is used to query the mesh for information.                                                                                                    |
| Metrics       | [MetricsRequest](#v1.MetricsRequest)             | [MetricsResponse](#v1.MetricsResponse)             | Metrics is used to retrieve interface metrics from the node.                                                                                        |
| Status        | [StatusRequest](#v1.StatusRequest)               | [StatusResponse](#v1.StatusResponse)               | Status is used to retrieve the status of the node.                                                                                                  |
| Subscribe     | [SubscribeRequest](#v1.SubscribeRequest)         | [SubscriptionEvent](#v1.SubscriptionEvent) stream  | Subscribe is used to subscribe to events in the mesh database.                                                                                      |
| Publish       | [PublishRequest](#v1.PublishRequest)             | [.google.protobuf.Empty](#google.protobuf.Empty)   | Publish is used to publish events to the mesh database. A restricted set of keys are allowed to be published to.                                    |

<div class="file-heading">

## v1/campfire.proto

[Top](#title)

</div>

### CampfireMessage

CampfireMessage is used to send messages between peers.

| Field  | Type                                                           | Label | Description                                                                                                                                                              |
|--------|----------------------------------------------------------------|-------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| type   | [CampfireMessage.MessageType](#v1.CampfireMessage.MessageType) |       | The type of the message.                                                                                                                                                 |
| id     | [string](#string)                                              |       | id is a unique identifier for the client. It is used to demultiplex messages from multiple clients. It should remain constant for the lifecycle of a WebRTC negotiation. |
| lufrag | [string](#string)                                              |       | The sending ufrag of the message.                                                                                                                                        |
| lpwd   | [string](#string)                                              |       | The sending password of the message.                                                                                                                                     |
| rufrag | [string](#string)                                              |       | The receiving ufrag of the message.                                                                                                                                      |
| rpwd   | [string](#string)                                              |       | The receiving password of the message.                                                                                                                                   |
| data   | [bytes](#bytes)                                                |       | The data of the message. It is recommended to be encrypted with a pre-shared key before sending.                                                                         |

### CampfireMessage.MessageType

MessageType is used to indicate the type of a CampfireMessage.

| Name      | Number | Description                                                                                                 |
|-----------|--------|-------------------------------------------------------------------------------------------------------------|
| UNKNOWN   | 0      | UNKNOWN is the default value and should not be used.                                                        |
| ANNOUNCE  | 1      | ANNOUNCE is used to announce presence at a campfire. This is only required when waiting for others to join. |
| OFFER     | 2      | OFFER is used to offer a WebRTC connection to another peer.                                                 |
| ANSWER    | 3      | ANSWER is used to answer a WebRTC connection from another peer.                                             |
| CANDIDATE | 4      | CANDIDATE is used to send a WebRTC candidate to another peer.                                               |

### Campfire

Campfire is the service definition for Campfire traffic. The protocol

is intended to be served over UDP alongside regular STUN and/or TURN

traffic, but can be used over any reliable transport.

| Method Name     | Request Type                                     | Response Type                                    | Description                                                                                                 |
|-----------------|--------------------------------------------------|--------------------------------------------------|-------------------------------------------------------------------------------------------------------------|
| Announce        | [CampfireMessage](#v1.CampfireMessage)           | [.google.protobuf.Empty](#google.protobuf.Empty) | Announce is used to announce presence at a campfire. This is only required when waiting for others to join. |
| SendOffer       | [CampfireMessage](#v1.CampfireMessage)           | [.google.protobuf.Empty](#google.protobuf.Empty) | SendOffer is used to send a WebRTC offer to another peer.                                                   |
| SendAnswer      | [CampfireMessage](#v1.CampfireMessage)           | [.google.protobuf.Empty](#google.protobuf.Empty) | SendAnswer is used to send a WebRTC answer to another peer.                                                 |
| SendCandidate   | [CampfireMessage](#v1.CampfireMessage)           | [.google.protobuf.Empty](#google.protobuf.Empty) | SendCandidate is used to send a WebRTC candidate to another peer.                                           |
| ReceiveMessages | [.google.protobuf.Empty](#google.protobuf.Empty) | [CampfireMessage](#v1.CampfireMessage) stream    | ReceiveMessages is used to receive messages from other peers.                                               |

<div class="file-heading">

## v1/peer_discovery.proto

[Top](#title)

</div>

### ListRaftPeersResponse

ListRaftPeersResponse is the response to ListPeers.

| Field | Type                     | Label    | Description                 |
|-------|--------------------------|----------|-----------------------------|
| peers | [RaftPeer](#v1.RaftPeer) | repeated | Peers is the list of peers. |

### RaftPeer

RaftPeer is a peer in the Raft cluster.

| Field   | Type              | Label | Description                                     |
|---------|-------------------|-------|-------------------------------------------------|
| id      | [string](#string) |       | ID is the ID of the peer.                       |
| address | [string](#string) |       | Address is the public gRPC address of the peer. |
| voter   | [bool](#bool)     |       | Voter is whether the peer is a voter.           |
| leader  | [bool](#bool)     |       | Leader is whether the peer is the leader.       |

### PeerDiscovery

PeerDiscovery is the service that provides peer discovery. This is a
service

that can optionally be exposed by nodes in the mesh to provide peer
discovery

to other nodes. Alternative methods of peer discovery can be used, such
as

static files or DNS, but running one or more publicly accessible nodes
with this

service registered is the simplest method.

It only makes sense to expose this service on a public address on a
member of the

Raft cluster. It is not necessary to expose this service on every node
in the mesh.

| Method Name | Request Type                                     | Response Type                                      | Description                                                           |
|-------------|--------------------------------------------------|----------------------------------------------------|-----------------------------------------------------------------------|
| ListPeers   | [.google.protobuf.Empty](#google.protobuf.Empty) | [ListRaftPeersResponse](#v1.ListRaftPeersResponse) | ListPeers returns a list of public peers currently known to the mesh. |

<div class="file-heading">

## v1/plugin.proto

[Top](#title)

</div>

### AllocateIPRequest

AllocateIPRequest is the message containing an IP allocation request.

| Field   | Type                                                           | Label | Description                                                |
|---------|----------------------------------------------------------------|-------|------------------------------------------------------------|
| node_id | [string](#string)                                              |       | node_id is the node that the IP should be allocated for.   |
| subnet  | [string](#string)                                              |       | subnet is the subnet that the IP should be allocated from. |
| version | [AllocateIPRequest.IPVersion](#v1.AllocateIPRequest.IPVersion) |       | version is the IP version that should be allocated.        |

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

### DataSnapshot

DataSnapshot is the message containing a snapshot of the data.

| Field | Type              | Label | Description                          |
|-------|-------------------|-------|--------------------------------------|
| term  | [uint64](#uint64) |       | term is the term of the log entry.   |
| index | [uint64](#uint64) |       | index is the index of the log entry. |
| data  | [bytes](#bytes)   |       | data is the snapshot of the data.    |

### Event

Event is the message containing a watch event.

| Field | Type                         | Label | Description                               |
|-------|------------------------------|-------|-------------------------------------------|
| type  | [WatchEvent](#v1.WatchEvent) |       | type is the type of the watch event.      |
| node  | [MeshNode](#v1.MeshNode)     |       | node is the node that the event is about. |

### PluginConfiguration

PluginConfiguration is the message containing the configuration of a
plugin.

| Field  | Type                                              | Label | Description                                |
|--------|---------------------------------------------------|-------|--------------------------------------------|
| config | [google.protobuf.Struct](#google.protobuf.Struct) |       | Config is the configuration of the plugin. |

### PluginInfo

PluginInfo is the information of a plugin.

| Field        | Type                                     | Label    | Description                                     |
|--------------|------------------------------------------|----------|-------------------------------------------------|
| name         | [string](#string)                        |          | Name is the name of the plugin.                 |
| version      | [string](#string)                        |          | Version is the version of the plugin.           |
| description  | [string](#string)                        |          | Description is the description of the plugin.   |
| capabilities | [PluginCapability](#v1.PluginCapability) | repeated | Capabilities is the capabilities of the plugin. |

### PluginQuery

PluginQuery is the message containing a storage query. It contains

a request ID that is used to correlate the query with the result.

| Field   | Type                                                     | Label | Description                              |
|---------|----------------------------------------------------------|-------|------------------------------------------|
| id      | [string](#string)                                        |       | id is the ID of the query.               |
| command | [PluginQuery.QueryCommand](#v1.PluginQuery.QueryCommand) |       | command is the command of the query.     |
| query   | [string](#string)                                        |       | query is the key or prefix of the query. |

### PluginQueryResult

PluginQueryResult is the message containing a storage query result. It
contains

a request ID that is used to correlate the query with the result.

| Field | Type              | Label    | Description                                                                                                                                                            |
|-------|-------------------|----------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| id    | [string](#string) |          | id is the ID of the query.                                                                                                                                             |
| key   | [string](#string) |          | key is the key of the query. For GET and ITER queries it will be the current key. For LIST queries it will be the prefix.                                              |
| value | [string](#string) | repeated | value is the value of the query. For GET and ITER queries it will be the value of the current key. For LIST queries it will be the list of keys that match the prefix. |
| error | [string](#string) |          | error is an error that occurred during the query. At the end of an ITER query it will be set to "EOF" to indicate that the iteration is complete.                      |

### ReleaseIPRequest

ReleaseIPRequest is the message containing an IP release request.

| Field   | Type              | Label | Description                                             |
|---------|-------------------|-------|---------------------------------------------------------|
| node_id | [string](#string) |       | node_id is the node that the IP should be released for. |
| ip      | [string](#string) |       | ip is the IP that should be released.                   |

### StoreLogRequest

StoreLogRequest is the message containing a raft log entry.

| Field | Type                             | Label | Description                          |
|-------|----------------------------------|-------|--------------------------------------|
| term  | [uint64](#uint64)                |       | term is the term of the log entry.   |
| index | [uint64](#uint64)                |       | index is the index of the log entry. |
| log   | [RaftLogEntry](#v1.RaftLogEntry) |       | log is the log entry.                |

### AllocateIPRequest.IPVersion

| Name               | Number | Description                                                      |
|--------------------|--------|------------------------------------------------------------------|
| IP_VERSION_UNKNOWN | 0      | IP_VERSION_UNKNOWN is the default value of IPVersion.            |
| IP_VERSION_4       | 4      | IP_VERSION_4 indicates that an IPv4 address should be allocated. |
| IP_VERSION_6       | 6      | IP_VERSION_6 indicates that an IPv6 address should be allocated. |

### PluginCapability

PluginCapability is the capabilities of a plugin.

| Name                      | Number | Description                                                                      |
|---------------------------|--------|----------------------------------------------------------------------------------|
| PLUGIN_CAPABILITY_UNKNOWN | 0      | PLUGIN_CAPABILITY_UNKNOWN is the default value of PluginCapability.              |
| PLUGIN_CAPABILITY_STORE   | 1      | PLUGIN_CAPABILITY_STORE indicates that the plugin is a raft store plugin.        |
| PLUGIN_CAPABILITY_AUTH    | 2      | PLUGIN_CAPABILITY_AUTH indicates that the plugin is an auth plugin.              |
| PLUGIN_CAPABILITY_WATCH   | 3      | PLUGIN_CAPABILITY_WATCH indicates that the plugin wants to receive watch events. |
| PLUGIN_CAPABILITY_IPAMV4  | 4      | PLUGIN_CAPABILITY_IPAMV4 indicates that the plugin is an IPv4 IPAM plugin.       |
| PLUGIN_CAPABILITY_IPAMV6  | 5      | PLUGIN_CAPABILITY_IPAMV6 indicates that the plugin is an IPv6 IPAM plugin.       |

### PluginQuery.QueryCommand

QueryCommand is the type of the query.

| Name | Number | Description                                                       |
|------|--------|-------------------------------------------------------------------|
| GET  | 0      | GET is the command to get a value.                                |
| LIST | 1      | LIST is the command to list keys with an optional prefix.         |
| ITER | 2      | ITER is the command to iterate over keys with an optional prefix. |

### WatchEvent

WatchEvent is the type of a watch event.

| Name                      | Number | Description                                                                     |
|---------------------------|--------|---------------------------------------------------------------------------------|
| WATCH_EVENT_UNKNOWN       | 0      | WATCH_EVENT_UNKNOWN is the default value of WatchEvent.                         |
| WATCH_EVENT_NODE_JOIN     | 1      | WATCH_EVENT_NODE_JOIN indicates that a node has joined the cluster.             |
| WATCH_EVENT_NODE_LEAVE    | 2      | WATCH_EVENT_NODE_LEAVE indicates that a node has left the cluster.              |
| WATCH_EVENT_LEADER_CHANGE | 3      | WATCH_EVENT_LEADER_CHANGE indicates that the leader of the cluster has changed. |

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

| Method Name   | Request Type                                      | Response Type                                    | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                         |
|---------------|---------------------------------------------------|--------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| GetInfo       | [.google.protobuf.Empty](#google.protobuf.Empty)  | [PluginInfo](#v1.PluginInfo)                     | GetInfo returns the information for the plugin.                                                                                                                                                                                                                                                                                                                                                                                                                     |
| Configure     | [PluginConfiguration](#v1.PluginConfiguration)    | [.google.protobuf.Empty](#google.protobuf.Empty) | Configure configures the plugin.                                                                                                                                                                                                                                                                                                                                                                                                                                    |
| InjectQuerier | [PluginQueryResult](#v1.PluginQueryResult) stream | [PluginQuery](#v1.PluginQuery) stream            | InjectQuerier is a stream opened by the node to faciliate read-only queries against the mesh state. The signature is misleading, but it is required to be able to stream the query results back to the node. The node will open a stream to the plugin and send a PluginSQLQueryResult message for every query that is received. The plugin can return an Unimplemented error or simply close the stream with no error it it does not wish to keep the stream open. |
| Close         | [.google.protobuf.Empty](#google.protobuf.Empty)  | [.google.protobuf.Empty](#google.protobuf.Empty) | Close closes the plugin. It is called when the node is shutting down.                                                                                                                                                                                                                                                                                                                                                                                               |

### StoragePlugin

StoragePlugin is the service definition for a Webmesh storage plugin.

| Method Name     | Request Type                           | Response Type                                    | Description                                                                   |
|-----------------|----------------------------------------|--------------------------------------------------|-------------------------------------------------------------------------------|
| Store           | [StoreLogRequest](#v1.StoreLogRequest) | [RaftApplyResponse](#v1.RaftApplyResponse)       | Store dispatches a Raft log entry for storage.                                |
| RestoreSnapshot | [DataSnapshot](#v1.DataSnapshot)       | [.google.protobuf.Empty](#google.protobuf.Empty) | RestoreSnapshot should drop any existing state and restore from the snapshot. |

### WatchPlugin

WatchPlugin is the service definition for a Webmesh watch plugin.

| Method Name | Request Type       | Response Type                                    | Description                 |
|-------------|--------------------|--------------------------------------------------|-----------------------------|
| Emit        | [Event](#v1.Event) | [.google.protobuf.Empty](#google.protobuf.Empty) | Emit handles a watch event. |

<div class="file-heading">

## v1/webrtc.proto

[Top](#title)

</div>

### DataChannelOffer

DataChannelOffer is an offer for a data channel. Candidates

are sent after the offer is sent.

| Field        | Type              | Label    | Description                                      |
|--------------|-------------------|----------|--------------------------------------------------|
| offer        | [string](#string) |          | offer is the offer.                              |
| stun_servers | [string](#string) | repeated | stun_servers is the list of STUN servers to use. |
| candidate    | [string](#string) |          | candidate is an ICE candidate.                   |

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

| Method Name      | Request Type                                                  | Response Type                                   | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          |
|------------------|---------------------------------------------------------------|-------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| StartDataChannel | [StartDataChannelRequest](#v1.StartDataChannelRequest) stream | [DataChannelOffer](#v1.DataChannelOffer) stream | StartDataChannel requests a new WebRTC connection to a node. The client speaks first with the request containing the node ID and where forwarded packets should be sent. The server responds with an offer and STUN servers to be used to establish a WebRTC connection. The client should then respond with an answer to the offer that matches the spec of the DataChannel.CHANNELS enum value. After the offer is accepted, the stream can be used to exchange ICE candidates to speed up the connection process. |

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
