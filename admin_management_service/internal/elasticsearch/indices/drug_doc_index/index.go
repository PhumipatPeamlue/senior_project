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
	Get(id string) (res models.DrugDocGetResult, err error)
	Search(query string) (res models.DrugDocSearchResult, err error)
	Insert(doc models.DrugDoc) (err error)
	Update(id string, doc models.DrugDoc) (err error)
	Delete(id string) (err error)
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

func (d *DrugDocIndex) Get(id string) (res models.DrugDocGetResult, err error) {
	resp, err := d.es.Get(d.index, id)
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

func (d *DrugDocIndex) Insert(doc models.DrugDoc) (err error) {
	b, err := json.Marshal(doc)
	if err != nil {
		return
	}

	resp, err := d.es.Index(d.index, bytes.NewReader(b))
	if err != nil {
		return
	}
	if resp.IsError() {
		err = fmt.Errorf("%s", resp.String())
	}
	return
}

func (d *DrugDocIndex) Search(query string) (res models.DrugDocSearchResult, err error) {
	resp, err := d.es.Search(
		d.es.Search.WithIndex(d.index),
		d.es.Search.WithBody(strings.NewReader(query)),
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

func (d *DrugDocIndex) Update(id string, doc models.DrugDoc) (err error) {
	b, err := json.Marshal(doc)
	if err != nil {
		return
	}

	resp, err := d.es.Update(d.index, id, bytes.NewReader([]byte(fmt.Sprintf(`{"doc":%s}`, b))))
	if err != nil {
		return
	}
	if resp.IsError() {
		err = fmt.Errorf("%s", resp.String())
	}
	return
}

func (d *DrugDocIndex) Delete(id string) (err error) {
	resp, err := d.es.Delete(d.index, id)
	if err != nil {
		return
	}
	if resp.IsError() {
		err = fmt.Errorf("%s", resp.String())
	}
	return
}
