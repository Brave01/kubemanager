package k8s

import (
	"github.com/gin-gonic/gin"
	"kubemanager/api"
)

type PodListRoute struct{}

func (p *PodListRoute) PodList(engine *gin.Engine) {
	group := engine.Group("/k8s")
	k8sApi := api.ApiGroupApp.K8sApiGroup
	group.GET("/podList", k8sApi.PodList)
}
