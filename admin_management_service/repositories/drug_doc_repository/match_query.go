package drug_doc_repository

import (
	"admin_management_service/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const (
	matchQuery = `
	{
		"from": %d,
		"size": %d,
		"query": {
		  "multi_match": {
			"query": "%s",
			"fields": ["trade_name", "drug_name", "description"]
		  }
		}
	}
	`
)

func (repo *DrugDocRepo) MatchQuery(from int, size int, keyword string) (statusCode int, searchResponse models.DrugDocSearchResponse, err error) {
	query := fmt.Sprintf(matchQuery, from, size, keyword)
	index := repo.es.Search.WithIndex(repo.indexName)
	body := repo.es.Search.WithBody(strings.NewReader(query))
	resp, err := repo.es.Search(index, body)
	if err != nil {
		statusCode = http.StatusInternalServerError
		return
	}
	if resp.IsError() {
		statusCode = resp.StatusCode
		err = fmt.Errorf("%s", resp.String())
		return
	}
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&searchResponse); err != nil {
		statusCode = http.StatusInternalServerError
		return
	}

	statusCode = resp.StatusCode
	return
}
