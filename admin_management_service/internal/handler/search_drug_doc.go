package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) SearchDrugDoc() func(c *gin.Context) {
	return func(c *gin.Context) {
		var err error
		defer func() {
			if err != nil {
				log.Println(err.Error())
			}
		}()

		page, err := strconv.Atoi(c.Query("page"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "page is invalid type"})
			return
		}
		pageSize, err := strconv.Atoi(c.Query("page_size"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "page_size is invalid type"})
			return
		}
		keyword := c.Query("keyword")

		var query string
		if len(keyword) == 0 {
			query = fmt.Sprintf(matchAllQuery, (page-1)*pageSize, pageSize)
		} else {
			query = fmt.Sprintf(searchQueryDrugDoc, (page-1)*pageSize, pageSize, keyword)
		}

		searchResult, err := h.drugDocIndex.Search(query)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"total": searchResult.Hits.Total.Value,
			"data":  searchResult.Hits.Hits,
		})
	}
}
