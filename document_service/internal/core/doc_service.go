package core

import "fmt"

type DocServiceInterface interface {
	SearchDoc(page int, pageSize int, keyword string) (docs []any, total int, err error)
}

type DocService struct {
	repository DocRepositoryInterface
}

// SearchDoc implements DocServiceInterface.
func (d *DocService) SearchDoc(page int, pageSize int, keyword string) (docs []any, total int, err error) {
	from := (page - 1) * pageSize
	var query string
	if keyword != "" {
		query = `{
			"from": %d,
			"size": %d,
			"query": {
				"multi_match": {
					"query": "%s",
					"fields": ["title", "description", "trade_name", "drug_name"]
				}
			}
		}`
		query = fmt.Sprintf(query, from, pageSize, keyword)
	} else {
		query = `{ "from": %d, "size": %d, "query": { "match_all": {} } }`
		query = fmt.Sprintf(query, from, pageSize)
	}

	docs, total, err = d.repository.Pagination(query)
	return
}

func NewDocService(repo DocRepositoryInterface) DocServiceInterface {
	return &DocService{
		repository: repo,
	}
}
