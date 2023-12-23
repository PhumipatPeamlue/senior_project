package handler

import (
	"admin_management_service/internal/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h *Handler) GetDrugDoc() func(c *gin.Context) {
	return func(c *gin.Context) {
		var err error
		defer func() {
			if err != nil {
				log.Println(err.Error())
			}
		}()

		id := c.Param("id")
		statusCode, err, getResult := h.drugDocIndex.Get(id)
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

		res := models.DrugDocDto{
			ID:      getResult.ID,
			DrugDoc: getResult.Source,
		}
		c.JSON(http.StatusOK, res)
	}
}
