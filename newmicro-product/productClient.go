package main

import (
	"context"
	"fmt"
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	opentracing2 "github.com/asim/go-micro/plugins/wrapper/trace/opentracing/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	"github.com/hjldev/newmicro-mall/newmicro-product/common"
	"github.com/hjldev/newmicro-mall/newmicro-product/proto/product"
	"github.com/opentracing/opentracing-go"
	"log"
)

func main() {
	//注册中心
	consulRegisty := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})
	//链路追踪
	t, io, err := common.NewTracer("client", "localhost:6831")
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	service := micro.NewService(
		micro.Name("client"),
		micro.Version("latest"),
		micro.Address("127.0.0.1:8085"),
		//添加注册中心
		micro.Registry(consulRegisty),
		//绑定链路追踪
		micro.WrapClient(opentracing2.NewClientWrapper(opentracing.GlobalTracer())),
	)

	productService := product.NewProductService("product", service.Client())

	productAdd := &product.ProductInfo{
		ProductName:        "imooc",
		ProductSku:         "cap",
		ProductPrice:       1.1,
		ProductDescription: "imooc-cap",
		ProductCategoryId:  1,
		ProductImage: []*product.ProductImage{
			{
				ImageName: "cap-image",
				ImageCode: "capimage01",
				ImageUrl:  "capimage01",
			},
			{
				ImageName: "cap-image02",
				ImageCode: "capimage02",
				ImageUrl:  "capimage02",
			},
		},
		ProductSize: []*product.ProductSize{
			{
				SizeName: "cap-size",
				SizeCode: "cap-size-code",
			},
		},
		ProductSeo: &product.ProductSeo{
			SeoTitle:       "cap-seo",
			SeoKeywords:    "cap-seo",
			SeoDescription: "cap-seo",
			SeoCode:        "cap-seo",
		},
	}
	response, err := productService.AddProduct(context.TODO(), productAdd)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response)
}
