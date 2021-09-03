module github.com/hjldev/newmicro-mall/user

go 1.16

require (
	github.com/asim/go-micro/plugins/registry/consul/v3 v3.0.0-20210903064949-80dbe510777a
	github.com/asim/go-micro/plugins/wrapper/ratelimiter/uber/v3 v3.0.0-20210903064949-80dbe510777a
	github.com/asim/go-micro/plugins/wrapper/trace/opentracing/v3 v3.0.0-20210903064949-80dbe510777a
	github.com/asim/go-micro/v3 v3.5.2-0.20210630062103-c13bb07171bc
	github.com/golang/protobuf v1.5.2
	github.com/hjldev/newmicro-mall/common v0.0.0-20210903071242-9f7226ff386e
	github.com/opentracing/opentracing-go v1.2.0
	google.golang.org/protobuf v1.26.0
	gorm.io/driver/mysql v1.1.0
	gorm.io/gorm v1.21.9
)

replace github.com/hjldev/newmicro-mall/common => ../common
