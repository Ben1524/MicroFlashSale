package router

import (
	"github.com/gin-gonic/gin"
	"web/controller/product"
	"web/controller/seckill"
	"web/controller/user"
)

func InitRouter(router *gin.Engine) {

	userGroup := router.Group("/user")
	productGroup := router.Group("/product")
	seckillGroup := router.Group("/seckill")

	user.Router(userGroup)       // 用户模块路由
	product.Router(productGroup) // 产品模块路由，调用参数不再是engine，而是group
	seckill.Router(seckillGroup)
}
