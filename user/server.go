package main

import (
	"context"
	"github.com/asim/go-micro/v3"
	pb "github.com/hjldev/newmicro-mall/user/proto"
	"log"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *pb.Request, rsp *pb.Response) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("helloword"),
	)
	service.Init()
	pb.RegisterGreeterHandler(service.Server(), new(Greeter))
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
