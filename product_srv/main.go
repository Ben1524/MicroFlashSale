package main

import (
	"context"
	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"product_srv/handler"
	pb "product_srv/proto/product"
	seckill "product_srv/proto/seckill"
	"time"
)

var (
	service = "product_srv"
	version = "latest"
)

func main() {
	// Create service
	consulReg := consul.NewRegistry()
	srv := micro.NewService()
	srv.Init(
		micro.Address(":8083"),
		micro.Name(service),
		micro.Version(version),
		micro.Registry(consulReg),
	)

	serviceSeckill := seckill.NewSeckillsrvService("seckill_srv", srv.Client())
	req := &seckill.CallRequest{
		Name: "abc",
	}

	// 重试机制
	maxRetries := 3
	for i := 0; i < maxRetries; i++ {
		rsp, err := serviceSeckill.Call(context.Background(), req)
		if err != nil {
			if i == maxRetries-1 {
				logger.Fatalf("Failed after %d retries: %v", maxRetries, err)
			}
			logger.Warnf("Retry %d: %v", i+1, err)
			time.Sleep(2 * time.Second) // 等待 2 秒后重试
			continue
		}
		logger.Infof("rsp: %v", rsp)
		break
	}

	// Register handler
	if err := pb.RegisterZhiliaoproductsrvHandler(srv.Server(), new(handler.Zhiliaoproductsrv)); err != nil {
		logger.Fatal(err)
	}

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
