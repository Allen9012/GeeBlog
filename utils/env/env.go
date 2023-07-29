package env

import (
	. "github.com/Allen9012/gee_blog/common"
)

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/7/29
  @desc: 区别本地和生成环境
  @modified by:
**/

func IsDev() bool {
	return Conf.Env == "dev"
}

func IsProd() bool {
	return Conf.Env == "prod"
}

func IsTesting() bool {
	return Conf.Env == "test"
}
