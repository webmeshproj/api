/**
 * Copyright 2023 Avi Zimmerman <avi.zimmerman@gmail.com>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import { PromiseClient, StreamResponse, CallOptions } from "@connectrpc/connect";
import { AppDaemon } from "../v1/app_connect.js";
import { MeshEdge } from "../v1/mesh_pb.js";
import { NetworkACL, Route } from "../v1/network_acls_pb.js";
import { MeshNode } from "../v1/node_pb.js";
import { Role, RoleBinding, Group } from "../v1/rbac_pb.js";
import { 
  QueryRequest_QueryCommand,
  QueryRequest_QueryType,
  QueryRequest,
  QueryResponse
} from "../v1/storage_query_pb.js";

/**
 * AppDaemonClient is the interface for working with an AppDaemon over gRPC.
 *
 * @generated From ts-rpcdb.ts.tmpl
 */
export type AppDaemonClient = PromiseClient<typeof AppDaemon>;

/**
 * TODO: Allow the below interfaces to be constructed with a plugin query stream as well.
 */

/**
 * MeshNode is the interface for working with MeshNodes over the AppDaemon query RPC.
 *
 * @generated From ts-rpcdb.ts.tmpl
 */
export class MeshNodes {
  /**
   * @param client - The client to use for RPC calls.
   * @param connID - The connection ID to use for RPC calls.
   */
  constructor(
    private readonly client: AppDaemonClient,
    private readonly connID: string
  ) {}

  /**
   * Queries the AppDaemon for MeshNodes.
   *
   * @param query - The query to run.
   * @returns The results of the query.
   */
  private query(query: QueryRequest): Promise<QueryResponse> {
    return new Promise((resolve, reject) => {
      this.client.query({
        id: this.connID,
        query: query
      }).then((res: QueryResponse) => {
        if (res.error.length > 0) {
          reject(new Error(res.error))
          return
        }
        resolve(res)
      }).catch((err: Error) => {
        reject(err)
      })
    });
  }

  /**
   * Returns the MeshNode with the given ID.
   *
   * @param id - The ID of the node.
   * @returns The MeshNode with the given ID.
   */
  get(id: string): Promise<MeshNode> {
    return new Promise((resolve, reject) => {
      this.query(new QueryRequest({
        command: QueryRequest_QueryCommand.GET,
        type: QueryRequest_QueryType.PEERS,
        query: `id=${id}`,
      })).then((res: QueryResponse) => {
        if (res.items.length == 0) {
          reject(new Error("MeshNode not found"))
          return
        }
        resolve(MeshNode.fromJson(res.items[0].toString()))
      }).catch((err: Error) => {
        reject(err)
      })
    });
  }
  
  /**
   * Returns the MeshNode with the given pubkey.
   *
   * @param pubkey - The base64 encoded public key of the node.
   * @returns The MeshNode with the given pubkey.
   */
  getByPubkey(pubkey: string): Promise<MeshNode> {
    return new Promise((resolve, reject) => {
      this.query(new QueryRequest({
        command: QueryRequest_QueryCommand.GET,
        type: QueryRequest_QueryType.PEERS,
        query: `pubkey=${pubkey}`,
      })).then((res: QueryResponse) => {
        if (res.error.length > 0) {
          reject(new Error(res.error))
          return
        }
        if (res.items.length == 0) {
          reject(new Error("MeshNode not found"))
          return
        }
        resolve(MeshNode.fromJson(res.items[0].toString()))
      }).catch((err: Error) => {
        reject(err)
      })
    });
  }
  
  /**
   * Deletes the MeshNode with the given ID.
   *
   * @param id - The ID of the node.
   */
  delete(id: string): Promise<void> {
    return new Promise((resolve, reject) => {
      this.query(new QueryRequest({
        command: QueryRequest_QueryCommand.DELETE,
        type: QueryRequest_QueryType.PEERS,
        query: `id=${id}`,
      })).then((res: QueryResponse) => {
        if (res.error.length > 0) {
          reject(new Error(res.error))
          return
        }
        resolve()
      }).catch((err: Error) => {
        reject(err)
      })
    });
  }

