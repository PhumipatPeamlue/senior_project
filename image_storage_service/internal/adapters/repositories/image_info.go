package repositories

import (
	"context"
	"database/sql"
	"errors"
	"image_storage_service/internal/core/domains"
	"image_storage_service/internal/core/ports"
	"image_storage_service/internal/core/services"
)

type imageInfoRepositoryMySQL struct {
	db *sql.DB
}

func (r *imageInfoRepositoryMySQL) FindByID(ctx context.Context, id string) (imgInfo domains.ImageInfo, err error) {
	query := "SELECT id, image_name FROM image_info WHERE id = ?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()

	if err = stmt.QueryRowContext(ctx, id).Scan(&imgInfo.ID, &imgInfo.ImageName); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = services.NotFoundError
		}
	}
	return
}

func (r *imageInfoRepositoryMySQL) Save(ctx context.Context, imgInfo domains.ImageInfo) (err error) {
	query := "INSERT INTO image_info (id, image_name) VALUES (?, ?)"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, imgInfo.ID, imgInfo.ImageName)
	if err != nil {
		return
	}

	num, err := res.RowsAffected()
	if err != nil {
		return
	} else if num == 0 {
		err = services.DuplicateError
	}
	return
}

func (r *imageInfoRepositoryMySQL) Update(ctx context.Context, info domains.ImageInfo) (err error) {
	query := "UPDATE image_info SET image_name = ? WHERE id = ?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, info.ImageName, info.ID)
	if err != nil {
		return
	}

	num, err := res.RowsAffected()
	if err != nil {
		return
	} else if num == 0 {
		err = services.NotFoundError
	}
	return
}

func (r *imageInfoRepositoryMySQL) Delete(ctx context.Context, id string) (err error) {
	query := "DELETE FROM image_info WHERE id = ?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, id)
	if err != nil {
		return
	}

	num, err := res.RowsAffected()
	if err != nil {
		return
	} else if num == 0 {
		err = services.NotFoundError
	}
	return
}

func NewImageInfoRepositoryMySQL(db *sql.DB) ports.ImageInfoRepository {
	return &imageInfoRepositoryMySQL{
		db: db,
	}
}
