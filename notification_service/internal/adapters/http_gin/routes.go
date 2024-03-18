package http_gin

import "github.com/gin-gonic/gin"

func NotificationRoutes(r *gin.Engine, h *notificationHandler) {
	notificationRouter := r.Group("/notification")
	hourRouter := notificationRouter.Group("/hour")
	{
		hourRouter.POST("/", h.AddNewFromHourReminder)
		hourRouter.PUT("/", h.ChangeNotifyTimeFromHourReminder)
	}
	periodRouter := notificationRouter.Group("/period")
	{
		periodRouter.POST("/", h.AddNewFromPeriodReminder)
		periodRouter.PUT("/", h.ChangeNotifyTimeFromPeriodReminder)
	}
}
