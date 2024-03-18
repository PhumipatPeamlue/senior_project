package core

import (
	"context"
	"reminder_service/internal/core/domains"
)

type NotificationServiceInterface interface {
	AddNewFromHourReminder(ctx context.Context, petID, reminderID string, info domains.HourNotifyInfo) error
	AddNewFromPeriodReminder(ctx context.Context, petID, reminderID string, info domains.PeriodNotifyInfo) error
	ChangeNotifyTimeFromHourReminder(ctx context.Context, petID, reminderID string, info domains.HourNotifyInfo) error
	ChangeNotifyTimeFromPeriodReminder(ctx context.Context, petID, reminderID string, info domains.PeriodNotifyInfo) error
}
