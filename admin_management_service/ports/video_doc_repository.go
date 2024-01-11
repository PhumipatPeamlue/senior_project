package ports

import "admin_management_service/models"

type VideoDocRepo interface {
	Get(id string) (doc models.VideoDocWithId, err error)
	Index(doc models.VideoDoc) (id string, err error)
	SearchMatchAll(from int, size int) (docs []models.VideoDocWithId, total int, err error)
	SearchMatchKeyword(from int, size int, keyword string) (docs []models.VideoDocWithId, total int, err error)
	Update(doc models.VideoDocWithId) (err error)
	Delete(id string) (err error)
}
