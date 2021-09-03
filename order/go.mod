module github.com/hjldev/newmicro-mall/order

go 1.15

require (
	github.com/asim/go-micro/plugins/registry/consul/v3 v3.0.0-20210902172428-e0807917878f
	github.com/asim/go-micro/plugins/wrapper/monitoring/prometheus/v3 v3.0.0-20210902172428-e0807917878f
	github.com/asim/go-micro/plugins/wrapper/ratelimiter/uber/v3 v3.0.0-20210902172428-e0807917878f
	github.com/asim/go-micro/plugins/wrapper/trace/opentracing/v3 v3.0.0-20210902172428-e0807917878f
	github.com/asim/go-micro/v3 v3.5.2-0.20210630062103-c13bb07171bc
	github.com/golang/protobuf v1.5.2
	github.com/hjldev/newmicro-mall/common v0.0.0-20201210084148-1ab1eced812e
	github.com/jinzhu/gorm v1.9.16
	github.com/opentracing/opentracing-go v1.2.0
	google.golang.org/protobuf v1.26.0
)

replace github.com/hjldev/newmicro-mall/common => ../common
