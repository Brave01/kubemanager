package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	SUCCESS = iota
	FAIL
)

func Success(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": SUCCESS,
		"msg":  "成功！",
	})
}

func SuccessWithMessage(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{
		"code": SUCCESS,
		"msg":  message,
	})
}

func SuccessWithDetail(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": SUCCESS,
		"msg":  message,
		"data": data,
	})
}

func Fail(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": FAIL,
		"msg":  "失败",
	})
}

func FailWithMessage(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{
		"code": FAIL,
		"msg":  message,
	})
}

func FailWithDetail(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": FAIL,
		"msg":  message,
		"data": data,
	})
}
