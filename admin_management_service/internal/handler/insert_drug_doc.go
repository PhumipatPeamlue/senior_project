package handler

import (
	"admin_management_service/internal/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func (h *Handler) InsertDrugDoc() func(c *gin.Context) {
	return func(c *gin.Context) {
		var err error
		defer func() {
			if err != nil {
				log.Println(err.Error())
			}
		}()

		body := models.DrugDoc{
			CreateAt: time.Now(),
		}
		if err = c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "can't bind JSON"})
			return
		}

		if err = h.drugDocIndex.Insert(body); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "can't insert the drug document"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "add the drug document successfully"})
	}
}
