package config

import (
	"fmt"
	. "github.com/Allen9012/gee_blog/common"
	"github.com/Allen9012/gee_blog/dao/mysql"
	"github.com/Allen9012/gee_blog/dao/redis"
	"github.com/Allen9012/gee_blog/utils"
	"github.com/Allen9012/gee_blog/utils/conf"
	"github.com/Allen9012/gee_blog/utils/logger"
	"go.uber.org/zap"
)

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/7/29
  @desc: 执行配置的初始化包括logger&&mysql&&redis&&雪花算法
  @modified by:
**/

// Init 初始化配置项
func Init() error {
	//	使用viper读取配置文件
	err := conf.InitWithViper()
	if err != nil {
		return err
	}

	// 设置日志级别
	if err = logger.Init(Conf.LogConfig, Conf.Env); err != nil {
		fmt.Printf("config_init logger error : %s \n", err)
		return err
	}
	defer zap.L().Sync()
	// 读取翻译文件
	if err := LoadLocales("config/locales/zh-cn.yaml"); err != nil {
		// zap输出错误日志使用到err
		zap.L().Panic("翻译文件加载失败 %v", zap.Error(err))
	}

	// 连接数据库
	//model.Database(os.Getenv("MYSQL_DSN"))
	//cache.Redis()

	//初始化数据库
	if err := mysql.Init(Conf.MySQLConfig); err != nil {
		zap.L().Panic("config_init mysql error : %s \n", zap.Error(err))
		return err
	}
	defer mysql.Close()
	//初始化redis
	if err := redis.Init(Conf.RedisConfig); err != nil {
		zap.L().Panic("config_init redis error : %s \n", zap.Error(err))
		return err
	}
	defer redis.Close()
	// 初始化雪花
	if err := utils.Init(Conf.StartTime, Conf.MachineID); err != nil {
		zap.L().Panic("config_init snowflake error : %s \n", zap.Error(err))
		return err
	}

	return nil

}
