module github.com/hjldev/newmicro-mall/cart

go 1.16

require (
	github.com/HdrHistogram/hdrhistogram-go v1.1.0 // indirect
	github.com/asim/go-micro/plugins/config/source/consul/v3 v3.0.0-20210517071652-f48911d2c3ef
	github.com/asim/go-micro/plugins/registry/consul/v3 v3.0.0-20210517071652-f48911d2c3ef
	github.com/asim/go-micro/plugins/wrapper/ratelimiter/uber/v3 v3.0.0-20210517071652-f48911d2c3ef
	github.com/asim/go-micro/plugins/wrapper/trace/opentracing/v3 v3.0.0-20210517071652-f48911d2c3ef
	github.com/asim/go-micro/v3 v3.5.1
	github.com/golang/protobuf v1.5.2
	github.com/opentracing/opentracing-go v1.2.0
	github.com/uber/jaeger-client-go v2.28.0+incompatible
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
	google.golang.org/protobuf v1.26.0
	gorm.io/driver/mysql v1.1.0
	gorm.io/gorm v1.21.10
)
