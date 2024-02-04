package domains

import "time"

type User struct {
	ID        string    `json:"user_id"`
	Morning   time.Time `json:"morning"`
	Noon      time.Time `json:"noon"`
	Evening   time.Time `json:"evening"`
	BeforeBed time.Time `json:"before_bed"`
}

type UserTimeSetting struct {
	Morning   time.Time `json:"morning"`
	Noon      time.Time `json:"noon"`
	Evening   time.Time `json:"evening"`
	BeforeBed time.Time `json:"before_bed"`
}
