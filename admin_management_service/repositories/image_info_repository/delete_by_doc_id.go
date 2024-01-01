package image_info_repository

func (repo *ImageInfoRepo) Delete(docID string) (err error) {
	query := "DELETE FROM image_info WHERE doc_id = ?"
	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(docID)
	return
}
