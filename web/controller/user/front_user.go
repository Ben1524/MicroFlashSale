package user

import (
	"context"
	"github.com/gin-gonic/gin"
	mgrpc "github.com/go-micro/plugins/v4/client/grpc"
	"github.com/go-micro/plugins/v4/registry/consul"
	mhttp "github.com/go-micro/plugins/v4/server/grpc"
	"go-micro.dev/v4"
	"log"
	"net/http"
	"sync"
	"time"
	front_user "user_srv/proto/front_user"
)

var (
	frontUserService front_user.FrontUserService
	once             sync.Once
)
var (
	name    = "user_srv"
	version = "latest"
)

func initService() {
	once.Do(func() {
		// 创建 Consul 注册中心
		consulReg := consul.NewRegistry()
		service := micro.NewService(
			micro.RegisterInterval(time.Second*10),
			micro.Client(mgrpc.NewClient()),
			micro.Server(mhttp.NewServer()),
		)
		opts := []micro.Option{
			micro.Registry(consulReg),
			micro.Name(name),
			micro.Version(version),
		}

		service.Init(opts...)

		// 创建 FrontUser 服务的客户端
		frontUserService = front_user.NewFrontUserService("user_srv", service.Client())
	})
}
func SendEmail(ctx *gin.Context) {
	// 确保服务已初始化（应该在应用启动时执行，而不是每次请求）
	initService()

	// 定义请求结构体，绑定 JSON 数据
	var request struct {
		Email string `json:"email" binding:"required,email"`
	}

	// 使用 ShouldBindJSON 解析 JSON 请求
	if err := ctx.ShouldBindJSON(&request); err != nil {
		log.Printf("JSON 解析失败: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "无效的请求格式",
			"details": err.Error(),
		})
		return
	}

	// 获取解析后的邮箱
	email := request.Email
	log.Printf("收到邮件请求: %s", email)

	// 调用微服务方法
	response, err := frontUserService.FrontUserSendEmail(context.Background(), &front_user.FrontUserMailRequest{
		Email: email,
	})

	if err != nil {
		log.Printf("微服务调用失败: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "邮件发送服务暂时不可用",
			"details": err.Error(),
		})
		return
	}

	// 检查微服务响应是否有效
	if response == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "邮件服务返回空响应",
		})
		return
	}

	// 返回成功响应
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "验证码已发送，请检查邮箱",
		"data":    response,
	})
}
