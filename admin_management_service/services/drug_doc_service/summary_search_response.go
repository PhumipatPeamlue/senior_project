package drug_doc_service

import "admin_management_service/models"

func (s *DrugDocService) SummarySearchResponse(searchResponse models.DrugDocSearchResponse) (listDrugDocs []models.DrugDoc, total int) {
	for _, hit := range searchResponse.Hits.Hits {
		var drugDoc models.DrugDoc
		drugDoc.ID = hit.ID
		drugDoc.TradeName = hit.Source.TradeName
		drugDoc.DrugName = hit.Source.DrugName
		drugDoc.Description = hit.Source.Description
		drugDoc.Preparation = hit.Source.Preparation
		drugDoc.Caution = hit.Source.Caution
		drugDoc.CreateAt = hit.Source.CreateAt
		drugDoc.UpdateAt = hit.Source.UpdateAt
		listDrugDocs = append(listDrugDocs, drugDoc)
	}
	total = searchResponse.Hits.Total.Value

	return
}
