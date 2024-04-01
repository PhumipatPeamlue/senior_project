package core

import (
	"time"

	"github.com/google/uuid"
)

type IHourNotification interface {
	INotification
	NotifyInfo() HourNotifyInfo
	changeNotifyInfo(updatedNotifyInfo HourNotifyInfo)
}

type hourNotification struct {
	notification
	notifyInfo HourNotifyInfo
}

// CreatedAt implements IHourNotification.
// Subtle: this method shadows the method (notification).CreatedAt of hourNotification.notification.
func (h *hourNotification) CreatedAt() time.Time {
	return h.createdAt
}

// DrugInfo implements IHourNotification.
// Subtle: this method shadows the method (notification).DrugInfo of hourNotification.notification.
func (h *hourNotification) DrugInfo() DrugInfo {
	return h.drugInfo
}

// FrequencyDayUsage implements IHourNotification.
// Subtle: this method shadows the method (notification).FrequencyDayUsage of hourNotification.notification.
func (h *hourNotification) FrequencyDayUsage() int {
	return h.frequencyDayUsage
}

// ID implements IHourNotification.
// Subtle: this method shadows the method (notification).ID of hourNotification.notification.
func (h *hourNotification) ID() string {
	return h.id
}

// NotifyInfo implements IHourNotification.
func (h *hourNotification) NotifyInfo() HourNotifyInfo {
	return h.notifyInfo
}

// PetID implements IHourNotification.
// Subtle: this method shadows the method (notification).PetID of hourNotification.notification.
func (h *hourNotification) PetID() string {
	return h.petID
}

// RenewIn implements IHourNotification.
// Subtle: this method shadows the method (notification).RenewIn of hourNotification.notification.
func (h *hourNotification) RenewIn() int {
	return h.renewIn
}

// Type implements IHourNotification.
// Subtle: this method shadows the method (notification).Type of hourNotification.notification.
func (h *hourNotification) Type() string {
	return h.notificationType
}

// UpdatedAt implements IHourNotification.
// Subtle: this method shadows the method (notification).UpdatedAt of hourNotification.notification.
func (h *hourNotification) UpdatedAt() time.Time {
	return h.updatedAt
}

// changeDrugInfo implements IHourNotification.
// Subtle: this method shadows the method (notification).changeDrugInfo of hourNotification.notification.
func (h *hourNotification) changeDrugInfo(updatedDrugInfo DrugInfo) {
	h.drugInfo = updatedDrugInfo
	h.updatedAt = time.Now().Local()
}

// changeFrequencyDayUsage implements IHourNotification.
// Subtle: this method shadows the method (notification).changeFrequencyDayUsage of hourNotification.notification.
func (h *hourNotification) changeFrequencyDayUsage(newFrequencyDayUsage int) {
	h.frequencyDayUsage = newFrequencyDayUsage
	h.updatedAt = time.Now().Local()
}

// changeNotifyInfo implements IHourNotification.
func (h *hourNotification) changeNotifyInfo(updatedNotifyInfo HourNotifyInfo) {
	h.notifyInfo = updatedNotifyInfo
	h.updatedAt = time.Now().Local()
}

func ScanHourNotification(id, petID string, drugInfo DrugInfo, frequencyDayUsage, renewIn int, createdAt, updatedAt time.Time, notifyInfo HourNotifyInfo) IHourNotification {
	notification := notification{
		id:                id,
		petID:             petID,
		notificationType:  "hour",
		drugInfo:          drugInfo,
		frequencyDayUsage: frequencyDayUsage,
		renewIn:           renewIn,
		createdAt:         createdAt,
		updatedAt:         updatedAt,
	}
	return &hourNotification{
		notification: notification,
		notifyInfo:   notifyInfo,
	}
}

func newHourNotification(petID string, drugInfo DrugInfo, frequencyDayUsage int, notifyInfo HourNotifyInfo) IHourNotification {
	id := uuid.New().String()
	now := time.Now().Local()
	return ScanHourNotification(id, petID, drugInfo, frequencyDayUsage, 0, now, now, notifyInfo)
}
