package routes

import (
	"document_service/internal/adapter/handlers"
	"github.com/gin-gonic/gin"
)

func DrugDoc(r *gin.Engine, h *handlers.DrugDocHandler) {
	rg := r.Group("/drug_doc")
	{
		rg.GET("/:doc_id", h.GetDrugDoc)
		rg.GET("/search", h.SearchDrugDoc)
		rg.POST("/", h.AddNewDrugDoc)
		rg.PUT("/", h.ChangeDrugDocInfo)
		rg.DELETE("/:doc_id", h.RemoveDrugDoc)
	}
}
