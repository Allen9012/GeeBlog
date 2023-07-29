package main

import (
	. "github.com/Allen9012/gee_blog/common"
	"github.com/Allen9012/gee_blog/config"
	"github.com/spf13/viper"
	"testing"
)

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/7/29
  @desc:
  @modified by:
**/

func TestConf(t *testing.T) {
	// 从配置文件读取配置
	//	设置读取的配置文件路径
	viper.AddConfigPath(".")
	//	设置读取的配置文件
	viper.SetConfigFile("config-dev.yaml")
	err := config.Init()
	if err != nil {
		t.Errorf("config_init config error : %s \n", err)
	}
	t.Log(Conf)
}
