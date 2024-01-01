package drug_doc_repository

import (
	"admin_management_service/models"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (repo *DrugDocRepo) Index(doc models.DrugDocES) (statusCode int, docID string, err error) {
	b, err := json.Marshal(doc)
	if err != nil {
		statusCode = http.StatusInternalServerError
		return
	}

	resp, err := repo.es.Index(repo.indexName, bytes.NewReader(b))
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

	var result map[string]interface{}
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		statusCode = http.StatusInternalServerError
		return
	}
	docID, ok := result["_id"].(string)
	if !ok {
		statusCode = http.StatusInternalServerError
		err = fmt.Errorf("_id was not found in response")
		return
	}

	statusCode = resp.StatusCode
	return
}
