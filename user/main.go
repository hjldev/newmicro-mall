package main

import (
	"fmt"
	"github.com/asim/go-micro/v3/server"
	"github.com/hjldev/newmicro-mall/user/domain/repository"
	"github.com/hjldev/newmicro-mall/user/domain/service"
	"github.com/hjldev/newmicro-mall/user/handler"
	userPb "github.com/hjldev/newmicro-mall/user/proto/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	server.Init(server.Name("user"))

	dsn := "root:123456@tcp(127.0.0.1:3306)/newmicro?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	rp := repository.NewUserRepository(db)
	userDataService := service.NewUserDataService(rp)

	err = userPb.RegisterUserHandler(server.NewServer(), &handler.User{UserDataService: userDataService})

	server.Handle(server.NewHandler(new(handler.User)))

	// Run server
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}

}
