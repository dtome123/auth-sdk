import * as jspb from 'google-protobuf'



export class ActionResource extends jspb.Message {
  getResource(): string;
  setResource(value: string): ActionResource;

  getAction(): string;
  setAction(value: string): ActionResource;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ActionResource.AsObject;
  static toObject(includeInstance: boolean, msg: ActionResource): ActionResource.AsObject;
  static serializeBinaryToWriter(message: ActionResource, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ActionResource;
  static deserializeBinaryFromReader(message: ActionResource, reader: jspb.BinaryReader): ActionResource;
}

export namespace ActionResource {
  export type AsObject = {
    resource: string,
    action: string,
  }
}

export class Permission extends jspb.Message {
  getId(): string;
  setId(value: string): Permission;

  getName(): string;
  setName(value: string): Permission;

  getDomain(): string;
  setDomain(value: string): Permission;

  getResource(): string;
  setResource(value: string): Permission;

  getAction(): string;
  setAction(value: string): Permission;

  getImpliedByActionsList(): Array<ActionResource>;
  setImpliedByActionsList(value: Array<ActionResource>): Permission;
  clearImpliedByActionsList(): Permission;
  addImpliedByActions(value?: ActionResource, index?: number): ActionResource;

  getDescription(): string;
  setDescription(value: string): Permission;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Permission.AsObject;
  static toObject(includeInstance: boolean, msg: Permission): Permission.AsObject;
  static serializeBinaryToWriter(message: Permission, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Permission;
  static deserializeBinaryFromReader(message: Permission, reader: jspb.BinaryReader): Permission;
}

export namespace Permission {
  export type AsObject = {
    id: string,
    name: string,
    domain: string,
    resource: string,
    action: string,
    impliedByActionsList: Array<ActionResource.AsObject>,
    description: string,
  }
}

export class Role extends jspb.Message {
  getId(): string;
  setId(value: string): Role;

  getName(): string;
  setName(value: string): Role;

  getPermissionIdsList(): Array<string>;
  setPermissionIdsList(value: Array<string>): Role;
  clearPermissionIdsList(): Role;
  addPermissionIds(value: string, index?: number): Role;

  getPermissionsList(): Array<Permission>;
  setPermissionsList(value: Array<Permission>): Role;
  clearPermissionsList(): Role;
  addPermissions(value?: Permission, index?: number): Permission;

  getDescription(): string;
  setDescription(value: string): Role;

  getScopeMap(): jspb.Map<string, string>;
  clearScopeMap(): Role;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Role.AsObject;
  static toObject(includeInstance: boolean, msg: Role): Role.AsObject;
  static serializeBinaryToWriter(message: Role, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Role;
  static deserializeBinaryFromReader(message: Role, reader: jspb.BinaryReader): Role;
}

export namespace Role {
  export type AsObject = {
    id: string,
    name: string,
    permissionIdsList: Array<string>,
    permissionsList: Array<Permission.AsObject>,
    description: string,
    scopeMap: Array<[string, string]>,
  }
}

export class PermissionPath extends jspb.Message {
  getId(): string;
  setId(value: string): PermissionPath;

  getDomain(): string;
  setDomain(value: string): PermissionPath;

  getPath(): string;
  setPath(value: string): PermissionPath;

  getResource(): string;
  setResource(value: string): PermissionPath;

  getAction(): string;
  setAction(value: string): PermissionPath;

  getType(): RouteScope;
  setType(value: RouteScope): PermissionPath;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PermissionPath.AsObject;
  static toObject(includeInstance: boolean, msg: PermissionPath): PermissionPath.AsObject;
  static serializeBinaryToWriter(message: PermissionPath, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PermissionPath;
  static deserializeBinaryFromReader(message: PermissionPath, reader: jspb.BinaryReader): PermissionPath;
}

export namespace PermissionPath {
  export type AsObject = {
    id: string,
    domain: string,
    path: string,
    resource: string,
    action: string,
    type: RouteScope,
  }
}

export class PageRequest extends jspb.Message {
  getPageNumber(): number;
  setPageNumber(value: number): PageRequest;

