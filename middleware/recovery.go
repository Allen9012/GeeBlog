package middleware

import (
	"github.com/Allen9012/gee_blog/initialize"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
	"time"
)

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/7/30
  @desc: recovery 中间件
  @modified by:
**/

// GinRecovery recover掉项目可能出现的panic，并使用zap记录相关日志
func GinRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				// 判断是否为网络断开
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				// 获取用户的请求信息,该方法只能用作debug mode
				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					initialize.Error(c.Request.URL.Path,
						zap.Time("time", time.Now()),
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}
				//	判断是否需要打印stack
				if stack {
					initialize.Error("[Recovery from panic]",
						zap.Time("time", time.Now()),               // 记录时间
						zap.Any("error", err),                      // 记录错误信息
						zap.String("request", string(httpRequest)), // 请求信息
						zap.Stack("stacktrace"),                    // 调用堆栈信息
					)
				} else {
					initialize.Error("[Recovery from panic]",
						zap.Time("time", time.Now()), // 记录时间
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				// 返回 500 状态码
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"message": "服务器内部错误，请稍后再试",
				})
			}
		}()
		c.Next()
	}
}
