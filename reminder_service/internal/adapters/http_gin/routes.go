package http_gin

import "github.com/gin-gonic/gin"

func ReminderRoutes(r *gin.Engine, h *reminderHandler) {
	reminderRouter := r.Group("/reminder")
	{
		reminderRouter.GET("/:reminder_id", h.FindReminder)
		reminderRouter.GET("/all/:pet_id", h.FindAllPetReminders)
		reminderRouter.DELETE("/:pet_id", h.RemoveAllPetReminders)
	}

	hourReminderRouter := reminderRouter.Group("/hour")
	{
		hourReminderRouter.POST("/", h.AddNewHourReminder)
		hourReminderRouter.PUT("/", h.ChangeHourReminderInfo)
		hourReminderRouter.GET("/:reminder_id", h.FindHourReminder)
		hourReminderRouter.DELETE("/:reminder_id", h.RemoveHourReminder)
	}

	periodReminderRouter := reminderRouter.Group("/period")
	{
		periodReminderRouter.POST("/", h.AddNewPeriodReminder)
		periodReminderRouter.PUT("/", h.ChangePeriodReminderInfo)
		periodReminderRouter.GET("/:reminder_id", h.FindPeriodReminder)
		periodReminderRouter.DELETE("/:reminder_id", h.RemovePeriodReminder)
	}
}
