package initiallize

import (
	"github.com/gin-gonic/gin"
	middleware "kubemanager/middlerware"
	"kubemanager/router"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())
	router.RouterGroupApp.ExampleRoute.PingRoute(r)
	k8sGroup := router.RouterGroupApp.K8sRouterGroup
	{
		// podList
		k8sGroup.PodList(r)
	}
	return r
}
