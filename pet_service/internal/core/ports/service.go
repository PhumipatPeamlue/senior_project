package ports

import (
	"context"
	"pet_service/internal/core/domains"
)

type PetService interface {
	GetPetInfo(ctx context.Context, petID string) (pet domains.Pet, err error)
	GetAllUserPet(ctx context.Context, userID string) (pets []domains.Pet, err error)
	AddNewPet(ctx context.Context, userID string, name string) (err error)
	ChangePetInfo(ctx context.Context, petID string, name string) (err error)
	RemovePet(ctx context.Context, petID string) (err error)
}
