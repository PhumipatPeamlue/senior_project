package domains

import (
	"time"
)

const (
	SentStatus = "sent"
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

func (n *Notification) Sent() {
	n.status = SentStatus
	n.updatedAt = time.Now().Local()
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

func NewNotification(reminder Reminder) []Notification {
	petID := reminder.petID
	reminderID := reminder.id
	notifications := make([]Notification, 0)

	switch reminder.reminderType {
	case "hour":
		hourNotifyInfo := reminder.hourNotifyInfo
		firstUsage, every := hourNotifyInfo.FirstUsage, *hourNotifyInfo.Every
		usePerDay := (24 - firstUsage.Hour()) / every

		for i := 0; i <= usePerDay; i++ {
			now := time.Now().Local()
			notifyAt := firstUsage.Add(time.Duration(i*every) * time.Hour)
			notifyAt = time.Date(now.Year(), now.Month(), now.Day(), notifyAt.Hour(), notifyAt.Minute(), 0, 0, time.Local)
			notification := ScanNotification(0, petID, reminderID, WaitStatus, notifyAt, now, now)
			notifications = append(notifications, notification)
		}
	case "period":
		periodNotifyInfo := reminder.periodNotifyInfo
		morning, noon, evening, beforeBed := periodNotifyInfo.Morning, periodNotifyInfo.Noon, periodNotifyInfo.Evening, periodNotifyInfo.BeforeBed
		periods := []*time.Time{morning, noon, evening, beforeBed}

		for _, period := range periods {
			if period == nil {
				continue
			}

			now := time.Now().Local()
			notifyAt := *period
			notifyAt = time.Date(now.Year(), now.Month(), now.Day(), notifyAt.Hour(), notifyAt.Minute(), 0, 0, time.Local)
			notification := ScanNotification(0, petID, reminderID, WaitStatus, notifyAt, now, now)
			notifications = append(notifications, notification)
		}
	}

	return notifications
}
