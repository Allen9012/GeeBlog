package redis

import (
	"context"
	"fmt"
	"github.com/Allen9012/gee_blog/utils/conf"
	"github.com/redis/go-redis/v9"
	"time"

	"go.uber.org/zap"
)

var (
	client *redis.Client
	ctx    context.Context
)

const (
	TokenPrefix  = "login:token:"
	TokenTimeout = time.Hour * 24
)

func Init(cfg *conf.RedisConfig) (err error) {
	client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password, // 密码
		DB:       cfg.Db,       // 数据库
		PoolSize: cfg.PoolSize, // 连接池大小
	})

	ctx = context.Background()
	_, err = client.Ping(ctx).Result()
	zap.L().Info("[dao redis Init] ping redis client failed ", zap.Error(err))
	return
}

func Close() {
	err := client.Close()
	zap.L().Info("[dao redis Close] close the redis connect failed ", zap.Error(err))
}
