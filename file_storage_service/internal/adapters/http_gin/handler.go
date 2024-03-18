package http_gin

import (
	"context"
	"errors"
	"file_storage_service/internal/core"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"time"
)

type localFileStorageHandler struct {
	service core.LocalFileStorageServiceInterface
}

func (l *localFileStorageHandler) handleError(c *gin.Context, err error) {
	c.Error(err)

	if errors.Is(err, context.DeadlineExceeded) {
		c.JSON(http.StatusRequestTimeout, gin.H{"error": "request timeout"})
		return
	}

	var errFileInfoNotFound *core.ErrFileInfoNotFound
	var ErrFileInfoDuplicate *core.ErrFileInfoDuplicate
	switch {
	case errors.As(err, &errFileInfoNotFound):
		c.JSON(http.StatusNotFound, gin.H{"error": "file info not found"})
	case errors.As(err, &ErrFileInfoDuplicate):
		c.JSON(http.StatusConflict, gin.H{"error": "file info already exists"})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
	}
}

func (l *localFileStorageHandler) GetImage(c *gin.Context) {
	fileName := c.Param("file_name")
	fmt.Println(fileName)
	c.File(filepath.Join("bucket", fileName))
}

func (l *localFileStorageHandler) GetURL(c *gin.Context) {
	id := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	url, err := l.service.GetURL(ctx, id)
	if err != nil {
		l.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"image_url": url,
	})
}

func (l *localFileStorageHandler) Save(c *gin.Context) {
	id := c.Param("id")
	file, header, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "image not found"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = l.service.Save(ctx, id, &file, header.Filename)
	if err != nil {
		l.handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "save image successfully"})
}

func (l *localFileStorageHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := l.service.Delete(ctx, id); err != nil {
		l.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "delete image successfully"})
}

func (l *localFileStorageHandler) ChangeFile(c *gin.Context) {
	id := c.Param("id")
	file, header, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "image not found"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = l.service.ChangeFile(ctx, id, &file, header.Filename)
	if err != nil {
		l.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "change the image successfully"})
}

func NewLocalFileStorageHandler(service core.LocalFileStorageServiceInterface) *localFileStorageHandler {
	return &localFileStorageHandler{
		service: service,
	}
}
