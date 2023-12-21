package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteImage() func(c *gin.Context) {
	return func(c *gin.Context) {
		var err error
		defer func() {
			if err != nil {
				log.Println(err.Error())
			}
		}()

		imageID, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "image's id is invalid type"})
			return
		}
		err = h.imageFileRepo.Delete(imageID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "delete image successfully"})
	}
}
