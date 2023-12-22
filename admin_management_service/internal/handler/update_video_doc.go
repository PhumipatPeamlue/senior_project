package handler

import (
	"admin_management_service/internal/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func (h *Handler) UpdateVideoDoc() func(c *gin.Context) {
	return func(c *gin.Context) {
		var err error
		defer func() {
			if err != nil {
				log.Println(err.Error())
			}
		}()

		body := models.VideoDocDto{}
		if err = c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "can't bind JSON"})
			return
		}

		getResult, err := h.videoDocIndex.Get(body.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("can't get the video document with id = %s", body.ID),
			})
		}

		body.VideoDoc.CreateAt = getResult.Source.CreateAt
		body.VideoDoc.UpdateAt = time.Now()
		if err = h.videoDocIndex.Update(body.ID, body.VideoDoc); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "can't update the video document"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "update the video document successfully"})
	}
}
