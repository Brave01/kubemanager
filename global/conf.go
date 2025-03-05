package global

import (
	"k8s.io/client-go/kubernetes"
	"kubemanager/config"
)

var (
	CONF          = new(config.Config)
	KubeConfigSet = new(kubernetes.Clientset)
)
