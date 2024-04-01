package core

import (
	"context"
	"time"
)

type INotificationRecordService interface {
	addNew(ctx context.Context, notificationRecord INotificationRecord) (err error)
	removeAllTodayAndWaitStatusNotification(ctx context.Context, notificationID string) (err error)
	AddNewFromHourNotification(ctx context.Context, petID string, notificationID string, info HourNotifyInfo) (err error)
	AddNewFromPeriodNotification(ctx context.Context, petID string, notificationID string, info PeriodNotifyInfo) (err error)
	ChangeNotifyTimeFromHourNotification(ctx context.Context, petID string, notificationID string, info HourNotifyInfo) (err error)
	ChangeNotifyTimeFromPeriodNotification(ctx context.Context, petID string, notificationID string, info PeriodNotifyInfo) (err error)
}

type notificationRecordService struct {
	repository INotificationRecordRepository
}

// AddNewFromHourNotification implements INotificationRecordService.
func (n *notificationRecordService) AddNewFromHourNotification(ctx context.Context, petID string, notificationID string, info HourNotifyInfo) (err error) {
	firstUsage := info.FirstUsage
	every := info.Every
	usagePerDay := (24 - firstUsage.Hour()) / every
	for i := 0; i <= usagePerDay; i++ {
		notificationRecord := newNotificationRecord(petID, notificationID, firstUsage.Add(time.Duration(i*every)*time.Hour))
		if err = n.addNew(ctx, notificationRecord); err != nil {
			break
		}
	}

	return
}

// AddNewFromPeriodNotification implements INotificationRecordService.
func (n *notificationRecordService) AddNewFromPeriodNotification(ctx context.Context, petID string, notificationID string, info PeriodNotifyInfo) (err error) {
	periods := []*time.Time{info.Morning, info.Noon, info.Evening, info.BeforeBed}

	for _, period := range periods {
		if period == nil {
			continue
		}

		notificationRecord := newNotificationRecord(petID, notificationID, *period)
		if err = n.addNew(ctx, notificationRecord); err != nil {
			break
		}
	}

	return
}

// ChangeNotifyTimeFromHourNotification implements INotificationRecordService.
func (n *notificationRecordService) ChangeNotifyTimeFromHourNotification(ctx context.Context, petID string, notificationID string, info HourNotifyInfo) (err error) {
	err = n.removeAllTodayAndWaitStatusNotification(ctx, notificationID)
	if err != nil {
		return
	}

	err = n.AddNewFromHourNotification(ctx, petID, notificationID, info)
	return
}

// ChangeNotifyTimeFromPeriodNotification implements INotificationRecordService.
func (n *notificationRecordService) ChangeNotifyTimeFromPeriodNotification(ctx context.Context, petID string, notificationID string, info PeriodNotifyInfo) (err error) {
	err = n.removeAllTodayAndWaitStatusNotification(ctx, notificationID)
	if err != nil {
		return
	}

	err = n.AddNewFromPeriodNotification(ctx, petID, notificationID, info)
	return
}

// addNew implements INotificationRecordService.
func (n *notificationRecordService) addNew(ctx context.Context, notificationRecord INotificationRecord) (err error) {
	err = n.repository.Create(ctx, notificationRecord)
	return
}

// removeAllTodayAndWaitStatusNotification implements INotificationRecordService.
func (n *notificationRecordService) removeAllTodayAndWaitStatusNotification(ctx context.Context, notificationID string) (err error) {
	err = n.repository.DeleteTodayAndWaitStatusByNotificationID(ctx, notificationID)
	return
}

func NewNotificationRecordService(r INotificationRecordRepository) INotificationRecordService {
	return &notificationRecordService{
		repository: r,
	}
}
