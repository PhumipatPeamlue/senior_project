package repositories

import (
	"database/sql"
	"document_service/internal/core"
	"document_service/internal/core/models"
	"errors"
)

type imageInfoMySQL struct {
	db *sql.DB
}

func NewImageInfoMySQL(db *sql.DB) core.ImageInfoRepository {
	return &imageInfoMySQL{
		db: db,
	}
}

func (r *imageInfoMySQL) ReadByDocID(docID string) (info models.ImageInfo, err error) {
	stmt, err := r.db.Prepare("SELECT * FROM image_info WHERE doc_id = ?")
	if err != nil {
		return
	}
	defer stmt.Close()

	if err = stmt.QueryRow(docID).Scan(&info.DocID, &info.ImageName); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = core.NewErrorNotFound(err)
		} else {
		}
	}
	return
}

func (r *imageInfoMySQL) Create(info models.ImageInfo) (err error) {
	stmt, err := r.db.Prepare("INSERT INTO image_info (doc_id, image_name) VALUES (?, ?)")
	if err != nil {
		return
	}
	defer stmt.Close()

	res, err := stmt.Exec(info.DocID, info.ImageName)
	if err != nil {
		return
	}

	num, err := res.RowsAffected()
	if err != nil {
		return
	} else if num == 0 {
		err = core.NewErrorDuplicate(err)
	}
	return
}

func (r *imageInfoMySQL) Update(info models.ImageInfo) (err error) {
	stmt, err := r.db.Prepare("UPDATE image_info SET image_name = ? WHERE doc_id = ?")
	if err != nil {
		return
	}
	defer stmt.Close()

	res, err := stmt.Exec(info.ImageName, info.DocID)
	if err != nil {
		return
	}

	num, err := res.RowsAffected()
	if err != nil {
		return
	} else if num == 0 {
		err = core.NewErrorNotFound(err)
	}
	return
}

func (r *imageInfoMySQL) Delete(docID string) (err error) {
	stmt, err := r.db.Prepare("DELETE FROM image_info WHERE doc_id = ?")
	if err != nil {
		return
	}
	defer stmt.Close()

	res, err := stmt.Exec(docID)
	if err != nil {
		return
	}

	num, err := res.RowsAffected()
	if err != nil {
		return
	} else if num == 0 {
		err = core.NewErrorNotFound(err)
	}
	return
}