  getPageSize(): number;
  setPageSize(value: number): PageRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PageRequest.AsObject;
  static toObject(includeInstance: boolean, msg: PageRequest): PageRequest.AsObject;
  static serializeBinaryToWriter(message: PageRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PageRequest;
  static deserializeBinaryFromReader(message: PageRequest, reader: jspb.BinaryReader): PageRequest;
}

export namespace PageRequest {
  export type AsObject = {
    pageNumber: number,
    pageSize: number,
  }
}

export class PageResponse extends jspb.Message {
  getTotalPage(): number;
  setTotalPage(value: number): PageResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PageResponse.AsObject;
  static toObject(includeInstance: boolean, msg: PageResponse): PageResponse.AsObject;
  static serializeBinaryToWriter(message: PageResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PageResponse;
  static deserializeBinaryFromReader(message: PageResponse, reader: jspb.BinaryReader): PageResponse;
}

export namespace PageResponse {
  export type AsObject = {
    totalPage: number,
  }
}

export class StringList extends jspb.Message {
  getValuesList(): Array<string>;
  setValuesList(value: Array<string>): StringList;
  clearValuesList(): StringList;
  addValues(value: string, index?: number): StringList;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): StringList.AsObject;
  static toObject(includeInstance: boolean, msg: StringList): StringList.AsObject;
  static serializeBinaryToWriter(message: StringList, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): StringList;
  static deserializeBinaryFromReader(message: StringList, reader: jspb.BinaryReader): StringList;
}

export namespace StringList {
  export type AsObject = {
    valuesList: Array<string>,
  }
}

export class MigratePermissionsRequest extends jspb.Message {
  getPermissionsList(): Array<Permission>;
  setPermissionsList(value: Array<Permission>): MigratePermissionsRequest;
  clearPermissionsList(): MigratePermissionsRequest;
  addPermissions(value?: Permission, index?: number): Permission;

  getPathsList(): Array<PermissionPath>;
  setPathsList(value: Array<PermissionPath>): MigratePermissionsRequest;
  clearPathsList(): MigratePermissionsRequest;
  addPaths(value?: PermissionPath, index?: number): PermissionPath;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): MigratePermissionsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: MigratePermissionsRequest): MigratePermissionsRequest.AsObject;
  static serializeBinaryToWriter(message: MigratePermissionsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): MigratePermissionsRequest;
  static deserializeBinaryFromReader(message: MigratePermissionsRequest, reader: jspb.BinaryReader): MigratePermissionsRequest;
}

export namespace MigratePermissionsRequest {
  export type AsObject = {
    permissionsList: Array<Permission.AsObject>,
    pathsList: Array<PermissionPath.AsObject>,
  }
}

export class MigratePermissionsResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): MigratePermissionsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: MigratePermissionsResponse): MigratePermissionsResponse.AsObject;
  static serializeBinaryToWriter(message: MigratePermissionsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): MigratePermissionsResponse;
  static deserializeBinaryFromReader(message: MigratePermissionsResponse, reader: jspb.BinaryReader): MigratePermissionsResponse;
}

export namespace MigratePermissionsResponse {
  export type AsObject = {
  }
}

export class GetUserPermissionsRequest extends jspb.Message {
  getUserId(): string;
  setUserId(value: string): GetUserPermissionsRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetUserPermissionsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetUserPermissionsRequest): GetUserPermissionsRequest.AsObject;
  static serializeBinaryToWriter(message: GetUserPermissionsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetUserPermissionsRequest;
  static deserializeBinaryFromReader(message: GetUserPermissionsRequest, reader: jspb.BinaryReader): GetUserPermissionsRequest;
}

export namespace GetUserPermissionsRequest {
  export type AsObject = {
    userId: string,
  }
}

export class GetUserPermissionsResponse extends jspb.Message {
  getPermissionsList(): Array<Permission>;
  setPermissionsList(value: Array<Permission>): GetUserPermissionsResponse;
  clearPermissionsList(): GetUserPermissionsResponse;
  addPermissions(value?: Permission, index?: number): Permission;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetUserPermissionsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetUserPermissionsResponse): GetUserPermissionsResponse.AsObject;
  static serializeBinaryToWriter(message: GetUserPermissionsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetUserPermissionsResponse;
  static deserializeBinaryFromReader(message: GetUserPermissionsResponse, reader: jspb.BinaryReader): GetUserPermissionsResponse;
}

