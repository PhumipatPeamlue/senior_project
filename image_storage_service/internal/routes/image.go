package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"image_storage_service/internal/adapters/handlers"
)

func Image(r *gin.Engine, h *handlers.ImageHandler) {
	rg := r.Group("/image")
	{
		rg.GET(fmt.Sprintf("%s/:image_name", h.StoragePath), h.GetImageHandler)
		rg.GET("/:id", h.GetURLHandler)
		rg.POST("/:id", h.SaveHandler)
		rg.DELETE("/:id", h.DeleteHandler)
	}
}
