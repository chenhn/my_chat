package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	respOk      = "OK"
	respFail    = "FAIL"
	respUnLogin = "UNLOGIN"
	respNoAuth  = "NOAUTH"
)

// resp Payload返回
type resp map[string]interface{}

func ok(c *gin.Context, r resp) {
	result := make(map[string]interface{})
	for key, value := range r {
		if fmt.Sprint(value) != "<nil>" {
			result[key] = value
		}
	}
	result["resultCode"] = respOk
	result["resultMsg"] = ""
	c.JSON(http.StatusOK, result)
}

// unLogin 未登录情况返回
func unLogin(c *gin.Context) {
	c.JSON(http.StatusOK, resp{
		"resultCode": respUnLogin,
		"resultMsg":  "未登录或登录失效",
	})
}

// noAuth 无权限
func noAuth(c *gin.Context) {
	c.JSON(http.StatusOK, resp{
		"resultCode": respNoAuth,
		"resultMsg":  "无权限",
	})
}

// fail 错误情况返回
func fail(c *gin.Context, e error) {
	c.JSON(http.StatusOK, resp{
		"resultCode": respFail,
		"resultMsg":  e.Error(),
	})
}
