package core

import "time"

type IUserTimeSetting interface {
	Morning() time.Time
	Noon() time.Time
	Evening() time.Time
	BeforeBed() time.Time
}

type userTimeSetting struct {
	morning   time.Time
	noon      time.Time
	evening   time.Time
	beforeBed time.Time
}

// BeforeBed implements UserTimeSetting.
func (u *userTimeSetting) BeforeBed() time.Time {
	return u.beforeBed
}

// Evening implements UserTimeSetting.
func (u *userTimeSetting) Evening() time.Time {
	return u.evening
}

// Morning implements UserTimeSetting.
func (u *userTimeSetting) Morning() time.Time {
	return u.morning
}

// Noon implements UserTimeSetting.
func (u *userTimeSetting) Noon() time.Time {
	return u.noon
}

func ScanUserTimeSetting(morning, noon, evening, beforeBed time.Time) IUserTimeSetting {
	return &userTimeSetting{
		morning:   morning,
		noon:      noon,
		evening:   evening,
		beforeBed: beforeBed,
	}
}

func defaultUserTimeSetting() IUserTimeSetting {
	now := time.Now().Local()
	year := now.Year()
	month := now.Month()
	day := now.Day()
	location := now.Location()
	return &userTimeSetting{
		morning:   time.Date(year, month, day, 8, 0, 0, 0, location),
		noon:      time.Date(year, month, day, 12, 0, 0, 0, location),
		evening:   time.Date(year, month, day, 16, 0, 0, 0, location),
		beforeBed: time.Date(year, month, day, 20, 0, 0, 0, location),
	}
}
