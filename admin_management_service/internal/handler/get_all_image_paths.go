package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h *Handler) GetAllImagePaths() func(c *gin.Context) {
	return func(c *gin.Context) {
		var err error
		defer func() {
			if err != nil {
				log.Println(err.Error())
			}
		}()

		docID := c.Param("id")
		res, err := h.imageFileRepo.Select(docID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "can't find the result with this document id"})
			return
		}

		c.JSON(http.StatusOK, res)
	}
}
