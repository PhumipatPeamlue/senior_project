package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"user_service/internal/core/domains"
	"user_service/internal/core/ports"
)

type UserTimeSettingsResponse struct {
	UserID string `json:"user_id"`
	domains.UserTimeSetting
}

type UpdateUserTimeSettingsRequest struct {
	UserID string `json:"user_id"`
	domains.UserTimeSetting
}

type UserHandler struct {
	service ports.UserService
}

func NewUserHandler(service ports.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) GetTimeSettingHandler(c *gin.Context) {
	userID := c.Param("user_id")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	alreadyExists, user, err := h.service.CheckUserExists(ctx, userID)
	if err != nil {
		c.Error(err)
		return
	}

	if !alreadyExists {
		user, err = h.service.AddNewUser(ctx, userID)
		if err != nil {
			c.Error(err)
			return
		}
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) ChangeTimeSettingsHandler(c *gin.Context) {
	var body UpdateUserTimeSettingsRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	userID := body.UserID
	uts := domains.UserTimeSetting{
		Morning:   body.Morning,
		Noon:      body.Noon,
		Evening:   body.Evening,
		BeforeBed: body.BeforeBed,
	}
	if err := h.service.ChangeTimeSettings(ctx, userID, uts); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "change user's time settings successfully"})
}

func (h *UserHandler) RemoveUserDataByIDHandler(c *gin.Context) {
	userID := c.Param("user_id")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := h.service.RemoveUserDataByID(ctx, userID); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "remove user's data successfully"})
}
