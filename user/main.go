package main

import (
	"fmt"
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	ratelimit "github.com/asim/go-micro/plugins/wrapper/ratelimiter/uber/v3"
	opentracing2 "github.com/asim/go-micro/plugins/wrapper/trace/opentracing/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	"github.com/hjldev/newmicro-mall/common"
	"github.com/hjldev/newmicro-mall/user/domain/repository"
	service2 "github.com/hjldev/newmicro-mall/user/domain/service"
	"github.com/hjldev/newmicro-mall/user/handler"
	userPb "github.com/hjldev/newmicro-mall/user/proto/user"
	"github.com/opentracing/opentracing-go"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var QPS = 100

func main() {
	consulConfig, err := common.GetConsulConfig("127.0.0.1", 8500, "/micro/config")
	if err != nil {
		log.Println(err)
		return
	}
	mysqlInfo := common.GetMysqlFromConsul(consulConfig, "mysql")
	dsn := mysqlInfo.User + ":" + mysqlInfo.Pwd + "@tcp(127.0.0.1:3306)/newmicro?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	//链路追踪
	t, io, err := common.NewTracer("top.hjlinfo.mall.user", "localhost:6831")
	if err != nil {
		log.Println(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	//注册中心
	consulRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	// New Service
	service := micro.NewService(
		micro.Name("top.hjlinfo.mall.user"),
		micro.Version("latest"),
		//暴露的服务地址
		micro.Address("0.0.0.0:8084"),
		//注册中心
		micro.Registry(consulRegistry),
		//链路追踪
		micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
		//添加限流
		micro.WrapHandler(ratelimit.NewHandlerWrapper(QPS)),
	)

	// Initialise service
	service.Init()

	rp := repository.NewUserRepository(db)
	userDataService := service2.NewUserDataService(rp)

	err = userPb.RegisterUserHandler(service.Server(), &handler.User{UserDataService: userDataService})

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