  /**
   * Returns all MeshNodes.
   *
   * @returns All MeshNodes known to the AppDaemon.
   */
  list(): Promise<MeshNode[]> {
    return new Promise((resolve, reject) => {
      this.query(new QueryRequest({
        command: QueryRequest_QueryCommand.LIST,
        type: QueryRequest_QueryType.PEERS,
      })).then((res: QueryResponse) => {
        if (res.error.length > 0) {
          reject(new Error(res.error))
          return
        }
        resolve(res.items.map((item) => MeshNode.fromJson(item.toString())))
      }).catch((err: Error) => {
        reject(err)
      })
    });
  }
  
  /**
   * Puts the given MeshNode.
   *
   * @param obj - The MeshNode to put into the mesh storage.
   */
  put(obj: MeshNode): Promise<void> {
    return new Promise((resolve, reject) => {
      const enc = new TextEncoder();
      this.query(new QueryRequest({
        command: QueryRequest_QueryCommand.PUT,
        type: QueryRequest_QueryType.PEERS,
        item: enc.encode(obj.toJsonString()),
      })).then((res: QueryResponse) => {
        if (res.error.length > 0) {
          reject(new Error(res.error))
          return
        }
        resolve()
      }).catch((err: Error) => {
        reject(err)
      })
    });
  }
}
/**
 * MeshEdge is the interface for working with MeshEdges over the AppDaemon query RPC.
 *
 * @generated From ts-rpcdb.ts.tmpl
 */
export class MeshEdges {
  /**
   * @param client - The client to use for RPC calls.
   * @param connID - The connection ID to use for RPC calls.
   */
  constructor(
    private readonly client: AppDaemonClient,
    private readonly connID: string
  ) {}

  /**
   * Queries the AppDaemon for MeshEdges.
   *
   * @param query - The query to run.
   * @returns The results of the query.
   */
  private query(query: QueryRequest): Promise<QueryResponse> {
    return new Promise((resolve, reject) => {
      this.client.query({
        id: this.connID,
        query: query
      }).then((res: QueryResponse) => {
        if (res.error.length > 0) {
          reject(new Error(res.error))
          return
        }
        resolve(res)
      }).catch((err: Error) => {
        reject(err)
      })
    });
  }

  /**
   * Returns the MeshEdge with the given Sourceid and Targetid.
   *
   * @param sourceid - The ID of the source node.
   * @param targetid - The ID of the target node.
   * @returns The MeshEdge with the given Targetid and Sourceid.
   */
  get(sourceid: string, targetid: string): Promise<MeshEdge> {
    return new Promise((resolve, reject) => {
      this.query(new QueryRequest({
        command: QueryRequest_QueryCommand.GET,
        type: QueryRequest_QueryType.EDGES,
        query: `sourceid=${sourceid},targetid=${targetid}`,
      })).then((res: QueryResponse) => {
        if (res.items.length == 0) {
          reject(new Error("MeshEdge not found"))
          return
        }
        resolve(MeshEdge.fromJson(res.items[0].toString()))
      }).catch((err: Error) => {
        reject(err)
      })
    });
  }
  
  /**
   * Deletes the MeshEdge with the given Sourceid and Targetid.
   *
   * @param sourceid - The ID of the source node.
   * @param targetid - The ID of the target node.
   */
  delete(sourceid: string, targetid: string): Promise<void> {
    return new Promise((resolve, reject) => {
      this.query(new QueryRequest({
        command: QueryRequest_QueryCommand.DELETE,
        type: QueryRequest_QueryType.EDGES,
        query: `targetid=${targetid},sourceid=${sourceid}`,
      })).then((res: QueryResponse) => {
        if (res.error.length > 0) {
          reject(new Error(res.error))
          return
        }
        resolve()
      }).catch((err: Error) => {
        reject(err)
      })
    });
  }

