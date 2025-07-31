// package: auth.v1
// file: proto/auth/v1/authorization.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "grpc";
import * as proto_auth_v1_authorization_pb from "../../../proto/auth/v1/authorization_pb";

interface IAuthorizationServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    getUserPermissions: IAuthorizationServiceService_IGetUserPermissions;
    hasPermission: IAuthorizationServiceService_IHasPermission;
    checkPermissions: IAuthorizationServiceService_ICheckPermissions;
    listPermissionPaths: IAuthorizationServiceService_IListPermissionPaths;
    listPermissions: IAuthorizationServiceService_IListPermissions;
    migratePermissions: IAuthorizationServiceService_IMigratePermissions;
    assignRolesToUser: IAuthorizationServiceService_IAssignRolesToUser;
    getUserRoles: IAuthorizationServiceService_IGetUserRoles;
    createRole: IAuthorizationServiceService_ICreateRole;
    updateRole: IAuthorizationServiceService_IUpdateRole;
    listRole: IAuthorizationServiceService_IListRole;
}

interface IAuthorizationServiceService_IGetUserPermissions extends grpc.MethodDefinition<proto_auth_v1_authorization_pb.GetUserPermissionsRequest, proto_auth_v1_authorization_pb.GetUserPermissionsResponse> {
    path: "/auth.v1.AuthorizationService/GetUserPermissions";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_auth_v1_authorization_pb.GetUserPermissionsRequest>;
    requestDeserialize: grpc.deserialize<proto_auth_v1_authorization_pb.GetUserPermissionsRequest>;
    responseSerialize: grpc.serialize<proto_auth_v1_authorization_pb.GetUserPermissionsResponse>;
    responseDeserialize: grpc.deserialize<proto_auth_v1_authorization_pb.GetUserPermissionsResponse>;
}
interface IAuthorizationServiceService_IHasPermission extends grpc.MethodDefinition<proto_auth_v1_authorization_pb.HasPermissionRequest, proto_auth_v1_authorization_pb.HasPermissionResponse> {
    path: "/auth.v1.AuthorizationService/HasPermission";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_auth_v1_authorization_pb.HasPermissionRequest>;
    requestDeserialize: grpc.deserialize<proto_auth_v1_authorization_pb.HasPermissionRequest>;
    responseSerialize: grpc.serialize<proto_auth_v1_authorization_pb.HasPermissionResponse>;
    responseDeserialize: grpc.deserialize<proto_auth_v1_authorization_pb.HasPermissionResponse>;
}
interface IAuthorizationServiceService_ICheckPermissions extends grpc.MethodDefinition<proto_auth_v1_authorization_pb.CheckPermissionsRequest, proto_auth_v1_authorization_pb.CheckPermissionsResponse> {
    path: "/auth.v1.AuthorizationService/CheckPermissions";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_auth_v1_authorization_pb.CheckPermissionsRequest>;
    requestDeserialize: grpc.deserialize<proto_auth_v1_authorization_pb.CheckPermissionsRequest>;
    responseSerialize: grpc.serialize<proto_auth_v1_authorization_pb.CheckPermissionsResponse>;
    responseDeserialize: grpc.deserialize<proto_auth_v1_authorization_pb.CheckPermissionsResponse>;
}
interface IAuthorizationServiceService_IListPermissionPaths extends grpc.MethodDefinition<proto_auth_v1_authorization_pb.ListPermissionPathsRequest, proto_auth_v1_authorization_pb.ListPermissionPathsResponse> {
    path: "/auth.v1.AuthorizationService/ListPermissionPaths";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_auth_v1_authorization_pb.ListPermissionPathsRequest>;
    requestDeserialize: grpc.deserialize<proto_auth_v1_authorization_pb.ListPermissionPathsRequest>;
    responseSerialize: grpc.serialize<proto_auth_v1_authorization_pb.ListPermissionPathsResponse>;
    responseDeserialize: grpc.deserialize<proto_auth_v1_authorization_pb.ListPermissionPathsResponse>;
}
interface IAuthorizationServiceService_IListPermissions extends grpc.MethodDefinition<proto_auth_v1_authorization_pb.ListPermissionsRequest, proto_auth_v1_authorization_pb.ListPermissionsResponse> {
    path: "/auth.v1.AuthorizationService/ListPermissions";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_auth_v1_authorization_pb.ListPermissionsRequest>;
    requestDeserialize: grpc.deserialize<proto_auth_v1_authorization_pb.ListPermissionsRequest>;
    responseSerialize: grpc.serialize<proto_auth_v1_authorization_pb.ListPermissionsResponse>;
    responseDeserialize: grpc.deserialize<proto_auth_v1_authorization_pb.ListPermissionsResponse>;
}
interface IAuthorizationServiceService_IMigratePermissions extends grpc.MethodDefinition<proto_auth_v1_authorization_pb.MigratePermissionsRequest, proto_auth_v1_authorization_pb.MigratePermissionsResponse> {
    path: "/auth.v1.AuthorizationService/MigratePermissions";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_auth_v1_authorization_pb.MigratePermissionsRequest>;
    requestDeserialize: grpc.deserialize<proto_auth_v1_authorization_pb.MigratePermissionsRequest>;
    responseSerialize: grpc.serialize<proto_auth_v1_authorization_pb.MigratePermissionsResponse>;
    responseDeserialize: grpc.deserialize<proto_auth_v1_authorization_pb.MigratePermissionsResponse>;
}
interface IAuthorizationServiceService_IAssignRolesToUser extends grpc.MethodDefinition<proto_auth_v1_authorization_pb.AssignRolesToUserRequest, proto_auth_v1_authorization_pb.AssignRolesToUserResponse> {
    path: "/auth.v1.AuthorizationService/AssignRolesToUser";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_auth_v1_authorization_pb.AssignRolesToUserRequest>;
    requestDeserialize: grpc.deserialize<proto_auth_v1_authorization_pb.AssignRolesToUserRequest>;
    responseSerialize: grpc.serialize<proto_auth_v1_authorization_pb.AssignRolesToUserResponse>;
    responseDeserialize: grpc.deserialize<proto_auth_v1_authorization_pb.AssignRolesToUserResponse>;
}
interface IAuthorizationServiceService_IGetUserRoles extends grpc.MethodDefinition<proto_auth_v1_authorization_pb.GetUserRolesRequest, proto_auth_v1_authorization_pb.GetUserRolesResponse> {
    path: "/auth.v1.AuthorizationService/GetUserRoles";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_auth_v1_authorization_pb.GetUserRolesRequest>;
    requestDeserialize: grpc.deserialize<proto_auth_v1_authorization_pb.GetUserRolesRequest>;
    responseSerialize: grpc.serialize<proto_auth_v1_authorization_pb.GetUserRolesResponse>;
    responseDeserialize: grpc.deserialize<proto_auth_v1_authorization_pb.GetUserRolesResponse>;
}
interface IAuthorizationServiceService_ICreateRole extends grpc.MethodDefinition<proto_auth_v1_authorization_pb.CreateRoleRequest, proto_auth_v1_authorization_pb.CreateRoleResponse> {
    path: "/auth.v1.AuthorizationService/CreateRole";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_auth_v1_authorization_pb.CreateRoleRequest>;
    requestDeserialize: grpc.deserialize<proto_auth_v1_authorization_pb.CreateRoleRequest>;
    responseSerialize: grpc.serialize<proto_auth_v1_authorization_pb.CreateRoleResponse>;
    responseDeserialize: grpc.deserialize<proto_auth_v1_authorization_pb.CreateRoleResponse>;
}
interface IAuthorizationServiceService_IUpdateRole extends grpc.MethodDefinition<proto_auth_v1_authorization_pb.UpdateRoleRequest, proto_auth_v1_authorization_pb.UpdateRoleResponse> {
    path: "/auth.v1.AuthorizationService/UpdateRole";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_auth_v1_authorization_pb.UpdateRoleRequest>;
    requestDeserialize: grpc.deserialize<proto_auth_v1_authorization_pb.UpdateRoleRequest>;
    responseSerialize: grpc.serialize<proto_auth_v1_authorization_pb.UpdateRoleResponse>;
    responseDeserialize: grpc.deserialize<proto_auth_v1_authorization_pb.UpdateRoleResponse>;
}
interface IAuthorizationServiceService_IListRole extends grpc.MethodDefinition<proto_auth_v1_authorization_pb.ListRoleRequest, proto_auth_v1_authorization_pb.ListRoleResponse> {
    path: "/auth.v1.AuthorizationService/ListRole";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_auth_v1_authorization_pb.ListRoleRequest>;
    requestDeserialize: grpc.deserialize<proto_auth_v1_authorization_pb.ListRoleRequest>;
    responseSerialize: grpc.serialize<proto_auth_v1_authorization_pb.ListRoleResponse>;
    responseDeserialize: grpc.deserialize<proto_auth_v1_authorization_pb.ListRoleResponse>;
}

