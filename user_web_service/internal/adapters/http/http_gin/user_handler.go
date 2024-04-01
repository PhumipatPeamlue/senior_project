package http_gin

import (
	"context"
	"errors"
	"net/http"
	"time"
	"user_web_service/internal/core"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service core.IUserService
}

func (u *UserHandler) handleError(c *gin.Context, err error) {
	c.Error(err)

	if errors.Is(err, context.DeadlineExceeded) {
		c.JSON(http.StatusRequestTimeout, gin.H{"error": "request time out"})
		return
	}

	var errNotFound *core.ErrNotFound
	var errDuplicate *core.ErrDuplicate
	if errors.As(err, &errNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "pet not found"})
	} else if errors.As(err, &errDuplicate) {
		c.JSON(http.StatusConflict, gin.H{"error": "pet already exists"})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
	}
}

func (u *UserHandler) FindUserByLineUserID(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	lineUserID := c.Param("line_user_id")

	user, err := u.service.FindUserByLineUserID(ctx, lineUserID)
	if err != nil {
		var errNotFound *core.ErrNotFound
		if !errors.As(err, &errNotFound) {
			u.handleError(c, err)
			return
		}

		if err = u.service.AddNewUser(ctx, lineUserID); err != nil {
			u.handleError(c, err)
			return
		}

		user, err = u.service.FindUserByLineUserID(ctx, lineUserID)
		if err != nil {
			u.handleError(c, err)
			return
		}
	}

	c.JSON(http.StatusOK, userResponse{
		UserID:    user.ID(),
		Morning:   user.TimeSetting().Morning(),
		Noon:      user.TimeSetting().Noon(),
		Evening:   user.TimeSetting().Evening(),
		BeforeBed: user.TimeSetting().BeforeBed(),
	})
}

func (u *UserHandler) ChangeTimeSetting(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var body changeTimeSettingRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "json decoding: " + err.Error()})
		return
	}

	timeSetting := core.ScanUserTimeSetting(body.Morning, body.Noon, body.Evening, body.BeforeBed)
	if err := u.service.ChangeTimeSetting(ctx, body.LineUserID, timeSetting); err != nil {
		u.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "change time setting successfully"})
}

func (u *UserHandler) RemoveUserByLineUserID(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	lineUserID := c.Param("line_user_id")

	if err := u.service.RemoveUserByLineUserID(ctx, lineUserID); err != nil {
		u.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "remove the user successfully"})
}

func NewUserHandler(s core.IUserService) *UserHandler {
	return &UserHandler{
		service: s,
	}
}
