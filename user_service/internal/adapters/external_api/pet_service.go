package external_api

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"user_service/internal/core"
)

type petService struct {
	client *http.Client
}

// DeleteUserPet implements core.PetServiceInterface.
func (p *petService) DeleteUserPet(ctx context.Context, userID string) (err error) {
	url := fmt.Sprintf("%s/%s", os.Getenv("DELETE_ALL_USER_PET_URL"), userID)
	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return
	}

	_, err = p.client.Do(req)
	return
}

func NewPetService(client *http.Client) core.PetServiceInterface {
	return &petService{
		client: client,
	}
}
