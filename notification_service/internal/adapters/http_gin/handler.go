package http_gin

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"notification_service/internal/core"
	"time"
)

type notificationHandler struct {
	service core.NotificationServiceInterface
}

func (n *notificationHandler) handleError(c *gin.Context, err error) {
	c.Error(err)

	if errors.Is(err, context.DeadlineExceeded) {
		c.JSON(http.StatusRequestTimeout, gin.H{"error": "request timeout"})
		return
	}

	var errNotificationDuplicate *core.ErrNotificationDuplicate
	switch {
	case errors.As(err, &errNotificationDuplicate):
		c.JSON(http.StatusConflict, gin.H{"error": "notification already exists"})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
	}
}

func (n *notificationHandler) AddNewFromHourReminder(c *gin.Context) {
	var body addNewFromHourReminderRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := n.service.AddNewFromHourReminder(ctx, body.PetID, body.ReminderID, body.HourNotifyInfo); err != nil {
		n.handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "add a new notification from hour reminder successfully"})
}

func (n *notificationHandler) AddNewFromPeriodReminder(c *gin.Context) {
	var body addNewFromPeriodReminderRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := n.service.AddNewFromPeriodReminder(ctx, body.PetID, body.ReminderID, body.PeriodNotifyInfo); err != nil {
		n.handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "add a new notification from period reminder successfully"})
}

func (n *notificationHandler) ChangeNotifyTimeFromHourReminder(c *gin.Context) {
	var body changeNotifyTimeFromHourReminderRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := n.service.ChangeNotifyTimeFromHourReminder(ctx, body.PetID, body.ReminderID, body.HourNotifyInfo); err != nil {
		n.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "change the notify time from hour reminder successfully"})
}

func (n *notificationHandler) ChangeNotifyTimeFromPeriodReminder(c *gin.Context) {
	var body changeNotifyTimeFromPeriodReminderRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := n.service.ChangeNotifyTimeFromPeriodReminder(ctx, body.PetID, body.ReminderID, body.PeriodNotifyInfo); err != nil {
		n.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "change the notify time from period reminder successfully"})
}

func NewNotificationHandler(service core.NotificationServiceInterface) *notificationHandler {
	return &notificationHandler{
		service: service,
	}
}
