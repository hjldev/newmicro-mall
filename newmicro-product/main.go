package main

import (
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	opentracing2 "github.com/asim/go-micro/plugins/wrapper/trace/opentracing/v3"
	"github.com/asim/go-micro/v3"
	log "github.com/asim/go-micro/v3/logger"
	"github.com/asim/go-micro/v3/registry"
	"github.com/hjldev/newmicro-mall/newmicro-product/common"
	"github.com/hjldev/newmicro-mall/newmicro-product/domain/repository"
	service2 "github.com/hjldev/newmicro-mall/newmicro-product/domain/service"
	"github.com/hjldev/newmicro-mall/newmicro-product/handler"
	"github.com/hjldev/newmicro-mall/newmicro-product/proto/product"
	"github.com/opentracing/opentracing-go"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	//配置中心
	consulConfig, err := common.GetConsulConfig("127.0.0.1", 8500, "/micro/config")
	if err != nil {
		log.Error(err)
	}
	//注册中心
	consulRegister := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	//链路追踪
	t, io, err := common.NewTracer("product", "127.0.0.1:6831")
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	//数据库设置
	mysqlInfo := common.GetMysqlFromConsul(consulConfig, "mysql")
	dsn := mysqlInfo.User + ":" + mysqlInfo.Pwd + "@tcp(127.0.0.1:3306)/newmicro?charset=utf8mb4&parseTime=True&loc=Local"
	//连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error(err)
	}

	//初始化
	repository.NewProductRepository(db).InitTable()

	productDataService := service2.NewProductDataService(repository.NewProductRepository(db))

	// 设置服务
	service := micro.NewService(
		micro.Name("product"),
		micro.Version("latest"),
		micro.Address("127.0.0.1:8085"),
		//添加注册中心
		micro.Registry(consulRegister),
		//绑定链路追踪
		micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
	)

	// Initialise service
	service.Init()

	// Register Handler
	product.RegisterProductHandler(service.Server(), &handler.Product{ProductDataService: productDataService})

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
