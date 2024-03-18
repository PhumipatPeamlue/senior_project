package core

import (
	"context"
	"file_storage_service/internal/core/domains"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

type LocalFileStorageServiceInterface interface {
	GetURL(ctx context.Context, id string) (url string, err error)
	Save(ctx context.Context, id string, file *multipart.File, fileName string) (err error)
	ChangeFile(ctx context.Context, id string, file *multipart.File, fileName string) (err error)
	Delete(ctx context.Context, id string) (err error)
}

type localFileStorageService struct {
	repository FileInfoRepositoryInterface
}

func (l *localFileStorageService) findByID(ctx context.Context, id string) (info domains.FileInfo, err error) {
	info, err = l.repository.ReadByID(ctx, id)
	return
}

func (l *localFileStorageService) ChangeFile(ctx context.Context, id string, file *multipart.File, fileName string) (err error) {
	info, err := l.findByID(ctx, id)
	if err != nil {
		return
	}

	err = os.Remove(filepath.Join("bucket", info.FileName()))
	if err != nil {
		return
	}

	info.ChangeFileName(fileName)
	f, err := os.Create(filepath.Join("bucket", info.FileName()))
	if err != nil {
		return
	}
	defer f.Close()

	_, err = io.Copy(f, *file)
	if err != nil {
		return
	}

	err = l.repository.Update(ctx, info)
	return
}

func (l *localFileStorageService) GetURL(ctx context.Context, id string) (url string, err error) {
	info, err := l.findByID(ctx, id)
	if err != nil {
		return
	}

	url = info.URL("bucket")
	return
}

func (l *localFileStorageService) Save(ctx context.Context, id string, file *multipart.File, fileName string) (err error) {
	fileInfo := domains.NewFileInfo(id, fileName)
	f, err := os.Create(filepath.Join("bucket", fileInfo.FileName()))
	if err != nil {
		return
	}
	defer f.Close()

	_, err = io.Copy(f, *file)
	if err != nil {
		return
	}

	err = l.repository.Create(ctx, fileInfo)
	return
}

func (l *localFileStorageService) Delete(ctx context.Context, id string) (err error) {
	fileInfo, err := l.repository.ReadByID(ctx, id)
	if err != nil {
		return
	}

	err = l.repository.Delete(ctx, id)
	if err != nil {
		return
	}

	err = os.Remove(filepath.Join("bucket", fileInfo.FileName()))
	return
}

func NewLocalFileStorageService(repository FileInfoRepositoryInterface) LocalFileStorageServiceInterface {
	return &localFileStorageService{
		repository: repository,
	}
}
