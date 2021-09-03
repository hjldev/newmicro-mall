.PHONY: proto
proto:
	protoc --micro_out=./proto/cartApi --go_out=./proto/cartApi ./proto/cartApi/cart_api.proto

.PHONY: build
build: 
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o cart-api *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t cart-api:latest

.PHONY: start-server
start:
	GOLANG_PROTOBUF_REGISTRATION_CONFLICT=warn go run main.go
