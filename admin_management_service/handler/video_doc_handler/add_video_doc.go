package video_doc_handler

import (
	"admin_management_service/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func (h *VideoDocHandler) AddVideoDoc(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			log.Println(err)
		}
	}()

	now := time.Now()
	var body models.VideoDocES
	body.CreateAt = &now
	body.UpdateAt = &now
	if err = c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't find JSON"})
		return
	}

	statusCode, docID, err := h.VideoDocRepo.Index(body)
	if statusCode == http.StatusInternalServerError {
		c.JSON(statusCode, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(statusCode, gin.H{"doc_id": docID})
}
