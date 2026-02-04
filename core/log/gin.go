package log

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

// GinLogger returns a gin handler func that logs requests using zerolog
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		c.Next()

		if raw != "" {
			path = path + "?" + raw
		}

		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		method := c.Request.Method
		errorMessage := c.Errors.ByType(gin.ErrorTypePrivate).String()
		bodySize := c.Writer.Size()

		var event *zerolog.Event
		if statusCode >= 500 {
			event = Error()
		} else if statusCode >= 400 {
			event = Warn()
		} else {
			event = Info()
		}

		if errorMessage != "" {
			event.Str("error", errorMessage)
		}

		event.
			Int("status", statusCode).
			Str("method", method).
			Str("path", path).
			Str("ip", clientIP).
			Dur("latency", time.Since(start)).
			Int("body_size", bodySize).
			Msg("GIN")
	}
}
