package services

import (
	"document_service/internal/core/domains"
	"document_service/internal/core/ports"
	"fmt"
	"github.com/google/uuid"
	"mime/multipart"
	"time"
)

type videoDocService struct {
	repo                ports.VideoDocRepository
	imageStorageService ports.ImageStorageService
}

func (s *videoDocService) AddNewVideoDoc(title string, videoURL string, description string, file *multipart.File, header *multipart.FileHeader) (err error) {
	docID := uuid.New().String()
	now := time.Now()
	doc := domains.VideoDoc{
		ID:          docID,
		Title:       title,
		VideoURL:    videoURL,
		Description: description,
		CreatedAt:   &now,
		UpdatedAt:   &now,
	}
	if err = s.repo.Save(doc); err != nil {
		return
	}

	if file != nil {
		err = s.imageStorageService.SaveImage(docID, file, header)
	}
	return
}

func (s *videoDocService) GetVideoDoc(docID string) (doc domains.VideoDoc, imageURL string, err error) {
	doc, err = s.repo.FindByID(docID)
	if err != nil {
		return
	}

	imageURL, err = s.imageStorageService.GetImageURL(docID)
	return
}

func (s *videoDocService) SearchVideoDoc(page int, pageSize int, keyword string) (docs []domains.VideoDoc, total int, err error) {
	from := (page - 1) * pageSize
	var query string
	if keyword != "" {
		query = `{
			"from": %d,
			"size": %d,
			"query": {
				"multi_match": {
					"query": "%s",
					"fields": ["title", "description"]
				}
			}
		}`
		query = fmt.Sprintf(query, from, pageSize, keyword)
	} else {
		query = `{ "from": %d, "size": %d, "query": { "match_all": {} } }`
		query = fmt.Sprintf(query, from, pageSize)
	}

	docs, total, err = s.repo.Pagination(query)
	return
}

func (s *videoDocService) ChangeVideoDocInfo(docID string, title string, videoURL string, description string, deleteImage bool, file *multipart.File, header *multipart.FileHeader) (err error) {
	doc, err := s.repo.FindByID(docID)
	if err != nil {
		return
	}

	now := time.Now()
	doc.Title = title
	doc.VideoURL = videoURL
	doc.Description = description
	doc.UpdatedAt = &now
	if err = s.repo.Update(doc); err != nil {
		return
	}

	if file != nil {
		if err = s.imageStorageService.DeleteImage(docID); err != nil {
			return
		}
		err = s.imageStorageService.SaveImage(docID, file, header)
	} else if deleteImage {
		err = s.imageStorageService.DeleteImage(docID)
	}

	return
}

func (s *videoDocService) RemoveVideoDoc(docID string) (err error) {
	if err = s.repo.Delete(docID); err != nil {
		return
	}

	err = s.imageStorageService.DeleteImage(docID)
	return
}

func NewVideoDocService(repo ports.VideoDocRepository, imageStorageService ports.ImageStorageService) ports.VideoDocService {
	return &videoDocService{
		repo:                repo,
		imageStorageService: imageStorageService,
	}
}
