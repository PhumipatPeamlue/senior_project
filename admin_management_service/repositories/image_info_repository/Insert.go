package image_info_repository

import "admin_management_service/models"

func (repo *ImageInfoRepo) Insert(image models.ImageInfo) (err error) {
	query := "INSERT INTO image_info (doc_id, filename, filepath) VALUES (?, ?, ?)"
	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(image.DocID, image.Filename, image.Filepath)
	return
}
