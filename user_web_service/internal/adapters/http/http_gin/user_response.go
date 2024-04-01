package http_gin

import "time"

type userResponse struct {
	UserID    string    `json:"user_id"`
	Morning   time.Time `json:"morning"`
	Noon      time.Time `json:"noon"`
	Evening   time.Time `json:"evening"`
	BeforeBed time.Time `json:"before_bed"`
}
