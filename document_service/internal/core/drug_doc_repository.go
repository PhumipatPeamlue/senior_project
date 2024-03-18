package core

import "document_service/internal/core/domains"

type DrugDocRepositoryInterface interface {
	Save(doc domains.DrugDoc) (err error)
	FindByID(id string) (doc domains.DrugDoc, err error)
	Pagination(query string) (docs []domains.DrugDoc, total int, err error)
	Update(doc domains.DrugDoc) (err error)
	Delete(id string) (err error)
}
