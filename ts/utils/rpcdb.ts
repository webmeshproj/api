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

import { PromiseClient } from "@connectrpc/connect";
import { AppDaemon } from "../v1/app_connect.js";
import { MeshEdge } from "../v1/mesh_pb.js";
import { NetworkACL, Route } from "../v1/network_acls_pb.js";
import { MeshNode } from "../v1/node_pb.js";
import { Role, RoleBinding, Group } from "../v1/rbac_pb.js";
import { 
  QueryRequest_QueryCommand,
  QueryRequest_QueryType,
  QueryResponse
} from "../v1/storage_query_pb.js";


/**
 * MeshNode is the interface for working with MeshNodes over the AppDaemon query RPC.
 *
 * @generated from ts-rpcdb.ts.tmpl
 */
export class MeshNodes {
  /**
   * @param client - the client to use for RPC calls
   * @param connID - the connection ID to use for RPC calls
   */
  constructor(private readonly client: PromiseClient<typeof AppDaemon>, private readonly connID: string) {
  }

  /**
   * get returns the MeshNode with the given ID.
   *
   * @param id - the ID of the MeshNode to get
   * @returns the MeshNode with the given ID
   */
  get(id: string): Promise<MeshNode> {
    return new Promise((resolve, reject) => {
      this.client.query({
        id: this.connID,
        query: {
          command: QueryRequest_QueryCommand.GET,
          type: QueryRequest_QueryType.PEERS,
          query: 'id=' + id,
        }
      }).then((res: QueryResponse) => {
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
   * delete deletes the MeshNode with the given ID.
   *
   * @param id - the ID of the MeshNode to delete
   */
  delete(id: string): Promise<void> {
    return new Promise((resolve, reject) => {
      this.client.query({
        id: this.connID,
        query: {
          command: QueryRequest_QueryCommand.DELETE,
          type: QueryRequest_QueryType.PEERS,
          query: 'id=' + id,
        }
      }).then(() => {
        resolve()
      }).catch((err: Error) => {
        reject(err)
      })
    });
  }

  /**
   * list returns all MeshNodes.
   *
   * @returns all MeshNodes
   */
  list(): Promise<MeshNode[]> {
    return new Promise((resolve, reject) => {
      this.client.query({
        id: this.connID,
        query: {
          command: QueryRequest_QueryCommand.LIST,
          type: QueryRequest_QueryType.PEERS,
        }
      }).then((res: QueryResponse) => {
        resolve(res.items.map((item) => MeshNode.fromJson(item.toString())))
      }).catch((err: Error) => {
        reject(err)
      })
    });
  }

  /**
   * put puts the given MeshNode.
   *
   * @param obj - the MeshNode to put
   */
  put(obj: MeshNode): Promise<void> {
    return new Promise((resolve, reject) => {
      const enc = new TextEncoder();
      this.client.query({
        id: this.connID,
        query: {
          command: QueryRequest_QueryCommand.PUT,
          type: QueryRequest_QueryType.PEERS,
          item: enc.encode(obj.toJsonString()),
        }
      }).then(() => {
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
 * @generated from ts-rpcdb.ts.tmpl
 */
export class MeshEdges {
  /**
   * @param client - the client to use for RPC calls
   * @param connID - the connection ID to use for RPC calls
   */
  constructor(private readonly client: PromiseClient<typeof AppDaemon>, private readonly connID: string) {
  }

  /**
   * get returns the MeshEdge with the given ID.
   *
   * @param id - the ID of the MeshEdge to get
   * @returns the MeshEdge with the given ID
   */
  get(id: string): Promise<MeshEdge> {
    return new Promise((resolve, reject) => {
      this.client.query({
        id: this.connID,
        query: {
          command: QueryRequest_QueryCommand.GET,
          type: QueryRequest_QueryType.EDGES,
          query: 'id=' + id,
        }
      }).then((res: QueryResponse) => {
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
   * delete deletes the MeshEdge with the given ID.
   *
   * @param id - the ID of the MeshEdge to delete
   */
  delete(id: string): Promise<void> {
    return new Promise((resolve, reject) => {
      this.client.query({
        id: this.connID,
        query: {
          command: QueryRequest_QueryCommand.DELETE,
          type: QueryRequest_QueryType.EDGES,
          query: 'id=' + id,
        }
      }).then(() => {
        resolve()
      }).catch((err: Error) => {
        reject(err)
      })
    });
  }

  /**
   * list returns all MeshEdges.
   *
   * @returns all MeshEdges
   */
  list(): Promise<MeshEdge[]> {
    return new Promise((resolve, reject) => {
      this.client.query({
        id: this.connID,
        query: {
          command: QueryRequest_QueryCommand.LIST,
          type: QueryRequest_QueryType.EDGES,
        }
      }).then((res: QueryResponse) => {
        resolve(res.items.map((item) => MeshEdge.fromJson(item.toString())))
      }).catch((err: Error) => {
        reject(err)
      })
    });
  }

  /**
   * put puts the given MeshEdge.
   *
   * @param obj - the MeshEdge to put
   */
  put(obj: MeshEdge): Promise<void> {
    return new Promise((resolve, reject) => {
      const enc = new TextEncoder();
      this.client.query({
        id: this.connID,
        query: {
          command: QueryRequest_QueryCommand.PUT,
          type: QueryRequest_QueryType.EDGES,
          item: enc.encode(obj.toJsonString()),
        }
      }).then(() => {
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
 * @generated from ts-rpcdb.ts.tmpl
 */
export class Roles {
  /**
   * @param client - the client to use for RPC calls
   * @param connID - the connection ID to use for RPC calls
   */
  constructor(private readonly client: PromiseClient<typeof AppDaemon>, private readonly connID: string) {
  }

  /**
   * get returns the Role with the given ID.
   *
   * @param id - the ID of the Role to get
   * @returns the Role with the given ID
   */
  get(id: string): Promise<Role> {
    return new Promise((resolve, reject) => {
      this.client.query({
        id: this.connID,
        query: {
          command: QueryRequest_QueryCommand.GET,
          type: QueryRequest_QueryType.ROLES,
          query: 'id=' + id,
        }
      }).then((res: QueryResponse) => {
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
   * delete deletes the Role with the given ID.
   *
   * @param id - the ID of the Role to delete
   */
  delete(id: string): Promise<void> {
    return new Promise((resolve, reject) => {
      this.client.query({
        id: this.connID,
        query: {
          command: QueryRequest_QueryCommand.DELETE,
          type: QueryRequest_QueryType.ROLES,
          query: 'id=' + id,
        }
      }).then(() => {
        resolve()
      }).catch((err: Error) => {
        reject(err)
      })
    });
  }

  /**
   * list returns all Roles.
   *
   * @returns all Roles
   */
  list(): Promise<Role[]> {
    return new Promise((resolve, reject) => {
      this.client.query({
        id: this.connID,
        query: {
          command: QueryRequest_QueryCommand.LIST,
          type: QueryRequest_QueryType.ROLES,
        }
      }).then((res: QueryResponse) => {
        resolve(res.items.map((item) => Role.fromJson(item.toString())))
      }).catch((err: Error) => {
        reject(err)
      })
    });
  }

  /**
   * put puts the given Role.
   *
   * @param obj - the Role to put
   */
  put(obj: Role): Promise<void> {
    return new Promise((resolve, reject) => {
      const enc = new TextEncoder();
      this.client.query({
        id: this.connID,
        query: {
          command: QueryRequest_QueryCommand.PUT,
          type: QueryRequest_QueryType.ROLES,
          item: enc.encode(obj.toJsonString()),
        }
      }).then(() => {
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
 * @generated from ts-rpcdb.ts.tmpl
 */
export class RoleBindings {
  /**
   * @param client - the client to use for RPC calls
   * @param connID - the connection ID to use for RPC calls
   */
  constructor(private readonly client: PromiseClient<typeof AppDaemon>, private readonly connID: string) {
  }

  /**
   * get returns the RoleBinding with the given ID.
   *
   * @param id - the ID of the RoleBinding to get
   * @returns the RoleBinding with the given ID
   */
  get(id: string): Promise<RoleBinding> {
    return new Promise((resolve, reject) => {
      this.client.query({
        id: this.connID,
        query: {
          command: QueryRequest_QueryCommand.GET,
          type: QueryRequest_QueryType.ROLEBINDINGS,
          query: 'id=' + id,
        }
      }).then((res: QueryResponse) => {
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
   * delete deletes the RoleBinding with the given ID.
   *
   * @param id - the ID of the RoleBinding to delete
   */
  delete(id: string): Promise<void> {
    return new Promise((resolve, reject) => {
      this.client.query({
        id: this.connID,
        query: {
          command: QueryRequest_QueryCommand.DELETE,
          type: QueryRequest_QueryType.ROLEBINDINGS,
          query: 'id=' + id,
        }
      }).then(() => {
        resolve()
      }).catch((err: Error) => {
        reject(err)
      })
    });
  }

  /**
   * list returns all RoleBindings.
   *
   * @returns all RoleBindings
   */
  list(): Promise<RoleBinding[]> {
    return new Promise((resolve, reject) => {
      this.client.query({
        id: this.connID,
        query: {
          command: QueryRequest_QueryCommand.LIST,
          type: QueryRequest_QueryType.ROLEBINDINGS,
        }
      }).then((res: QueryResponse) => {
        resolve(res.items.map((item) => RoleBinding.fromJson(item.toString())))
      }).catch((err: Error) => {
        reject(err)
      })
    });
  }

  /**
   * put puts the given RoleBinding.
   *
   * @param obj - the RoleBinding to put
   */
  put(obj: RoleBinding): Promise<void> {
    return new Promise((resolve, reject) => {
      const enc = new TextEncoder();
      this.client.query({
        id: this.connID,
        query: {
          command: QueryRequest_QueryCommand.PUT,
          type: QueryRequest_QueryType.ROLEBINDINGS,
          item: enc.encode(obj.toJsonString()),
        }
      }).then(() => {
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
 * @generated from ts-rpcdb.ts.tmpl
 */
export class Groups {
  /**
   * @param client - the client to use for RPC calls
   * @param connID - the connection ID to use for RPC calls
   */
  constructor(private readonly client: PromiseClient<typeof AppDaemon>, private readonly connID: string) {
  }

  /**
   * get returns the Group with the given ID.
   *
   * @param id - the ID of the Group to get
   * @returns the Group with the given ID
   */
  get(id: string): Promise<Group> {
    return new Promise((resolve, reject) => {
      this.client.query({
        id: this.connID,
        query: {
          command: QueryRequest_QueryCommand.GET,
          type: QueryRequest_QueryType.GROUPS,
          query: 'id=' + id,
        }
      }).then((res: QueryResponse) => {
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
   * delete deletes the Group with the given ID.
   *
   * @param id - the ID of the Group to delete
   */
  delete(id: string): Promise<void> {
    return new Promise((resolve, reject) => {
      this.client.query({
        id: this.connID,
        query: {
          command: QueryRequest_QueryCommand.DELETE,
          type: QueryRequest_QueryType.GROUPS,
          query: 'id=' + id,
        }
      }).then(() => {
        resolve()
      }).catch((err: Error) => {
        reject(err)
      })
    });
  }

  /**
   * list returns all Groups.
   *
   * @returns all Groups
   */
  list(): Promise<Group[]> {
    return new Promise((resolve, reject) => {
      this.client.query({
        id: this.connID,
        query: {
          command: QueryRequest_QueryCommand.LIST,
          type: QueryRequest_QueryType.GROUPS,
        }
      }).then((res: QueryResponse) => {
        resolve(res.items.map((item) => Group.fromJson(item.toString())))
      }).catch((err: Error) => {
        reject(err)
      })
    });
  }

  /**
   * put puts the given Group.
   *
   * @param obj - the Group to put
   */
  put(obj: Group): Promise<void> {
    return new Promise((resolve, reject) => {
      const enc = new TextEncoder();
      this.client.query({
        id: this.connID,
        query: {
          command: QueryRequest_QueryCommand.PUT,
          type: QueryRequest_QueryType.GROUPS,
          item: enc.encode(obj.toJsonString()),
        }
      }).then(() => {
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
 * @generated from ts-rpcdb.ts.tmpl
 */
export class NetworkACLs {
  /**
   * @param client - the client to use for RPC calls
   * @param connID - the connection ID to use for RPC calls
   */
  constructor(private readonly client: PromiseClient<typeof AppDaemon>, private readonly connID: string) {
  }

  /**
   * get returns the NetworkACL with the given ID.
   *
   * @param id - the ID of the NetworkACL to get
   * @returns the NetworkACL with the given ID
   */
  get(id: string): Promise<NetworkACL> {
    return new Promise((resolve, reject) => {
      this.client.query({
        id: this.connID,
        query: {
          command: QueryRequest_QueryCommand.GET,
          type: QueryRequest_QueryType.ACLS,
          query: 'id=' + id,
        }
      }).then((res: QueryResponse) => {
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
   * delete deletes the NetworkACL with the given ID.
   *
   * @param id - the ID of the NetworkACL to delete
   */
  delete(id: string): Promise<void> {
    return new Promise((resolve, reject) => {
      this.client.query({
        id: this.connID,
        query: {
          command: QueryRequest_QueryCommand.DELETE,
          type: QueryRequest_QueryType.ACLS,
          query: 'id=' + id,
        }
      }).then(() => {
        resolve()
      }).catch((err: Error) => {
        reject(err)
      })
    });
  }

  /**
   * list returns all NetworkACLs.
   *
   * @returns all NetworkACLs
   */
  list(): Promise<NetworkACL[]> {
    return new Promise((resolve, reject) => {
      this.client.query({
        id: this.connID,
        query: {
          command: QueryRequest_QueryCommand.LIST,
          type: QueryRequest_QueryType.ACLS,
        }
      }).then((res: QueryResponse) => {
        resolve(res.items.map((item) => NetworkACL.fromJson(item.toString())))
      }).catch((err: Error) => {
        reject(err)
      })
    });
  }

  /**
   * put puts the given NetworkACL.
   *
   * @param obj - the NetworkACL to put
   */
  put(obj: NetworkACL): Promise<void> {
    return new Promise((resolve, reject) => {
      const enc = new TextEncoder();
      this.client.query({
        id: this.connID,
        query: {
          command: QueryRequest_QueryCommand.PUT,
          type: QueryRequest_QueryType.ACLS,
          item: enc.encode(obj.toJsonString()),
        }
      }).then(() => {
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
 * @generated from ts-rpcdb.ts.tmpl
 */
export class Routes {
  /**
   * @param client - the client to use for RPC calls
   * @param connID - the connection ID to use for RPC calls
   */
  constructor(private readonly client: PromiseClient<typeof AppDaemon>, private readonly connID: string) {
  }

  /**
   * get returns the Route with the given ID.
   *
   * @param id - the ID of the Route to get
   * @returns the Route with the given ID
   */
  get(id: string): Promise<Route> {
    return new Promise((resolve, reject) => {
      this.client.query({
        id: this.connID,
        query: {
          command: QueryRequest_QueryCommand.GET,
          type: QueryRequest_QueryType.ROUTES,
          query: 'id=' + id,
        }
      }).then((res: QueryResponse) => {
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
   * delete deletes the Route with the given ID.
   *
   * @param id - the ID of the Route to delete
   */
  delete(id: string): Promise<void> {
    return new Promise((resolve, reject) => {
      this.client.query({
        id: this.connID,
        query: {
          command: QueryRequest_QueryCommand.DELETE,
          type: QueryRequest_QueryType.ROUTES,
          query: 'id=' + id,
        }
      }).then(() => {
        resolve()
      }).catch((err: Error) => {
        reject(err)
      })
    });
  }

  /**
   * list returns all Routes.
   *
   * @returns all Routes
   */
  list(): Promise<Route[]> {
    return new Promise((resolve, reject) => {
      this.client.query({
        id: this.connID,
        query: {
          command: QueryRequest_QueryCommand.LIST,
          type: QueryRequest_QueryType.ROUTES,
        }
      }).then((res: QueryResponse) => {
        resolve(res.items.map((item) => Route.fromJson(item.toString())))
      }).catch((err: Error) => {
        reject(err)
      })
    });
  }

  /**
   * put puts the given Route.
   *
   * @param obj - the Route to put
   */
  put(obj: Route): Promise<void> {
    return new Promise((resolve, reject) => {
      const enc = new TextEncoder();
      this.client.query({
        id: this.connID,
        query: {
          command: QueryRequest_QueryCommand.PUT,
          type: QueryRequest_QueryType.ROUTES,
          item: enc.encode(obj.toJsonString()),
        }
      }).then(() => {
        resolve()
      }).catch((err: Error) => {
        reject(err)
      })
    });
  }
}

