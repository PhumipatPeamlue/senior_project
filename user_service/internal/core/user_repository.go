package core

import "context"

type UserRepositoryInterface interface {
	ReadByID(ctx context.Context, lineUserID string) (user User, err error)
	Create(ctx context.Context, user User) (err error)
	Update(ctx context.Context, user User) (err error)
	DeleteByID(ctx context.Context, lineUserID string) (err error)
}