  /**
   * Returns all MeshEdges.
   *
   * @returns All MeshEdges known to the AppDaemon.
   */
  list(): Promise<MeshEdge[]> {
    return new Promise((resolve, reject) => {
      this.query(new QueryRequest({
        command: QueryRequest_QueryCommand.LIST,
        type: QueryRequest_QueryType.EDGES,
      })).then((res: QueryResponse) => {
        if (res.error.length > 0) {
          reject(new Error(res.error))
          return
        }
        resolve(res.items.map((item) => MeshEdge.fromJson(item.toString())))
      }).catch((err: Error) => {
        reject(err)
      })
    });
  }
  
  /**
   * Puts the given MeshEdge.
   *
   * @param obj - The MeshEdge to put into the mesh storage.
   */
  put(obj: MeshEdge): Promise<void> {
    return new Promise((resolve, reject) => {
      const enc = new TextEncoder();
      this.query(new QueryRequest({
        command: QueryRequest_QueryCommand.PUT,
        type: QueryRequest_QueryType.EDGES,
        item: enc.encode(obj.toJsonString()),
      })).then((res: QueryResponse) => {
        if (res.error.length > 0) {
          reject(new Error(res.error))
          return
        }
        resolve()
      }).catch((err: Error) => {
        reject(err)
      })
    });
  }
}
/**
 * Role is the interface for working with Roles over the AppDaemon query RPC.
 *
 * @generated From ts-rpcdb.ts.tmpl
 */
export class Roles {
  /**
   * @param client - The client to use for RPC calls.
   * @param connID - The connection ID to use for RPC calls.
   */
  constructor(
    private readonly client: AppDaemonClient,
    private readonly connID: string
  ) {}

  /**
   * Queries the AppDaemon for Roles.
   *
   * @param query - The query to run.
   * @returns The results of the query.
   */
  private query(query: QueryRequest): Promise<QueryResponse> {
    return new Promise((resolve, reject) => {
      this.client.query({
        id: this.connID,
        query: query
      }).then((res: QueryResponse) => {
        if (res.error.length > 0) {
          reject(new Error(res.error))
          return
        }
        resolve(res)
      }).catch((err: Error) => {
        reject(err)
      })
    });
  }

  /**
   * Returns the Role with the given ID.
   *
   * @param id - The name of the role.
   * @returns The Role with the given ID.
   */
  get(id: string): Promise<Role> {
    return new Promise((resolve, reject) => {
      this.query(new QueryRequest({
        command: QueryRequest_QueryCommand.GET,
        type: QueryRequest_QueryType.ROLES,
        query: `id=${id}`,
      })).then((res: QueryResponse) => {
        if (res.items.length == 0) {
          reject(new Error("Role not found"))
          return
        }
        resolve(Role.fromJson(res.items[0].toString()))
      }).catch((err: Error) => {
        reject(err)
      })
    });
  }
  
  /**
   * Deletes the Role with the given ID.
   *
   * @param id - The name of the role.
   */
  delete(id: string): Promise<void> {
    return new Promise((resolve, reject) => {
      this.query(new QueryRequest({
        command: QueryRequest_QueryCommand.DELETE,
        type: QueryRequest_QueryType.ROLES,
        query: `id=${id}`,
      })).then((res: QueryResponse) => {
        if (res.error.length > 0) {
          reject(new Error(res.error))
          return
        }
        resolve()
      }).catch((err: Error) => {
        reject(err)
      })
    });
  }

  /**
   * Returns all Roles.
   *
   * @returns All Roles known to the AppDaemon.
   */
  list(): Promise<Role[]> {
    return new Promise((resolve, reject) => {
      this.query(new QueryRequest({
        command: QueryRequest_QueryCommand.LIST,
        type: QueryRequest_QueryType.ROLES,
      })).then((res: QueryResponse) => {
        if (res.error.length > 0) {
          reject(new Error(res.error))
          return
        }
        resolve(res.items.map((item) => Role.fromJson(item.toString())))
      }).catch((err: Error) => {
        reject(err)
      })
    });
  }
  
