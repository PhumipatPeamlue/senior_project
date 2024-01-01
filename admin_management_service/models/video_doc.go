package models

import "time"

type VideoDocES struct {
	Title       string     `json:"title"`
	VideoUrl    string     `json:"video_url"`
	Description string     `json:"description"`
	CreateAt    *time.Time `json:"create_at"`
	UpdateAt    *time.Time `json:"update_at"`
}

type VideoDoc struct {
	ID string `json:"id"`
	VideoDocES
}

type VideoDocGetResponse struct {
	Index       string     `json:"_index"`
	ID          string     `json:"_id"`
	Version     int        `json:"_version"`
	SeqNo       int        `json:"_seq_no"`
	PrimaryTerm int        `json:"_primary_term"`
	Found       bool       `json:"found"`
	Source      VideoDocES `json:"_source"`
}

type VideoDocSearchResponse struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	Shards   struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Skipped    int `json:"skipped"`
		Failed     int `json:"failed"`
	} `json:"_shards"`
	Hits struct {
		Total struct {
			Value    int    `json:"value"`
			Relation string `json:"relation"`
		} `json:"total"`
		MaxScore float64 `json:"max_score"`
		Hits     []struct {
			Index  string     `json:"_index"`
			Type   string     `json:"_type"`
			ID     string     `json:"_id"`
			Score  float64    `json:"_score"`
			Source VideoDocES `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

type VideoDocUpdatedBody struct {
	Doc VideoDocES `json:"doc"`
}
