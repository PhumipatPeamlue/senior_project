package domains

import "time"

type PeriodReminder struct {
	Reminder
	Morning   *time.Time `json:"morning"`
	Noon      *time.Time `json:"noon"`
	Evening   *time.Time `json:"evening"`
	BeforeBed *time.Time `json:"before_bed"`
}
