package http_gin

import "time"

type changeTimeSettingRequest struct {
	LineUserID string    `json:"user_id" binding:"required"`
	Morning    time.Time `json:"morning"`
	Noon       time.Time `json:"noon"`
	Evening    time.Time `json:"evening"`
	BeforeBed  time.Time `json:"before_bed"`
}
