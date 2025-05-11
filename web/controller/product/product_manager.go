package product

import "github.com/gin-gonic/gin"

func GetProductList(ctx *gin.Context) {
	currentPage := ctx.DefaultQuery("currentPage", "1") // 带有默认值的查询参数
	pageSize := ctx.DefaultQuery("pageSize", "10")
	ctx.Qu
}
