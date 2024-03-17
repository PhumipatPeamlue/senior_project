package http_gin

import "github.com/gin-gonic/gin"

func PetRoutes(r *gin.Engine, h *petHandler) {
	petRouter := r.Group("/pet")
	{
		petRouter.POST("/", h.AddNewPet)
		petRouter.PUT("/", h.ChangePetName)
		petRouter.GET("/all/:user_id", h.FindAllUserPets)
		petRouter.GET("/:pet_id", h.FindPet)
		petRouter.DELETE("/all/:user_id", h.RemoveAllUserPets)
		petRouter.DELETE("/:pet_id", h.RemovePet)
	}
}
