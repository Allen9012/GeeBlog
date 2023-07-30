package conf

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/7/29
  @desc: 配置相关结构体
  @modified by:
**/

type AppConfig struct {
	Name         string `mapstructure:"name"`
	Env          string `mapstructure:"env"`
	Version      string `mapstructure:"version"`
	StartTime    string `mapstructure:"start_time"`
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	MachineID    int64  `mapstructure:"machine_id"`
	*LogConfig   `mapstructure:"logger"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}
type LogConfig struct {
	Level string `mapstructure:"level"`
	//prefix       string `mapstructure:"prefix"`
	//ShowLine     bool   `mapstructure:"show_line"` //	是否显示行号
	Directory string `mapstructure:"directory"` //	日志文件目录
	Filename  string `mapstructure:"filename"`
	//LogInConsole bool   `mapstructure:"log_in_console"` //	是否显示打印的路径
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
	Compress   bool   `mapstructure:"compress"`
	LogType    string `mapstructure:"log_type"`
}

type MySQLConfig struct {
	Host        string `mapstructure:"host"`
	Port        int    `mapstructure:"port"`
	Dbname      string `mapstructure:"dbname"`
	User        string `mapstructure:"user"`
	Password    string `mapstructure:"password"`
	LogLevel    string `mapstructure:"log_level"` //	日志级别， debug是全部输出
	MaxOpenConn int    `mapstructure:"max_open_conn"`
	MaxIdleConn int    `mapstructure:"max_idle_conn"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	Db       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}
