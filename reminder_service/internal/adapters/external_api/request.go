package external_api

import "reminder_service/internal/core/domains"

type addNewFromHourReminderRequest struct {
	PetID      string `json:"pet_id"`
	ReminderID string `json:"reminder_id"`
	domains.HourNotifyInfo
}

type addNewFromPeriodReminderRequest struct {
	PetID      string `json:"pet_id"`
	ReminderID string `json:"reminder_id"`
	domains.PeriodNotifyInfo
}

type changeNotifyTimeFromHourReminder struct {
	PetID      string `json:"pet_id"`
	ReminderID string `json:"reminder_id"`
	domains.HourNotifyInfo
}

type changeNotifyTimeFromPeriodReminder struct {
	PetID      string `json:"pet_id"`
	ReminderID string `json:"reminder_id"`
	domains.PeriodNotifyInfo
}
