package core

import (
	"strconv"
	"time"
)

const (
	NOTIFICATION_RECORD_WAIT_STATUS = "wait"
)

type INotificationRecord interface {
	BaseModel
	PetID() string
	NotificationID() string
	NotifyAt() time.Time
	Status() string
}

type notificationRecord struct {
	id             int
	petID          string
	notificationID string
	notifyAt       time.Time
	status         string
	createdAt      time.Time
	updatedAt      time.Time
}

// CreatedAt implements INotificationRecord.
func (n *notificationRecord) CreatedAt() time.Time {
	return n.createdAt
}

// ID implements INotificationRecord.
func (n *notificationRecord) ID() string {
	return strconv.Itoa(n.id)
}

// NotificationID implements INotificationRecord.
func (n *notificationRecord) NotificationID() string {
	return n.notificationID
}

// NotifyAt implements INotificationRecord.
func (n *notificationRecord) NotifyAt() time.Time {
	return n.notifyAt
}

// PetID implements INotificationRecord.
func (n *notificationRecord) PetID() string {
	return n.petID
}

// Status implements INotificationRecord.
func (n *notificationRecord) Status() string {
	return n.status
}

// UpdatedAt implements INotificationRecord.
func (n *notificationRecord) UpdatedAt() time.Time {
	return n.updatedAt
}

func ScanNotificationRecord(id int, petID, notificationID, status string, notifyAt, createdAt, updatedAt time.Time) INotificationRecord {
	return &notificationRecord{
		id:             id,
		petID:          petID,
		notificationID: notificationID,
		status:         status,
		notifyAt:       notifyAt,
		createdAt:      createdAt,
		updatedAt:      updatedAt,
	}
}

func newNotificationRecord(petID, notificationID string, notifyAt time.Time) INotificationRecord {
	status := NOTIFICATION_RECORD_WAIT_STATUS
	now := time.Now().Local()
	return &notificationRecord{
		petID:          petID,
		notificationID: notificationID,
		status:         status,
		notifyAt:       notifyAt,
		createdAt:      now,
		updatedAt:      now,
	}
}
