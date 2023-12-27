package models

import "time"

type Pet struct {
	ID           string    `json:"pet_id"`
	UserID       string    `json:"user_id"`
	AnimalTypeID string    `json:"animaltype_type_id"`
	Name         string    `json:"pet_name"`
	Weight       float64   `json:"weight"`
	BeSterile    int       `json:"be sterile"`
	Age          string    `json:"age"`
	Activities   string    `json:"activitie"`
	CreatedDate  time.Time `json:"created_date"`
	UpdatedDate  time.Time `json:"updated_date"`
}
