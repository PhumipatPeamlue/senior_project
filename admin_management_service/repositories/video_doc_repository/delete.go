package video_doc_repository

import (
	"fmt"
	"net/http"
)

func (repo *VideoDocRepo) Delete(docID string) (statusCode int, err error) {
	resp, err := repo.es.Delete(repo.indexName, docID)
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
