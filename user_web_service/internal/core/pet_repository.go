package core

import "context"

type IPetRepository interface {
	ReadByID(ctx context.Context, id string) (pet IPet, err error)
	ReadByUserID(ctx context.Context, userID string) (pets []IPet, err error)
	Create(ctx context.Context, pet IPet) (err error)
	Update(ctx context.Context, pet IPet) (err error)
	DeleteByID(ctx context.Context, id string) (err error)
	DeleteByUserID(ctx context.Context, userID string) (err error)
}
