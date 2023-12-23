package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h *Handler) SearchDrugDoc() func(c *gin.Context) {
	return func(c *gin.Context) {
		var err error
		defer func() {
			if err != nil {
				log.Println(err.Error())
			}
		}()

		page, err := h.handleIntQuery(c, "page")
		if err != nil {
			return
		}
		pageSize, err := h.handleIntQuery(c, "page_size")
		if err != nil {
			return
		}
		keyword := c.Query("keyword")

		var query string
		if len(keyword) == 0 {
			query = fmt.Sprintf(matchAllQuery, (page-1)*pageSize, pageSize)
		} else {
			query = fmt.Sprintf(searchQueryDrugDoc, (page-1)*pageSize, pageSize, keyword)
		}

		statusCode, err, searchResult := h.drugDocIndex.Search(query)
		if statusCode == 500 {
			h.handleInternalServerError(c)
			return
		}
		if statusCode != 200 {
			c.JSON(statusCode, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"total": searchResult.Hits.Total.Value,
			"data":  searchResult.Hits.Hits,
		})
	}
}
