package services

import (
	"document_service/internal/core"
	"document_service/internal/core/models"
	"fmt"
	"mime/multipart"
	"time"

	"github.com/google/uuid"
)

type AddNewDrugDocRequest struct {
	TradeName   string `form:"trade_name" binding:"required"`
	DrugName    string `form:"drug_name" binding:"required"`
	Description string `form:"description" binding:"required"`
	Preparation string `form:"preparation" binding:"required"`
	Caution     string `form:"caution" binding:"required"`
}

type ChangeDrugDocInfoRequest struct {
	ID          string `form:"id" binding:"required"`
	TradeName   string `form:"trade_name"`
	DrugName    string `form:"drug_name"`
	Description string `form:"description"`
	Preparation string `form:"preparation"`
	Caution     string `form:"caution"`
	DeleteImage bool   `form:"delete_image"`
}

type GetDrugDocResponse struct {
	Doc      models.DrugDoc `json:"doc"`
	ImageURL string         `json:"image_url"`
}

type SearchDrugDocResponse struct {
	Data  []models.DrugDoc `json:"data"`
	Total int              `json:"total"`
}

type DrugDocService interface {
	AddNewDrugDoc(req AddNewDrugDocRequest, file *multipart.File, header *multipart.FileHeader) (err error)
	GetDrugDoc(id string) (res GetDrugDocResponse, err error)
	SearchDrugDoc(page int, pageSize int, keyword string) (res SearchDrugDocResponse, err error)
	ChangeDrugDocInfo(req ChangeDrugDocInfoRequest, file *multipart.File, header *multipart.FileHeader) (err error)
	RemoveDrugDoc(id string) (err error)
}

type drugDocService struct {
	repo                core.DrugDocRepository
	imageStorageService ImageStorageService
}

func NewDrugDocService(repo core.DrugDocRepository, imageStorageService ImageStorageService) DrugDocService {
	return &drugDocService{
		repo:                repo,
		imageStorageService: imageStorageService,
	}
}

func (s *drugDocService) AddNewDrugDoc(req AddNewDrugDocRequest, file *multipart.File, header *multipart.FileHeader) (err error) {
	id := uuid.New().String()
	now := time.Now()
	doc := models.DrugDoc{
		ID:          id,
		TradeName:   req.TradeName,
		DrugName:    req.DrugName,
		Description: req.Description,
		Preparation: req.Preparation,
		Caution:     req.Caution,
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

func (s *drugDocService) GetDrugDoc(id string) (res GetDrugDocResponse, err error) {
	doc, err := s.repo.Read(id)
	if err != nil {
		return
	}

	res.Doc = doc
	res.ImageURL = s.imageStorageService.GetURL(doc.ID)
	return
}

func (s *drugDocService) SearchDrugDoc(page int, pageSize int, keyword string) (res SearchDrugDocResponse, err error) {
	from := (page - 1) * pageSize
	var query string
	if keyword != "" {
		query = `{
			"from": %d,
			"size": %d,
			"query": {
				"multi_match": {
					"query": "%s",
					"fields": ["trade_name", "drug_name", "description"]
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
		res.Data = []models.DrugDoc{}
	} else {
		res.Data = docs
	}
	res.Total = total
	return
}

func (s *drugDocService) ChangeDrugDocInfo(req ChangeDrugDocInfoRequest, file *multipart.File, header *multipart.FileHeader) (err error) {
	docID := req.ID
	oldDoc, err := s.repo.Read(docID)
	if err != nil {
		return
	}

	now := time.Now()
	doc := models.DrugDoc{
		ID:          docID,
		TradeName:   req.TradeName,
		DrugName:    req.DrugName,
		Description: req.Description,
		Preparation: req.Preparation,
		Caution:     req.Caution,
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

func (s *drugDocService) RemoveDrugDoc(id string) (err error) {
	if err = s.repo.Delete(id); err != nil {
		return
	}

	err = s.imageStorageService.Delete(id)
	return
}
