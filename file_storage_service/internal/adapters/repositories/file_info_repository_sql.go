package repositories

import (
	"context"
	"database/sql"
	"errors"
	"file_storage_service/internal/core"
	"file_storage_service/internal/core/domains"
	"github.com/go-sql-driver/mysql"
)

type fileInfoRepositorySQL struct {
	db *sql.DB
}

func (f *fileInfoRepositorySQL) Update(ctx context.Context, info domains.FileInfo) (err error) {
	query := "UPDATE file_info SET file_name = ? WHERE id = ?"

	_, err = f.db.ExecContext(ctx, query, info.FileName(), info.ID())
	if errors.Is(err, sql.ErrNoRows) {
		err = core.NewErrFileInfoNotFound(err)
	}

	return
}

func (f *fileInfoRepositorySQL) ReadByID(ctx context.Context, id string) (info domains.FileInfo, err error) {
	var query = "SELECT file_name FROM file_info WHERE id = ?"
	var fileName string

	row := f.db.QueryRowContext(ctx, query, id)
	err = row.Scan(&fileName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = core.NewErrFileInfoNotFound(err)
		}

		return
	}

	info = domains.ScanFileInfo(id, fileName)
	return
}

func (f *fileInfoRepositorySQL) Create(ctx context.Context, info domains.FileInfo) (err error) {
	query := "INSERT INTO file_info (id, file_name) VALUES (?, ?)"

	_, err = f.db.ExecContext(ctx, query, info.ID(), info.FileName())
	if err != nil {
		var mysqlError *mysql.MySQLError
		if errors.As(err, &mysqlError); mysqlError.Number == 1062 {
			err = core.NewErrFileInfoDuplicate(err)
		}
	}

	return
}

func (f *fileInfoRepositorySQL) Delete(ctx context.Context, id string) (err error) {
	query := "DELETE FROM file_info WHERE id = ?"
	_, err = f.db.ExecContext(ctx, query, id)
	return
}

func NewFileInfoRepositorySQL(db *sql.DB) core.FileInfoRepositoryInterface {
	return &fileInfoRepositorySQL{
		db: db,
	}
}
