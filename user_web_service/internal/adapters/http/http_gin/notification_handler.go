package http_gin

import (
	"context"
	"errors"
	"net/http"
	"time"
	"user_web_service/internal/core"

	"github.com/gin-gonic/gin"
)

type NotificationHandler struct {
	service core.INotificationService
}

func (r *NotificationHandler) handleError(c *gin.Context, err error) {
	c.Error(err)

	if errors.Is(err, context.DeadlineExceeded) {
		c.JSON(http.StatusRequestTimeout, gin.H{"error": "request timeout"})
		return
	}

	var errReminderNotFound *core.ErrNotFound
	var errReminderDuplicate *core.ErrDuplicate
	switch {
	case errors.As(err, &errReminderNotFound):
		c.JSON(http.StatusNotFound, gin.H{"error": "reminder not found"})
	case errors.As(err, &errReminderDuplicate):
		c.JSON(http.StatusConflict, gin.H{"error": "reminder already exists"})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
	}
}

func (r *NotificationHandler) AddNewHourNotification(c *gin.Context) {
	var body addNewHourNotificationRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.service.AddNewHourNotification(ctx, body.PetID, body.DrugInfo, body.Frequency, body.HourNotifyInfo)
	if err != nil {
		r.handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "add a new hour notification successfully"})
}

func (r *NotificationHandler) AddNewPeriodNotification(c *gin.Context) {
	var body addNewPeriodNotificationRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()

	err := r.service.AddNewPeriodNotification(ctx, body.PetID, body.DrugInfo, body.Frequency, body.PeriodNotifyInfo)
	if err != nil {
		r.handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "add a new period notification successfully"})
}

func (r *NotificationHandler) ChangeHourNotificationInfo(c *gin.Context) {
	var body changeHourDrugLabelInfoRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.service.ChangeHourNotificationInfo(ctx, body.NotificationID, body.DrugInfo, body.Frequency, body.HourNotifyInfo)
	if err != nil {
		r.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "change hour notification's information successfully"})
}

func (r *NotificationHandler) ChangePeriodNotificationInfo(c *gin.Context) {
	var body changePeriodDrugLabelInfoRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()

	err := r.service.ChangePeriodNotificationInfo(ctx, body.NotificationID, body.DrugInfo, body.Frequency, body.PeriodNotifyInfo)
	if err != nil {
		r.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "change period notification's information successfully"})
}

func (r *NotificationHandler) FindAllPetNotifications(c *gin.Context) {
	petID := c.Param("pet_id")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	notifications, err := r.service.FindAllPetNotifications(ctx, petID)
	if err != nil {
		r.handleError(c, err)
		return
	}

	var petNotifications []notificationResponse
	for _, notification := range notifications {
		rr := notificationResponse{
			NotificationID:   notification.ID(),
			PetID:            notification.PetID(),
			NotificationType: notification.Type(),
			DrugInfo:         notification.DrugInfo(),
		}
		petNotifications = append(petNotifications, rr)
	}

	c.JSON(http.StatusOK, petNotifications)
}

func (r *NotificationHandler) FindHourNotification(c *gin.Context) {
	notificationID := c.Param("reminder_id")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	notification, err := r.service.FindHourNotification(ctx, notificationID)
	if err != nil {
		r.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, hourNotificationResponse{
		NotificationID: notification.ID(),
		PetID:          notification.PetID(),
		Frequency:      notification.FrequencyDayUsage(),
		DrugInfo:       notification.DrugInfo(),
		HourNotifyInfo: notification.NotifyInfo(),
	})
}

func (r *NotificationHandler) FindPeriodNotification(c *gin.Context) {
	notificationID := c.Param("reminder_id")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	reminder, err := r.service.FindPeriodNotification(ctx, notificationID)
	if err != nil {
		r.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, periodNotificationResponse{
		NotificationID:   reminder.ID(),
		PetID:            reminder.PetID(),
		Frequency:        reminder.FrequencyDayUsage(),
		DrugInfo:         reminder.DrugInfo(),
		PeriodNotifyInfo: reminder.NotifyInfo(),
	})
}

func (r *NotificationHandler) RemoveAllPetNotifications(c *gin.Context) {
	petID := c.Param("pet_id")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.service.RemoveAllPetNotifications(ctx, petID)
	if err != nil {
		r.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "remove all pet's notifications successfully"})
}

func (r *NotificationHandler) RemoveHourNotification(c *gin.Context) {
	reminderID := c.Param("reminder_id")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.service.RemoveHourNotification(ctx, reminderID)
	if err != nil {
		r.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "remove hour notification successfully"})
}

func (r *NotificationHandler) RemovePeriodNotification(c *gin.Context) {
	reminderID := c.Param("reminder_id")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.service.RemovePeriodNotification(ctx, reminderID)
	if err != nil {
		r.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "remove period notification successfully"})
}

func (r *NotificationHandler) FindNotification(c *gin.Context) {
	notificationID := c.Param("reminder_id")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	notification, err := r.service.FindNotification(ctx, notificationID)
	if err != nil {
		r.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, notificationResponse{
		NotificationID:   notificationID,
		PetID:            notification.PetID(),
		NotificationType: notification.Type(),
		DrugInfo:         notification.DrugInfo(),
	})
}

func NewNotificationHandler(service core.INotificationService) *NotificationHandler {
	return &NotificationHandler{
		service: service,
	}
}
