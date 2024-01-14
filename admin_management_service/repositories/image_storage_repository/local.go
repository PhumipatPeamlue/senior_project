package image_storage_repository

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

type localImageStorageRepository struct {
	basePath string
}

func NewLocalImageStorageRepository(basePath string) ImageStorageRepository {
	return &localImageStorageRepository{
		basePath: basePath,
	}
}

func (r *localImageStorageRepository) Save(file *multipart.File, imageName string) (err error) {
	filePath := filepath.Join(r.basePath, imageName)
	f, err := os.Create(filePath)
	if err != nil {
		err = ErrInternalImageStorageRepository
		return
	}
	defer f.Close()

	_, err = io.Copy(f, *file)
	if err != nil {
		err = ErrInternalImageStorageRepository
	}
	return
}

func (r *localImageStorageRepository) Delete(imageName string) (err error) {
	filePath := filepath.Join(r.basePath, imageName)
	if err = os.Remove(filePath); err != nil {
		err = ErrInternalImageStorageRepository
	}
	return
}