  /**
   * Returns the Roles with the given nodeid.
   *
   * @param nodeid - The ID of the node
   * @returns Any Roles found with the given nodeid.
   */
  listByNodeID(nodeid: string): Promise<Role[]> {
    return new Promise((resolve, reject) => {
      this.query(new QueryRequest({
        command: QueryRequest_QueryCommand.LIST,
        type: QueryRequest_QueryType.ROLES,
        query: `nodeid=${nodeid}`,
      })).then((res: QueryResponse) => {
        if (res.error.length > 0) {
          reject(new Error(res.error))
          return
        }
        resolve(res.items.map((item) => Role.fromJson(item.toString())))
      }).catch((err: Error) => {
        reject(err)
      })
    });
  }
  
  /**
   * Puts the given Role.
   *
   * @param obj - The Role to put into the mesh storage.
   */
  put(obj: Role): Promise<void> {
    return new Promise((resolve, reject) => {
      const enc = new TextEncoder();
      this.query(new QueryRequest({
        command: QueryRequest_QueryCommand.PUT,
        type: QueryRequest_QueryType.ROLES,
        item: enc.encode(obj.toJsonString()),
      })).then((res: QueryResponse) => {
        if (res.error.length > 0) {
          reject(new Error(res.error))
          return
        }
        resolve()
      }).catch((err: Error) => {
        reject(err)
      })
    });
  }
}
/**
 * RoleBinding is the interface for working with RoleBindings over the AppDaemon query RPC.
 *
 * @generated From ts-rpcdb.ts.tmpl
 */
export class RoleBindings {
  /**
   * @param client - The client to use for RPC calls.
   * @param connID - The connection ID to use for RPC calls.
   */
  constructor(
    private readonly client: AppDaemonClient,
    private readonly connID: string
  ) {}

  /**
   * Queries the AppDaemon for RoleBindings.
   *
   * @param query - The query to run.
   * @returns The results of the query.
   */
  private query(query: QueryRequest): Promise<QueryResponse> {
    return new Promise((resolve, reject) => {
      this.client.query({
        id: this.connID,
        query: query
      }).then((res: QueryResponse) => {
        if (res.error.length > 0) {
          reject(new Error(res.error))
          return
        }
        resolve(res)
      }).catch((err: Error) => {
        reject(err)
      })
    });
  }

  /**
   * Returns the RoleBinding with the given ID.
   *
   * @param id - The name of the rolebinding.
   * @returns The RoleBinding with the given ID.
   */
  get(id: string): Promise<RoleBinding> {
    return new Promise((resolve, reject) => {
      this.query(new QueryRequest({
        command: QueryRequest_QueryCommand.GET,
        type: QueryRequest_QueryType.ROLEBINDINGS,
        query: `id=${id}`,
      })).then((res: QueryResponse) => {
        if (res.items.length == 0) {
          reject(new Error("RoleBinding not found"))
          return
        }
        resolve(RoleBinding.fromJson(res.items[0].toString()))
      }).catch((err: Error) => {
        reject(err)
      })
    });
  }
  
  /**
   * Deletes the RoleBinding with the given ID.
   *
   * @param id - The name of the rolebinding.
   */
  delete(id: string): Promise<void> {
    return new Promise((resolve, reject) => {
      this.query(new QueryRequest({
        command: QueryRequest_QueryCommand.DELETE,
        type: QueryRequest_QueryType.ROLEBINDINGS,
        query: `id=${id}`,
      })).then((res: QueryResponse) => {
        if (res.error.length > 0) {
          reject(new Error(res.error))
          return
        }
        resolve()
      }).catch((err: Error) => {
        reject(err)
      })
    });
  }

  /**
   * Returns all RoleBindings.
   *
   * @returns All RoleBindings known to the AppDaemon.
   */
  list(): Promise<RoleBinding[]> {
    return new Promise((resolve, reject) => {
      this.query(new QueryRequest({
        command: QueryRequest_QueryCommand.LIST,
        type: QueryRequest_QueryType.ROLEBINDINGS,
      })).then((res: QueryResponse) => {
        if (res.error.length > 0) {
          reject(new Error(res.error))
          return
        }
        resolve(res.items.map((item) => RoleBinding.fromJson(item.toString())))
      }).catch((err: Error) => {
        reject(err)
      })
    });
  }
  
