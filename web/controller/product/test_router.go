package product

import "github.com/gin-gonic/gin"

func Test(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"code": 200,
		"msg":  "ok",
	})
}
