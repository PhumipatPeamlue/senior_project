package core

import (
	"time"

	"github.com/google/uuid"
)

type IPet interface {
	BaseModel
	UserID() string
	Name() string
	changeName(newName string)
}

type pet struct {
	id        string
	userID    string
	name      string
	createdAt time.Time
	updatedAt time.Time
}

// CreatedAt implements IPet.
func (p *pet) CreatedAt() time.Time {
	return p.createdAt
}

// ID implements IPet.
func (p *pet) ID() string {
	return p.id
}

// Name implements IPet.
func (p *pet) Name() string {
	return p.name
}

// UpdatedAt implements IPet.
func (p *pet) UpdatedAt() time.Time {
	return p.updatedAt
}

// UserID implements IPet.
func (p *pet) UserID() string {
	return p.userID
}

// changeName implements IPet.
func (p *pet) changeName(newName string) {
	p.name = newName
	p.updatedAt = time.Now().Local()
}

func ScanPet(id, userID, name string, createdAt, updatedAt time.Time) IPet {
	return &pet{
		id:        id,
		userID:    userID,
		name:      name,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

func newPet(userID, name string) IPet {
	id := uuid.New().String()
	now := time.Now().Local()
	return &pet{
		id:        id,
		userID:    userID,
		name:      name,
		createdAt: now,
		updatedAt: now,
	}
}
