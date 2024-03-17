package http_gin

import (
	"time"
	"user_service/internal/core"
)

type userResponse struct {
	LineUserID string    `json:"user_id"`
	Morning    time.Time `json:"morning"`
	Noon       time.Time `json:"noon"`
	Evening    time.Time `json:"evening"`
	BeforeBed  time.Time `json:"before_bed"`
}

func newUserResponse(user core.User) userResponse {
	lineUserID := user.ID()
	ts := user.TimeSetting()
	morning, noon, evening, beforeBed := ts.Morning(), ts.Noon(), ts.Evening(), ts.BeforeBed()
	return userResponse{
		LineUserID: lineUserID,
		Morning:    morning,
		Noon:       noon,
		Evening:    evening,
		BeforeBed:  beforeBed,
	}
}
