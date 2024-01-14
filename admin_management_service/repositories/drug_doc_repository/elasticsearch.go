package drug_doc_repository

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

type esDrugDocRepository struct {
	es        *elasticsearch.Client
	indexName string
}

func NewEsDrugDocRepository(es *elasticsearch.Client, indexName string) DrugDocRepository {
	return &esDrugDocRepository{
		es:        es,
		indexName: indexName,
	}
}

func (r *esDrugDocRepository) handleEsApiErr(res *esapi.Response) (err error) {
	if res.IsError() {
		log.Println(res.String())
		switch res.StatusCode {
		case http.StatusNotFound:
			err = ErrDrugDocNotFound
		default:
			err = ErrInternalDrugDocRepo
		}
		return
	}
	return
}

func (r *esDrugDocRepository) search(query string) (docs *[]DrugDocWithID, total int, err error) {
	res, err := r.es.Search(
		r.es.Search.WithIndex(r.indexName),
		r.es.Search.WithBody(strings.NewReader(query)),
	)
	if err != nil {
		log.Println(err)
		err = ErrInternalDrugDocRepo
		return
	}
	if err = r.handleEsApiErr(res); err != nil {
		return
	}
	defer res.Body.Close()

	var searchResponseBody SearchResponseBody
	if err = json.NewDecoder(res.Body).Decode(&searchResponseBody); err != nil {
		log.Println(err)
		err = ErrInternalDrugDocRepo
		return
	}

	var slice []DrugDocWithID
	for _, hit := range searchResponseBody.Hits.Hits {
		doc := DrugDocWithID{
			ID:      hit.ID,
			DrugDoc: hit.Source,
		}
		slice = append(slice, doc)
	}
	docs = &slice
	total = searchResponseBody.Hits.Total.Value
	return
}

// Get implements DrugDocRepository.
func (r *esDrugDocRepository) Get(docID string) (doc *DrugDocWithID, err error) {
	res, err := r.es.Get(r.indexName, docID)
	if err != nil {
		log.Println(err)
		err = ErrInternalDrugDocRepo
		return
	}
	if err = r.handleEsApiErr(res); err != nil {
		return
	}
	defer res.Body.Close()

	var getResponseBody GetResponseBody
	if err = json.NewDecoder(res.Body).Decode(&getResponseBody); err != nil {
		log.Println(err)
		err = ErrInternalDrugDocRepo
		return
	}

	doc = &DrugDocWithID{
		ID:      getResponseBody.ID,
		DrugDoc: getResponseBody.Source,
	}
	return
}

// MatchAll implements DrugDocRepository.
func (r *esDrugDocRepository) MatchAll(from int, size int) (docs *[]DrugDocWithID, total int, err error) {
	query := `{
		"from": %d,
		"size": %d,
		"query": { "match_all": {} }
	}`
	docs, total, err = r.search(fmt.Sprintf(query, from, size))
	return
}

// MatchKeyword implements DrugDocRepository.
func (r *esDrugDocRepository) MatchKeyword(from int, size int, keyword string) (docs *[]DrugDocWithID, total int, err error) {
	query := `{
		"from": %d,
		"size": %d,
		"multi_match": {
			"query": "%s",
			"fields": ["trade_name", "drug_name", "description"]
		}
	}`
	docs, total, err = r.search(fmt.Sprintf(query, from, size, keyword))
	return
}

// Create implements DrugDocRepository.
func (r *esDrugDocRepository) Create(doc DrugDoc) (err error) {
	b, err := json.Marshal(doc)
	if err != nil {
		log.Println(err)
		err = ErrInternalDrugDocRepo
		return
	}

	res, err := r.es.Index(r.indexName, bytes.NewReader(b))
	if err != nil {
		log.Println(err)
		err = ErrInternalDrugDocRepo
		return
	}
	err = r.handleEsApiErr(res)
	return
}

// Update implements DrugDocRepository.
func (r *esDrugDocRepository) Update(doc *DrugDocWithID) (err error) {
	updateBody := UpdateBody{
		Doc: doc.DrugDoc,
	}
	b, err := json.Marshal(updateBody)
	if err != nil {
		log.Println(err)
		err = ErrInternalDrugDocRepo
		return
	}

	res, err := r.es.Update(r.indexName, doc.ID, bytes.NewReader(b))
	if err != nil {
		log.Println(err)
		err = ErrInternalDrugDocRepo
		return
	}
	err = r.handleEsApiErr(res)
	return
}

// Delete implements DrugDocRepository.
func (r *esDrugDocRepository) Delete(docID string) (err error) {
	res, err := r.es.Delete(r.indexName, docID)
	if err != nil {
		log.Println(err)
		err = ErrInternalDrugDocRepo
		return
	}
	err = r.handleEsApiErr(res)
	return
}
