package api

import "github.com/gin-gonic/gin"

func ServeAPI(ginServer *gin.Engine) {

	ginServer.NoRoute(notFound)
	ginServer.NoRoute(notFound)

	ginServer.Handle("GET", "/api/system/ping", ping)

}
