package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"reminder_service/internal/core/ports"
	"time"
)

type AddNewPeriodReminderRequest struct {
	UserID    string     `json:"user_id" binding:"required"`
	PetID     string     `json:"pet_id" binding:"required"`
	DrugName  string     `json:"drug_name" binding:"required"`
	DrugUsage string     `json:"drug_usage" binding:"required"`
	Frequency string     `json:"frequency" binding:"required"`
	Morning   *time.Time `json:"morning"`
	Noon      *time.Time `json:"noon"`
	Evening   *time.Time `json:"evening"`
	BeforeBed *time.Time `json:"before_bed"`
}

type AddNewHourReminderRequest struct {
	UserID     string    `json:"user_id" binding:"required"`
	PetID      string    `json:"pet_id" binding:"required"`
	DrugName   string    `json:"drug_name" binding:"required"`
	DrugUsage  string    `json:"drug_usage" binding:"required"`
	Frequency  string    `json:"frequency" binding:"required"`
	FirstUsage time.Time `json:"first_usage" binding:"required"`
	Every      int       `json:"every" binding:"required"`
}

type ChangePeriodReminderInfoRequest struct {
	UserID     string     `json:"user_id" binding:"required"`
	ReminderID string     `json:"reminder_id" binding:"required"`
	DrugName   string     `json:"drug_name"`
	DrugUsage  string     `json:"drug_usage"`
	Frequency  string     `json:"frequency"`
	Morning    *time.Time `json:"morning"`
	Noon       *time.Time `json:"noon"`
	Evening    *time.Time `json:"evening"`
	BeforeBed  *time.Time `json:"before_bed"`
}

type ChangeHourReminderInfoRequest struct {
	UserID     string    `json:"user_id" binding:"required"`
	ReminderID string    `json:"reminder_id" binding:"required"`
	DrugName   string    `json:"drug_name"`
	DrugUsage  string    `json:"drug_usage"`
	Frequency  string    `json:"frequency"`
	FirstUsage time.Time `json:"first_usage"`
	Every      int       `json:"every"`
}

type ReminderHandler struct {
	service ports.ReminderService
}

func (h *ReminderHandler) GetAllRemindersHandler(c *gin.Context) {
	petID := c.Param("pet_id")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	reminders, err := h.service.GetAllPetReminders(ctx, petID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, reminders)
}

func (h *ReminderHandler) GetPeriodReminderInfoHandler(c *gin.Context) {
	reminderID := c.Param("reminder_id")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	reminder, err := h.service.GetPeriodReminderInfo(ctx, reminderID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, reminder)
}

func (h *ReminderHandler) GetHourReminderInfoHandler(c *gin.Context) {
	reminderID := c.Param("reminder_id")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	reminder, err := h.service.GetHourReminderInfo(ctx, reminderID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, reminder)
}

func (h *ReminderHandler) AddNewPeriodReminderHandler(c *gin.Context) {
	var body AddNewPeriodReminderRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := h.service.AddNewPeriodReminder(ctx, body.UserID, body.PetID, body.DrugName, body.DrugUsage, body.Frequency, body.Morning, body.Noon, body.Evening, body.BeforeBed); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "add new period reminder successfully"})
}

func (h *ReminderHandler) AddNewHourReminderHandler(c *gin.Context) {
	var body AddNewHourReminderRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := h.service.AddNewHourReminder(ctx, body.UserID, body.PetID, body.DrugName, body.DrugUsage, body.Frequency, body.FirstUsage, body.Every); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "add new hour reminder successfully"})
}

func (h *ReminderHandler) ChangePeriodReminderInfoHandler(c *gin.Context) {
	var body ChangePeriodReminderInfoRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := h.service.ChangePeriodReminderInfo(ctx, body.UserID, body.ReminderID, body.DrugName, body.DrugUsage, body.Frequency, body.Morning, body.Noon, body.Evening, body.BeforeBed); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "change period reminder's information successfully"})
}

func (h *ReminderHandler) ChangeHourReminderInfoHandler(c *gin.Context) {
	var body ChangeHourReminderInfoRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := h.service.ChangeHourReminderInfo(ctx, body.UserID, body.ReminderID, body.DrugName, body.DrugUsage, body.Frequency, body.FirstUsage, body.Every); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "change hour reminder's information successfully"})
}

func (h *ReminderHandler) RemoveHandler(c *gin.Context) {
	reminderID := c.Param("reminder_id")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := h.service.Remove(ctx, reminderID); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "remove reminder successfully"})
}

func NewReminderHandler(s ports.ReminderService) *ReminderHandler {
	return &ReminderHandler{
		service: s,
	}
}
