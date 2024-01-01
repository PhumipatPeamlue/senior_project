package video_doc_repository

import (
	"admin_management_service/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const (
	matchAll = `
	{
		"from": %d,
		"size": %d,
		"query": {
			"match_all": {} 
		} 
	}
	`
)

func (repo *VideoDocRepo) MatchAll(from int, size int) (statusCode int, searchResponse models.VideoDocSearchResponse, err error) {
	query := fmt.Sprintf(matchAll, from, size)
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
