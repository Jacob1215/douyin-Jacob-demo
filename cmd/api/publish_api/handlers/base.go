package handlers

import (
	"douyin-Jacob/cmd/api/publish_api/global"
	"douyin-Jacob/pkg/jwt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"net/http"
)

var (Jwt *jwt.JWT)

type Response struct {
	StatusCode int64 `json:"status_code"`
	StatusMsg string `json:"status_msg"`
	Data interface{} `json:"data"`
}

func InitJwt()  {
	zap.S().Info(global.ServerConfig.JWTInfo.SigningKey)
	Jwt = jwt.NewJWT(global.ServerConfig.JWTInfo.SigningKey)
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
	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg: "get user info success",
		Data: data,
	})
}


