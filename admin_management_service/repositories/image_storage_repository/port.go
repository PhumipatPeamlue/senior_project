package image_storage_repository

import "mime/multipart"

type ImageStorageRepository interface {
	Save(file *multipart.File, imageName string) (err error)
	Delete(imageName string) (err error)
}
