package core

import "time"

type BaseModel interface {
	ID() string
	CreatedAt() time.Time
	UpdatedAt() time.Time
}
