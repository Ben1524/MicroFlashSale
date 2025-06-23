package main

import (
	"fmt"
	grpcc "github.com/go-micro/plugins/v4/client/grpc"
	"github.com/go-micro/plugins/v4/registry/consul"
	grpcs "github.com/go-micro/plugins/v4/server/grpc"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/registry"
	conf "product_srv/conf"
	"product_srv/controller"
	product "product_srv/proto/product"
	//seckill "product_srv/proto/seckill"
)

var (
	service = conf.CONFIG.MicroServices[0].Name
	version = conf.CONFIG.MicroServices[0].Version
	address = fmt.Sprintf("%s:%d", conf.CONFIG.MicroServices[0].Host, conf.CONFIG.MicroServices[0].Port)
)

func main() {
	// 创建 Consul 注册中心，明确指定地址
	consulReg := consul.NewRegistry(
		// 如果 Consul 不在本地运行，需要指定实际地址
		registry.Addrs("localhost:8500"),
	)

	// 创建微服务实例
	srv := micro.NewService(
		micro.Client(grpcc.NewClient()),
		micro.Server(grpcs.NewServer()),
	)
	opts := []micro.Option{
		micro.Registry(consulReg), // 使用 Consul 作为服务发现
		micro.Name(service),
		micro.Version(version),
		micro.Address(address),
	}
	// 初始化服务
	srv.Init(opts...)

	// Register handlers
	if err := product.RegisterProductsHandler(srv.Server(), new(controller.ProductHandler)); err != nil {
		logger.Fatal(err)
	}
	//if err := seckill.RegisterSecKillsHandler(srv.Server(), new(controller.SecKillsHandler)); err != nil {
	//	logger.Fatal(err)
	//}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
