package models

type User struct {
	ID      int64  `json:"user_id"`
	UserID  string `json:"line_user_id"`
	Morning string `json:"morning"`
	Noon    string `json:"noon"`
	Evening string `json:"evening"`
	Night   string `json:"night"`
}
