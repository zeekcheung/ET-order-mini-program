package logger

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 中间件函数，创建 logger 实例并保存到 Gin 上下文中
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger, _ := zap.NewDevelopment()
		defer logger.Sync()

		c.Set("logger", logger)

		c.Next()
	}
}

func GetLoggerFromCtx(c *gin.Context) *zap.Logger {
	return c.MustGet("logger").(*zap.Logger)
}
