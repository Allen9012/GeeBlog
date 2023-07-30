package config

import (
	"fmt"
	. "github.com/Allen9012/gee_blog/common"
	"github.com/Allen9012/gee_blog/dao/mysql"
	"github.com/Allen9012/gee_blog/dao/redis"
	"github.com/Allen9012/gee_blog/utils"
	"github.com/Allen9012/gee_blog/utils/logger"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/7/29
  @desc: 执行配置的初始化包括logger&&mysql&&redis&&雪花算法
  @modified by:
**/

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

// Init 初始化配置项
func Init() error {
	//	使用viper读取配置文件
	err := InitWithViper()
	if err != nil {
		return err
	}

	// 设置日志级别
	if err = logger.Init(Conf.LogConfig, Conf.Env); err != nil {
		fmt.Printf("config_init logger error : %s \n", err)
		return err
	}
	defer logger.Sync()
	// 读取翻译文件
	if err := LoadLocales("config/locales/zh-cn.yaml"); err != nil {
		// zap输出错误日志使用到err
		logger.Panic("翻译文件加载失败 %v", zap.Error(err))
	}

	// 连接数据库
	//model.Database(os.Getenv("MYSQL_DSN"))
	//cache.Redis()

	//初始化数据库
	if err := mysql.Init(Conf.MySQLConfig); err != nil {
		logger.Panic("config_init mysql error : %s \n", zap.Error(err))
		return err
	}
	defer mysql.Close()
	//初始化redis
	if err := redis.Init(Conf.RedisConfig); err != nil {
		logger.Panic("config_init redis error : %s \n", zap.Error(err))
		return err
	}
	defer redis.Close()
	// 初始化雪花
	if err := utils.Init(Conf.StartTime, Conf.MachineID); err != nil {
		logger.Panic("config_init snowflake error : %s \n", zap.Error(err))
		return err
	}

	return nil

}
