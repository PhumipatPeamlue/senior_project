package ports

import (
	"context"
	"user_service/internal/core/domains"
)

type UserService interface {
	GetTimeSettingByID(ctx context.Context, userID string) (uts domains.UserTimeSetting, err error)
	CheckUserExists(ctx context.Context, userID string) (exists bool, user domains.User, err error)
	AddNewUser(ctx context.Context, userID string) (user domains.User, err error)
	ChangeTimeSettings(ctx context.Context, userID string, newUts domains.UserTimeSetting) (err error)
	RemoveUserDataByID(ctx context.Context, userID string) (err error)
}
