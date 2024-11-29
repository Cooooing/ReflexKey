package api

import (
	"github.com/gin-gonic/gin"
	"kernel/api/system"
	"kernel/api/tool"
	"kernel/conf"
	"net/http/pprof"
)

func ServeAPI(ginServer *gin.Engine) {

	// 静态资源路由
	systemController := system.NewSystemController()
	systemController.InitRoutes(ginServer)

	// api 路由
	api := ginServer.Group("/api")
	{
		systemGroup := api.Group("/system")
		{
			systemController.Routes(systemGroup)
		}
		tools := api.Group("/tools")
		{
			tool.NewEncodingController().Routes(tools.Group("/encoding"))
			tool.NewCryptoController().Routes(tools.Group("/crypto"))
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
