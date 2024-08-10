package middleware

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"time"
)

func Logger(log *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()
		latency := endTime.Sub(startTime)

		// 记录请求日志
		log.Info("Request",
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
			"status", c.Writer.Status(),
			"latency", latency,
			"client_ip", c.ClientIP(),
		)
	}
}
