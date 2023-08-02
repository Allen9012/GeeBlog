package redis

import (
	"errors"
	. "github.com/Allen9012/gee_blog/common"
	"go.uber.org/zap"
)

var (
	ErrorTokenNotExist = errors.New("Token 不存在")
)

// InsertTokenByUserId
func InsertTokenByUserId(token string, userId int64, userRole uint8) (err error) {
	// 使用 pipeline 减少 RTT
	pipeline := GEE_REDIS.TxPipeline()

	// 把 token 插入到 redis中
	key := TokenPrefix + token
	pipeline.HSet(GEE_REDIS_CTX, key, KeyUserId, userId, KeyUserRole, userRole)
	// 为 token 设置过期时间
	pipeline.Expire(GEE_REDIS_CTX, key, TokenTimeout)

	// 执行 pipeline
	_, err = pipeline.Exec(GEE_REDIS_CTX)

	return
}

// RefreshToken
func RefreshToken(token string) {
	key := TokenPrefix + token

	err := GEE_REDIS.HMGet(GEE_REDIS_CTX, key, KeyUserId, KeyUserRole).Err()
	if err != nil {
		zap.L().Error("[middleware token] GEE_REDIS hmget key ", zap.Error(err))
		return
	}

	err = GEE_REDIS.Expire(GEE_REDIS_CTX, key, TokenTimeout).Err()
	if err != nil {
		zap.L().Error("[middleware token] GEE_REDIS expire key ", zap.Error(err))
	}
	return
}

// CheckTokenExist
func CheckTokenExist(token string) ([]interface{}, error) {
	key := TokenPrefix + token
	res, err := GEE_REDIS.HMGet(GEE_REDIS_CTX, key, KeyUserId, KeyUserRole).Result()
	if err != nil {
		zap.L().Error("[middleware token] GEE_REDIS hmget key ", zap.Error(err))
		return nil, err
	}
	if res == nil {
		return nil, ErrorTokenNotExist
	}
	return res, nil
}

// DeleteToken
func DeleteToken(token string) error {
	return GEE_REDIS.HDel(GEE_REDIS_CTX, TokenPrefix+token, KeyUserId, KeyUserRole).Err()
}
