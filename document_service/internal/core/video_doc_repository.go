package core

import "document_service/internal/core/domains"

type VideoDocRepositoryInterface interface {
	Save(doc domains.VideoDoc) (err error)
	FindByID(id string) (doc domains.VideoDoc, err error)
	Pagination(query string) (docs []domains.VideoDoc, total int, err error)
	Update(doc domains.VideoDoc) (err error)
	Delete(id string) (err error)
}
