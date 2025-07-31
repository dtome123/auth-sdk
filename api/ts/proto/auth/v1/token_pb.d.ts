import * as jspb from 'google-protobuf'

import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb'; // proto import: "google/protobuf/timestamp.proto"


export class SignRequest extends jspb.Message {
  getUserId(): string;
  setUserId(value: string): SignRequest;

  getUserType(): string;
  setUserType(value: string): SignRequest;

  getDeviceId(): string;
  setDeviceId(value: string): SignRequest;

  getMetadata(): string;
  setMetadata(value: string): SignRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SignRequest.AsObject;
  static toObject(includeInstance: boolean, msg: SignRequest): SignRequest.AsObject;
  static serializeBinaryToWriter(message: SignRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SignRequest;
  static deserializeBinaryFromReader(message: SignRequest, reader: jspb.BinaryReader): SignRequest;
}

export namespace SignRequest {
  export type AsObject = {
    userId: string,
    userType: string,
    deviceId: string,
    metadata: string,
  }
}

export class RefreshRequest extends jspb.Message {
  getRefreshToken(): string;
  setRefreshToken(value: string): RefreshRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RefreshRequest.AsObject;
  static toObject(includeInstance: boolean, msg: RefreshRequest): RefreshRequest.AsObject;
  static serializeBinaryToWriter(message: RefreshRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RefreshRequest;
  static deserializeBinaryFromReader(message: RefreshRequest, reader: jspb.BinaryReader): RefreshRequest;
}

export namespace RefreshRequest {
  export type AsObject = {
    refreshToken: string,
  }
}

export class TokenRequest extends jspb.Message {
  getGrantType(): string;
  setGrantType(value: string): TokenRequest;

  getClientAssertionType(): string;
  setClientAssertionType(value: string): TokenRequest;

  getClientAssertion(): string;
  setClientAssertion(value: string): TokenRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TokenRequest.AsObject;
  static toObject(includeInstance: boolean, msg: TokenRequest): TokenRequest.AsObject;
  static serializeBinaryToWriter(message: TokenRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TokenRequest;
  static deserializeBinaryFromReader(message: TokenRequest, reader: jspb.BinaryReader): TokenRequest;
}

export namespace TokenRequest {
  export type AsObject = {
    grantType: string,
    clientAssertionType: string,
    clientAssertion: string,
  }
}

export class SignResponse extends jspb.Message {
  getAccessToken(): string;
  setAccessToken(value: string): SignResponse;

  getRefreshToken(): string;
  setRefreshToken(value: string): SignResponse;

  getExpiresIn(): number;
  setExpiresIn(value: number): SignResponse;

  getExpiresAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setExpiresAt(value?: google_protobuf_timestamp_pb.Timestamp): SignResponse;
  hasExpiresAt(): boolean;
  clearExpiresAt(): SignResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SignResponse.AsObject;
  static toObject(includeInstance: boolean, msg: SignResponse): SignResponse.AsObject;
  static serializeBinaryToWriter(message: SignResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SignResponse;
  static deserializeBinaryFromReader(message: SignResponse, reader: jspb.BinaryReader): SignResponse;
}

export namespace SignResponse {
  export type AsObject = {
    accessToken: string,
    refreshToken: string,
    expiresIn: number,
    expiresAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }
}

export class RefreshResponse extends jspb.Message {
  getAccessToken(): string;
  setAccessToken(value: string): RefreshResponse;

  getRefreshToken(): string;
  setRefreshToken(value: string): RefreshResponse;

  getExpiresIn(): number;
  setExpiresIn(value: number): RefreshResponse;

  getExpiresAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setExpiresAt(value?: google_protobuf_timestamp_pb.Timestamp): RefreshResponse;
  hasExpiresAt(): boolean;
  clearExpiresAt(): RefreshResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RefreshResponse.AsObject;
  static toObject(includeInstance: boolean, msg: RefreshResponse): RefreshResponse.AsObject;
  static serializeBinaryToWriter(message: RefreshResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RefreshResponse;
  static deserializeBinaryFromReader(message: RefreshResponse, reader: jspb.BinaryReader): RefreshResponse;
}

export namespace RefreshResponse {
  export type AsObject = {
    accessToken: string,
    refreshToken: string,
    expiresIn: number,
    expiresAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }
}

export class TokenResponse extends jspb.Message {
  getAccessToken(): string;
  setAccessToken(value: string): TokenResponse;

  getTokenType(): string;
  setTokenType(value: string): TokenResponse;

  getExpiresIn(): number;
  setExpiresIn(value: number): TokenResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TokenResponse.AsObject;
  static toObject(includeInstance: boolean, msg: TokenResponse): TokenResponse.AsObject;
  static serializeBinaryToWriter(message: TokenResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TokenResponse;
  static deserializeBinaryFromReader(message: TokenResponse, reader: jspb.BinaryReader): TokenResponse;
}

export namespace TokenResponse {
  export type AsObject = {
    accessToken: string,
    tokenType: string,
    expiresIn: number,
  }
}

