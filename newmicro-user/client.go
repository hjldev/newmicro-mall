package main

import (
	"context"
	"fmt"
	"github.com/asim/go-micro/v3"
	pb "github.com/hjldev/newmicro-mall/newmicro-user/proto"
)

func main() {

	service := micro.NewService(
		micro.Name("helloword"),
	)
	service.Init()

	greeterService := pb.NewGreeterService("helloword", service.Client())
	res, err := greeterService.Hello(context.TODO(), &pb.Request{
		Name: "micro",
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}
