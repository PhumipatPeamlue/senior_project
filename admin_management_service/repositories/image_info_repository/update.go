package image_info_repository

import "admin_management_service/models"

func (repo *ImageInfoRepo) Update(imageInfo models.ImageInfo) (err error) {
	query := "UPDATE image_info SET filename = ?, filepath = ? WHERE doc_id = ?"
	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return
	}

	_, err = stmt.Exec(imageInfo.Filename, imageInfo.Filepath, imageInfo.DocID)
	return
}
