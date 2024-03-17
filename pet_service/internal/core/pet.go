package core

import (
	"time"

	"github.com/google/uuid"
)

type Pet struct {
	id        string
	userID    string
	name      string
	createdAt time.Time
	updatedAt time.Time
}

func (p *Pet) ID() string {
	return p.id
}

func (p *Pet) UserID() string {
	return p.userID
}

func (p *Pet) Name() string {
	return p.name
}

func (p *Pet) CreatedAt() time.Time {
	return p.createdAt
}

func (p *Pet) UpdatedAt() time.Time {
	return p.updatedAt
}

func (p *Pet) ChangeName(newName string) {
	p.name = newName
	p.updatedAt = time.Now().Local()
}

func ScanPet(id, userID, name string, createdAt, updatedAt time.Time) Pet {
	return Pet{
		id:        id,
		userID:    userID,
		name:      name,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

func newPet(userID, name string) Pet {
	id := uuid.New().String()
	now := time.Now().Local()
	pet := ScanPet(id, userID, name, now, now)
	return pet
}
