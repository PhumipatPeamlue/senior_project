package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
	"user_service/internal/core"

	"github.com/go-sql-driver/mysql"
)

type userRepositorySQL struct {
	db *sql.DB
}

// ReadByID implements core.UserRepositoryInterface.
func (u *userRepositorySQL) ReadByID(ctx context.Context, lineUserID string) (user core.User, err error) {
	var query string = `
	SELECT morning, noon, evening, before_bed, created_at, updated_at
	FROM users
	WHERE id = ?
`
	var morning, noon, evening, beforeBed, createdAt, updatedAt time.Time

	row := u.db.QueryRowContext(ctx, query, lineUserID)
	err = row.Scan(&morning, &noon, &evening, &beforeBed, &createdAt, &updatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = core.NewErrUserNotFound(fmt.Errorf("error at userRepositorySQL.ReadByID: %v", err))
			return
		}

		err = fmt.Errorf("error at userRepositorySQL.ReadByID: %v", err)
		return
	}

	ts := core.ScanUserTimeSetting(morning, noon, evening, beforeBed)
	user = core.ScanUser(lineUserID, ts, createdAt, updatedAt)
	return
}

// Create implements core.UserRepositoryInterface.
func (u *userRepositorySQL) Create(ctx context.Context, user core.User) (err error) {
	query := `
		INSERT INTO users (id, morning, noon, evening, before_bed, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`
	id := user.ID()
	ts := user.TimeSetting()
	createdAt, updatedAt := user.CreatedAt(), user.UpdatedAt()
	morning, noon, evening, beforeBed := ts.Morning(), ts.Noon(), ts.Evening(), ts.BeforeBed()

	_, err = u.db.ExecContext(ctx, query, id, morning, noon, evening, beforeBed, createdAt, updatedAt)
	if err != nil {
		var mysqlError *mysql.MySQLError
		if errors.As(err, &mysqlError) {
			switch mysqlError.Number {
			case 1062:
				err = core.NewErrUserDuplicate(fmt.Errorf("error at userRepositorySQL.Create: %v", err))
			}
			return
		}

		err = fmt.Errorf("error at userRepositorySQL.Create: %v", err)
		return
	}

	return
}

// Update implements core.UserRepositoryInterface.
func (u *userRepositorySQL) Update(ctx context.Context, user core.User) (err error) {
	query := `
		UPDATE users SET morning = ?, noon = ?, evening = ?, before_bed = ?, updated_at = ?
		WHERE id = ?
	`
	id := user.ID()
	ts := user.TimeSetting()
	updatedAt := user.UpdatedAt()
	morning, noon, evening, beforeBed := ts.Morning(), ts.Noon(), ts.Evening(), ts.BeforeBed()

	_, err = u.db.ExecContext(ctx, query, morning, noon, evening, beforeBed, updatedAt, id)
	if errors.Is(err, sql.ErrNoRows) {
		err = core.NewErrUserNotFound(fmt.Errorf("error at userRepositorySQL.Update: %v", err))
	}

	return
}

// DeleteByID implements core.UserRepositoryInterface.
func (u *userRepositorySQL) DeleteByID(ctx context.Context, lineUserID string) (err error) {
	query := "DELETE FROM users WHERE id = ?"

	_, err = u.db.ExecContext(ctx, query, lineUserID)
	if err != nil {
		err = fmt.Errorf("error at userRepositorySQL.DeleteByID: %v", err)
	}

	return
}

func NewUserRepositorySQL(db *sql.DB) core.UserRepositoryInterface {
	return &userRepositorySQL{
		db: db,
	}
}