export namespace GetUserPermissionsResponse {
  export type AsObject = {
    permissionsList: Array<Permission.AsObject>,
  }
}

export class HasPermissionRequest extends jspb.Message {
  getUserId(): string;
  setUserId(value: string): HasPermissionRequest;

  getResource(): string;
  setResource(value: string): HasPermissionRequest;

  getAction(): string;
  setAction(value: string): HasPermissionRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): HasPermissionRequest.AsObject;
  static toObject(includeInstance: boolean, msg: HasPermissionRequest): HasPermissionRequest.AsObject;
  static serializeBinaryToWriter(message: HasPermissionRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): HasPermissionRequest;
  static deserializeBinaryFromReader(message: HasPermissionRequest, reader: jspb.BinaryReader): HasPermissionRequest;
}

export namespace HasPermissionRequest {
  export type AsObject = {
    userId: string,
    resource: string,
    action: string,
  }
}

export class HasPermissionResponse extends jspb.Message {
  getAllowed(): boolean;
  setAllowed(value: boolean): HasPermissionResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): HasPermissionResponse.AsObject;
  static toObject(includeInstance: boolean, msg: HasPermissionResponse): HasPermissionResponse.AsObject;
  static serializeBinaryToWriter(message: HasPermissionResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): HasPermissionResponse;
  static deserializeBinaryFromReader(message: HasPermissionResponse, reader: jspb.BinaryReader): HasPermissionResponse;
}

export namespace HasPermissionResponse {
  export type AsObject = {
    allowed: boolean,
  }
}

export class CheckPermissionsRequest extends jspb.Message {
  getUserId(): string;
  setUserId(value: string): CheckPermissionsRequest;

  getChecksList(): Array<ActionResource>;
  setChecksList(value: Array<ActionResource>): CheckPermissionsRequest;
  clearChecksList(): CheckPermissionsRequest;
  addChecks(value?: ActionResource, index?: number): ActionResource;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CheckPermissionsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CheckPermissionsRequest): CheckPermissionsRequest.AsObject;
  static serializeBinaryToWriter(message: CheckPermissionsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CheckPermissionsRequest;
  static deserializeBinaryFromReader(message: CheckPermissionsRequest, reader: jspb.BinaryReader): CheckPermissionsRequest;
}

export namespace CheckPermissionsRequest {
  export type AsObject = {
    userId: string,
    checksList: Array<ActionResource.AsObject>,
  }
}

export class CheckPermissionsResponse extends jspb.Message {
  getResultsMap(): jspb.Map<string, boolean>;
  clearResultsMap(): CheckPermissionsResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CheckPermissionsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CheckPermissionsResponse): CheckPermissionsResponse.AsObject;
  static serializeBinaryToWriter(message: CheckPermissionsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CheckPermissionsResponse;
  static deserializeBinaryFromReader(message: CheckPermissionsResponse, reader: jspb.BinaryReader): CheckPermissionsResponse;
}

export namespace CheckPermissionsResponse {
  export type AsObject = {
    resultsMap: Array<[string, boolean]>,
  }
}

export class ListPermissionPathsRequest extends jspb.Message {
  getFilter(): ListPermissionPathsRequest.Filter | undefined;
  setFilter(value?: ListPermissionPathsRequest.Filter): ListPermissionPathsRequest;
  hasFilter(): boolean;
  clearFilter(): ListPermissionPathsRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListPermissionPathsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListPermissionPathsRequest): ListPermissionPathsRequest.AsObject;
  static serializeBinaryToWriter(message: ListPermissionPathsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListPermissionPathsRequest;
  static deserializeBinaryFromReader(message: ListPermissionPathsRequest, reader: jspb.BinaryReader): ListPermissionPathsRequest;
}

export namespace ListPermissionPathsRequest {
  export type AsObject = {
    filter?: ListPermissionPathsRequest.Filter.AsObject,
  }

