package api

import (
	"github.com/gin-gonic/gin"
	"kernel/model"
	"kernel/model/param"
	"kernel/util"
	"net/http"
)

// AesEncrypt AES加密
// @Summary AES加密
// @Description 使用AES算法加密数据
// @Tags 加密解密
// @Accept json
// @Produce json
// @Param params body param.AesEncryptParam true "加密参数"
// @Success 200 {object} result.Result "加密成功"
// @Router /api/crypto/aes/encrypt [post]
func AesEncrypt(c *gin.Context) {
	var params param.AesEncryptParam
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, model.BadRequest())
		return
	}

	cipherText, err := util.AesEncrypt(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Fail(err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.Success(cipherText))
}

// AesDecrypt AES解密
// @Summary AES解密
// @Description 使用AES算法解密数据
// @Tags 加密解密
// @Accept json
// @Produce json
// @Param params body param.AesDecryptParam true "解密参数"
// @Success 200 {object} result.Result "解密成功"
// @Router /api/crypto/aes/decrypt [post]
func AesDecrypt(c *gin.Context) {
	var params param.AesDecryptParam
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, model.BadRequest())
		return
	}

	plainText, err := util.AesDecrypt(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Fail(err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.Success(plainText))
}
