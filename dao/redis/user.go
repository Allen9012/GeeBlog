package redis

import (
	"errors"
	"github.com/Allen9012/gee_blog/common"
	"go.uber.org/zap"
)

var (
	ErrorTokenNotExist = errors.New("Token 不存在")
)

// InsertTokenByUserId
func InsertTokenByUserId(token string, userId int64, userRole uint8) (err error) {
	// 使用 pipeline 减少 RTT
	pipeline := client.TxPipeline()

	// 把 token 插入到 redis中
	key := TokenPrefix + token
	pipeline.HSet(ctx, key, common.KeyUserId, userId, common.KeyUserRole, userRole)
	// 为 token 设置过期时间
	pipeline.Expire(ctx, key, TokenTimeout)

	// 执行 pipeline
	_, err = pipeline.Exec(ctx)

	return
}

// RefreshToken
func RefreshToken(token string) {
	key := TokenPrefix + token

	err := client.HMGet(ctx, key, common.KeyUserId, common.KeyUserRole).Err()
	if err != nil {
		zap.L().Error("[middleware token] client hmget key ", zap.Error(err))
		return
	}

	err = client.Expire(ctx, key, TokenTimeout).Err()
	if err != nil {
		zap.L().Error("[middleware token] client expire key ", zap.Error(err))
	}
	return
}

// CheckTokenExist
func CheckTokenExist(token string) ([]interface{}, error) {
	key := TokenPrefix + token
	res, err := client.HMGet(ctx, key, common.KeyUserId, common.KeyUserRole).Result()
	if err != nil {
		zap.L().Error("[middleware token] client hmget key ", zap.Error(err))
		return nil, err
	}
	if res == nil {
		return nil, ErrorTokenNotExist
	}
	return res, nil
}

// DeleteToken
func DeleteToken(token string) error {
	return client.HDel(ctx, TokenPrefix+token, common.KeyUserId, common.KeyUserRole).Err()
}
