// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var proto_auth_v1_authorization_pb = require('../../../proto/auth/v1/authorization_pb.js');

function serialize_auth_v1_AssignRolesToUserRequest(arg) {
  if (!(arg instanceof proto_auth_v1_authorization_pb.AssignRolesToUserRequest)) {
    throw new Error('Expected argument of type auth.v1.AssignRolesToUserRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_auth_v1_AssignRolesToUserRequest(buffer_arg) {
  return proto_auth_v1_authorization_pb.AssignRolesToUserRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_auth_v1_AssignRolesToUserResponse(arg) {
  if (!(arg instanceof proto_auth_v1_authorization_pb.AssignRolesToUserResponse)) {
    throw new Error('Expected argument of type auth.v1.AssignRolesToUserResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_auth_v1_AssignRolesToUserResponse(buffer_arg) {
  return proto_auth_v1_authorization_pb.AssignRolesToUserResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_auth_v1_CheckPermissionsRequest(arg) {
  if (!(arg instanceof proto_auth_v1_authorization_pb.CheckPermissionsRequest)) {
    throw new Error('Expected argument of type auth.v1.CheckPermissionsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_auth_v1_CheckPermissionsRequest(buffer_arg) {
  return proto_auth_v1_authorization_pb.CheckPermissionsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_auth_v1_CheckPermissionsResponse(arg) {
  if (!(arg instanceof proto_auth_v1_authorization_pb.CheckPermissionsResponse)) {
    throw new Error('Expected argument of type auth.v1.CheckPermissionsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_auth_v1_CheckPermissionsResponse(buffer_arg) {
  return proto_auth_v1_authorization_pb.CheckPermissionsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_auth_v1_CreateRoleRequest(arg) {
  if (!(arg instanceof proto_auth_v1_authorization_pb.CreateRoleRequest)) {
    throw new Error('Expected argument of type auth.v1.CreateRoleRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_auth_v1_CreateRoleRequest(buffer_arg) {
  return proto_auth_v1_authorization_pb.CreateRoleRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_auth_v1_CreateRoleResponse(arg) {
  if (!(arg instanceof proto_auth_v1_authorization_pb.CreateRoleResponse)) {
    throw new Error('Expected argument of type auth.v1.CreateRoleResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_auth_v1_CreateRoleResponse(buffer_arg) {
  return proto_auth_v1_authorization_pb.CreateRoleResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_auth_v1_GetUserPermissionsRequest(arg) {
  if (!(arg instanceof proto_auth_v1_authorization_pb.GetUserPermissionsRequest)) {
    throw new Error('Expected argument of type auth.v1.GetUserPermissionsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_auth_v1_GetUserPermissionsRequest(buffer_arg) {
  return proto_auth_v1_authorization_pb.GetUserPermissionsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_auth_v1_GetUserPermissionsResponse(arg) {
  if (!(arg instanceof proto_auth_v1_authorization_pb.GetUserPermissionsResponse)) {
    throw new Error('Expected argument of type auth.v1.GetUserPermissionsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_auth_v1_GetUserPermissionsResponse(buffer_arg) {
  return proto_auth_v1_authorization_pb.GetUserPermissionsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_auth_v1_GetUserRolesRequest(arg) {
  if (!(arg instanceof proto_auth_v1_authorization_pb.GetUserRolesRequest)) {
    throw new Error('Expected argument of type auth.v1.GetUserRolesRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_auth_v1_GetUserRolesRequest(buffer_arg) {
  return proto_auth_v1_authorization_pb.GetUserRolesRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_auth_v1_GetUserRolesResponse(arg) {
  if (!(arg instanceof proto_auth_v1_authorization_pb.GetUserRolesResponse)) {
    throw new Error('Expected argument of type auth.v1.GetUserRolesResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_auth_v1_GetUserRolesResponse(buffer_arg) {
  return proto_auth_v1_authorization_pb.GetUserRolesResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_auth_v1_HasPermissionRequest(arg) {
  if (!(arg instanceof proto_auth_v1_authorization_pb.HasPermissionRequest)) {
    throw new Error('Expected argument of type auth.v1.HasPermissionRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_auth_v1_HasPermissionRequest(buffer_arg) {
  return proto_auth_v1_authorization_pb.HasPermissionRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_auth_v1_HasPermissionResponse(arg) {
  if (!(arg instanceof proto_auth_v1_authorization_pb.HasPermissionResponse)) {
    throw new Error('Expected argument of type auth.v1.HasPermissionResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_auth_v1_HasPermissionResponse(buffer_arg) {
  return proto_auth_v1_authorization_pb.HasPermissionResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_auth_v1_ListPermissionPathsRequest(arg) {
  if (!(arg instanceof proto_auth_v1_authorization_pb.ListPermissionPathsRequest)) {
    throw new Error('Expected argument of type auth.v1.ListPermissionPathsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_auth_v1_ListPermissionPathsRequest(buffer_arg) {
  return proto_auth_v1_authorization_pb.ListPermissionPathsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_auth_v1_ListPermissionPathsResponse(arg) {
  if (!(arg instanceof proto_auth_v1_authorization_pb.ListPermissionPathsResponse)) {
    throw new Error('Expected argument of type auth.v1.ListPermissionPathsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_auth_v1_ListPermissionPathsResponse(buffer_arg) {
  return proto_auth_v1_authorization_pb.ListPermissionPathsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_auth_v1_ListPermissionsRequest(arg) {
  if (!(arg instanceof proto_auth_v1_authorization_pb.ListPermissionsRequest)) {
    throw new Error('Expected argument of type auth.v1.ListPermissionsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_auth_v1_ListPermissionsRequest(buffer_arg) {
  return proto_auth_v1_authorization_pb.ListPermissionsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_auth_v1_ListPermissionsResponse(arg) {
  if (!(arg instanceof proto_auth_v1_authorization_pb.ListPermissionsResponse)) {
    throw new Error('Expected argument of type auth.v1.ListPermissionsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_auth_v1_ListPermissionsResponse(buffer_arg) {
  return proto_auth_v1_authorization_pb.ListPermissionsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_auth_v1_ListRoleRequest(arg) {
  if (!(arg instanceof proto_auth_v1_authorization_pb.ListRoleRequest)) {
    throw new Error('Expected argument of type auth.v1.ListRoleRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_auth_v1_ListRoleRequest(buffer_arg) {
  return proto_auth_v1_authorization_pb.ListRoleRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_auth_v1_ListRoleResponse(arg) {
  if (!(arg instanceof proto_auth_v1_authorization_pb.ListRoleResponse)) {
    throw new Error('Expected argument of type auth.v1.ListRoleResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_auth_v1_ListRoleResponse(buffer_arg) {
  return proto_auth_v1_authorization_pb.ListRoleResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_auth_v1_MigratePermissionsRequest(arg) {
  if (!(arg instanceof proto_auth_v1_authorization_pb.MigratePermissionsRequest)) {
    throw new Error('Expected argument of type auth.v1.MigratePermissionsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_auth_v1_MigratePermissionsRequest(buffer_arg) {
  return proto_auth_v1_authorization_pb.MigratePermissionsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_auth_v1_MigratePermissionsResponse(arg) {
  if (!(arg instanceof proto_auth_v1_authorization_pb.MigratePermissionsResponse)) {
    throw new Error('Expected argument of type auth.v1.MigratePermissionsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_auth_v1_MigratePermissionsResponse(buffer_arg) {
  return proto_auth_v1_authorization_pb.MigratePermissionsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_auth_v1_UpdateRoleRequest(arg) {
  if (!(arg instanceof proto_auth_v1_authorization_pb.UpdateRoleRequest)) {
    throw new Error('Expected argument of type auth.v1.UpdateRoleRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_auth_v1_UpdateRoleRequest(buffer_arg) {
  return proto_auth_v1_authorization_pb.UpdateRoleRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_auth_v1_UpdateRoleResponse(arg) {
  if (!(arg instanceof proto_auth_v1_authorization_pb.UpdateRoleResponse)) {
    throw new Error('Expected argument of type auth.v1.UpdateRoleResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_auth_v1_UpdateRoleResponse(buffer_arg) {
  return proto_auth_v1_authorization_pb.UpdateRoleResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var AuthorizationServiceService = exports.AuthorizationServiceService = {
  // ========= Permission APIs ==========
getUserPermissions: {
    path: '/auth.v1.AuthorizationService/GetUserPermissions',
    requestStream: false,
    responseStream: false,
    requestType: proto_auth_v1_authorization_pb.GetUserPermissionsRequest,
    responseType: proto_auth_v1_authorization_pb.GetUserPermissionsResponse,
    requestSerialize: serialize_auth_v1_GetUserPermissionsRequest,
    requestDeserialize: deserialize_auth_v1_GetUserPermissionsRequest,
    responseSerialize: serialize_auth_v1_GetUserPermissionsResponse,
    responseDeserialize: deserialize_auth_v1_GetUserPermissionsResponse,
  },
  hasPermission: {
    path: '/auth.v1.AuthorizationService/HasPermission',
    requestStream: false,
    responseStream: false,
    requestType: proto_auth_v1_authorization_pb.HasPermissionRequest,
    responseType: proto_auth_v1_authorization_pb.HasPermissionResponse,
    requestSerialize: serialize_auth_v1_HasPermissionRequest,
    requestDeserialize: deserialize_auth_v1_HasPermissionRequest,
    responseSerialize: serialize_auth_v1_HasPermissionResponse,
    responseDeserialize: deserialize_auth_v1_HasPermissionResponse,
  },
  checkPermissions: {
    path: '/auth.v1.AuthorizationService/CheckPermissions',
    requestStream: false,
    responseStream: false,
    requestType: proto_auth_v1_authorization_pb.CheckPermissionsRequest,
    responseType: proto_auth_v1_authorization_pb.CheckPermissionsResponse,
    requestSerialize: serialize_auth_v1_CheckPermissionsRequest,
    requestDeserialize: deserialize_auth_v1_CheckPermissionsRequest,
    responseSerialize: serialize_auth_v1_CheckPermissionsResponse,
    responseDeserialize: deserialize_auth_v1_CheckPermissionsResponse,
  },
  listPermissionPaths: {
    path: '/auth.v1.AuthorizationService/ListPermissionPaths',
    requestStream: false,
    responseStream: false,
    requestType: proto_auth_v1_authorization_pb.ListPermissionPathsRequest,
    responseType: proto_auth_v1_authorization_pb.ListPermissionPathsResponse,
    requestSerialize: serialize_auth_v1_ListPermissionPathsRequest,
    requestDeserialize: deserialize_auth_v1_ListPermissionPathsRequest,
    responseSerialize: serialize_auth_v1_ListPermissionPathsResponse,
    responseDeserialize: deserialize_auth_v1_ListPermissionPathsResponse,
  },
  listPermissions: {
    path: '/auth.v1.AuthorizationService/ListPermissions',
    requestStream: false,
    responseStream: false,
    requestType: proto_auth_v1_authorization_pb.ListPermissionsRequest,
    responseType: proto_auth_v1_authorization_pb.ListPermissionsResponse,
    requestSerialize: serialize_auth_v1_ListPermissionsRequest,
    requestDeserialize: deserialize_auth_v1_ListPermissionsRequest,
    responseSerialize: serialize_auth_v1_ListPermissionsResponse,
    responseDeserialize: deserialize_auth_v1_ListPermissionsResponse,
  },
  migratePermissions: {
    path: '/auth.v1.AuthorizationService/MigratePermissions',
    requestStream: false,
    responseStream: false,
    requestType: proto_auth_v1_authorization_pb.MigratePermissionsRequest,
    responseType: proto_auth_v1_authorization_pb.MigratePermissionsResponse,
    requestSerialize: serialize_auth_v1_MigratePermissionsRequest,
    requestDeserialize: deserialize_auth_v1_MigratePermissionsRequest,
    responseSerialize: serialize_auth_v1_MigratePermissionsResponse,
    responseDeserialize: deserialize_auth_v1_MigratePermissionsResponse,
  },
  // ========= Role APIs ===============
assignRolesToUser: {
    path: '/auth.v1.AuthorizationService/AssignRolesToUser',
    requestStream: false,
    responseStream: false,
    requestType: proto_auth_v1_authorization_pb.AssignRolesToUserRequest,
    responseType: proto_auth_v1_authorization_pb.AssignRolesToUserResponse,
    requestSerialize: serialize_auth_v1_AssignRolesToUserRequest,
    requestDeserialize: deserialize_auth_v1_AssignRolesToUserRequest,
    responseSerialize: serialize_auth_v1_AssignRolesToUserResponse,
    responseDeserialize: deserialize_auth_v1_AssignRolesToUserResponse,
  },
  getUserRoles: {
    path: '/auth.v1.AuthorizationService/GetUserRoles',
    requestStream: false,
    responseStream: false,
    requestType: proto_auth_v1_authorization_pb.GetUserRolesRequest,
    responseType: proto_auth_v1_authorization_pb.GetUserRolesResponse,
    requestSerialize: serialize_auth_v1_GetUserRolesRequest,
    requestDeserialize: deserialize_auth_v1_GetUserRolesRequest,
    responseSerialize: serialize_auth_v1_GetUserRolesResponse,
    responseDeserialize: deserialize_auth_v1_GetUserRolesResponse,
  },
  createRole: {
    path: '/auth.v1.AuthorizationService/CreateRole',
    requestStream: false,
    responseStream: false,
    requestType: proto_auth_v1_authorization_pb.CreateRoleRequest,
    responseType: proto_auth_v1_authorization_pb.CreateRoleResponse,
    requestSerialize: serialize_auth_v1_CreateRoleRequest,
    requestDeserialize: deserialize_auth_v1_CreateRoleRequest,
    responseSerialize: serialize_auth_v1_CreateRoleResponse,
    responseDeserialize: deserialize_auth_v1_CreateRoleResponse,
  },
  updateRole: {
    path: '/auth.v1.AuthorizationService/UpdateRole',
    requestStream: false,
    responseStream: false,
    requestType: proto_auth_v1_authorization_pb.UpdateRoleRequest,
    responseType: proto_auth_v1_authorization_pb.UpdateRoleResponse,
    requestSerialize: serialize_auth_v1_UpdateRoleRequest,
    requestDeserialize: deserialize_auth_v1_UpdateRoleRequest,
    responseSerialize: serialize_auth_v1_UpdateRoleResponse,
    responseDeserialize: deserialize_auth_v1_UpdateRoleResponse,
  },
  listRole: {
    path: '/auth.v1.AuthorizationService/ListRole',
    requestStream: false,
    responseStream: false,
    requestType: proto_auth_v1_authorization_pb.ListRoleRequest,
    responseType: proto_auth_v1_authorization_pb.ListRoleResponse,
    requestSerialize: serialize_auth_v1_ListRoleRequest,
    requestDeserialize: deserialize_auth_v1_ListRoleRequest,
    responseSerialize: serialize_auth_v1_ListRoleResponse,
    responseDeserialize: deserialize_auth_v1_ListRoleResponse,
  },
};

exports.AuthorizationServiceClient = grpc.makeGenericClientConstructor(AuthorizationServiceService, 'AuthorizationService');
