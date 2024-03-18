package http_gin

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"reminder_service/internal/core"
	"time"
)

type reminderHandler struct {
	service core.ReminderServiceInterface
}

func (r *reminderHandler) handleError(c *gin.Context, err error) {
	c.Error(err)

	if errors.Is(err, context.DeadlineExceeded) {
		c.JSON(http.StatusRequestTimeout, gin.H{"error": "request timeout"})
		return
	}

	var errReminderNotFound *core.ErrReminderNotFound
	var errReminderDuplicate *core.ErrReminderDuplicate
	switch {
	case errors.As(err, &errReminderNotFound):
		c.JSON(http.StatusNotFound, gin.H{"error": "reminder not found"})
	case errors.As(err, &errReminderDuplicate):
		c.JSON(http.StatusConflict, gin.H{"error": "reminder already exists"})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
	}
}

func (r *reminderHandler) AddNewHourReminder(c *gin.Context) {
	var body addNewHourReminderRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.service.AddNewHourReminder(ctx, body.PetID, body.DrugInfo, body.Frequency, body.HourNotifyInfo)
	if err != nil {
		r.handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "add a new hour reminder successfully"})
}

func (r *reminderHandler) AddNewPeriodReminder(c *gin.Context) {
	var body addNewPeriodReminderRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()

	err := r.service.AddNewPeriodReminder(ctx, body.PetID, body.DrugInfo, body.Frequency, body.PeriodNotifyInfo)
	if err != nil {
		r.handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "add a new period reminder successfully"})
}

func (r *reminderHandler) ChangeHourReminderInfo(c *gin.Context) {
	var body changeHourDrugLabelInfoRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.service.ChangeHourReminderInfo(ctx, body.ReminderID, body.DrugInfo, body.Frequency, body.HourNotifyInfo)
	if err != nil {
		r.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "change hour reminder's information successfully"})
}

func (r *reminderHandler) ChangePeriodReminderInfo(c *gin.Context) {
	var body changePeriodDrugLabelInfoRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()

	err := r.service.ChangePeriodReminderInfo(ctx, body.ReminderID, body.DrugInfo, body.Frequency, body.PeriodNotifyInfo)
	if err != nil {
		r.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "change period reminder's information successfully"})
}

func (r *reminderHandler) FindAllPetReminders(c *gin.Context) {
	petID := c.Param("pet_id")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	reminders, err := r.service.FindAllPetReminders(ctx, petID)
	if err != nil {
		r.handleError(c, err)
		return
	}

	var petReminders []reminderResponse
	for _, reminder := range reminders {
		rr := reminderResponse{
			ReminderID:   reminder.ID(),
			PetID:        reminder.PetID(),
			ReminderType: reminder.Type(),
			DrugInfo:     reminder.DrugInfo(),
		}
		petReminders = append(petReminders, rr)
	}

	c.JSON(http.StatusOK, petReminders)
}

func (r *reminderHandler) FindHourReminder(c *gin.Context) {
	reminderID := c.Param("reminder_id")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	reminder, err := r.service.FindHourReminder(ctx, reminderID)
	if err != nil {
		r.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, hourReminderResponse{
		ReminderID:     reminder.ID(),
		PetID:          reminder.PetID(),
		Frequency:      reminder.FrequencyDayUsage(),
		DrugInfo:       reminder.DrugInfo(),
		HourNotifyInfo: reminder.NotifyInfo(),
	})
}

func (r *reminderHandler) FindPeriodReminder(c *gin.Context) {
	reminderID := c.Param("reminder_id")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	reminder, err := r.service.FindPeriodReminder(ctx, reminderID)
	if err != nil {
		r.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, periodReminderResponse{
		ReminderID:       reminder.ID(),
		PetID:            reminder.PetID(),
		Frequency:        reminder.FrequencyDayUsage(),
		DrugInfo:         reminder.DrugInfo(),
		PeriodNotifyInfo: reminder.NotifyInfo(),
	})
}

func (r *reminderHandler) RemoveAllPetReminders(c *gin.Context) {
	petID := c.Param("pet_id")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.service.RemoveAllPetReminders(ctx, petID)
	if err != nil {
		r.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "remove all pet's reminders successfully"})
}

func (r *reminderHandler) RemoveHourReminder(c *gin.Context) {
	reminderID := c.Param("reminder_id")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.service.RemoveHourReminder(ctx, reminderID)
	if err != nil {
		r.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "remove hour reminder successfully"})
}

func (r *reminderHandler) RemovePeriodReminder(c *gin.Context) {
	reminderID := c.Param("reminder_id")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.service.RemovePeriodReminder(ctx, reminderID)
	if err != nil {
		r.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "remove period reminder successfully"})
}

func (r *reminderHandler) FindReminder(c *gin.Context) {
	reminderID := c.Param("reminder_id")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	reminder, err := r.service.FindReminder(ctx, reminderID)
	if err != nil {
		r.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, reminderResponse{
		ReminderID:   reminderID,
		PetID:        reminder.PetID(),
		ReminderType: reminder.Type(),
		DrugInfo:     reminder.DrugInfo(),
	})
}

func NewReminderHandler(service core.ReminderServiceInterface) *reminderHandler {
	return &reminderHandler{
		service: service,
	}
}