export const AuthorizationServiceService: IAuthorizationServiceService;

export interface IAuthorizationServiceServer {
    getUserPermissions: grpc.handleUnaryCall<proto_auth_v1_authorization_pb.GetUserPermissionsRequest, proto_auth_v1_authorization_pb.GetUserPermissionsResponse>;
    hasPermission: grpc.handleUnaryCall<proto_auth_v1_authorization_pb.HasPermissionRequest, proto_auth_v1_authorization_pb.HasPermissionResponse>;
    checkPermissions: grpc.handleUnaryCall<proto_auth_v1_authorization_pb.CheckPermissionsRequest, proto_auth_v1_authorization_pb.CheckPermissionsResponse>;
    listPermissionPaths: grpc.handleUnaryCall<proto_auth_v1_authorization_pb.ListPermissionPathsRequest, proto_auth_v1_authorization_pb.ListPermissionPathsResponse>;
    listPermissions: grpc.handleUnaryCall<proto_auth_v1_authorization_pb.ListPermissionsRequest, proto_auth_v1_authorization_pb.ListPermissionsResponse>;
    migratePermissions: grpc.handleUnaryCall<proto_auth_v1_authorization_pb.MigratePermissionsRequest, proto_auth_v1_authorization_pb.MigratePermissionsResponse>;
    assignRolesToUser: grpc.handleUnaryCall<proto_auth_v1_authorization_pb.AssignRolesToUserRequest, proto_auth_v1_authorization_pb.AssignRolesToUserResponse>;
    getUserRoles: grpc.handleUnaryCall<proto_auth_v1_authorization_pb.GetUserRolesRequest, proto_auth_v1_authorization_pb.GetUserRolesResponse>;
    createRole: grpc.handleUnaryCall<proto_auth_v1_authorization_pb.CreateRoleRequest, proto_auth_v1_authorization_pb.CreateRoleResponse>;
    updateRole: grpc.handleUnaryCall<proto_auth_v1_authorization_pb.UpdateRoleRequest, proto_auth_v1_authorization_pb.UpdateRoleResponse>;
    listRole: grpc.handleUnaryCall<proto_auth_v1_authorization_pb.ListRoleRequest, proto_auth_v1_authorization_pb.ListRoleResponse>;
}

