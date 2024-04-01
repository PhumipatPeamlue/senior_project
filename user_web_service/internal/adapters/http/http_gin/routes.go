package http_gin

import "github.com/gin-gonic/gin"

func UserRoutes(r *gin.Engine, h *UserHandler) {
	userRouter := r.Group("/user")
	{
		userRouter.GET("/:line_user_id", h.FindUserByLineUserID)
		userRouter.PUT("/", h.ChangeTimeSetting)
		userRouter.DELETE("/:line_user_id", h.RemoveUserByLineUserID)
	}
}

func PetRoutes(r *gin.Engine, h *PetHandler) {
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

func NotificationRoutes(r *gin.Engine, h *NotificationHandler) {
	notificationRouter := r.Group("/reminder")
	{
		notificationRouter.GET("/:reminder_id", h.FindNotification)
		notificationRouter.GET("/all/:pet_id", h.FindAllPetNotifications)
		notificationRouter.DELETE("/:pet_id", h.RemoveAllPetNotifications)
	}

	hourNotificationRouter := notificationRouter.Group("/hour")
	{
		hourNotificationRouter.POST("/", h.AddNewHourNotification)
		hourNotificationRouter.PUT("/", h.ChangeHourNotificationInfo)
		hourNotificationRouter.GET("/:reminder_id", h.FindHourNotification)
		hourNotificationRouter.DELETE("/:reminder_id", h.RemoveHourNotification)
	}

	periodNotificationRouter := notificationRouter.Group("/period")
	{
		periodNotificationRouter.POST("/", h.AddNewPeriodNotification)
		periodNotificationRouter.PUT("/", h.ChangePeriodNotificationInfo)
		periodNotificationRouter.GET("/:reminder_id", h.FindPeriodNotification)
		periodNotificationRouter.DELETE("/:reminder_id", h.RemovePeriodNotification)
	}
}
