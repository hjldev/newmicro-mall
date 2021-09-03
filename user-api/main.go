package main

import (
	"context"
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/plugins/wrapper/select/roundrobin/v3"
	opentracing2 "github.com/asim/go-micro/plugins/wrapper/trace/opentracing/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/client"
	"github.com/asim/go-micro/v3/registry"
	"github.com/hjldev/newmicro-mall/common"
	"github.com/hjldev/newmicro-mall/user-api/handler"
	pb "github.com/hjldev/newmicro-mall/user/proto/user"
	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/common/log"
	"net"
	"net/http"
)

func main() {
	//注册中心
	consulRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	//链路追踪
	t, io, err := common.NewTracer("top.hjlinfo.mall.userApi", "localhost:6831")
	if err != nil {
		log.Error(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	//熔断器
	hystrixStreamHandler := hystrix.NewStreamHandler()
	hystrixStreamHandler.Start()
	//启动熔断器端口端口
	go func() {
		err = http.ListenAndServe(net.JoinHostPort("0.0.0.0", " "), hystrixStreamHandler)
		if err != nil {
			log.Error(err)
		}
	}()

	// New Service
	service := micro.NewService(
		micro.Name("top.hjlinfo.mall.userApi"),
		micro.Version("latest"),
		micro.Address("0.0.0.0:8085"),
		//添加 consul 注册中心
		micro.Registry(consulRegistry),
		//添加链路追踪
		micro.WrapClient(opentracing2.NewClientWrapper(opentracing.GlobalTracer())),
		//添加熔断
		micro.WrapClient(NewClientHystrixWrapper()),
		//添加负载均衡
		micro.WrapClient(roundrobin.NewClientWrapper()),
	)

	// Initialise service
	service.Init()

	common.PrometheusBoot(9095)

	userService := pb.NewUserService("top.hjlinfo.mall.user", service.Client())

	userService.Register(context.TODO(), &pb.UserInfoRequest{
		LoginName: "test-micro",
		Password:  "123456",
	})

	// Register Handler
	if err := pb.RegisterUserHandler(service.Server(), &handler.UserApi{UserService: userService}); err != nil {
		log.Error(err)
		return
	}

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

type clientWrapper struct {
	client.Client
}

func (c *clientWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	return hystrix.Do(req.Service()+"."+req.Endpoint(), func() error {
		//run 正常执行
		fmt.Println(req.Service() + "." + req.Endpoint())
		return c.Client.Call(ctx, req, rsp, opts...)
	}, func(err error) error {
		fmt.Println(err)
		return err
	})
}

func NewClientHystrixWrapper() client.Wrapper {
	return func(i client.Client) client.Client {
		return &clientWrapper{i}
	}
}
