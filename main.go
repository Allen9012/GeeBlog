package main

import (
	"fmt"
	"github.com/Allen9012/gee_blog/common"
	"github.com/Allen9012/gee_blog/initialize"
	"github.com/Allen9012/gee_blog/router"
	"github.com/Allen9012/gee_blog/utils/conf"
	"go.uber.org/zap"
)

// // Viper 初始化配置项
//
//	func init() error {
//		//	使用viper读取配置文件
//		err := Viper()
//		if err != nil {
//			return err
//		}
//
//		// 设置日志级别
//		if err = initialize.Init(Conf.LogConfig, Conf.Env); err != nil {
//			fmt.Printf("config_init logger error : %s \n", err)
//			return err
//		}
//		defer initialize.Sync()
//		// 读取翻译文件
//		if err := LoadLocales("config/locales/zh-cn.yaml"); err != nil {
//			// zap输出错误日志使用到err
//			initialize.Panic("翻译文件加载失败 %v", zap.Error(err))
//		}
//
//		// 连接数据库
//		//model.Database(os.Getenv("MYSQL_DSN"))
//		//cache.Redis()
//
//		//初始化数据库
//		if err := mysql.Init(Conf.MySQLConfig); err != nil {
//			initialize.Panic("config_init mysql error : %s \n", zap.Error(err))
//			return err
//		}
//		defer mysql.Close()
//		//初始化redis
//		if err := redis.Init(Conf.RedisConfig); err != nil {
//			initialize.Panic("config_init redis error : %s \n", zap.Error(err))
//			return err
//		}
//		defer redis.Close()
//		// 初始化雪花
//		if err := utils.Init(Conf.StartTime, Conf.MachineID); err != nil {
//			initialize.Panic("config_init snowflake error : %s \n", zap.Error(err))
//			return err
//		}
//
//		return nil
//
// }
func init() {
	common.GEE_CONFIG = new(conf.AppConfig)
	common.GEE_VP = initialize.Viper()
	common.GEE_DB = initialize.Gorm()
	//GEE_ELASTIC *elasticsearch.Client
	common.GEE_REDIS = initialize.Redis()
	//GEE_I18N    *i18n.Bundle
	common.GEE_LOG = initialize.Zap()
	//GEE_QUEST   *dataframe.DataFrame
	common.GEE_REDIS_CTX = initialize.RedisCtx()
}

func main() {
	if common.GEE_DB != nil {
		db, _ := common.GEE_DB.DB()
		defer db.Close()
	}
	if common.GEE_REDIS != nil {
		common.GEE_LOG.Info("[main] redis closed")
		defer common.GEE_REDIS.Close()
	}
	// 检查redis ctx是否初始化成功
	if _, err := common.GEE_REDIS.Ping(common.GEE_REDIS_CTX).Result(); err != nil {
		panic(fmt.Errorf("[dao redis Init] ping redis client failed: %v", zap.Error(err)))
	}
	// gin-swagger middleware
	//docs.SwaggerInfo.BasePath = "/api/v1"
	// 装载路由
	r := router.NewRouter()

	r.Run(":3000")
}
