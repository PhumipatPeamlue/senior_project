package video_doc_service

import "admin_management_service/models"

type VideoDocServiceInterface interface {
	SummarySearchResponse(searchResponse models.VideoDocSearchResponse) (listVideoDocs []models.VideoDoc, total int)
}

type VideoDocService struct{}

func New() *VideoDocService {
	return &VideoDocService{}
}
