package core

import (
	"context"
	"cronjob/internal/core/domains"
)

type ReminderRepositoryInterface interface {
	ReadByID(ctx context.Context, id string) (reminder domains.Reminder, err error)
	ReadAll(ctx context.Context) (reminders []domains.Reminder, err error)
	ReadAllZeroRenew(ctx context.Context) (reminders []domains.Reminder, err error)
	UpdateRenew(ctx context.Context, reminder domains.Reminder) (err error)
}
