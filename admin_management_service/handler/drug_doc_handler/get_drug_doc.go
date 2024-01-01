package drug_doc_handler

import (
	"admin_management_service/models"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *DrugDocHandler) GetDrugDoc(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			log.Println(err)
		}
	}()

	docID := c.Param("doc_id")

	statusCode, getResponse, err := h.DrugDocRepo.Get(docID)
	if statusCode == http.StatusInternalServerError {
		c.JSON(statusCode, gin.H{"error": "internal server error"})
		return
	}
	if statusCode == http.StatusNotFound {
		msg := fmt.Sprintf("%s was not found", docID)
		c.JSON(statusCode, gin.H{"error": msg})
		return
	}

	var drugDoc models.DrugDoc
	drugDoc.ID = docID
	drugDoc.TradeName = getResponse.Source.TradeName
	drugDoc.DrugName = getResponse.Source.DrugName
	drugDoc.Description = getResponse.Source.Description
	drugDoc.Preparation = getResponse.Source.Preparation
	drugDoc.Caution = getResponse.Source.Caution
	drugDoc.CreateAt = getResponse.Source.CreateAt
	drugDoc.UpdateAt = getResponse.Source.UpdateAt
	c.JSON(statusCode, drugDoc)
}
