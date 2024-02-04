package ports

import (
	"context"
	"user_service/internal/core/domains"
)

type UserRepository interface {
	FindByID(ctx context.Context, id string) (user domains.User, err error)
	Save(ctx context.Context, user domains.User) (err error)
	Update(ctx context.Context, user domains.User) (err error)
	DeleteByID(ctx context.Context, id string) (err error)
}
