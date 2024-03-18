package external_api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"reminder_service/internal/core"
	"reminder_service/internal/core/domains"
)

type notificationService struct {
	client *http.Client
}

// AddNewFromHourReminder implements core.NotificationServiceInterface.
func (n *notificationService) AddNewFromHourReminder(ctx context.Context, petID, reminderID string, info domains.HourNotifyInfo) (err error) {
	body := addNewFromHourReminderRequest{
		PetID:          petID,
		ReminderID:     reminderID,
		HourNotifyInfo: info,
	}
	b, err := json.Marshal(body)
	if err != nil {
		return
	}

	req, err := http.NewRequestWithContext(ctx, "POST", os.Getenv("ADD_NEW_FROM_HOUR_REMINDER_URL"), bytes.NewBuffer(b))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := n.client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusConflict:
		err = core.NewErrNotificationStatusConflict(errors.New(resp.Status))
	case http.StatusInternalServerError:
		err = core.NewErrNotficationStatusInternalServerError(errors.New(resp.Status))
	}

	return
}

// AddNewFromPeriodReminder implements core.NotificationServiceInterface.
func (n *notificationService) AddNewFromPeriodReminder(ctx context.Context, petID, reminderID string, info domains.PeriodNotifyInfo) (err error) {
	body := addNewFromPeriodReminderRequest{
		PetID:            petID,
		ReminderID:       reminderID,
		PeriodNotifyInfo: info,
	}
	b, err := json.Marshal(body)
	if err != nil {
		return
	}

	req, err := http.NewRequestWithContext(ctx, "POST", os.Getenv("ADD_NEW_FROM_PERIOD_REMINDER_URL"), bytes.NewBuffer(b))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := n.client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusConflict:
		err = core.NewErrNotificationStatusConflict(errors.New(resp.Status))
	case http.StatusInternalServerError:
		err = core.NewErrNotficationStatusInternalServerError(errors.New(resp.Status))
	}

	return
}

// ChangeNotifyTimeFromHourReminder implements core.NotificationServiceInterface.
func (n *notificationService) ChangeNotifyTimeFromHourReminder(ctx context.Context, petID, reminderID string, info domains.HourNotifyInfo) (err error) {
	body := changeNotifyTimeFromHourReminder{
		PetID:          petID,
		ReminderID:     reminderID,
		HourNotifyInfo: info,
	}
	b, err := json.Marshal(body)
	if err != nil {
		return
	}

	req, err := http.NewRequestWithContext(ctx, "PUT", os.Getenv("CHANGE_NOTIFY_TIME_FROM_HOUR_REMINDER_URL"), bytes.NewBuffer(b))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := n.client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusInternalServerError:
		err = core.NewErrNotficationStatusInternalServerError(errors.New(resp.Status))
	}

	return
}

// ChangeNotifyTimeFromPeriodReminder implements core.NotificationServiceInterface.
func (n *notificationService) ChangeNotifyTimeFromPeriodReminder(ctx context.Context, petID, reminderID string, info domains.PeriodNotifyInfo) (err error) {
	body := changeNotifyTimeFromPeriodReminder{
		PetID:            petID,
		ReminderID:       reminderID,
		PeriodNotifyInfo: info,
	}
	b, err := json.Marshal(body)
	if err != nil {
		return
	}

	req, err := http.NewRequestWithContext(ctx, "PUT", os.Getenv("CHANGE_NOTIFY_TIME_FROM_PERIOD_REMINDER_URL"), bytes.NewBuffer(b))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := n.client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusInternalServerError:
		err = core.NewErrNotficationStatusInternalServerError(errors.New(resp.Status))
	}

	return
}

func NewNotificationService(client *http.Client) core.NotificationServiceInterface {
	return &notificationService{
		client: client,
	}
}
