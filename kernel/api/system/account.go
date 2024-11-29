package system

import (
	"github.com/gin-gonic/gin"
	"kernel/model"
	"kernel/model/entity"
	"kernel/sql"
	"net/http"
)

func addAccount(c *gin.Context) {
	var account entity.Account
	if err := c.ShouldBindJSON(&account); err != nil {
		model.BadRequest(c)
		return
	}
	result := sql.DB.Create(&account) // 通过数据的指针来创建
	c.JSON(http.StatusOK, model.Success(result))
}
