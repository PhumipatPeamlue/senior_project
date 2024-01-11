package ports

import "admin_management_service/models"

type DrugDocRepo interface {
	Get(id string) (doc models.DrugDocWithID, err error)
	Index(doc models.DrugDoc) (id string, err error)
	SearchMatchAll(from int, size int) (docs []models.DrugDocWithID, total int, err error)
	SearchMatchKeyword(from int, size int, keyword string) (docs []models.DrugDocWithID, total int, err error)
	Update(doc models.DrugDocWithID) (err error)
	Delete(id string) (err error)
}
