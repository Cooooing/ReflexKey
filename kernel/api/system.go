package api

import (
	"fmt"
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
	arg := map[string]interface{}{}
	c.BindJSON(&arg)
	s := arg["sql"].(string)
	page := int64(arg["page"].(float64))
	size := int64(arg["size"].(float64))

	list := sql.QueryForPage(page, size, s)
	c.JSON(http.StatusOK, model.Success(list))
}

func db1(c *gin.Context) {
	arg := map[string]interface{}{}
	c.BindJSON(&arg)
	s := arg["sql"].(string)

	var s3 []s2
	sql.DB.Raw(s).Scan(&s3)
	for i := range s3 {
		fmt.Println(s3[i])
	}
	c.JSON(http.StatusOK, model.Success(s3))
}

type s2 struct {
	Id int64
}
