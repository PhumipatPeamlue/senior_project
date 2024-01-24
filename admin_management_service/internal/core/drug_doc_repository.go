package core

import "document_service/internal/core/models"

type DrugDocRepository interface {
	Create(doc models.DrugDoc) (err error)
	Read(id string) (doc models.DrugDoc, err error)
	Pagination(query string) (docs []models.DrugDoc, total int, err error)
	Update(doc models.DrugDoc) (err error)
	Delete(id string) (err error)
}
