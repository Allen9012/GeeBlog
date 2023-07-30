package common

import (
	"github.com/Allen9012/gee_blog/utils/conf"
)

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/7/29
  @desc: 主要放置一些Global的变量
  @modified by:
**/

var (
	Conf = new(conf.AppConfig)
)

//var (
//	DB      *gorm.DB
//	ELASTIC *elasticsearch.Client
//	REIDS   *redis.Client
//	VP      *viper.Viper
//	CONFIG  config.Server
//	I18N    *i18n.Bundle
//	LOG     *zap.Logger
//	QUEST   *dataframe.DataFrame
//)
