GOPATH:=$(shell go env GOPATH)

.PHONY: proto
proto:
	protoc --proto_path=. --micro_out=.  --go_out=:. bbm-api/test/test.proto

