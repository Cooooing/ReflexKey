package tool

import (
	"github.com/gin-gonic/gin"
	"kernel/common"
	"kernel/model"
	"kernel/model/param"
	"kernel/util"
	"net/http"
)

type CryptoController struct {
}

func NewCryptoController() *CryptoController {
	return &CryptoController{}
}

func (crypto *CryptoController) Routes(rg *gin.RouterGroup) {
	rg.POST("/aesEncrypt", crypto.aesEncrypt)
	rg.POST("/aesDecrypt", crypto.aesDecrypt)
	rg.POST("/bcryptHash", crypto.bcryptHash)
	rg.POST("/bcryptCompare", crypto.bcryptCompare)
	rg.POST("/bcryptCost", crypto.bcryptCost)
}

func (crypto *CryptoController) aesEncrypt(c *gin.Context) {
	var (
		params param.AesEncryptParam
		result string
		err    error
	)
	if err := c.ShouldBindJSON(&params); err != nil {
		common.Log.Error("aesEncrypt params error:", err)
		model.BadRequest(c)
		return
	}

	cipherText, err := util.AesEncrypt(params)
	if err != nil {
		common.Log.Error("aesEncrypt error:", err)
		c.JSON(http.StatusInternalServerError, model.Fail(err.Error()))
		return
	}

	result = cipherText
	c.JSON(http.StatusOK, model.Success(result))
}

func (crypto *CryptoController) aesDecrypt(c *gin.Context) {
	var (
		params param.AesDecryptParam
		result string
		err    error
	)
	if err := c.ShouldBindJSON(&params); err != nil {
		common.Log.Error("aesDecrypt params error:", err)
		model.BadRequest(c)
		return
	}

	plainText, err := util.AesDecrypt(params)
	if err != nil {
		common.Log.Error("aesDecrypt error:", err)
		c.JSON(http.StatusInternalServerError, model.Fail(err.Error()))
		return
	}

	result = plainText
	c.JSON(http.StatusOK, model.Success(result))
}

func (crypto *CryptoController) bcryptHash(c *gin.Context) {
	var (
		params param.BcryptHashParam
		result string
		err    error
	)
	if err := c.ShouldBindJSON(&params); err != nil {
		common.Log.Error("bcryptHash params error:", err)
		model.BadRequest(c)
		return
	}

	hashedPassword, err := util.BcryptHash([]byte(params.Password), params.Cost)
	if err != nil {
		common.Log.Error("bcryptHash error:", err)
		c.JSON(http.StatusInternalServerError, model.Fail(err.Error()))
		return
	}

	result = string(hashedPassword)
	c.JSON(http.StatusOK, model.Success(result))
}

func (crypto *CryptoController) bcryptCompare(c *gin.Context) {
	var (
		params param.BcryptCompareParam
		result string
		err    error
	)
	if err := c.ShouldBindJSON(&params); err != nil {
		common.Log.Error("bcryptCompare params error:", err)
		model.BadRequest(c)
		return
	}

	err = util.BcryptCompare([]byte(params.HashedPassword), []byte(params.Password))
	if err != nil {
		common.Log.Error("bcryptCompare error:", err)
		c.JSON(http.StatusInternalServerError, model.Fail(err.Error()))
		return
	}

	result = "crypto/bcrypt: hashedPassword is the hash of the given password"
	c.JSON(http.StatusOK, model.Success(result))
}

func (crypto *CryptoController) bcryptCost(c *gin.Context) {
	var (
		params param.BcryptCostParam
		result int
		err    error
	)
	if err := c.ShouldBindJSON(&params); err != nil {
		common.Log.Error("bcryptCost params error:", err)
		model.BadRequest(c)
		return
	}

	cost, err := util.BcryptCost([]byte(params.HashedPassword))
	if err != nil {
		common.Log.Error("bcryptCost error:", err)
		c.JSON(http.StatusInternalServerError, model.Fail(err.Error()))
		return
	}

	result = cost
	c.JSON(http.StatusOK, model.Success(result))
}
