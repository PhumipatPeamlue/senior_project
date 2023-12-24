package handler

import (
	"admin_management_service/internal/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func (h *Handler) InsertVideoDoc() func(c *gin.Context) {
	return func(c *gin.Context) {
		var err error
		defer func() {
			if err != nil {
				log.Println(err.Error())
			}
		}()

		now := time.Now()
		body := models.VideoDoc{
			CreateAt: now,
			UpdateAt: now,
		}
		if err = c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "can't bind JSON"})
			return
		}

		statusCode, err := h.videoDocIndex.Insert(body)
		if statusCode == 500 {
			h.handleInternalServerError(c)
			return
		}
		if statusCode != 200 {
			c.JSON(statusCode, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "add the video document successfully"})
	}
}
