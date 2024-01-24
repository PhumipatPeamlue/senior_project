package routes

import (
	"document_service/internal/adapter/handlers"
	"github.com/gin-gonic/gin"
)

func VideoDoc(r *gin.Engine, h *handlers.VideoDocHandler) {
	rg := r.Group("/video_doc")
	{
		rg.GET("/:doc_id", h.GetVideoDoc)
		rg.GET("/search", h.SearchVideoDoc)
		rg.POST("/", h.AddNewVideoDoc)
		rg.PUT("/", h.ChangeVideoDocInfo)
		rg.DELETE("/:doc_id", h.RemoveVideoDoc)
	}
}
