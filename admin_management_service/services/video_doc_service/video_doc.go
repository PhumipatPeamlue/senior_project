package video_doc_service

import "time"

type VideoDoc struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	VideoURL    string     `json:"video_url"`
	Description string     `json:"description"`
	CreateAt    *time.Time `json:"create_at"`
	UpdateAt    *time.Time `json:"update_at"`
}

type GetResponse struct {
	Doc      VideoDoc `json:"doc"`
	ImageURL string   `json:"image_url"`
}

type SearchResponse struct {
	Total int        `json:"total"`
	Data  []VideoDoc `json:"data"`
}

type NewVideoDocRequest struct {
	Title       string `form:"title"`
	VideoURL    string `form:"video_url"`
	Description string `form:"description"`
}

type UpdateVideoDocRequest struct {
	ID          string `form:"id"`
	Title       string `form:"title"`
	VideoURL    string `form:"video_url"`
	Description string `form:"description"`
	DeleteImage bool   `form:"delete_image"`
}
