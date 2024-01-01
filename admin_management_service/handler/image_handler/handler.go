package image_handler

import "admin_management_service/repositories/image_info_repository"

type ImageHandler struct {
	repo image_info_repository.ImageInfoRepoInterface
}

func New(repo image_info_repository.ImageInfoRepoInterface) *ImageHandler {
	return &ImageHandler{
		repo: repo,
	}
}
