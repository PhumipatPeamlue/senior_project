package doc_image_repository

import (
	"admin_management_service/models"
	"admin_management_service/ports"
	"database/sql"
)

type docImageRepo struct {
	db *sql.DB
}

func (r *docImageRepo) SelectByDocID(docID string) (docImage models.DocImage, err error) {
	query := "SELECT * FROM doc_image WHERE doc_id = ?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()

	row := stmt.QueryRow(docID)
	err = row.Scan(&docImage.DocID, &docImage.Name, &docImage.Path)

	return
}

func (r *docImageRepo) Insert(docImage models.DocImage) (err error) {
	query := "INSERT INTO doc_image (doc_id, name, path) VALUES (?, ?, ?)"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(docImage.DocID, docImage.Name, docImage.Path)
	return
}

func (r *docImageRepo) Update(docImage models.DocImage) (err error) {
	query := "UPDATE doc_image SET name = ?, path = ? WHERE doc_id = ?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return
	}

	_, err = stmt.Exec(docImage.Name, docImage.Path, docImage.DocID)
	return
}

func (r *docImageRepo) Delete(docID string) (err error) {
	query := "DELETE FROM doc_image WHERE doc_id = ?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(docID)
	return
}

func NewDocImageRepo(db *sql.DB) ports.DocImageRepo {
	return &docImageRepo{
		db: db,
	}
}
