package models

import "time"

type VideoDoc struct {
	Title       string     `json:"title" form:"title"`
	Description string     `json:"description" form:"description"`
	VideoURL    string     `json:"video_url" form:"video_url"`
	CreateAt    *time.Time `json:"create_at"`
	UpdateAt    *time.Time `json:"update_at"`
}

type VideoDocWithId struct {
	ID string `json:"id" form:"id"`
	VideoDoc
}