  export class Filter extends jspb.Message {
    getKeyword(): string;
    setKeyword(value: string): Filter;

    getDomainsList(): Array<string>;
    setDomainsList(value: Array<string>): Filter;
    clearDomainsList(): Filter;
    addDomains(value: string, index?: number): Filter;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Filter.AsObject;
    static toObject(includeInstance: boolean, msg: Filter): Filter.AsObject;
    static serializeBinaryToWriter(message: Filter, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Filter;
    static deserializeBinaryFromReader(message: Filter, reader: jspb.BinaryReader): Filter;
  }

  export namespace Filter {
    export type AsObject = {
      keyword: string,
      domainsList: Array<string>,
    }
  }

}

export class ListPermissionPathsResponse extends jspb.Message {
  getPathsList(): Array<PermissionPath>;
  setPathsList(value: Array<PermissionPath>): ListPermissionPathsResponse;
  clearPathsList(): ListPermissionPathsResponse;
  addPaths(value?: PermissionPath, index?: number): PermissionPath;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListPermissionPathsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListPermissionPathsResponse): ListPermissionPathsResponse.AsObject;
  static serializeBinaryToWriter(message: ListPermissionPathsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListPermissionPathsResponse;
  static deserializeBinaryFromReader(message: ListPermissionPathsResponse, reader: jspb.BinaryReader): ListPermissionPathsResponse;
}

export namespace ListPermissionPathsResponse {
  export type AsObject = {
    pathsList: Array<PermissionPath.AsObject>,
  }
}

export class ListPermissionsRequest extends jspb.Message {
  getPagination(): PageRequest | undefined;
  setPagination(value?: PageRequest): ListPermissionsRequest;
  hasPagination(): boolean;
  clearPagination(): ListPermissionsRequest;

  getFilter(): ListPermissionsRequest.Filter | undefined;
  setFilter(value?: ListPermissionsRequest.Filter): ListPermissionsRequest;
  hasFilter(): boolean;
  clearFilter(): ListPermissionsRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListPermissionsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListPermissionsRequest): ListPermissionsRequest.AsObject;
  static serializeBinaryToWriter(message: ListPermissionsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListPermissionsRequest;
  static deserializeBinaryFromReader(message: ListPermissionsRequest, reader: jspb.BinaryReader): ListPermissionsRequest;
}

export namespace ListPermissionsRequest {
  export type AsObject = {
    pagination?: PageRequest.AsObject,
    filter?: ListPermissionsRequest.Filter.AsObject,
  }

  export class Filter extends jspb.Message {
    getKeyword(): string;
    setKeyword(value: string): Filter;

    getDomainsList(): Array<string>;
    setDomainsList(value: Array<string>): Filter;
    clearDomainsList(): Filter;
    addDomains(value: string, index?: number): Filter;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Filter.AsObject;
    static toObject(includeInstance: boolean, msg: Filter): Filter.AsObject;
    static serializeBinaryToWriter(message: Filter, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Filter;
    static deserializeBinaryFromReader(message: Filter, reader: jspb.BinaryReader): Filter;
  }

  export namespace Filter {
    export type AsObject = {
      keyword: string,
      domainsList: Array<string>,
    }
  }

}

export class ListPermissionsResponse extends jspb.Message {
  getPermissionsList(): Array<Permission>;
  setPermissionsList(value: Array<Permission>): ListPermissionsResponse;
  clearPermissionsList(): ListPermissionsResponse;
  addPermissions(value?: Permission, index?: number): Permission;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListPermissionsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListPermissionsResponse): ListPermissionsResponse.AsObject;
  static serializeBinaryToWriter(message: ListPermissionsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListPermissionsResponse;
  static deserializeBinaryFromReader(message: ListPermissionsResponse, reader: jspb.BinaryReader): ListPermissionsResponse;
}

export namespace ListPermissionsResponse {
  export type AsObject = {
    permissionsList: Array<Permission.AsObject>,
  }
}

export class CreateRoleRequest extends jspb.Message {
  getName(): string;
  setName(value: string): CreateRoleRequest;

