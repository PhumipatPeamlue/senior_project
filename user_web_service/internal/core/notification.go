package core

import "time"

type INotification interface {
	BaseModel
	PetID() string
	Type() string
	DrugInfo() DrugInfo
	FrequencyDayUsage() int
	RenewIn() int
	CreatedAt() time.Time
	UpdatedAt() time.Time
	changeFrequencyDayUsage(newFrequencyDayUsage int)
	changeDrugInfo(updatedDrugInfo DrugInfo)
}

type notification struct {
	id                string
	petID             string
	notificationType  string
	drugInfo          DrugInfo
	frequencyDayUsage int
	renewIn           int
	createdAt         time.Time
	updatedAt         time.Time
}

// changeDrugInfo implements INotification.
func (n *notification) changeDrugInfo(updatedDrugInfo DrugInfo) {
	n.drugInfo = updatedDrugInfo
	n.updatedAt = time.Now().Local()
}

// changeFrequencyDayUsage implements INotification.
func (n *notification) changeFrequencyDayUsage(newFrequencyDayUsage int) {
	n.frequencyDayUsage = newFrequencyDayUsage
	n.renewIn = 0
	n.updatedAt = time.Now().Local()
}

// CreatedAt implements INotification.
func (n *notification) CreatedAt() time.Time {
	return n.createdAt
}

// DrugInfo implements INotification.
func (n *notification) DrugInfo() DrugInfo {
	return n.drugInfo
}

// FrequencyDayUsage implements INotification.
func (n *notification) FrequencyDayUsage() int {
	return n.frequencyDayUsage
}

// ID implements INotification.
func (n *notification) ID() string {
	return n.id
}

// PetID implements INotification.
func (n *notification) PetID() string {
	return n.petID
}

// RenewIn implements INotification.
func (n *notification) RenewIn() int {
	return n.renewIn
}

// Type implements INotification.
func (n *notification) Type() string {
	return n.notificationType
}

// UpdatedAt implements INotification.
func (n *notification) UpdatedAt() time.Time {
	return n.updatedAt
}

func ScanNotification(id, petID, notificationType string, drugInfo DrugInfo, frequencyDayUsage, renewIn int, createdAt, updatedAt time.Time) INotification {
	return &notification{
		id:                id,
		petID:             petID,
		notificationType:  notificationType,
		drugInfo:          drugInfo,
		frequencyDayUsage: frequencyDayUsage,
		renewIn:           renewIn,
		createdAt:         createdAt,
		updatedAt:         updatedAt,
	}
}
