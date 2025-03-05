package initiallize

import (
	"fmt"
	"github.com/spf13/viper"
	"kubemanager/global"
)

func ViperConfig() {
	v := viper.New()
	v.SetConfigFile("config.yaml")
	v.SetConfigType("yaml")
	err2 := v.ReadInConfig()
	if err2 != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err2))
	}
	err := v.Unmarshal(global.CONF)
	if err != nil {
		panic(err)
	}
}
