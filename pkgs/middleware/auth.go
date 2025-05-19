package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}	
}

func CustomLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		if c.FullPath() != "/health" {
			latency := time.Since(start).Milliseconds()
			userID := "userID"
			uID := "c.GetString"
			logger.Info(path,
				zap.String("requestID", uID),
				zap.String("leadId", "ucc"),
				zap.String("userId", userID),
				zap.Int("status", c.Writer.Status()),
				zap.String("method", c.Request.Method),
				zap.String("path", path),
				zap.String("query", query),
				zap.String("user-agent", c.Request.UserAgent()),
				zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
				zap.Int64("latency", latency),
			)
		}
	}
}
