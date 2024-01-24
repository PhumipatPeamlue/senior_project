package core

import "document_service/internal/core/models"

type ImageInfoRepository interface {
	ReadByDocID(docID string) (info models.ImageInfo, err error)
	Create(info models.ImageInfo) (err error)
	Update(info models.ImageInfo) (err error)
	Delete(docID string) (err error)
}
