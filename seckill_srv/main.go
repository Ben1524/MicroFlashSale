package main

import (
	"github.com/go-micro/plugins/v4/registry/consul"
	"seckill_srv/handler"
	pb "seckill_srv/proto"

	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
)

var (
	service = "seckill_srv"
	version = "latest"
)

func main() {
	// Create service
	consulReg := consul.NewRegistry()
	srv := micro.NewService()
	srv.Init(
		micro.Address(":8082"),
		micro.Name(service),
		micro.Version(version),
		micro.Registry(consulReg),
	)
	// Register handler
	if err := pb.RegisterZhiliaoseckillsrvHandler(srv.Server(), new(handler.Zhiliaoseckillsrv)); err != nil {
		logger.Fatal(err)
	}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
