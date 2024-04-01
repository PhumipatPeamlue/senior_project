package core

import (
	"time"

	"github.com/google/uuid"
)

type IPeriodNotification interface {
	INotification
	NotifyInfo() PeriodNotifyInfo
	changeNotifyInfo(updatedNotifyInfo PeriodNotifyInfo)
}

type periodNotification struct {
	notification
	notifyInfo PeriodNotifyInfo
}

// CreatedAt implements IPeriodNotification.
// Subtle: this method shadows the method (notification).CreatedAt of periodNotification.notification.
func (p *periodNotification) CreatedAt() time.Time {
	return p.createdAt
}

// DrugInfo implements IPeriodNotification.
// Subtle: this method shadows the method (notification).DrugInfo of periodNotification.notification.
func (p *periodNotification) DrugInfo() DrugInfo {
	return p.drugInfo
}

// FrequencyDayUsage implements IPeriodNotification.
// Subtle: this method shadows the method (notification).FrequencyDayUsage of periodNotification.notification.
func (p *periodNotification) FrequencyDayUsage() int {
	return p.frequencyDayUsage
}

// ID implements IPeriodNotification.
// Subtle: this method shadows the method (notification).ID of periodNotification.notification.
func (p *periodNotification) ID() string {
	return p.id
}

// NotifyInfo implements IPeriodNotification.
func (p *periodNotification) NotifyInfo() PeriodNotifyInfo {
	return p.notifyInfo
}

// PetID implements IPeriodNotification.
// Subtle: this method shadows the method (notification).PetID of periodNotification.notification.
func (p *periodNotification) PetID() string {
	return p.petID
}

// RenewIn implements IPeriodNotification.
// Subtle: this method shadows the method (notification).RenewIn of periodNotification.notification.
func (p *periodNotification) RenewIn() int {
	return p.renewIn
}

// Type implements IPeriodNotification.
// Subtle: this method shadows the method (notification).Type of periodNotification.notification.
func (p *periodNotification) Type() string {
	return p.notificationType
}

// UpdatedAt implements IPeriodNotification.
// Subtle: this method shadows the method (notification).UpdatedAt of periodNotification.notification.
func (p *periodNotification) UpdatedAt() time.Time {
	return p.updatedAt
}

// changeDrugInfo implements IPeriodNotification.
// Subtle: this method shadows the method (notification).changeDrugInfo of periodNotification.notification.
func (p *periodNotification) changeDrugInfo(updatedDrugInfo DrugInfo) {
	p.drugInfo = updatedDrugInfo
	p.updatedAt = time.Now().Local()
}

// changeFrequencyDayUsage implements IPeriodNotification.
// Subtle: this method shadows the method (notification).changeFrequencyDayUsage of periodNotification.notification.
func (p *periodNotification) changeFrequencyDayUsage(newFrequencyDayUsage int) {
	p.frequencyDayUsage = newFrequencyDayUsage
	p.updatedAt = time.Now().Local()
}

// changeNotifyInfo implements IPeriodNotification.
func (p *periodNotification) changeNotifyInfo(updatedNotifyInfo PeriodNotifyInfo) {
	p.notifyInfo = updatedNotifyInfo
	p.updatedAt = time.Now().Local()
}

func ScanPeriodNotification(id, petID string, drugInfo DrugInfo, frequencyDayUsage, renewIn int, createdAt, updatedAt time.Time, notifyInfo PeriodNotifyInfo) IPeriodNotification {
	notification := notification{
		id:                id,
		petID:             petID,
		notificationType:  "period",
		drugInfo:          drugInfo,
		frequencyDayUsage: frequencyDayUsage,
		renewIn:           renewIn,
		createdAt:         createdAt,
		updatedAt:         updatedAt,
	}
	return &periodNotification{
		notification: notification,
		notifyInfo:   notifyInfo,
	}
}

func newPeriodNotification(petID string, drugInfo DrugInfo, frequencyDayUsage int, notifyInfo PeriodNotifyInfo) IPeriodNotification {
	id := uuid.New().String()
	now := time.Now().Local()
	return ScanPeriodNotification(id, petID, drugInfo, frequencyDayUsage, 0, now, now, notifyInfo)
}
