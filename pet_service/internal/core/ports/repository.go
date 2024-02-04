package ports

import (
	"context"
	"pet_service/internal/core/domains"
)

type PetRepository interface {
	FindByID(ctx context.Context, id string) (pet domains.Pet, err error)
	FindByUserID(ctx context.Context, userID string) (pets []domains.Pet, err error)
	Save(ctx context.Context, pet domains.Pet) (err error)
	Update(ctx context.Context, pet domains.Pet) (err error)
	DeleteByID(ctx context.Context, id string) (err error)
}
