package services

import (
	"context"
	"reminder_service/internal/core/domains"
	"reminder_service/internal/core/ports"
	"time"
)

type notificationService struct {
	repo ports.NotificationRepository
}

func (s *notificationService) RemoveByReminderID(ctx context.Context, reminderID string) (err error) {
	err = s.repo.DeleteByReminderID(ctx, reminderID)
	return
}

func (s *notificationService) AddNewFromPeriodReminder(ctx context.Context, reminderID string, userID string, morning *time.Time, noon *time.Time, evening *time.Time, beforeBed *time.Time) (err error) {
	list := []*time.Time{morning, noon, evening, beforeBed}

	for _, t := range list {
		if t == nil {
			continue
		}
		n := domains.Notification{
			ReminderID: reminderID,
			UserID:     userID,
			Time:       *t,
			Status:     "not sent",
		}
		if err = s.repo.Save(ctx, n); err != nil {
			return
		}
	}
	return
}

func (s *notificationService) AddNewFromHourReminder(ctx context.Context, reminderID string, userID string, firstUsage time.Time, every int) (err error) {
	bkkLocation, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		return
	}

	firstUsage = firstUsage.In(bkkLocation)
	numUsagePerDay := (24 - firstUsage.Hour()) / every
	for i := 0; i <= numUsagePerDay; i++ {
		n := domains.Notification{
			ID:         0,
			ReminderID: reminderID,
			UserID:     userID,
			Time:       firstUsage.Add(time.Duration(i*every) * time.Hour).UTC(),
			Status:     "not sent",
		}
		if err = s.repo.Save(ctx, n); err != nil {
			return
		}
	}
	return
}

func NewNotificationService(repo ports.NotificationRepository) ports.NotificationService {
	return &notificationService{
		repo: repo,
	}
}
