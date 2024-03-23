package core

import "context"

type PetServiceInterface interface {
	FindPet(ctx context.Context, petID string) (Pet, error)
	FindAllUserPets(ctx context.Context, userID string) ([]Pet, error)
	AddNewPet(ctx context.Context, userID, petName string) (string, error)
	ChangePetName(ctx context.Context, petID, newName string) error
	RemovePet(ctx context.Context, petID string) error
	RemoveAllUserPets(ctx context.Context, userID string) error
}

type petService struct {
	repository PetRepositoryInterface
}

func (p *petService) FindPet(ctx context.Context, petID string) (Pet, error) {
	pet, err := p.repository.ReadByID(ctx, petID)
	return pet, err
}

func (p *petService) FindAllUserPets(ctx context.Context, userID string) ([]Pet, error) {
	pets, err := p.repository.ReadByUserID(ctx, userID)
	return pets, err
}

func (p *petService) AddNewPet(ctx context.Context, userID string, petName string) (string, error) {
	pet := newPet(userID, petName)
	err := p.repository.Create(ctx, pet)
	return pet.id, err
}

func (p *petService) ChangePetName(ctx context.Context, petID string, newName string) error {
	pet, err := p.FindPet(ctx, petID)
	if err != nil {
		return err
	}

	pet.ChangeName(newName)
	err = p.repository.Update(ctx, pet)
	return err
}

func (p *petService) RemovePet(ctx context.Context, petID string) error {
	err := p.repository.DeleteByID(ctx, petID)
	return err
}

func (p *petService) RemoveAllUserPets(ctx context.Context, userID string) error {
	err := p.repository.DeleteByUserID(ctx, userID)
	return err
}

func NewPetService(repo PetRepositoryInterface) PetServiceInterface {
	return &petService{
		repository: repo,
	}
}
