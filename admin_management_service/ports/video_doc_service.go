package ports

import "admin_management_service/models"

type VideoDocService interface {
	CreateVideoDoc(doc models.VideoDoc, imageName string, imagePath string) (err error)
	GetVideoDoc(docID string) (doc models.VideoDocWithId, imageURL string, err error)
	SearchVideoDoc(from int, size int, keyword string) (docs []models.VideoDocWithId, total int, err error)
	UpdateVideoDoc(doc models.VideoDocWithId, imageName string, imagePath string) (err error)
	DeleteVideoDoc(docID string) (err error)
}
