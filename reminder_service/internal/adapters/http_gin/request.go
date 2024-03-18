package http_gin

import "reminder_service/internal/core/domains"

type addNewHourReminderRequest struct {
	PetID     string `json:"pet_id" binding:"required"`
	Frequency int    `json:"frequency"`
	domains.DrugInfo
	domains.HourNotifyInfo
}

type changeHourDrugLabelInfoRequest struct {
	ReminderID string `json:"reminder_id" binding:"required"`
	Frequency  int    `json:"frequency"`
	domains.DrugInfo
	domains.HourNotifyInfo
}

type addNewPeriodReminderRequest struct {
	PetID     string `json:"pet_id" binding:"required"`
	Frequency int    `json:"frequency"`
	domains.DrugInfo
	domains.PeriodNotifyInfo
}

type changePeriodDrugLabelInfoRequest struct {
	ReminderID string `json:"reminder_id" binding:"required"`
	Frequency  int    `json:"frequency"`
	domains.DrugInfo
	domains.PeriodNotifyInfo
}
