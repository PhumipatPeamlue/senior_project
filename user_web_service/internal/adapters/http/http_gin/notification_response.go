package http_gin

import "user_web_service/internal/core"

type notificationResponse struct {
	NotificationID   string `json:"reminder_id"`
	PetID            string `json:"pet_id"`
	NotificationType string `json:"type"`
	core.DrugInfo
}

type hourNotificationResponse struct {
	NotificationID string `json:"reminder_id"`
	PetID          string `json:"pet_id"`
	Frequency      int    `json:"frequency"`
	core.DrugInfo
	core.HourNotifyInfo
}

type periodNotificationResponse struct {
	NotificationID string `json:"reminder_id"`
	PetID          string `json:"pet_id"`
	Frequency      int    `json:"frequency"`
	core.DrugInfo
	core.PeriodNotifyInfo
}
