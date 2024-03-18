package core

import (
	"context"
	"cronjob/internal/core/domains"
)

type NotificationRepositoryInterface interface {
	Create(ctx context.Context, notification domains.Notification) (err error)
	ReadByWaitStatus(ctx context.Context) (notifications []domains.Notification, err error)
	Update(ctx context.Context, notification domains.Notification) (err error)
}
