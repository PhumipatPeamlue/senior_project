package ports

import "admin_management_service/models"

type DocImageService interface {
	CreateDocImage(docImage models.DocImage) (err error)
	GetDocImage(docID string) (docImage models.DocImage, err error)
	GetImageURL(docID string) (imageURL string, err error)
	UpdateDocImage(docImage models.DocImage) (err error)
	DeleteDocImage(docID string) (err error)
}
