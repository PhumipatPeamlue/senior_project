package video_doc_handler

import (
	"admin_management_service/models"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *VideoDocHandler) GetVideoDoc(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			log.Println(err)
		}
	}()

	docID := c.Param("doc_id")

	statusCode, getResponse, err := h.VideoDocRepo.Get(docID)
	if statusCode == http.StatusInternalServerError {
		c.JSON(statusCode, gin.H{"error": "internal server error"})
		return
	}
	if statusCode == http.StatusNotFound {
		msg := fmt.Sprintf("%s was not found", docID)
		c.JSON(statusCode, gin.H{"error": msg})
		return
	}

	var videoDoc models.VideoDoc
	videoDoc.ID = docID
	videoDoc.Title = getResponse.Source.Title
	videoDoc.VideoUrl = getResponse.Source.VideoUrl
	videoDoc.Description = getResponse.Source.Description
	videoDoc.CreateAt = getResponse.Source.CreateAt
	videoDoc.UpdateAt = getResponse.Source.UpdateAt
	c.JSON(statusCode, videoDoc)
}
