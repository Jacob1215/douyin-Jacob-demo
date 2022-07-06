package handlers

import (
"github.com/gin-gonic/gin"
"github.com/go-playground/validator/v10"
"google.golang.org/grpc/codes"
"google.golang.org/grpc/status"

"net/http"
)


type Response struct {
	StatusCode int64 `json:"status_code"`
	StatusMsg string `json:"status_msg"`
	Data interface{} `json:"data"`
}

//SendResponseToHttp send the success and errors to http
func SendResponseToHttp(err error, c *gin.Context,data interface{}) {
	//将grpc的code转换成http的状态码
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				c.JSON(http.StatusNotFound, Response{
					StatusCode: http.StatusNotFound,
					StatusMsg: e.Message(),
					Data: data,
				})
			case codes.Internal:
				c.JSON(http.StatusInternalServerError, Response{
					StatusCode: http.StatusInternalServerError,
					StatusMsg: e.Message(),
				})
			case codes.InvalidArgument:
				c.JSON(http.StatusBadRequest, Response{
					StatusCode: http.StatusBadRequest,
					StatusMsg: e.Message(),
				})
			case codes.Unavailable:
				c.JSON(http.StatusInternalServerError, Response{
					StatusCode: http.StatusInternalServerError,
					StatusMsg: e.Message(),
				})

			default:
				c.JSON(http.StatusInternalServerError, Response{
					StatusCode: http.StatusInternalServerError,
					StatusMsg: e.Message(),
				})
			}
			return
		}
	}
	c.JSON(http.StatusOK,Response{
		StatusCode: 0,
		StatusMsg: "get user info success",
		Data: data,
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

