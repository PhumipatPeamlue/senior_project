package drug_doc_repository

import (
	"admin_management_service/models"

	"github.com/elastic/go-elasticsearch/v8"
)

type DrugDocRepoInterface interface {
	Get(docID string) (statusCode int, getResponse models.DrugDocGetResponse, err error)
	Index(doc models.DrugDocES) (statusCode int, docID string, err error)
	MatchAll(from int, size int) (statusCode int, searchResponse models.DrugDocSearchResponse, err error)
	MatchQuery(from int, size int, keyword string) (statusCode int, searchResponse models.DrugDocSearchResponse, err error)
	Update(docID string, updatedBody models.DrugDocUpdatedBody) (statusCode int, err error)
	Delete(docID string) (statusCode int, err error)
}

type DrugDocRepo struct {
	es        *elasticsearch.Client
	indexName string
}

func New(es *elasticsearch.Client, indexName string) *DrugDocRepo {
	return &DrugDocRepo{
		es:        es,
		indexName: indexName,
	}
}
