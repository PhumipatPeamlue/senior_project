package domains

import "time"

type Notification struct {
	ID         int
	ReminderID string
	UserID     string
	Time       time.Time
	Status     string
}
