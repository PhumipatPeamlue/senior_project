package handler

import (
	"admin_management_service/internal/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
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

		getResult, err := h.drugDocIndex.Get(body.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("can't get the drug document with id = %s", body.ID),
			})
		}

		body.DrugDoc.CreateAt = getResult.Source.CreateAt
		body.DrugDoc.UpdateAt = time.Now()
		if err = h.drugDocIndex.Update(body.ID, body.DrugDoc); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "can't update the drug document"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "update the drug document successfully"})
	}
}
