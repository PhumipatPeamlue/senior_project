package models

import "time"

type VideoDoc struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	VideoURL    string     `json:"video_url"`
	Description string     `json:"description"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}