  getPermissionIdsList(): Array<string>;
  setPermissionIdsList(value: Array<string>): CreateRoleRequest;
  clearPermissionIdsList(): CreateRoleRequest;
  addPermissionIds(value: string, index?: number): CreateRoleRequest;

  getDescription(): string;
  setDescription(value: string): CreateRoleRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateRoleRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateRoleRequest): CreateRoleRequest.AsObject;
  static serializeBinaryToWriter(message: CreateRoleRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateRoleRequest;
  static deserializeBinaryFromReader(message: CreateRoleRequest, reader: jspb.BinaryReader): CreateRoleRequest;
}

export namespace CreateRoleRequest {
  export type AsObject = {
    name: string,
    permissionIdsList: Array<string>,
    description: string,
  }
}

export class CreateRoleResponse extends jspb.Message {
  getId(): string;
  setId(value: string): CreateRoleResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateRoleResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CreateRoleResponse): CreateRoleResponse.AsObject;
  static serializeBinaryToWriter(message: CreateRoleResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateRoleResponse;
  static deserializeBinaryFromReader(message: CreateRoleResponse, reader: jspb.BinaryReader): CreateRoleResponse;
}

export namespace CreateRoleResponse {
  export type AsObject = {
    id: string,
  }
}

export class UpdateRoleRequest extends jspb.Message {
  getId(): string;
  setId(value: string): UpdateRoleRequest;

  getName(): string;
  setName(value: string): UpdateRoleRequest;

  getPermissionIdsList(): Array<string>;
  setPermissionIdsList(value: Array<string>): UpdateRoleRequest;
  clearPermissionIdsList(): UpdateRoleRequest;
  addPermissionIds(value: string, index?: number): UpdateRoleRequest;

  getDescription(): string;
  setDescription(value: string): UpdateRoleRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateRoleRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateRoleRequest): UpdateRoleRequest.AsObject;
  static serializeBinaryToWriter(message: UpdateRoleRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateRoleRequest;
  static deserializeBinaryFromReader(message: UpdateRoleRequest, reader: jspb.BinaryReader): UpdateRoleRequest;
}

export namespace UpdateRoleRequest {
  export type AsObject = {
    id: string,
    name: string,
    permissionIdsList: Array<string>,
    description: string,
  }
}

export class UpdateRoleResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateRoleResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateRoleResponse): UpdateRoleResponse.AsObject;
  static serializeBinaryToWriter(message: UpdateRoleResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateRoleResponse;
  static deserializeBinaryFromReader(message: UpdateRoleResponse, reader: jspb.BinaryReader): UpdateRoleResponse;
}

export namespace UpdateRoleResponse {
  export type AsObject = {
  }
}

export class AssignRolesToUserRequest extends jspb.Message {
  getUserId(): string;
  setUserId(value: string): AssignRolesToUserRequest;

