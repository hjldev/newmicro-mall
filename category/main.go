package main

import (
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3"
	log "github.com/asim/go-micro/v3/logger"
	"github.com/asim/go-micro/v3/registry"
	"github.com/hjldev/newmicro-mall/category/common"
	"github.com/hjldev/newmicro-mall/category/domain/repository"
	service2 "github.com/hjldev/newmicro-mall/category/domain/service"
	"github.com/hjldev/newmicro-mall/category/handler"
	"github.com/hjldev/newmicro-mall/category/proto/category"
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
	consulRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	// New Service
	service := micro.NewService(
		micro.Name("category"),
		micro.Version("latest"),
		//这里设置地址和需要暴露的端口
		micro.Address("127.0.0.1:8082"),
		//添加consul 作为注册中心
		micro.Registry(consulRegistry),
	)

	//获取mysql配置,路径中不带前缀
	mysqlInfo := common.GetMysqlFromConsul(consulConfig, "mysql")

	dsn := mysqlInfo.User + ":" + mysqlInfo.Pwd + "@tcp(127.0.0.1:3306)/newmicro?charset=utf8mb4&parseTime=True&loc=Local"

	//连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error(err)
	}

	rp := repository.NewCategoryRepository(db)
	rp.InitTable()
	// Initialise service
	service.Init()

	categoryDataService := service2.NewCategoryDataService(rp)

	err = category.RegisterCategoryHandler(service.Server(), &handler.Category{CategoryDataService: categoryDataService})
	if err != nil {
		log.Error(err)
	}

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
