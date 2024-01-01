package drug_doc_handler

import (
	"admin_management_service/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *DrugDocHandler) Search(c *gin.Context) {
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
	var searchResponse models.DrugDocSearchResponse
	if len(keyword) == 0 {
		statusCode, searchResponse, err = h.DrugDocRepo.MatchAll((page-1)*pageSize, pageSize)
	} else {
		statusCode, searchResponse, err = h.DrugDocRepo.MatchQuery((page-1)*pageSize, pageSize, keyword)
	}

	listDrugDocs, total := h.DrugDocService.SummarySearchResponse(searchResponse)
	if statusCode == http.StatusInternalServerError {
		c.JSON(statusCode, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  listDrugDocs,
		"total": total,
	})
}
