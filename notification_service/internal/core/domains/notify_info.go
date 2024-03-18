package domains

import "time"

type HourNotifyInfo struct {
	FirstUsage time.Time `json:"first_usage"`
	Every      int       `json:"every"`
}

type PeriodNotifyInfo struct {
	Morning   *time.Time `json:"morning"`
	Noon      *time.Time `json:"noon"`
	Evening   *time.Time `json:"evening"`
	BeforeBed *time.Time `json:"before_bed"`
}
