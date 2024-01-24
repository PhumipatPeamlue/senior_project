package routes

import (
	"document_service/internal/adapter/handlers"
	"github.com/gin-gonic/gin"
)

func Image(r *gin.Engine, h *handlers.ImageHandler) {
	rg := r.Group("/" + h.StoragePath)
	{
		rg.GET("/:image_name", h.GetImage)
	}
}
