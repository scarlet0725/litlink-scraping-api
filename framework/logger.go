package framework

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Logger interface {
	GinLogger() gin.HandlerFunc
}

type logger struct {
	zap *zap.Logger
}

func NewLogger(zap *zap.Logger) Logger {
	return &logger{zap: zap}
}

func (l *logger) GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()
		l.zap.Info("Logger",
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.String("query", c.Request.URL.RawQuery),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("elapsed", time.Since(start)),
		)
	}
}
