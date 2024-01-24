package services

import (
	"document_service/internal/core"
	"document_service/internal/core/models"
	"fmt"
	"mime/multipart"
	"time"

	"github.com/google/uuid"
)

type AddNewVideoDocRequest struct {
	Title       string `json:"title" form:"title" binding:"required"`
	VideoURL    string `json:"video_url" form:"video_url" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
}

type ChangeVideoDocInfoRequest struct {
	ID          string `json:"id" form:"id" binding:"required"`
	Title       string `json:"title" form:"title"`
	VideoURL    string `json:"video_url" form:"video_url"`
	Description string `json:"description" form:"description"`
	DeleteImage bool   `form:"delete_image"`
}

type GetVideoDocResponse struct {
	Doc      models.VideoDoc `json:"doc"`
	ImageURL string          `json:"image_url"`
}

type SearchVideoDocResponse struct {
	Data  []models.VideoDoc `json:"data"`
	Total int               `json:"total"`
}

type VideoDocService interface {
	AddNewVideoDoc(req AddNewVideoDocRequest, file *multipart.File, header *multipart.FileHeader) (err error)
	GetVideoDoc(id string) (res GetVideoDocResponse, err error)
	SearchVideoDoc(page int, pageSize int, keyword string) (res SearchVideoDocResponse, err error)
	ChangeVideoDocInfo(req ChangeVideoDocInfoRequest, file *multipart.File, header *multipart.FileHeader) (err error)
	RemoveVideoDoc(id string) (err error)
}

type videoDocService struct {
	repo                core.VideoDocRepository
	imageStorageService ImageStorageService
}

func NewVideoDocService(repo core.VideoDocRepository, imageStorageService ImageStorageService) VideoDocService {
	return &videoDocService{
		repo:                repo,
		imageStorageService: imageStorageService,
	}
}

func (s *videoDocService) AddNewVideoDoc(req AddNewVideoDocRequest, file *multipart.File, header *multipart.FileHeader) (err error) {
	id := uuid.New().String()
	now := time.Now()
	doc := models.VideoDoc{
		ID:          id,
		Title:       req.Title,
		VideoURL:    req.VideoURL,
		Description: req.Description,
		CreatedAt:   &now,
		UpdatedAt:   &now,
	}
	if err = s.repo.Create(doc); err != nil {
		return
	}

	if file != nil {
		imageName := fmt.Sprintf("%s_%s", time.Now().String(), header.Filename)
		err = s.imageStorageService.Save(id, file, imageName)
	}
	return
}

func (s *videoDocService) GetVideoDoc(id string) (res GetVideoDocResponse, err error) {
	doc, err := s.repo.Read(id)
	if err != nil {
		return
	}

	res.Doc = doc
	res.ImageURL = s.imageStorageService.GetURL(doc.ID)
	return
}

func (s *videoDocService) SearchVideoDoc(page int, pageSize int, keyword string) (res SearchVideoDocResponse, err error) {
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

	docs, total, err := s.repo.Pagination(query)
	if err != nil {
		return
	}

	if total == 0 {
		res.Data = []models.VideoDoc{}
	} else {
		res.Data = docs
	}
	res.Total = total
	return
}

func (s *videoDocService) ChangeVideoDocInfo(req ChangeVideoDocInfoRequest, file *multipart.File, header *multipart.FileHeader) (err error) {
	docID := req.ID
	oldDoc, err := s.repo.Read(docID)
	if err != nil {
		return
	}

	now := time.Now()
	doc := models.VideoDoc{
		ID:          docID,
		Title:       req.Title,
		VideoURL:    req.VideoURL,
		Description: req.Description,
		CreatedAt:   oldDoc.CreatedAt,
		UpdatedAt:   &now,
	}
	if err = s.repo.Update(doc); err != nil {
		return
	}

	if file != nil {
		if err = s.imageStorageService.Delete(docID); err != nil {
			return
		}

		imageName := fmt.Sprintf("%s_%s", time.Now().String(), header.Filename)
		err = s.imageStorageService.Save(docID, file, imageName)
	} else if req.DeleteImage {
		err = s.imageStorageService.Delete(docID)
	}

	return
}

func (s *videoDocService) RemoveVideoDoc(id string) (err error) {
	if err = s.repo.Delete(id); err != nil {
		return
	}

	err = s.imageStorageService.Delete(id)
	return
}
