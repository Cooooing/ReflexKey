package api

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"kernel/model"
	"kernel/model/param"
	"net/http"
)

func aesEncrypt(c *gin.Context) {
	var params param.AesEncryptParam
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, model.BadRequest())
		return
	}

	// 校验参数

	c.JSON(http.StatusOK, model.Success(""))
}

func aesDecrypt(c *gin.Context) {

	data := c.Query("data")
	res := base64.StdEncoding.EncodeToString([]byte(data))
	c.JSON(http.StatusOK, model.Success(res))
}
