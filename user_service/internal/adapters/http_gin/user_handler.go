package http_gin

import (
	"context"
	"errors"
	"net/http"
	"time"
	"user_service/internal/core"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service core.UserServiceInterface
}

func (u *UserHandler) handleError(c *gin.Context, err error) {
	c.Error(err)

	if errors.Is(err, context.DeadlineExceeded) {
		c.JSON(http.StatusRequestTimeout, gin.H{"error": "request time out"})
		return
	}

	var errUserNotFound *core.ErrUserNotFound
	var errUserDuplicate *core.ErrUserDuplicate
	if errors.As(err, &errUserNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
	} else if errors.As(err, &errUserDuplicate) {
		c.JSON(http.StatusConflict, gin.H{"error": "user already exists"})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
	}
}

func (u *UserHandler) FindUserByLineUserID(c *gin.Context) {
	lineUserID := c.Param("line_user_id")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user, err := u.service.FindUserByLineUserID(ctx, lineUserID)
	if err != nil {
		var errUserNotFound *core.ErrUserNotFound
		if errors.As(err, &errUserNotFound) {
			err = nil
		} else {
			u.handleError(c, err)
			return
		}
	}

	if user.CreatedAt().IsZero() {
		err = u.service.AddNewUser(ctx, lineUserID)
		if err != nil {
			u.handleError(c, err)
			return
		}

		user, err = u.service.FindUserByLineUserID(ctx, lineUserID)
		if err != nil {
			u.handleError(c, err)
			return
		}
	}

	c.JSON(http.StatusOK, newUserResponse(user))
}

func (u *UserHandler) AddNewUser(c *gin.Context) {
	lineUserID := c.Param("line_user_id")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := u.service.AddNewUser(ctx, lineUserID)
	if err != nil {
		u.handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "add a new user successfully"})
}

func (u *UserHandler) ChangeTimeSetting(c *gin.Context) {
	var body changeTimeSettingRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "json decoding: " + err.Error()})
		return
	}

	newTimeSetting := core.ScanUserTimeSetting(body.Morning, body.Noon, body.Evening, body.BeforeBed)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := u.service.ChangeTimeSetting(ctx, body.LineUserID, newTimeSetting)
	if err != nil {
		u.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "change time setting successfully"})
}

func (u *UserHandler) RemoveUserByLineUserID(c *gin.Context) {
	lineUserID := c.Param("line_user_id")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := u.service.RemoveUserByLineUserID(ctx, lineUserID)
	if err != nil {
		u.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "remove the user successfully"})
}

func NewUserHandler(service core.UserServiceInterface) *UserHandler {
	return &UserHandler{
		service: service,
	}
}
