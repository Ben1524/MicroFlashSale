package product

import (
	"github.com/gin-gonic/gin"
	"web/middle_ware"
)

func Router(router *gin.RouterGroup) {

	router.GET("/products", middle_ware.JwtTokenAdminValid, GetProductList) // 获取商品列表

	router.POST("/products", middle_ware.JwtTokenAdminValid, AddProduct) // 添加商品

	//// 商品列表
	//router.GET("/get_product_list", GetProductList)
	//// 商品详情
	//router.GET("/get_product_detail", GetProductDetail)
	//// 商品搜索
	//router.GET("/search_product", SearchProduct)
	//// 商品分类列表
	//router.GET("/get_product_category_list", GetProductCategoryList)
	//// 商品分类详情
	//router.GET("/get_product_category_detail", GetProductCategoryDetail)
}