  /**
   * Puts the given RoleBinding.
   *
   * @param obj - The RoleBinding to put into the mesh storage.
   */
  put(obj: RoleBinding): Promise<void> {
    return new Promise((resolve, reject) => {
      const enc = new TextEncoder();
      this.query(new QueryRequest({
        command: QueryRequest_QueryCommand.PUT,
        type: QueryRequest_QueryType.ROLEBINDINGS,
        item: enc.encode(obj.toJsonString()),
      })).then((res: QueryResponse) => {
        if (res.error.length > 0) {
          reject(new Error(res.error))
          return
        }
        resolve()
      }).catch((err: Error) => {
        reject(err)
      })
    });
  }
}
/**
 * Group is the interface for working with Groups over the AppDaemon query RPC.
 *
 * @generated From ts-rpcdb.ts.tmpl
 */
export class Groups {
  /**
   * @param client - The client to use for RPC calls.
   * @param connID - The connection ID to use for RPC calls.
   */
  constructor(
    private readonly client: AppDaemonClient,
    private readonly connID: string
  ) {}

  /**
   * Queries the AppDaemon for Groups.
   *
   * @param query - The query to run.
   * @returns The results of the query.
   */
  private query(query: QueryRequest): Promise<QueryResponse> {
    return new Promise((resolve, reject) => {
      this.client.query({
        id: this.connID,
        query: query
      }).then((res: QueryResponse) => {
        if (res.error.length > 0) {
          reject(new Error(res.error))
          return
        }
        resolve(res)
      }).catch((err: Error) => {
        reject(err)
      })
    });
  }

  /**
   * Returns the Group with the given ID.
   *
   * @param id - The name of the group.
   * @returns The Group with the given ID.
   */
  get(id: string): Promise<Group> {
    return new Promise((resolve, reject) => {
      this.query(new QueryRequest({
        command: QueryRequest_QueryCommand.GET,
        type: QueryRequest_QueryType.GROUPS,
        query: `id=${id}`,
      })).then((res: QueryResponse) => {
        if (res.items.length == 0) {
          reject(new Error("Group not found"))
          return
        }
        resolve(Group.fromJson(res.items[0].toString()))
      }).catch((err: Error) => {
        reject(err)
      })
    });
  }
  
  /**
   * Deletes the Group with the given ID.
   *
   * @param id - The name of the group.
   */
  delete(id: string): Promise<void> {
    return new Promise((resolve, reject) => {
      this.query(new QueryRequest({
        command: QueryRequest_QueryCommand.DELETE,
        type: QueryRequest_QueryType.GROUPS,
        query: `id=${id}`,
      })).then((res: QueryResponse) => {
        if (res.error.length > 0) {
          reject(new Error(res.error))
          return
        }
        resolve()
      }).catch((err: Error) => {
        reject(err)
      })
    });
  }

  /**
   * Returns all Groups.
   *
   * @returns All Groups known to the AppDaemon.
   */
  list(): Promise<Group[]> {
    return new Promise((resolve, reject) => {
      this.query(new QueryRequest({
        command: QueryRequest_QueryCommand.LIST,
        type: QueryRequest_QueryType.GROUPS,
      })).then((res: QueryResponse) => {
        if (res.error.length > 0) {
          reject(new Error(res.error))
          return
        }
        resolve(res.items.map((item) => Group.fromJson(item.toString())))
      }).catch((err: Error) => {
        reject(err)
      })
    });
  }
  
  /**
   * Puts the given Group.
   *
   * @param obj - The Group to put into the mesh storage.
   */
  put(obj: Group): Promise<void> {
    return new Promise((resolve, reject) => {
      const enc = new TextEncoder();
      this.query(new QueryRequest({
        command: QueryRequest_QueryCommand.PUT,
        type: QueryRequest_QueryType.GROUPS,
        item: enc.encode(obj.toJsonString()),
      })).then((res: QueryResponse) => {
        if (res.error.length > 0) {
          reject(new Error(res.error))
          return
        }
        resolve()
      }).catch((err: Error) => {
        reject(err)
      })
    });
  }
}
/**
 * NetworkACL is the interface for working with NetworkACLs over the AppDaemon query RPC.
 *
 * @generated From ts-rpcdb.ts.tmpl
 */
