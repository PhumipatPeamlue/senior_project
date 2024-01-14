package config

import (
	"log"
	"net/http"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
)

var (
	videoDocIndex = "video_doc"
	drugDocIndex  = "drug_doc"
)

func InitElasticsearch(cfg elasticsearch.Config) (es *elasticsearch.Client) {
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalln(err)
	}

	if alreadyExists := checkIndexExists(es); alreadyExists {
		return
	}

	createIndex(es)

	return
}

func checkIndexExists(es *elasticsearch.Client) bool {
	res, err := es.Indices.Exists([]string{videoDocIndex, drugDocIndex})
	if err != nil {
		log.Fatalln(err)
	}
	if res.StatusCode != http.StatusNotFound {
		log.Println("elasticsearch index already exists")
		return true
	}
	return false
}

func createIndex(es *elasticsearch.Client) {
	videoDocMapping := `{
		"mappings": {
			"properties": {
				"title": { "type": "text" },
				"description": { "type" : "text" },
				"video_url": { "type" : "text" },
				"image_name": { "type" : "text" },
				"create_at": { "type" : "date" },
				"update_at": { "type" : "date" }
			}
		}
	}`
	res, err := es.Indices.Create(
		videoDocIndex,
		es.Indices.Create.WithBody(strings.NewReader(videoDocMapping)),
	)
	if err != nil {
		log.Fatalln(err)
	}
	if res.IsError() {
		log.Fatalln(res.String())
	}

	drugDocMapping := `{
		"mappings": {
			"properties": {
				"trade_name": { "type": "text" },
				"drug_name": { "type" : "text" },
				"description": { "type" : "text" },
				"preparation": { "type" : "text" },
				"caution": { "type" : "text" },
				"create_at": { "type" : "date" },
				"update_at": { "type" : "date" }
			}
		}
	}`
	res, err = es.Indices.Create(
		drugDocIndex,
		es.Indices.Create.WithBody(strings.NewReader(drugDocMapping)),
	)
	if err != nil {
		log.Fatalln(err)
	}
	if res.IsError() {
		log.Fatalln(res.String())
	}
}
