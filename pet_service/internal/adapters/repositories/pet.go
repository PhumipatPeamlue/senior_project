package repositories

import (
	"context"
	"database/sql"
	"errors"
	"github.com/go-sql-driver/mysql"
	"pet_service/internal/core/domains"
	"pet_service/internal/core/ports"
	"pet_service/internal/core/services"
)

type petRepositoryMySQL struct {
	db *sql.DB
}

func (r *petRepositoryMySQL) FindByID(ctx context.Context, id string) (pet domains.Pet, err error) {
	query := "SELECT id, user_id, name FROM pets WHERE id = ?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()

	if err = stmt.QueryRowContext(ctx, id).Scan(&pet.ID, &pet.UserID, &pet.Name); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = services.NotFoundError
		}
	}

	return
}

func (r *petRepositoryMySQL) FindByUserID(ctx context.Context, userID string) (pets []domains.Pet, err error) {
	query := "SELECT id, user_id, name FROM pets WHERE user_id = ?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, userID)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var pet domains.Pet
		if err = rows.Scan(&pet.ID, &pet.UserID, &pet.Name); err != nil {
			return
		}
		pets = append(pets, pet)
	}

	if len(pets) == 0 {
		err = services.NotFoundError
	}
	return
}

func (r *petRepositoryMySQL) Save(ctx context.Context, pet domains.Pet) (err error) {
	query := "INSERT INTO pets (id, user_id, name) VALUES (?, ?, ?)"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()

	if _, err = stmt.ExecContext(ctx, pet.ID, pet.UserID, pet.Name); err != nil {
		var e *mysql.MySQLError
		if errors.As(err, &e); e.Number == 1062 {
			err = services.DuplicateError
		}
	}
	return
}

func (r *petRepositoryMySQL) Update(ctx context.Context, pet domains.Pet) (err error) {
	query := "UPDATE pets SET name = ? WHERE id = ?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, pet.Name, pet.ID)
	return
}

func (r *petRepositoryMySQL) DeleteByID(ctx context.Context, id string) (err error) {
	query := "DELETE FROM pets WHERE id = ?"
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

func NewPetRepositoryMySQL(db *sql.DB) ports.PetRepository {
	return &petRepositoryMySQL{
		db: db,
	}
}
