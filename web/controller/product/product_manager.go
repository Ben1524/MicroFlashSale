package product

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
	product "product_srv/proto/product"
	"sync"
	"time"
)

var (
	productService product.ProductsService
	productOnce    sync.Once
	name           = "product_srv"
	version        = "latest"
)

func initProductService() {
	productOnce.Do(func() {
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

		// 创建 Product 服务的客户端
		productService = product.NewProductsService("product_srv", service.Client())
	})
}

func GetProductList(ctx *gin.Context) {
	initProductService()
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

	response, err := productService.ProductsList(context.Background(), &product.ProductsRequest{
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
		"success":  true,
		"message":  "获取用户列表成功",
		"products": response.Products,
		"total":    response.Total,
	})
}
