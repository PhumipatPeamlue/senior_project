package core

import "context"

type IUserRepository interface {
	ReadByID(ctx context.Context, id string) (user IUser, err error)
	Create(ctx context.Context, user IUser) (err error)
	Update(ctx context.Context, user IUser) (err error)
	DeleteByID(ctx context.Context, id string) (err error)
}
