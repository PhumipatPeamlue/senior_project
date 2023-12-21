package handler

import (
	"admin_management_service/internal/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path/filepath"
)

func (h *Handler) InsertImage() func(c *gin.Context) {
	return func(c *gin.Context) {
		var err error
		defer func() {
			if err != nil {
				log.Println(err.Error())
			}
		}()

		file, err := c.FormFile("image")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "image key not found"})
			return
		}

		path := filepath.Join("uploads", file.Filename)
		err = c.SaveUploadedFile(file, path)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "can't save the image in sever's file system"})
			return
		}

		docID := c.Param("id")
		row := models.ImageFile{
			DocID:    docID,
			Filename: file.Filename,
			FilePath: path,
		}
		_, err = h.imageFileRepo.Insert(row)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "can't insert image's path to database"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "insert the image successfully"})
	}
}
