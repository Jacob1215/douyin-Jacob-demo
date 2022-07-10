package handlers

import (
	"douyin-Jacob/cmd/api/favorite_api/global"
	"douyin-Jacob/pkg/errno"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"
)


type Response struct {
	HttpCode int `json:"Http_code"`
	StatusCode int `json:"status_code"`
	StatusMsg string `json:"status_msg"`
}

func SendHttpResponse(err errno.HttpErr,c *gin.Context)  {
	c.JSON(err.StatusCode,Response{
		HttpCode: err.StatusCode,
		StatusCode: err.Errno.Code,
		StatusMsg: err.Errno.Message,
	})
}

//处理validator的错误
func HandleValidatorError(c *gin.Context, err error) {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"error": removeTopStruct(errs.Translate(global.Trans)),
	})
}
func removeTopStruct(fields map[string]string) map[string]string {
	rsp := map[string]string{}
	for field, err := range fields {
		rsp[field[strings.Index(field, ".")+1:]] = err
	}
	return rsp
}
