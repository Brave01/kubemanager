package example

import (
	"github.com/gin-gonic/gin"
	"kubemanager/api"
)

type Example struct {
}

func (e *Example) PingRoute(engine *gin.Engine) {
	group := engine.Group("/example")
	group.GET("/ping", api.ApiGroupApp.ExampleApi.PingUrl)
}
