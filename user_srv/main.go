package main

import (
	"fmt"
	grpcc "github.com/go-micro/plugins/v4/client/grpc"
	"github.com/go-micro/plugins/v4/registry/consul"
	grpcs "github.com/go-micro/plugins/v4/server/grpc"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/registry"
	"log"
	"user_srv/conf"
	"user_srv/controller"
	_ "user_srv/data_source"
	admin_user "user_srv/proto/admin_user"
	front_user "user_srv/proto/front_user"
)

var (
	service = conf.CONFIG.MicroServices[2].Name
	version = conf.CONFIG.MicroServices[2].Version
	address = fmt.Sprintf("%s:%d", conf.CONFIG.MicroServices[2].Host, conf.CONFIG.MicroServices[2].Port)
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

	// 注册前端用户服务处理器
	if err := front_user.RegisterFrontUserHandler(srv.Server(), new(controller.FrontUserHandler)); err != nil {
		log.Fatalf("注册 FrontUser 处理器失败: %v", err)
	}

	// 注册管理员服务处理器
	if err := admin_user.RegisterAdminUserHandler(srv.Server(), new(controller.AdminUserHandler)); err != nil {
		log.Fatalf("注册 AdminUser 处理器失败: %v", err)
	}

	// 启动服务
	logger.Info("服务启动中...")
	if err := srv.Run(); err != nil {
		logger.Fatalf("服务运行失败: %v", err)
	}
}
