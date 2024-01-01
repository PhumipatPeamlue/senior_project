package drug_doc_repository

import (
	"admin_management_service/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func (repo *DrugDocRepo) Get(docID string) (statusCode int, getResponse models.DrugDocGetResponse, err error) {
	resp, err := repo.es.Get(repo.indexName, docID)
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

	if err = json.NewDecoder(resp.Body).Decode(&getResponse); err != nil {
		statusCode = http.StatusInternalServerError
		return
	}

	statusCode = resp.StatusCode
	return
}
