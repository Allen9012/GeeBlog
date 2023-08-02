package initialize

import (
	"fmt"
	. "github.com/Allen9012/gee_blog/common"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

const CONFIGPATH = "."
const CONFIGFILE = "config-dev.yaml"

func Viper() *viper.Viper {
	v := viper.New()
	//	设置读取的配置文件路径
	v.AddConfigPath(CONFIGPATH)
	//	设置读取的配置文件
	v.SetConfigFile(CONFIGFILE)
	//TODO 设置默认值
	viper.SetDefault("app.Name", "gee-blog")
	viper.SetDefault("app.port", 8080)
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", 3306)
	viper.SetDefault("REDIS_ADDR", "127.0.0.1:6379")
	//	读取配置文件
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	//	将读取的配置信息保存至全局变量Conf
	if err = v.Unmarshal(GEE_CONFIG); err != nil {
		panic(fmt.Errorf("viper.Unmarshal failed, err : %v", zap.Error(err)))
	}
	//	监控配置文件变化
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了...")
		if err = v.Unmarshal(GEE_CONFIG); err != nil {
			panic(fmt.Errorf("viper.Unmarshal failed, err : %v", zap.Error(err)))
		}
	})
	return v
}
