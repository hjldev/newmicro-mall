
GOPATH:=$(shell go env GOPATH)
.PHONY: proto
proto:
	protoc --micro_out=./proto --go_out=:./proto proto/user_api.proto
	
.PHONY: build
build:
	go build -o user-api *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t user-api:latest
