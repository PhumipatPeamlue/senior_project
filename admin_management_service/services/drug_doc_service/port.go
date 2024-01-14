package drug_doc_service

import "mime/multipart"

type DrugDocService interface {
	GetDrugDoc(docID string) (getREs *GetResponse, err error)
	CreateDrugDoc(newDocReq *NewDrugDocRequest) (err error)
	CreateDrugDocWithImage(newDocReq *NewDrugDocRequest, file *multipart.File, header *multipart.FileHeader) (err error)
	SearchDrugDoc(page int, pageSize int, keyword string) (searchRes *SearchResponse, err error)
	UpdateDrugDoc(updateDocReq *UpdateDrugDocRequest) (err error)
	UpdateDrugDocWithImage(updateDocReq *UpdateDrugDocRequest, file *multipart.File, header *multipart.FileHeader) (err error)
	DeleteDrugDoc(docID string) (err error)
}
