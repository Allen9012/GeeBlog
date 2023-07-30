package main

import (
	"fmt"
	"github.com/Allen9012/gee_blog/config"
	"github.com/Allen9012/gee_blog/router"
	"github.com/spf13/viper"
)

const CONFIGPATH = "."
const CONFIGFILE = "config-dev.yaml"

func main() {
	//	设置读取的配置文件路径
	viper.AddConfigPath(CONFIGPATH)
	//	设置读取的配置文件
	viper.SetConfigFile(CONFIGFILE)
	// 初始化配置
	err := config.Init()
	if err != nil {
		fmt.Printf("config_init config error : %s \n", err)
		return
	}
	// gin-swagger middleware
	//docs.SwaggerInfo.BasePath = "/api/v1"
	// 装载路由
	r := router.NewRouter()

	r.Run(":3000")
}
