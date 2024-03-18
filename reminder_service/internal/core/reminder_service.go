package core

import (
	"context"
	"reminder_service/internal/core/domains"
)

type hourReminderServiceInterface interface {
	FindHourReminder(ctx context.Context, reminderID string) (hr domains.HourReminder, err error)
	AddNewHourReminder(ctx context.Context, petID string, drugInfo domains.DrugInfo, frequencyDayUsage int, notifyInfo domains.HourNotifyInfo) (err error)
	ChangeHourReminderInfo(ctx context.Context, reminderID string, drugInfo domains.DrugInfo, frequencyDayUsage int, notifyInfo domains.HourNotifyInfo) (err error)
	RemoveHourReminder(ctx context.Context, reminderID string) (err error)
}

type periodReminderServiceInterface interface {
	FindPeriodReminder(ctx context.Context, reminderID string) (pr domains.PeriodReminder, err error)
	AddNewPeriodReminder(ctx context.Context, petID string, drugInfo domains.DrugInfo, frequencyDayUsage int, notifyInfo domains.PeriodNotifyInfo) (err error)
	ChangePeriodReminderInfo(ctx context.Context, reminderID string, drugInfo domains.DrugInfo, frequencyDayUsage int, notifyInfo domains.PeriodNotifyInfo) (err error)
	RemovePeriodReminder(ctx context.Context, reminderID string) (err error)
}

type ReminderServiceInterface interface {
	FindReminder(ctx context.Context, reminderID string) (reminder domains.Reminder, err error)
	FindAllPetReminders(ctx context.Context, petID string) (reminders []domains.Reminder, err error)
	RemoveAllPetReminders(ctx context.Context, petID string) (err error)
	hourReminderServiceInterface
	periodReminderServiceInterface
}

type reminderService struct {
	repository          ReminderRepositoryInterface
	notificationService NotificationServiceInterface
}

func (r *reminderService) FindReminder(ctx context.Context, reminderID string) (reminder domains.Reminder, err error) {
	reminder, err = r.repository.ReadByID(ctx, reminderID)
	return
}

func (r *reminderService) FindAllPetReminders(ctx context.Context, petID string) (reminders []domains.Reminder, err error) {
	reminders, err = r.repository.ReadByPetID(ctx, petID)
	return
}

func (r *reminderService) RemoveAllPetReminders(ctx context.Context, petID string) (err error) {
	err = r.repository.DeleteByPetID(ctx, petID)
	return
}

func (r *reminderService) FindHourReminder(ctx context.Context, reminderID string) (hr domains.HourReminder, err error) {
	hr, err = r.repository.ReadHourReminderByID(ctx, reminderID)
	return
}

func (r *reminderService) AddNewHourReminder(ctx context.Context, petID string, drugInfo domains.DrugInfo, frequencyDayUsage int, notifyInfo domains.HourNotifyInfo) (err error) {
	hr := domains.NewHourReminder(petID, drugInfo, frequencyDayUsage, notifyInfo)
	err = r.repository.CreateHourReminder(ctx, hr)
	if err != nil {
		return
	}

	err = r.notificationService.AddNewFromHourReminder(ctx, petID, hr.ID(), notifyInfo)
	return
}

func (r *reminderService) ChangeHourReminderInfo(ctx context.Context, reminderID string, drugInfo domains.DrugInfo, frequencyDayUsage int, notifyInfo domains.HourNotifyInfo) (err error) {
	hr, err := r.FindHourReminder(ctx, reminderID)
	if err != nil {
		return
	}

	updateNotification := false
	if hr.NotifyInfo() != notifyInfo {
		updateNotification = true
	}

	hr.ChangeFrequencyDayUsage(frequencyDayUsage)
	hr.ChangeDrugInfo(drugInfo)
	hr.ChangeNotifyInfo(notifyInfo)

	err = r.repository.UpdateHourReminder(ctx, hr)
	if err != nil {
		return
	}

	if updateNotification {
		err = r.notificationService.ChangeNotifyTimeFromHourReminder(ctx, hr.PetID(), reminderID, notifyInfo)
	}
	return
}

func (r *reminderService) RemoveHourReminder(ctx context.Context, reminderID string) (err error) {
	err = r.repository.DeleteByID(ctx, reminderID)
	return
}

func (r *reminderService) FindPeriodReminder(ctx context.Context, reminderID string) (pr domains.PeriodReminder, err error) {
	pr, err = r.repository.ReadPeriodReminderByID(ctx, reminderID)
	return
}

func (r *reminderService) AddNewPeriodReminder(ctx context.Context, petID string, drugInfo domains.DrugInfo, frequencyDayUsage int, notifyInfo domains.PeriodNotifyInfo) (err error) {
	pr := domains.NewPeriodReminder(petID, drugInfo, frequencyDayUsage, notifyInfo)
	err = r.repository.CreatePeriodReminder(ctx, pr)
	if err != nil {
		return
	}

	err = r.notificationService.AddNewFromPeriodReminder(ctx, petID, pr.ID(), notifyInfo)
	return
}

func (r *reminderService) ChangePeriodReminderInfo(ctx context.Context, reminderID string, drugInfo domains.DrugInfo, frequencyDayUsage int, notifyInfo domains.PeriodNotifyInfo) (err error) {
	pr, err := r.FindPeriodReminder(ctx, reminderID)
	if err != nil {
		return
	}

	updatedNotification := false
	if pr.NotifyInfo() != notifyInfo {
		pr.ChangeNotifyInfo(notifyInfo)
		updatedNotification = true
	}

	pr.ChangeFrequencyDayUsage(frequencyDayUsage)
	pr.ChangeDrugInfo(drugInfo)

	err = r.repository.UpdatePeriodReminder(ctx, pr)
	if err != nil {
		return
	}

	if updatedNotification {
		err = r.notificationService.ChangeNotifyTimeFromPeriodReminder(ctx, pr.PetID(), reminderID, notifyInfo)
	}

	return
}

func (r *reminderService) RemovePeriodReminder(ctx context.Context, reminderID string) (err error) {
	err = r.repository.DeleteByID(ctx, reminderID)
	return
}

func NewReminderService(repository ReminderRepositoryInterface, notificationService NotificationServiceInterface) ReminderServiceInterface {
	return &reminderService{
		repository:          repository,
		notificationService: notificationService,
	}
}