export interface IAuthorizationServiceClient {
    getUserPermissions(request: proto_auth_v1_authorization_pb.GetUserPermissionsRequest, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_authorization_pb.GetUserPermissionsResponse) => void): grpc.ClientUnaryCall;
    getUserPermissions(request: proto_auth_v1_authorization_pb.GetUserPermissionsRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_authorization_pb.GetUserPermissionsResponse) => void): grpc.ClientUnaryCall;
    getUserPermissions(request: proto_auth_v1_authorization_pb.GetUserPermissionsRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_authorization_pb.GetUserPermissionsResponse) => void): grpc.ClientUnaryCall;
    hasPermission(request: proto_auth_v1_authorization_pb.HasPermissionRequest, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_authorization_pb.HasPermissionResponse) => void): grpc.ClientUnaryCall;
    hasPermission(request: proto_auth_v1_authorization_pb.HasPermissionRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_authorization_pb.HasPermissionResponse) => void): grpc.ClientUnaryCall;
    hasPermission(request: proto_auth_v1_authorization_pb.HasPermissionRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_authorization_pb.HasPermissionResponse) => void): grpc.ClientUnaryCall;
    checkPermissions(request: proto_auth_v1_authorization_pb.CheckPermissionsRequest, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_authorization_pb.CheckPermissionsResponse) => void): grpc.ClientUnaryCall;
    checkPermissions(request: proto_auth_v1_authorization_pb.CheckPermissionsRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_authorization_pb.CheckPermissionsResponse) => void): grpc.ClientUnaryCall;
    checkPermissions(request: proto_auth_v1_authorization_pb.CheckPermissionsRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_authorization_pb.CheckPermissionsResponse) => void): grpc.ClientUnaryCall;
    listPermissionPaths(request: proto_auth_v1_authorization_pb.ListPermissionPathsRequest, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_authorization_pb.ListPermissionPathsResponse) => void): grpc.ClientUnaryCall;
    listPermissionPaths(request: proto_auth_v1_authorization_pb.ListPermissionPathsRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_authorization_pb.ListPermissionPathsResponse) => void): grpc.ClientUnaryCall;
    listPermissionPaths(request: proto_auth_v1_authorization_pb.ListPermissionPathsRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_authorization_pb.ListPermissionPathsResponse) => void): grpc.ClientUnaryCall;
    listPermissions(request: proto_auth_v1_authorization_pb.ListPermissionsRequest, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_authorization_pb.ListPermissionsResponse) => void): grpc.ClientUnaryCall;
    listPermissions(request: proto_auth_v1_authorization_pb.ListPermissionsRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_authorization_pb.ListPermissionsResponse) => void): grpc.ClientUnaryCall;
    listPermissions(request: proto_auth_v1_authorization_pb.ListPermissionsRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_authorization_pb.ListPermissionsResponse) => void): grpc.ClientUnaryCall;
    migratePermissions(request: proto_auth_v1_authorization_pb.MigratePermissionsRequest, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_authorization_pb.MigratePermissionsResponse) => void): grpc.ClientUnaryCall;
    migratePermissions(request: proto_auth_v1_authorization_pb.MigratePermissionsRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_authorization_pb.MigratePermissionsResponse) => void): grpc.ClientUnaryCall;
    migratePermissions(request: proto_auth_v1_authorization_pb.MigratePermissionsRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_authorization_pb.MigratePermissionsResponse) => void): grpc.ClientUnaryCall;
    assignRolesToUser(request: proto_auth_v1_authorization_pb.AssignRolesToUserRequest, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_authorization_pb.AssignRolesToUserResponse) => void): grpc.ClientUnaryCall;
    assignRolesToUser(request: proto_auth_v1_authorization_pb.AssignRolesToUserRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_authorization_pb.AssignRolesToUserResponse) => void): grpc.ClientUnaryCall;
    assignRolesToUser(request: proto_auth_v1_authorization_pb.AssignRolesToUserRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_authorization_pb.AssignRolesToUserResponse) => void): grpc.ClientUnaryCall;
    getUserRoles(request: proto_auth_v1_authorization_pb.GetUserRolesRequest, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_authorization_pb.GetUserRolesResponse) => void): grpc.ClientUnaryCall;
    getUserRoles(request: proto_auth_v1_authorization_pb.GetUserRolesRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_authorization_pb.GetUserRolesResponse) => void): grpc.ClientUnaryCall;
    getUserRoles(request: proto_auth_v1_authorization_pb.GetUserRolesRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_authorization_pb.GetUserRolesResponse) => void): grpc.ClientUnaryCall;
    createRole(request: proto_auth_v1_authorization_pb.CreateRoleRequest, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_authorization_pb.CreateRoleResponse) => void): grpc.ClientUnaryCall;
    createRole(request: proto_auth_v1_authorization_pb.CreateRoleRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_authorization_pb.CreateRoleResponse) => void): grpc.ClientUnaryCall;
    createRole(request: proto_auth_v1_authorization_pb.CreateRoleRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_authorization_pb.CreateRoleResponse) => void): grpc.ClientUnaryCall;
    updateRole(request: proto_auth_v1_authorization_pb.UpdateRoleRequest, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_authorization_pb.UpdateRoleResponse) => void): grpc.ClientUnaryCall;
    updateRole(request: proto_auth_v1_authorization_pb.UpdateRoleRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_authorization_pb.UpdateRoleResponse) => void): grpc.ClientUnaryCall;
    updateRole(request: proto_auth_v1_authorization_pb.UpdateRoleRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_authorization_pb.UpdateRoleResponse) => void): grpc.ClientUnaryCall;
    listRole(request: proto_auth_v1_authorization_pb.ListRoleRequest, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_authorization_pb.ListRoleResponse) => void): grpc.ClientUnaryCall;
    listRole(request: proto_auth_v1_authorization_pb.ListRoleRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_authorization_pb.ListRoleResponse) => void): grpc.ClientUnaryCall;
    listRole(request: proto_auth_v1_authorization_pb.ListRoleRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_authorization_pb.ListRoleResponse) => void): grpc.ClientUnaryCall;
}

