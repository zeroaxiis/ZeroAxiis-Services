package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func JSONBodyLimit(maxBytes int64) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.ContentType() == "application/json" {
			c.Request.Body = http.MaxBytesReader(
				c.Writer,
				c.Request.Body,
				maxBytes,
			)
		}
		c.Next()
	}
}
