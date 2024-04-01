package core

import "context"

type iHourNotificationRepository interface {
	ReadHourNotificationByID(ctx context.Context, id string) (hn IHourNotification, err error)
	CreateHourNotification(ctx context.Context, hn IHourNotification) (err error)
	UpdateHourNotification(ctx context.Context, hn IHourNotification) (err error)
}

type iPeriodNotificationRepository interface {
	ReadPeriodNotificationByID(ctx context.Context, id string) (pn IPeriodNotification, err error)
	CreatePeriodNotification(ctx context.Context, pn IPeriodNotification) (err error)
	UpdatePeriodNotification(ctx context.Context, pn IPeriodNotification) (err error)
}

type INotificationRepository interface {
	iHourNotificationRepository
	iPeriodNotificationRepository
	ReadByID(ctx context.Context, id string) (notification INotification, err error)
	ReadByPetID(ctx context.Context, petID string) (notifications []INotification, err error)
	DeleteByPetID(ctx context.Context, petID string) (err error)
	DeleteByID(ctx context.Context, id string) (err error)
}
