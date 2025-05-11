package main

import (
	"github.com/gin-gonic/gin"
	allRouter "web/router"
)

func main() {

	router := gin.Default()
	// 注册路由组
	allRouter.InitRouter(router)

	// 启动gin
	if err := router.Run(":8081"); err != nil {
		panic(err)
	}

	//userToken, _ := utils.GenToken("user", utils.FrontUserExpireDuration, utils.FrontUserSecret)
	//fmt.Println(userToken)
	//token, _ := utils.AuthToken(userToken, utils.FrontUserSecret)
	//fmt.Println(token.UserName)
	//// unix-->标准时间
	//fmt.Println(time.Unix(token.ExpiresAt, 0).Format("2006-01-02 15:04:05"))
}
