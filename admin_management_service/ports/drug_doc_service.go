package ports

import "admin_management_service/models"

type DrugDocService interface {
	CreateDrugDoc(doc models.DrugDoc, imageName string, imagePath string) (err error)
	GetDrugDoc(docID string) (doc models.DrugDocWithID, imageURL string, err error)
	SearchDrugDoc(from int, size int, keyword string) (docs []models.DrugDocWithID, total int, err error)
	UpdateDrugDoc(doc models.DrugDocWithID, imageName string, imagePath string) (err error)
	DeleteDrugDoc(docID string) (err error)
}
