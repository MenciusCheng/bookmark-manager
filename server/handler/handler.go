package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ping(c *gin.Context) {
	JSON(c, "pong")
}

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data,omitempty"`
}

func JSON(c *gin.Context, v any) {
	c.JSON(http.StatusOK, Response{Data: v})
}

func JSONErr(c *gin.Context, code ErrCode, v ...any) {
	msg := ErrMsg[code]
	if len(v) > 0 {
		msg = fmt.Sprint(append([]any{msg, ":"}, v...)...)
	}

	c.JSON(http.StatusOK, Response{Code: int(code), Msg: msg})
}

type ErrCode int

const (
	ErrCodeParam    = 400
	ErrCodeInternal = 500
)

var ErrMsg = map[ErrCode]string{
	ErrCodeParam:    "param error",
	ErrCodeInternal: "internal server error",
}
