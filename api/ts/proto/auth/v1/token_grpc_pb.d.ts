// package: auth.v1
// file: proto/auth/v1/token.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "grpc";
import * as proto_auth_v1_token_pb from "../../../proto/auth/v1/token_pb";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";

interface ITokenServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    token: ITokenServiceService_IToken;
    sign: ITokenServiceService_ISign;
    refresh: ITokenServiceService_IRefresh;
}

interface ITokenServiceService_IToken extends grpc.MethodDefinition<proto_auth_v1_token_pb.TokenRequest, proto_auth_v1_token_pb.TokenResponse> {
    path: "/auth.v1.TokenService/Token";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_auth_v1_token_pb.TokenRequest>;
    requestDeserialize: grpc.deserialize<proto_auth_v1_token_pb.TokenRequest>;
    responseSerialize: grpc.serialize<proto_auth_v1_token_pb.TokenResponse>;
    responseDeserialize: grpc.deserialize<proto_auth_v1_token_pb.TokenResponse>;
}
interface ITokenServiceService_ISign extends grpc.MethodDefinition<proto_auth_v1_token_pb.SignRequest, proto_auth_v1_token_pb.SignResponse> {
    path: "/auth.v1.TokenService/Sign";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_auth_v1_token_pb.SignRequest>;
    requestDeserialize: grpc.deserialize<proto_auth_v1_token_pb.SignRequest>;
    responseSerialize: grpc.serialize<proto_auth_v1_token_pb.SignResponse>;
    responseDeserialize: grpc.deserialize<proto_auth_v1_token_pb.SignResponse>;
}
interface ITokenServiceService_IRefresh extends grpc.MethodDefinition<proto_auth_v1_token_pb.RefreshRequest, proto_auth_v1_token_pb.RefreshResponse> {
    path: "/auth.v1.TokenService/Refresh";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<proto_auth_v1_token_pb.RefreshRequest>;
    requestDeserialize: grpc.deserialize<proto_auth_v1_token_pb.RefreshRequest>;
    responseSerialize: grpc.serialize<proto_auth_v1_token_pb.RefreshResponse>;
    responseDeserialize: grpc.deserialize<proto_auth_v1_token_pb.RefreshResponse>;
}

export const TokenServiceService: ITokenServiceService;

export interface ITokenServiceServer {
    token: grpc.handleUnaryCall<proto_auth_v1_token_pb.TokenRequest, proto_auth_v1_token_pb.TokenResponse>;
    sign: grpc.handleUnaryCall<proto_auth_v1_token_pb.SignRequest, proto_auth_v1_token_pb.SignResponse>;
    refresh: grpc.handleUnaryCall<proto_auth_v1_token_pb.RefreshRequest, proto_auth_v1_token_pb.RefreshResponse>;
}

export interface ITokenServiceClient {
    token(request: proto_auth_v1_token_pb.TokenRequest, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_token_pb.TokenResponse) => void): grpc.ClientUnaryCall;
    token(request: proto_auth_v1_token_pb.TokenRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_token_pb.TokenResponse) => void): grpc.ClientUnaryCall;
    token(request: proto_auth_v1_token_pb.TokenRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_token_pb.TokenResponse) => void): grpc.ClientUnaryCall;
    sign(request: proto_auth_v1_token_pb.SignRequest, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_token_pb.SignResponse) => void): grpc.ClientUnaryCall;
    sign(request: proto_auth_v1_token_pb.SignRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_token_pb.SignResponse) => void): grpc.ClientUnaryCall;
    sign(request: proto_auth_v1_token_pb.SignRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_token_pb.SignResponse) => void): grpc.ClientUnaryCall;
    refresh(request: proto_auth_v1_token_pb.RefreshRequest, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_token_pb.RefreshResponse) => void): grpc.ClientUnaryCall;
    refresh(request: proto_auth_v1_token_pb.RefreshRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_token_pb.RefreshResponse) => void): grpc.ClientUnaryCall;
    refresh(request: proto_auth_v1_token_pb.RefreshRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_token_pb.RefreshResponse) => void): grpc.ClientUnaryCall;
}

export class TokenServiceClient extends grpc.Client implements ITokenServiceClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
    public token(request: proto_auth_v1_token_pb.TokenRequest, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_token_pb.TokenResponse) => void): grpc.ClientUnaryCall;
    public token(request: proto_auth_v1_token_pb.TokenRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_token_pb.TokenResponse) => void): grpc.ClientUnaryCall;
    public token(request: proto_auth_v1_token_pb.TokenRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_token_pb.TokenResponse) => void): grpc.ClientUnaryCall;
    public sign(request: proto_auth_v1_token_pb.SignRequest, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_token_pb.SignResponse) => void): grpc.ClientUnaryCall;
    public sign(request: proto_auth_v1_token_pb.SignRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_token_pb.SignResponse) => void): grpc.ClientUnaryCall;
    public sign(request: proto_auth_v1_token_pb.SignRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_token_pb.SignResponse) => void): grpc.ClientUnaryCall;
    public refresh(request: proto_auth_v1_token_pb.RefreshRequest, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_token_pb.RefreshResponse) => void): grpc.ClientUnaryCall;
    public refresh(request: proto_auth_v1_token_pb.RefreshRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_token_pb.RefreshResponse) => void): grpc.ClientUnaryCall;
    public refresh(request: proto_auth_v1_token_pb.RefreshRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: proto_auth_v1_token_pb.RefreshResponse) => void): grpc.ClientUnaryCall;
}
