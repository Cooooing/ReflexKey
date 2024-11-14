package api

import (
	"github.com/gin-gonic/gin"
	"kernel/model"
	"kernel/sql"
	"net/http"
)

func addAccount(c *gin.Context) {
	var account model.Account
	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, model.BadRequest())
		return
	}
	result := sql.DB.Create(&account) // 通过数据的指针来创建
	c.JSON(http.StatusOK, model.Success(result))
}
