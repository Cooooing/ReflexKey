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

	var scan []model.Config
	list := sql.QueryForPage(page, size, &scan, s)
	c.JSON(http.StatusOK, model.Success(list))
}
func db1(c *gin.Context) {
	ret := sql.GetDatabaseVer()
	c.JSON(http.StatusOK, model.Success(ret))
}
func db2(c *gin.Context) {
	sql.SetDatabaseVer()
	c.JSON(http.StatusOK, model.Success(""))
}

func err(c *gin.Context) {
	panic("err")
	// 无意抛出 panic
	var slice = []int{1, 2, 3, 4, 5}
	slice[6] = 6
	//c.JSON(http.StatusOK, model.Success(true))
}
