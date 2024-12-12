package system

import (
	"embed"
	"github.com/gin-gonic/gin"
	"kernel/conf"
	"kernel/model"
	"net/http"
	"strings"
)

//go:embed static/*
var frontend embed.FS

type SystemController struct {
}

func NewSystemController() *SystemController {
	return &SystemController{}
}

func (system *SystemController) InitRoutes(ginServer *gin.Engine) {
	// 首页路由
	ginServer.GET("/", system.index)
	// 静态资源路由
	ginServer.Any("/static/*filepath", system.static)
	// 捕获所有未匹配路由，前端路由返回 index.html 后端路由（/api/*）返回 404
	ginServer.NoRoute(system.noRoute)
}

func (system *SystemController) Routes(rg *gin.RouterGroup) {
	rg.GET("/ping", system.ping)
	rg.GET("/databaseVersion", system.databaseVersion)
}

func (system *SystemController) static(c *gin.Context) {
	staticServer := http.FileServer(http.FS(frontend))
	staticServer.ServeHTTP(c.Writer, c.Request)
}

func (system *SystemController) index(c *gin.Context) {
	file, err := frontend.ReadFile("static/index.html")
	if err != nil {
		c.String(http.StatusInternalServerError, "Error loading index.html")
		return
	}
	c.Data(http.StatusOK, "text/html; charset=utf-8", file)
}

func (system *SystemController) noRoute(c *gin.Context) {
	path := c.Request.URL.Path
	if strings.HasPrefix(path, "/api/") {
		c.JSON(http.StatusNotFound, model.Fail("404 page not found"))
		return
	}

	c.Redirect(http.StatusFound, "/")
}

// ShowAccount godoc
//	@Summary	服务器状态
//	@Tags		system
//	@Router		/system/ping [get]
func (system *SystemController) ping(c *gin.Context) {
	c.JSON(http.StatusOK, model.Success("pong"))
}

// ShowAccount godoc
//	@Summary	获取数据库版本
//	@Tags		system
//	@Router		/system/databaseVersion [get]
func (system *SystemController) databaseVersion(c *gin.Context) {
	c.JSON(http.StatusOK, model.Success(conf.DatabaseVersion))
}
