package drug_doc_repository

type DrugDocRepository interface {
	Get(docID string) (doc *DrugDocWithID, err error)
	MatchAll(from int, size int) (docs *[]DrugDocWithID, total int, err error)
	MatchKeyword(from int, size int, keyword string) (docs *[]DrugDocWithID, total int, err error)
	Create(doc DrugDoc) (err error)
	Update(doc *DrugDocWithID) (err error)
	Delete(docID string) (err error)
}
