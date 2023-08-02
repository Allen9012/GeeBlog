package common

import (
	"context"
	"github.com/Allen9012/gee_blog/utils/conf"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/7/29
  @desc: 主要放置一些Global的变量
  @modified by:
**/

var (
	GEE_DB *gorm.DB
	//GEE_ELASTIC *elasticsearch.Client
	GEE_REDIS  *redis.Client
	GEE_VP     *viper.Viper
	GEE_CONFIG *conf.AppConfig
	//GEE_I18N    *i18n.Bundle
	GEE_LOG *zap.Logger
	//GEE_QUEST   *dataframe.DataFrame
	GEE_REDIS_CTX context.Context
)
