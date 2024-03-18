package core

import (
	"context"
	"reminder_service/internal/core/domains"
)

type hourReminderRepositoryInterface interface {
	ReadHourReminderByID(ctx context.Context, id string) (hr domains.HourReminder, err error)
	CreateHourReminder(ctx context.Context, hr domains.HourReminder) (err error)
	UpdateHourReminder(ctx context.Context, hr domains.HourReminder) (err error)
}

type periodReminderRepositoryInterface interface {
	ReadPeriodReminderByID(ctx context.Context, id string) (pr domains.PeriodReminder, err error)
	CreatePeriodReminder(ctx context.Context, pr domains.PeriodReminder) (err error)
	UpdatePeriodReminder(ctx context.Context, pr domains.PeriodReminder) (err error)
}

type ReminderRepositoryInterface interface {
	hourReminderRepositoryInterface
	periodReminderRepositoryInterface
	ReadByID(ctx context.Context, id string) (reminder domains.Reminder, err error)
	ReadByPetID(ctx context.Context, petID string) (reminders []domains.Reminder, err error)
	DeleteByPetID(ctx context.Context, petID string) (err error)
	DeleteByID(ctx context.Context, id string) (err error)
}
