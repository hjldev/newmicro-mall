.PHONY: proto
proto:
	protoc --micro_out=./proto/cart --go_out=./proto/cart ./proto/cart/cart.proto

.PHONY: build
build: 

	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o cart-service *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t cart-service:latest

.PHONY: start-server
start:
	go run main.go
