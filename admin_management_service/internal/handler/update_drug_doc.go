package handler

import (
	"admin_management_service/internal/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h *Handler) UpdateDrugDoc() func(c *gin.Context) {
	return func(c *gin.Context) {
		var err error
		defer func() {
			if err != nil {
				log.Println(err.Error())
			}
		}()

		body := models.DrugDocDto{}
		if err = c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "can't bind JSON"})
			return
		}

		if err = h.drugDocIndex.Update(body.ID, body.DrugDoc); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "can't update the drug document"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "update the drug document successfully"})
	}
}
