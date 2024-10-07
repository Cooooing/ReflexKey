package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"kernel/common"
	"kernel/model"
	"kernel/util"
	"net/http"
	"time"
)

// Logging 打印请求信息
func Logging(c *gin.Context) {
	start := time.Now()
	path := c.Request.URL.Path
	raw := c.Request.URL.RawQuery
	defer func() {
		end := time.Now()
		latency := end.Sub(start)

		if raw != "" {
			path = path + "?" + raw
		}

		common.Log.Info(" %d | %d | %s | %s | %s %s ", c.Writer.Status(), c.Writer.Size(), latency, c.ClientIP(), c.Request.Method, path)
	}()
	c.Next()
}

// Recover 异常处理、日志记录
func Recover(c *gin.Context) {
	defer func() {
		if e := recover(); nil != e {
			common.Log.RecoverError(e)
			if util.Dev == util.Mode {
				c.AbortWithStatusJSON(http.StatusInternalServerError, model.Fail(fmt.Sprintf("%v", e)))
			} else {
				c.AbortWithStatusJSON(http.StatusInternalServerError, model.Fail("Internal Server Error"))
			}
		}
	}()
	c.Next()
}

// CorsMiddleware 配置跨域请求
func CorsMiddleware(c *gin.Context) {
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
