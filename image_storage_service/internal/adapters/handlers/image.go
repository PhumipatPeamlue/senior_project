package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"image_storage_service/internal/core/ports"
	"net/http"
	"path/filepath"
	"time"
)

type ImageHandler struct {
	service     ports.ImageStorageService
	StoragePath string
}

func (h *ImageHandler) GetImageHandler(c *gin.Context) {
	imageName := c.Param("image_name")
	c.File(filepath.Join(h.StoragePath, imageName))
}

func (h *ImageHandler) GetURLHandler(c *gin.Context) {
	id := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	url, err := h.service.GetURL(ctx, id)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"image_url": url,
	})
}

func (h *ImageHandler) SaveHandler(c *gin.Context) {
	id := c.Param("id")
	file, header, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "image not found"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = h.service.Save(ctx, id, &file, header); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "save image successfully"})
}

func (h *ImageHandler) DeleteHandler(c *gin.Context) {
	id := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := h.service.Delete(ctx, id); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "delete image successfully"})
}

func NewImageHandler(service ports.ImageStorageService, storagePath string) *ImageHandler {
	return &ImageHandler{
		service:     service,
		StoragePath: storagePath,
	}
}
