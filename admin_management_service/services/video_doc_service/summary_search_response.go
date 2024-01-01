package video_doc_service

import "admin_management_service/models"

func (s *VideoDocService) SummarySearchResponse(searchResponse models.VideoDocSearchResponse) (listVideoDocs []models.VideoDoc, total int) {
	for _, hit := range searchResponse.Hits.Hits {
		var videoDoc models.VideoDoc
		videoDoc.ID = hit.ID
		videoDoc.Title = hit.Source.Title
		videoDoc.VideoUrl = hit.Source.VideoUrl
		videoDoc.Description = hit.Source.Description
		videoDoc.CreateAt = hit.Source.CreateAt
		videoDoc.UpdateAt = hit.Source.UpdateAt
		listVideoDocs = append(listVideoDocs, videoDoc)
	}
	total = searchResponse.Hits.Total.Value

	return
}
