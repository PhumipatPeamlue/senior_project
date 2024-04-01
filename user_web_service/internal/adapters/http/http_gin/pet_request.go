package http_gin

type addNewPetRequest struct {
	UserID string `json:"user_id" binding:"required"`
	Name   string `json:"name" binding:"required"`
}

type changePetNameRequest struct {
	PetID string `json:"pet_id" binding:"required"`
	Name  string `json:"name" binding:"required"`
}
