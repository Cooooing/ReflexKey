package api

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"kernel/model"
	"net/http"
)

func base64Encode(c *gin.Context) {
	data := c.Query("data")
	res := base64.StdEncoding.EncodeToString([]byte(data))
	c.JSON(http.StatusOK, model.Success(res))
}

func base64Decode(c *gin.Context) {
	data := c.Query("data")
	res, err := base64.StdEncoding.DecodeString(data)
	if nil != err {
		c.JSON(http.StatusOK, model.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.Success(string(res)))
}
