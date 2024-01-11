package ports

import "admin_management_service/models"

type DocImageRepo interface {
	SelectByDocID(docID string) (docImage models.DocImage, err error)
	Insert(docImage models.DocImage) (err error)
	Update(docImage models.DocImage) (err error)
	Delete(docID string) (err error)
}
