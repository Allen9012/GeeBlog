package initialize

import (
	"context"
	"fmt"
	"github.com/Allen9012/gee_blog/common"
	"github.com/Allen9012/gee_blog/utils/conf"
	"github.com/redis/go-redis/v9"
)

func RedisCtx() context.Context {
	return context.Background()
}

func Redis() *redis.Client {
	return InitRedis(common.GEE_CONFIG.RedisConfig)
}
func InitRedis(cfg *conf.RedisConfig) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password, // 密码
		DB:       cfg.Db,       // 数据库
		PoolSize: cfg.PoolSize, // 连接池大小
	})
	return client
}
