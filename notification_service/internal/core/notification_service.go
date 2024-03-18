package core

import (
	"context"
	"notification_service/internal/core/domains"
	"time"
)

type NotificationServiceInterface interface {
	addNew(ctx context.Context, notification domains.Notification) (err error)
	removeAllTodayAndWaitStatusNotification(ctx context.Context, reminderID string) (err error)
	AddNewFromHourReminder(ctx context.Context, petID string, reminderID string, info domains.HourNotifyInfo) (err error)
	AddNewFromPeriodReminder(ctx context.Context, petID string, reminderID string, info domains.PeriodNotifyInfo) (err error)
	ChangeNotifyTimeFromHourReminder(ctx context.Context, petID string, reminderID string, info domains.HourNotifyInfo) (err error)
	ChangeNotifyTimeFromPeriodReminder(ctx context.Context, petID string, reminderID string, info domains.PeriodNotifyInfo) (err error)
}

type notificationService struct {
	repository NotificationRepositoryInterface
}

func (n *notificationService) addNew(ctx context.Context, notification domains.Notification) (err error) {
	err = n.repository.Create(ctx, notification)
	return
}

func (n *notificationService) removeAllTodayAndWaitStatusNotification(ctx context.Context, reminderID string) (err error) {
	err = n.repository.DeleteTodayAndWaitStatusByReminderID(ctx, reminderID)
	return
}

func (n *notificationService) AddNewFromHourReminder(ctx context.Context, petID string, reminderID string, info domains.HourNotifyInfo) (err error) {
	firstUsage := info.FirstUsage
	every := info.Every
	usagePerDay := (24 - firstUsage.Hour()) / every
	for i := 0; i <= usagePerDay; i++ {
		notification := domains.NewNotification(petID, reminderID, firstUsage.Add(time.Duration(i*every)*time.Hour))
		if err = n.addNew(ctx, notification); err != nil {
			break
		}
	}

	return
}

func (n *notificationService) AddNewFromPeriodReminder(ctx context.Context, petID string, reminderID string, info domains.PeriodNotifyInfo) (err error) {
	periods := []*time.Time{info.Morning, info.Noon, info.Evening, info.BeforeBed}

	for _, period := range periods {
		if period == nil {
			continue
		}

		notification := domains.NewNotification(petID, reminderID, *period)
		if err = n.addNew(ctx, notification); err != nil {
			break
		}
	}

	return
}

func (n *notificationService) ChangeNotifyTimeFromHourReminder(ctx context.Context, petID string, reminderID string, info domains.HourNotifyInfo) (err error) {
	err = n.removeAllTodayAndWaitStatusNotification(ctx, reminderID)
	if err != nil {
		return
	}

	err = n.AddNewFromHourReminder(ctx, petID, reminderID, info)
	return
}

func (n *notificationService) ChangeNotifyTimeFromPeriodReminder(ctx context.Context, petID string, reminderID string, info domains.PeriodNotifyInfo) (err error) {
	err = n.removeAllTodayAndWaitStatusNotification(ctx, reminderID)
	if err != nil {
		return
	}

	err = n.AddNewFromPeriodReminder(ctx, petID, reminderID, info)
	return
}

func NewNotificationService(repository NotificationRepositoryInterface) NotificationServiceInterface {
	return &notificationService{
		repository: repository,
	}
}
