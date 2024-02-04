package services

import (
	"document_service/internal/core/domains"
	"document_service/internal/core/ports"
	"fmt"
	"github.com/google/uuid"
	"mime/multipart"
	"time"
)

type drugDocService struct {
	repo                ports.DrugDocRepository
	imageStorageService ports.ImageStorageService
}

func (s *drugDocService) AddNewDrugDoc(tradeName string, drugName string, description string, preparation string, caution string, file *multipart.File, header *multipart.FileHeader) (err error) {
	docID := uuid.New().String()
	now := time.Now()
	doc := domains.DrugDoc{
		ID:          docID,
		TradeName:   tradeName,
		DrugName:    drugName,
		Description: description,
		Preparation: preparation,
		Caution:     caution,
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

func (s *drugDocService) GetDrugDoc(docID string) (doc domains.DrugDoc, imageURL string, err error) {
	doc, err = s.repo.FindByID(docID)
	if err != nil {
		return
	}

	imageURL, err = s.imageStorageService.GetImageURL(docID)
	return
}

func (s *drugDocService) SearchDrugDoc(page int, pageSize int, keyword string) (docs []domains.DrugDoc, total int, err error) {
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

	docs, total, err = s.repo.Pagination(query)
	return
}

func (s *drugDocService) ChangeDrugDocInfo(docID string, tradeName string, drugName string, description string, preparation string, caution string, deleteImage bool, file *multipart.File, header *multipart.FileHeader) (err error) {
	doc, err := s.repo.FindByID(docID)
	if err != nil {
		return
	}

	now := time.Now()
	doc.TradeName = tradeName
	doc.DrugName = drugName
	doc.Description = description
	doc.Preparation = preparation
	doc.Caution = caution
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

func (s *drugDocService) RemoveDrugDoc(docID string) (err error) {
	if err = s.repo.Delete(docID); err != nil {
		return
	}

	err = s.imageStorageService.DeleteImage(docID)
	return
}

func NewDrugDocService(repo ports.DrugDocRepository, imageStorageService ports.ImageStorageService) ports.DrugDocService {
	return &drugDocService{
		repo:                repo,
		imageStorageService: imageStorageService,
	}
}
