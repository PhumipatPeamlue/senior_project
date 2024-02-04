package routes

import (
	"github.com/gin-gonic/gin"
	"pet_service/internal/adapters/handlers"
)

func Pet(r *gin.Engine, h *handlers.PetHandler) {
	rg := r.Group("/pet")
	{
		rg.GET("/:pet_id", h.GetPetInfoHandler)
		rg.GET("/all/:user_id", h.GetAllUserPetHandler)
		rg.POST("/", h.AddNewPetHandler)
		rg.PUT("/", h.ChangePetInfoHandler)
		rg.DELETE("/:pet_id", h.RemovePetHandler)
	}

}
