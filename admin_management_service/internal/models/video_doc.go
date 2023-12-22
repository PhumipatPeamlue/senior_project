package models

import (
	"time"
)

type VideoDoc struct {
	Title       string    `json:"title"`
	VideoUrl    string    `json:"video_url"`
	Description string    `json:"description"`
	CreateAt    time.Time `json:"create_at"`
	UpdateAt    time.Time `json:"update_at"`
}

type VideoDocDto struct {
	ID string `json:"id"`
	VideoDoc
}

type VideoDocGetResult struct {
	Index       string   `json:"_index"`
	ID          string   `json:"_id"`
	Version     int      `json:"_version"`
	SeqNo       int      `json:"_seq_no"`
	PrimaryTerm int      `json:"_primary_term"`
	Found       bool     `json:"found"`
	Source      VideoDoc `json:"_source"`
}

type VideoDocSearchResult struct {
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
			Index  string   `json:"_index"`
			Type   string   `json:"_type"`
			ID     string   `json:"_id"`
			Score  float64  `json:"_score"`
			Source VideoDoc `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}
