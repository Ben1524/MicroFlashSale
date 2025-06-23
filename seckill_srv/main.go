package main

import (
	"fmt"
	"github.com/go-micro/plugins/v4/registry/consul"
	"seckill_srv/conf"
	"seckill_srv/handler"
	pb "seckill_srv/proto"

	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
)

var (
	service = conf.CONFIG.MicroServices[1].Name
	version = conf.CONFIG.MicroServices[1].Version
	address = fmt.Sprintf("%s:%d", conf.CONFIG.MicroServices[1].Host, conf.CONFIG.MicroServices[1].Port)
)

func main() {
	// Create service
	consulReg := consul.NewRegistry()
	srv := micro.NewService()
	srv.Init(
		micro.Address(address),
		micro.Name(service),
		micro.Version(version),
		micro.Registry(consulReg),
	)
	// Register handler
	if err := pb.RegisterSeckillsrvHandler(srv.Server(), new(handler.Seckillsrv)); err != nil {
		logger.Fatal(err)
	}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
