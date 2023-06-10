# Protocol Documentation

## Table of Contents

<div id="toc-container">

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
  - [<span class="badge">E</span>RuleVerbs](#v1.RuleVerbs)
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
- [v1/node.proto](#v1%2fnode.proto)
  - [<span class="badge">M</span>DataChannelNegotiation](#v1.DataChannelNegotiation)
  - [<span class="badge">M</span>GetStatusRequest](#v1.GetStatusRequest)
  - [<span class="badge">M</span>InterfaceMetrics](#v1.InterfaceMetrics)
  - [<span class="badge">M</span>JoinRequest](#v1.JoinRequest)
  - [<span class="badge">M</span>JoinResponse](#v1.JoinResponse)
  - [<span class="badge">M</span>LeaveRequest](#v1.LeaveRequest)
  - [<span class="badge">M</span>PeerMetrics](#v1.PeerMetrics)
  - [<span class="badge">M</span>Status](#v1.Status)
  - [<span class="badge">M</span>WireGuardPeer](#v1.WireGuardPeer)
  - [<span class="badge">E</span>ClusterStatus](#v1.ClusterStatus)
  - [<span class="badge">E</span>DataChannel](#v1.DataChannel)
  - [<span class="badge">E</span>Feature](#v1.Feature)
  - [<span class="badge">S</span>Node](#v1.Node)
- [v1/mesh.proto](#v1%2fmesh.proto)
  - [<span class="badge">M</span>GetNodeRequest](#v1.GetNodeRequest)
  - [<span class="badge">M</span>MeshEdge](#v1.MeshEdge)
  - [<span class="badge">M</span>MeshGraph](#v1.MeshGraph)
  - [<span class="badge">M</span>MeshNode](#v1.MeshNode)
  - [<span class="badge">M</span>NodeList](#v1.NodeList)
  - [<span class="badge">S</span>Mesh](#v1.Mesh)
- [v1/peer_discovery.proto](#v1%2fpeer_discovery.proto)
  - [<span class="badge">M</span>ListRaftPeersResponse](#v1.ListRaftPeersResponse)
  - [<span class="badge">M</span>RaftPeer](#v1.RaftPeer)
  - [<span class="badge">S</span>PeerDiscovery](#v1.PeerDiscovery)
- [v1/raft.proto](#v1%2fraft.proto)
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
- [v1/webrtc.proto](#v1%2fwebrtc.proto)
  - [<span class="badge">M</span>DataChannelOffer](#v1.DataChannelOffer)
  - [<span class="badge">M</span>StartDataChannelRequest](#v1.StartDataChannelRequest)
  - [<span class="badge">S</span>WebRTC](#v1.WebRTC)
- [Scalar Value Types](#scalar-value-types)

</div>

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
| verb          | [RuleVerbs](#v1.RuleVerbs)       |       | verb is the verb that is performed on the resource.                         |

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
| verbs          | [RuleVerbs](#v1.RuleVerbs)       | repeated | verbs is the list of verbs that apply to the resource.                  |

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
| RESOURCE_ALL           | 999    | RESOURCE_ALL is a wildcard resource that matches all resources.                                                |

### RuleVerbs

RuleVerbs is the verb type for a rule.

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
| nodes             | [string](#string) |          | node is the node that broadcasts the route. A group can be specified with the prefix "group:". |
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

| Method Name       | Request Type                                     | Response Type                                    | Description                                       |
|-------------------|--------------------------------------------------|--------------------------------------------------|---------------------------------------------------|
| PutRole           | [Role](#v1.Role)                                 | [.google.protobuf.Empty](#google.protobuf.Empty) | PutRole creates or updates a role.                |
| DeleteRole        | [Role](#v1.Role)                                 | [.google.protobuf.Empty](#google.protobuf.Empty) | DeleteRole deletes a role.                        |
| GetRole           | [Role](#v1.Role)                                 | [Role](#v1.Role)                                 | GetRole gets a role.                              |
| ListRoles         | [.google.protobuf.Empty](#google.protobuf.Empty) | [Roles](#v1.Roles)                               | ListRoles gets all roles.                         |
| PutRoleBinding    | [RoleBinding](#v1.RoleBinding)                   | [.google.protobuf.Empty](#google.protobuf.Empty) | PutRoleBinding creates or updates a role binding. |
| DeleteRoleBinding | [RoleBinding](#v1.RoleBinding)                   | [.google.protobuf.Empty](#google.protobuf.Empty) | DeleteRoleBinding deletes a role binding.         |
| GetRoleBinding    | [RoleBinding](#v1.RoleBinding)                   | [RoleBinding](#v1.RoleBinding)                   | GetRoleBinding gets a role binding.               |
| ListRoleBindings  | [.google.protobuf.Empty](#google.protobuf.Empty) | [RoleBindings](#v1.RoleBindings)                 | ListRoleBindings gets all role bindings.          |
| PutGroup          | [Group](#v1.Group)                               | [.google.protobuf.Empty](#google.protobuf.Empty) | PutGroup creates or updates a group.              |
| DeleteGroup       | [Group](#v1.Group)                               | [.google.protobuf.Empty](#google.protobuf.Empty) | DeleteGroup deletes a group.                      |
| GetGroup          | [Group](#v1.Group)                               | [Group](#v1.Group)                               | GetGroup gets a group.                            |
| ListGroups        | [.google.protobuf.Empty](#google.protobuf.Empty) | [Groups](#v1.Groups)                             | ListGroups gets all groups.                       |
| PutNetworkACL     | [NetworkACL](#v1.NetworkACL)                     | [.google.protobuf.Empty](#google.protobuf.Empty) | PutNetworkACL creates or updates a network ACL.   |
| DeleteNetworkACL  | [NetworkACL](#v1.NetworkACL)                     | [.google.protobuf.Empty](#google.protobuf.Empty) | DeleteNetworkACL deletes a network ACL.           |
| GetNetworkACL     | [NetworkACL](#v1.NetworkACL)                     | [NetworkACL](#v1.NetworkACL)                     | GetNetworkACL gets a network ACL.                 |
| ListNetworkACLs   | [.google.protobuf.Empty](#google.protobuf.Empty) | [NetworkACLs](#v1.NetworkACLs)                   | ListNetworkACLs gets all network ACLs.            |
| PutRoute          | [Route](#v1.Route)                               | [.google.protobuf.Empty](#google.protobuf.Empty) | PutRoute creates or updates a route.              |
| DeleteRoute       | [Route](#v1.Route)                               | [.google.protobuf.Empty](#google.protobuf.Empty) | DeleteRoute deletes a route.                      |
| GetRoute          | [Route](#v1.Route)                               | [Route](#v1.Route)                               | GetRoute gets a route.                            |
| ListRoutes        | [.google.protobuf.Empty](#google.protobuf.Empty) | [Routes](#v1.Routes)                             | ListRoutes gets all routes.                       |

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

| Field               | Type              | Label    | Description                                                                                                                             |
|---------------------|-------------------|----------|-----------------------------------------------------------------------------------------------------------------------------------------|
| id                  | [string](#string) |          | id is the ID of the node.                                                                                                               |
| public_key          | [string](#string) |          | public_key is the public wireguard key of the node to broadcast to peers.                                                               |
| raft_port           | [int32](#int32)   |          | raft_port is the Raft listen port of the node.                                                                                          |
| grpc_port           | [int32](#int32)   |          | grpc_port is the gRPC listen port of the node.                                                                                          |
| primary_endpoint    | [string](#string) |          | primary_endpoint is a routable address for the node. If left unset, the node is assumed to be behind a NAT and not directly accessible. |
| wireguard_endpoints | [string](#string) | repeated | wireguard_endpoints is a list of WireGuard endpoints for the node.                                                                      |
| zone_awareness_id   | [string](#string) |          | zone_awareness_id is the zone awareness ID of the node.                                                                                 |
| assign_ipv4         | [bool](#bool)     |          | assign_ipv4 is whether an IPv4 address should be assigned to the node.                                                                  |
| prefer_raft_ipv6    | [bool](#bool)     |          | prefer_raft_ipv6 is whether IPv6 should be preferred over IPv4 for raft communication. This is only used if assign_ipv4 is true.        |
| as_voter            | [bool](#bool)     |          | as_voter is whether the node should receive a vote in elections.                                                                        |
| routes              | [string](#string) | repeated | routes is a list of routes to advertise to peers. The request will be denied if the node is not allowed to put routes.                  |

### JoinResponse

JoinResponse is a response to a join request.

| Field        | Type                               | Label    | Description                                                                                                                                                                                                                                                    |
|--------------|------------------------------------|----------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| address_ipv4 | [string](#string)                  |          | address_ipv4 is the private IPv4 wireguard address of the node in CIDR format representing the network. This is only set if assign_ipv4 was set in the request or no network_ipv6 was provided. The bits are set to the network bits of the Mesh IPv4 network. |
| address_ipv6 | [string](#string)                  |          | address_ipv6 is the IPv6 network assigned to the node.                                                                                                                                                                                                         |
| network_ipv6 | [string](#string)                  |          | network_ipv6 is the IPv6 network of the Mesh.                                                                                                                                                                                                                  |
| peers        | [WireGuardPeer](#v1.WireGuardPeer) | repeated | peers is a list of wireguard peers to connect to.                                                                                                                                                                                                              |

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

### WireGuardPeer

WireGuardPeer is a peer in the Wireguard network.

| Field               | Type              | Label    | Description                                                                  |
|---------------------|-------------------|----------|------------------------------------------------------------------------------|
| id                  | [string](#string) |          | id is the ID of the peer.                                                    |
| public_key          | [string](#string) |          | public_key is the public key of the peer.                                    |
| primary_endpoint    | [string](#string) |          | primary_endpoint is the primary endpoint of the peer.                        |
| wireguard_endpoints | [string](#string) | repeated | wireguard_endpoints are the WireGuard endpoints for the peer, if applicable. |
| zone_awareness_id   | [string](#string) |          | zone_awareness_id is the zone awareness ID of the peer.                      |
| address_ipv4        | [string](#string) |          | address_ipv4 is the private IPv4 wireguard address of the peer.              |
| address_ipv6        | [string](#string) |          | address_ipv6 is the private IPv6 wireguard address of the peer.              |
| allowed_ips         | [string](#string) | repeated | allowed_ips is the list of allowed IPs for the peer.                         |

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

| Name            | Number | Description                                               |
|-----------------|--------|-----------------------------------------------------------|
| FEATURE_NONE    | 0      | FEATURE_NONE is the default feature set.                  |
| NODES           | 1      | NODES is the feature for nodes. This is always supported. |
| LEADER_PROXY    | 2      | LEADER_PROXY is the feature for leader proxying.          |
| MESH_API        | 3      | MESH_API is the feature for the mesh API.                 |
| ADMIN_API       | 4      | ADMIN_API is the feature for the admin API.               |
| PEER_DISCOVERY  | 5      | PEER_DISCOVERY is the feature for peer discovery.         |
| METRICS         | 6      | METRICS is the feature for exposing metrics.              |
| ICE_NEGOTIATION | 7      | ICE_NEGOTIATION is the feature for ICE negotiation.       |
| TURN_SERVER     | 8      | TURN_SERVER is the feature for TURN server.               |
| MESH_DNS        | 9      | MESH_DNS is the feature for mesh DNS.                     |

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

| Field  | Type              | Label | Description                       |
|--------|-------------------|-------|-----------------------------------|
| source | [string](#string) |       | source is the source node.        |
| target | [string](#string) |       | target is the target node.        |
| weight | [int32](#int32)   |       | weight is the weight of the edge. |

### MeshGraph

MeshGraph is a graph of nodes.

| Field | Type                     | Label    | Description                                 |
|-------|--------------------------|----------|---------------------------------------------|
| nodes | [string](#string)        | repeated | nodes is the list of nodes.                 |
| edges | [MeshEdge](#v1.MeshEdge) | repeated | edges is the list of edges.                 |
| dot   | [string](#string)        |          | dot is the DOT representation of the graph. |

### MeshNode

MeshNode is a node that has been registered with the controller.

| Field               | Type                                                    | Label    | Description                                                        |
|---------------------|---------------------------------------------------------|----------|--------------------------------------------------------------------|
| id                  | [string](#string)                                       |          | id is the ID of the node.                                          |
| primary_endpoint    | [string](#string)                                       |          | primary_endpoint is the primary endpoint of the node.              |
| wireguard_endpoints | [string](#string)                                       | repeated | wireguard_endpoints is a list of WireGuard endpoints for the node. |
| zone_awareness_id   | [string](#string)                                       |          | zone_awareness_id is the zone awareness ID of the node.            |
| raft_port           | [int32](#int32)                                         |          | raft_port is the Raft listen port of the node.                     |
| grpc_port           | [int32](#int32)                                         |          | grpc_port is the gRPC listen port of the node.                     |
| public_key          | [string](#string)                                       |          | public_key is the public key of the node.                          |
| private_ipv4        | [string](#string)                                       |          | private_ipv4 is the private IPv4 address of the node.              |
| private_ipv6        | [string](#string)                                       |          | private_ipv6 is the private IPv6 address of the node.              |
| updated_at          | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |          | updated_at is the last time the node joined the cluster.           |
| created_at          | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |          | created_at is the creation time for the node.                      |
| cluster_status      | [ClusterStatus](#v1.ClusterStatus)                      |          | cluster_status is the status of the node in the cluster.           |

### NodeList

NodeList is a list of nodes.

| Field | Type                     | Label    | Description                 |
|-------|--------------------------|----------|-----------------------------|
| nodes | [MeshNode](#v1.MeshNode) | repeated | nodes is the list of nodes. |

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

## v1/raft.proto

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

| Field     | Type              | Label | Description                                        |
|-----------|-------------------|-------|----------------------------------------------------|
| node_id   | [string](#string) |       | node_id is the ID of the node to send the data to. |
| proto     | [string](#string) |       | proto is the protocol of the traffic.              |
| dst       | [string](#string) |       | dst is the destination address of the traffic.     |
| port      | [uint32](#uint32) |       | port is the destination port of the traffic.       |
| answer    | [string](#string) |       | answer is the answer to the offer.                 |
| candidate | [string](#string) |       | candidate is an ICE candidate.                     |

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
