package video_doc_repository

import (
	"admin_management_service/models"
	"github.com/elastic/go-elasticsearch/v8"
)

type VideoDocRepoInterface interface {
	Get(docID string) (statusCode int, res models.VideoDocGetResponse, err error)
	MatchAll(from int, size int) (statusCode int, searchResponse models.VideoDocSearchResponse, err error)
	MatchQuery(from int, size int, keyword string) (statusCode int, searchResponse models.VideoDocSearchResponse, err error)
	Index(doc models.VideoDocES) (statusCode int, docID string, err error)
	Update(docID string, updatedBody models.VideoDocUpdatedBody) (statusCode int, err error)
	Delete(docID string) (statusCode int, err error)
}

type VideoDocRepo struct {
	es        *elasticsearch.Client
	indexName string
}

func New(es *elasticsearch.Client, indexName string) *VideoDocRepo {
	return &VideoDocRepo{
		es:        es,
		indexName: indexName,
	}
}
