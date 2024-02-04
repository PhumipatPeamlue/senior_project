package middlewares

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"image_storage_service/internal/core/services"
	"net/http"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		err := c.Errors.Last()
		if err == nil {
			return
		}

		if errors.Is(err, context.DeadlineExceeded) {
			c.JSON(http.StatusRequestTimeout, gin.H{"error": "request timeout"})
			return
		}

		if errors.Is(err, services.NotFoundError) {
			c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
			return
		}

		if errors.Is(err, services.DuplicateError) {
			c.JSON(http.StatusConflict, gin.H{"error": "already exists"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
}
