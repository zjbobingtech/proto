GOPATH:=$(shell go env GOPATH)

.PHONY: proto
proto:
	protoc --proto_path=. --micro_out=.  --go_out=:. bbm/instance/instance.proto

	protoc --proto_path=. --micro_out=.  --go_out=:. bbm/resource/resource.proto

	protoc --proto_path=. --micro_out=.  --go_out=:. bbm/basic/basic.proto

	protoc --proto_path=. --micro_out=.  --go_out=:. bbm/bomber/bomber.proto

