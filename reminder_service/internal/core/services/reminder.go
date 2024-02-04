package services

import (
	"context"
	"github.com/google/uuid"
	"reminder_service/internal/core/domains"
	"reminder_service/internal/core/ports"
	"time"
)

type reminderService struct {
	reminderRepository         ports.ReminderRepository
	periodReminderRepository   ports.PeriodReminderRepository
	hourReminderRepoRepository ports.HourReminderRepository

	notificationService ports.NotificationService
}

func (s *reminderService) GetAllPetReminders(ctx context.Context, petID string) (reminders []domains.Reminder, err error) {
	reminders, err = s.reminderRepository.FindByPetID(ctx, petID)
	return
}

func (s *reminderService) GetPeriodReminderInfo(ctx context.Context, reminderID string) (reminder domains.PeriodReminder, err error) {
	reminder, err = s.periodReminderRepository.FindByID(ctx, reminderID)
	return
}

func (s *reminderService) GetHourReminderInfo(ctx context.Context, reminderID string) (reminder domains.HourReminder, err error) {
	reminder, err = s.hourReminderRepoRepository.FindByID(ctx, reminderID)
	return
}

func (s *reminderService) AddNewPeriodReminder(ctx context.Context, userID string, petID string, drugName string, drugUsage string, frequency string, morning *time.Time, noon *time.Time, evening *time.Time, beforeBed *time.Time) (err error) {
	reminderID := uuid.New().String()
	pr := domains.PeriodReminder{
		Reminder: domains.Reminder{
			ID:        reminderID,
			PetID:     petID,
			Type:      "period",
			DrugName:  drugName,
			DrugUsage: drugUsage,
			Frequency: frequency,
		},
		Morning:   morning,
		Noon:      noon,
		Evening:   evening,
		BeforeBed: beforeBed,
	}
	if err = s.periodReminderRepository.Save(ctx, pr); err != nil {
		return
	}

	err = s.notificationService.AddNewFromPeriodReminder(ctx, reminderID, userID, morning, noon, evening, beforeBed)
	return
}

func (s *reminderService) AddNewHourReminder(ctx context.Context, userID string, petID string, drugName string, drugUsage string, frequency string, firstUsage time.Time, every int) (err error) {
	reminderID := uuid.New().String()
	hr := domains.HourReminder{
		Reminder: domains.Reminder{
			ID:        reminderID,
			PetID:     petID,
			Type:      "hour",
			DrugName:  drugName,
			DrugUsage: drugUsage,
			Frequency: frequency,
		},
		FirstUsage: firstUsage,
		Every:      every,
	}
	if err = s.hourReminderRepoRepository.Save(ctx, hr); err != nil {
		return
	}

	err = s.notificationService.AddNewFromHourReminder(ctx, reminderID, userID, firstUsage, every)
	return
}

func (s *reminderService) ChangePeriodReminderInfo(ctx context.Context, userID string, reminderID string, drugName string, drugUsage string, frequency string, morning *time.Time, noon *time.Time, evening *time.Time, beforeBed *time.Time) (err error) {
	pr, err := s.periodReminderRepository.FindByID(ctx, reminderID)
	if err != nil {
		return
	}
	pr.DrugName = drugName
	pr.DrugUsage = drugUsage
	pr.Frequency = frequency
	pr.Morning = morning
	pr.Noon = noon
	pr.Evening = evening
	pr.BeforeBed = beforeBed
	if err = s.periodReminderRepository.Update(ctx, pr); err != nil {
		return
	}

	if err = s.notificationService.RemoveByReminderID(ctx, reminderID); err != nil {
		return
	}

	err = s.notificationService.AddNewFromPeriodReminder(ctx, reminderID, userID, morning, noon, evening, beforeBed)
	return
}

func (s *reminderService) ChangeHourReminderInfo(ctx context.Context, userID string, reminderID string, drugName string, drugUsage string, frequency string, firstUsage time.Time, every int) (err error) {
	hr, err := s.hourReminderRepoRepository.FindByID(ctx, reminderID)
	if err != nil {
		return
	}
	hr.DrugName = drugName
	hr.DrugUsage = drugUsage
	hr.Frequency = frequency
	hr.FirstUsage = firstUsage
	hr.Every = every
	if err = s.hourReminderRepoRepository.Update(ctx, hr); err != nil {
		return
	}

	if err = s.notificationService.RemoveByReminderID(ctx, reminderID); err != nil {
		return
	}

	err = s.notificationService.AddNewFromHourReminder(ctx, reminderID, userID, firstUsage, every)
	return
}

func (s *reminderService) Remove(ctx context.Context, reminderID string) (err error) {
	err = s.reminderRepository.Delete(ctx, reminderID)
	return
}

func NewReminderService(rr ports.ReminderRepository, prr ports.PeriodReminderRepository, hrr ports.HourReminderRepository, ns ports.NotificationService) ports.ReminderService {
	return &reminderService{
		reminderRepository:         rr,
		periodReminderRepository:   prr,
		hourReminderRepoRepository: hrr,
		notificationService:        ns,
	}
}
