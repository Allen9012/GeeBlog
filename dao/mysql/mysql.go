package mysql

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

//Host        string
//Port        int
//Dbname      string
//User        string
//Password    string
//LogLevel    string
//MaxOpenConn int
//MaxIdleConn int

var db *gorm.DB

func Init(cfg *conf.MySQLConfig) (err error) {
	if cfg == nil {
		zap.L().Error("[dao mysql Init] invalid config")
		return
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Dbname)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.NewGormLogger(),
		// 默认情况下，GORM会在每个操作上启动一个新事务，如果该操作已经位于事务中，则不会启动新事务。
		SkipDefaultTransaction: true,
	})
	if err != nil {
		zap.L().Error("[dao mysql Init] connect mysql error ", zap.Error(err))
		return
	}
	//	自动cc迁移
	err = db.AutoMigrate(&model.User{}, &model.Post{})
	if err != nil {
		zap.L().Info("[dao mysql Init] create table failed ", zap.Error(err))
		return err
	}

	conn, err := db.DB()
	if err != nil {
		zap.L().Info("[dao mysql Init] get sql instance failed ", zap.Error(err))
		return err
	}
	conn.SetMaxOpenConns(cfg.MaxOpenConn)
	conn.SetMaxIdleConns(cfg.MaxIdleConn)
	return
}

func Close() {
	conn, err := db.DB()
	zap.L().Info("[dao mysql Close] get sql instance failed ", zap.Error(err))
	err = conn.Close()
	zap.L().Info("[dao mysql Close] close the mysql connect failed ", zap.Error(err))
}
