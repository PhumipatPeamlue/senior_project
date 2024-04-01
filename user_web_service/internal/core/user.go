package core

import "time"

type IUser interface {
	BaseModel
	TimeSetting() IUserTimeSetting
	changeTimeSetting(newTimeSetting IUserTimeSetting)
}

type user struct {
	id          string
	timeSetting IUserTimeSetting
	createdAt   time.Time
	updatedAt   time.Time
}

// changeTimeSetting implements User.
func (u *user) changeTimeSetting(newTimeSetting IUserTimeSetting) {
	u.timeSetting = newTimeSetting
	u.updatedAt = time.Now().Local()
}

// CreatedAt implements User.
func (u *user) CreatedAt() time.Time {
	return u.createdAt
}

// ID implements User.
func (u *user) ID() string {
	return u.id
}

// TimeSetting implements User.
func (u *user) TimeSetting() IUserTimeSetting {
	return u.timeSetting
}

// UpdatedAt implements User.
func (u *user) UpdatedAt() time.Time {
	return u.updatedAt
}

func ScanUser(id string, timeSetting IUserTimeSetting, createdAt, updatedAt time.Time) IUser {
	return &user{
		id:          id,
		timeSetting: timeSetting,
		createdAt:   createdAt,
		updatedAt:   updatedAt,
	}
}

func newUser(lineUserID string) IUser {
	timeSetting := defaultUserTimeSetting()
	now := time.Now().Local()
	return &user{
		id:          lineUserID,
		timeSetting: timeSetting,
		createdAt:   now,
		updatedAt:   now,
	}
}
