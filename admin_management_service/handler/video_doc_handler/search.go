package video_doc_handler

import (
	"admin_management_service/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *VideoDocHandler) Search(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			log.Println(err)
		}
	}()

	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "page is not integer"})
		return
	}
	pageSize, err := strconv.Atoi(c.Query("page_size"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "page_size is not integer"})
		return
	}
	keyword := c.Query("keyword")

	var statusCode int
	var searchResponse models.VideoDocSearchResponse
	if len(keyword) == 0 {
		statusCode, searchResponse, err = h.VideoDocRepo.MatchAll((page-1)*pageSize, pageSize)
	} else {
		statusCode, searchResponse, err = h.VideoDocRepo.MatchQuery((page-1)*pageSize, pageSize, keyword)
	}

	listVideoDocs, total := h.VideoDocService.SummarySearchResponse(searchResponse)
	if statusCode == http.StatusInternalServerError {
		c.JSON(statusCode, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  listVideoDocs,
		"total": total,
	})
}
