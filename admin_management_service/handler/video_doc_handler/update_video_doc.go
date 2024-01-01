package video_doc_handler

import (
	"admin_management_service/models"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *VideoDocHandler) UpdateVideoDoc(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			log.Println(err)
		}
	}()

	now := time.Now()
	var body models.VideoDoc
	body.UpdateAt = &now
	if err = c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't bind JSON"})
		return
	}

	statusCode, getResponse, err := h.VideoDocRepo.Get(body.ID)
	if statusCode == http.StatusInternalServerError {
		c.JSON(statusCode, gin.H{"error": "internal server error"})
		return
	}
	if statusCode == http.StatusNotFound {
		msg := fmt.Sprintf("%s was not found", body.ID)
		c.JSON(statusCode, gin.H{"error": msg})
		return
	}

	body.CreateAt = getResponse.Source.CreateAt
	updatedBody := models.VideoDocUpdatedBody{
		Doc: body.VideoDocES,
	}
	statusCode, err = h.VideoDocRepo.Update(body.ID, updatedBody)
	if statusCode == http.StatusInternalServerError {
		c.JSON(statusCode, gin.H{"error": "internal server error"})
		return
	}
	if statusCode == http.StatusNotFound {
		msg := fmt.Sprintf("%s was not found", body.ID)
		c.JSON(statusCode, gin.H{"error": msg})
		return
	}

	c.JSON(statusCode, gin.H{"message": "update the video document successfully"})
}
