package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"pet_service/internal/core/ports"
	"time"
)

type AddNewPetRequest struct {
	UserID string `json:"user_id" binding:"required"`
	Name   string `json:"name" binding:"required"`
}

type ChangePetInfoRequest struct {
	PetID string `json:"pet_id" binding:"required"`
	Name  string `json:"name"`
}

type PetHandler struct {
	service ports.PetService
}

func (h *PetHandler) GetPetInfoHandler(c *gin.Context) {
	petID := c.Param("pet_id")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pet, err := h.service.GetPetInfo(ctx, petID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, pet)
}

func (h *PetHandler) GetAllUserPetHandler(c *gin.Context) {
	userID := c.Param("user_id")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pets, err := h.service.GetAllUserPet(ctx, userID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, pets)
}

func (h *PetHandler) AddNewPetHandler(c *gin.Context) {
	var body AddNewPetRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := h.service.AddNewPet(ctx, body.UserID, body.Name); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "add new pet successfully"})
}

func (h *PetHandler) ChangePetInfoHandler(c *gin.Context) {
	var body ChangePetInfoRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := h.service.ChangePetInfo(ctx, body.PetID, body.Name); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "change pet's information successfully"})
}

func (h *PetHandler) RemovePetHandler(c *gin.Context) {
	petID := c.Param("pet_id")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := h.service.RemovePet(ctx, petID); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "remove pet successfully"})
}

func NewPetHandler(service ports.PetService) *PetHandler {
	return &PetHandler{
		service: service,
	}
}
