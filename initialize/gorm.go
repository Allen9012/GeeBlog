package initialize

import (
	"fmt"
	. "github.com/Allen9012/gee_blog/common"
	"github.com/Allen9012/gee_blog/model"
	"github.com/Allen9012/gee_blog/utils/conf"
	"github.com/Allen9012/gee_blog/utils/logger"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	return InitGorm(GEE_CONFIG.MySQLConfig)
}

func InitGorm(cfg *conf.MySQLConfig) *gorm.DB {
	if cfg == nil {
		panic("[dao mysql Init] invalid config")
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Dbname)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.NewGormLogger(),
		// 默认情况下，GORM会在每个操作上启动一个新事务，如果该操作已经位于事务中，则不会启动新事务。
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic(fmt.Errorf("[dao mysql Init] connect mysql error ", zap.Error(err)))
	}
	//	自动cc迁移
	err = db.AutoMigrate(&model.User{}, &model.Post{})
	if err != nil {
		panic(fmt.Errorf("[dao mysql Init] create table failed ", zap.Error(err)))
	}

	conn, err := db.DB()
	if err != nil {
		panic(fmt.Errorf("[dao mysql Init] get sql instance failed ", zap.Error(err)))
	}
	conn.SetMaxOpenConns(cfg.MaxOpenConn)
	conn.SetMaxIdleConns(cfg.MaxIdleConn)
	return db
}

func Close() {
	conn, err := GEE_DB.DB()
	GEE_LOG.Info("[dao mysql Close] get sql instance failed ", zap.Error(err))
	err = conn.Close()
	GEE_LOG.Info("[dao mysql Close] close the mysql connect failed ", zap.Error(err))
}
