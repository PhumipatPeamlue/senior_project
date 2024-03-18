package http_gin

import (
	"github.com/gin-gonic/gin"
)

func ImageRoutes(r *gin.Engine, h *localFileStorageHandler) {
	imageRouter := r.Group("/image")
	{
		imageRouter.GET("/bucket/:file_name", h.GetImage)
		imageRouter.GET("/:id", h.GetURL)
		imageRouter.POST("/:id", h.Save)
		imageRouter.PUT("/:id", h.ChangeFile)
		imageRouter.DELETE("/:id", h.Delete)
	}
}
