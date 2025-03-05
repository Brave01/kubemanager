package k8s

import (
	"fmt"
	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubemanager/global"
)

type PodListApi struct{}

func (p *PodListApi) PodList(c *gin.Context) {
	list, err := global.KubeConfigSet.CoreV1().Pods("").List(c, metav1.ListOptions{})
	if err != nil {
		fmt.Println(err)
	}
	for _, item := range list.Items {
		fmt.Println(item.Namespace, item.Name)
	}
	c.JSON(200, gin.H{"message": "pong"})
}
