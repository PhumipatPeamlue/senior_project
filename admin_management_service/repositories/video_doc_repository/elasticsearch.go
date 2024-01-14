package video_doc_repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

type esVideoDocRepository struct {
	es        *elasticsearch.Client
	indexName string
}

func NewEsVideoDocRepository(es *elasticsearch.Client, indexName string) VideoDocRepository {
	return &esVideoDocRepository{
		es:        es,
		indexName: indexName,
	}
}

func (r *esVideoDocRepository) handleEsApiErr(res *esapi.Response) (err error) {
	if res.IsError() {
		log.Println(res.String())
		switch res.StatusCode {
		case http.StatusNotFound:
			err = ErrVideoDocNotFound
		default:
			err = ErrInternalVideoDocRepo
		}
		return
	}
	return
}

func (r *esVideoDocRepository) search(query string) (docs *[]VideoDocWithID, total int, err error) {
	res, err := r.es.Search(
		r.es.Search.WithIndex(r.indexName),
		r.es.Search.WithBody(strings.NewReader(query)),
	)
	if err != nil {
		log.Println(err)
		err = ErrInternalVideoDocRepo
		return
	}
	if err = r.handleEsApiErr(res); err != nil {
		return
	}
	defer res.Body.Close()

	var searchResponseBody SearchResponseBody
	if err = json.NewDecoder(res.Body).Decode(&searchResponseBody); err != nil {
		log.Println(err)
		err = ErrInternalVideoDocRepo
		return
	}

	var slice []VideoDocWithID
	for _, hit := range searchResponseBody.Hits.Hits {
		doc := VideoDocWithID{
			ID:       hit.ID,
			VideoDoc: hit.Source,
		}
		slice = append(slice, doc)
	}
	docs = &slice
	total = searchResponseBody.Hits.Total.Value
	return
}

func (r *esVideoDocRepository) Get(docID string) (doc *VideoDocWithID, err error) {
	res, err := r.es.Get(r.indexName, docID)
	if err != nil {
		log.Println(err)
		err = ErrInternalVideoDocRepo
		return
	}
	if err = r.handleEsApiErr(res); err != nil {
		return
	}
	defer res.Body.Close()

	var getResponseBody GetResponseBody
	if err = json.NewDecoder(res.Body).Decode(&getResponseBody); err != nil {
		log.Println(err)
		err = ErrInternalVideoDocRepo
		return
	}

	doc = &VideoDocWithID{
		ID:       getResponseBody.ID,
		VideoDoc: getResponseBody.Source,
	}
	return
}

func (r *esVideoDocRepository) MatchAll(from int, size int) (docs *[]VideoDocWithID, total int, err error) {
	query := `{
		"from": %d,
		"size": %d,
		"query": { "match_all": {} }
	}`
	docs, total, err = r.search(fmt.Sprintf(query, from, size))
	return
}

func (r *esVideoDocRepository) MatchKeyword(from int, size int, keyword string) (docs *[]VideoDocWithID, total int, err error) {
	query := `{
		"from": %d,
		"size": %d,
		"multi_match": {
			"query": "%s",
			"fields": ["title", "description"]
		}
	}`
	docs, total, err = r.search(fmt.Sprintf(query, from, size, keyword))
	return
}

func (r *esVideoDocRepository) Create(doc VideoDoc) (err error) {
	b, err := json.Marshal(doc)
	if err != nil {
		log.Println(err)
		err = ErrInternalVideoDocRepo
		return
	}

	res, err := r.es.Index(r.indexName, bytes.NewReader(b))
	if err != nil {
		log.Println(err)
		err = ErrInternalVideoDocRepo
		return
	}
	err = r.handleEsApiErr(res)
	return
}

func (r *esVideoDocRepository) Update(doc *VideoDocWithID) (err error) {
	updateBody := UpdateBody{
		Doc: doc.VideoDoc,
	}
	b, err := json.Marshal(updateBody)
	if err != nil {
		log.Println(err)
		err = ErrInternalVideoDocRepo
		return
	}

	res, err := r.es.Update(r.indexName, doc.ID, bytes.NewReader(b))
	if err != nil {
		log.Println(err)
		err = ErrInternalVideoDocRepo
		return
	}
	err = r.handleEsApiErr(res)
	return
}

func (r *esVideoDocRepository) Delete(docID string) (err error) {
	res, err := r.es.Delete(r.indexName, docID)
	if err != nil {
		log.Println(err)
		err = ErrInternalVideoDocRepo
		return
	}
	err = r.handleEsApiErr(res)
	return
}
