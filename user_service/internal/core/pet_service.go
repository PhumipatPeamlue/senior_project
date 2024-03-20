package core

import "context"

type PetServiceInterface interface {
	DeleteUserPet(ctx context.Context, userID string) (err error)
}
