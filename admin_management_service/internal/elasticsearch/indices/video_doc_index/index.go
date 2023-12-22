package video_doc_index

import (
	"admin_management_service/internal/models"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"log"
	"strings"
)

type VideoDocIndexInterface interface {
	CreateIndex() (err error)
	Get(id string) (statusCode int, err error, res models.VideoDocGetResult)
	Search(query string) (statusCode int, err error, res models.VideoDocSearchResult)
	Insert(doc models.VideoDoc) (statusCode int, err error)
	Update(id string, doc models.VideoDoc) (statusCode int, err error)
	Delete(id string) (statusCode int, err error)
}

type VideoDocIndex struct {
	es    *elasticsearch.Client
	index string
}

func New(es *elasticsearch.Client, index string) *VideoDocIndex {
	return &VideoDocIndex{
		es:    es,
		index: index,
	}
}

func (v *VideoDocIndex) CreateIndex() (err error) {
	resp, err := v.es.Indices.Exists([]string{v.index})
	if err != nil {
		return
	}
	if resp.StatusCode == 200 {
		log.Printf("%s (index) is already exist\n", v.index)
		return
	}

	resp, err = v.es.Indices.Create(v.index)
	if err != nil {
		return
	}
	if resp.StatusCode == 400 {
		log.Printf("%s (index) is already exist\n", v.index)
	}
	if resp.IsError() {
		err = fmt.Errorf("%s", resp.String())
	}
	return
}

func (v *VideoDocIndex) Get(id string) (statusCode int, err error, res models.VideoDocGetResult) {
	statusCode = 200
	resp, err := v.es.Get(v.index, id)
	if err != nil {
		statusCode = 500
		return
	}
	if resp.IsError() {
		statusCode = resp.StatusCode
		err = fmt.Errorf("%s", resp.String())
		return
	}

	if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
		statusCode = 500
	}
	return
}

func (v *VideoDocIndex) Insert(doc models.VideoDoc) (statusCode int, err error) {
	statusCode = 200
	b, err := json.Marshal(doc)
	if err != nil {
		statusCode = 500
		return
	}

	resp, err := v.es.Index(v.index, bytes.NewReader(b))
	if err != nil {
		statusCode = 500
		return
	}
	if resp.IsError() {
		statusCode = resp.StatusCode
		err = fmt.Errorf("%s", resp.String())
	}
	return
}

func (v *VideoDocIndex) Search(query string) (statusCode int, err error, res models.VideoDocSearchResult) {
	statusCode = 200
	resp, err := v.es.Search(
		v.es.Search.WithIndex(v.index),
		v.es.Search.WithBody(strings.NewReader(query)),
	)
	if err != nil {
		statusCode = 500
		return
	}
	if resp.IsError() {
		statusCode = resp.StatusCode
		err = fmt.Errorf("%s", resp.String())
		return
	}

	if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
		statusCode = 500
	}
	return
}

func (v *VideoDocIndex) Update(id string, doc models.VideoDoc) (statusCode int, err error) {
	statusCode = 200
	b, err := json.Marshal(doc)
	if err != nil {
		statusCode = 500
		return
	}

	resp, err := v.es.Update(v.index, id, bytes.NewReader([]byte(fmt.Sprintf(`{"doc":%s}`, b))))
	if err != nil {
		statusCode = 500
		return
	}
	if resp.IsError() {
		statusCode = resp.StatusCode
		err = fmt.Errorf("%s", resp.String())
	}
	return
}

func (v *VideoDocIndex) Delete(id string) (statusCode int, err error) {
	statusCode = 200
	resp, err := v.es.Delete(v.index, id)
	if err != nil {
		statusCode = 500
		return
	}
	if resp.IsError() {
		statusCode = resp.StatusCode
		err = fmt.Errorf("%s", resp.String())
	}
	return
}
