package conf

import (
	"fmt"
	. "github.com/Allen9012/gee_blog/common"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/7/29
  @desc: 配置相关结构体
  @modified by:
**/

type AppConfig struct {
	Name         string `mapstructure:"name"`
	Env          string `mapstructure:"env"`
	Version      string `mapstructure:"version"`
	StartTime    string `mapstructure:"start_time"`
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	MachineID    int64  `mapstructure:"machine_id"`
	*LogConfig   `mapstructure:"logger"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

type LogConfig struct {
	Level string `mapstructure:"level"`
	//prefix       string `mapstructure:"prefix"`
	//ShowLine     bool   `mapstructure:"show_line"` //	是否显示行号
	Directory string `mapstructure:"directory"` //	日志文件目录
	Filename  string `mapstructure:"filename"`
	//LogInConsole bool   `mapstructure:"log_in_console"` //	是否显示打印的路径
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
	Compress   bool   `mapstructure:"compress"`
	LogType    string `mapstructure:"log_type"`
}

type MySQLConfig struct {
	Host        string `mapstructure:"host"`
	Port        int    `mapstructure:"port"`
	Dbname      string `mapstructure:"dbname"`
	User        string `mapstructure:"user"`
	Password    string `mapstructure:"password"`
	LogLevel    string `mapstructure:"log_level"` //	日志级别， debug是全部输出
	MaxOpenConn int    `mapstructure:"max_open_conn"`
	MaxIdleConn int    `mapstructure:"max_idle_conn"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	Db       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

func InitWithViper() (err error) {
	//TODO 设置默认值
	viper.SetDefault("app.Name", "gee-blog")
	viper.SetDefault("app.port", 8080)
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", 3306)
	viper.SetDefault("REDIS_ADDR", "127.0.0.1:6379")
	//	读取配置文件
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	//	将读取的配置信息保存至全局变量Conf
	if err = viper.Unmarshal(&Conf); err != nil {
		zap.L().Error("viper.Unmarshal failed, err : %v", zap.Error(err))
		return
	}
	//	监控配置文件变化
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了...")
		if err = viper.Unmarshal(&Conf); err != nil {
			zap.L().Error("viper.Unmarshal failed, err : %v", zap.Error(err))
			return
		}
	})
	return
}
