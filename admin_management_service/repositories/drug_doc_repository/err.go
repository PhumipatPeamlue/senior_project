package drug_doc_repository

import "errors"

var (
	ErrDrugDocNotFound     = errors.New("drug document repository: drug document not found")
	ErrInternalDrugDocRepo = errors.New("drug document repository: internal repository error")
)
