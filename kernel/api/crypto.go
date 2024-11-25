package api

import (
	"github.com/gin-gonic/gin"
	"kernel/model"
	"kernel/model/param"
	"kernel/util"
	"net/http"
)

func AesEncrypt(c *gin.Context) {
	var params param.AesEncryptParam
	if err := c.ShouldBindJSON(&params); err != nil {
		model.BadRequest(c)
		return
	}

	cipherText, err := util.AesEncrypt(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Fail(err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.Success(cipherText))
}

func AesDecrypt(c *gin.Context) {
	var params param.AesDecryptParam
	if err := c.ShouldBindJSON(&params); err != nil {
		model.BadRequest(c)
		return
	}

	plainText, err := util.AesDecrypt(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Fail(err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.Success(plainText))
}

func BcryptHash(c *gin.Context) {
	var params param.BcryptHashParam
	if err := c.ShouldBindJSON(&params); err != nil {
		model.BadRequest(c)
		return
	}

	hashedPassword, err := util.BcryptHash([]byte(params.Password), params.Cost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Fail(err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.Success(hashedPassword))
}

func BcryptCompare(c *gin.Context) {
	var params param.BcryptCompareParam
	if err := c.ShouldBindJSON(&params); err != nil {
		model.BadRequest(c)
		return
	}

	err := util.BcryptCompare([]byte(params.HashedPassword), []byte(params.Password))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.Success("crypto/bcrypt: hashedPassword is the hash of the given password"))
}

func BcryptCost(c *gin.Context) {
	var params param.BcryptCostParam
	if err := c.ShouldBindJSON(&params); err != nil {
		model.BadRequest(c)
		return
	}
	cost, err := util.BcryptCost([]byte(params.HashedPassword))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.Success(cost))
}
