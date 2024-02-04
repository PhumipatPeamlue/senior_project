package routes

import (
	"document_service/internal/adapters/handlers"
	"github.com/gin-gonic/gin"
)

func VideoDoc(r *gin.Engine, h *handlers.VideoDocHandler) {
	rg := r.Group("/document/video_doc")
	{
		rg.GET("/:doc_id", h.GetVideoDocHandler)
		rg.GET("/search", h.SearchVideoDocHandler)
		rg.POST("/", h.AddNewVideoDocHandler)
		rg.PUT("/", h.ChangeVideoDocInfoHandler)
		rg.DELETE("/:doc_id", h.RemoveVideoDocHandler)
	}
}
