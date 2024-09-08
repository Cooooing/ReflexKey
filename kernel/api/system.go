package api

import (
	"github.com/gin-gonic/gin"
	"kernel/model"
	"net/http"
)

func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, model.Fail("404 page not found"))
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, model.Success("pong"))
}
