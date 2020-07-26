#!/usr/bin/env bash

set -e

if ! [[ "$0" =~ scripts/genproto.sh ]]; then
	echo "must be run from repository root"
	exit 255
fi

if [[ $(protoc --version | cut -f2 -d' ') != "3.7.1" ]]; then
	echo "could not find protoc 3.7.1, is it installed + in PATH?"
	exit 255
fi

#=$1
export GO111MODULE=off
# set up self-contained GOPATH for building
export GOPATH=${PWD}/gopath.proto
export GOBIN=${PWD}/bin
export PATH="${GOBIN}:${PATH}"

GOGOPROTO_ROOT="${GOPATH}/src/github.com/gogo/protobuf"
GOGOPROTO_PATH="${GOGOPROTO_ROOT}:${GOGOPROTO_ROOT}/protobuf"
GRPC_GATEWAY_ROOT="${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway"
function cleanup {
  # Remove the whole fake GOPATH which can really confuse go mod.
  rm -rf "${PWD}/gopath.proto"
}

cleanup
trap cleanup EXIT
echo $GOPATH
go get -u github.com/gogo/protobuf/{proto,protoc-gen-gogo,gogoproto}
go get -u golang.org/x/tools/cmd/goimports
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

protoc --gofast_out=plugins=grpc:. -I=".:${GOGOPROTO_PATH}:${GRPC_GATEWAY_ROOT}/third_party/googleapis" ${1}/*.proto