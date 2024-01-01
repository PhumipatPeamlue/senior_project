package config

import (
	"log"

	"github.com/elastic/go-elasticsearch/v8"
)

func ConnectES(cfg elasticsearch.Config) (es *elasticsearch.Client) {
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func CreateIndex(es *elasticsearch.Client, index string) {
	resp, err := es.Indices.Exists([]string{index})
	if err != nil {
		return
	}
	if resp.StatusCode == 200 {
		log.Printf("%s (index) is already exist\n", index)
		return
	}

	resp, err = es.Indices.Create(index)
	if err != nil {
		log.Fatal(err)
		return
	}
	if resp.StatusCode == 400 {
		log.Printf("%s (index) is already exist\n", index)
		return
	}
	if resp.IsError() {
		log.Fatal(resp.String())
	}

}
