package elasticsearch

import "github.com/elastic/go-elasticsearch/v8"

func Connect(cfg elasticsearch.Config) (es *elasticsearch.Client, err error) {
	es, err = elasticsearch.NewClient(cfg)
	return
}
