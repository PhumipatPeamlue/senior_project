package doc_image_handler

import (
	"github.com/gin-gonic/gin"
	"path"
)

type DocImageHandler struct{}

func NewDocImageHandler() *DocImageHandler {
	return &DocImageHandler{}
}

func (h *DocImageHandler) GetImage(c *gin.Context) {
	imageName := c.Param("filename")
	c.File(path.Join("uploads", imageName))
}
