package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	ERROR   = 7
	SUCCESS = 0
)

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func Result(code int, data any, msg string, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func OK(data any, msg string, c *gin.Context) {
	Result(SUCCESS, data, msg, c)
}

func Okwith(c *gin.Context) {
	Result(SUCCESS, map[string]any{}, "成功", c)
}

func OkwithData(data any, c *gin.Context) {
	Result(SUCCESS, data, "成功", c)
}

func OkwithMessage(messge string, c *gin.Context) {
	Result(SUCCESS, map[string]any{}, messge, c)
}

func Fail(data any, msg string, c *gin.Context) {
	Result(ERROR, data, msg, c)
}

func FailwithMessage(messge string, c *gin.Context) {
	Result(ERROR, map[string]any{}, messge, c)
}

func FailwithCode(code ErrorCode, c *gin.Context) {
	msg, ok := ErrorMap[code]
	if ok {
		Result(int(code), map[string]any{}, msg, c)
		return
	}
	Result(ERROR, map[string]any{}, "未知错误", c)
}
