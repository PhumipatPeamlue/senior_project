package drug_doc_handler

import (
	"admin_management_service/models"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *DrugDocHandler) UpdateDrugDoc(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			log.Println(err)
		}
	}()

	now := time.Now()
	var body models.DrugDoc
	body.UpdateAt = &now
	if err = c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't bind JSON"})
		return
	}

	statusCode, getResponse, err := h.DrugDocRepo.Get(body.ID)
	if statusCode == http.StatusInternalServerError {
		c.JSON(statusCode, gin.H{"error": "internal server error"})
		return
	}
	if statusCode == http.StatusNotFound {
		msg := fmt.Sprintf("%s was not found", body.ID)
		c.JSON(statusCode, gin.H{"error": msg})
		return
	}

	body.CreateAt = getResponse.Source.CreateAt
	updatedBody := models.DrugDocUpdatedBody{
		Doc: body.DrugDocES,
	}
	statusCode, err = h.DrugDocRepo.Update(body.ID, updatedBody)
	if statusCode == http.StatusInternalServerError {
		c.JSON(statusCode, gin.H{"error": "internal server error"})
		return
	}
	if statusCode == http.StatusNotFound {
		msg := fmt.Sprintf("%s was not found", body.ID)
		c.JSON(statusCode, gin.H{"error": msg})
		return
	}

	c.JSON(statusCode, gin.H{"message": "update the drug document successfully"})
}
