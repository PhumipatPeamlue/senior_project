package routes

import (
	"document_service/internal/adapters/handlers"
	"github.com/gin-gonic/gin"
)

func DrugDoc(r *gin.Engine, h *handlers.DrugDocHandler) {
	rg := r.Group("/document/drug_doc")
	{
		rg.GET("/:doc_id", h.GetDrugDocHandler)
		rg.GET("/search", h.SearchDrugDocHandler)
		rg.POST("/", h.AddNewDrugDocHandler)
		rg.PUT("/", h.ChangeDrugDocInfoHandler)
		rg.DELETE("/:doc_id", h.RemoveDrugDocHandler)
	}
}
