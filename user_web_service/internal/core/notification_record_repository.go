package core

import "context"

type INotificationRecordRepository interface {
	Create(ctx context.Context, notificationRecord INotificationRecord) (err error)
	DeleteTodayAndWaitStatusByNotificationID(ctx context.Context, notificationID string) (err error)
}
