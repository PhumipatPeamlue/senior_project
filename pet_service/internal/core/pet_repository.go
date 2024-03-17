package core

import "context"

type PetRepositoryInterface interface {
	ReadByID(ctx context.Context, id string) (pet Pet, err error)
	ReadByUserID(ctx context.Context, userID string) (pets []Pet, err error)
	Create(ctx context.Context, pet Pet) (err error)
	Update(ctx context.Context, pet Pet) (err error)
	DeleteByID(ctx context.Context, id string) (err error)
	DeleteByUserID(ctx context.Context, userID string) (err error)
}
