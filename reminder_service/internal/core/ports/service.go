package ports

import (
	"context"
	"reminder_service/internal/core/domains"
	"time"
)

type NotificationService interface {
	AddNewFromPeriodReminder(ctx context.Context, reminderID string, userID string, morning *time.Time, noon *time.Time, evening *time.Time, beforeBed *time.Time) (err error)
	AddNewFromHourReminder(ctx context.Context, reminderID string, userID string, firstUsage time.Time, every int) (err error)
	RemoveByReminderID(ctx context.Context, reminderID string) (err error)
}

type ReminderService interface {
	GetAllPetReminders(ctx context.Context, petID string) (reminders []domains.Reminder, err error)
	GetPeriodReminderInfo(ctx context.Context, reminderID string) (reminder domains.PeriodReminder, err error)
	GetHourReminderInfo(ctx context.Context, reminderID string) (reminder domains.HourReminder, err error)
	AddNewPeriodReminder(ctx context.Context, userID string, petID string, drugName string, drugUsage string, frequency string, morning *time.Time, noon *time.Time, evening *time.Time, beforeBed *time.Time) (err error)
	AddNewHourReminder(ctx context.Context, userID string, petID string, drugName string, drugUsage string, frequency string, firstUsage time.Time, every int) (err error)
	ChangePeriodReminderInfo(ctx context.Context, userID string, reminderID string, drugName string, drugUsage string, frequency string, morning *time.Time, noon *time.Time, evening *time.Time, beforeBed *time.Time) (err error)
	ChangeHourReminderInfo(ctx context.Context, userID string, reminderID string, drugName string, drugUsage string, frequency string, firstUsage time.Time, every int) (err error)
	Remove(ctx context.Context, reminderID string) (err error)
}
