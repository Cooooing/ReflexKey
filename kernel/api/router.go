package api

import (
	"embed"
	"github.com/gin-gonic/gin"
	"kernel/conf"
	"net/http/pprof"
)

//go:embed static/*
var frontend embed.FS

func ServeAPI(ginServer *gin.Engine) {

	// api 路由
	api := ginServer.Group("/api")
	{
		system := api.Group("/system")
		{
			system.GET("/ping", ping)
			system.GET("/databaseVersion", databaseVersion)
			system.GET("/db", db)
			system.GET("/err", err)
		}
		auth := api.Group("/auth")
		{
			auth.POST("/login", login)
		}
		account := api.Group("/account")
		{
			account.POST("/addAccount", addAccount)
		}
		tools := api.Group("/tools")
		{
			encoding := tools.Group("/encoding")
			{
				encoding.GET("/base64Encode", base64Encode)
				encoding.GET("/base64Decode", base64Decode)
			}
			crypto := tools.Group("/crypto")
			{
				crypto.POST("/aesEncrypt", AesEncrypt)
				crypto.POST("/aesDecrypt", AesDecrypt)
				crypto.POST("/bcryptHash", BcryptHash)
				crypto.POST("/bcryptCompare", BcryptCompare)
				crypto.POST("/bcryptCost", BcryptCost)
			}
		}
	}

	// 静态资源路由
	ginServer.Any("/static/*filepath", static)

	// 首页路由
	ginServer.GET("/", index)

	// 捕获所有未匹配路由，前端路由返回 index.html 后端路由（/api/*）返回 404
	ginServer.NoRoute(noRoute)

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
