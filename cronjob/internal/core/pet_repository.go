package core

import (
	"context"
	"cronjob/internal/core/domains"
)

type PetRepositoryInterface interface {
	ReadByID(ctx context.Context, id string) (pet domains.Pet, err error)
}
