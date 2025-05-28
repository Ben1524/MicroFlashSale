package user

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	mgrpc "github.com/go-micro/plugins/v4/client/grpc"
	"github.com/go-micro/plugins/v4/registry/consul"
	mhttp "github.com/go-micro/plugins/v4/server/grpc"
	"go-micro.dev/v4"
	"log"
	"net/http"
	"sync"
	"time"
	admin_user "user_srv/proto/admin_user"
	"web/utils"
)

var (
	adminUserService admin_user.AdminUserService
	adminOnce        sync.Once
)

func initAdminService() {
	adminOnce.Do(func() {
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
		adminUserService = admin_user.NewAdminUserService("user_srv", service.Client())
	})
}
func AdminLogin(ctx *gin.Context) {
	initAdminService()
	// 定义请求结构体，绑定 JSON 数据
	var request struct {
		Username string `json:"username" binding:"required"`
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

	fmt.Printf("接收到的请求: %+v\n", request)

	// 获取解析后的用户名和密码
	username := request.Username
	password := request.Password

	// 调用微服务方法
	response, err := adminUserService.AdminUserLogin(context.Background(), &admin_user.AdminUserRequest{
		Username: username,
		Password: password,
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

	// 生成jwt
	jwt, err := utils.GenToken(username, utils.AdminUserExpireDuration, utils.AdminUserSecretKey)

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
		"username": username,
	})
}

func GetUserList(ctx *gin.Context) {
	initAdminService()
	// 定义请求结构体，绑定查询参数（适用于GET请求）
	var request struct {
		CurrentPage int64  `form:"CurrentPage" binding:"required,gte=1"` // gte=1 确保页码≥1
		PageSize    int64  `form:"PageSize" binding:"required,gte=1"`    // 每页数量≥1
		Search      string `form:"Search"`
		Status      int    `form:"Status"`
		SortField   string `form:"SortField"`
		SortOrder   string `form:"SortOrder"` // "asc" 或 "desc"
	}

	// 使用 ShouldBindQuery 解析查询参数（适用于GET请求）
	if err := ctx.ShouldBindQuery(&request); err != nil {
		log.Printf("参数解析失败: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "无效的请求参数",
			"details": err.Error(),
		})
		return
	}

	fmt.Printf("接收到的请求: %+v\n", request)

	currentPage := request.CurrentPage
	pageSize := request.PageSize

	response, err := adminUserService.FrontUserList(context.Background(), &admin_user.FrontUsersRequest{
		CurrentPage: currentPage,
		Pagesize:    pageSize,
	})

	if err != nil {
		log.Printf("微服务调用失败: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "获取用户列表服务暂时不可用",
			"details": err.Error(),
		})
		return
	}

	if response == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "获取用户列表服务返回空响应",
			"details": "请稍后再试",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "获取用户列表成功",
		"users":   response.FrontUsers,
		"total":   response.Total,
	})
}
