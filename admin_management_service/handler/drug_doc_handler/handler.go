package drug_doc_handler

import (
	"admin_management_service/repositories/drug_doc_repository"
	"admin_management_service/services/drug_doc_service"
)

type DrugDocHandler struct {
	DrugDocRepo    drug_doc_repository.DrugDocRepoInterface
	DrugDocService drug_doc_service.DrugDocServiceInterface
}

func New(repo drug_doc_repository.DrugDocRepoInterface, service drug_doc_service.DrugDocServiceInterface) *DrugDocHandler {
	return &DrugDocHandler{
		DrugDocRepo:    repo,
		DrugDocService: service,
	}
}
