package repositories

import (
	"context"
	"cronjob/internal/core"
	"cronjob/internal/core/domains"
	"database/sql"
	"errors"
	"time"
)

type petRepositorySQL struct {
	db *sql.DB
}

func (p *petRepositorySQL) ReadByID(ctx context.Context, id string) (pet domains.Pet, err error) {
	var query = `
		SELECT user_id, name, created_at, updated_at
		FROM pets
		WHERE id = ?
	`
	var userID, name string
	var createdAt, updatedAt time.Time

	row := p.db.QueryRowContext(ctx, query, id)
	err = row.Scan(&userID, &name, &createdAt, &updatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = core.NewErrPetNotFound(err)
		}

		return
	}

	pet = domains.ScanPet(id, userID, name, createdAt, updatedAt)
	return
}

func NewPetRepositorySQL(db *sql.DB) core.PetRepositoryInterface {
	return &petRepositorySQL{
		db: db,
	}
}
