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
	"web/utils"
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

func FrontUserRegister(ctx *gin.Context) {
	// 确保服务已初始化（应该在应用启动时执行，而不是每次请求）
	initService()

	// 定义请求结构体，绑定 JSON 数据
	var request struct {
		Email    string `json:"email" binding:"required,email"`
		Code     string `json:"code" binding:"required"`
		Password string `json:"password" binding:"required"`
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

	// 获取解析后的邮箱、验证码和密码
	email := request.Email
	code := request.Code
	password := request.Password
	log.Printf("收到注册请求: %s", email)

	// 调用微服务方法
	response, err := frontUserService.FrontUserRegister(context.Background(), &front_user.FrontUserRequest{
		Email:    email,
		Code:     code,
		Password: password,
	})

	if err != nil {
		log.Printf("微服务调用失败: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "注册服务暂时不可用",
			"details": err.Error(),
		})
		return
	}

	if response == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "注册服务返回空响应",
		})
		return
	}

	if response.Code != 200 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "注册失败，验证码错误或邮箱已被注册",
			"data":  response,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func FrontUserLogin(ctx *gin.Context) {
	// 确保服务已初始化（应该在应用启动时执行，而不是每次请求）
	initService()

	// 定义请求结构体，绑定 JSON 数据
	var request struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
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

	// 获取解析后的邮箱和密码
	email := request.Email
	password := request.Password
	log.Printf("收到登录请求: %s", email)

	// 调用微服务方法
	response, err := frontUserService.FrontUserLogin(context.Background(), &front_user.FrontUserRequest{
		Email:    email,
		Password: password,
	})

	if err != nil {
		log.Printf("微服务调用失败: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "登录服务暂时不可用",
			"details": err.Error(),
		})
		return
	}

	if response == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "登录服务返回空响应",
		})
		return
	}

	if response.Code != 200 {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "登录失败，用户名或密码错误",
			"data":  response,
		})
		return
	}

	// 生成jwt
	jwt, err := utils.GenToken(email, utils.FrontUserExpireDuration, utils.FrontUserSecretKey)

	if err != nil {
		log.Printf("JWT 生成失败: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "登录服务暂时不可用",
			"details": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success":  true,
		"message":  "登录成功",
		"data":     response,
		"token":    jwt,
		"username": email,
	})
}
