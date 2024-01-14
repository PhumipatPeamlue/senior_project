package video_doc_repository

import "time"

type VideoDoc struct {
	Title       string     `json:"title"`
	VideoURL    string     `json:"video_url"`
	Description string     `json:"description"`
	ImageName   string     `json:"image_name"`
	CreateAt    *time.Time `json:"create_at"`
	UpdateAt    *time.Time `json:"update_at"`
}

type VideoDocWithID struct {
	ID string `json:"id"`
	VideoDoc
}

type GetResponseBody struct {
	ID     string   `json:"_id"`
	Source VideoDoc `json:"_source"`
}

type SearchResponseBody struct {
	Hits struct {
		Total struct {
			Value int `json:"value"`
		} `json:"total"`
		Hits []struct {
			ID     string   `json:"_id"`
			Source VideoDoc `json:"_source"`
		}
	}
}

type UpdateBody struct {
	Doc VideoDoc `json:"doc"`
}
