package api

import (
	"kubemanager/api/v1/example"
	"kubemanager/api/v1/k8s"
)

type ApiGroup struct {
	ExampleApi  example.ExampleApi
	K8sApiGroup k8s.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
