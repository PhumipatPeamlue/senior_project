package services

import (
	"context"
	"github.com/google/uuid"
	"pet_service/internal/core/domains"
	"pet_service/internal/core/ports"
)

type petService struct {
	repo ports.PetRepository
}

func (s *petService) GetPetInfo(ctx context.Context, petID string) (pet domains.Pet, err error) {
	pet, err = s.repo.FindByID(ctx, petID)
	return
}

func (s *petService) GetAllUserPet(ctx context.Context, userID string) (pets []domains.Pet, err error) {
	pets, err = s.repo.FindByUserID(ctx, userID)
	return
}

func (s *petService) AddNewPet(ctx context.Context, userID string, name string) (err error) {
	pet := domains.Pet{
		ID:     uuid.New().String(),
		UserID: userID,
		Name:   name,
	}
	err = s.repo.Save(ctx, pet)
	return
}

func (s *petService) ChangePetInfo(ctx context.Context, petID string, name string) (err error) {
	pet, err := s.repo.FindByID(ctx, petID)
	if err != nil {
		return
	}

	pet.Name = name
	err = s.repo.Update(ctx, pet)
	return
}

func (s *petService) RemovePet(ctx context.Context, petID string) (err error) {
	err = s.repo.DeleteByID(ctx, petID)
	return
}

func NewPetService(repo ports.PetRepository) ports.PetService {
	return &petService{
		repo: repo,
	}
}
