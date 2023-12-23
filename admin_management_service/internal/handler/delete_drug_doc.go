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
		statusCode, err := h.drugDocIndex.Delete(id)
		if statusCode == 500 {
			h.handleInternalServerError(c)
			return
		}
		if statusCode == 404 {
			h.handleNotFound(c, id)
			return
		}
		if statusCode != 200 {
			c.JSON(statusCode, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "delete the drug document successfully"})
	}
}
