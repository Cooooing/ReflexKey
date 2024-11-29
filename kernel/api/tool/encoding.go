package tool

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"kernel/common"
	"kernel/model"
	"kernel/model/param"
	"kernel/util"
	"net/http"
	"strconv"
)

type EncodingController struct {
}

func NewEncodingController() EncodingController {
	return EncodingController{}
}

func (encoding EncodingController) Routes(rg *gin.RouterGroup) {
	rg.POST("/base64Encode", encoding.base64Encode)
	rg.POST("/base64Decode", encoding.base64Decode)
	rg.POST("/base64EncodeImg", encoding.base64EncodeImg)
	rg.POST("/base64DecodeImg", encoding.base64DecodeImg)
}

func (encoding EncodingController) base64Encode(c *gin.Context) {
	var params model.Tuple[string]
	if err := c.ShouldBindJSON(&params); err != nil {
		common.Log.Error("base64Encode params error:", err)
		model.BadRequest(c)
		return
	}
	res := util.Base64Encode([]byte(params.T1))
	c.JSON(http.StatusOK, model.Success(res))
}

func (encoding EncodingController) base64Decode(c *gin.Context) {
	var params model.Tuple[string]
	if err := c.ShouldBindJSON(&params); err != nil {
		common.Log.Error("base64Decode params error:", err)
		model.BadRequest(c)
		return
	}
	res, err := util.Base64Decode(params.T1)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.Success(res))
}

func (encoding EncodingController) base64EncodeImg(c *gin.Context) {
	//file, err := c.FormFile("data")
	var params param.Base64EncodeParam
	err := c.ShouldBind(&params)
	fmt.Println(params)
	if err != nil {
		common.Log.Error("base64EncodeImg params error:", err)
		model.BadRequest(c)
		return
	}
	fileRead, err := params.File.Open()
	defer fileRead.Close()
	bytes, err := io.ReadAll(fileRead)
	data := util.Base64Encode(bytes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.Success(data))
}

func (encoding EncodingController) base64DecodeImg(c *gin.Context) {
	var params model.Tuple[string]
	if err := c.ShouldBindJSON(&params); err != nil {
		common.Log.Error("base64Decode params error:", err)
		model.BadRequest(c)
		return
	}
	img, err := util.Base64Decode(params.T1)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Fail(err.Error()))
		return
	}

	// 设置响应头
	c.Header("Content-Type", "image/jpeg")
	c.Header("Content-Disposition", "attachment; filename=temp.jpg")
	c.Header("Content-Length", strconv.FormatInt(int64(len(img)), 10))
	// 返回图像数据
	c.Writer.Write(img)
}
