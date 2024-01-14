package drug_doc_service

import (
	"errors"
	"fmt"
	"log"
	"mime/multipart"
	"senior_project/admin_management_service/repositories/drug_doc_repository"
	"senior_project/admin_management_service/repositories/image_storage_repository"
	"time"
)

type service struct {
	drugDocRepo      drug_doc_repository.DrugDocRepository
	imageStorageRepo image_storage_repository.ImageStorageRepository
}

func New(drugDocRepo drug_doc_repository.DrugDocRepository, imageStorageRepo image_storage_repository.ImageStorageRepository) DrugDocService {
	return &service{
		drugDocRepo:      drugDocRepo,
		imageStorageRepo: imageStorageRepo,
	}
}

func (s *service) handleErr(err error) error {
	log.Println(err)
	if errors.Is(err, drug_doc_repository.ErrDrugDocNotFound) {
		return ErrDrugDocNotFound
	}
	return ErrInternalDrugDocService
}

// GetDrugDoc implements DrugDocService.
func (s *service) GetDrugDoc(docID string) (getRes *GetResponse, err error) {
	doc, err := s.drugDocRepo.Get(docID)
	if err != nil {
		err = s.handleErr(err)
		return
	}

	var imageURL string
	if len(doc.ImageName) != 0 {
		imageURL = fmt.Sprintf("http://localhost:8080/image/%s", doc.ImageName)
	}
	getRes = &GetResponse{
		Doc: DrugDoc{
			ID:          docID,
			TradeName:   doc.TradeName,
			DrugName:    doc.DrugName,
			Description: doc.Description,
			Preparation: doc.Preparation,
			Caution:     doc.Caution,
			CreateAt:    doc.CreateAt,
			UpdateAt:    doc.UpdateAt,
		},
		ImageURL: imageURL,
	}
	return
}

// CreateDrugDoc implements DrugDocService.
func (s *service) CreateDrugDoc(newDocReq *NewDrugDocRequest) (err error) {
	now := time.Now()
	doc := drug_doc_repository.DrugDoc{
		TradeName:   newDocReq.TradeName,
		DrugName:    newDocReq.DrugName,
		Description: newDocReq.Description,
		Preparation: newDocReq.Preparation,
		Caution:     newDocReq.Caution,
		CreateAt:    &now,
		UpdateAt:    &now,
	}
	if err = s.drugDocRepo.Create(doc); err != nil {
		err = s.handleErr(err)
	}
	return
}

// CreateDrugDocWithImage implements DrugDocService.
func (s *service) CreateDrugDocWithImage(newDocReq *NewDrugDocRequest, file *multipart.File, header *multipart.FileHeader) (err error) {
	now := time.Now()
	imageName := fmt.Sprintf("%s_%s", now.String(), header.Filename)
	doc := drug_doc_repository.DrugDoc{
		TradeName:   newDocReq.TradeName,
		DrugName:    newDocReq.DrugName,
		Description: newDocReq.Description,
		Preparation: newDocReq.Preparation,
		Caution:     newDocReq.Caution,
		ImageName:   imageName,
		CreateAt:    &now,
		UpdateAt:    &now,
	}
	if err = s.drugDocRepo.Create(doc); err != nil {
		err = s.handleErr(err)
		return
	}

	if err = s.imageStorageRepo.Save(file, imageName); err != nil {
		err = s.handleErr(err)
	}
	return
}

// SearchDrugDoc implements DrugDocService.
func (s *service) SearchDrugDoc(page int, pageSize int, keyword string) (searchRes *SearchResponse, err error) {
	var docs *[]drug_doc_repository.DrugDocWithID
	var total int
	from := (page - 1) * pageSize
	if len(keyword) == 0 {
		docs, total, err = s.drugDocRepo.MatchAll(from, pageSize)
	} else {
		docs, total, err = s.drugDocRepo.MatchKeyword(from, pageSize, keyword)
	}

	searchRes = &SearchResponse{
		Data:  []DrugDoc{},
		Total: total,
	}

	for _, doc := range *docs {
		data := DrugDoc{
			ID:          doc.ID,
			TradeName:   doc.TradeName,
			DrugName:    doc.DrugName,
			Description: doc.Description,
			Preparation: doc.Preparation,
			Caution:     doc.Caution,
			CreateAt:    doc.CreateAt,
			UpdateAt:    doc.UpdateAt,
		}
		searchRes.Data = append(searchRes.Data, data)
	}
	return
}

// UpdateDrugDoc implements DrugDocService.
func (s *service) UpdateDrugDoc(updateDocReq *UpdateDrugDocRequest) (err error) {
	var doc *drug_doc_repository.DrugDocWithID
	now := time.Now()
	doc, err = s.drugDocRepo.Get(updateDocReq.ID)
	if err != nil {
		err = s.handleErr(err)
		return
	}

	doc.TradeName = updateDocReq.TradeName
	doc.DrugName = updateDocReq.DrugName
	doc.Description = updateDocReq.Description
	doc.Preparation = updateDocReq.Preparation
	doc.Caution = updateDocReq.Caution
	doc.UpdateAt = &now
	if updateDocReq.DeleteImage && len(doc.ImageName) != 0 {
		if err = s.imageStorageRepo.Delete(doc.ImageName); err != nil {
			err = s.handleErr(err)
			return
		}
		doc.ImageName = ""
	}
	if err = s.drugDocRepo.Update(doc); err != nil {
		err = s.handleErr(err)
	}
	return
}

// UpdateDrugDocWithImage implements DrugDocService.
func (s *service) UpdateDrugDocWithImage(updateDocReq *UpdateDrugDocRequest, file *multipart.File, header *multipart.FileHeader) (err error) {
	var doc *drug_doc_repository.DrugDocWithID
	now := time.Now()
	imageName := fmt.Sprintf("%s_%s", now.String(), header.Filename)
	doc, err = s.drugDocRepo.Get(updateDocReq.ID)
	if err != nil {
		err = s.handleErr(err)
		return
	}
	if len(doc.ImageName) != 0 {
		if err = s.imageStorageRepo.Delete(doc.ImageName); err != nil {
			err = s.handleErr(err)
			return
		}
	}
	if err = s.imageStorageRepo.Save(file, imageName); err != nil {
		err = s.handleErr(err)
		return
	}

	doc.TradeName = updateDocReq.TradeName
	doc.DrugName = updateDocReq.DrugName
	doc.Description = updateDocReq.Description
	doc.Preparation = updateDocReq.Preparation
	doc.Caution = updateDocReq.Caution
	doc.UpdateAt = &now
	doc.ImageName = imageName
	if err = s.drugDocRepo.Update(doc); err != nil {
		err = s.handleErr(err)
	}
	return
}

// DeleteDrugDoc implements DrugDocService.
func (s *service) DeleteDrugDoc(docID string) (err error) {
	doc, err := s.drugDocRepo.Get(docID)
	if err != nil {
		err = s.handleErr(err)
		return
	}
	if len(doc.ImageName) != 0 {
		if err = s.imageStorageRepo.Delete(doc.ImageName); err != nil {
			err = s.handleErr(err)
			return
		}
	}

	if err = s.drugDocRepo.Delete(docID); err != nil {
		err = s.handleErr(err)
	}
	return
}
