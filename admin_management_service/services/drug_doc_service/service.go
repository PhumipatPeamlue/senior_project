package drug_doc_service

import "admin_management_service/models"

type DrugDocServiceInterface interface {
	SummarySearchResponse(searchResponse models.DrugDocSearchResponse) (listDrugDocs []models.DrugDoc, total int)
}

type DrugDocService struct{}

func New() *DrugDocService {
	return &DrugDocService{}
}
