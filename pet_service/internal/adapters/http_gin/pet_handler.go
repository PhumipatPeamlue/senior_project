package http_gin

import (
	"context"
	"errors"
	"net/http"
	"pet_service/internal/core"
	"time"

	"github.com/gin-gonic/gin"
)

type petHandler struct {
	service core.PetServiceInterface
}

func (p *petHandler) handleError(c *gin.Context, err error) {
	c.Error(err)

	if errors.Is(err, context.DeadlineExceeded) {
		c.JSON(http.StatusRequestTimeout, gin.H{"error": "request timeout"})
		return
	}

	var errPetNotFound *core.ErrPetNotFound
	var errPetDuplicate *core.ErrPetDuplicate
	switch {
	case errors.As(err, &errPetNotFound):
		c.JSON(http.StatusNotFound, gin.H{"error": "pet not found"})
	case errors.As(err, &errPetDuplicate):
		c.JSON(http.StatusConflict, gin.H{"error": "pet already exists"})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
	}
}

func (p *petHandler) AddNewPet(c *gin.Context) {
	var body addNewPetRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	petID, err := p.service.AddNewPet(ctx, body.UserID, body.Name)
	if err != nil {
		p.handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "add a new pet successfully",
		"pet_id":  petID,
	})
}

func (p *petHandler) ChangePetName(c *gin.Context) {
	var body changePetNameRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := p.service.ChangePetName(ctx, body.PetID, body.Name)
	if err != nil {
		p.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "change the name successfully"})
}

func (p *petHandler) FindAllUserPets(c *gin.Context) {
	userID := c.Param("user_id")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pets, err := p.service.FindAllUserPets(ctx, userID)
	if err != nil {
		p.handleError(c, err)
		return
	}

	var userPets []petResponse
	for _, pet := range pets {
		pr := petResponse{
			PetID:  pet.ID(),
			UserID: pet.UserID(),
			Name:   pet.Name(),
		}
		userPets = append(userPets, pr)
	}

	c.JSON(http.StatusOK, userPets)
}

func (p *petHandler) FindPet(c *gin.Context) {
	petID := c.Param("pet_id")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pet, err := p.service.FindPet(ctx, petID)
	if err != nil {
		p.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, petResponse{
		PetID:  pet.ID(),
		UserID: pet.UserID(),
		Name:   pet.Name(),
	})
}

func (p *petHandler) RemoveAllUserPets(c *gin.Context) {
	userID := c.Param("user_id")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := p.service.RemoveAllUserPets(ctx, userID)
	if err != nil {
		p.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "remove all user's pets successfully"})
}

func (p *petHandler) RemovePet(c *gin.Context) {
	petID := c.Param("pet_id")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := p.service.RemovePet(ctx, petID)
	if err != nil {
		p.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "remove the pet successfully"})
}

func NewPetHandler(service core.PetServiceInterface) *petHandler {
	return &petHandler{
		service: service,
	}
}
