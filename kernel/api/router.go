package api

import (
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"kernel/conf"
	"net/http"
	"net/http/pprof"
	"strings"
)

//go:embed static/*
var frontend embed.FS

func ServeAPI(ginServer *gin.Engine) {

	api := ginServer.Group("/api")
	{
		system := api.Group("/system")
		{
			system.GET("/ping", ping)
			system.GET("/databaseVersion", databaseVersion)
			system.GET("/db", db)
			system.GET("/err", err)
		}
		encoding := api.Group("/encoding")
		{
			encoding.GET("/base64Encode", base64Encode)
			encoding.GET("/base64Decode", base64Decode)
		}
		crypto := api.Group("/crypto")
		{
			crypto.POST("/aesEncrypt", AesEncrypt)
			crypto.POST("/aesDecrypt", AesDecrypt)
		}
	}

	ginServer.Any("/static/*filepath", func(c *gin.Context) {
		staticServer := http.FileServer(http.FS(frontend))
		staticServer.ServeHTTP(c.Writer, c.Request)
	})

	// 加载前端的 index.html
	ginServer.GET("/", func(c *gin.Context) {
		file, err := frontend.ReadFile("static/index.html")
		if err != nil {
			c.String(http.StatusInternalServerError, "Error loading index.html")
			return
		}
		c.Data(http.StatusOK, "text/html; charset=utf-8", file)
	})

	// 捕获所有未匹配路由，返回 index.html（支持前端路由）
	ginServer.NoRoute(func(c *gin.Context) {
		path := c.Request.URL.Path
		fmt.Println(path)
		if strings.HasPrefix(path, "/api/") {
			notFound(c)
			return
		}

		file, err := frontend.ReadFile("static/index.html")
		if err != nil {
			c.String(http.StatusInternalServerError, "Error loading index.html")
			return
		}
		c.Data(http.StatusOK, "text/html; charset=utf-8", file)
	})

	// serveDebug 生产模式下关闭 pprof
	if conf.Prod == conf.Mode {
		return
	}

	ginServer.GET("/debug/pprof/", gin.WrapF(pprof.Index))
	ginServer.GET("/debug/pprof/allocs", gin.WrapF(pprof.Index))
	ginServer.GET("/debug/pprof/block", gin.WrapF(pprof.Index))
	ginServer.GET("/debug/pprof/goroutine", gin.WrapF(pprof.Index))
	ginServer.GET("/debug/pprof/heap", gin.WrapF(pprof.Index))
	ginServer.GET("/debug/pprof/mutex", gin.WrapF(pprof.Index))
	ginServer.GET("/debug/pprof/threadcreate", gin.WrapF(pprof.Index))
	ginServer.GET("/debug/pprof/cmdline", gin.WrapF(pprof.Cmdline))
	ginServer.GET("/debug/pprof/profile", gin.WrapF(pprof.Profile))
	ginServer.GET("/debug/pprof/symbol", gin.WrapF(pprof.Symbol))
	ginServer.GET("/debug/pprof/trace", gin.WrapF(pprof.Trace))

}
