package services

import (
	"context"
	"fmt"
	"image_storage_service/internal/core/domains"
	"image_storage_service/internal/core/ports"
	"io"
	"mime/multipart"
	"net/url"
	"os"
	"path/filepath"
	"time"
)

type imageStorageService struct {
	repo        ports.ImageInfoRepository
	host        string
	port        string
	storagePath string
}

func (s *imageStorageService) GetURL(ctx context.Context, id string) (imageUrl string, err error) {
	imgInfo, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return
	}

	path := fmt.Sprintf("/%s/%s/%s", "image", s.storagePath, imgInfo.ImageName)
	u := url.URL{
		Scheme: "http",
		Host:   s.host + ":" + s.port,
		Path:   path,
	}
	imageUrl = u.String()
	return
}

func (s *imageStorageService) Save(ctx context.Context, id string, file *multipart.File, header *multipart.FileHeader) (err error) {
	name := fmt.Sprintf("%s_%s", time.Now().String(), header.Filename)
	imgInfo := domains.ImageInfo{
		ID:        id,
		ImageName: name,
	}
	if err = s.repo.Save(ctx, imgInfo); err != nil {
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

func (s *imageStorageService) Delete(ctx context.Context, id string) (err error) {
	imgInfo, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return
	}

	path := filepath.Join(s.storagePath, imgInfo.ImageName)
	if err = os.Remove(path); err != nil {
		return
	}

	err = s.repo.Delete(ctx, id)
	return
}

func NewImageStorageService(repo ports.ImageInfoRepository, host string, port string, storagePath string) ports.ImageStorageService {
	return &imageStorageService{
		repo:        repo,
		host:        host,
		port:        port,
		storagePath: storagePath,
	}
}
