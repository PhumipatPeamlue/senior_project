package core

import "document_service/internal/core/models"

type VideoDocRepository interface {
	Create(doc models.VideoDoc) (err error)
	Read(id string) (doc models.VideoDoc, err error)
	Pagination(query string) (docs []models.VideoDoc, total int, err error)
	Update(doc models.VideoDoc) (err error)
	Delete(id string) (err error)
}
