.PHONY: proto
proto:
	protoc --micro_out=./proto/user --go_out=./proto/user ./proto/user/user.proto

.PHONY: build
build: 

	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o user-service *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t user-service:latest

.PHONY: start-server
start:
	go run main.go
