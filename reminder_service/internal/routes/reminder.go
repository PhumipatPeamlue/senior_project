package routes

import (
	"github.com/gin-gonic/gin"
	"reminder_service/internal/adapters/handlers"
)

func Reminder(r *gin.Engine, h *handlers.ReminderHandler) {
	rg := r.Group("/reminder")
	{
		rg.GET("/period/:reminder_id", h.GetPeriodReminderInfoHandler)
		rg.GET("/hour/:reminder_id", h.GetHourReminderInfoHandler)
		rg.GET("/all/:pet_id", h.GetAllRemindersHandler)
		rg.POST("/period", h.AddNewPeriodReminderHandler)
		rg.POST("/hour", h.AddNewHourReminderHandler)
		rg.PUT("/period", h.ChangePeriodReminderInfoHandler)
		rg.PUT("/hour", h.ChangeHourReminderInfoHandler)
		rg.DELETE("/:reminder_id", h.RemoveHandler)
	}
}
