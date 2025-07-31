// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var proto_auth_v1_token_pb = require('../../../proto/auth/v1/token_pb.js');
var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js');

function serialize_auth_v1_RefreshRequest(arg) {
  if (!(arg instanceof proto_auth_v1_token_pb.RefreshRequest)) {
    throw new Error('Expected argument of type auth.v1.RefreshRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_auth_v1_RefreshRequest(buffer_arg) {
  return proto_auth_v1_token_pb.RefreshRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_auth_v1_RefreshResponse(arg) {
  if (!(arg instanceof proto_auth_v1_token_pb.RefreshResponse)) {
    throw new Error('Expected argument of type auth.v1.RefreshResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_auth_v1_RefreshResponse(buffer_arg) {
  return proto_auth_v1_token_pb.RefreshResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_auth_v1_SignRequest(arg) {
  if (!(arg instanceof proto_auth_v1_token_pb.SignRequest)) {
    throw new Error('Expected argument of type auth.v1.SignRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_auth_v1_SignRequest(buffer_arg) {
  return proto_auth_v1_token_pb.SignRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_auth_v1_SignResponse(arg) {
  if (!(arg instanceof proto_auth_v1_token_pb.SignResponse)) {
    throw new Error('Expected argument of type auth.v1.SignResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_auth_v1_SignResponse(buffer_arg) {
  return proto_auth_v1_token_pb.SignResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_auth_v1_TokenRequest(arg) {
  if (!(arg instanceof proto_auth_v1_token_pb.TokenRequest)) {
    throw new Error('Expected argument of type auth.v1.TokenRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_auth_v1_TokenRequest(buffer_arg) {
  return proto_auth_v1_token_pb.TokenRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_auth_v1_TokenResponse(arg) {
  if (!(arg instanceof proto_auth_v1_token_pb.TokenResponse)) {
    throw new Error('Expected argument of type auth.v1.TokenResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_auth_v1_TokenResponse(buffer_arg) {
  return proto_auth_v1_token_pb.TokenResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var TokenServiceService = exports.TokenServiceService = {
  token: {
    path: '/auth.v1.TokenService/Token',
    requestStream: false,
    responseStream: false,
    requestType: proto_auth_v1_token_pb.TokenRequest,
    responseType: proto_auth_v1_token_pb.TokenResponse,
    requestSerialize: serialize_auth_v1_TokenRequest,
    requestDeserialize: deserialize_auth_v1_TokenRequest,
    responseSerialize: serialize_auth_v1_TokenResponse,
    responseDeserialize: deserialize_auth_v1_TokenResponse,
  },
  sign: {
    path: '/auth.v1.TokenService/Sign',
    requestStream: false,
    responseStream: false,
    requestType: proto_auth_v1_token_pb.SignRequest,
    responseType: proto_auth_v1_token_pb.SignResponse,
    requestSerialize: serialize_auth_v1_SignRequest,
    requestDeserialize: deserialize_auth_v1_SignRequest,
    responseSerialize: serialize_auth_v1_SignResponse,
    responseDeserialize: deserialize_auth_v1_SignResponse,
  },
  refresh: {
    path: '/auth.v1.TokenService/Refresh',
    requestStream: false,
    responseStream: false,
    requestType: proto_auth_v1_token_pb.RefreshRequest,
    responseType: proto_auth_v1_token_pb.RefreshResponse,
    requestSerialize: serialize_auth_v1_RefreshRequest,
    requestDeserialize: deserialize_auth_v1_RefreshRequest,
    responseSerialize: serialize_auth_v1_RefreshResponse,
    responseDeserialize: deserialize_auth_v1_RefreshResponse,
  },
};

exports.TokenServiceClient = grpc.makeGenericClientConstructor(TokenServiceService, 'TokenService');