export class NetworkACLs {
  /**
   * @param client - The client to use for RPC calls.
   * @param connID - The connection ID to use for RPC calls.
   */
  constructor(
    private readonly client: AppDaemonClient,
    private readonly connID: string
  ) {}

  /**
   * Queries the AppDaemon for NetworkACLs.
   *
   * @param query - The query to run.
   * @returns The results of the query.
   */
  private query(query: QueryRequest): Promise<QueryResponse> {
    return new Promise((resolve, reject) => {
      this.client.query({
        id: this.connID,
        query: query
      }).then((res: QueryResponse) => {
        if (res.error.length > 0) {
          reject(new Error(res.error))
          return
        }
        resolve(res)
      }).catch((err: Error) => {
        reject(err)
      })
    });
  }

  /**
   * Returns the NetworkACL with the given ID.
   *
   * @param id - The name of the network ACL.
   * @returns The NetworkACL with the given ID.
   */
  get(id: string): Promise<NetworkACL> {
    return new Promise((resolve, reject) => {
      this.query(new QueryRequest({
        command: QueryRequest_QueryCommand.GET,
        type: QueryRequest_QueryType.ACLS,
        query: `id=${id}`,
      })).then((res: QueryResponse) => {
        if (res.items.length == 0) {
          reject(new Error("NetworkACL not found"))
          return
        }
        resolve(NetworkACL.fromJson(res.items[0].toString()))
      }).catch((err: Error) => {
        reject(err)
      })
    });
  }
  
  /**
   * Deletes the NetworkACL with the given ID.
   *
   * @param id - The name of the network ACL.
   */
  delete(id: string): Promise<void> {
    return new Promise((resolve, reject) => {
      this.query(new QueryRequest({
        command: QueryRequest_QueryCommand.DELETE,
        type: QueryRequest_QueryType.ACLS,
        query: `id=${id}`,
      })).then((res: QueryResponse) => {
        if (res.error.length > 0) {
          reject(new Error(res.error))
          return
        }
        resolve()
      }).catch((err: Error) => {
        reject(err)
      })
    });
  }

  /**
   * Returns all NetworkACLs.
   *
   * @returns All NetworkACLs known to the AppDaemon.
   */
  list(): Promise<NetworkACL[]> {
    return new Promise((resolve, reject) => {
      this.query(new QueryRequest({
        command: QueryRequest_QueryCommand.LIST,
        type: QueryRequest_QueryType.ACLS,
      })).then((res: QueryResponse) => {
        if (res.error.length > 0) {
          reject(new Error(res.error))
          return
        }
        resolve(res.items.map((item) => NetworkACL.fromJson(item.toString())))
      }).catch((err: Error) => {
        reject(err)
      })
    });
  }
  
  /**
   * Puts the given NetworkACL.
   *
   * @param obj - The NetworkACL to put into the mesh storage.
   */
  put(obj: NetworkACL): Promise<void> {
    return new Promise((resolve, reject) => {
      const enc = new TextEncoder();
      this.query(new QueryRequest({
        command: QueryRequest_QueryCommand.PUT,
        type: QueryRequest_QueryType.ACLS,
        item: enc.encode(obj.toJsonString()),
      })).then((res: QueryResponse) => {
        if (res.error.length > 0) {
          reject(new Error(res.error))
          return
        }
        resolve()
      }).catch((err: Error) => {
        reject(err)
      })
    });
  }
}
/**
 * Route is the interface for working with Routes over the AppDaemon query RPC.
 *
 * @generated From ts-rpcdb.ts.tmpl
 */
export class Routes {
  /**
   * @param client - The client to use for RPC calls.
   * @param connID - The connection ID to use for RPC calls.
   */
  constructor(
    private readonly client: AppDaemonClient,
    private readonly connID: string
  ) {}

  /**
   * Queries the AppDaemon for Routes.
   *
   * @param query - The query to run.
   * @returns The results of the query.
   */
  private query(query: QueryRequest): Promise<QueryResponse> {
    return new Promise((resolve, reject) => {
      this.client.query({
        id: this.connID,
        query: query
      }).then((res: QueryResponse) => {
        if (res.error.length > 0) {
          reject(new Error(res.error))
          return
        }
        resolve(res)
      }).catch((err: Error) => {
        reject(err)
      })
    });
  }

