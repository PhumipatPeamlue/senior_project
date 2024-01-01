package image_info_repository

import (
	"admin_management_service/models"
	"database/sql"
)

type ImageInfoRepoInterface interface {
	SelectByDocID(docID string) (image models.ImageInfo, err error)
	Insert(imageInfo models.ImageInfo) (err error)
	Delete(docID string) (err error)
	Update(imageInfo models.ImageInfo) (err error)
}

type ImageInfoRepo struct {
	db *sql.DB
}

func New(db *sql.DB) *ImageInfoRepo {
	return &ImageInfoRepo{
		db: db,
	}
}
