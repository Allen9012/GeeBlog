package controller

import (
	"errors"
	"github.com/Allen9012/gee_blog/utils"
	"github.com/gin-gonic/gin"
)

var ErrorIdNotExist = errors.New("用户不可用")

func getUserId(c *gin.Context) (int64, error) {
	value, exist := c.Get(utils.KeyUserId)
	if !exist {
		return -1, ErrorIdNotExist
	}
	userId, ok := value.(int64)
	if !ok {
		return -1, ErrorIdNotExist
	}
	return userId, nil
}
