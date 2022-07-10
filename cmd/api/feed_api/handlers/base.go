package handlers

import (
	"douyin-Jacob/pkg/errno"
	"github.com/gin-gonic/gin"

)


type Response struct {
	HttpCode int `json:"Http_code"`
	StatusCode int `json:"status_code"`
	StatusMsg string `json:"status_msg"`
}

func SendHttpResponse(err errno.HttpErr,c *gin.Context)  {
	c.JSON(err.StatusCode,Response{
		StatusCode: err.Errno.Code,
		StatusMsg: err.Errno.Message,
	})
}
