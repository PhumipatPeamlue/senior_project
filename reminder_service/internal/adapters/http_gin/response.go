package http_gin

import "reminder_service/internal/core/domains"

type reminderResponse struct {
	ReminderID   string `json:"reminder_id"`
	PetID        string `json:"pet_id"`
	ReminderType string `json:"type"`
	domains.DrugInfo
}

type hourReminderResponse struct {
	ReminderID string `json:"reminder_id"`
	PetID      string `json:"pet_id"`
	Frequency  int    `json:"frequency"`
	domains.DrugInfo
	domains.HourNotifyInfo
}

type periodReminderResponse struct {
	ReminderID string `json:"reminder_id"`
	PetID      string `json:"pet_id"`
	Frequency  int    `json:"frequency"`
	domains.DrugInfo
	domains.PeriodNotifyInfo
}
