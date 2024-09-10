package api

import (
	"github.com/gin-gonic/gin"
	"kernel/util"
	"net/http"
	"path"
	"path/filepath"
)

func ServeAPI(ginServer *gin.Engine) {

	ginServer.NoRoute(notFound)
	ginServer.NoRoute(notFound)

	ginServer.GET("/assets/*path", func(c *gin.Context) {
		requestPath := c.Param("path")
		relativePath := path.Join("assets", requestPath)
		p := filepath.Join(util.DataDir, relativePath)
		http.ServeFile(c.Writer, c.Request, p)
		return
	})

	ginServer.StaticFile("/", "./workspace/data/index.html")
	ginServer.StaticFile("/index.html", "./workspace/data/index.html")

	ginServer.StaticFS("/more", http.Dir("./workspace"))
	ginServer.StaticFile("/favicon.ico", "./workspace/data/resources/icon.png")

	ginServer.Handle("GET", "/api/system/ping", ping)
	ginServer.Handle("GET", "/api/system/db", db)

}
