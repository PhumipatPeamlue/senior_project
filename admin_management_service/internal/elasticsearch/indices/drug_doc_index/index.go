package drug_doc_index

import (
	"admin_management_service/internal/models"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"log"
	"strings"
)

type DrugDocIndexInterface interface {
	CreateIndex() (err error)
	Get(id string) (statusCode int, err error, res models.DrugDocGetResult)
	Search(query string) (statusCode int, err error, res models.DrugDocSearchResult)
	Insert(doc models.DrugDoc) (statusCode int, err error)
	Update(id string, doc models.DrugDoc) (statusCode int, err error)
	Delete(id string) (statusCode int, err error)
}

type DrugDocIndex struct {
	es    *elasticsearch.Client
	index string
}

func New(es *elasticsearch.Client, index string) *DrugDocIndex {
	return &DrugDocIndex{
		es:    es,
		index: index,
	}
}

func (d *DrugDocIndex) CreateIndex() (err error) {
	resp, err := d.es.Indices.Exists([]string{d.index})
	if err != nil {
		return
	}
	if resp.StatusCode == 200 {
		log.Printf("%s (index) is already exist\n", d.index)
		return
	}

	resp, err = d.es.Indices.Create(d.index)
	if err != nil {
		return
	}
	if resp.StatusCode == 400 {
		log.Printf("%s (index) is already exist\n", d.index)
	}
	if resp.IsError() {
		err = fmt.Errorf("%s", resp.String())
	}
	return
}

func (d *DrugDocIndex) Get(id string) (statusCode int, err error, res models.DrugDocGetResult) {
	statusCode = 200
	resp, err := d.es.Get(d.index, id)
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

func (d *DrugDocIndex) Insert(doc models.DrugDoc) (statusCode int, err error) {
	statusCode = 200
	b, err := json.Marshal(doc)
	if err != nil {
		statusCode = 500
		return
	}

	resp, err := d.es.Index(d.index, bytes.NewReader(b))
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

func (d *DrugDocIndex) Search(query string) (statusCode int, err error, res models.DrugDocSearchResult) {
	statusCode = 200
	resp, err := d.es.Search(
		d.es.Search.WithIndex(d.index),
		d.es.Search.WithBody(strings.NewReader(query)),
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

func (d *DrugDocIndex) Update(id string, doc models.DrugDoc) (statusCode int, err error) {
	statusCode = 200
	b, err := json.Marshal(doc)
	if err != nil {
		statusCode = 500
		return
	}

	resp, err := d.es.Update(d.index, id, bytes.NewReader([]byte(fmt.Sprintf(`{"doc":%s}`, b))))
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

func (d *DrugDocIndex) Delete(id string) (statusCode int, err error) {
	statusCode = 200
	resp, err := d.es.Delete(d.index, id)
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
