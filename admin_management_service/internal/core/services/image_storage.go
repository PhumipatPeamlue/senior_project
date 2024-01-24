package services

import (
	"document_service/internal/core"
	"document_service/internal/core/models"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/url"
	"os"
	"path/filepath"
)

type ImageStorageService interface {
	GetURL(docID string) (imageURL string)
	Save(docID string, file *multipart.File, name string) (err error)
	Delete(docID string) (err error)
}

type imageStorageFileSystem struct {
	repo        core.ImageInfoRepository
	host        string
	port        string
	storagePath string
}

func NewImageStorageFileSystem(repo core.ImageInfoRepository, host string, port string, storagePath string) ImageStorageService {
	return &imageStorageFileSystem{
		repo:        repo,
		host:        host,
		port:        port,
		storagePath: storagePath,
	}
}

func (s *imageStorageFileSystem) GetURL(docID string) (imageURL string) {
	info, err := s.repo.ReadByDocID(docID)
	if err != nil {
		var e *core.Error
		if errors.As(err, &e); e.Code() == core.CodeErrorNotFound {
			return
		}
	}

	path := fmt.Sprintf("/%s/%s", s.storagePath, info.ImageName)
	u := &url.URL{
		Scheme: "http",
		Host:   s.host + ":" + s.port,
		Path:   path,
	}
	imageURL = u.String()
	return
}

func (s *imageStorageFileSystem) Save(docID string, file *multipart.File, name string) (err error) {
	info := models.ImageInfo{
		DocID:     docID,
		ImageName: name,
	}
	if err = s.repo.Create(info); err != nil {
		return
	}

	path := filepath.Join(s.storagePath, name)
	f, err := os.Create(path)
	if err != nil {
		return
	}
	defer f.Close()

	_, err = io.Copy(f, *file)

	return
}

func (s *imageStorageFileSystem) Delete(docID string) (err error) {
	info, err := s.repo.ReadByDocID(docID)
	if err != nil {
		var e *core.Error
		if errors.As(err, &e); e.Code() == core.CodeErrorNotFound {
			err = nil
			return
		}
	}

	if err = s.repo.Delete(info.DocID); err != nil {
		return
	}

	path := filepath.Join(s.storagePath, info.ImageName)
	err = os.Remove(path)
	return
}
