package http_gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		allowedOrigins := []string{"http://localhost:3000", "http://localhost:3001"}
		origin := c.Request.Header.Get("Origin")

		for _, allowedOrigin := range allowedOrigins {
			if origin == allowedOrigin {
				c.Header("Access-Control-Allow-Origin", origin)
				c.Header("Access-Control-Allow-Credentials", "true")
				c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
				c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")

				if c.Request.Method == "OPTIONS" {
					c.AbortWithStatus(http.StatusNoContent)
					return
				}

				break
			}
		}

		c.Next()
	}
}
