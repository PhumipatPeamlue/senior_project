package core

import "time"

type UserTimeSetting struct {
	morning   time.Time
	noon      time.Time
	evening   time.Time
	beforeBed time.Time
}

func (u *UserTimeSetting) Morning() time.Time {
	return u.morning
}

func (u *UserTimeSetting) Noon() time.Time {
	return u.noon
}

func (u *UserTimeSetting) Evening() time.Time {
	return u.evening
}

func (u *UserTimeSetting) BeforeBed() time.Time {
	return u.beforeBed
}

func ScanUserTimeSetting(morning, noon, evening, beforeBed time.Time) UserTimeSetting {
	return UserTimeSetting{
		morning:   morning,
		noon:      noon,
		evening:   evening,
		beforeBed: beforeBed,
	}
}

func defaultUserTimeSetting() UserTimeSetting {
	now := time.Now().Local()
	year := now.Year()
	month := now.Month()
	day := now.Day()
	location := now.Location()
	return UserTimeSetting{
		morning:   time.Date(year, month, day, 8, 0, 0, 0, location),
		noon:      time.Date(year, month, day, 12, 0, 0, 0, location),
		evening:   time.Date(year, month, day, 16, 0, 0, 0, location),
		beforeBed: time.Date(year, month, day, 20, 0, 0, 0, location),
	}
}

type User struct {
	id          string
	timeSetting UserTimeSetting
	createdAt   time.Time
	updatedAt   time.Time
}

func (u *User) ID() string {
	return u.id
}

func (u *User) TimeSetting() UserTimeSetting {
	return u.timeSetting
}

func (u *User) CreatedAt() time.Time {
	return u.createdAt
}

func (u *User) UpdatedAt() time.Time {
	return u.updatedAt
}

func (u *User) changeTimeSetting(newTimeSetting UserTimeSetting) {
	u.timeSetting = newTimeSetting
	u.updatedAt = time.Now().Local()
}

func ScanUser(lineUserID string, timeSetting UserTimeSetting, createdAt time.Time, updatedAt time.Time) User {
	return User{
		id:          lineUserID,
		timeSetting: timeSetting,
		createdAt:   createdAt,
		updatedAt:   updatedAt,
	}
}

func newUser(lineUserID string) User {
	ts := defaultUserTimeSetting()
	now := time.Now().Local()
	return ScanUser(lineUserID, ts, now, now)
}
