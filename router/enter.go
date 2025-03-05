package router

import (
	"kubemanager/router/example"
	"kubemanager/router/k8s"
)

type RouterGroup struct {
	ExampleRoute   example.ExampleRoute
	K8sRouterGroup k8s.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
