package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h *Handler) DeleteDrugDoc() func(c *gin.Context) {
	return func(c *gin.Context) {
		var err error
		defer func() {
			if err != nil {
				log.Println(err.Error())
			}
		}()

		id := c.Param("id")
		if err = h.drugDocIndex.Delete(id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "can't delete the drug document"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "delete the drug document successfully"})
	}
}
