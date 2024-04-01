package core

import "context"

type IPetService interface {
	FindPet(ctx context.Context, petID string) (pet IPet, err error)
	FindAllUserPets(ctx context.Context, userID string) (pets []IPet, err error)
	AddNewPet(ctx context.Context, userID, petName string) (petID string, err error)
	ChangePetName(ctx context.Context, petID, newName string) (err error)
	RemovePet(ctx context.Context, petID string) (err error)
	RemoveAllUserPets(ctx context.Context, userID string) (err error)
}

type petService struct {
	repository IPetRepository
}

// AddNewPet implements IPetService.
func (p *petService) AddNewPet(ctx context.Context, userID string, petName string) (petID string, err error) {
	pet := newPet(userID, petName)
	petID = pet.ID()
	err = p.repository.Create(ctx, pet)
	return
}

// ChangePetName implements IPetService.
func (p *petService) ChangePetName(ctx context.Context, petID string, newName string) (err error) {
	pet, err := p.FindPet(ctx, petID)
	if err != nil {
		return
	}

	pet.changeName(newName)
	err = p.repository.Update(ctx, pet)
	return err
}

// FindAllUserPets implements IPetService.
func (p *petService) FindAllUserPets(ctx context.Context, userID string) (pets []IPet, err error) {
	pets, err = p.repository.ReadByUserID(ctx, userID)
	return pets, err
}

// FindPet implements IPetService.
func (p *petService) FindPet(ctx context.Context, petID string) (pet IPet, err error) {
	pet, err = p.repository.ReadByID(ctx, petID)
	return pet, err
}

// RemoveAllUserPets implements IPetService.
func (p *petService) RemoveAllUserPets(ctx context.Context, userID string) (err error) {
	err = p.repository.DeleteByUserID(ctx, userID)
	return
}

// RemovePet implements IPetService.
func (p *petService) RemovePet(ctx context.Context, petID string) (err error) {
	err = p.repository.DeleteByID(ctx, petID)
	return
}

func NewPetService(r IPetRepository) IPetService {
	return &petService{
		repository: r,
	}
}
