package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"user_web_service/internal/core"

	"github.com/go-sql-driver/mysql"
)

type petRepositorySQL struct {
	db *sql.DB
}

func (p *petRepositorySQL) ReadByID(ctx context.Context, id string) (pet core.IPet, err error) {
	var query string = `
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
			err = &core.ErrNotFound{OriginalError: fmt.Errorf("error at petRepositorySQL.ReadByID: %v", err)}
			return
		}

		err = fmt.Errorf("error at petRepositorySQL.ReadByID: %v", err)
		return
	}

	pet = core.ScanPet(id, userID, name, createdAt, updatedAt)
	return
}

func (p *petRepositorySQL) ReadByUserID(ctx context.Context, userID string) (pets []core.IPet, err error) {
	query := `
		SELECT id, name, created_at, updated_at
		FROM pets
		WHERE user_id = ?
	`

	rows, err := p.db.QueryContext(ctx, query, userID)
	if err != nil {
		err = fmt.Errorf("error at db.QueryContext in petRepository.ReadByUserID: %v", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		var createdAt, updatedAt time.Time
		if err = rows.Scan(&id, &name, &createdAt, &updatedAt); err != nil {
			err = fmt.Errorf("error at rows.Scan in petRepository.ReadByUserID: %v", err)
			return
		}

		pet := core.ScanPet(id, userID, name, createdAt, updatedAt)
		pets = append(pets, pet)
	}

	if len(pets) == 0 {
		err = &core.ErrNotFound{OriginalError: sql.ErrNoRows}
	}

	return
}

func (p *petRepositorySQL) Create(ctx context.Context, pet core.IPet) (err error) {
	query := `
		INSERT INTO pets (id, user_id, name, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?)
	`
	id := pet.ID()
	userID := pet.UserID()
	name := pet.Name()
	createdAt, updatedAt := pet.CreatedAt(), pet.UpdatedAt()

	_, err = p.db.ExecContext(ctx, query, id, userID, name, createdAt, updatedAt)
	if err != nil {
		var e *mysql.MySQLError
		if errors.As(err, &e); e.Number == 1062 {
			err = &core.ErrDuplicate{OriginalError: fmt.Errorf("error at petRepositorySQL.Create: %v", err)}
			return
		}

		err = fmt.Errorf("error at petRepositorySQL.Create: %v", err)
	}

	return
}

func (p *petRepositorySQL) Update(ctx context.Context, pet core.IPet) (err error) {
	query := `
		UPDATE pets SET name = ?, updated_at = ?
		WHERE id = ?
	`
	id := pet.ID()
	name := pet.Name()
	updatedAt := pet.UpdatedAt()

	_, err = p.db.ExecContext(ctx, query, name, updatedAt, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = &core.ErrNotFound{OriginalError: fmt.Errorf("error at petRepositorySQL.Update: %v", err)}
			return
		}

		err = fmt.Errorf("error at petRepositorySQL.Update: %v", err)
	}

	return
}

func (p *petRepositorySQL) DeleteByID(ctx context.Context, id string) (err error) {
	query := "DELETE FROM pets WHERE id = ?"
	_, err = p.db.ExecContext(ctx, query, id)
	if err != nil {
		err = fmt.Errorf("error at petRepositorySQL.DeleteByID: %v", err)
	}
	return
}

func (p *petRepositorySQL) DeleteByUserID(ctx context.Context, userID string) (err error) {
	query := "DELETE FROM pets WHERE user_id = ?"
	_, err = p.db.ExecContext(ctx, query, userID)
	if err != nil {
		err = fmt.Errorf("error at petRepositorySQL.DeleteByUserID: %v", err)
	}
	return
}

func NewPetRepositorySQL(db *sql.DB) core.IPetRepository {
	return &petRepositorySQL{
		db: db,
	}
}
