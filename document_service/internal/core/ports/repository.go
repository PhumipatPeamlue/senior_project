package ports

import "document_service/internal/core/domains"

type VideoDocRepository interface {
	Save(doc domains.VideoDoc) (err error)
	FindByID(id string) (doc domains.VideoDoc, err error)
	Pagination(query string) (docs []domains.VideoDoc, total int, err error)
	Update(doc domains.VideoDoc) (err error)
	Delete(id string) (err error)
}

type DrugDocRepository interface {
	Save(doc domains.DrugDoc) (err error)
	FindByID(id string) (doc domains.DrugDoc, err error)
	Pagination(query string) (docs []domains.DrugDoc, total int, err error)
	Update(doc domains.DrugDoc) (err error)
	Delete(id string) (err error)
}
