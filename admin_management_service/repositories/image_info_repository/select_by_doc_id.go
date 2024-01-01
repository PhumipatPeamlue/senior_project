package image_info_repository

import "admin_management_service/models"

func (repo *ImageInfoRepo) SelectByDocID(docID string) (imageInfo models.ImageInfo, err error) {
	query := "SELECT * FROM image_info WHERE doc_id = ?"
	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()

	row := stmt.QueryRow(docID)
	err = row.Scan(&imageInfo.DocID, &imageInfo.Filename, &imageInfo.Filepath)

	return
}
