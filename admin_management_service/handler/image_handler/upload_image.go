package image_handler

import (
	"admin_management_service/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func (h *ImageHandler) UploadImage(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			log.Println(err)
		}
	}()

	docID := c.Param("doc_id")
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "image (form-data's key) was not found"})
		return
	}

	now := time.Now().String()
	saveName := fmt.Sprintf("%s_%s", now, file.Filename)
	savePath := filepath.Join("uploads", saveName)
	err = c.SaveUploadedFile(file, savePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	image := models.ImageInfo{
		DocID:    docID,
		Filename: saveName,
		Filepath: savePath,
	}
	if err = h.repo.Insert(image); err != nil {
		if err = os.Remove(savePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}
		c.JSON(http.StatusConflict, gin.H{"error": "this document has already added a image"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "insert the image successfully"})
}
