package infrastructures

import (
	"github.com/elastic/go-elasticsearch/v8"
	"log"
)

func ConnectES(esURL string) *elasticsearch.Client {
	cfg := elasticsearch.Config{
		Addresses: []string{esURL},
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatal(err)
	}

	_, err = es.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return es
}
