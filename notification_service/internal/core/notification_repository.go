package core

import (
	"context"
	"notification_service/internal/core/domains"
)

type NotificationRepositoryInterface interface {
	Create(ctx context.Context, notification domains.Notification) (err error)
	DeleteTodayAndWaitStatusByReminderID(ctx context.Context, reminderID string) (err error)
}
