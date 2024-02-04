package ports

import (
	"document_service/internal/core/domains"
	"mime/multipart"
)

type VideoDocService interface {
	AddNewVideoDoc(title string, videoURL string, description string, file *multipart.File, header *multipart.FileHeader) (err error)
	GetVideoDoc(docID string) (doc domains.VideoDoc, imageURL string, err error)
	SearchVideoDoc(page int, pageSize int, keyword string) (docs []domains.VideoDoc, total int, err error)
	ChangeVideoDocInfo(docID string, title string, videoURL string, description string, deleteImage bool, file *multipart.File, header *multipart.FileHeader) (err error)
	RemoveVideoDoc(docID string) (err error)
}

type DrugDocService interface {
	AddNewDrugDoc(tradeName string, drugName string, description string, preparation string, caution string, file *multipart.File, header *multipart.FileHeader) (err error)
	GetDrugDoc(docID string) (doc domains.DrugDoc, imageURL string, err error)
	SearchDrugDoc(page int, pageSize int, keyword string) (docs []domains.DrugDoc, total int, err error)
	ChangeDrugDocInfo(docID string, tradeName string, drugName string, description string, preparation string, caution string, deleteImage bool, file *multipart.File, header *multipart.FileHeader) (err error)
	RemoveDrugDoc(docID string) (err error)
}

type ImageStorageService interface {
	GetImageURL(docID string) (imageURL string, err error)
	SaveImage(docID string, file *multipart.File, header *multipart.FileHeader) (err error)
	DeleteImage(docID string) (err error)
}
