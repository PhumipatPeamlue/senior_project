package repositories

import (
	"bytes"
	"document_service/internal/core"
	"document_service/internal/core/models"
	"encoding/json"
	"errors"
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

type videoDocES struct {
	es        *elasticsearch.Client
	indexName string
}

func NewVideoDocES(es *elasticsearch.Client, indexName string) core.VideoDocRepository {
	return &videoDocES{
		es:        es,
		indexName: indexName,
	}
}

func (r *videoDocES) Create(doc models.VideoDoc) (err error) {
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
		err = core.HandleEsErrorResponse(res)
		return
	}

	return
}

func (r *videoDocES) Read(id string) (doc models.VideoDoc, err error) {
	res, err := r.es.Get(r.indexName, id)
	if err != nil {
		return
	} else if res.IsError() {
		err = core.HandleEsErrorResponse(res)
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

func (r *videoDocES) Pagination(query string) (docs []models.VideoDoc, total int, err error) {
	res, err := r.es.Search(r.es.Search.WithIndex(r.indexName), r.es.Search.WithBody(strings.NewReader(query)))
	if err != nil {
		return
	} else if res.IsError() {
		err = core.HandleEsErrorResponse(res)
		return
	}
	defer res.Body.Close()

	var searchResponse VideoDocSearchResponse
	if err = json.NewDecoder(res.Body).Decode(&searchResponse); err != nil {
		return
	}

	if total = searchResponse.Hits.Total.Value; total == 0 {
		err = core.NewErrorNotFound(errors.New("search video documents not found"))
		return
	}

	for _, hit := range searchResponse.Hits.Hits {
		doc := models.VideoDoc{
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

func (r *videoDocES) Update(doc models.VideoDoc) (err error) {
	b, err := json.Marshal(doc)
	if err != nil {
		return
	}
	body := fmt.Sprintf("{\"doc\":%s}", string(b))
	res, err := r.es.Update(r.indexName, doc.ID, strings.NewReader(body))
	if err != nil {
		return
	} else if res.IsError() {
		err = core.HandleEsErrorResponse(res)
	}
	return
}

func (r *videoDocES) Delete(id string) (err error) {
	res, err := r.es.Delete(r.indexName, id)
	if err != nil {
		return
	} else if res.IsError() {
		err = core.HandleEsErrorResponse(res)
	}
	return
}
