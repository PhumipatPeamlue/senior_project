package core

import (
	"context"
	"cronjob/internal/core/domains"
	"log"
)

type ServiceInterface interface {
	Renew(ctx context.Context) (err error)
	SendNotification(ctx context.Context) (err error)
}

type service struct {
	reminderRepository      ReminderRepositoryInterface
	notificationRepository  NotificationRepositoryInterface
	petRepository           PetRepositoryInterface
	lineNotificationService LineNotificationServiceInterface
}

func (r *service) SendNotification(ctx context.Context) (err error) {
	notifications, err := r.notificationRepository.ReadByWaitStatus(ctx)
	if err != nil {
		return
	}

	for _, notification := range notifications {
		var reminder domains.Reminder
		reminder, err = r.reminderRepository.ReadByID(ctx, notification.ReminderID())
		if err != nil {
			log.Println("error at reminder's repository:", err)
			continue
		}

		notification.Sent()
		err = r.notificationRepository.Update(ctx, notification)
		if err != nil {
			log.Println("error at notification's repository:", err)
			continue
		}

		var pet domains.Pet
		pet, err = r.petRepository.ReadByID(ctx, notification.PetID())
		if err != nil {
			log.Println("error at pet's repository:", err)
			continue
		}

		err = r.lineNotificationService.PushMessage(ctx, pet.UserID(), pet.Name(), reminder.DrugInfo(), notification.NotifyAt())
		if err != nil {
			log.Println("error at LINE notification service:", err)
			continue
		}
	}

	return
}

func (r *service) Renew(ctx context.Context) (err error) {
	reminders, err := r.reminderRepository.ReadAllZeroRenew(ctx)
	if err != nil {
		return
	}

	for _, reminder := range reminders {
		reminder.ReNew()
		err = r.reminderRepository.UpdateRenew(ctx, reminder)
		if err != nil {
			return
		}

		notifications := domains.NewNotification(reminder)
		for _, notification := range notifications {
			err = r.notificationRepository.Create(ctx, notification)
			if err != nil {
				return
			}
		}
	}

	reminders, err = r.reminderRepository.ReadAll(ctx)
	if err != nil {
		return
	}

	for _, reminder := range reminders {
		reminder.DecrementRenew()
		err = r.reminderRepository.UpdateRenew(ctx, reminder)
		if err != nil {
			return
		}
	}
	return
}

func NewService(repositoryInterface ReminderRepositoryInterface, notificationRepository NotificationRepositoryInterface, petRepository PetRepositoryInterface, lineNotificationService LineNotificationServiceInterface) ServiceInterface {
	return &service{
		reminderRepository:      repositoryInterface,
		notificationRepository:  notificationRepository,
		petRepository:           petRepository,
		lineNotificationService: lineNotificationService,
	}
}
