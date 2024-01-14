package drug_doc_service

import "errors"

var (
	ErrDrugDocNotFound        = errors.New("drug document service: drug document not found")
	ErrInternalDrugDocService = errors.New("drug document service: internal service error")
)
