.PHONY: proto
proto:
	protoc --micro_out=./proto/order --go_out=./proto/order ./proto/order/order.proto

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o order-service *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t order-service:latest
