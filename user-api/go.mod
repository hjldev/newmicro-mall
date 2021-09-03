module github.com/hjldev/newmicro-mall/user-api

go 1.15

require (
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/asim/go-micro/plugins/registry/consul/v3 v3.0.0-20210903064949-80dbe510777a
	github.com/asim/go-micro/plugins/wrapper/select/roundrobin/v3 v3.0.0-20210903064949-80dbe510777a
	github.com/asim/go-micro/plugins/wrapper/trace/opentracing/v3 v3.0.0-20210903064949-80dbe510777a
	github.com/asim/go-micro/v3 v3.6.0
	github.com/golang/protobuf v1.5.2
	github.com/hjldev/newmicro-mall/common v0.0.0-20210903071242-9f7226ff386e
	github.com/hjldev/newmicro-mall/user v0.0.0-20210903071242-9f7226ff386e
	github.com/opentracing/opentracing-go v1.2.0
	github.com/prometheus/common v0.24.0
	google.golang.org/protobuf v1.26.0
)

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace (
	github.com/hjldev/newmicro-mall/user => ../user
	github.com/hjldev/newmicro-mall/common => ../common
)
