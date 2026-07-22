#!/usr/bin/env bash

DOMAIN_PATH="github.com/dtome123"
MY_REPO="$DOMAIN_PATH/auth-sdk"
GO_GEN="./api/go/"
TS_GEN="./api/ts/"

for dir in ./api/*; do rm -rf $dir/* ; done

for f in ./proto/*; do
  PROTO_NAME=$(basename $f)

  mkdir -p ${GO_GEN} && mkdir -p ${TS_GEN}

  ./node_modules/.bin/grpc_tools_node_protoc \
  --descriptor_set_out=bin/descriptor_set.bin \
  --include_imports \
  proto/$PROTO_NAME/**/*.proto

  protoc \
	--plugin="protoc-gen-ts=./node_modules/.bin/protoc-gen-ts" \
	--plugin="protoc-gen-grpc=./node_modules/.bin/grpc_tools_node_protoc_plugin" \
	--grpc_out="grpc_js:${TS_GEN}" \
	--ts_out="service=grpc-node,mode=grpc-js:${TS_GEN}" \
	--descriptor_set_out=bin/descriptor_set.bin \
  --include_imports \
	proto/$PROTO_NAME/**/*.proto

  protoc --go_out ${GO_GEN} --go-grpc_out ${GO_GEN} --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative  \
	--js_out=import_style=commonjs,binary:${TS_GEN} \
	--grpc-web_out=import_style=typescript,mode=grpcweb:${TS_GEN} \
	--plugin="protoc-gen-ts=./node_modules/.bin/protoc-gen-ts" \
	proto/$PROTO_NAME/**/*.proto

  mv ${GO_GEN}/proto/* ${GO_GEN}
  rm -rf ${GO_GEN}/proto

done
