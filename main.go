package main

import (
	"kubemanager/global"
	"kubemanager/initiallize"
)

func main() {
	initiallize.ViperConfig()
	router := initiallize.Router()
	initiallize.K8S()
	panic(router.Run(global.CONF.System.Addr))
}
