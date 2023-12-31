package initialize

import (
	"encoding/json"
	"fmt"
	. "github.com/Allen9012/gee_blog/common"
	"github.com/Allen9012/gee_blog/utils/conf"
	"github.com/Allen9012/gee_blog/utils/env"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
	"time"
)

func Zap() *zap.Logger {
	var lg *zap.Logger
	var err error
	if lg, err = InitZap(GEE_CONFIG.LogConfig, GEE_CONFIG.Env); err != nil {
		panic(fmt.Errorf("config_init zap error : %s \n", zap.Error(err)))
	}
	return lg
}

// Zap 初始化lg
func InitZap(cfg *conf.LogConfig, mode string) (*zap.Logger, error) {
	// 配置lumberjack
	writeSyncer := getLogWriter(
		cfg.Filename,
		cfg.MaxSize,
		cfg.MaxBackups,
		cfg.MaxAge,
		cfg.Compress,
		cfg.LogType,
	)
	logLevel := new(zapcore.Level)
	// 将配置文件中的日志级别转换成zapcore.Level类型
	err := logLevel.UnmarshalText([]byte(cfg.Level))
	if err != nil {
		fmt.Println("日志级别设置错误")
		return nil, err
	}
	// 设置日志编码格式
	var core zapcore.Core
	if mode == "dev" {
		// 进入开发模式，日志输出到终端
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		core = zapcore.NewTee(
			// 生产模式
			zapcore.NewCore(getEncoder(), writeSyncer, logLevel),
			// dev模式
			zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel),
		)
	} else {
		core = zapcore.NewCore(getEncoder(), writeSyncer, logLevel)
	}

	lg := zap.New(core,
		zap.AddCaller(),
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zap.ErrorLevel))

	zap.ReplaceGlobals(lg)
	lg.Info("config_init logger success")
	return lg, nil
}

// getEncoder
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeTime = CustomTimeEncoder
	//encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

// 自定义日志输出时间格式
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006/01/02 - 15:04:05.000"))
}

// getLogWriter
func getLogWriter(filename string, maxSize, maxBackup, maxAge int, compress bool, logType string) zapcore.WriteSyncer {
	// 改成时间格式的log
	if logType == "daily" {
		logname := time.Now().Format("2006-01-02.log")
		filename = strings.ReplaceAll(filename, "gee.log", logname)
	}
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
		Compress:   compress,
	}
	if env.IsDev() {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
	} else {
		return zapcore.AddSync(lumberJackLogger)
	}
}

/* 辅助logger函数 */

// Dump 调试专用，不会中断程序，会在终端打印出 warning 消息。
// 第一个参数会使用 json.Marshal 进行渲染，第二个参数消息（可选）
//
//	Logger.Dump(user.User{Name:"test"})
//	Logger.Dump(user.User{Name:"test"}, "用户信息")
func Dump(value interface{}, msg ...string) {
	valueString := jsonString(value)
	// 判断第二个参数是否传参 msg
	if len(msg) > 0 {
		GEE_LOG.Warn("Dump", zap.String(msg[0], valueString))
	} else {
		GEE_LOG.Warn("Dump", zap.String("data", valueString))
	}
}

// LogIf 当 err != nil 时记录 error 等级的日志
func LogIf(err error) {
	if err != nil {
		GEE_LOG.Error("Error Occurred:", zap.Error(err))
	}
}

// LogWarnIf 当 err != nil 时记录 warning 等级的日志
func LogWarnIf(err error) {
	if err != nil {
		GEE_LOG.Warn("Error Occurred:", zap.Error(err))
	}
}

// LogInfoIf 当 err != nil 时记录 info 等级的日志
func LogInfoIf(err error) {
	if err != nil {
		GEE_LOG.Info("Error Occurred:", zap.Error(err))
	}
}
func Sync() {
	GEE_LOG.Sync()
}

// Debug 调试日志，详尽的程序日志
// 调用示例：
//
//	logger.Debug("Database", zap.String("sql", sql))
func Debug(moduleName string, fields ...zap.Field) {
	GEE_LOG.Debug(moduleName, fields...)
}

// Info 告知类日志
func Info(moduleName string, fields ...zap.Field) {
	GEE_LOG.Info(moduleName, fields...)
}

// Warn 警告类
func Warn(moduleName string, fields ...zap.Field) {
	GEE_LOG.Warn(moduleName, fields...)
}

// Error 错误时记录，不应该中断程序，查看日志时重点关注
func Error(moduleName string, fields ...zap.Field) {
	GEE_LOG.Error(moduleName, fields...)
}

// Fatal 级别同 Error(), 写完 log 后调用 os.Exit(1) 退出程序
func Fatal(moduleName string, fields ...zap.Field) {
	GEE_LOG.Fatal(moduleName, fields...)
}
func Panic(moduleName string, fields ...zap.Field) {
	GEE_LOG.Panic(moduleName, fields...)
}

// DebugString 记录一条字符串类型的 debug 日志，调用示例：
//
//	logger.DebugString("SMS", "短信内容", string(result.RawResponse))
func DebugString(moduleName, name, msg string) {
	GEE_LOG.Debug(moduleName, zap.String(name, msg))
}

func InfoString(moduleName, name, msg string) {
	GEE_LOG.Info(moduleName, zap.String(name, msg))
}

func WarnString(moduleName, name, msg string) {
	GEE_LOG.Warn(moduleName, zap.String(name, msg))
}

func ErrorString(moduleName, name, msg string) {
	GEE_LOG.Error(moduleName, zap.String(name, msg))
}

func FatalString(moduleName, name, msg string) {
	GEE_LOG.Fatal(moduleName, zap.String(name, msg))
}

// DebugJSON 记录对象类型的 debug 日志，使用 json.Marshal 进行编码。调用示例：
//
//	logger.DebugJSON("Auth", "读取登录用户", auth.CurrentUser())
func DebugJSON(moduleName, name string, value interface{}) {
	GEE_LOG.Debug(moduleName, zap.String(name, jsonString(value)))
}

func InfoJSON(moduleName, name string, value interface{}) {
	GEE_LOG.Info(moduleName, zap.String(name, jsonString(value)))
}

func WarnJSON(moduleName, name string, value interface{}) {
	GEE_LOG.Warn(moduleName, zap.String(name, jsonString(value)))
}

func ErrorJSON(moduleName, name string, value interface{}) {
	GEE_LOG.Error(moduleName, zap.String(name, jsonString(value)))
}

func FatalJSON(moduleName, name string, value interface{}) {
	GEE_LOG.Fatal(moduleName, zap.String(name, jsonString(value)))
}

func jsonString(value interface{}) string {
	b, err := json.Marshal(value)
	if err != nil {
		GEE_LOG.Error("Logger", zap.String("JSON marshal error", err.Error()))
	}
	return string(b)
}
