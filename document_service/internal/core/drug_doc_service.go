package core

import (
	"document_service/internal/core/domains"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type DrugDocServiceInterface interface {
	AddNewDrugDoc(tradeName string, drugName string, description string, preparation string, caution string) (docID string, err error)
	GetDrugDoc(docID string) (doc domains.DrugDoc, err error)
	SearchDrugDoc(page int, pageSize int, keyword string) (docs []domains.DrugDoc, total int, err error)
	ChangeDrugDocInfo(docID string, tradeName string, drugName string, description string, preparation string, caution string) (err error)
	RemoveDrugDoc(docID string) (err error)
}

type drugDocService struct {
	repo DrugDocRepositoryInterface
}

func (s *drugDocService) AddNewDrugDoc(tradeName string, drugName string, description string, preparation string, caution string) (docID string, err error) {
	docID = uuid.New().String()
	now := time.Now().Local()
	doc := domains.DrugDoc{
		ID:          docID,
		TradeName:   tradeName,
		DrugName:    drugName,
		Description: description,
		Preparation: preparation,
		Caution:     caution,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	err = s.repo.Save(doc)
	return
}

func (s *drugDocService) GetDrugDoc(docID string) (doc domains.DrugDoc, err error) {
	doc, err = s.repo.FindByID(docID)

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

func (s *drugDocService) ChangeDrugDocInfo(docID string, tradeName string, drugName string, description string, preparation string, caution string) (err error) {
	doc, err := s.repo.FindByID(docID)
	if err != nil {
		return
	}

	now := time.Now().Local()
	doc.TradeName = tradeName
	doc.DrugName = drugName
	doc.Description = description
	doc.Preparation = preparation
	doc.Caution = caution
	doc.UpdatedAt = now
	err = s.repo.Update(doc)

	return
}

func (s *drugDocService) RemoveDrugDoc(docID string) (err error) {
	err = s.repo.Delete(docID)
	return
}

func NewDrugDocService(repo DrugDocRepositoryInterface) DrugDocServiceInterface {
	return &drugDocService{
		repo: repo,
	}
}
