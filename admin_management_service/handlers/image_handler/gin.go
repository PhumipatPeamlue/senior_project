package image_handler

import (
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type GinImageHandler struct{}

func NewGinImageHandler() *GinImageHandler {
	return &GinImageHandler{}
}

func (h *GinImageHandler) HandlerGetImage(c *gin.Context) {
	imageName := c.Param("image_name")
	path := filepath.Join("uploads", imageName)
	c.File(path)
}