  getRoleIdsList(): Array<string>;
  setRoleIdsList(value: Array<string>): AssignRolesToUserRequest;
  clearRoleIdsList(): AssignRolesToUserRequest;
  addRoleIds(value: string, index?: number): AssignRolesToUserRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AssignRolesToUserRequest.AsObject;
  static toObject(includeInstance: boolean, msg: AssignRolesToUserRequest): AssignRolesToUserRequest.AsObject;
  static serializeBinaryToWriter(message: AssignRolesToUserRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AssignRolesToUserRequest;
  static deserializeBinaryFromReader(message: AssignRolesToUserRequest, reader: jspb.BinaryReader): AssignRolesToUserRequest;
}

export namespace AssignRolesToUserRequest {
  export type AsObject = {
    userId: string,
    roleIdsList: Array<string>,
  }
}

export class AssignRolesToUserResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AssignRolesToUserResponse.AsObject;
  static toObject(includeInstance: boolean, msg: AssignRolesToUserResponse): AssignRolesToUserResponse.AsObject;
  static serializeBinaryToWriter(message: AssignRolesToUserResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AssignRolesToUserResponse;
  static deserializeBinaryFromReader(message: AssignRolesToUserResponse, reader: jspb.BinaryReader): AssignRolesToUserResponse;
}

export namespace AssignRolesToUserResponse {
  export type AsObject = {
  }
}

export class GetUserRolesRequest extends jspb.Message {
  getUserId(): string;
  setUserId(value: string): GetUserRolesRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetUserRolesRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetUserRolesRequest): GetUserRolesRequest.AsObject;
  static serializeBinaryToWriter(message: GetUserRolesRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetUserRolesRequest;
  static deserializeBinaryFromReader(message: GetUserRolesRequest, reader: jspb.BinaryReader): GetUserRolesRequest;
}

export namespace GetUserRolesRequest {
  export type AsObject = {
    userId: string,
  }
}

export class GetUserRolesResponse extends jspb.Message {
  getRolesList(): Array<Role>;
  setRolesList(value: Array<Role>): GetUserRolesResponse;
  clearRolesList(): GetUserRolesResponse;
  addRoles(value?: Role, index?: number): Role;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetUserRolesResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetUserRolesResponse): GetUserRolesResponse.AsObject;
  static serializeBinaryToWriter(message: GetUserRolesResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetUserRolesResponse;
  static deserializeBinaryFromReader(message: GetUserRolesResponse, reader: jspb.BinaryReader): GetUserRolesResponse;
}

export namespace GetUserRolesResponse {
  export type AsObject = {
    rolesList: Array<Role.AsObject>,
  }
}

export class ListRoleRequest extends jspb.Message {
  getPagination(): PageRequest | undefined;
  setPagination(value?: PageRequest): ListRoleRequest;
  hasPagination(): boolean;
  clearPagination(): ListRoleRequest;

  getFilter(): ListRoleRequest.Filter | undefined;
  setFilter(value?: ListRoleRequest.Filter): ListRoleRequest;
  hasFilter(): boolean;
  clearFilter(): ListRoleRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListRoleRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListRoleRequest): ListRoleRequest.AsObject;
  static serializeBinaryToWriter(message: ListRoleRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListRoleRequest;
  static deserializeBinaryFromReader(message: ListRoleRequest, reader: jspb.BinaryReader): ListRoleRequest;
}

export namespace ListRoleRequest {
  export type AsObject = {
    pagination?: PageRequest.AsObject,
    filter?: ListRoleRequest.Filter.AsObject,
  }

  export class Filter extends jspb.Message {
    getKeyword(): string;
    setKeyword(value: string): Filter;

    getScopeMap(): jspb.Map<string, StringList>;
    clearScopeMap(): Filter;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Filter.AsObject;
    static toObject(includeInstance: boolean, msg: Filter): Filter.AsObject;
    static serializeBinaryToWriter(message: Filter, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Filter;
    static deserializeBinaryFromReader(message: Filter, reader: jspb.BinaryReader): Filter;
  }

  export namespace Filter {
    export type AsObject = {
      keyword: string,
      scopeMap: Array<[string, StringList.AsObject]>,
    }
  }

}

export class ListRoleResponse extends jspb.Message {
  getRolesList(): Array<Role>;
  setRolesList(value: Array<Role>): ListRoleResponse;
  clearRolesList(): ListRoleResponse;
  addRoles(value?: Role, index?: number): Role;

  getPagination(): PageResponse | undefined;
  setPagination(value?: PageResponse): ListRoleResponse;
  hasPagination(): boolean;
  clearPagination(): ListRoleResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListRoleResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListRoleResponse): ListRoleResponse.AsObject;
  static serializeBinaryToWriter(message: ListRoleResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListRoleResponse;
  static deserializeBinaryFromReader(message: ListRoleResponse, reader: jspb.BinaryReader): ListRoleResponse;
}

export namespace ListRoleResponse {
  export type AsObject = {
    rolesList: Array<Role.AsObject>,
    pagination?: PageResponse.AsObject,
  }
}

export enum RouteScope { 
  ROUTE_SCOPE_UNSPECIFIED = 0,
  ROUTE_SCOPE_PUBLIC = 1,
  ROUTE_SCOPE_AUTHENTICATED = 2,
  ROUTE_SCOPE_AUTHORIZED = 3,
}
