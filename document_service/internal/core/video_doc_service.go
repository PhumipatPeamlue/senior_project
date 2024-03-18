package core

import (
	"document_service/internal/core/domains"
	"fmt"
	"github.com/google/uuid"
	"time"
)

type VideoDocServiceInterface interface {
	AddNewVideoDoc(title string, videoURL string, description string) (err error)
	GetVideoDoc(docID string) (doc domains.VideoDoc, err error)
	SearchVideoDoc(page int, pageSize int, keyword string) (docs []domains.VideoDoc, total int, err error)
	ChangeVideoDocInfo(docID string, title string, videoURL string, description string) (err error)
	RemoveVideoDoc(docID string) (err error)
}

type videoDocService struct {
	repo VideoDocRepositoryInterface
}

func (s *videoDocService) AddNewVideoDoc(title string, videoURL string, description string) (err error) {
	docID := uuid.New().String()
	now := time.Now().Local()
	doc := domains.VideoDoc{
		ID:          docID,
		Title:       title,
		VideoURL:    videoURL,
		Description: description,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	err = s.repo.Save(doc)
	return
}

func (s *videoDocService) GetVideoDoc(docID string) (doc domains.VideoDoc, err error) {
	doc, err = s.repo.FindByID(docID)
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

func (s *videoDocService) ChangeVideoDocInfo(docID string, title string, videoURL string, description string) (err error) {
	doc, err := s.repo.FindByID(docID)
	if err != nil {
		return
	}

	now := time.Now().Local()
	doc.Title = title
	doc.VideoURL = videoURL
	doc.Description = description
	doc.UpdatedAt = now
	err = s.repo.Update(doc)
	return
}

func (s *videoDocService) RemoveVideoDoc(docID string) (err error) {
	err = s.repo.Delete(docID)

	return
}

func NewVideoDocService(repo VideoDocRepositoryInterface) VideoDocServiceInterface {
	return &videoDocService{
		repo: repo,
	}
}
