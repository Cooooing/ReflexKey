package api

import (
	"github.com/gin-gonic/gin"
	"kernel/conf"
	"net/http"
	"net/http/pprof"
	"path"
	"path/filepath"
)

func ServeAPI(ginServer *gin.Engine) {

	ginServer.NoRoute(notFound)
	ginServer.NoRoute(notFound)

	ginServer.GET("/assets/*path", func(c *gin.Context) {
		requestPath := c.Param("path")
		relativePath := path.Join("assets", requestPath)
		p := filepath.Join(conf.DataDir, relativePath)
		http.ServeFile(c.Writer, c.Request, p)
		return
	})

	ginServer.StaticFile("/", "./workspace/data/index.html")
	ginServer.StaticFile("/index.html", "./workspace/data/index.html")

	ginServer.StaticFS("/more", http.Dir("./workspace"))
	ginServer.StaticFile("/favicon.ico", "./workspace/data/resources/icon.png")

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
