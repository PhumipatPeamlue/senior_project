package video_doc_repository

import (
	"admin_management_service/models"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (repo *VideoDocRepo) Update(docID string, updatedBody models.VideoDocUpdatedBody) (statusCode int, err error) {
	b, err := json.Marshal(updatedBody)
	if err != nil {
		statusCode = http.StatusInternalServerError
		return
	}

	resp, err := repo.es.Update(repo.indexName, docID, bytes.NewReader(b))
	if err != nil {
		statusCode = http.StatusInternalServerError
		return
	}
	if resp.IsError() {
		statusCode = resp.StatusCode
		err = fmt.Errorf("%s", resp.String())
		return
	}

	statusCode = resp.StatusCode
	return
}
