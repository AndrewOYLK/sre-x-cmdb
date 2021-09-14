package utils

import "github.com/gin-gonic/gin"

// APIError API接口错误返回
type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// APISuccess API接口正确返回
type APISuccess struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// NewAPIError Gin API错误返回
func NewAPIError(ctx *gin.Context, status int, err error) {
	e := APIError{
		Code:    1,
		Message: err.Error(),
	}
	ctx.JSON(status, e)
}

// NewAPISuccess Gin API成功返回
func NewAPISuccess(ctx *gin.Context, status int, message string, data interface{}) {
	s := APISuccess{
		Code:    0,
		Message: message,
		Data:    data,
	}
	ctx.JSON(status, s)
}
