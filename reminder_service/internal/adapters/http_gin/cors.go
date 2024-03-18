package http_gin

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func Cors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:  []string{"*"},
		ExposeHeaders: []string{"Content-Length"},
		// AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	})
}
