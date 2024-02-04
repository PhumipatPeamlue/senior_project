package routes

import (
	"github.com/gin-gonic/gin"
	"user_service/internal/adapters/handlers"
)

func User(r *gin.Engine, h *handlers.UserHandler) {
	rg := r.Group("/user")
	{
		rg.GET("/:user_id", h.GetTimeSettingHandler)
		rg.PUT("/", h.ChangeTimeSettingsHandler)
		rg.DELETE("/:user_id", h.RemoveUserDataByIDHandler)
	}
}
