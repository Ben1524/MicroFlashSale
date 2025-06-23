package middle_ware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// 处理跨域请求的中间件

func CrosMiddleWare(c *gin.Context) {
	origin := c.Request.Header.Get("origin")
	if len(origin) == 0 {
		origin = c.Request.Header.Get("Origin") // 兼容处理
	}

	if origin == "" {
		log.Println("origin is empty")
	}

	// 设置响应头，允许指定的来源进行跨域请求
	c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
	// 设置响应头，允许携带凭证（如 cookies、HTTP 认证等）进行跨域请求
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	// 设置响应头，允许的请求头字段，客户端可以在请求中包含这些字段
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	// 设置响应头，允许的请求方法，客户端可以使用这些方法进行跨域请求
	c.Writer.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST")
	// 设置响应头，指定响应内容的类型为 JSON 格式，字符编码为 UTF-8
	c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	// 如果请求方法是 OPTIONS，通常是浏览器在发送实际请求之前的预检请求
	if c.Request.Method == "OPTIONS" {
		// 直接返回 204 状态码，表示请求已成功处理，但没有返回内容
		c.AbortWithStatus(http.StatusNoContent)
		// 结束当前处理流程，不再继续执行后续的中间件或处理函数
		return
	}
	// 输出新的origin
	log.Println("origin:", origin)
	// 继续执行后续的中间件或处理函数
	c.Next()
}
