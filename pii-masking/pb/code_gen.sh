#!/bin/bash

set -eu

protoc --proto_path=. --go_out=paths=source_relative:. test.proto
protoc-go-inject-tag -input="*.pb.go"