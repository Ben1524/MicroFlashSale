// package main
//
// import (
//
//	"github.com/gin-gonic/gin"
//	"web/middle_ware"
//	allRouter "web/router"
//
// )
//
// func main() {
//
//		router := gin.Default()
//		// 中间件
//		router.Use(middle_ware.CrosMiddleWare)
//		// 注册路由组
//		allRouter.InitRouter(router)
//
//		//router
//
//		// 启动gin
//		if err := router.Run(":8082"); err != nil {
//			panic(err)
//		}
//
//		//userToken, _ := utils.GenToken("user", utils.FrontUserExpireDuration, utils.FrontUserSecret)
//		//fmt.Println(userToken)
//		//token, _ := utils.AuthToken(userToken, utils.FrontUserSecret)
//		//fmt.Println(token.UserName)
//		//// unix-->标准时间
//		//fmt.Println(time.Unix(token.ExpiresAt, 0).Format("2006-01-02 15:04:05"))
//	}
package main

import (
	"context"
	"github.com/go-micro/plugins/v4/client/grpc"
	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4"
	"log"
	front_user "user_srv/proto/front_user" // 根据实际路径调整
)

func main() {
	// 创建 Consul 注册中心
	consulReg := consul.NewRegistry()

	// 创建使用 gRPC 协议的微服务客户端
	service := micro.NewService(
		micro.Name("user_client"),
		micro.Registry(consulReg),
		micro.Client(grpc.NewClient(
		//client.WithRequestTimeout(10*time.Second), // 正确使用WithRequestTimeout
		)),
	)
	service.Init()

	// 创建 FrontUser 服务的客户端
	client := front_user.NewFrontUserService("user_srv", service.Client())

	// 准备请求
	request := &front_user.FrontUserMailRequest{
		Email: "test@example.com",
	}

	// 调用服务方法
	response, err := client.FrontUserSendEmail(context.Background(), request)
	if err != nil {
		log.Fatalf("调用服务失败: %v", err)
	}

	log.Printf("服务响应: %+v", response)
}
