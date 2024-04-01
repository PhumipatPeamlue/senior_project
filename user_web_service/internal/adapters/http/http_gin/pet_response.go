package http_gin

type petResponse struct {
	PetID  string `json:"pet_id"`
	UserID string `json:"user_id"`
	Name   string `json:"name"`
}
