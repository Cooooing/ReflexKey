package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"kernel/api"
	"kernel/common"
	"kernel/conf"
	"kernel/middlewares"
	"kernel/util"
	"net"
	"net/http"
	"net/http/pprof"
	"net/url"
	"os"
	"time"
)

func Start() {
	if util.ServerIsRunning {
		return
	}

	gin.SetMode(gin.ReleaseMode)
	ginServer := gin.New()
	ginServer.UseH2C = true
	ginServer.MaxMultipartMemory = 1024 * 1024 * 32 // 表示处理上传的文件时，最多将32MB的数据保存在内存中，超出部分会保存到临时文件中。这样可以避免大文件上传时占用过多内存。
	ginServer.Use(
		middlewares.Logging,
		middlewares.Recover,
		middlewares.CorsMiddleware, // 后端服务支持 CORS 预检请求验证
	)

	serveDebug(ginServer)
	api.ServeAPI(ginServer)

	var host string
	if conf.Conf.System.NetworkServe {
		host = "0.0.0.0"
	} else {
		host = "127.0.0.1"
	}

	ln, err := net.Listen("tcp", host+":"+util.ServerPort)
	if nil != err {
		common.Log.Error("boot kernel failed: %s", err)
	}

	_, port, err := net.SplitHostPort(ln.Addr().String())
	if nil != err {
		common.Log.Error("boot kernel failed: %s", err)
	}
	util.ServerPort = port

	util.ServerURL, err = url.Parse("http://127.0.0.1:" + port)
	if err != nil {
		common.Log.Error("parse server url failed: %s", err)
	}

	pid := fmt.Sprintf("%d", os.Getpid())
	common.Log.Info("kernel [pid=%s] http server [%s] is booting", pid, host+":"+port)
	//util.HttpServing = true

	//go util.HookUILoaded()

	//go func() {
	//	time.Sleep(1 * time.Second)
	//	go proxy.InitFixedPortService(host)
	//	go proxy.InitPublishService()
	//	// 反代服务器启动失败不影响核心服务器启动
	//}()

	util.Server = &http.Server{
		Addr:    host + ":" + port,
		Handler: ginServer,
	}

	go func() {
		if err = util.Server.Serve(ln); nil != err && !errors.Is(http.ErrServerClosed, err) {
			util.ServerIsRunning = false
			common.Log.Fatal(common.ExitCodeUnavailablePort, "boot kernel failed: %s", err)
		}
	}()

	util.ServerIsRunning = true
	common.Log.Info("http server started")
}

func Stop() {
	common.Log.Info("shutdown server ...")

	// 创建一个 5 秒的超时上下文
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 关闭 HTTP Server
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := util.Server.Shutdown(ctx); err != nil {
		common.Log.Fatal(0, "server shutdown:", err)
	}
	util.ServerIsRunning = false
	common.Log.Info("server exited")
}

func Restart() {
	Stop()
	Start()
}

// serveDebug 生产模式下关闭 pprof
func serveDebug(ginServer *gin.Engine) {
	if util.Prod == util.Mode {
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
