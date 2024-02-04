package repositories

import (
	"bytes"
	"document_service/internal/core/domains"
	"document_service/internal/core/ports"
	"document_service/internal/core/services"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"strings"
	"time"
)

type VideoDocSource struct {
	Title       string     `json:"title"`
	VideoURL    string     `json:"video_url"`
	Description string     `json:"description"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

type VideoDocGetResponse struct {
	ID     string         `json:"_id"`
	Source VideoDocSource `json:"_source"`
}

type VideoDocSearchResponse struct {
	Hits struct {
		Total struct {
			Value int `json:"value"`
		} `json:"total"`
		Hits []struct {
			ID     string         `json:"_id"`
			Source VideoDocSource `json:"_source"`
		}
	}
}

type videoDocRepositoryES struct {
	es        *elasticsearch.Client
	indexName string
}

func (r *videoDocRepositoryES) Save(doc domains.VideoDoc) (err error) {
	source := VideoDocSource{
		Title:       doc.Title,
		VideoURL:    doc.VideoURL,
		Description: doc.Description,
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
		err = services.HandleEsErrorResponse(res)
	}

	return
}

func (r *videoDocRepositoryES) FindByID(id string) (doc domains.VideoDoc, err error) {
	res, err := r.es.Get(r.indexName, id)
	if err != nil {
		return
	} else if res.IsError() {
		err = services.HandleEsErrorResponse(res)
		return
	}
	defer res.Body.Close()

	var getResponse VideoDocGetResponse
	if err = json.NewDecoder(res.Body).Decode(&getResponse); err != nil {
		return
	}

	doc.ID = getResponse.ID
	doc.Title = getResponse.Source.Title
	doc.VideoURL = getResponse.Source.VideoURL
	doc.Description = getResponse.Source.Description
	doc.CreatedAt = getResponse.Source.CreatedAt
	doc.UpdatedAt = getResponse.Source.UpdatedAt
	return
}

func (r *videoDocRepositoryES) Pagination(query string) (docs []domains.VideoDoc, total int, err error) {
	res, err := r.es.Search(r.es.Search.WithIndex(r.indexName), r.es.Search.WithBody(strings.NewReader(query)))
	if err != nil {
		return
	} else if res.IsError() {
		err = services.HandleEsErrorResponse(res)
		return
	}
	defer res.Body.Close()

	var searchResponse VideoDocSearchResponse
	if err = json.NewDecoder(res.Body).Decode(&searchResponse); err != nil {
		return
	}

	// if total = searchResponse.Hits.Total.Value; total == 0 {
	// 	err = core.NewErrorNotFound(errors.New("search video documents not found"))
	// 	return
	// }
	total = searchResponse.Hits.Total.Value

	for _, hit := range searchResponse.Hits.Hits {
		doc := domains.VideoDoc{
			ID:          hit.ID,
			Title:       hit.Source.Title,
			VideoURL:    hit.Source.VideoURL,
			Description: hit.Source.Description,
			CreatedAt:   hit.Source.CreatedAt,
			UpdatedAt:   hit.Source.UpdatedAt,
		}
		docs = append(docs, doc)
	}
	return
}

func (r *videoDocRepositoryES) Update(doc domains.VideoDoc) (err error) {
	b, err := json.Marshal(doc)
	if err != nil {
		return
	}
	body := fmt.Sprintf("{\"doc\":%s}", string(b))
	res, err := r.es.Update(r.indexName, doc.ID, strings.NewReader(body))
	if err != nil {
		return
	} else if res.IsError() {
		err = services.HandleEsErrorResponse(res)
	}
	return
}

func (r *videoDocRepositoryES) Delete(id string) (err error) {
	res, err := r.es.Delete(r.indexName, id)
	if err != nil {
		return
	} else if res.IsError() {
		err = services.HandleEsErrorResponse(res)
	}
	return
}

func NewVideoDocRepositoryES(es *elasticsearch.Client, indexName string) ports.VideoDocRepository {
	return &videoDocRepositoryES{
		es:        es,
		indexName: indexName,
	}
}
