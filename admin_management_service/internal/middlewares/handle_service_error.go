package middlewares

import (
	"document_service/internal/core"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleServiceError() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		err := c.Errors.Last()
		if err == nil {
			return
		}

		//log.Println(err.Error())
		var e *core.Error
		if errors.As(err, &e) {
			if e.Code() == core.CodeErrorNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
			} else if e.Code() == core.CodeErrorDuplicate {
				c.JSON(http.StatusConflict, gin.H{"error": "already exists"})
			} else if e.Code() == core.CodeErrorBadRequest {
				c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
			}
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		}
	}
}