  /**
   * Returns the Route with the given ID.
   *
   * @param id - The name of the route.
   * @returns The Route with the given ID.
   */
  get(id: string): Promise<Route> {
    return new Promise((resolve, reject) => {
      this.query(new QueryRequest({
        command: QueryRequest_QueryCommand.GET,
        type: QueryRequest_QueryType.ROUTES,
        query: `id=${id}`,
      })).then((res: QueryResponse) => {
        if (res.items.length == 0) {
          reject(new Error("Route not found"))
          return
        }
        resolve(Route.fromJson(res.items[0].toString()))
      }).catch((err: Error) => {
        reject(err)
      })
    });
  }
  
  /**
   * Deletes the Route with the given ID.
   *
   * @param id - The name of the route.
   */
  delete(id: string): Promise<void> {
    return new Promise((resolve, reject) => {
      this.query(new QueryRequest({
        command: QueryRequest_QueryCommand.DELETE,
        type: QueryRequest_QueryType.ROUTES,
        query: `id=${id}`,
      })).then((res: QueryResponse) => {
        if (res.error.length > 0) {
          reject(new Error(res.error))
          return
        }
        resolve()
      }).catch((err: Error) => {
        reject(err)
      })
    });
  }

  /**
   * Returns all Routes.
   *
   * @returns All Routes known to the AppDaemon.
   */
  list(): Promise<Route[]> {
    return new Promise((resolve, reject) => {
      this.query(new QueryRequest({
        command: QueryRequest_QueryCommand.LIST,
        type: QueryRequest_QueryType.ROUTES,
      })).then((res: QueryResponse) => {
        if (res.error.length > 0) {
          reject(new Error(res.error))
          return
        }
        resolve(res.items.map((item) => Route.fromJson(item.toString())))
      }).catch((err: Error) => {
        reject(err)
      })
    });
  }
  
  /**
   * Returns the Routes with the given cidr.
   *
   * @param cidr - The CIDR of the route
   * @returns Any Routes found with the given cidr.
   */
  listByCidr(cidr: string): Promise<Route[]> {
    return new Promise((resolve, reject) => {
      this.query(new QueryRequest({
        command: QueryRequest_QueryCommand.LIST,
        type: QueryRequest_QueryType.ROUTES,
        query: `cidr=${cidr}`,
      })).then((res: QueryResponse) => {
        if (res.error.length > 0) {
          reject(new Error(res.error))
          return
        }
        resolve(res.items.map((item) => Route.fromJson(item.toString())))
      }).catch((err: Error) => {
        reject(err)
      })
    });
  }
  
  /**
   * Returns the Routes with the given nodeid.
   *
   * @param nodeid - The ID of the node
   * @returns Any Routes found with the given nodeid.
   */
  listByNodeID(nodeid: string): Promise<Route[]> {
    return new Promise((resolve, reject) => {
      this.query(new QueryRequest({
        command: QueryRequest_QueryCommand.LIST,
        type: QueryRequest_QueryType.ROUTES,
        query: `nodeid=${nodeid}`,
      })).then((res: QueryResponse) => {
        if (res.error.length > 0) {
          reject(new Error(res.error))
          return
        }
        resolve(res.items.map((item) => Route.fromJson(item.toString())))
      }).catch((err: Error) => {
        reject(err)
      })
    });
  }
  
  /**
   * Puts the given Route.
   *
   * @param obj - The Route to put into the mesh storage.
   */
  put(obj: Route): Promise<void> {
    return new Promise((resolve, reject) => {
      const enc = new TextEncoder();
      this.query(new QueryRequest({
        command: QueryRequest_QueryCommand.PUT,
        type: QueryRequest_QueryType.ROUTES,
        item: enc.encode(obj.toJsonString()),
      })).then((res: QueryResponse) => {
        if (res.error.length > 0) {
          reject(new Error(res.error))
          return
        }
        resolve()
      }).catch((err: Error) => {
        reject(err)
      })
    });
  }
}

