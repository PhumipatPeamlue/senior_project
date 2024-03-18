package http_gin

import (
	"notification_service/internal/core/domains"
)

type addNewFromHourReminderRequest struct {
	PetID      string `json:"pet_id" binding:"required"`
	ReminderID string `json:"reminder_id" binding:"required"`
	domains.HourNotifyInfo
}

type addNewFromPeriodReminderRequest struct {
	PetID      string `json:"pet_id" binding:"required"`
	ReminderID string `json:"reminder_id" binding:"required"`
	domains.PeriodNotifyInfo
}

type changeNotifyTimeFromHourReminderRequest struct {
	PetID      string `json:"pet_id" binding:"required"`
	ReminderID string `json:"reminder_id" binding:"required"`
	domains.HourNotifyInfo
}

type changeNotifyTimeFromPeriodReminderRequest struct {
	PetID      string `json:"pet_id" binding:"required"`
	ReminderID string `json:"reminder_id" binding:"required"`
	domains.PeriodNotifyInfo
}
