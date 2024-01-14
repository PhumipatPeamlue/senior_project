package video_doc_repository

type VideoDocRepository interface {
	Get(docID string) (doc *VideoDocWithID, err error)
	MatchAll(from int, size int) (docs *[]VideoDocWithID, total int, err error)
	MatchKeyword(from int, size int, keyword string) (docs *[]VideoDocWithID, total int, err error)
	Create(doc VideoDoc) (err error)
	Update(doc *VideoDocWithID) (err error)
	Delete(docID string) (err error)
}
