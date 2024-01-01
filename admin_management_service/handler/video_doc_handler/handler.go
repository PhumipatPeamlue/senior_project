package video_doc_handler

import (
	"admin_management_service/repositories/video_doc_repository"
	"admin_management_service/services/video_doc_service"
)

type VideoDocHandler struct {
	VideoDocRepo    video_doc_repository.VideoDocRepoInterface
	VideoDocService video_doc_service.VideoDocServiceInterface
}

func New(repo video_doc_repository.VideoDocRepoInterface, service video_doc_service.VideoDocServiceInterface) *VideoDocHandler {
	return &VideoDocHandler{
		VideoDocRepo:    repo,
		VideoDocService: service,
	}
}
