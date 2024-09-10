package api

import (
	"github.com/gin-gonic/gin"
	"kernel/model"
	"kernel/sql"
	"net/http"
)

func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, model.Fail("404 page not found"))
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, model.Success("pong"))
}

func db(c *gin.Context) {
	arg := map[string]any{}
	c.BindJSON(&arg)
	s := arg["sql"].(string)
	page := int64(arg["page"].(float64))
	size := int64(arg["size"].(float64))

	list := sql.QueryForPage(page, size, s)
	c.JSON(http.StatusOK, model.Success(list))
}
