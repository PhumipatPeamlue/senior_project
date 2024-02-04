package domains

type Pet struct {
	ID     string `json:"pet_id"`
	UserID string `json:"user_id"`
	Name   string `json:"name"`
}
