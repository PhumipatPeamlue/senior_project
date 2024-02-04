package domains

import "time"

type HourReminder struct {
	Reminder
	FirstUsage time.Time `json:"first_usage"`
	Every      int       `json:"every"`
}
