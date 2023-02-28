package resp

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	ParamMiss  = 500
	TokenMiss  = -1000
	RedisError = 502
)

type response struct {
	ErrCode int         `json:"err_code"`
	ErrMsg  string      `json:"err_msg"`
	Data    interface{} `json:"data"`
}

func OK(c *gin.Context, data ...interface{}) {
	r := response{
		ErrCode: 0,
		ErrMsg:  "",
		Data:    nil,
	}
	for _, value := range data {
		switch value.(type) {
		case string:
			r.ErrMsg = value.(string)
		case int:
			r.ErrCode = value.(int)
		case interface{}:
			r.Data = value.(interface{})
		}
	}

	c.JSON(http.StatusOK, r)
	return
}

func ParamError(c *gin.Context, data ...interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"err_code": ParamMiss,
		"err_msg":  "参数绑定异常",
		"data":     data,
	})

	return
}

func Error(c *gin.Context, data ...interface{}) {
	r := response{
		ErrCode: -1,
		ErrMsg:  "操作失败",
		Data:    nil,
	}
	for _, value := range data {
		switch value.(type) {
		case int:
			r.ErrCode = value.(int)
		case string:
			r.ErrMsg = value.(string)
		case interface{}:
			r.Data = value.(interface{})
		}
	}

	c.JSON(http.StatusOK, r)
}

// Success 正确返回
func Success(c *gin.Context, msg ...string) {
	r := response{
		ErrCode: 0,
		ErrMsg:  "",
	}
	if len(msg) > 0 {
		r.Data = map[string]string{"msg": msg[0]}
	}
	fmt.Printf("%+v", r)

	c.JSON(http.StatusOK, r)
}
