package domains

import (
	"time"
)

const (
	WaitStatus = "wait"
)

type Notification struct {
	id         int
	petID      string
	reminderID string
	notifyAt   time.Time
	status     string
	createdAt  time.Time
	updatedAt  time.Time
}

func (n *Notification) ID() int {
	return n.id
}

func (n *Notification) PetID() string {
	return n.petID
}

func (n *Notification) ReminderID() string {
	return n.reminderID
}

func (n *Notification) NotifyAt() time.Time {
	return n.notifyAt
}

func (n *Notification) Status() string {
	return n.status
}

func (n *Notification) CreatedAt() time.Time {
	return n.createdAt
}

func (n *Notification) UpdatedAt() time.Time {
	return n.updatedAt
}

func ScanNotification(id int, petID, reminderID, status string, notifyAt, createdAt, updatedAt time.Time) Notification {
	return Notification{
		id:         id,
		petID:      petID,
		reminderID: reminderID,
		status:     status,
		notifyAt:   notifyAt,
		createdAt:  createdAt,
		updatedAt:  updatedAt,
	}
}

func NewNotification(petID, reminderID string, notifyAt time.Time) Notification {
	status := WaitStatus
	now := time.Now().Local()
	return ScanNotification(0, petID, reminderID, status, notifyAt, now, now)
}
