package handlers

import (
	"github.com/gin-gonic/gin"
	"path/filepath"
)

type ImageHandler struct {
	StoragePath string
}

func NewImageHandler(storagePath string) *ImageHandler {
	return &ImageHandler{
		StoragePath: storagePath,
	}
}

func (h *ImageHandler) GetImage(c *gin.Context) {
	imageName := c.Param("image_name")
	c.File(filepath.Join(h.StoragePath, imageName))
}
