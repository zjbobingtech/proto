GOPATH:=$(shell go env GOPATH)

.PHONY: proto
proto:
	protoc --proto_path=. --go-grpc_out=. --go_out=:. bbm-api/test/test.proto

