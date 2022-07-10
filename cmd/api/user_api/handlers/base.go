package handlers

import (
	"douyin-Jacob/pkg/errno"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
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


//validator error
func HandleValidatorError(c *gin.Context, err error) {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		c.JSON(http.StatusOK, Response{
			StatusCode: http.StatusOK,
			StatusMsg: err.Error(),
		})
	}
	c.JSON(http.StatusBadRequest, Response{
		StatusCode: http.StatusBadRequest,
		StatusMsg: errs.Error(),
	})
}

