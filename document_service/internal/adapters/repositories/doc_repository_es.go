package repositories

import (
	"document_service/internal/core"
	"document_service/internal/core/domains"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

type DocSearchResponse struct {
	Hits struct {
		Total struct {
			Value int `json:"value"`
		} `json:"total"`
		Hits []struct {
			Index  string `json:"_index"`
			ID     string `json:"_id"`
			Source any    `json:"_source"`
		}
	}
}

type DocRepositoryES struct {
	es                *elasticsearch.Client
	videoDocIndexName string
	drugDocIndexName  string
}

func (d *DocRepositoryES) handleError(res *esapi.Response) error {
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

// pagination implements core.DocRepositoryInterface.
func (d *DocRepositoryES) Pagination(query string) (docs []any, total int, err error) {
	res, err := d.es.Search(d.es.Search.WithIndex(d.videoDocIndexName, d.drugDocIndexName), d.es.Search.WithBody(strings.NewReader(query)))
	if err != nil {
		return
	} else if res.IsError() {
		err = d.handleError(res)
		return
	}
	defer res.Body.Close()

	var searchResponse DocSearchResponse
	if err = json.NewDecoder(res.Body).Decode(&searchResponse); err != nil {
		return
	}

	total = searchResponse.Hits.Total.Value

	for _, hit := range searchResponse.Hits.Hits {
		sourceMap, ok := hit.Source.(map[string]any)
		if !ok {
			err = errors.New("error type assertion at DocRepositoryES.Pagination")
			return
		}

		var createdAt, updatedAt time.Time
		createdAt, err = time.Parse("2006-01-02T15:04:05-07:00", sourceMap["created_at"].(string))
		if err != nil {
			return
		}
		updatedAt, err = time.Parse("2006-01-02T15:04:05-07:00", sourceMap["updated_at"].(string))
		if err != nil {
			return
		}

		if hit.Index == os.Getenv("VIDEO_DOC_INDEX") {
			doc := domains.VideoDoc{
				ID:          hit.ID,
				Title:       sourceMap["title"].(string),
				VideoURL:    sourceMap["video_url"].(string),
				Description: sourceMap["description"].(string),
				CreatedAt:   createdAt,
				UpdatedAt:   updatedAt,
			}
			docs = append(docs, doc)
		} else if hit.Index == os.Getenv("DRUG_DOC_INDEX") {
			doc := domains.DrugDoc{
				ID:          hit.ID,
				TradeName:   sourceMap["trade_name"].(string),
				DrugName:    sourceMap["drug_name"].(string),
				Description: sourceMap["description"].(string),
				Preparation: sourceMap["preparation"].(string),
				Caution:     sourceMap["caution"].(string),
				CreatedAt:   createdAt,
				UpdatedAt:   updatedAt,
			}
			docs = append(docs, doc)
		} else {
			continue
		}
	}

	if len(docs) == 0 {
		err = core.NewErrDocNotFound(errors.New("no document in video_doc index and drug_doc index"))
	}

	return
}

func NewDocRepositoryES(es *elasticsearch.Client, videoDocIndexName, drugDocIndexName string) core.DocRepositoryInterface {
	return &DocRepositoryES{
		es:                es,
		videoDocIndexName: videoDocIndexName,
		drugDocIndexName:  drugDocIndexName,
	}
}