export class AuthorizationServiceClient extends grpc.Client implements IAuthorizationServiceClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
    public getUserPermissions(request: proto_auth_v1_authorization_pb.GetUserPermissionsRequest, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_authorization_pb.GetUserPermissionsResponse) => void): grpc.ClientUnaryCall;
    public getUserPermissions(request: proto_auth_v1_authorization_pb.GetUserPermissionsRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_authorization_pb.GetUserPermissionsResponse) => void): grpc.ClientUnaryCall;
    public getUserPermissions(request: proto_auth_v1_authorization_pb.GetUserPermissionsRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_authorization_pb.GetUserPermissionsResponse) => void): grpc.ClientUnaryCall;
    public hasPermission(request: proto_auth_v1_authorization_pb.HasPermissionRequest, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_authorization_pb.HasPermissionResponse) => void): grpc.ClientUnaryCall;
    public hasPermission(request: proto_auth_v1_authorization_pb.HasPermissionRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_authorization_pb.HasPermissionResponse) => void): grpc.ClientUnaryCall;
    public hasPermission(request: proto_auth_v1_authorization_pb.HasPermissionRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_authorization_pb.HasPermissionResponse) => void): grpc.ClientUnaryCall;
    public checkPermissions(request: proto_auth_v1_authorization_pb.CheckPermissionsRequest, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_authorization_pb.CheckPermissionsResponse) => void): grpc.ClientUnaryCall;
    public checkPermissions(request: proto_auth_v1_authorization_pb.CheckPermissionsRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_authorization_pb.CheckPermissionsResponse) => void): grpc.ClientUnaryCall;
    public checkPermissions(request: proto_auth_v1_authorization_pb.CheckPermissionsRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_authorization_pb.CheckPermissionsResponse) => void): grpc.ClientUnaryCall;
    public listPermissionPaths(request: proto_auth_v1_authorization_pb.ListPermissionPathsRequest, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_authorization_pb.ListPermissionPathsResponse) => void): grpc.ClientUnaryCall;
    public listPermissionPaths(request: proto_auth_v1_authorization_pb.ListPermissionPathsRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_authorization_pb.ListPermissionPathsResponse) => void): grpc.ClientUnaryCall;
    public listPermissionPaths(request: proto_auth_v1_authorization_pb.ListPermissionPathsRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_authorization_pb.ListPermissionPathsResponse) => void): grpc.ClientUnaryCall;
    public listPermissions(request: proto_auth_v1_authorization_pb.ListPermissionsRequest, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_authorization_pb.ListPermissionsResponse) => void): grpc.ClientUnaryCall;
    public listPermissions(request: proto_auth_v1_authorization_pb.ListPermissionsRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_authorization_pb.ListPermissionsResponse) => void): grpc.ClientUnaryCall;
    public listPermissions(request: proto_auth_v1_authorization_pb.ListPermissionsRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_authorization_pb.ListPermissionsResponse) => void): grpc.ClientUnaryCall;
    public migratePermissions(request: proto_auth_v1_authorization_pb.MigratePermissionsRequest, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_authorization_pb.MigratePermissionsResponse) => void): grpc.ClientUnaryCall;
    public migratePermissions(request: proto_auth_v1_authorization_pb.MigratePermissionsRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_authorization_pb.MigratePermissionsResponse) => void): grpc.ClientUnaryCall;
    public migratePermissions(request: proto_auth_v1_authorization_pb.MigratePermissionsRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_authorization_pb.MigratePermissionsResponse) => void): grpc.ClientUnaryCall;
    public assignRolesToUser(request: proto_auth_v1_authorization_pb.AssignRolesToUserRequest, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_authorization_pb.AssignRolesToUserResponse) => void): grpc.ClientUnaryCall;
    public assignRolesToUser(request: proto_auth_v1_authorization_pb.AssignRolesToUserRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_authorization_pb.AssignRolesToUserResponse) => void): grpc.ClientUnaryCall;
    public assignRolesToUser(request: proto_auth_v1_authorization_pb.AssignRolesToUserRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_authorization_pb.AssignRolesToUserResponse) => void): grpc.ClientUnaryCall;
    public getUserRoles(request: proto_auth_v1_authorization_pb.GetUserRolesRequest, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_authorization_pb.GetUserRolesResponse) => void): grpc.ClientUnaryCall;
    public getUserRoles(request: proto_auth_v1_authorization_pb.GetUserRolesRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_authorization_pb.GetUserRolesResponse) => void): grpc.ClientUnaryCall;
    public getUserRoles(request: proto_auth_v1_authorization_pb.GetUserRolesRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_authorization_pb.GetUserRolesResponse) => void): grpc.ClientUnaryCall;
    public createRole(request: proto_auth_v1_authorization_pb.CreateRoleRequest, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_authorization_pb.CreateRoleResponse) => void): grpc.ClientUnaryCall;
    public createRole(request: proto_auth_v1_authorization_pb.CreateRoleRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_authorization_pb.CreateRoleResponse) => void): grpc.ClientUnaryCall;
    public createRole(request: proto_auth_v1_authorization_pb.CreateRoleRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_authorization_pb.CreateRoleResponse) => void): grpc.ClientUnaryCall;
    public updateRole(request: proto_auth_v1_authorization_pb.UpdateRoleRequest, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_authorization_pb.UpdateRoleResponse) => void): grpc.ClientUnaryCall;
    public updateRole(request: proto_auth_v1_authorization_pb.UpdateRoleRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_authorization_pb.UpdateRoleResponse) => void): grpc.ClientUnaryCall;
    public updateRole(request: proto_auth_v1_authorization_pb.UpdateRoleRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_authorization_pb.UpdateRoleResponse) => void): grpc.ClientUnaryCall;
    public listRole(request: proto_auth_v1_authorization_pb.ListRoleRequest, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_authorization_pb.ListRoleResponse) => void): grpc.ClientUnaryCall;
    public listRole(request: proto_auth_v1_authorization_pb.ListRoleRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_authorization_pb.ListRoleResponse) => void): grpc.ClientUnaryCall;
    public listRole(request: proto_auth_v1_authorization_pb.ListRoleRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_authorization_pb.ListRoleResponse) => void): grpc.ClientUnaryCall;
}
