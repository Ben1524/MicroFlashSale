package user

import (
	"github.com/gin-gonic/gin"
	"web/middle_ware"
)

func Router(router *gin.RouterGroup) {

	//// 用户端注册登录
	router.POST("/send_email", SendEmail)
	router.POST("/front_user_login", FrontUserLogin)
	router.POST("/front_user_register", FrontUserRegister)
	//
	//// 管理端登录
	router.POST("/admin_login", AdminLogin)
	//中间件类似处理函数
	router.GET("/users", middle_ware.JwtTokenAdminValid, GetUserList)

}
