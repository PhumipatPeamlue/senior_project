package drug_doc_repository

import (
	"admin_management_service/models"
	"admin_management_service/ports"
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/update"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"net/http"
)

type drugDocRepo struct {
	es        *elasticsearch.TypedClient
	indexName string
}

func (r *drugDocRepo) Get(id string) (doc models.DrugDocWithID, err error) {
	res, err := r.es.Get(r.indexName, id).Do(context.Background())
	if err != nil {
		return
	}
	if !res.Found {
		err = types.ElasticsearchError{Status: http.StatusNotFound}
		return
	}

	if err = json.Unmarshal(res.Source_, &doc); err != nil {
		return
	}

	doc.ID = res.Id_
	return
}

func (r *drugDocRepo) Index(doc models.DrugDoc) (id string, err error) {
	res, err := r.es.Index(r.indexName).Request(doc).Do(context.Background())
	if err != nil {
		return
	}

	id = res.Id_
	return
}

func (r *drugDocRepo) SearchMatchAll(from int, size int) (docs []models.DrugDocWithID, total int, err error) {
	res, err := r.es.Search().Index(r.indexName).Request(&search.Request{
		From: &from,
		Size: &size,
		Query: &types.Query{
			MatchAll: &types.MatchAllQuery{},
		},
	}).Do(context.Background())
	if err != nil {
		return
	}

	for _, hit := range res.Hits.Hits {
		var source models.DrugDoc
		if err = json.Unmarshal(hit.Source_, &source); err != nil {
			return
		}
		doc := models.DrugDocWithID{
			ID:      hit.Id_,
			DrugDoc: source,
		}
		docs = append(docs, doc)
	}
	total = int(res.Hits.Total.Value)

	return
}

func (r *drugDocRepo) SearchMatchKeyword(from int, size int, keyword string) (docs []models.DrugDocWithID, total int, err error) {
	res, err := r.es.Search().Index(r.indexName).Request(&search.Request{
		From: &from,
		Size: &size,
		Query: &types.Query{
			MultiMatch: &types.MultiMatchQuery{
				Query:  keyword,
				Fields: []string{"trade_name", "drug_name", "description"},
			},
		},
	}).Do(context.Background())
	if err != nil {
		return
	}

	for _, hit := range res.Hits.Hits {
		var source models.DrugDoc
		if err = json.Unmarshal(hit.Source_, &source); err != nil {
			return
		}
		doc := models.DrugDocWithID{
			ID:      hit.Id_,
			DrugDoc: source,
		}
		docs = append(docs, doc)
	}
	total = int(res.Hits.Total.Value)

	return
}

func (r *drugDocRepo) Update(doc models.DrugDocWithID) (err error) {
	exists, err := r.es.Exists(r.indexName, doc.ID).IsSuccess(nil)
	if !exists {
		err = types.ElasticsearchError{Status: http.StatusNotFound}
		return
	} else if err != nil {
		return
	}

	jsonData, err := json.Marshal(doc.DrugDoc)
	if err != nil {
		return
	}

	rawMessage := json.RawMessage(jsonData)
	_, err = r.es.Update(r.indexName, doc.ID).Request(&update.Request{
		Doc: rawMessage,
	}).Do(context.Background())
	return
}

func (r *drugDocRepo) Delete(id string) (err error) {
	exists, err := r.es.Exists(r.indexName, id).IsSuccess(nil)
	if !exists {
		err = types.ElasticsearchError{Status: http.StatusNotFound}
		return
	} else if err != nil {
		return
	}
	_, err = r.es.Delete(r.indexName, id).Do(context.Background())
	return
}

func NewDrugDocRepo(es *elasticsearch.TypedClient, indexName string) ports.DrugDocRepo {
	return &drugDocRepo{
		es:        es,
		indexName: indexName,
	}
}
