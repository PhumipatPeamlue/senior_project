package repositories

import (
	"context"
	"database/sql"
	"errors"
	"github.com/go-sql-driver/mysql"
	"user_service/internal/core/domains"
	"user_service/internal/core/ports"
	"user_service/internal/core/services"
)

type userRepositoryMySQL struct {
	db *sql.DB
}

func (r *userRepositoryMySQL) FindByID(ctx context.Context, id string) (user domains.User, err error) {
	query := "SELECT id, morning, noon, evening, before_bed FROM users WHERE id = ?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()

	if err = stmt.QueryRowContext(ctx, id).Scan(&user.ID, &user.Morning, &user.Noon, &user.Evening, &user.BeforeBed); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = services.NotFoundError
		}
	}
	return
}

func (r *userRepositoryMySQL) Save(ctx context.Context, user domains.User) (err error) {
	query := "INSERT INTO users (id, morning, noon, evening, before_bed) VALUES (?, ?, ?, ?, ?)"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()

	if _, err = stmt.ExecContext(ctx, user.ID, user.Morning, user.Noon, user.Evening, user.BeforeBed); err != nil {
		var e *mysql.MySQLError
		if errors.As(err, &e); e.Number == 1062 {
			err = services.DuplicateError
		}
	}
	return
}

func (r *userRepositoryMySQL) Update(ctx context.Context, user domains.User) (err error) {
	query := "UPDATE users SET morning = ?, noon = ?, evening = ?, before_bed = ? WHERE id = ?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, user.Morning, user.Noon, user.Evening, user.BeforeBed, user.ID)
	return
}

func (r *userRepositoryMySQL) DeleteByID(ctx context.Context, id string) (err error) {
	query := "DELETE FROM users WHERE id = ?"
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

func NewUserRepositoryMySQL(db *sql.DB) ports.UserRepository {
	return &userRepositoryMySQL{
		db: db,
	}
}
