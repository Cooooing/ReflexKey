package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"kernel/common/log"
	"kernel/model"
	"net/http"
	"time"
)

// Logging 打印请求信息
func Logging(c *gin.Context) {
	start := time.Now()
	path := c.Request.URL.Path
	raw := c.Request.URL.RawQuery

	c.Next()

	end := time.Now()
	latency := end.Sub(start)

	if raw != "" {
		path = path + "?" + raw
	}

	log.Info(" %d | %d | %s | %s | %s %s ", c.Writer.Status(), c.Writer.Size(), latency, c.ClientIP(), c.Request.Method, path)
}

// Recover 异常处理、日志记录
func Recover(c *gin.Context) {
	defer func() {
		if e := recover(); nil != e {
			log.RecoverError(e)
			c.JSON(http.StatusInternalServerError, model.Fail(fmt.Sprintf("%v", e)))
			c.Abort()
		}
	}()
	c.Next()
}

// CorsMiddleware 配置跨域请求
func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "origin, Content-Length, Content-Type, Authorization")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, HEAD, OPTIONS")
		c.Header("Access-Control-Allow-Private-Network", "true")

		if c.Request.Method == "OPTIONS" {
			c.Header("Access-Control-Max-Age", "600")
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
