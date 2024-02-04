package ports

import (
	"context"
	"reminder_service/internal/core/domains"
)

type ReminderRepository interface {
	FindByPetID(ctx context.Context, petID string) (reminders []domains.Reminder, err error)
	Save(ctx context.Context, reminder domains.Reminder) (err error)
	Update(ctx context.Context, reminder domains.Reminder) (err error)
	Delete(ctx context.Context, id string) (err error)
}

type PeriodReminderRepository interface {
	FindByID(ctx context.Context, id string) (pr domains.PeriodReminder, err error)
	Save(ctx context.Context, pr domains.PeriodReminder) (err error)
	Update(ctx context.Context, pr domains.PeriodReminder) (err error)
}

type HourReminderRepository interface {
	FindByID(ctx context.Context, id string) (hr domains.HourReminder, err error)
	Save(ctx context.Context, hr domains.HourReminder) (err error)
	Update(ctx context.Context, hr domains.HourReminder) (err error)
}

type NotificationRepository interface {
	Save(ctx context.Context, notification domains.Notification) (err error)
	DeleteByReminderID(ctx context.Context, reminderID string) (err error)
}
