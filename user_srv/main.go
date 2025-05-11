package main

import (
	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4"
	"user_srv/controller"
	_ "user_srv/data_source"
	admin_user "user_srv/proto/admin_user"
	front_user "user_srv/proto/front_user"
)

var (
	service = "user_srv"
	version = "latest"
)

func main() {
	// Create service
	consulReg := consul.NewRegistry()
	srv := micro.NewService()
	srv.Init(
		micro.Address(":8081"),
		micro.Name(service),
		micro.Version(version),
		micro.Registry(consulReg),
	)

	// Register
	err := front_user.RegisterFrontUserHandler(srv.Server(), new(controller.FrontUserHandler))
	if err != nil {
		return
	}
	err = admin_user.RegisterAdminUserHandler(srv.Server(), new(controller.AdminUserHandler))
	if err != nil {
		return
	}

	// Run service
	if err := srv.Run(); err != nil {
		panic(err)
	}

}
