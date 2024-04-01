package http_gin

import "user_web_service/internal/core"

type addNewHourNotificationRequest struct {
	PetID     string `json:"pet_id" binding:"required"`
	Frequency int    `json:"frequency"`
	core.DrugInfo
	core.HourNotifyInfo
}

type changeHourDrugLabelInfoRequest struct {
	NotificationID string `json:"reminder_id" binding:"required"`
	Frequency      int    `json:"frequency"`
	core.DrugInfo
	core.HourNotifyInfo
}

type addNewPeriodNotificationRequest struct {
	PetID     string `json:"pet_id" binding:"required"`
	Frequency int    `json:"frequency"`
	core.DrugInfo
	core.PeriodNotifyInfo
}

type changePeriodDrugLabelInfoRequest struct {
	NotificationID string `json:"reminder_id" binding:"required"`
	Frequency      int    `json:"frequency"`
	core.DrugInfo
	core.PeriodNotifyInfo
}
