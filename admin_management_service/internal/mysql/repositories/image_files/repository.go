package image_files

import (
	"admin_management_service/internal/models"
	"database/sql"
)

type ImageFilesRepoInterface interface {
	CreateTable() (err error)
	Select(docID string) (res models.ImageFiles, err error)
	Insert(data models.ImageFile) (id int64, err error)
	Delete(id int64) (err error)
}

type ImageFilesRepo struct {
	db *sql.DB
}

func New(db *sql.DB) *ImageFilesRepo {
	return &ImageFilesRepo{
		db: db,
	}
}

func (i *ImageFilesRepo) CreateTable() (err error) {
	query := `
	CREATE TABLE IF NOT EXISTS image_files (
  		id INT AUTO_INCREMENT PRIMARY KEY,
    	doc_id VARCHAR(255),
    	filename VARCHAR(255),
  		file_path VARCHAR(255)
	)
	`
	_, err = i.db.Exec(query)
	return
}

func (i *ImageFilesRepo) Select(docID string) (res models.ImageFiles, err error) {
	query := "SELECT * FROM image_files WHERE doc_id = ?"
	rows, err := i.db.Query(query, docID)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var temp models.ImageFile
		err = rows.Scan(&temp.ID, &temp.DocID, &temp.Filename, &temp.FilePath)
		if err != nil {
			return
		}
		res.Data = append(res.Data, temp)
	}

	return
}

func (i *ImageFilesRepo) Insert(data models.ImageFile) (id int64, err error) {
	query := "INSERT INTO image_files(doc_id, filename, file_path) VALUES(?, ?, ?)"
	stmt, err := i.db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()

	res, err := stmt.Exec(data.DocID, data.Filename, data.FilePath)
	if err != nil {
		return
	}

	id, err = res.LastInsertId()
	return
}

func (i *ImageFilesRepo) Delete(id int64) (err error) {
	query := "DELETE FROM image_files WHERE id = ?"
	_, err = i.db.Exec(query, id)
	return
}
