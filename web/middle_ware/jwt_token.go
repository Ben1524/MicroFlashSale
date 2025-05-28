package middle_ware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
	"web/utils"
)

// 对请求进行JWT验证的中间件
func JwtTokenAdminValid(ctx *gin.Context) {

	jwt_head := ctx.Request.Header.Get("Authorization") // Authorization 字段的格式通常为 Bearer <token>
	if jwt_head == "" {
		ctx.JSON(401, gin.H{
			"code": 401,
			"msg":  "请携带token",
		})
		ctx.Abort() // 中止请求，不再执行后续的处理函数
		log.Fatal("请携带token")
		return
	}
	auths := strings.Split(jwt_head, " ")

	bearer := auths[0] // bearer表示携带token，
	token := auths[1]
	if len(token) == 0 || len(bearer) == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 401,
			"msg":  "请携带正确格式的token",
		})
		ctx.Abort()
		log.Fatal("请携带正确格式的token")
		return
	}
	user, err := utils.AuthToken(token, utils.AdminUserSecretKey)

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 401,
			"msg":  "无效的token",
		})
		ctx.Abort()
		return
	}

	ctx.Set("admin_user_name", user.UserName) // 将解析后的用户名存储到gin的上下文中，
	ctx.Next()                                // 调用下一个中间件或处理函数
}

func JwtTokenFrontValid(ctx *gin.Context) {
	jwt_head := ctx.Request.Header.Get("Authorization")
	if jwt_head == "" {
		ctx.JSON(401, gin.H{
			"code": 401,
			"msg":  "请携带token",
		})
		ctx.Abort() // 中止请求，不再执行后续的处理函数
		log.Fatal("请携带token")
		return
	}
	auths := strings.Split(jwt_head, " ")

	bearer := auths[0] // 获取请求头中的Bearer类型，这是JWT的标准格式，表示
	token := auths[1]
	if len(token) == 0 || len(bearer) == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 401,
			"msg":  "请携带正确格式的token",
		})
		ctx.Abort()
		log.Fatal("请携带正确格式的token")
		return
	}
	user, err := utils.AuthToken(token, utils.FrontUserSecretKey)

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 401,
			"msg":  "无效的token",
		})
		ctx.Abort()
		return
	}

	// 将解析后的用户名存储到gin的上下文中，
	ctx.Set("front_user_name", user.UserName) // 将解析后的用户名存储到gin的上下文中，
	ctx.Next()                                // 调用下一个中间件或处理函数
}
