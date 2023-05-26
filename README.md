# Protocol Documentation

## Table of Contents

<div id="toc-container">

- [v1/admin_messages.proto](#v1%2fadmin_messages.proto)
  - [<span class="badge">M</span>DeleteRaftACLRequest](#.DeleteRaftACLRequest)
  - [<span class="badge">M</span>GetRaftACLRequest](#.GetRaftACLRequest)
  - [<span class="badge">M</span>RaftACL](#.RaftACL)
  - [<span class="badge">M</span>RaftACLList](#.RaftACLList)
  - [<span class="badge">E</span>ACLAction](#.ACLAction)
- [v1/admin.proto](#v1%2fadmin.proto)
  - [<span class="badge">S</span>Admin](#v1.Admin)
- [v1/mesh_messages.proto](#v1%2fmesh_messages.proto)
  - [<span class="badge">M</span>GetNodeRequest](#v1.GetNodeRequest)
  - [<span class="badge">M</span>MeshNode](#v1.MeshNode)
  - [<span class="badge">M</span>NodeList](#v1.NodeList)
  - [<span class="badge">E</span>ClusterStatus](#v1.ClusterStatus)
  - [<span class="badge">E</span>Feature](#v1.Feature)
- [v1/mesh.proto](#v1%2fmesh.proto)
  - [<span class="badge">S</span>Mesh](#v1.Mesh)
- [v1/node_messages.proto](#v1%2fnode_messages.proto)
  - [<span class="badge">M</span>DataChannelNegotiation](#v1.DataChannelNegotiation)
  - [<span class="badge">M</span>GetStatusRequest](#v1.GetStatusRequest)
  - [<span class="badge">M</span>JoinRequest](#v1.JoinRequest)
  - [<span class="badge">M</span>JoinResponse](#v1.JoinResponse)
  - [<span class="badge">M</span>LeaveRequest](#v1.LeaveRequest)
  - [<span class="badge">M</span>Status](#v1.Status)
  - [<span class="badge">M</span>WireguardPeer](#v1.WireguardPeer)
  - [<span class="badge">E</span>DataChannel](#v1.DataChannel)
- [v1/node.proto](#v1%2fnode.proto)
  - [<span class="badge">S</span>Node](#v1.Node)
- [v1/node_metrics.proto](#v1%2fnode_metrics.proto)
  - [<span class="badge">M</span>CPUTimes](#v1.CPUTimes)
  - [<span class="badge">M</span>DiskInfo](#v1.DiskInfo)
  - [<span class="badge">M</span>HostInfo](#v1.HostInfo)
  - [<span class="badge">M</span>HostMetrics](#v1.HostMetrics)
  - [<span class="badge">M</span>LoadAverage](#v1.LoadAverage)
  - [<span class="badge">M</span>MemoryInfo](#v1.MemoryInfo)
  - [<span class="badge">M</span>NodeMetrics](#v1.NodeMetrics)
  - [<span class="badge">M</span>OSInfo](#v1.OSInfo)
  - [<span class="badge">M</span>PeerMetrics](#v1.PeerMetrics)
- [v1/node_raft.proto](#v1%2fnode_raft.proto)
  - [<span class="badge">M</span>RaftApplyResponse](#v1.RaftApplyResponse)
  - [<span class="badge">M</span>RaftLogEntry](#v1.RaftLogEntry)
  - [<span class="badge">M</span>SQLExec](#v1.SQLExec)
  - [<span class="badge">M</span>SQLExecResult](#v1.SQLExecResult)
  - [<span class="badge">M</span>SQLParameter](#v1.SQLParameter)
  - [<span class="badge">M</span>SQLQuery](#v1.SQLQuery)
  - [<span class="badge">M</span>SQLQueryResult](#v1.SQLQueryResult)
  - [<span class="badge">M</span>SQLStatement](#v1.SQLStatement)
  - [<span class="badge">M</span>SQLValues](#v1.SQLValues)
  - [<span class="badge">E</span>RaftCommandType](#v1.RaftCommandType)
  - [<span class="badge">E</span>SQLParameterType](#v1.SQLParameterType)
- [v1/peer_discovery.proto](#v1%2fpeer_discovery.proto)
  - [<span class="badge">M</span>ListRaftPeersResponse](#v1.ListRaftPeersResponse)
  - [<span class="badge">M</span>RaftPeer](#v1.RaftPeer)
  - [<span class="badge">S</span>PeerDiscovery](#v1.PeerDiscovery)
- [v1/webrtc_messages.proto](#v1%2fwebrtc_messages.proto)
  - [<span class="badge">M</span>DataChannelOffer](#v1.DataChannelOffer)
  - [<span class="badge">M</span>StartDataChannelRequest](#v1.StartDataChannelRequest)
- [v1/webrtc.proto](#v1%2fwebrtc.proto)
  - [<span class="badge">S</span>WebRTC](#v1.WebRTC)
- [Scalar Value Types](#scalar-value-types)

</div>

<div class="file-heading">

## v1/admin_messages.proto

[Top](#title)

</div>

### DeleteRaftACLRequest

DeleteRaftACLRequest is a request to delete a RaftACL.

| Field | Type              | Label | Description                  |
|-------|-------------------|-------|------------------------------|
| name  | [string](#string) |       | name is the name of the ACL. |

### GetRaftACLRequest

GetRaftACLRequest is a request to get a RaftACL.

| Field | Type              | Label | Description                  |
|-------|-------------------|-------|------------------------------|
| name  | [string](#string) |       | name is the name of the ACL. |

### RaftACL

RaftACL is an ACL that applies to Raft nodes.

| Field    | Type                    | Label    | Description                                                                                                           |
|----------|-------------------------|----------|-----------------------------------------------------------------------------------------------------------------------|
| name     | [string](#string)       |          | name is the name of the ACL.                                                                                          |
| patterns | [string](#string)       | repeated | patterns is a list of node patterns this ACL applies to.                                                              |
| action   | [ACLAction](#ACLAction) |          | action is the action of this ACL. For RaftACLs This is if the matching nodes are allowed to vote in elections or not. |

### RaftACLList

RaftACLList is a list of RaftACLs.

| Field | Type                | Label    | Description                  |
|-------|---------------------|----------|------------------------------|
| items | [RaftACL](#RaftACL) | repeated | items is a list of RaftACLs. |

### ACLAction

ACLAction is the action of an ACL.

| Name           | Number | Description                       |
|----------------|--------|-----------------------------------|
| ACLActionDeny  | 0      | ACLActionDeny denies the action.  |
| ACLActionAllow | 1      | ACLActionAllow allows the action. |

<div class="file-heading">

## v1/admin.proto

[Top](#title)

</div>

### Admin

Admin is the service that provides cluster admin operations. Most
methods

require the leader to be contacted.

| Method Name   | Request Type                                     | Response Type                                    | Description                                         |
|---------------|--------------------------------------------------|--------------------------------------------------|-----------------------------------------------------|
| PutRaftACL    | [.RaftACL](#RaftACL)                             | [.google.protobuf.Empty](#google.protobuf.Empty) | PutRaftACL creates or updates an ACL.               |
| DeleteRaftACL | [.DeleteRaftACLRequest](#DeleteRaftACLRequest)   | [.google.protobuf.Empty](#google.protobuf.Empty) | DeleteRaftACL deletes an ACL by name.               |
| GetRaftACL    | [.GetRaftACLRequest](#GetRaftACLRequest)         | [.RaftACL](#RaftACL)                             | GetRaftACL returns an ACL by name.                  |
| ListRaftACLs  | [.google.protobuf.Empty](#google.protobuf.Empty) | [.RaftACLList](#RaftACLList)                     | ListRaftACLs returns the ACLs for the Raft cluster. |

<div class="file-heading">

## v1/mesh_messages.proto

[Top](#title)

</div>

### GetNodeRequest

GetNodeRequest is a request to get a node.

| Field | Type              | Label | Description               |
|-------|-------------------|-------|---------------------------|
| id    | [string](#string) |       | id is the ID of the node. |

### MeshNode

MeshNode is a node that has been registered with the controller.

| Field            | Type                                                    | Label    | Description                                              |
|------------------|---------------------------------------------------------|----------|----------------------------------------------------------|
| id               | [string](#string)                                       |          | id is the ID of the node.                                |
| primary_endpoint | [string](#string)                                       |          | primary_endpoint is the primary_endpoint of the node.    |
| endpoints        | [string](#string)                                       | repeated | endpoints are alternate endpoints for the node.          |
| raft_port        | [int32](#int32)                                         |          | raft_port is the Raft listen port of the node.           |
| grpc_port        | [int32](#int32)                                         |          | grpc_port is the gRPC listen port of the node.           |
| wireguard_port   | [int32](#int32)                                         |          | wireguard_port is the Wireguard listen port of the node. |
| public_key       | [string](#string)                                       |          | public_key is the public key of the node.                |
| private_ipv4     | [string](#string)                                       |          | private_ipv4 is the private IPv4 address of the node.    |
| private_ipv6     | [string](#string)                                       |          | private_ipv6 is the private IPv6 address of the node.    |
| updated_at       | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |          | updated_at is the last time the node joined the cluster. |
| created_at       | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |          | created_at is the creation time for the node.            |
| cluster_status   | [ClusterStatus](#v1.ClusterStatus)                      |          | cluster_status is the status of the node in the cluster. |

### NodeList

NodeList is a list of nodes.

| Field | Type                     | Label    | Description                 |
|-------|--------------------------|----------|-----------------------------|
| nodes | [MeshNode](#v1.MeshNode) | repeated | nodes is the list of nodes. |

### ClusterStatus

ClusterStatus is the status of the node in the cluster.

| Name                   | Number | Description                                           |
|------------------------|--------|-------------------------------------------------------|
| CLUSTER_STATUS_UNKNOWN | 0      | CLUSTER_STATUS_UNKNOWN is the default status.         |
| CLUSTER_LEADER         | 1      | CLUSTER_LEADER is the status for the leader node.     |
| CLUSTER_VOTER          | 2      | CLUSTER_VOTER is the status for a voter node.         |
| CLUSTER_NON_VOTER      | 3      | CLUSTER_NON_VOTER is the status for a non-voter node. |

### Feature

Feature is a list of features supported by a node.

| Name            | Number | Description                                               |
|-----------------|--------|-----------------------------------------------------------|
| FEATURE_NONE    | 0      | FEATURE_NONE is the default feature set.                  |
| NODES           | 1      | NODES is the feature for nodes. This is always supported. |
| LEADER_PROXY    | 2      | LEADER_PROXY is the feature for leader proxying.          |
| MESH_API        | 3      | MESH_API is the feature for the mesh API.                 |
| PEER_DISCOVERY  | 4      | PEER_DISCOVERY is the feature for peer discovery.         |
| METRICS         | 5      | METRICS is the feature for exposing metrics.              |
| ICE_NEGOTIATION | 6      | ICE_NEGOTIATION is the feature for ICE negotiation.       |
| TURN_SERVER     | 7      | TURN_SERVER is the feature for TURN server.               |
| MESH_DNS        | 8      | MESH_DNS is the feature for mesh DNS.                     |

<div class="file-heading">

## v1/mesh.proto

[Top](#title)

</div>

### Mesh

Mesh is a service that can optionally be exposed by a node. It provides
methods for

interfacing with the webmesh from the outside. Some methods are only
available on the

leader. Nodes can enable the leader proxy to expose the leader's Mesh
service.

| Method Name | Request Type                                     | Response Type            | Description                |
|-------------|--------------------------------------------------|--------------------------|----------------------------|
| GetNode     | [GetNodeRequest](#v1.GetNodeRequest)             | [MeshNode](#v1.MeshNode) | GetNode gets a node by ID. |
| ListNodes   | [.google.protobuf.Empty](#google.protobuf.Empty) | [NodeList](#v1.NodeList) | ListNodes lists all nodes. |

<div class="file-heading">

## v1/node_messages.proto

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

### JoinRequest

JoinRequest is a request to join the cluster.

| Field            | Type              | Label    | Description                                                                                                                             |
|------------------|-------------------|----------|-----------------------------------------------------------------------------------------------------------------------------------------|
| id               | [string](#string) |          | id is the ID of the node.                                                                                                               |
| public_key       | [string](#string) |          | public_key is the public wireguard key of the node to broadcast to peers.                                                               |
| raft_port        | [int32](#int32)   |          | raft_port is the Raft listen port of the node.                                                                                          |
| grpc_port        | [int32](#int32)   |          | grpc_port is the gRPC listen port of the node.                                                                                          |
| wireguard_port   | [int32](#int32)   |          | wireguard_port is the Wireguard listen port of the node.                                                                                |
| primary_endpoint | [string](#string) |          | primary_endpoint is a routable address for the node. If left unset, the node is assumed to be behind a NAT and not directly accessible. |
| endpoints        | [string](#string) | repeated | endpoints are additional endpoints routable to the node. See primary_endpoint.                                                          |
| assign_ipv4      | [bool](#bool)     |          | assign_ipv4 is whether an IPv4 address should be assigned to the node.                                                                  |
| prefer_raft_ipv6 | [bool](#bool)     |          | prefer_raft_ipv6 is whether IPv6 should be preferred over IPv4 for raft communication. This is only used if assign_ipv4 is true.        |
| as_voter         | [bool](#bool)     |          | as_voter is whether the node should receive a vote in elections.                                                                        |

### JoinResponse

JoinResponse is a response to a join request.

| Field        | Type                               | Label    | Description                                                                                                                                                                                     |
|--------------|------------------------------------|----------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| address_ipv4 | [string](#string)                  |          | address_ipv4 is the private IPv4 wireguard address of the node in CIDR format representing the network. This is only set if assign_ipv4 was set in the request or no network_ipv6 was provided. |
| network_ipv6 | [string](#string)                  |          | network_ipv6 is the IPv6 network assigned to the node.                                                                                                                                          |
| peers        | [WireguardPeer](#v1.WireguardPeer) | repeated | peers is a list of wireguard peers to connect to.                                                                                                                                               |

### LeaveRequest

LeaveRequest is a request to leave the cluster.

| Field | Type              | Label | Description               |
|-------|-------------------|-------|---------------------------|
| id    | [string](#string) |       | id is the ID of the node. |

### Status

Status represents the status of a node.

| Field           | Type                                                    | Label    | Description                                                 |
|-----------------|---------------------------------------------------------|----------|-------------------------------------------------------------|
| id              | [string](#string)                                       |          | id is the ID of the node.                                   |
| version         | [string](#string)                                       |          | version is the version of the node.                         |
| commit          | [string](#string)                                       |          | commit is the commit of the node.                           |
| build_date      | [string](#string)                                       |          | build_date is the build date of the node.                   |
| uptime          | [string](#string)                                       |          | uptime is the uptime of the node.                           |
| started_at      | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |          | started_at is the time the node started.                    |
| features        | [Feature](#v1.Feature)                                  | repeated | features is the list of features currently enabled.         |
| wireguard_peers | [uint32](#uint32)                                       |          | wireguard_peers is the number of wireguard peers connected. |
| cluster_status  | [ClusterStatus](#v1.ClusterStatus)                      |          | cluster_status is the status of the node in the cluster.    |
| current_leader  | [string](#string)                                       |          | current_leader is the current leader of the cluster.        |

### WireguardPeer

WireguardPeer is a peer in the Wireguard network.

| Field            | Type              | Label    | Description                                                          |
|------------------|-------------------|----------|----------------------------------------------------------------------|
| id               | [string](#string) |          | id is the ID of the peer.                                            |
| public_key       | [string](#string) |          | public_key is the public key of the peer.                            |
| primary_endpoint | [string](#string) |          | primary_endpoint is the primary endpoint of the peer, if applicable. |
| endpoints        | [string](#string) | repeated | endpoints are alternative endpoints of the peer, if applicable.      |
| address_ipv4     | [string](#string) |          | address_ipv4 is the private IPv4 wireguard address of the peer.      |
| address_ipv6     | [string](#string) |          | address_ipv6 is the private IPv6 wireguard address of the peer.      |

### DataChannel

DataChannel are the data channels used when communicating over ICE

with a node.

| Name        | Number | Description                                                                                                                                                                                                             |
|-------------|--------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| CHANNELS    | 0      | CHANNELS is the data channel used for negotiating new channels. This is the first channel that is opened. The ID of the channel should be 0.                                                                            |
| CONNECTIONS | 1      | CONNECTIONS is the data channel used for negotiating new connections. This is a channel that is opened for each incoming connection from a client. The ID should start at 0 and be incremented for each new connection. |

<div class="file-heading">

## v1/node.proto

[Top](#title)

</div>

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
| Leave                | [LeaveRequest](#v1.LeaveRequest)                            | [.google.protobuf.Empty](#google.protobuf.Empty)            | Leave is used to remove a node from the mesh. The node will be removed from the mesh and will no longer be able to query the mesh state or vote in elections.                                                                                                                                                                                                                                                                                                                                                       |
| GetStatus            | [GetStatusRequest](#v1.GetStatusRequest)                    | [Status](#v1.Status)                                        | GetStatus gets the status of a node in the cluster.                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
| NegotiateDataChannel | [DataChannelNegotiation](#v1.DataChannelNegotiation) stream | [DataChannelNegotiation](#v1.DataChannelNegotiation) stream | NegotiateDataChannel is used to negotiate a WebRTC connection between a webmesh client and a node in the cluster. The handling server will send the target node the source address, the destination for traffic, and STUN/TURN servers to use for the negotiation. The node responds with an offer to be forwarded to the client. When the handler receives an answer from the client, it forwards it to the node. Once the node receives the answer, the stream can optionally be used to exchange ICE candidates. |

<div class="file-heading">

## v1/node_metrics.proto

[Top](#title)

</div>

### CPUTimes

CPUTimes is the CPU times.

| Field        | Type                           | Label | Description                        |
|--------------|--------------------------------|-------|------------------------------------|
| load_average | [LoadAverage](#v1.LoadAverage) |       |                                    |
| user         | [string](#string)              |       | user is the user CPU time.         |
| system       | [string](#string)              |       | system is the system CPU time.     |
| idle         | [string](#string)              |       | idle is the idle CPU time.         |
| io_wait      | [string](#string)              |       | io_wait is the IO wait CPU time.   |
| irq          | [string](#string)              |       | irq is the IRQ CPU time.           |
| nice         | [string](#string)              |       | nice is the nice CPU time.         |
| soft_irq     | [string](#string)              |       | soft_irq is the soft IRQ CPU time. |
| steal        | [string](#string)              |       | steal is the steal CPU time.       |

### DiskInfo

DiskInfo is the disk information.

| Field               | Type              | Label | Description                                         |
|---------------------|-------------------|-------|-----------------------------------------------------|
| filesystem_path     | [string](#string) |       | filesystem_path is the path the disk is mounted to. |
| device              | [string](#string) |       | device is the device of the disk.                   |
| total               | [uint64](#uint64) |       | total is the total disk space.                      |
| free                | [uint64](#uint64) |       | free is the free disk space.                        |
| used                | [uint64](#uint64) |       | used is the used disk space.                        |
| used_percent        | [float](#float)   |       | used_percent is the used disk space percentage.     |
| inodes_total        | [uint64](#uint64) |       | inodes_total is the total inodes.                   |
| inodes_free         | [uint64](#uint64) |       | inodes_free is the free inodes.                     |
| inodes_used         | [uint64](#uint64) |       | inodes_used is the used inodes.                     |
| inodes_used_percent | [float](#float)   |       | inodes_used_percent is the used inodes percentage.  |

### HostInfo

HostInfo is the host information.

| Field          | Type                 | Label    | Description                                         |
|----------------|----------------------|----------|-----------------------------------------------------|
| architecture   | [string](#string)    |          | architecture is the architecture of the host.       |
| boot_time      | [string](#string)    |          | boot_time is the boot time of the host.             |
| containerized  | [bool](#bool)        |          | containerized is whether the host is containerized. |
| hostname       | [string](#string)    |          | hostname is the hostname of the host.               |
| ips            | [string](#string)    | repeated | ips is the list of IP addresses of the host.        |
| kernel_version | [string](#string)    |          | kernel_version is the kernel version of the host.   |
| macs           | [string](#string)    | repeated | macs is the list of MAC addresses of the host.      |
| os             | [OSInfo](#v1.OSInfo) |          | os is the OS of the host.                           |
| timezone       | [string](#string)    |          | timezone is the timezone of the host.               |
| uptime         | [string](#string)    |          | uptime is the uptime of the host.                   |

### HostMetrics

HostMetrics is the metrics for the host.

| Field  | Type                         | Label    | Description                   |
|--------|------------------------------|----------|-------------------------------|
| host   | [HostInfo](#v1.HostInfo)     |          | host is the host metrics.     |
| cpu    | [CPUTimes](#v1.CPUTimes)     |          | cpu is the CPU metrics.       |
| memory | [MemoryInfo](#v1.MemoryInfo) |          | memory is the memory metrics. |
| disks  | [DiskInfo](#v1.DiskInfo)     | repeated | disks is the disk metrics.    |

### LoadAverage

LoadAverage is the load average.

| Field   | Type            | Label | Description                            |
|---------|-----------------|-------|----------------------------------------|
| one     | [float](#float) |       | one is the 1 minute load average.      |
| five    | [float](#float) |       | five is the 5 minute load average.     |
| fifteen | [float](#float) |       | fifteen is the 15 minute load average. |

### MemoryInfo

MemoryInfo is the memory information.

| Field         | Type              | Label | Description                                |
|---------------|-------------------|-------|--------------------------------------------|
| total         | [uint64](#uint64) |       | total is the total memory.                 |
| used          | [uint64](#uint64) |       | used is the used memory.                   |
| available     | [uint64](#uint64) |       | available is the available memory.         |
| free          | [uint64](#uint64) |       | free is the free memory.                   |
| virtual_total | [uint64](#uint64) |       | virtual_total is the total virtual memory. |
| virtual_used  | [uint64](#uint64) |       | virtual_used is the used virtual memory.   |
| virtual_free  | [uint64](#uint64) |       | virtual_free is the free virtual memory.   |

### NodeMetrics

NodeMetrics is the metrics for the node.

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
| peers                | [PeerMetrics](#v1.PeerMetrics) | repeated | peers are the per-peer statistics.                             |
| system               | [HostMetrics](#v1.HostMetrics) |          | system is the system metrics for the node.                     |

### OSInfo

OSInfo is the OS information.

| Field    | Type              | Label | Description                       |
|----------|-------------------|-------|-----------------------------------|
| type     | [string](#string) |       | type is the type of OS.           |
| family   | [string](#string) |       | family is the family of OS.       |
| platform | [string](#string) |       | platform is the platform of OS.   |
| name     | [string](#string) |       | name is the name of OS.           |
| version  | [string](#string) |       | version is the version of OS.     |
| major    | [int64](#int64)   |       | major is the major version of OS. |
| minor    | [int64](#int64)   |       | minor is the minor version of OS. |
| patch    | [int64](#int64)   |       | patch is the patch version of OS. |
| build    | [string](#string) |       | build is the build version of OS. |
| codename | [string](#string) |       | codename is the codename of OS.   |

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

<div class="file-heading">

## v1/node_raft.proto

[Top](#title)

</div>

### RaftApplyResponse

RaftApplyResponse is the response to an apply request. It

contains the result of applying the log entry.

| Field        | Type                                 | Label | Description                                            |
|--------------|--------------------------------------|-------|--------------------------------------------------------|
| time         | [string](#string)                    |       | time is the total time it took to apply the log entry. |
| query_result | [SQLQueryResult](#v1.SQLQueryResult) |       | query is the query result of the log entry.            |
| exec_result  | [SQLExecResult](#v1.SQLExecResult)   |       | exec is the exec result of the log entry.              |
| error        | [string](#string)                    |       | error is an error that occurred during the apply.      |

### RaftLogEntry

RaftLogEntry is the data of an entry in the Raft log.

| Field     | Type                                   | Label | Description                                                                     |
|-----------|----------------------------------------|-------|---------------------------------------------------------------------------------|
| type      | [RaftCommandType](#v1.RaftCommandType) |       | type is the type of the log entry.                                              |
| sql_query | [SQLQuery](#v1.SQLQuery)               |       | data is the data of the log entry. sql_query is the SQL query of the log entry. |
| sql_exec  | [SQLExec](#v1.SQLExec)                 |       | sql_exec is the SQL exec of the log entry.                                      |

### SQLExec

SQLExec is a SQL exec.

| Field       | Type                             | Label | Description                                          |
|-------------|----------------------------------|-------|------------------------------------------------------|
| transaction | [bool](#bool)                    |       | transaction flags whether the exec is a transaction. |
| statement   | [SQLStatement](#v1.SQLStatement) |       | statement is the statement of the SQL exec.          |

### SQLExecResult

SQLExecResult is the result of a SQL exec.

| Field          | Type              | Label | Description                                      |
|----------------|-------------------|-------|--------------------------------------------------|
| last_insert_id | [int64](#int64)   |       | last_insert_id is the last insert ID.            |
| rows_affected  | [int64](#int64)   |       | rows_affected is the number of rows affected.    |
| error          | [string](#string) |       | error is an error that occurred during the exec. |
| time           | [string](#string) |       | time is the time it took to execute the exec.    |

### SQLParameter

SQLParameter is a parameter of a SQL query.

| Field  | Type                                                    | Label | Description                        |
|--------|---------------------------------------------------------|-------|------------------------------------|
| name   | [string](#string)                                       |       | name is the name of the parameter. |
| type   | [SQLParameterType](#v1.SQLParameterType)                |       |                                    |
| int64  | [sint64](#sint64)                                       |       |                                    |
| double | [double](#double)                                       |       |                                    |
| bool   | [bool](#bool)                                           |       |                                    |
| bytes  | [bytes](#bytes)                                         |       |                                    |
| str    | [string](#string)                                       |       |                                    |
| time   | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |       |                                    |

### SQLQuery

SQLQuery is a SQL query.

| Field       | Type                             | Label | Description                                           |
|-------------|----------------------------------|-------|-------------------------------------------------------|
| transaction | [bool](#bool)                    |       | transaction flags whether the query is a transaction. |
| statement   | [SQLStatement](#v1.SQLStatement) |       | statement is the statement of the SQL query.          |

### SQLQueryResult

SQLQueryResult contains the rows of a SQL query.

| Field   | Type                       | Label    | Description                                       |
|---------|----------------------------|----------|---------------------------------------------------|
| columns | [string](#string)          | repeated | columns is the list of columns.                   |
| types   | [string](#string)          | repeated | types is the list of types.                       |
| values  | [SQLValues](#v1.SQLValues) | repeated | values is the list of values.                     |
| error   | [string](#string)          |          | error is an error that occurred during the query. |
| time    | [string](#string)          |          | time is the time it took to execute the query.    |

### SQLStatement

SQLStatement is a SQL statement.

| Field      | Type                             | Label    | Description                                        |
|------------|----------------------------------|----------|----------------------------------------------------|
| sql        | [string](#string)                |          | sql is the SQL statement.                          |
| parameters | [SQLParameter](#v1.SQLParameter) | repeated | parameters is the parameters of the SQL statement. |

### SQLValues

SQLValues is a list of values.

| Field  | Type                             | Label    | Description                   |
|--------|----------------------------------|----------|-------------------------------|
| values | [SQLParameter](#v1.SQLParameter) | repeated | values is the list of values. |

### RaftCommandType

RaftCommandType is the type of command being sent to the

Raft log.

| Name    | Number | Description                          |
|---------|--------|--------------------------------------|
| UNKNOWN | 0      | UNKNOWN is the unknown command type. |
| QUERY   | 1      | QUERY is the query command type.     |
| EXECUTE | 2      | EXECUTE is the execute command type. |

### SQLParameterType

| Name              | Number | Description                            |
|-------------------|--------|----------------------------------------|
| SQL_PARAM_UNKNOWN | 0      | UNKNOWN is the unknown parameter type. |
| SQL_PARAM_INT64   | 1      | INT64 is the int64 parameter type.     |
| SQL_PARAM_DOUBLE  | 2      | DOUBLE is the double parameter type.   |
| SQL_PARAM_BOOL    | 3      | BOOL is the bool parameter type.       |
| SQL_PARAM_BYTES   | 4      | BYTES is the bytes parameter type.     |
| SQL_PARAM_STRING  | 5      | STRING is the string parameter type.   |
| SQL_PARAM_TIME    | 6      | TIME is the time parameter type.       |
| SQL_PARAM_NULL    | 7      | NULL is the null parameter type.       |

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

## v1/webrtc_messages.proto

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

| Field     | Type              | Label | Description                                        |
|-----------|-------------------|-------|----------------------------------------------------|
| node_id   | [string](#string) |       | node_id is the ID of the node to send the data to. |
| proto     | [string](#string) |       | proto is the protocol of the traffic.              |
| dst       | [string](#string) |       | dst is the destination address of the traffic.     |
| port      | [uint32](#uint32) |       | port is the destination port of the traffic.       |
| answer    | [string](#string) |       | answer is the answer to the offer.                 |
| candidate | [string](#string) |       | candidate is an ICE candidate.                     |

<div class="file-heading">

## v1/webrtc.proto

[Top](#title)

</div>

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
