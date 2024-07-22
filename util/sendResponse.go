package util

import "github.com/gin-gonic/gin"

type ResponseDto struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func SendSuccessRes(c *gin.Context, code int, message string, data any) {
	c.JSON(code, ResponseDto{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

func SendErrorRes(c *gin.Context, code int, message string) {
	c.JSON(code, ResponseDto{
		Code:    code,
		Message: message,
		Data:    nil,
	})
}
