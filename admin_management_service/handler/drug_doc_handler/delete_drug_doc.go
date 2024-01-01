package drug_doc_handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *DrugDocHandler) DeleteDrugDoc(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			log.Println(err)
		}
	}()

	docID := c.Param("doc_id")

	statusCode, err := h.DrugDocRepo.Delete(docID)
	if statusCode == http.StatusInternalServerError {
		c.JSON(statusCode, gin.H{"error": "internal server error"})
		return
	}
	if statusCode == http.StatusNotFound {
		msg := fmt.Sprintf("%s was not found", docID)
		c.JSON(statusCode, gin.H{"error": msg})
		return
	}

	c.JSON(statusCode, gin.H{"message": "delete the drug document successfully"})
}
