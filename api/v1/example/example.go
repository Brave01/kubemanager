package example

import (
	"github.com/gin-gonic/gin"
	"kubemanager/response"
)

type Example struct{}

func (e *Example) PingUrl(ctx *gin.Context) {
	response.SuccessWithDetail(ctx, "成功", map[string]string{"url": "http://www.google.com"})
}
