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
	Get(id string) (res models.VideoDocGetResult, err error)
	Search(query string) (res models.VideoDocSearchResult, err error)
	Insert(doc models.VideoDoc) (err error)
	Update(id string, doc models.VideoDoc) (err error)
	Delete(id string) (err error)
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

func (v *VideoDocIndex) Get(id string) (res models.VideoDocGetResult, err error) {
	resp, err := v.es.Get(v.index, id)
	if err != nil {
		return
	}
	if resp.IsError() {
		err = fmt.Errorf("%s", resp.String())
		return
	}

	err = json.NewDecoder(resp.Body).Decode(&res)
	return
}

func (v *VideoDocIndex) Insert(doc models.VideoDoc) (err error) {
	b, err := json.Marshal(doc)
	if err != nil {
		return
	}

	resp, err := v.es.Index(v.index, bytes.NewReader(b))
	if err != nil {
		return
	}
	if resp.IsError() {
		err = fmt.Errorf("%s", resp.String())
	}
	return
}

func (v *VideoDocIndex) Search(query string) (res models.VideoDocSearchResult, err error) {
	resp, err := v.es.Search(
		v.es.Search.WithIndex(v.index),
		v.es.Search.WithBody(strings.NewReader(query)),
	)
	if err != nil {
		return
	}
	if resp.IsError() {
		err = fmt.Errorf("%s", resp.String())
		return
	}

	err = json.NewDecoder(resp.Body).Decode(&res)
	return
}

func (v *VideoDocIndex) Update(id string, doc models.VideoDoc) (err error) {
	b, err := json.Marshal(doc)
	if err != nil {
		return
	}

	resp, err := v.es.Update(v.index, id, bytes.NewReader([]byte(fmt.Sprintf(`{"doc":%s}`, b))))
	if err != nil {
		return
	}
	if resp.IsError() {
		err = fmt.Errorf("%s", resp.String())
	}
	return
}

func (v *VideoDocIndex) Delete(id string) (err error) {
	resp, err := v.es.Delete(v.index, id)
	if err != nil {
		return
	}
	if resp.IsError() {
		err = fmt.Errorf("%s", resp.String())
	}
	return
}
