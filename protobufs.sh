#!/bin/bash

go get -u github.com/golang/protobuf/protoc-gen-go
pushd vendor/github.com/Blizzard/s2client-proto/s2clientprotocol/
protoc --go_out=../../../../../proto/ *.proto
popd