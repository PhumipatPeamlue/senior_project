package repositories

import (
	"bytes"
	"document_service/internal/core"
	"document_service/internal/core/domains"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"net/http"
	"strings"
	"time"
)

type DrugDocSource struct {
	TradeName   string    `json:"trade_name"`
	DrugName    string    `json:"drug_name"`
	Description string    `json:"description"`
	Preparation string    `json:"preparation"`
	Caution     string    `json:"caution"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type DrugDocGetResponse struct {
	ID     string        `json:"_id"`
	Source DrugDocSource `json:"_source"`
}

type DrugDocSearchResponse struct {
	Hits struct {
		Total struct {
			Value int `json:"value"`
		} `json:"total"`
		Hits []struct {
			ID     string        `json:"_id"`
			Source DrugDocSource `json:"_source"`
		}
	}
}

type drugDocRepositoryES struct {
	es        *elasticsearch.Client
	indexName string
}

func (r *drugDocRepositoryES) handleError(res *esapi.Response) error {
	err := errors.New(res.String())

	switch res.StatusCode {
	case http.StatusNotFound:
		err = core.NewErrDocNotFound(err)
	case http.StatusConflict:
		err = core.NewErrDocDuplicate(err)
	case http.StatusBadRequest:
		err = core.NewErrDocBadRequest(err)
	}

	return err
}

func (r *drugDocRepositoryES) Save(doc domains.DrugDoc) (err error) {
	source := DrugDocSource{
		TradeName:   doc.TradeName,
		DrugName:    doc.DrugName,
		Description: doc.Description,
		Preparation: doc.Preparation,
		Caution:     doc.Caution,
		CreatedAt:   doc.CreatedAt,
		UpdatedAt:   doc.UpdatedAt,
	}
	b, err := json.Marshal(source)
	if err != nil {
		return
	}

	res, err := r.es.Index(r.indexName, bytes.NewReader(b), r.es.Index.WithDocumentID(doc.ID))
	if err != nil {
		return
	} else if res.IsError() {
		err = r.handleError(res)
	}

	return
}

func (r *drugDocRepositoryES) FindByID(id string) (doc domains.DrugDoc, err error) {
	res, err := r.es.Get(r.indexName, id)
	if err != nil {
		return
	} else if res.IsError() {
		err = r.handleError(res)
		return
	}
	defer res.Body.Close()

	var getResponse DrugDocGetResponse
	if err = json.NewDecoder(res.Body).Decode(&getResponse); err != nil {
		return
	}

	doc.ID = getResponse.ID
	doc.TradeName = getResponse.Source.TradeName
	doc.DrugName = getResponse.Source.DrugName
	doc.Description = getResponse.Source.Description
	doc.Preparation = getResponse.Source.Preparation
	doc.Caution = getResponse.Source.Caution
	doc.CreatedAt = getResponse.Source.CreatedAt
	doc.UpdatedAt = getResponse.Source.UpdatedAt
	return
}

func (r *drugDocRepositoryES) Pagination(query string) (docs []domains.DrugDoc, total int, err error) {
	res, err := r.es.Search(r.es.Search.WithIndex(r.indexName), r.es.Search.WithBody(strings.NewReader(query)))
	if err != nil {
		return
	} else if res.IsError() {
		err = r.handleError(res)
		return
	}
	defer res.Body.Close()

	var searchResponse DrugDocSearchResponse
	if err = json.NewDecoder(res.Body).Decode(&searchResponse); err != nil {
		return
	}

	total = searchResponse.Hits.Total.Value

	for _, hit := range searchResponse.Hits.Hits {
		doc := domains.DrugDoc{
			ID:          hit.ID,
			TradeName:   hit.Source.TradeName,
			DrugName:    hit.Source.DrugName,
			Description: hit.Source.Description,
			Preparation: hit.Source.Preparation,
			Caution:     hit.Source.Caution,
			CreatedAt:   hit.Source.CreatedAt,
			UpdatedAt:   hit.Source.UpdatedAt,
		}
		docs = append(docs, doc)
	}

	if len(docs) == 0 {
		err = core.NewErrDocNotFound(errors.New("no document in drug_doc index"))
	}

	return
}

func (r *drugDocRepositoryES) Update(doc domains.DrugDoc) (err error) {
	b, err := json.Marshal(doc)
	if err != nil {
		return
	}
	body := fmt.Sprintf("{\"doc\":%s}", string(b))
	res, err := r.es.Update(r.indexName, doc.ID, strings.NewReader(body))
	if err != nil {
		return
	} else if res.IsError() {
		err = r.handleError(res)
	}
	return
}

func (r *drugDocRepositoryES) Delete(id string) (err error) {
	res, err := r.es.Delete(r.indexName, id)
	if err != nil {
		return
	} else if res.IsError() {
		err = r.handleError(res)
	}
	return
}

func NewDrugDocRepositoryEs(es *elasticsearch.Client, indexName string) core.DrugDocRepositoryInterface {
	return &drugDocRepositoryES{
		es:        es,
		indexName: indexName,
	}
}
